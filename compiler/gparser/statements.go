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
	p.parseOptional(lexer.TknSemicolon)
}

func parseAssignmentStatement(p *Parser) {
	node := p.parseAssignment()
	p.Scope.AddSequential(node)
	p.parseOptional(lexer.TknSemicolon)
}

func (p *Parser) parseOptionalAssignment() *ast.AssignmentStatementNode {
	// all optional assignments must be simple
	if isSimpleAssignmentStatement(p) {
		assigned := p.parseSimpleAssignment()
		p.parseOptional(lexer.TknSemicolon)
		return &assigned
	}
	return nil
}

func (p *Parser) parseSimpleAssignment() ast.AssignmentStatementNode {

	modifiers := p.parseModifiers(lexer.TknIdentifier)

	var assigned []ast.ExpressionNode
	assigned = append(assigned, p.parseSimpleExpression())
	for p.parseOptional(lexer.TknComma) {
		assigned = append(assigned, p.parseSimpleExpression())
	}

	if !p.parseOptional(lexer.GetAssignments()...) {
		if p.parseOptional(lexer.TknIncrement, lexer.TknDecrement) {
			return ast.AssignmentStatementNode{
				Modifiers: modifiers,
				Left:      assigned,
				Right:     nil,
			}
		}
	}

	var to []ast.ExpressionNode
	to = append(to, p.parseSimpleExpression())
	for p.parseOptional(lexer.TknComma) {
		to = append(to, p.parseSimpleExpression())
	}

	return ast.AssignmentStatementNode{
		Modifiers: modifiers,
		Left:      assigned,
		Right:     to,
	}
}

func (p *Parser) parseAssignment() ast.AssignmentStatementNode {

	modifiers := p.parseModifiers(lexer.TknIdentifier)

	var assigned []ast.ExpressionNode
	assigned = append(assigned, p.parseExpression())
	for p.parseOptional(lexer.TknComma) {
		assigned = append(assigned, p.parseExpression())
	}

	if !p.parseOptional(lexer.GetAssignments()...) {
		if p.parseOptional(lexer.TknIncrement, lexer.TknDecrement) {
			return ast.AssignmentStatementNode{
				Modifiers: modifiers,
				Left:      assigned,
				Right:     nil,
			}
		}
	}

	var to []ast.ExpressionNode
	to = append(to, p.parseExpression())
	for p.parseOptional(lexer.TknComma) {
		to = append(to, p.parseExpression())
	}

	return ast.AssignmentStatementNode{
		Modifiers: modifiers,
		Left:      assigned,
		Right:     to,
	}
}

func parseIfStatement(p *Parser) {

	p.parseRequired(lexer.TknIf)

	// parse init expr, can be nil
	init := p.parseOptionalAssignment()

	var conditions []ast.ConditionNode

	// parse initial if condition, required
	cond := p.parseSimpleExpression()

	body := p.parseEnclosedScope()

	conditions = append(conditions, ast.ConditionNode{
		Condition: cond,
		Body:      body,
	})

	// parse elif cases
	for p.parseOptional(lexer.TknElif) {
		condition := p.parseSimpleExpression()
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
		Init:       init,
		Conditions: conditions,
		Else:       elseBlock,
	}

	// BUG: again, why is this necessary?
	if init == nil {
		node.Init = nil
	}

	p.Scope.AddSequential(node)
}

func parseForStatement(p *Parser) {

	p.parseRequired(lexer.TknFor)

	// parse init expr, can be nil
	init := p.parseOptionalAssignment()
	// parse condition, required

	// must evaluate to boolean, checked at validation

	cond := p.parseSimpleExpression()

	// TODO: parse post statement properly

	var post *ast.AssignmentStatementNode

	if p.parseOptional(lexer.TknSemicolon) {
		post = p.parseOptionalAssignment()
	}

	body := p.parseEnclosedScope()

	node := ast.ForStatementNode{
		Init:  init,
		Cond:  cond,
		Post:  post,
		Block: body,
	}

	// BUG: I have absolutely no idea why this is necessary, BUT IT IS
	// BUG: surely the literal should do the job
	// TODO: Is this a golang bug?
	if post == nil {
		node.Post = nil
	}
	p.Scope.AddSequential(node)
}

func parseFlowStatement(p *Parser) {
	node := ast.FlowStatementNode{
		Token: p.current().Type,
	}
	p.next()
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

	// TODO: currently only works with identifier
	target := p.parseIdentifierExpression()

	cases := p.parseEnclosedScope(ast.CaseStatement)

	node := ast.SwitchStatementNode{
		IsExclusive: exclusive,
		Target:      target,
		Cases:       cases,
	}

	p.Scope.AddSequential(node)

}