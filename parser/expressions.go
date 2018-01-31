package parser

import (
	"fmt"

	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/ast"
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

var operators = map[token.Type]Operator{
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

func parseSimpleExpression(p *Parser) ast.ExpressionNode {
	return p.parseSimpleExpression()
}

func parseExpression(p *Parser) ast.ExpressionNode {
	return p.parseExpression()
}

func pushNode(stack []ast.ExpressionNode, op token.Type) []ast.ExpressionNode {
	n := ast.BinaryExpressionNode{}
	n.Right, stack = stack[len(stack)-1], stack[:len(stack)-1]
	n.Left, stack = stack[len(stack)-1], stack[:len(stack)-1]
	n.Operator = op
	return append(stack, &n)
}

func (p *Parser) parseExpression() ast.ExpressionNode {
	p.seenCastOperator = false
	lastWasExpression := false
	var opStack []token.Type
	var expStack []ast.ExpressionNode
	var current token.Type
main:
	for p.hasTokens(1) {
		current = p.current().Type
		switch current {
		case token.LineComment:
			return p.finalise(expStack, opStack)
		case token.MultilineComment:
			p.next()
			break
		case token.Semicolon, token.NewLine, token.Comma:
			return p.finalise(expStack, opStack)
		case token.OpenBracket:
			p.parseRequired(token.OpenBracket)
			opStack = append(opStack, current)
			break
		case token.CloseBracket:
			var op token.Type
			for len(opStack) > 0 {
				op, opStack = opStack[len(opStack)-1], opStack[:len(opStack)-1]
				if op == token.OpenBracket {
					p.next()
					continue main
				} else {
					if len(expStack) < 2 {
						p.addError(p.getCurrentTokenLocation(), errIncompleteExpression)
						return nil
					}
					expStack = pushNode(expStack, op)
				}
			}
			// unmatched parens
			return p.finalise(expStack, opStack)
		default:
			parseIgnored(p)
			if o1, ok := operators[current]; ok {
				lastWasExpression = false
				if current == token.As || current == token.Is {
					p.seenCastOperator = true
				}
				p.next()
				for len(opStack) > 0 {
					// consider top item on stack
					op := opStack[len(opStack)-1]
					if o2, isOp := operators[op]; !isOp || o1.prec > o2.prec ||
						o1.prec == o2.prec && o1.rightAssoc {
						break
					} else {
						opStack = opStack[:len(opStack)-1]
						if len(expStack) < 2 {
							p.addError(p.getCurrentTokenLocation(), errIncompleteExpression)
							return nil
						}
						expStack = pushNode(expStack, op)
					}
				}

				opStack = append(opStack, current)
			} else {
				saved := *p
				expr := p.parseExpressionComponent()
				// if it isn't an expression
				if expr == nil {
					*p = saved
					return p.finalise(expStack, opStack)
				} else {
					if lastWasExpression {
						*p = saved
						p.addError(p.getCurrentTokenLocation(), fmt.Sprintf(errConsecutiveExpression, p.current().String(p.lexer)))
						return p.finalise(expStack, opStack)
					}
				}
				lastWasExpression = true
				expStack = append(expStack, expr)
			}
			break
		}
	}
	if len(expStack) == 0 {
		return nil
	}
	return p.finalise(expStack, opStack)
}

func (p *Parser) finalise(expStack []ast.ExpressionNode, opStack []token.Type) ast.ExpressionNode {
	for len(opStack) > 0 {
		if len(expStack) < 2 {
			p.addError(p.getCurrentTokenLocation(), errIncompleteExpression)
			return nil
		}
		n := ast.BinaryExpressionNode{}
		n.Right, expStack = expStack[len(expStack)-1], expStack[:len(expStack)-1]
		n.Left, expStack = expStack[len(expStack)-1], expStack[:len(expStack)-1]
		n.Operator, opStack = opStack[len(opStack)-1], opStack[:len(opStack)-1]
		n.Begin = n.Right.Start()
		n.Final = n.Left.End()
		expStack = append(expStack, &n)
	}
	if len(expStack) == 0 {
		return nil
	}
	return expStack[len(expStack)-1]
}

func (p *Parser) parseExpressionComponent() ast.ExpressionNode {

	if p.seenCastOperator {
		p.seenCastOperator = false
		if p.isNextAType() {
			n := p.parseType()
			// all types can be converted to expressions for this purpose
			// they resolve to unknown for now
			return n.(ast.ExpressionNode)
		}
		p.addError(p.getCurrentTokenLocation(), errInvalidTypeAfterCast)
		return nil
	}

	var expr ast.ExpressionNode
	//fmt.Printf("index: %d, tokens: %d\n", p.index, len(p.lexer.Tokens))
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
	case token.New:
		return p.parseKeywordExpression()
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
	n.Begin = p.getCurrentTokenLocation()
	n.Operator = p.current().Type
	p.next()
	n.Operand = p.parseExpression()
	n.Final = p.getLastTokenLocation()
	return n
}

func (p *Parser) parseReference(expr ast.ExpressionNode) *ast.ReferenceNode {
	n := new(ast.ReferenceNode)
	n.Begin = p.getCurrentTokenLocation()
	p.parseRequired(token.Dot)
	n.Parent = expr
	n.Reference = p.parseExpressionComponent()
	n.Final = p.getLastTokenLocation()
	return n
}

func (p *Parser) parseIdentifierList() []string {
	var ids []string
	ids = append(ids, p.parseIdentifier())
	for p.parseOptional(token.Comma) {
		ids = append(ids, p.parseIdentifier())
	}
	return ids
}

func (p *Parser) parseKeywordExpression() *ast.KeywordNode {
	n := new(ast.KeywordNode)
	n.Begin = p.getCurrentTokenLocation()
	n.Keyword = p.parseRequired(token.New)
	n.TypeNode = p.parseType()
	p.parseRequired(token.OpenBracket)
	if !p.parseOptional(token.CloseBracket) {
		p.ignoreNewLines()
		n.Arguments = p.parseExpressionList()
		p.ignoreNewLines()
		p.parseRequired(token.CloseBracket)
	}
	n.Final = p.getLastTokenLocation()
	return n
}

func (p *Parser) parseCallExpression(expr ast.ExpressionNode) *ast.CallExpressionNode {
	n := new(ast.CallExpressionNode)
	n.Begin = expr.Start()
	n.Call = expr
	p.parseRequired(token.OpenBracket)
	if !p.parseOptional(token.CloseBracket) {
		p.ignoreNewLines()
		n.Arguments = p.parseExpressionList()
		p.ignoreNewLines()
		p.parseRequired(token.CloseBracket)
	}
	n.Final = p.getLastTokenLocation()
	return n
}

func (p *Parser) parseArrayLiteral() *ast.ArrayLiteralNode {

	n := new(ast.ArrayLiteralNode)
	n.Begin = p.getCurrentTokenLocation()
	// [string:3]{"Dog", "Cat", ""}

	n.Signature = p.parseArrayType()

	p.parseRequired(token.OpenBrace)
	if !p.parseOptional(token.CloseBrace) {
		p.ignoreNewLines()
		// TODO: check this is right
		n.Data = p.parseExpressionList()
		p.ignoreNewLines()
		p.parseRequired(token.CloseBrace)
	}
	n.Final = p.getLastTokenLocation()
	return n
}

func (p *Parser) parseFuncLiteral() *ast.FuncLiteralNode {
	n := new(ast.FuncLiteralNode)

	n.Begin = p.getCurrentTokenLocation()

	p.parseRequired(token.Func)

	n.Generics = p.parsePossibleGenerics()

	n.Parameters = p.parseParameters()

	n.Results = p.parseResults()

	n.Scope = p.parseBracesScope()
	n.Final = p.getLastTokenLocation()
	return n
}

func (p *Parser) parseMapLiteral() *ast.MapLiteralNode {
	n := new(ast.MapLiteralNode)
	n.Begin = p.getCurrentTokenLocation()
	n.Signature = p.parseMapType()

	p.parseRequired(token.OpenBrace)
	if !p.parseOptional(token.CloseBrace) {
		p.ignoreNewLines()
		firstKey := p.parseExpression()
		p.parseRequired(token.Colon)
		firstValue := p.parseExpression()
		p.ignoreNewLines()
		n.Data = make(map[ast.ExpressionNode]ast.ExpressionNode)
		n.Data[firstKey] = firstValue
		for p.parseOptional(token.Comma) {
			p.ignoreNewLines()
			if p.parseOptional(token.CloseBrace) {
				return n
			}
			key := p.parseExpression()
			p.parseRequired(token.Colon)
			value := p.parseExpression()
			n.Data[key] = value
			p.ignoreNewLines()
		}
		p.ignoreNewLines()
		p.parseRequired(token.CloseBrace)
	}
	n.Final = p.getLastTokenLocation()
	return n
}

func (p *Parser) parseExpressionList() (list []ast.ExpressionNode) {
	first := p.parseExpression()
	if first == nil {
		return list
	}
	list = append(list, first)
	for p.parseOptional(token.Comma) {
		p.ignoreNewLines()
		list = append(list, p.parseExpression())
	}
	return list
}

func (p *Parser) parseIndexExpression(expr ast.ExpressionNode) ast.ExpressionNode {
	n := ast.IndexExpressionNode{}
	n.Begin = expr.Start()
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
	n.Final = p.getLastTokenLocation()
	return &n
}

func (p *Parser) parseSliceExpression(expr ast.ExpressionNode,
	first ast.ExpressionNode) *ast.SliceExpressionNode {
	n := ast.SliceExpressionNode{}
	n.Begin = p.getCurrentTokenLocation()
	n.Expression = expr
	n.Low = first
	if !p.parseOptional(token.CloseSquare) {
		n.High = p.parseExpression()
		p.parseRequired(token.CloseSquare)
	}
	n.Final = p.getLastTokenLocation()
	return &n
}

func (p *Parser) parseIdentifierExpression() *ast.IdentifierNode {
	n := new(ast.IdentifierNode)
	n.Begin = p.getCurrentTokenLocation()
	n.Name = p.parseIdentifier()
	n.Final = p.getLastTokenLocation()
	return n
}

func (p *Parser) parseLiteral() *ast.LiteralNode {
	n := new(ast.LiteralNode)
	n.Begin = p.getCurrentTokenLocation()
	n.LiteralType = p.current().Type
	n.Data = p.current().TrimmedString(p.lexer)
	p.next()
	n.Final = p.getLastTokenLocation()
	return n
}

func (p *Parser) parseCompositeLiteral() *ast.CompositeLiteralNode {
	n := new(ast.CompositeLiteralNode)
	// expr must be a reference or identifier node
	n.Begin = p.getCurrentTokenLocation()
	n.TypeName = p.parsePlainType()

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
	n.Final = p.getLastTokenLocation()
	return n
}
