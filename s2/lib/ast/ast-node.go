package ast

type ASTNode interface {
	GetParent() ASTNode
	GetChildren() []ASTNode
	GetType() ASTNodeType
	GetText() string
}

type SimpleASTNode struct {
	Parent          ASTNode
	Children        []ASTNode
	NodeType        ASTNodeType
	Text            string
}

func (node *SimpleASTNode) GetParent() ASTNode {
	return node.Parent
}

func (node *SimpleASTNode) GetChildren() []ASTNode {
	return node.Children
}

func (node *SimpleASTNode) GetType() ASTNodeType {
	return node.NodeType
}

func (node *SimpleASTNode) GetText() string {
	return node.Text
}

func (node *SimpleASTNode) AddChild(child *SimpleASTNode) {
	node.Children = append(node.Children, child)
	child.Parent = node
}