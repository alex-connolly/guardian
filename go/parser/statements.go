package parser

import (
	"fmt"

	"github.com/end-r/guardian/go/ast"
	"github.com/end-r/guardian/go/lexer"
)

func parseReturnStatement(p *Parser) {

	p.parseRequired(lexer.TknReturn)
	fmt.Println("return")
	node := ast.ReturnStatementNode{
		Results: p.parseExpressionList(),
	}
	fmt.Println("done")
	p.Scope.Declare(flowKey, node)
}

func parseAssignmentStatement(p *Parser) {

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

	node := ast.AssignmentStatementNode{
		Left:  assigned,
		Right: to,
	}
	p.Scope.Declare(flowKey, node)
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
	//init := parseAssignmentStatement(p)
	// parse condition, required
	cond := p.parseExpression()
	// parse statement
	//stat := p.parseAssignment()

	body := ast.ScopeNode{}

	p.parseEnclosedScope(&body)

	node := ast.ForStatementNode{
		Cond:  cond,
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
