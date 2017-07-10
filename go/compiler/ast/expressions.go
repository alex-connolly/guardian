package ast

// BinaryExpressionNode ...
type BinaryExpressionNode struct {
	Left, Right Node
	Operator    lexer.TokenType
}

// Type ...
func (n *BinaryExpressionNode) Type() NodeType { return BinaryExpression }

type UnaryExpressionNode struct {
	Operator lexer.TokenType
	Operand  Node
}

func (n *UnaryExpressionNode) Type() NodeType { return UnaryExpression }

type LiteralNode struct {
	LiteralType lexer.TokenType
}

func (n *LiteralNode) Type() NodeType { return Literal }

type CompositeLiteralNode struct {
}

func (n *CompositeLiteralNode) Type() NodeType { return CompositeLiteral }

type IndexExpressionNode struct {
	Expression Node
}

func (n *IndexExpressionNode) Type() NodeType { return IndexExpression }

type GenericExpressionNode struct {
	Expression Node
}

func (n *GenericExpressionNode) Type() NodeType { return GenericExpression }

type SliceExpressionNode struct {
	Expression Node
	Low, High  Node
	Max        Node
}

func (n *SliceExpressionNode) Type() NodeType { return SliceExpression }
