package parser

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/compiler/lexer"
)

func (p *parser) parseExpression() ast.Node {
	// should be checked in decreasing order of complexity
	switch {
	case p.isMapLiteral(0):
		return p.parseMapLiteral()
	case p.isArrayLiteral():
		return p.parseArrayLiteral()
	case p.isCompositeLiteral():
		return p.parseCompositeLiteralExpression()
	case p.isCallExpression():
		return p.parseCallExpression()
	case p.isLiteral():
		return p.parseLiteralExpression()
	case p.isUnaryExpression():
		return p.parseUnaryExpression()
	case p.isBinaryExpression():
		return p.parseBinaryExpression()
	case p.isReference():
		return p.parseReference()
	default:
		p.addError("Required expression, not found.")
		return nil
	}
}

func (p *parser) parseCallExpression() ast.Node {

	name := p.parseIdentifier()

	p.parseRequired(lexer.TknOpenBracket)

	var args []ast.Node
	if !p.parseOptional(lexer.TknCloseBracket) {
		args = append(args, p.parseExpression())
		for p.parseOptional(lexer.TknComma) {
			args = append(args, p.parseExpression())
		}
		p.parseRequired(lexer.TknCloseBracket)
	}

	p.validate(ast.CallExpression)

	return ast.CallExpressionNode{
		Name:      name,
		Arguments: args,
	}
}

func (p *parser) parseBinaryExpression() ast.Node {

	left := p.parseExpression()

	op := p.current().Type
	p.next()

	right := p.parseExpression()

	p.validate(ast.BinaryExpression)

	return ast.BinaryExpressionNode{
		Left:     left,
		Operator: op,
		Right:    right,
	}
}

func (p *parser) parseUnaryExpression() ast.Node {

	op := p.current().Type
	p.next()

	operand := p.parseExpression()

	p.validate(ast.UnaryExpression)

	return ast.UnaryExpressionNode{
		Operator: op,
		Operand:  operand,
	}
}

func (p *parser) parseIndexExpression() ast.Node {

	expr := p.parseExpression()

	p.parseRequired(lexer.TknOpenSquare)

	index := p.parseExpression()

	p.parseRequired(lexer.TknCloseSquare)

	p.validate(ast.IndexExpression)

	return ast.IndexExpressionNode{
		Expression: expr,
		Index:      index,
	}
}

func (p *parser) parseSliceExpression() ast.Node {

	expr := p.parseExpression()

	p.parseRequired(lexer.TknOpenSquare)

	var low, high ast.Node

	first := p.parseExpression()

	if p.parseOptional(lexer.TknColon) {

	} else {
		p.parseRequired(lexer.TknColon)
		low = first
	}

	p.validate(ast.SliceExpression)

	return ast.SliceExpressionNode{
		Expression: expr,
		Low:        low,
		High:       high,
	}
}

func (p *parser) parseLiteralExpression() ast.Node {

	p.validate(ast.Literal)

	return ast.LiteralNode{}
}

func (p *parser) parseCompositeLiteralExpression() ast.Node {

	p.validate(ast.CompositeLiteral)

	return ast.CompositeLiteralNode{}

}

func (p *parser) parseArrayLiteral() ast.Node {
	return ast.ArrayLiteralNode{}
}

func (p *parser) parseMapLiteral() ast.Node {
	return ast.MapLiteralNode{}
}

func (p *parser) parseReference() ast.Node {
	name := p.parseIdentifier()
	return ast.ReferenceNode{
		Name: name,
	}
}
