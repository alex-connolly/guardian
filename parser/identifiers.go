package parser

import (
	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/lexer"
)

func isNewLine(p *Parser) bool {
	return p.isNextToken(lexer.TknNewLine)
}

// e.g. name string
func isExplicitVarDeclaration(p *Parser) bool {
	if !p.hasTokens(2) {
		return false
	}
	saved := *p
	for p.parseOptional(lexer.GetModifiers()...) {
	}
	if !p.parseOptional(lexer.TknIdentifier) {
		*p = saved
		return false
	}
	for p.parseOptional(lexer.TknComma) {
		if !p.parseOptional(lexer.TknIdentifier) {
			*p = saved
			return false
		}
	}
	if !p.hasTokens(1) {
		*p = saved
		return false
	}
	if !p.isNextAType() {
		*p = saved
		return false
	}
	p.parseType()
	if p.hasTokens(1) {
		if !p.isNextTerminating() {
			*p = saved
			return false
		}
	}
	*p = saved
	return true
}

func (p *Parser) isNextTerminating() bool {
	return p.parseOptional(lexer.TknNewLine, lexer.TknSemicolon, lexer.TknComma)
}

func (p *Parser) isNextAType() bool {
	return p.isPlainType() || p.isArrayType() || p.isMapType() || p.isFuncType()
}

func (p *Parser) isPlainType() bool {
	saved := *p
	p.parseOptional(lexer.TknEllipsis)
	expr := p.parseExpressionComponent()
	*p = saved
	if expr == nil {
		return false
	}
	return expr.Type() == ast.Reference || expr.Type() == ast.Identifier
}

func (p *Parser) isArrayType() bool {
	immediate := p.nextTokens(lexer.TknOpenSquare, lexer.TknCloseSquare)
	variable := p.nextTokens(lexer.TknEllipsis, lexer.TknOpenSquare, lexer.TknCloseSquare)
	return immediate || variable
}

func (p *Parser) isFuncType() bool {
	immediate := p.nextTokens(lexer.TknFunc)
	variable := p.nextTokens(lexer.TknEllipsis, lexer.TknFunc)
	return immediate || variable
}

func (p *Parser) isMapType() bool {
	immediate := p.nextTokens(lexer.TknMap)
	variable := p.nextTokens(lexer.TknEllipsis, lexer.TknMap)
	return immediate || variable
}

func (p *Parser) nextTokens(tokens ...lexer.TokenType) bool {
	if !p.hasTokens(len(tokens)) {
		return false
	}
	for i, t := range tokens {
		if p.token(i).Type != t {
			return false
		}
	}
	return true
}

func (p *Parser) modifiersUntilToken(types ...lexer.TokenType) bool {
	saved := p.index
	if p.hasTokens(1) {
		for p.current().Type.IsModifier() {
			p.next()
		}
		for _, t := range types {
			if p.current().Type == t {
				p.index = saved
				return true
			}
		}
	}
	p.index = saved
	return false
}

func isClassDeclaration(p *Parser) bool {
	return p.modifiersUntilToken(lexer.TknClass)
}

func isInterfaceDeclaration(p *Parser) bool {
	return p.modifiersUntilToken(lexer.TknInterface)
}

func isLifecycleDeclaration(p *Parser) bool {
	return p.modifiersUntilToken(lexer.GetLifecycles()...)
}

func isEnumDeclaration(p *Parser) bool {
	return p.modifiersUntilToken(lexer.TknEnum)
}

func isContractDeclaration(p *Parser) bool {
	return p.modifiersUntilToken(lexer.TknContract)
}

func isFuncDeclaration(p *Parser) bool {
	return p.modifiersUntilToken(lexer.TknFunc)
}

func (p *Parser) isNextToken(types ...lexer.TokenType) bool {
	if p.hasTokens(1) {
		for _, t := range types {
			if p.current().Type == t {
				return true
			}
		}
	}
	return false
}

func isEventDeclaration(p *Parser) bool {
	return p.modifiersUntilToken(lexer.TknEvent)
}

func isTypeDeclaration(p *Parser) bool {
	return p.modifiersUntilToken(lexer.TknType)
}

func isForStatement(p *Parser) bool {
	return p.isNextToken(lexer.TknFor)
}

func isIfStatement(p *Parser) bool {
	return p.isNextToken(lexer.TknIf)
}

func isAssignmentStatement(p *Parser) bool {
	return p.preserveState(func(p *Parser) bool {
		for p.parseOptional(lexer.GetModifiers()...) {
		}
		expr := p.parseExpression()
		if expr == nil {
			return false
		}
		for p.parseOptional(lexer.TknComma) {
			// assume these will be expressions
			p.parseExpression()
		}
		return p.isNextTokenAssignment()
	})
}

// performs operations and then returns the parser to its initial state
func (p *Parser) preserveState(a func(p *Parser) bool) bool {
	saved := *p
	flag := a(p)
	*p = saved
	return flag
}

func isSimpleAssignmentStatement(p *Parser) bool {

	return p.preserveState(func(p *Parser) bool {
		for p.parseOptional(lexer.GetModifiers()...) {
		}
		expr := p.parseSimpleExpression()
		if expr == nil {
			return false
		}
		for p.parseOptional(lexer.TknComma) {
			// assume these will be expressions
			p.parseSimpleExpression()
		}
		return p.isNextTokenAssignment()
	})

}

func (p *Parser) isNextTokenAssignment() bool {
	return p.isNextToken(lexer.TknAssign, lexer.TknAddAssign, lexer.TknSubAssign, lexer.TknMulAssign,
		lexer.TknDivAssign, lexer.TknShrAssign, lexer.TknShlAssign, lexer.TknModAssign, lexer.TknAndAssign,
		lexer.TknOrAssign, lexer.TknXorAssign, lexer.TknIncrement, lexer.TknDecrement, lexer.TknDefine)
}

func isFlowStatement(p *Parser) bool {
	return p.isNextToken(lexer.TknBreak, lexer.TknContinue)
}

func isSingleLineComment(p *Parser) bool {
	return p.isNextToken(lexer.TknLineComment)
}

func isMultiLineComment(p *Parser) bool {
	return p.isNextToken(lexer.TknCommentOpen)
}

func isSwitchStatement(p *Parser) bool {
	ex := p.nextTokens(lexer.TknExclusive, lexer.TknSwitch)
	dir := p.nextTokens(lexer.TknSwitch)
	return ex || dir
}

func isReturnStatement(p *Parser) bool {
	return p.isNextToken(lexer.TknReturn)
}

func isCaseStatement(p *Parser) bool {
	return p.isNextToken(lexer.TknCase)
}

func isModifierList(p *Parser) bool {

	return p.preserveState(func(p *Parser) bool {
		for p.parseOptional(lexer.GetModifiers()...) {
		}
		return p.parseOptional(lexer.TknOpenBracket)
	})
}

func isPackageStatement(p *Parser) bool {
	return p.isNextToken(lexer.TknPackage)
}

func isImportStatement(p *Parser) bool {
	return p.isNextToken(lexer.TknImport)
}
