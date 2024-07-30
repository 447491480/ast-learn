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