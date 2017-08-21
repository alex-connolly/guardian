package parser

import (
	"github.com/end-r/guardian/go/compiler/lexer"
)

func isClassDeclaration(p *Parser) bool {
	if p.index+1 < len(p.lexer.Tokens) {
		return p.current().Type == lexer.TknClass ||
			(p.current().Type == lexer.TknAbstract && p.token(1).Type == lexer.TknClass)
	}
	return p.current().Type == lexer.TknClass
}

func isInterfaceDeclaration(p *Parser) bool {
	if p.index+1 < len(p.lexer.Tokens) {
		return p.current().Type == lexer.TknInterface ||
			(p.current().Type == lexer.TknAbstract && p.token(1).Type == lexer.TknInterface)
	}
	return p.current().Type == lexer.TknInterface
}

func isContractDeclaration(p *Parser) bool {
	if p.index+1 < len(p.lexer.Tokens) {
		return p.current().Type == lexer.TknContract ||
			(p.current().Type == lexer.TknAbstract && p.token(1).Type == lexer.TknContract)
	}
	return p.current().Type == lexer.TknContract
}

func isFuncDeclaration(p *Parser) bool {
	if p.index+2 < len(p.lexer.Tokens) {
		return p.current().Type == lexer.TknIdentifier && p.token(1).Type == lexer.TknOpenBracket ||
			p.current().Type == lexer.TknAbstract && p.token(1).Type == lexer.TknIdentifier
	}
	return false
}

func isTypeDeclaration(p *Parser) bool {
	return p.current().Type == lexer.TknType
}

func isForStatement(p *Parser) bool {
	return p.current().Type == lexer.TknFor
}

func isIfStatement(p *Parser) bool {
	return p.current().Type == lexer.TknIf
}

func isAssignmentStatement(p *Parser) bool {
	return false
}

func isSwitchStatement(p *Parser) bool {
	if p.index+2 < len(p.lexer.Tokens) {
		return p.current().Type == lexer.TknSwitch ||
			(p.current().Type == lexer.TknExclusive && p.token(1).Type == lexer.TknSwitch)
	}
	return p.current().Type == lexer.TknSwitch
}

func isReturnStatement(p *Parser) bool {
	return p.current().Type == lexer.TknReturn
}

func isCaseStatement(p *Parser) bool {
	return p.current().Type == lexer.TknCase
}
