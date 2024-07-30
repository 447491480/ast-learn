package main

import (
	"fmt"
	"play/lib"
)

type SimpleCalculator struct {
	lib.SimpleParser
}

func (calculator *SimpleCalculator) evaluate(script string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	tree := calculator.Parse(script)
	calculator.DumpAST(tree, "")
	result := calculator.EvaluateNode(tree, "")
	fmt.Println("Result:", result)
}

func main() {
	calculator := &SimpleCalculator{}

	script := "int a = b+3;"
	fmt.Println("解析变量声明语句:", script)
	lexer := &lib.SimpleLexer{}
	tokens := lexer.Tokenize(script)
	node := calculator.IntDeclare(&lib.SimpleTokenReader{Tokens: tokens})
	calculator.DumpAST(node, "")

	script = "2+3*5"
	fmt.Println("\n计算:", script, "，看上去一切正常。")
	calculator.evaluate(script)

	script = "2+"
	fmt.Println("\n:", script, "，应该有语法错误。")
	calculator.evaluate(script)

	script = "2+3+4"
	fmt.Println("\n计算:", script, "，结合性出现错误。")
	calculator.evaluate(script)
}
