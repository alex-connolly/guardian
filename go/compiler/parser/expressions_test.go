package parser

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/compiler/lexer"
	"axia/guardian/go/util"
	"testing"
)

func TestParseUnaryExpression(t *testing.T) {
	p := createParser("!me")
	n := p.parseUnaryExpression()
	util.Assert(t, n.Type() == ast.UnaryExpression, "wrong node type")
	u := n.(ast.UnaryExpressionNode)
	util.Assert(t, u.Operator == lexer.TknNot, "wrong operator")
}

func TestParseBinaryExpression(t *testing.T) {
	p := createParser("3 + 5")
	n := p.parseBinaryExpression()
	util.Assert(t, n.Type() == ast.BinaryExpression, "wrong node type")
	u := n.(ast.BinaryExpressionNode)
	util.Assert(t, u.Operator == lexer.TknAdd, "wrong operator")
}

func TestParseIndexExpression(t *testing.T) {
	p := createParser("array[5]")
	n := p.parseIndexExpression()
	util.Assert(t, n.Type() == ast.IndexExpression, "wrong node type")
	u := n.(ast.IndexExpressionNode)
	util.Assert(t, u.Index.Type() == ast.Literal, "wrong index type")
	p = createParser(`array[getNumber("hi")]`)
	n = p.parseIndexExpression()
	util.Assert(t, n.Type() == ast.IndexExpression, "wrong node type")
	u = n.(ast.IndexExpressionNode)
	util.Assert(t, u.Index.Type() == ast.CallExpression, "wrong index type")
}

func TestParseSliceExpression(t *testing.T) {
	p := createParser("array[5:]")
	n := p.parseSliceExpression()
	util.Assert(t, n.Type() == ast.SliceExpression, "wrong node type")
	u := n.(ast.SliceExpressionNode)
	util.Assert(t, u.Low.Type() == ast.Literal, "wrong slice type")
	p = createParser("array[:5]")
	n = p.parseSliceExpression()
	util.Assert(t, n.Type() == ast.SliceExpression, "wrong node type")
	u = n.(ast.SliceExpressionNode)
	util.Assert(t, u.High.Type() == ast.Literal, "wrong slice type")
	p = createParser("array[1:10]")
	n = p.parseSliceExpression()
	util.Assert(t, n.Type() == ast.SliceExpression, "wrong node type")
	u = n.(ast.SliceExpressionNode)
	util.Assert(t, u.Low.Type() == ast.Literal, "wrong slice low type")
	util.Assert(t, u.High.Type() == ast.Literal, "wrong slice high type")
}
