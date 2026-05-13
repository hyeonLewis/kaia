// Modifications Copyright 2024 The Kaia Authors

package blockchain

import (
	"context"

	"github.com/kaiachain/kaia/blockchain/state"
	"github.com/kaiachain/kaia/blockchain/types"
	"github.com/kaiachain/kaia/common"
)

type stateAtReader interface {
	StateAt(root common.Hash) (*state.StateDB, error)
}

// PrefetchBlockState warms the trieDB node cache via a disposable StateDB.
// ctx cancels on shutdown/superseded spec-exec so the goroutine doesn't leak.
func PrefetchBlockState(ctx context.Context, stateReader stateAtReader, parentRoot common.Hash, blockNumber uint64, txs []*types.Transaction, signer types.Signer) {
	if stateReader == nil || len(txs) == 0 {
		return
	}
	go func() {
		defer func() {
			if r := recover(); r != nil {
				logger.Warn("PrefetchBlockState recovered from panic", "err", r)
			}
		}()
		statedb, err := stateReader.StateAt(parentRoot)
		if err != nil {
			return
		}
		for _, tx := range txs {
			select {
			case <-ctx.Done():
				return
			default:
			}
			// ValidateSender derives the sender via signature recovery (Ethereum)
			// or AccountKey-based validation (Kaia) and reads the sender account
			// key from state, populating tx.validatedSender.
			if _, err := tx.ValidateSender(signer, statedb, blockNumber); err != nil {
				continue
			}
			from := tx.ValidatedSender()
			statedb.Exist(from)
			if tx.IsFeeDelegatedTransaction() {
				if _, err := tx.ValidateFeePayer(signer, statedb, blockNumber); err == nil {
					statedb.Exist(tx.ValidatedFeePayer())
				}
			}
			// EIP-7702: sender may carry delegated code; warm its code + storage too.
			if statedb.GetCodeHash(from) != types.EmptyCodeHash {
				statedb.GetCode(from)
				statedb.GetState(from, common.Hash{})
			}
			to := tx.To()
			if to == nil {
				continue
			}
			if statedb.GetCodeHash(*to) != types.EmptyCodeHash {
				statedb.GetCode(*to)
				statedb.GetState(*to, common.Hash{})
			}
		}
	}()
}
