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
