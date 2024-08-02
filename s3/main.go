package main

import (
	"antlr-test/parser" // Adjust the import path according to your project structure
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"strconv"
)

type EvalVisitor struct {
	parser.BaseCalcParserVisitor
}

func (v *EvalVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch ctx := tree.(type) {
	case *parser.ExpressionContext:
		if ctx.INT() != nil {
			val, _ := strconv.Atoi(ctx.INT().GetText())
			return val
		}
		left := v.Visit(ctx.Expression(0)).(int)
		right := v.Visit(ctx.Expression(1)).(int)
		switch ctx.GetOp().GetTokenType() {
		case parser.CalcParserADD:
			return left + right
		case parser.CalcParserSUB:
			return left - right
		case parser.CalcParserMUL:
			return left * right
		case parser.CalcParserDIV:
			return left / right
		}
	}
	return 0
}

func main() {
	input := "3 + 5 * 2"
	is := antlr.NewInputStream(input)
	lexer := parser.NewCalcLexer(is)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewCalcParser(tokens)
	tree := p.Expression()

	visitor := &EvalVisitor{}
	result := visitor.Visit(tree).(int)
	fmt.Println("Result:", result)
}
