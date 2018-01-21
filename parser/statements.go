package parser

import (
	"fmt"

	"github.com/end-r/guardian/token"

	"github.com/blang/semver"
	"github.com/end-r/guardian/ast"
)

func parseReturnStatement(p *Parser) {

	start := p.getCurrentTokenLocation()

	p.parseRequired(token.Return)
	node := ast.ReturnStatementNode{
		Begin:   start,
		Final:   p.getLastTokenLocation(),
		Results: p.parseExpressionList(),
	}
	p.scope.AddSequential(&node)
}

func (p *Parser) parseOptionalAssignment() *ast.AssignmentStatementNode {
	// all optional assignments must be simple
	if isSimpleAssignmentStatement(p) {
		assigned := p.parseSimpleAssignment()
		p.parseOptional(token.Semicolon)
		return assigned
	}
	return nil
}

func (p *Parser) parsePostAssignment(assigned []ast.ExpressionNode) *ast.AssignmentStatementNode {
	if len(assigned) > 1 {
		p.addError(p.getCurrentTokenLocation(), errInvalidIncDec)
	}
	b := &ast.BinaryExpressionNode{
		Left: assigned[0],
		Right: &ast.LiteralNode{
			Begin:       p.getCurrentTokenLocation(),
			Final:       p.getLastTokenLocation(),
			LiteralType: token.Integer,
			Data:        "1",
		},
	}
	if p.parseOptional(token.Increment) {
		b.Operator = token.Add
	} else if p.parseOptional(token.Decrement) {
		b.Operator = token.Sub
	}
	p.parseOptional(token.Semicolon)
	return &ast.AssignmentStatementNode{
		Begin: assigned[0].Start(),
		Final: p.getLastTokenLocation(),
		Left:  assigned,
		Right: []ast.ExpressionNode{b},
	}
}

var assigns = map[token.Type]token.Type{
	token.AddAssign: token.Add,
	token.SubAssign: token.Sub,
	token.MulAssign: token.Mul,
	token.DivAssign: token.Div,
	token.ModAssign: token.Mod,
	token.ExpAssign: token.Exp,
	token.AndAssign: token.And,
	token.OrAssign:  token.Or,
	token.XorAssign: token.Xor,
	token.ShlAssign: token.Shl,
	token.ShrAssign: token.Shr,
}

func (p *Parser) parseAssignmentOf(assigned []ast.ExpressionNode, expressionParser func(p *Parser) ast.ExpressionNode) *ast.AssignmentStatementNode {
	switch p.current().Type {
	case token.Increment, token.Decrement:
		return p.parsePostAssignment(assigned)
	case token.Assign:
		p.parseRequired(token.Assign)
		var to []ast.ExpressionNode

		to = append(to, expressionParser(p))
		for p.parseOptional(token.Comma) {
			to = append(to, expressionParser(p))
		}
		p.parseOptional(token.Semicolon)
		return &ast.AssignmentStatementNode{
			Begin: assigned[0].Start(),
			Final: p.getLastTokenLocation(),
			Left:  assigned,
			Right: to,
		}
	default:
		a, ok := assigns[p.current().Type]
		if ok {
			p.next()
			var to []ast.ExpressionNode
			to = append(to, expressionParser(p))
			for p.parseOptional(token.Comma) {
				to = append(to, expressionParser(p))
			}
			p.parseOptional(token.Semicolon)
			return &ast.AssignmentStatementNode{
				Begin:    assigned[0].Start(),
				Final:    p.getLastTokenLocation(),
				Left:     assigned,
				Right:    to,
				Operator: a,
			}
		}
		break
	}
	return nil
}

func (p *Parser) parseAssignment(expressionParser func(p *Parser) ast.ExpressionNode) *ast.AssignmentStatementNode {

	var assigned []ast.ExpressionNode
	assigned = append(assigned, expressionParser(p))
	for p.parseOptional(token.Comma) {
		assigned = append(assigned, expressionParser(p))
	}

	return p.parseAssignmentOf(assigned, expressionParser)
}

func (p *Parser) parseSimpleAssignment() *ast.AssignmentStatementNode {
	return p.parseAssignment(parseSimpleExpression)
}

func (p *Parser) parseFullAssignment() *ast.AssignmentStatementNode {
	return p.parseAssignment(parseExpression)
}

