package lib

import (
	"fmt"
	"play/lib/ast"
	"strconv"
)

type SimpleParser struct{}

func (parser *SimpleParser) Parse(code string) *ast.SimpleASTNode {
	lexer := &SimpleLexer{}
	tokens := lexer.Tokenize(code)
	return parser.Prog(&SimpleTokenReader{Tokens: tokens})
}

func (parser *SimpleParser) Prog(tokens *SimpleTokenReader) *ast.SimpleASTNode {
	node := &ast.SimpleASTNode{NodeType: ast.Programm, Text: "Calculator"}
	child := parser.Additive(tokens)
	if child != nil {
		node.AddChild(child)
	}
	return node
}

func (parser *SimpleParser) IntDeclare(tokens *SimpleTokenReader) *ast.SimpleASTNode {
	node := &ast.SimpleASTNode{NodeType: ast.IntDeclaration}
	token := tokens.Peek()
	if token != nil && token.Type == Int {
		tokens.Read()
		token = tokens.Peek()
		if token != nil && token.Type == Identifier {
			tokens.Read()
			node.Text = token.Text
			token = tokens.Peek()
			if token != nil && token.Type == Assignment {
				tokens.Read()
				child := parser.Additive(tokens)
				if child == nil {
					panic("invalid variable initialization, expecting an expression")
				}
				node.AddChild(child)
			}
		} else {
			panic("variable name expected")
		}
		token = tokens.Peek()
		if token != nil && token.Type == SemiColon {
			tokens.Read()
		} else {
			panic("invalid statement, expecting semicolon")
		}
	}
	return node
}

func (parser *SimpleParser) Additive(tokens *SimpleTokenReader) *ast.SimpleASTNode {
	child1 := parser.Multiplicative(tokens)
	node := child1
	token := tokens.Peek()
	if child1 != nil && token != nil {
		if token.Type == Plus || token.Type == Minus {
			tokens.Read()
			child2 := parser.Additive(tokens)
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

func (parser *SimpleParser) Multiplicative(tokens *SimpleTokenReader) *ast.SimpleASTNode {
	child1 := parser.Primary(tokens)
	node := child1
	token := tokens.Peek()
	if child1 != nil && token != nil {
		if token.Type == Star || token.Type == Slash {
			tokens.Read()
			child2 := parser.Multiplicative(tokens)
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

func (parser *SimpleParser) Primary(tokens *SimpleTokenReader) *ast.SimpleASTNode {
	node := &ast.SimpleASTNode{}
	token := tokens.Peek()
	if token != nil {
		if token.Type == IntLiteral {
			tokens.Read()
			node = &ast.SimpleASTNode{NodeType: ast.IntLiteral, Text: token.Text}
		} else if token.Type == Identifier {
			tokens.Read()
			node = &ast.SimpleASTNode{NodeType: ast.Identifier, Text: token.Text}
		} else if token.Type == LeftParen {
			tokens.Read()
			node = parser.Additive(tokens)
			if node != nil {
				token = tokens.Peek()
				if token != nil && token.Type == RightParen {
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

func (parser *SimpleParser) EvaluateNode(node *ast.SimpleASTNode, indent string) int {
	result := 0
	fmt.Println(indent+"Calculating:", node.NodeType)
	switch node.NodeType {
	case ast.Programm:
		for _, child := range node.Children {
			result = parser.EvaluateNode(child.(*ast.SimpleASTNode), indent+"\t")
		}
	case ast.Additive:
		child1 := node.Children[0].(*ast.SimpleASTNode)
		value1 := parser.EvaluateNode(child1, indent+"\t")
		child2 := node.Children[1].(*ast.SimpleASTNode)
		value2 := parser.EvaluateNode(child2, indent+"\t")
		if node.Text == "+" {
			result = value1 + value2
		} else {
			result = value1 - value2
		}
	case ast.Multiplicative:
		child1 := node.Children[0].(*ast.SimpleASTNode)
		value1 := parser.EvaluateNode(child1, indent+"\t")
		child2 := node.Children[1].(*ast.SimpleASTNode)
		value2 := parser.EvaluateNode(child2, indent+"\t")
		if node.Text == "*" {
			result = value1 * value2
		} else {
			result = value1 / value2
		}
	case ast.IntLiteral:
		result, _ = strconv.Atoi(node.Text)
	default:
		panic("unhandled default case")
	}
	fmt.Println(indent+"Result:", result)
	return result
}

func (parser *SimpleParser) DumpAST(node *ast.SimpleASTNode, indent string) {
	fmt.Println(indent + node.NodeType.String() + " " + node.Text)
	for _, child := range node.Children {
		parser.DumpAST(child.(*ast.SimpleASTNode), indent+"\t")
	}
}
