package parser

import (
	"github.com/end-r/guardian/go/compiler/ast"
	"github.com/end-r/guardian/go/compiler/lexer"
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

	node := ast.AssignmentStatementNode{
		Left:  assigned,
		Right: to,
	}
}

func parseIfStatement(p *Parser) {

	p.parseRequired(lexer.TknIf)
	//	init := p.parseStatement()
	//	expr := p.parseExpression()

	// parse elif cases

	// parse else case

}

func parseForStatement(p *Parser) {

	p.parseRequired(lexer.TknFor)
	// parse init expr, can be nil
	//init := p.parseExpression()
	// parse condition, required
	cond := p.parseExpression()
	// parse statement
	//	stat := p.parseStatement()

	body := ast.ScopeNode{}

	p.parseScope(&body)

	node := ast.ForStatementNode{
		//Init:  init,
		Cond:  cond,
		Block: body,
	}

}

func parseCaseStatement(p *Parser) {

	p.parseRequired(lexer.TknCase)

	exprs := p.parseExpressionList()

	body := ast.ScopeNode{}

	p.parseScope(&body)

	node := ast.CaseStatementNode{
		Expressions: exprs,
		Block:       body,
	}
}

func parseSwitchStatement(p *Parser) {

	p.parseRequired(lexer.TknSwitch)
	//expr := p.parseExpression()
	p.parseRequired(lexer.TknOpenBrace)

	//s := ast.SwitchStatementNode{}

	//p.Scope.Declare("", s)

}
