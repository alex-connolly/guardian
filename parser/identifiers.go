package parser

import (
	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/lexer"
)

func isNewLine(p *Parser) bool {
	return p.isNextToken(token.NewLine)
}

// e.g. name string
func isExplicitVarDeclaration(p *Parser) bool {
	if !p.hasTokens(2) {
		return false
	}
	saved := *p
	for p.parseOptional(lexer.GetModifiers()...) {
	}
	if !p.parseOptional(token.Identifier) {
		*p = saved
		return false
	}
	for p.parseOptional(token.Comma) {
		if !p.parseOptional(token.Identifier) {
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
	return p.parseOptional(token.NewLine, token.Semicolon, token.Comma, token.CloseBracket)
}

func (p *Parser) isNextAType() bool {
	return p.isPlainType() || p.isArrayType() || p.isMapType() || p.isFuncType()
}

func (p *Parser) isPlainType() bool {
	saved := *p
	p.parseOptional(token.Ellipsis)
	expr := p.parseExpressionComponent()
	*p = saved
	if expr == nil {
		return false
	}
	return expr.Type() == ast.Reference || expr.Type() == ast.Identifier
}

func (p *Parser) isArrayType() bool {
	immediate := p.nextTokens(token.OpenSquare)
	variable := p.nextTokens(token.Ellipsis, token.OpenSquare)
	return immediate || variable
}

func (p *Parser) isFuncType() bool {
	immediate := p.nextTokens(token.Func)
	variable := p.nextTokens(token.Ellipsis, token.Func)
	return immediate || variable
}

func (p *Parser) isMapType() bool {
	immediate := p.nextTokens(token.Map)
	variable := p.nextTokens(token.Ellipsis, token.Map)
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
	return p.modifiersUntilToken(token.Class)
}

func isInterfaceDeclaration(p *Parser) bool {
	return p.modifiersUntilToken(token.Interface)
}

func isLifecycleDeclaration(p *Parser) bool {
	return p.modifiersUntilToken(lexer.GetLifecycles()...)
}

func isEnumDeclaration(p *Parser) bool {
	return p.modifiersUntilToken(token.Enum)
}

func isContractDeclaration(p *Parser) bool {
	return p.modifiersUntilToken(token.Contract)
}

func isFuncDeclaration(p *Parser) bool {
	return p.modifiersUntilToken(token.Func)
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
	return p.modifiersUntilToken(token.Event)
}

func isTypeDeclaration(p *Parser) bool {
	return p.modifiersUntilToken(token.Type)
}

func isForStatement(p *Parser) bool {
	return p.isNextToken(token.For)
}

func isIfStatement(p *Parser) bool {
	return p.isNextToken(token.If)
}

func isAssignmentStatement(p *Parser) bool {
	return p.preserveState(func(p *Parser) bool {
		for p.parseOptional(lexer.GetModifiers()...) {
		}
		expr := p.parseExpression()
		if expr == nil {
			return false
		}
		for p.parseOptional(token.Comma) {
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
		for p.parseOptional(token.Comma) {
			// assume these will be expressions
			p.parseSimpleExpression()
		}
		return p.isNextTokenAssignment()
	})

}

func (p *Parser) isNextTokenAssignment() bool {
	return p.isNextToken(token.Assign, token.AddAssign, token.SubAssign, token.MulAssign,
		token.DivAssign, token.ShrAssign, token.ShlAssign, token.ModAssign, token.AndAssign,
		token.OrAssign, token.XorAssign, token.Increment, token.Decrement, token.Define)
}

func isFlowStatement(p *Parser) bool {
	return p.isNextToken(token.Break, token.Continue)
}

func isSingleLineComment(p *Parser) bool {
	return p.isNextToken(token.LineComment)
}

func isMultiLineComment(p *Parser) bool {
	return p.isNextToken(token.CommentOpen)
}

func isSwitchStatement(p *Parser) bool {
	ex := p.nextTokens(token.Exclusive, token.Switch)
	dir := p.nextTokens(token.Switch)
	return ex || dir
}

func isReturnStatement(p *Parser) bool {
	return p.isNextToken(token.Return)
}

func isCaseStatement(p *Parser) bool {
	return p.isNextToken(token.Case)
}

func isKeywordGroup(p *Parser) bool {

	return p.preserveState(func(p *Parser) bool {

		for p.parseOptional(append(lexer.GetModifiers(), lexer.GetDeclarations()...)...) {
		}
		return p.parseOptional(token.OpenBracket)
	})
}

func isPackageStatement(p *Parser) bool {
	return p.isNextToken(token.Package)
}

func isImportStatement(p *Parser) bool {
	return p.isNextToken(token.Import)
}

func isForEachStatement(p *Parser) bool {
	return p.preserveState(func(p *Parser) bool {

		if !p.parseOptional(token.For) {
			return false
		}
		p.parseOptional(token.Identifier)
		for p.parseOptional(token.Comma) {
			p.parseOptional(token.Identifier)
		}
		return p.isNextToken(token.In)
	})
}
