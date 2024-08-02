// Code generated from /Users/wjchang/CodeKits/github/ast-learn/s3/parser/CalcParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // CalcParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by CalcParser.
type CalcParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CalcParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}
}
