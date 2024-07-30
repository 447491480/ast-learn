package lib

type DfaState int

const (
	Initial DfaState = iota
	Id
	GTState
	GEState
	PlusState
	MinusState
	StarState
	SlashState
	SemiColonState
	LeftParenState
	RightParenState
	AssignmentState
	IntLiteralState
	IdInt1
	IdInt2
	IdInt3
)
