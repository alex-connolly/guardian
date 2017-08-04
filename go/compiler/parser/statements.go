package parser

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/compiler/lexer"
)

func parseReturnStatement(p *Parser) {

	p.parseRequired(lexer.TknReturn)

	if p.parseOptional(lexer.TknOpenBracket) {
		var tuple []ast.Node
		tuple = append(tuple, p.parseExpression())
		for p.parseOptional(lexer.TknComma) {
			tuple = append(tuple, p.parseExpression())
		}
	}
	p.parseRequired(lexer.TknCloseBracket)

	p.scope.Validate(ast.ReturnStatement)
}

func parseAssignmentStatement(p *Parser) {

	var assigned []ast.Node
	assigned = append(assigned, p.parseExpression())
	for p.parseOptional(lexer.TknComma) {
		assigned = append(assigned, p.parseExpression())
	}

	p.parseRequired(lexer.TknAssign)

	var to []ast.Node
	to = append(to, p.parseExpression())
	for p.parseOptional(lexer.TknComma) {
		to = append(to, p.parseExpression())
	}

	p.scope.Validate(ast.AssignmentStatement)

	p.scope.Declare("Assignment", ast.AssignmentStatementNode{
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
