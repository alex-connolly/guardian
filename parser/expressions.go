package parser

import (
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
	lexer.TknShl: exponentiveOperator,
	lexer.TknShr: exponentiveOperator,
	// multiplicative operators
	lexer.TknMul: multiplicativeOperator,
	lexer.TknDiv: multiplicativeOperator,
	lexer.TknMod: multiplicativeOperator,
	lexer.TknAnd: multiplicativeOperator,
	// additive operators
	lexer.TknAdd: additiveOperator,
	lexer.TknSub: additiveOperator,
	lexer.TknOr:  additiveOperator,
	lexer.TknXor: additiveOperator,
	// cast operators
	lexer.TknIs: castOperator,
	lexer.TknAs: castOperator,
	// comparative operators
	lexer.TknLss: comparativeOperator,
	lexer.TknGtr: comparativeOperator,
	lexer.TknGeq: comparativeOperator,
	lexer.TknLeq: comparativeOperator,
	lexer.TknEql: comparativeOperator,
	lexer.TknNeq: comparativeOperator,
	// logical operators
	lexer.TknLogicalAnd: conjunctiveOperator,
	lexer.TknLogicalOr:  disjunctiveOperator,
	// ternary operators
	lexer.TknTernary: ternaryOperator,
	/* assignment operators
	lexer.TknAssign:    assignmentOperator,
	lexer.TknMulAssign: assignmentOperator,
	lexer.TknDivAssign: assignmentOperator,
	lexer.TknModAssign: assignmentOperator,
	lexer.TknAddAssign: assignmentOperator,
	lexer.TknSubAssign: assignmentOperator,
	lexer.TknShlAssign: assignmentOperator,
	lexer.TknShrAssign: assignmentOperator,
	lexer.TknAndAssign: assignmentOperator,
	lexer.TknOrAssign:  assignmentOperator,
	lexer.TknXorAssign: assignmentOperator,*/
}

func pushNode(stack []ast.ExpressionNode, op lexer.TokenType) []ast.ExpressionNode {
	n := ast.BinaryExpressionNode{}
	n.Right, stack = stack[len(stack)-1], stack[:len(stack)-1]
	n.Left, stack = stack[len(stack)-1], stack[:len(stack)-1]
	n.Operator = op
	return append(stack, n)
}

