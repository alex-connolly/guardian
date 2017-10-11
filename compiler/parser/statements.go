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
	p.Scope.AddSequential(node)
}

func parseAssignmentStatement(p *Parser) {
	node := p.parseAssignment()
	p.Scope.AddSequential(node)
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

	// parse init expr, can be nil
	// init := p.parseAssignment()

	var conditions []ast.ConditionNode

	// parse initial if condition, required
	cond := p.parseExpression()
	body := p.parseEnclosedScope()

	conditions = append(conditions, ast.ConditionNode{
		Condition: cond,
		Body:      body,
	})

	// parse elif cases
	for p.parseOptional(lexer.TknElif) {
		condition := p.parseExpression()
		body := p.parseEnclosedScope()
		conditions = append(conditions, ast.ConditionNode{
			Condition: condition,
			Body:      body,
		})
	}

	var elseBlock *ast.ScopeNode
	// parse else case
	if p.parseOptional(lexer.TknElse) {
		elseBlock = p.parseEnclosedScope()
	}

	node := ast.IfStatementNode{
		Init:       nil,
		Conditions: conditions,
		Else:       elseBlock,
	}
	p.Scope.AddSequential(node)
}

func parseForStatement(p *Parser) {

	p.parseRequired(lexer.TknFor)
	// parse init expr, can be nil
	//init := p.parseAssignment()
	// parse condition, required
	cond := p.parseExpression()
	// parse statement
	post := p.parseAssignment()

	body := p.parseEnclosedScope()

	node := ast.ForStatementNode{
		Init:  nil,
		Cond:  cond,
		Post:  post,
		Block: body,
	}
	p.Scope.AddSequential(node)

}

func parseCaseStatement(p *Parser) {

	p.parseRequired(lexer.TknCase)

	exprs := p.parseExpressionList()

	body := p.parseScope()

	node := ast.CaseStatementNode{
		Expressions: exprs,
		Block:       body,
	}
	p.Scope.AddSequential(node)
}

func parseSwitchStatement(p *Parser) {

	exclusive := p.parseOptional(lexer.TknExclusive)

	p.parseRequired(lexer.TknSwitch)

	target := p.parseExpression()

	cases := p.parseEnclosedScope(ast.CaseStatement)

	node := ast.SwitchStatementNode{
		IsExclusive: exclusive,
		Target:      target,
		Cases:       cases,
	}

	p.Scope.AddSequential(node)

}
