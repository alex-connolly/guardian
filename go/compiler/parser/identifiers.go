package parser

import (
	"axia/guardian/go/compiler/lexer"
)

func isClassDeclaration(p *parser) bool {
	if p.index+1 < len(p.lexer.Tokens) {
		return p.current().Type == lexer.TknClass ||
			(p.current().Type == lexer.TknAbstract && p.token(1).Type == lexer.TknClass)
	}
	return p.current().Type == lexer.TknClass
}

func isInterfaceDeclaration(p *parser) bool {
	if p.index+1 < len(p.lexer.Tokens) {
		return p.current().Type == lexer.TknInterface ||
			(p.current().Type == lexer.TknAbstract && p.token(1).Type == lexer.TknInterface)
	}
	return p.current().Type == lexer.TknInterface
}

func isContractDeclaration(p *parser) bool {
	if p.index+1 < len(p.lexer.Tokens) {
		return p.current().Type == lexer.TknContract ||
			(p.current().Type == lexer.TknAbstract && p.token(1).Type == lexer.TknContract)
	}
	return p.current().Type == lexer.TknContract
}

func isFuncDeclaration(p *parser) bool {
	if p.index+2 < len(p.lexer.Tokens) {
		return p.current().Type == lexer.TknIdentifier && p.token(1).Type == lexer.TknOpenBracket
	}
	return false
}

func isTypeDeclaration(p *parser) bool {
	return p.current().Type == lexer.TknType
}

func isForStatement(p *parser) bool {
	return p.current().Type == lexer.TknFor
}

func isIfStatement(p *parser) bool {
	return p.current().Type == lexer.TknIf
}

func isAssignmentStatement(p *parser) bool {
	return false
}

func isSwitchStatement(p *parser) bool {
	if p.index+2 < len(p.lexer.Tokens) {
		return p.current().Type == lexer.TknSwitch ||
			(p.current().Type == lexer.TknExclusive && p.token(1).Type == lexer.TknSwitch)
	}
	return p.current().Type == lexer.TknSwitch
}

func isReturnStatement(p *parser) bool {
	return p.current().Type == lexer.TknReturn
}

func isCaseStatement(p *parser) bool {
	return p.current().Type == lexer.TknCase
}

func (p *parser) isExpression(offset int) (bool, int) {
	// possibly recursive
	return p.isBinaryExpression(offset) || p.isUnaryExpression(offset) || p.isReference(offset) ||
		p.isArrayLiteral(offset) || p.isMapLiteral(offset) || p.isLiteral(offset) ||
		p.isCallExpression(offset) || p.isCompositeLiteral(offset)
}

func isOperator(offset int) (bool, int) {
	switch t {
	case lexer.TknAdd, lexer.TknSub, lexer.TknDiv, lexer.TknAddress:
		return true, 1
	case lexer.TknEql:
		return true, 2
	}
	return false, 0
}

func (p *parser) isBinaryExpression(offset int) (bool, int) {
	expr, len := p.isExpression(0)
	if !expr {
		return false, 0
	}
	op, len2 := isOperator(len)
	if !op {
		return false, 0
	}
	expr2, len3 := p.isExpression(len + len2)
	if !expr2 {
		return false
	}
	return true, len + len2 + len3
}

func (p *parser) isUnaryExpression(offset int) (bool, int) {
	op, len := isOperator(0, p.current().Type)
	if !op {
		return false
	}
	return p.isExpression(len)
}

func (p *parser) isCallExpression(offset int) bool {
	return p.token(offset).Type == lexer.TknIdentifier && p.token(offset+1).Type == lexer.TknOpenBracket
}

func (p *parser) isMapLiteral(offset int) bool {
	// map[key]value{}
	return p.current().Type == lexer.TknMap
}

func (p *parser) isArrayLiteral(offset int) bool {
	// []type{}
	return p.current().Type == lexer.TknOpenSquare && p.token(3).Type == lexer.TknOpenBrace
}

func (p *parser) isLiteral(offset int) bool {
	return (p.current().Type == lexer.TknString) || (p.current().Type == lexer.TknNumber)
}

func (p *parser) isCompositeLiteral(offset int) bool {
	if p.index+1 > len(p.lexer.Tokens) {
		return false
	}
	return (p.current().Type == lexer.TknIdentifier) && (p.token(1).Type == lexer.TknOpenBrace)
}

func (p *parser) isReference(offset int) bool {
	return p.current().Type == lexer.TknIdentifier
}
