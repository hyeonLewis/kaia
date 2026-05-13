package consensus

// NewSequenceEvent wakes the worker for the next block-building cycle.
// RoundChange=true requests an immediate rebuild after becoming proposer
// via round change, skipping the ideal-block-time wait.
type NewSequenceEvent struct {
	RoundChange bool
}
