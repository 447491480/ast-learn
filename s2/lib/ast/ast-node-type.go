package ast

type NodeType int

func (t NodeType) String() string {
	switch t {
	case Programm:
		return "Programm"
	case IntDeclaration:
		return "IntDeclaration"
	case Additive:
		return "Additive"
	case Multiplicative:
		return "Multiplicative"
	case IntLiteral:
		return "IntLiteral"
	case Identifier:
		return "Identifier"
	default:
		return "Unknown"
	}
}

const (
	Programm NodeType = iota
	IntDeclaration
	Additive
	Multiplicative
	IntLiteral
	Identifier
)
