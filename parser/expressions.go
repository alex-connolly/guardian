package parser

import (
	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/lexer"
)

// Operator ...
type Operator struct {
	rightAssoc bool
	prec       int
}

// Guardian uses Swift's operator precedence table
var (
	//assignmentOperator     = Operator{true, 90}
	ternaryOperator        = Operator{true, 100}
	disjunctiveOperator    = Operator{true, 110}
	conjunctiveOperator    = Operator{true, 120}
	comparativeOperator    = Operator{false, 130}
	castOperator           = Operator{false, 132}
	rangeOperator          = Operator{false, 135}
	additiveOperator       = Operator{false, 140}
	multiplicativeOperator = Operator{false, 150}
	exponentiveOperator    = Operator{false, 160}
)

var operators = map[lexer.TokenType]Operator{
	// exponentive operators
	token.Shl: exponentiveOperator,
	token.Shr: exponentiveOperator,
	// multiplicative operators
	token.Mul: multiplicativeOperator,
	token.Div: multiplicativeOperator,
	token.Mod: multiplicativeOperator,
	token.And: multiplicativeOperator,
	// additive operators
	token.Add: additiveOperator,
	token.Sub: additiveOperator,
	token.Or:  additiveOperator,
	token.Xor: additiveOperator,
	// cast operators
	token.Is: castOperator,
	token.As: castOperator,
	// comparative operators
	token.Lss: comparativeOperator,
	token.Gtr: comparativeOperator,
	token.Geq: comparativeOperator,
	token.Leq: comparativeOperator,
	token.Eql: comparativeOperator,
	token.Neq: comparativeOperator,
	// logical operators
	token.LogicalAnd: conjunctiveOperator,
	token.LogicalOr:  disjunctiveOperator,
	// ternary operators
	token.Ternary: ternaryOperator,
	/* assignment operators
	token.Assign:    assignmentOperator,
	token.MulAssign: assignmentOperator,
	token.DivAssign: assignmentOperator,
	token.ModAssign: assignmentOperator,
	token.AddAssign: assignmentOperator,
	token.SubAssign: assignmentOperator,
	token.ShlAssign: assignmentOperator,
	token.ShrAssign: assignmentOperator,
	token.AndAssign: assignmentOperator,
	token.OrAssign:  assignmentOperator,
	token.XorAssign: assignmentOperator,*/
}

func pushNode(stack []ast.ExpressionNode, op lexer.TokenType) []ast.ExpressionNode {
	n := ast.BinaryExpressionNode{}
	n.Right, stack = stack[len(stack)-1], stack[:len(stack)-1]
	n.Left, stack = stack[len(stack)-1], stack[:len(stack)-1]
	n.Operator = op
	return append(stack, &n)
}

func (p *Parser) parseExpression() ast.ExpressionNode {
	var opStack []lexer.TokenType
	var expStack []ast.ExpressionNode
	var current lexer.TokenType
main:
	for p.hasTokens(1) {
		current = p.current().Type
		switch current {
		case token.OpenBracket:
			p.next()
			opStack = append(opStack, current)
			break
		case token.CloseBracket:
			var op lexer.TokenType
			for len(opStack) > 0 {
				op, opStack = opStack[len(opStack)-1], opStack[:len(opStack)-1]
				if op == token.OpenBracket {
					p.next()
					continue main
				} else {
					expStack = pushNode(expStack, op)
				}
			}
			// unmatched parens
			return finalise(expStack, opStack)
		default:
			if o1, ok := operators[current]; ok {
				p.next()
				for len(opStack) > 0 {
					// consider top item on stack
					op := opStack[len(opStack)-1]
					if o2, isOp := operators[op]; !isOp || o1.prec > o2.prec ||
						o1.prec == o2.prec && o1.rightAssoc {
						break
					} else {
						opStack = opStack[:len(opStack)-1]
						expStack = pushNode(expStack, op)
					}
				}

				opStack = append(opStack, current)
			} else {
				expr := p.parseExpressionComponent()

				// if it isn't an expression
				if expr == nil {
					return finalise(expStack, opStack)
				}
				expStack = append(expStack, expr)
			}
			break
		}
	}
	if len(expStack) == 0 {
		return nil
	}
	return finalise(expStack, opStack)
}