func (p *Parser) parseExpression() ast.ExpressionNode {
	var opStack []lexer.TokenType
	var expStack []ast.ExpressionNode
	var current lexer.TokenType
main:
	for p.hasTokens(1) {
		current = p.current().Type
		switch current {
		case lexer.TknOpenBracket:
			p.next()
			opStack = append(opStack, current)
			break
		case lexer.TknCloseBracket:
			var op lexer.TokenType
			for len(opStack) > 0 {
				op, opStack = opStack[len(opStack)-1], opStack[:len(opStack)-1]
				if op == lexer.TknOpenBracket {
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
		expStack = append(expStack, n)
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
	case lexer.TknMap:
		return p.parseMapLiteral()
	case lexer.TknOpenSquare:
		return p.parseArrayLiteral()
	case lexer.TknString, lexer.TknCharacter, lexer.TknInteger, lexer.TknFloat,
		lexer.TknTrue, lexer.TknFalse:
		return p.parseLiteral()
	case lexer.TknIdentifier:
		return p.parseIdentifierExpression()
	case lexer.TknNot, lexer.TknTypeOf:
		return p.parsePrefixUnaryExpression()
	case lexer.TknFunc:
		return p.parseFuncLiteral()
	}
	return nil
}

func (p *Parser) parseSecondaryComponent(expr ast.ExpressionNode) (e ast.ExpressionNode, done bool) {
	switch p.current().Type {
	case lexer.TknOpenBracket:
		return p.parseCallExpression(expr), false
	case lexer.TknOpenSquare:
		return p.parseIndexExpression(expr), false
	case lexer.TknDot:
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
	return p.current().Type == lexer.TknIdentifier && p.token(1).Type == lexer.TknOpenBrace
}

func (p *Parser) parsePrefixUnaryExpression() (n ast.UnaryExpressionNode) {
	n.Operator = p.current().Type
	p.next()
	n.Operand = p.parseExpression()
	return n
}

func (p *Parser) parseReference(expr ast.ExpressionNode) (n ast.ReferenceNode) {
	p.parseRequired(lexer.TknDot)
	n.Parent = expr
	n.Reference = p.parseExpressionComponent()
	return n
}

func (p *Parser) parseCallExpression(expr ast.ExpressionNode) (n ast.CallExpressionNode) {
	n.Call = expr
	p.parseRequired(lexer.TknOpenBracket)
	if !p.parseOptional(lexer.TknCloseBracket) {
		n.Arguments = p.parseExpressionList()
		p.parseRequired(lexer.TknCloseBracket)
	}
	return n
}

func (p *Parser) parseArrayLiteral() (n ast.ArrayLiteralNode) {
	// [string:3]{"Dog", "Cat", ""}

	n.Signature = p.parseArrayType()

	p.parseRequired(lexer.TknOpenBrace)
	if !p.parseOptional(lexer.TknCloseBrace) {
		// TODO: check this is right
		n.Data = p.parseExpressionList()
		p.parseRequired(lexer.TknCloseBrace)
	}
	return n
}

func (p *Parser) parseFuncLiteral() (n ast.FuncLiteralNode) {
	n.Parameters = p.parseParameters()

	n.Results = p.parseResults()

	n.Scope = p.parseBracesScope()
	return n
}

func (p *Parser) parseMapLiteral() (n ast.MapLiteralNode) {
	n.Signature = p.parseMapType()

	p.parseRequired(lexer.TknOpenBrace)
	if !p.parseOptional(lexer.TknCloseBrace) {
		firstKey := p.parseExpression()
		p.parseRequired(lexer.TknColon)
		firstValue := p.parseExpression()
		n.Data = make(map[ast.ExpressionNode]ast.ExpressionNode)
		n.Data[firstKey] = firstValue
		for p.parseOptional(lexer.TknComma) {
			key := p.parseExpression()
			p.parseRequired(lexer.TknColon)
			value := p.parseExpression()
			n.Data[key] = value
		}
		p.parseRequired(lexer.TknCloseBrace)
	}
	return n
}

func (p *Parser) parseExpressionList() (list []ast.ExpressionNode) {
	list = append(list, p.parseExpression())
	for p.parseOptional(lexer.TknComma) {
		list = append(list, p.parseExpression())
	}
	return list
}

func (p *Parser) parseIndexExpression(expr ast.ExpressionNode) ast.ExpressionNode {
	n := ast.IndexExpressionNode{}
	p.parseRequired(lexer.TknOpenSquare)
	if p.parseOptional(lexer.TknColon) {
		return p.parseSliceExpression(expr, nil)
	}
	index := p.parseExpression()
	if p.parseOptional(lexer.TknColon) {
		return p.parseSliceExpression(expr, index)
	}
	p.parseRequired(lexer.TknCloseSquare)
	n.Expression = expr
	n.Index = index

	return n
}

func (p *Parser) parseSliceExpression(expr ast.ExpressionNode,
	first ast.ExpressionNode) ast.SliceExpressionNode {
	n := ast.SliceExpressionNode{}
	n.Expression = expr
	n.Low = first
	if !p.parseOptional(lexer.TknCloseSquare) {
		n.High = p.parseExpression()
		p.parseRequired(lexer.TknCloseSquare)
	}
	return n
}

func (p *Parser) parseIdentifierExpression() (n ast.IdentifierNode) {
	n.Name = p.parseIdentifier()
	return n
}

func (p *Parser) parseLiteral() (n ast.LiteralNode) {
	n.LiteralType = p.current().Type
	n.Data = p.current().String()
	p.next()
	return n
}

func (p *Parser) parseCompositeLiteral() (n ast.CompositeLiteralNode) {
	// expr must be a reference or identifier node
	n.TypeName = p.parseIdentifier()

	p.parseRequired(lexer.TknOpenBrace)
	for p.parseOptional(lexer.TknNewLine) {
	}
	if !p.parseOptional(lexer.TknCloseBrace) {
		firstKey := p.parseIdentifier()
		p.parseRequired(lexer.TknColon)
		expr := p.parseExpression()
		if n.Fields == nil {
			n.Fields = make(map[string]ast.ExpressionNode)
		}
		n.Fields[firstKey] = expr
		for p.parseOptional(lexer.TknComma) {
			for p.parseOptional(lexer.TknNewLine) {
			}
			if p.parseOptional(lexer.TknCloseBrace) {
				return n
			}
			key := p.parseIdentifier()
			p.parseRequired(lexer.TknColon)
			exp := p.parseExpression()
			n.Fields[key] = exp
		}
		// TODO: allow more than one?
		p.parseOptional(lexer.TknNewLine)
		p.parseRequired(lexer.TknCloseBrace)
	}
	return n
}
