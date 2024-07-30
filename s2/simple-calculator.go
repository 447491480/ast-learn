package main

import (
	"fmt"
	"strconv"

	lib "play/lib"
	ast "play/lib/ast"
)

type SimpleCalculator struct{}

func (calculator *SimpleCalculator) evaluate(script string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	tree := calculator.parse(script)
	calculator.dumpAST(tree, "")
	result := calculator.evaluateNode(tree, "")
	fmt.Println("Result:", result)
}

func (calculator *SimpleCalculator) parse(code string) *ast.SimpleASTNode {
	lexer := &lib.SimpleLexer{}
	tokens := lexer.Tokenize(code)
	return calculator.prog(&lib.SimpleTokenReader{Tokens: tokens})
}

func (calculator *SimpleCalculator) prog(tokens *lib.SimpleTokenReader) *ast.SimpleASTNode {
	node := &ast.SimpleASTNode{NodeType: ast.Programm, Text: "Calculator"}
	child := calculator.additive(tokens)
	if child != nil {
		node.AddChild(child)
	}
	return node
}

func (calculator *SimpleCalculator) intDeclare(tokens *lib.SimpleTokenReader) *ast.SimpleASTNode {
	node := &ast.SimpleASTNode{NodeType: ast.IntDeclaration}
	token := tokens.Peek()
	if token != nil && token.Type == lib.Int {
		tokens.Read()
		token = tokens.Peek()
		if token != nil && token.Type == lib.Identifier {
			tokens.Read()
			node.Text = token.Text
			token = tokens.Peek()
			if token != nil && token.Type == lib.Assignment {
				tokens.Read()
				child := calculator.additive(tokens)
				if child == nil {
					panic("invalid variable initialization, expecting an expression")
				}
				node.AddChild(child)
			}
		} else {
			panic("variable name expected")
		}
		token = tokens.Peek()
		if token != nil && token.Type == lib.SemiColon {
			tokens.Read()
		} else {
			panic("invalid statement, expecting semicolon")
		}
	}
	return node
}

func (calculator *SimpleCalculator) additive(tokens *lib.SimpleTokenReader) *ast.SimpleASTNode {
	child1 := calculator.multiplicative(tokens)
	node := child1
	token := tokens.Peek()
	if child1 != nil && token != nil {
		if token.Type == lib.Plus || token.Type == lib.Minus {
			tokens.Read()
			child2 := calculator.additive(tokens)
			if child2 != nil {
				node = &ast.SimpleASTNode{NodeType: ast.Additive, Text: token.Text}
				node.AddChild(child1)
				node.AddChild(child2)
			} else {
				panic("invalid additive expression, expecting the right part")
			}
		}
	}
	return node
}

func (calculator *SimpleCalculator) multiplicative(tokens *lib.SimpleTokenReader) *ast.SimpleASTNode {
	child1 := calculator.primary(tokens)
	node := child1
	token := tokens.Peek()
	if child1 != nil && token != nil {
		if token.Type == lib.Star || token.Type == lib.Slash {
			tokens.Read()
			child2 := calculator.multiplicative(tokens)
			if child2 != nil {
				node = &ast.SimpleASTNode{NodeType: ast.Multiplicative, Text: token.Text}
				node.AddChild(child1)
				node.AddChild(child2)
			} else {
				panic("invalid multiplicative expression, expecting the right part")
			}
		}
	}
	return node
}

func (calculator *SimpleCalculator) primary(tokens *lib.SimpleTokenReader) *ast.SimpleASTNode {
	node := &ast.SimpleASTNode{}
	token := tokens.Peek()
	if token != nil {
		if token.Type == lib.IntLiteral {
			tokens.Read()
			node = &ast.SimpleASTNode{NodeType: ast.IntLiteral, Text: token.Text}
		} else if token.Type == lib.Identifier {
			tokens.Read()
			node = &ast.SimpleASTNode{NodeType: ast.Identifier, Text: token.Text}
		} else if token.Type == lib.LeftParen {
			tokens.Read()
			node = calculator.additive(tokens)
			if node != nil {
				token = tokens.Peek()
				if token != nil && token.Type == lib.RightParen {
					tokens.Read()
				} else {
					panic("expecting right parenthesis")
				}
			} else {
				panic("expecting an additive expression inside parenthesis")
			}
		}
	}
	return node
}

func (calculator *SimpleCalculator) evaluateNode(node *ast.SimpleASTNode, indent string) int {
	result := 0
	fmt.Println(indent+"Calculating:", node.NodeType)
	switch node.NodeType {
	case ast.Programm:
		for _, child := range node.Children {
			result = calculator.evaluateNode(child.(*ast.SimpleASTNode), indent+"\t")
		}
	case ast.Additive:
		child1 := node.Children[0].(*ast.SimpleASTNode)
		value1 := calculator.evaluateNode(child1, indent+"\t")
		child2 := node.Children[1].(*ast.SimpleASTNode)
		value2 := calculator.evaluateNode(child2, indent+"\t")
		if node.Text == "+" {
			result = value1 + value2
		} else {
			result = value1 - value2
		}
	case ast.Multiplicative:
		child1 := node.Children[0].(*ast.SimpleASTNode)
		value1 := calculator.evaluateNode(child1, indent+"\t")
		child2 := node.Children[1].(*ast.SimpleASTNode)
		value2 := calculator.evaluateNode(child2, indent+"\t")
		if node.Text == "*" {
			result = value1 * value2
		} else {
			result = value1 / value2
		}
	case ast.IntLiteral:
		result, _ = strconv.Atoi(node.Text)
	}
	fmt.Println(indent+"Result:", result)
	return result
}

func (calculator *SimpleCalculator) dumpAST(node *ast.SimpleASTNode, indent string) {
	fmt.Println(indent + node.NodeType.String() + " " + node.Text)
	for _, child := range node.Children {
		calculator.dumpAST(child.(*ast.SimpleASTNode), indent+"\t")
	}
}

func main() {
	calculator := &SimpleCalculator{}

	script := "int a = b+3;"
	fmt.Println("解析变量声明语句:", script)
	lexer := &lib.SimpleLexer{}
	tokens := lexer.Tokenize(script)
	node := calculator.intDeclare(&lib.SimpleTokenReader{Tokens: tokens})
	calculator.dumpAST(node, "")

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