func finalise(expStack []ast.ExpressionNode, opStack []lexer.TokenType) ast.ExpressionNode {
	for len(opStack) > 0 {

		n := ast.BinaryExpressionNode{}
		n.Right, expStack = expStack[len(expStack)-1], expStack[:len(expStack)-1]
		n.Left, expStack = expStack[len(expStack)-1], expStack[:len(expStack)-1]
		n.Operator, opStack = opStack[len(opStack)-1], opStack[:len(opStack)-1]
		expStack = append(expStack, &n)
	}
	if len(expStack) == 0 {
		return nil
	}

	return expStack[len(expStack)-1]
}

func (p *Parser) parseExpressionComponent() ast.ExpressionNode {

	var expr ast.ExpressionNode
	//fmt.Printf("index: %d, tokens: %d\n", p.index, len(p.tokens))
	if p.hasTokens(1) {
		expr = p.parsePrimaryComponent()
	}
	// parse composite expressions
	var done bool
	for p.hasTokens(1) && expr != nil && !done {
		expr, done = p.parseSecondaryComponent(expr)
	}

	return expr
}

func (p *Parser) parsePrimaryComponent() ast.ExpressionNode {
	if p.isCompositeLiteral() {
		return p.parseCompositeLiteral()
	}
	switch p.current().Type {
	case token.Map:
		return p.parseMapLiteral()
	case token.OpenSquare:
		return p.parseArrayLiteral()
	case token.String, token.Character, token.Integer, token.Float,
		token.True, token.False:
		return p.parseLiteral()
	case token.Identifier:
		return p.parseIdentifierExpression()
	case token.Not, token.TypeOf:
		return p.parsePrefixUnaryExpression()
	case token.Func:
		return p.parseFuncLiteral()
	}
	return nil
}

func (p *Parser) parseSecondaryComponent(expr ast.ExpressionNode) (e ast.ExpressionNode, done bool) {
	switch p.current().Type {
	case token.OpenBracket:
		return p.parseCallExpression(expr), false
	case token.OpenSquare:
		return p.parseIndexExpression(expr), false
	case token.Dot:
		return p.parseReference(expr), false
	}
	return expr, true
}

func (p *Parser) parseSimpleExpression() ast.ExpressionNode {
	p.simple = true
	expr := p.parseExpression()
	p.simple = false
	return expr
}

// Dog{age: 6}.age {} = true
// 4 {} = false
// dog {} = false
// if 5 > Dog{} {} = true
// x = dog {} = false
func (p *Parser) isCompositeLiteral() bool {
	if p.simple {
		return false
	}
	if !p.hasTokens(2) {
		return false
	}
	return p.current().Type == token.Identifier && p.token(1).Type == token.OpenBrace
}

func (p *Parser) parsePrefixUnaryExpression() *ast.UnaryExpressionNode {
	n := new(ast.UnaryExpressionNode)
	n.Operator = p.current().Type
	p.next()
	n.Operand = p.parseExpression()
	return n
}

func (p *Parser) parseReference(expr ast.ExpressionNode) *ast.ReferenceNode {
	n := new(ast.ReferenceNode)
	p.parseRequired(token.Dot)
	n.Parent = expr
	n.Reference = p.parseExpressionComponent()
	return n
}

func (p *Parser) parseIdentifierList() []string {
	var ids []string
	if p.current().Type != token.String {
		return nil
	}
	ids = append(ids, p.parseIdentifier())
	for p.parseOptional(token.Comma) {
		ids = append(ids, p.parseIdentifier())
	}
	return ids
}

func (p *Parser) parseCallExpression(expr ast.ExpressionNode) *ast.CallExpressionNode {
	n := new(ast.CallExpressionNode)
	n.Call = expr
	p.parseRequired(token.OpenBracket)
	if !p.parseOptional(token.CloseBracket) {
		n.Arguments = p.parseExpressionList()
		p.parseRequired(token.CloseBracket)
	}
	return n
}