func parseIfStatement(p *Parser) {

	start := p.getCurrentTokenLocation()

	p.parseRequired(token.If)

	// parse init expr, can be nil
	init := p.parseOptionalAssignment()

	var conditions []*ast.ConditionNode

	condStart := p.getCurrentTokenLocation()

	// parse initial if condition, required
	cond := p.parseSimpleExpression()

	body := p.parseBracesScope()

	conditions = append(conditions, &ast.ConditionNode{
		Begin:     condStart,
		Final:     p.getLastTokenLocation(),
		Condition: cond,
		Body:      body,
	})

	// parse elif cases
	for p.parseOptional(token.ElseIf) {
		condStart := p.getCurrentTokenLocation()
		condition := p.parseSimpleExpression()
		body := p.parseBracesScope()
		conditions = append(conditions, &ast.ConditionNode{
			Begin:     condStart,
			Final:     p.getLastTokenLocation(),
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
		Begin:      start,
		Final:      p.getLastTokenLocation(),
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

	start := p.getCurrentTokenLocation()

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
		Begin: start,
		Final: p.getLastTokenLocation(),
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

	start := p.getCurrentTokenLocation()

	p.parseRequired(token.For)

	vars := p.parseIdentifierList()

	p.parseRequired(token.In)

	producer := p.parseSimpleExpression()

	body := p.parseBracesScope()

	node := ast.ForEachStatementNode{
		Begin:     start,
		Final:     p.getLastTokenLocation(),
		Variables: vars,
		Producer:  producer,
		Block:     body,
	}

	p.scope.AddSequential(&node)
}

func parseFlowStatement(p *Parser) {
	node := ast.FlowStatementNode{
		Token: p.current().Type,
		Begin: p.getCurrentTokenLocation(),
		Final: p.getLastTokenLocation(),
	}
	p.next()
	p.scope.AddSequential(&node)
}

func parseCaseStatement(p *Parser) {

	start := p.getCurrentTokenLocation()

	p.parseRequired(token.Case)

	exprs := p.parseExpressionList()

	p.parseRequired(token.Colon)

	saved := p.scope
	p.scope = new(ast.ScopeNode)
	for p.hasTokens(1) && !p.isNextToken(token.Case, token.Default, token.CloseBrace) {
		p.parseNextConstruct()
	}

	node := ast.CaseStatementNode{
		Begin:       start,
		Final:       p.getLastTokenLocation(),
		Expressions: exprs,
		Block:       p.scope,
	}
	p.scope = saved
	p.scope.AddSequential(&node)
}

func parseSwitchStatement(p *Parser) {

	start := p.getCurrentTokenLocation()

	exclusive := p.parseOptional(token.Exclusive)

	p.parseRequired(token.Switch)

	target := p.parseSimpleExpression()

	cases := p.parseBracesScope(ast.CaseStatement)

	node := ast.SwitchStatementNode{
		Begin:       start,
		Final:       p.getLastTokenLocation(),
		IsExclusive: exclusive,
		Target:      target,
		Cases:       cases,
	}

	p.scope.AddSequential(&node)

}

func parseImportStatement(p *Parser) {

	p.parseGroupable(token.Import, func(arg2 *Parser) {

		start := p.getCurrentTokenLocation()
		var alias, path string

		if p.isNextToken(token.String) {
			path = p.current().TrimmedString(p.lexer)
		} else {
			p.addError(p.current().Start, fmt.Sprintf(errInvalidImportPath, p.current().String(p.lexer)))
		}

		p.next()

		if p.parseOptional(token.As) {
			alias = p.parseIdentifier()
		}
		p.ignoreNewLines()

		node := ast.ImportStatementNode{
			Begin: start,
			Final: p.getLastTokenLocation(),
			Alias: alias,
			Path:  path,
		}

		p.scope.AddSequential(&node)
	})
}

func parsePackageStatement(p *Parser) {

	start := p.getCurrentTokenLocation()

	p.parseRequired(token.Package)

	name := p.parseIdentifier()

	p.parseRequired(token.Guardian)

	version := p.parseSemver()

	node := ast.PackageStatementNode{
		Begin:   start,
		Final:   p.getLastTokenLocation(),
		Name:    name,
		Version: version,
	}

	p.scope.AddSequential(&node)
}

func (p *Parser) parseSemver() semver.Version {
	s := ""
	for p.hasTokens(1) && !p.isNextToken(token.NewLine, token.Semicolon) {
		s += p.current().String(p.lexer)
		p.next()
	}
	v, err := semver.Parse(s)
	if err != nil {
		p.addError(p.getCurrentTokenLocation(), fmt.Sprintf("Invalid semantic version %s", s))
	}
	return v
}
