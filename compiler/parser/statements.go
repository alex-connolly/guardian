package parser

import (
	"github.com/end-r/guardian/compiler/ast"
	"github.com/end-r/guardian/compiler/lexer"
)

func parseReturnStatement(p *Parser) {

	p.parseRequired(lexer.TknReturn)
	node := ast.ReturnStatementNode{
		Results: p.parseExpressionList(),
	}
	p.Scope.Declare(flowKey, node)
}

func parseAssignmentStatement(p *Parser) {
	node := p.parseAssignment()
	p.Scope.Declare(flowKey, node)
}

func (p *Parser) parseAssignment() ast.AssignmentStatementNode {
	var assigned []ast.ExpressionNode
	assigned = append(assigned, p.parseExpression())
	for p.parseOptional(lexer.TknComma) {
		assigned = append(assigned, p.parseExpression())
	}

	p.parseRequired(lexer.TknAssign, lexer.TknAddAssign, lexer.TknSubAssign, lexer.TknMulAssign,
		lexer.TknDivAssign, lexer.TknShrAssign, lexer.TknShlAssign, lexer.TknModAssign, lexer.TknAndAssign,
		lexer.TknOrAssign, lexer.TknXorAssign)

	var to []ast.ExpressionNode
	to = append(to, p.parseExpression())
	for p.parseOptional(lexer.TknComma) {
		to = append(to, p.parseExpression())
	}

	return ast.AssignmentStatementNode{
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
	init := p.parseAssignment()
	// parse condition, required
	cond := p.parseExpression()
	// parse statement
	post := p.parseAssignment()

	body := ast.ScopeNode{}

	p.parseEnclosedScope(&body)

	node := ast.ForStatementNode{
		Init:  init,
		Cond:  cond,
		Post:  post,
		Block: body,
	}
	p.Scope.Declare(flowKey, node)

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
	p.Scope.Declare(flowKey, node)
}

func parseSwitchStatement(p *Parser) {

	p.parseRequired(lexer.TknSwitch)
	//expr := p.parseExpression()
	p.parseRequired(lexer.TknOpenBrace)

	//s := ast.SwitchStatementNode{}

	//p.Scope.Declare("", s)

}
