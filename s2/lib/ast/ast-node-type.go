package ast

type ASTNodeType int

const (
	Programm ASTNodeType = iota
	IntDeclaration
	Additive
	Multiplicative
	IntLiteral
	Identifier
)