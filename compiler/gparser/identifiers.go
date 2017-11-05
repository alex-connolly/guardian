package gparser

import (
	"github.com/end-r/guardian/compiler/ast"

	"github.com/end-r/guardian/compiler/lexer"
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
	expr := p.parseExpressionComponent()
	*p = saved
	if expr == nil {
		return false
	}
	return expr.Type() == ast.Reference || expr.Type() == ast.Identifier
}

func (p *Parser) isArrayType() bool {
	return p.isNextToken(lexer.TknOpenSquare) && p.token(1).Type == lexer.TknCloseSquare
}

func (p *Parser) isFuncType() bool {
	return p.isNextToken(lexer.TknFunc)
}

func (p *Parser) isMapType() bool {
	return p.isNextToken(lexer.TknMap)
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
	saved := *p
	for p.parseOptional(lexer.GetModifiers()...) {
	}
	expr := p.parseExpression()
	if expr == nil {
		*p = saved
		return false
	}
	for p.parseOptional(lexer.TknComma) {
		// assume these will be expressions
		p.parseExpression()
	}
	flag := p.isNextTokenAssignment()
	*p = saved
	return flag
}

func isSimpleAssignmentStatement(p *Parser) bool {
	saved := *p
	for p.parseOptional(lexer.GetModifiers()...) {
	}
	expr := p.parseSimpleExpression()
	if expr == nil {
		*p = saved
		return false
	}
	for p.parseOptional(lexer.TknComma) {
		// assume these will be expressions
		p.parseSimpleExpression()
	}
	flag := p.isNextTokenAssignment()
	*p = saved
	return flag
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
	if p.hasTokens(2) {
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
