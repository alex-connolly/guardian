package parser

import (
	"github.com/end-r/guardian/go/compiler/lexer"
)

func isScopeClosure(p *Parser) bool {
	return p.isNextToken(lexer.TknCloseBrace)
}

// e.g. name string
func isExplicitVarDeclaration(p *Parser) bool {
	return p.isNextToken(lexer.TknIdentifier) && p.token(1).Type == lexer.TknIdentifier
}

func isClassDeclaration(p *Parser) bool {
	if p.hasTokens(2) {
		return p.isNextToken(lexer.TknClass) ||
			(p.current().Type == lexer.TknAbstract && p.token(1).Type == lexer.TknClass)
	}
	return p.isNextToken(lexer.TknClass)
}

func isInterfaceDeclaration(p *Parser) bool {
	if p.hasTokens(2) {
		return p.isNextToken(lexer.TknInterface) ||
			(p.current().Type == lexer.TknAbstract && p.token(1).Type == lexer.TknInterface)
	}
	return p.isNextToken(lexer.TknInterface)
}

func isContractDeclaration(p *Parser) bool {
	if p.hasTokens(2) {
		return p.isNextToken(lexer.TknContract) ||
			p.isNextToken(lexer.TknAbstract) && p.token(1).Type == lexer.TknContract
	}
	return p.isNextToken(lexer.TknContract)
}

func isFuncDeclaration(p *Parser) bool {
	if p.hasTokens(2) {
		return p.isNextToken(lexer.TknFunc) ||
			(p.current().Type == lexer.TknAbstract && p.token(1).Type == lexer.TknFunc)
	}
	return p.isNextToken(lexer.TknFunc)
}

func (p *Parser) isNextToken(t lexer.TokenType) bool {
	if p.hasTokens(1) {
		return p.current().Type == t
	}
	return false
}

func isNewLine(p *Parser) bool {
	return p.isNextToken(lexer.TknNewLine)
}

func isEventDeclaration(p *Parser) bool {
	return p.isNextToken(lexer.TknEvent)
}

func isTypeDeclaration(p *Parser) bool {
	return p.isNextToken(lexer.TknType)
}

func isForStatement(p *Parser) bool {
	return p.isNextToken(lexer.TknFor)
}

func isIfStatement(p *Parser) bool {
	return p.isNextToken(lexer.TknIf)
}

func isAssignmentStatement(p *Parser) bool {
	return false
}

func isSwitchStatement(p *Parser) bool {
	if p.index+2 < len(p.lexer.Tokens) {
		return p.isNextToken(lexer.TknSwitch) ||
			(p.current().Type == lexer.TknExclusive && p.token(1).Type == lexer.TknSwitch)
	}
	return p.isNextToken(lexer.TknSwitch)
}

func isReturnStatement(p *Parser) bool {
	return p.isNextToken(lexer.TknReturn)
}

func isCaseStatement(p *Parser) bool {
	return p.isNextToken(lexer.TknCase)
}
