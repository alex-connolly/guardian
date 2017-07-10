package parser

import "axia/guardian/go/compiler/ast"

func (p *parser) parseBinaryExpression() {

	left := p.parseExpression()

	op := p.current().Type
	p.next()

	right := p.parseExpression()

	p.scope.Validate(ast.BinaryExpressionNode)

	p.scope.Declare("expression", ast.BinaryExpressionNode{
		left:  left,
		op:    op,
		right: right,
	})

}

func (p *parser) parseUnaryExpression() {

	op := p.current().Type
	p.next()

	operand := p.parseExpression()

	p.scope.Validate(ast.UnaryExpressionNode)

	p.scope.Declare("expression", ast.UnaryExpressionNode{
		op:      op,
		operand: operand,
	})
}

func (p *parser) parseIndexExpression() {

	expr := p.parseExpression()

	p.parseRequired(lexer.TknOpenSquare)

	index := p.parseExpression()

	p.parseRequired(lexer.TknClosedSquare)

	p.scope.Validate(ast.IndexExpressionNode)

	p.scope.Declare("expression", ast.IndexExpressionNode{
		expr:  expr,
		index: index,
	})
}

func (p *parser) parseLiteralExpression() {

	p.scope.Validate(ast.LiteralExpressionNode)

	p.scope.Declare("expression", ast.LiteralExpressionNode{})
}
