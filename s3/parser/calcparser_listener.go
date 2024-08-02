// Code generated from /Users/wjchang/CodeKits/github/ast-learn/s3/parser/CalcParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // CalcParser

import "github.com/antlr4-go/antlr/v4"

// CalcParserListener is a complete listener for a parse tree produced by CalcParser.
type CalcParserListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)
}
