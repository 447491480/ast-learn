package ast

type Node interface {
	GetParent() Node
	GetChildren() []Node
	GetType() NodeType
	GetText() string
}

type SimpleASTNode struct {
	Parent   Node
	Children []Node
	NodeType NodeType
	Text     string
}

func (node *SimpleASTNode) GetParent() Node {
	return node.Parent
}

func (node *SimpleASTNode) GetChildren() []Node {
	return node.Children
}

func (node *SimpleASTNode) GetType() NodeType {
	return node.NodeType
}

func (node *SimpleASTNode) GetText() string {
	return node.Text
}

func (node *SimpleASTNode) AddChild(child *SimpleASTNode) {
	node.Children = append(node.Children, child)
	child.Parent = node
}
