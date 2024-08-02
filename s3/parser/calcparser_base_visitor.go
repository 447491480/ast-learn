// Code generated from /Users/wjchang/CodeKits/github/ast-learn/s3/parser/CalcParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // CalcParser

import "github.com/antlr4-go/antlr/v4"

type BaseCalcParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCalcParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}
