// Code generated from /Users/wjchang/CodeKits/github/ast-learn/s3/parser/CalcParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // CalcParser

import "github.com/antlr4-go/antlr/v4"

// BaseCalcParserListener is a complete listener for a parse tree produced by CalcParser.
type BaseCalcParserListener struct{}

var _ CalcParserListener = &BaseCalcParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCalcParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCalcParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCalcParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCalcParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCalcParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCalcParserListener) ExitExpression(ctx *ExpressionContext) {}
