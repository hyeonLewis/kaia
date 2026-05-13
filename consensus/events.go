package consensus

type NewSequenceEvent struct {
	RoundChange bool
	IsProposer  bool
}
