package parser

import (
	"github.com/end-r/guardian/go/compiler/ast"
	"github.com/end-r/guardian/go/compiler/lexer"
)

func (p *Parser) parseExpression() ast.ExpressionNode {
	// Guardian expressions can be arbitrarily chained
	// e.g. array[expr]
	// the expr could be 5 + 4 + 3, or 5 + 4 + getNumber()
	// this is all stored in one expression Node
	// however, ORDER IS IMPORTANT
	// (5 + 4) * 3 vs 5 + 4 * 3
	// these expression are not evaluated at compile time
	// actually, maybe evaluate constants fully

	// (dog() - 5) + 6
	// !((dog() - 5) + 6 > 10)

	// some expressions are 'complete'
	// map, array, compositeLiteral
	// some are 'incomplete'
	// literals, references, calls
	// incomplete expressions can have further ops performed
	// a() + b() * c()
	// TODO: improve the logical flow here: it's horrendous

	switch p.current().Type {
	case lexer.TknMap:
		return p.parseMapLiteral()
	case lexer.TknOpenSquare:
		return p.parseArrayLiteral()
	case lexer.TknString, lexer.TknCharacter, lexer.TknNumber:
		expr := p.parseLiteral()
		if p.current().Type.IsBinaryOperator() {
			p.parseBinaryExpression(expr)
		} else if p.current().Type.IsUnaryOperator() {
			p.parsePostfixUnaryExpression(expr)
		}
		break
	case lexer.TknIdentifier:
		expr := p.parseReference()
		switch p.current().Type {
		case lexer.TknOpenBracket:
			p.parseCallExpression(expr)
			break
		case lexer.TknOpenBrace:
			return p.parseCompositeLiteral(expr)
		case lexer.TknOpenSquare:
			p.parseIndexExpression() // TODO: could be slice
		}
		if p.current().Type.IsBinaryOperator() {
			p.parseBinaryExpression(expr)
		} else if p.current().Type.IsUnaryOperator() {
			p.parsePostfixUnaryExpression(expr)
		}
		break
	case lexer.TknNot, lexer.TknIncrement, lexer.TknDecrement:
		// prefix unary operator
		p.parsePrefixUnaryExpression()
		break
	}
	return p.parseReference()
}

func (p *Parser) parsePrefixUnaryExpression() (n ast.UnaryExpressionNode) {
	n.Operator = p.current().Type
	p.next()
	n.Operand = p.parseExpression()
	return n
}

func (p *Parser) parsePostfixUnaryExpression(expr ast.ExpressionNode) (n ast.UnaryExpressionNode) {
	n.Operand = expr
	n.Operator = p.current().Type
	return n
}

func (p *Parser) parseBinaryExpression(expr ast.ExpressionNode) (n ast.BinaryExpressionNode) {
	n.Left = expr
	n.Operator = p.current().Type
	n.Right = p.parseExpression()
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
	p.parseRequired(lexer.TknOpenSquare)
	n.Key = p.parseReference()
	if !p.parseOptional(lexer.TknCloseSquare) {
		n.Data = append(n.Data, p.parseExpression())
		for p.parseOptional(lexer.TknComma) {
			n.Data = append(n.Data, p.parseExpression())
		}
	}
	return n
}

func (p *Parser) parseMapLiteral() (n ast.MapLiteralNode) {
	p.parseRequired(lexer.TknMap)
	p.parseRequired(lexer.TknOpenSquare)
	n.Key = p.parseReference()
	p.parseRequired(lexer.TknCloseSquare)
	n.Value = p.parseReference()
	p.parseRequired(lexer.TknOpenBrace)
	if !p.parseOptional(lexer.TknCloseBrace) {
		firstKey := p.parseExpression()
		p.parseRequired(lexer.TknColon)
		firstValue := p.parseExpression()
		n.Data[firstKey] = firstValue
		for p.parseOptional(lexer.TknComma) {
			key := p.parseExpression()
			p.parseRequired(lexer.TknColon)
			value := p.parseExpression()
			n.Data[key] = value
		}
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

func (p *Parser) parseIndexExpression() (n ast.IndexExpressionNode) {
	n.Expression = p.parseExpression()
	n.Index = p.parseExpression()
	return n
}

func (p *Parser) parseSliceExpression() (n ast.SliceExpressionNode) {
	n.Expression = p.parseExpression()
	p.parseRequired(lexer.TknOpenSquare)
	n.Low = p.parseExpression()
	p.parseRequired(lexer.TknColon)
	if !p.parseOptional(lexer.TknCloseSquare) {
		n.High = p.parseExpression()
		p.parseRequired(lexer.TknCloseSquare)
	}
	return n
}

func (p *Parser) parseReference() (n ast.ReferenceNode) {
	n.Names = make([]string, 0)
	n.Names = append(n.Names, p.parseIdentifier())
	for p.parseOptional(lexer.TknDot) {
		n.Names = append(n.Names, p.parseIdentifier())
	}
	return n
}

func (p *Parser) parseLiteral() (n ast.LiteralNode) {
	n.LiteralType = p.current().Type
	n.Data = p.lexer.TokenString(p.current())
	return n
}

func (p *Parser) parseCompositeLiteral(ref ast.ReferenceNode) (n ast.CompositeLiteralNode) {
	n.Reference = ref
	return n
}