func (p *Parser) parseArrayLiteral() *ast.ArrayLiteralNode {

	n := new(ast.ArrayLiteralNode)
	// [string:3]{"Dog", "Cat", ""}

	n.Signature = p.parseArrayType()

	p.parseRequired(token.OpenBrace)
	if !p.parseOptional(token.CloseBrace) {
		// TODO: check this is right
		n.Data = p.parseExpressionList()
		p.parseRequired(token.CloseBrace)
	}
	return n
}

func (p *Parser) parseFuncLiteral() *ast.FuncLiteralNode {
	n := new(ast.FuncLiteralNode)
	n.Parameters = p.parseParameters()

	n.Results = p.parseResults()

	n.Scope = p.parseBracesScope()
	return n
}

func (p *Parser) parseMapLiteral() *ast.MapLiteralNode {
	n := new(ast.MapLiteralNode)
	n.Signature = p.parseMapType()

	p.parseRequired(token.OpenBrace)
	if !p.parseOptional(token.CloseBrace) {
		firstKey := p.parseExpression()
		p.parseRequired(token.Colon)
		firstValue := p.parseExpression()
		n.Data = make(map[ast.ExpressionNode]ast.ExpressionNode)
		n.Data[firstKey] = firstValue
		for p.parseOptional(token.Comma) {
			key := p.parseExpression()
			p.parseRequired(token.Colon)
			value := p.parseExpression()
			n.Data[key] = value
		}
		p.parseRequired(token.CloseBrace)
	}
	return n
}

func (p *Parser) parseExpressionList() (list []ast.ExpressionNode) {
	list = append(list, p.parseExpression())
	for p.parseOptional(token.Comma) {
		list = append(list, p.parseExpression())
	}
	return list
}

func (p *Parser) parseIndexExpression(expr ast.ExpressionNode) ast.ExpressionNode {
	n := ast.IndexExpressionNode{}
	p.parseRequired(token.OpenSquare)
	if p.parseOptional(token.Colon) {
		return p.parseSliceExpression(expr, nil)
	}
	index := p.parseExpression()
	if p.parseOptional(token.Colon) {
		return p.parseSliceExpression(expr, index)
	}
	p.parseRequired(token.CloseSquare)
	n.Expression = expr
	n.Index = index
	return &n
}

func (p *Parser) parseSliceExpression(expr ast.ExpressionNode,
	first ast.ExpressionNode) *ast.SliceExpressionNode {
	n := ast.SliceExpressionNode{}
	n.Expression = expr
	n.Low = first
	if !p.parseOptional(token.CloseSquare) {
		n.High = p.parseExpression()
		p.parseRequired(token.CloseSquare)
	}
	return &n
}

func (p *Parser) parseIdentifierExpression() *ast.IdentifierNode {
	n := new(ast.IdentifierNode)
	n.Name = p.parseIdentifier()
	return n
}

func (p *Parser) parseLiteral() *ast.LiteralNode {
	n := new(ast.LiteralNode)
	n.LiteralType = p.current().Type
	n.Data = p.current().String()
	p.next()
	return n
}

func (p *Parser) parseCompositeLiteral() *ast.CompositeLiteralNode {
	n := new(ast.CompositeLiteralNode)
	// expr must be a reference or identifier node
	n.TypeName = p.parseIdentifier()

	p.parseRequired(token.OpenBrace)
	for p.parseOptional(token.NewLine) {
	}
	if !p.parseOptional(token.CloseBrace) {
		firstKey := p.parseIdentifier()
		p.parseRequired(token.Colon)
		expr := p.parseExpression()
		if n.Fields == nil {
			n.Fields = make(map[string]ast.ExpressionNode)
		}
		n.Fields[firstKey] = expr
		for p.parseOptional(token.Comma) {
			for p.parseOptional(token.NewLine) {
			}
			if p.parseOptional(token.CloseBrace) {
				return n
			}
			key := p.parseIdentifier()
			p.parseRequired(token.Colon)
			exp := p.parseExpression()
			n.Fields[key] = exp
		}
		// TODO: allow more than one?
		p.parseOptional(token.NewLine)
		p.parseRequired(token.CloseBrace)
	}
	return n
}
