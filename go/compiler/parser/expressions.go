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

	var expr ast.ExpressionNode
	expr = nil
	//	fmt.Printf("index: %d, tokens: %d\n", p.index, len(p.lexer.Tokens))
	if p.hasTokens(1) {
		switch p.current().Type {
		case lexer.TknMap:
			expr = p.parseMapLiteral()
			break
		case lexer.TknOpenSquare:
			expr = p.parseArrayLiteral()
			break
		case lexer.TknString, lexer.TknCharacter, lexer.TknNumber:
			expr = p.parseLiteral()
			return expr
		case lexer.TknIdentifier:
			expr = p.parseReference()
			break
		case lexer.TknNot, lexer.TknIncrement, lexer.TknDecrement:
			expr = p.parsePrefixUnaryExpression()
		}
	}
	// parse composite expressiona
	for p.hasTokens(1) && expr != nil {
		t := p.current().Type
		switch {
		case t == lexer.TknOpenBracket:
			expr = p.parseCallExpression(expr)
			break
		case t == lexer.TknOpenBrace:
			expr = p.parseCompositeLiteral(expr)
			break
		case t == lexer.TknOpenSquare:
			expr = p.parseIndexExpression(expr)
			break
		case t.IsBinaryOperator():
			expr = p.parseBinaryExpression(expr)
			break
		case t.IsUnaryOperator():
			expr = p.parsePostfixUnaryExpression(expr)
			break
		default:
			return expr
		}
	}
	return expr
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
		n.Size = p.parseExpression()
		p.parseRequired(lexer.TknCloseSquare)
	}
	p.parseRequired(lexer.TknOpenBrace)
	if !p.parseOptional(lexer.TknCloseBrace) {
		n.Data = append(n.Data, p.parseExpression())
		for p.parseOptional(lexer.TknComma) {
			n.Data = append(n.Data, p.parseExpression())
		}
		p.parseRequired(lexer.TknCloseBrace)
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

func (p *Parser) parseReference() (n ast.ReferenceNode) {
	n.Names = []string{p.parseIdentifier()}
	for p.parseOptional(lexer.TknDot) {
		n.Names = append(n.Names, p.parseIdentifier())
	}
	return n
}

func (p *Parser) parseLiteral() (n ast.LiteralNode) {
	n.LiteralType = p.current().Type
	n.Data = p.lexer.TokenString(p.current())
	p.next()
	return n
}

func (p *Parser) parseCompositeLiteral(expr ast.ExpressionNode) (n ast.CompositeLiteralNode) {
	// expr must be a reference node
	n.Reference = expr.(ast.ReferenceNode)
	p.parseRequired(lexer.TknOpenBrace)
	if !p.parseOptional(lexer.TknCloseBrace) {
		firstKey := p.parseIdentifier()
		p.parseRequired(lexer.TknColon)
		expr := p.parseExpression()
		if n.Fields == nil {
			n.Fields = make(map[string]ast.ExpressionNode)
		}
		n.Fields[firstKey] = expr
		for p.parseOptional(lexer.TknComma) {
			key := p.parseIdentifier()
			p.parseRequired(lexer.TknColon)
			exp := p.parseExpression()
			n.Fields[key] = exp
		}
		p.parseRequired(lexer.TknCloseBrace)
	}
	return n
}
