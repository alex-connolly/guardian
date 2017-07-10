package parser

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/compiler/lexer"
)

func (p *parser) parseReturnStatement() {

	p.parseRequired(lexer.TknReturn)

	if p.parseOptional(lexer.TknOpenBracket) {
		var tuple []Node
		tuple = append(tuple, p.parseExpression())
		for p.parseOptional(lexer.TknComma) {
			tuple = append(tuple, p.parseExpression())
		}
	}
	p.parseRequired(lexer.TknCloseBracket)

	p.scope.Validate(ast.ReturnStatement)
}

func (p *parser) parseAssignmentStatement() {

	var assigned []Node
	assigned = append(assigned, p.parseExpression())
	for p.parseOptional(lexer.TknComma) {
		assigned = append(assigned, p.parseExpression())
	}

	p.parseRequired(lexer.TknAssign)

	var to []Node
	to = append(to, p.parseExpression())
	for p.parseOptional(lexer.TknComma) {
		to = append(to, p.parseExpression())
	}

	p.scope.Validate(ast.AssignmentStatement)

	p.scope.Declare("Assignment", ast.AssignmentStatement{
		assigned: assigned,
		to:       to,
	})
}

func (p *parser) parseIfStatement() {

	p.parseRequired(lexer.TknIf)
	p := p.parseExpression
}
