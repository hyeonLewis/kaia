// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package misc

import (
	"errors"
	"math/big"
	"strings"

	"github.com/klaytn/klaytn"
	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = klaytn.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// CypressCreditMetaData contains all meta data concerning the CypressCredit contract.
var CypressCreditMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"getPhoto\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNames\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"cd838f0f": "getNames()",
		"b2ec005e": "getPhoto()",
	},
}

// CypressCreditABI is the input ABI used to generate the binding from.
// Deprecated: Use CypressCreditMetaData.ABI instead.
var CypressCreditABI = CypressCreditMetaData.ABI

// CypressCreditBinRuntime is the compiled bytecode used for adding genesis block without deploying code.

// CypressCreditFuncSigs maps the 4-byte function signature to its string representation.
// Deprecated: Use CypressCreditMetaData.Sigs instead.
var CypressCreditFuncSigs = CypressCreditMetaData.Sigs

// CypressCreditBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CypressCreditMetaData.Bin instead.
var CypressCreditBin = CypressCreditMetaData.Bin

// DeployCypressCredit deploys a new Klaytn contract, binding an instance of CypressCredit to it.
func DeployCypressCredit(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CypressCredit, error) {
	parsed, err := CypressCreditMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CypressCreditBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CypressCredit{CypressCreditCaller: CypressCreditCaller{contract: contract}, CypressCreditTransactor: CypressCreditTransactor{contract: contract}, CypressCreditFilterer: CypressCreditFilterer{contract: contract}}, nil
}

// CypressCredit is an auto generated Go binding around a Klaytn contract.
type CypressCredit struct {
	CypressCreditCaller     // Read-only binding to the contract
	CypressCreditTransactor // Write-only binding to the contract
	CypressCreditFilterer   // Log filterer for contract events
}

// CypressCreditCaller is an auto generated read-only Go binding around a Klaytn contract.
type CypressCreditCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CypressCreditTransactor is an auto generated write-only Go binding around a Klaytn contract.
type CypressCreditTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CypressCreditFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type CypressCreditFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CypressCreditSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type CypressCreditSession struct {
	Contract     *CypressCredit    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CypressCreditCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type CypressCreditCallerSession struct {
	Contract *CypressCreditCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CypressCreditTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type CypressCreditTransactorSession struct {
	Contract     *CypressCreditTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CypressCreditRaw is an auto generated low-level Go binding around a Klaytn contract.
type CypressCreditRaw struct {
	Contract *CypressCredit // Generic contract binding to access the raw methods on
}

// CypressCreditCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type CypressCreditCallerRaw struct {
	Contract *CypressCreditCaller // Generic read-only contract binding to access the raw methods on
}

// CypressCreditTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type CypressCreditTransactorRaw struct {
	Contract *CypressCreditTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCypressCredit creates a new instance of CypressCredit, bound to a specific deployed contract.
func NewCypressCredit(address common.Address, backend bind.ContractBackend) (*CypressCredit, error) {
	contract, err := bindCypressCredit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CypressCredit{CypressCreditCaller: CypressCreditCaller{contract: contract}, CypressCreditTransactor: CypressCreditTransactor{contract: contract}, CypressCreditFilterer: CypressCreditFilterer{contract: contract}}, nil
}

// NewCypressCreditCaller creates a new read-only instance of CypressCredit, bound to a specific deployed contract.
func NewCypressCreditCaller(address common.Address, caller bind.ContractCaller) (*CypressCreditCaller, error) {
	contract, err := bindCypressCredit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CypressCreditCaller{contract: contract}, nil
}

// NewCypressCreditTransactor creates a new write-only instance of CypressCredit, bound to a specific deployed contract.
func NewCypressCreditTransactor(address common.Address, transactor bind.ContractTransactor) (*CypressCreditTransactor, error) {
	contract, err := bindCypressCredit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CypressCreditTransactor{contract: contract}, nil
}

// NewCypressCreditFilterer creates a new log filterer instance of CypressCredit, bound to a specific deployed contract.
func NewCypressCreditFilterer(address common.Address, filterer bind.ContractFilterer) (*CypressCreditFilterer, error) {
	contract, err := bindCypressCredit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CypressCreditFilterer{contract: contract}, nil
}

// bindCypressCredit binds a generic wrapper to an already deployed contract.
func bindCypressCredit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CypressCreditMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CypressCredit *CypressCreditRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CypressCredit.Contract.CypressCreditCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CypressCredit *CypressCreditRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CypressCredit.Contract.CypressCreditTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CypressCredit *CypressCreditRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CypressCredit.Contract.CypressCreditTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CypressCredit *CypressCreditCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CypressCredit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CypressCredit *CypressCreditTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CypressCredit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CypressCredit *CypressCreditTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CypressCredit.Contract.contract.Transact(opts, method, params...)
}

// GetNames is a free data retrieval call binding the contract method 0xcd838f0f.
//
// Solidity: function getNames() pure returns(string)
func (_CypressCredit *CypressCreditCaller) GetNames(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CypressCredit.contract.Call(opts, &out, "getNames")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetNames is a free data retrieval call binding the contract method 0xcd838f0f.
//
// Solidity: function getNames() pure returns(string)
func (_CypressCredit *CypressCreditSession) GetNames() (string, error) {
	return _CypressCredit.Contract.GetNames(&_CypressCredit.CallOpts)
}

// GetNames is a free data retrieval call binding the contract method 0xcd838f0f.
//
// Solidity: function getNames() pure returns(string)
func (_CypressCredit *CypressCreditCallerSession) GetNames() (string, error) {
	return _CypressCredit.Contract.GetNames(&_CypressCredit.CallOpts)
}

// GetPhoto is a free data retrieval call binding the contract method 0xb2ec005e.
//
// Solidity: function getPhoto() pure returns(string)
func (_CypressCredit *CypressCreditCaller) GetPhoto(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CypressCredit.contract.Call(opts, &out, "getPhoto")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetPhoto is a free data retrieval call binding the contract method 0xb2ec005e.
//
// Solidity: function getPhoto() pure returns(string)
func (_CypressCredit *CypressCreditSession) GetPhoto() (string, error) {
	return _CypressCredit.Contract.GetPhoto(&_CypressCredit.CallOpts)
}

// GetPhoto is a free data retrieval call binding the contract method 0xb2ec005e.
//
// Solidity: function getPhoto() pure returns(string)
func (_CypressCredit *CypressCreditCallerSession) GetPhoto() (string, error) {
	return _CypressCredit.Contract.GetPhoto(&_CypressCredit.CallOpts)
}