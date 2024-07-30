package lib

type TokenType int

const (
	Identifier TokenType = iota
	IntLiteral
	GT
	GE
	Plus
	Minus
	Star
	Slash
	SemiColon
	LeftParen
	RightParen
	Assignment
	Int
)

func (t TokenType) String() string {
	switch t {
	case Identifier:
		return "Identifier"
	case IntLiteral:
		return "IntLiteral"
	case GT:
		return "GT"
	case GE:
		return "GE"
	case Plus:
		return "Plus"
	case Minus:
		return "Minus"
	case Star:
		return "Star"
	case Slash:
		return "Slash"
	case SemiColon:
		return "SemiColon"
	case LeftParen:
		return "LeftParen"
	case RightParen:
		return "RightParen"
	case Assignment:
		return "Assignment"
	case Int:
		return "Int"
	default:
		return "Unknown"
	}
}
