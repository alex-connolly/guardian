package parser

import (
	"github.com/end-r/guardian/ast"

	"github.com/end-r/guardian/token"
)

func isNewLine(p *Parser) bool {
	return p.isNextToken(token.NewLine)
}

// e.g. name string
func isExplicitVarDeclaration(p *Parser) bool {
	return p.isNextToken(token.Var, token.Const)
}

func (p *Parser) isNextTerminating() bool {
	return p.parseOptional(token.NewLine, token.Semicolon, token.Comma, token.CloseBracket)
}

func (p *Parser) isNextAType() bool {
	return p.isFuncType() || p.isPlainType() || p.isArrayType() || p.isMapType()
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

func (p *Parser) nextTokens(tokens ...token.Type) bool {
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

func isClassDeclaration(p *Parser) bool {
	return p.isNextToken(token.Class)
}

func isInterfaceDeclaration(p *Parser) bool {
	return p.isNextToken(token.Interface)
}

func isLifecycleDeclaration(p *Parser) bool {
	return p.isNextToken(token.GetLifecycles()...)
}

func isEnumDeclaration(p *Parser) bool {
	return p.isNextToken(token.Enum)
}

func isContractDeclaration(p *Parser) bool {
	return p.isNextToken(token.Contract)
}

func isFuncDeclaration(p *Parser) bool {
	return p.isNextToken(token.Func)
}

func isGroup(p *Parser) bool {
	return p.isNextToken(token.OpenBracket)
}

func (p *Parser) isNextToken(types ...token.Type) bool {
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
	return p.isNextToken(token.Event)
}

func isTypeDeclaration(p *Parser) bool {
	return p.isNextToken(token.KWType)
}

func isAnnotation(p *Parser) bool {
	return p.isNextToken(token.At)
}

func isForStatement(p *Parser) bool {
	return p.isNextToken(token.For)
}

func isIfStatement(p *Parser) bool {
	return p.isNextToken(token.If)
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
		/*for p.parseOptional(token.GetModifiers()...) {
		}*/
		expr := p.parseSimpleExpression()
		if expr == nil {
			return false
		}
		for p.parseOptional(token.Comma) {
			// assume these will be expressions
			p.parseSimpleExpression()
		}
		t := p.isNextTokenAssignment()
		return t
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

func isSwitchStatement(p *Parser) bool {
	ex := p.nextTokens(token.Exclusive, token.Switch)
	dir := p.nextTokens(token.Switch)
	return ex || dir
}

func isIgnored(p *Parser) bool {
	return p.isNextToken(token.LineComment, token.MultilineComment)
}

func isReturnStatement(p *Parser) bool {
	return p.isNextToken(token.Return)
}

func isCaseStatement(p *Parser) bool {
	return p.isNextToken(token.Case)
}

func (p *Parser) isRecursiveModifier() bool {
	if p.hasTokens(1) {
		switch p.current().Type {
		case token.Func, token.Var, token.Const, token.Enum,
			token.Interface, token.Contract, token.Class, token.Event:
			return true
		case token.Identifier:
			p.next()
			return p.isRecursiveModifier()
		case token.OpenBracket:
			p.next()
			p.ignoreNewLines()
			return p.isRecursiveModifier()
		}
		return false
	}
	return false
}

func isModifier(p *Parser) bool {
	// have to deal with nested calls
	// assert(now(assert(now())))
	return p.preserveState(func(p *Parser) bool {
		if !p.parseOptional(token.Identifier) {
			return false
		}
		for p.parseOptional(token.Identifier) {
		}
		return p.isRecursiveModifier()
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
