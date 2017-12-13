package parser

import (
	"fmt"

	"github.com/end-r/guardian/token"

	"github.com/blang/semver"
	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/lexer"
)

func parseReturnStatement(p *Parser) {

	p.parseRequired(token.Return)
	node := ast.ReturnStatementNode{
		Results: p.parseExpressionList(),
	}
	p.scope.AddSequential(&node)
	p.parseOptional(token.Semicolon)
}

func parseAssignmentStatement(p *Parser) {
	node := p.parseAssignment()
	p.scope.AddSequential(&node)
	p.parseOptional(token.Semicolon)
}

func (p *Parser) parseOptionalAssignment() *ast.AssignmentStatementNode {
	// all optional assignments must be simple
	if isSimpleAssignmentStatement(p) {
		assigned := p.parseSimpleAssignment()
		p.parseOptional(token.Semicolon)
		return &assigned
	}
	return nil
}

func (p *Parser) parseSimpleAssignment() ast.AssignmentStatementNode {

	modifiers := p.parseKeywords(token.Identifier)

	var assigned []ast.ExpressionNode
	assigned = append(assigned, p.parseSimpleExpression())
	for p.parseOptional(token.Comma) {
		assigned = append(assigned, p.parseSimpleExpression())
	}

	if !p.parseOptional(lexer.GetAssignments()...) {
		if p.parseOptional(token.Increment, token.Decrement) {
			return ast.AssignmentStatementNode{
				Modifiers: modifiers,
				Left:      assigned,
				Right:     nil,
			}
		}
	}

	var to []ast.ExpressionNode
	to = append(to, p.parseSimpleExpression())
	for p.parseOptional(token.Comma) {
		to = append(to, p.parseSimpleExpression())
	}

	return ast.AssignmentStatementNode{
		Modifiers: modifiers,
		Left:      assigned,
		Right:     to,
	}
}

func (p *Parser) parseAssignment() ast.AssignmentStatementNode {

	modifiers := p.parseKeywords(token.Identifier)

	var assigned []ast.ExpressionNode
	assigned = append(assigned, p.parseExpression())
	for p.parseOptional(token.Comma) {
		assigned = append(assigned, p.parseExpression())
	}

	if !p.parseOptional(lexer.GetAssignments()...) {
		if p.parseOptional(token.Increment, token.Decrement) {
			return ast.AssignmentStatementNode{
				Modifiers: modifiers,
				Left:      assigned,
				Right:     nil,
			}
		}
	}

	var to []ast.ExpressionNode
	to = append(to, p.parseExpression())
	for p.parseOptional(token.Comma) {
		to = append(to, p.parseExpression())
	}

	p.parseOptional(token.Semicolon)

	return ast.AssignmentStatementNode{
		Modifiers: modifiers,
		Left:      assigned,
		Right:     to,
	}
}

func parseIfStatement(p *Parser) {

	p.parseRequired(token.If)

	// parse init expr, can be nil
	init := p.parseOptionalAssignment()

	var conditions []*ast.ConditionNode

	// parse initial if condition, required
	cond := p.parseSimpleExpression()

	body := p.parseBracesScope()

	conditions = append(conditions, &ast.ConditionNode{
		Condition: cond,
		Body:      body,
	})

	// parse elif cases
	for p.parseOptional(token.Elif) {
		condition := p.parseSimpleExpression()
		body := p.parseBracesScope()
		conditions = append(conditions, &ast.ConditionNode{
			Condition: condition,
			Body:      body,
		})
	}

	var elseBlock *ast.ScopeNode
	// parse else case
	if p.parseOptional(token.Else) {
		elseBlock = p.parseBracesScope()
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

	p.scope.AddSequential(&node)
}

func parseForStatement(p *Parser) {

	p.parseRequired(token.For)

	// parse init expr, can be nil
	// TODO: should be able to be any statement
	init := p.parseOptionalAssignment()

	// parse condition, required

	// must evaluate to boolean, checked at validation

	cond := p.parseSimpleExpression()

	// TODO: parse post statement properly

	var post *ast.AssignmentStatementNode

	if p.parseOptional(token.Semicolon) {
		post = p.parseOptionalAssignment()
	}

	body := p.parseBracesScope()

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
	p.scope.AddSequential(&node)
}

func parseForEachStatement(p *Parser) {

	p.parseRequired(token.For)

	vars := p.parseIdentifierList()

	p.parseRequired(token.In)

	producer := p.parseExpression()

	body := p.parseBracesScope()

	node := ast.ForEachStatementNode{
		Variables: vars,
		Producer:  producer,
		Block:     body,
	}

	p.scope.AddSequential(&node)
}

func parseFlowStatement(p *Parser) {
	node := ast.FlowStatementNode{
		Token: p.current().Type,
	}
	p.next()
	p.scope.AddSequential(&node)
}

func parseCaseStatement(p *Parser) {

	p.parseRequired(token.Case)

	exprs := p.parseExpressionList()

	p.parseRequired(token.Colon)

	body := p.parseBracesScope()

	node := ast.CaseStatementNode{
		Expressions: exprs,
		Block:       body,
	}
	p.scope.AddSequential(&node)
}

func parseSwitchStatement(p *Parser) {

	exclusive := p.parseOptional(token.Exclusive)

	p.parseRequired(token.Switch)

	// TODO: currently only works with identifier
	target := p.parseIdentifierExpression()

	cases := p.parseBracesScope(ast.CaseStatement)

	node := ast.SwitchStatementNode{
		IsExclusive: exclusive,
		Target:      target,
		Cases:       cases,
	}

	p.scope.AddSequential(&node)

}

func parseImportStatement(p *Parser) {
	p.parseRequired(token.Import)

	var alias, path string

	if p.isNextToken(token.Identifier) {
		alias = p.parseIdentifier()
	}
	path = p.current().TrimmedString()

	node := ast.ImportStatementNode{
		Alias: alias,
		Path:  path,
	}

	p.scope.AddSequential(&node)
}

func parsePackageStatement(p *Parser) {
	p.parseRequired(token.Package)

	name := p.parseIdentifier()

	p.parseRequired(token.At)

	version := p.parseSemver()

	node := ast.PackageStatementNode{
		Name:    name,
		Version: version,
	}

	p.scope.AddSequential(&node)
}

func (p *Parser) parseSemver() semver.Version {
	s := ""
	for p.hasTokens(1) && p.current().Type != token.NewLine {
		s += p.current().String()
		p.next()
	}
	v, err := semver.Make(s)
	if err != nil {
		p.addError(fmt.Sprintf("Invalid semantic version %s", s))
	}
	return v
}
