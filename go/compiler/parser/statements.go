package parser

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/compiler/lexer"
)

func parseReturnStatement(p *Parser) {

	p.parseRequired(lexer.TknReturn)

	if p.parseOptional(lexer.TknOpenBracket) {
		var tuple []ast.ExpressionNode
		tuple = append(tuple, p.parseExpression())
		for p.parseOptional(lexer.TknComma) {
			tuple = append(tuple, p.parseExpression())
		}
	}
	p.parseRequired(lexer.TknCloseBracket)

	p.Scope.Validate(ast.ReturnStatement)
}

func parseAssignmentStatement(p *Parser) {

	var assigned []ast.ExpressionNode
	assigned = append(assigned, p.parseExpression())
	for p.parseOptional(lexer.TknComma) {
		assigned = append(assigned, p.parseExpression())
	}

	p.parseRequired(lexer.TknAssign)

	var to []ast.ExpressionNode
	to = append(to, p.parseExpression())
	for p.parseOptional(lexer.TknComma) {
		to = append(to, p.parseExpression())
	}

	p.Scope.Validate(ast.AssignmentStatement)

	p.Scope.Declare("Assignment", ast.AssignmentStatementNode{
		Left:  assigned,
		Right: to,
	})
}

func parseIfStatement(p *Parser) {

	p.parseRequired(lexer.TknIf)
	//one := p.parseExpression()

}

func parseForStatement(p *Parser) {

	p.parseRequired(lexer.TknFor)

}

func parseCaseStatement(p *Parser) {

}

func parseSwitchStatement(p *Parser) {

}
