package parser

import (
	"fmt"

	"github.com/end-r/guardian/util"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/lexer"
)

// Parser ...
type Parser struct {
	scope      *ast.ScopeNode
	Expression ast.ExpressionNode
	tokens     []lexer.Token
	modifiers  []lexer.TokenType
	index      int
	errs       util.Errors
	line       int
	simple     bool
}

func createParser(data string) *Parser {
	p := new(Parser)
	p.tokens, _ = lexer.LexString(data)
	p.scope = &ast.ScopeNode{
		ValidTypes: []ast.NodeType{
			ast.InterfaceDeclaration, ast.ClassDeclaration,
			ast.FuncDeclaration,
		},
	}
	return p
}

func (p *Parser) current() lexer.Token {
	return p.token(0)
}

func (p *Parser) next() {
	p.index++
}

func (p *Parser) token(offset int) lexer.Token {
	return p.tokens[p.index+offset]
}

// index 0, tknlength 1
// hasTokens(0) yes always
// hasTokens(1) yes
// hasTokens(2) no
func (p *Parser) hasTokens(offset int) bool {
	return p.index+offset <= len(p.tokens)
}

func (p *Parser) parseOptional(types ...lexer.TokenType) bool {
	if !p.hasTokens(1) {
		return false
	}
	for _, t := range types {
		if p.current().Type == t {
			p.next()
			return true
		}
	}
	return false
}

// TODO: clarify what this actually returns
func (p *Parser) parseRequired(types ...lexer.TokenType) lexer.TokenType {
	if !p.hasTokens(1) {
		p.addError(fmt.Sprintf("Required %s, found nothing", "x"))
		// TODO: what should be returned here
		return lexer.TknReturn
	}
	for _, t := range types {
		if p.current().Type == t {
			p.next()
			return t
		}
	}
	p.addError(fmt.Sprintf("Required %s, found %s", "x", p.current().Name()))
	return p.current().Type
}

func (p *Parser) parseIdentifier() string {
	if !p.hasTokens(1) {
		p.addError("Required identifier, nothing found")
		return ""
	}
	if p.current().Type != lexer.TknIdentifier {
		p.addError(fmt.Sprintf("Required identifier, found %s", p.current().Name()))
		return ""
	}
	s := p.current().String()
	p.next()
	return s
}

func (p *Parser) validate(t ast.NodeType) {
	if p.scope != nil {
		if !p.scope.IsValid(t) {
			p.addError("Invalid declaration in scope")
		}
	}
}

func parseNewLine(p *Parser) {
	p.line++
	p.next()
}

func parseSingleLineComment(p *Parser) {
	p.parseRequired(lexer.TknLineComment)
	for p.hasTokens(1) && p.current().Type != lexer.TknNewLine {
		p.next()
	}
	p.parseOptional(lexer.TknNewLine)
}

func parseMultiLineComment(p *Parser) {
	p.parseRequired(lexer.TknCommentOpen)
	for p.hasTokens(1) && p.current().Type != lexer.TknCommentClose {
		p.next()
	}
	p.parseOptional(lexer.TknCommentClose)
}

func parseModifierList(p *Parser) {
	old := p.modifiers
	p.modifiers = p.parseModifiers(lexer.TknOpenBracket)
	p.parseEnclosedScope(lexer.TknOpenBracket, lexer.TknCloseBracket)
	p.modifiers = old
}

func (p *Parser) addError(message string) {
	err := util.Error{
		Message:    message,
		LineNumber: p.line,
	}
	p.errs = append(p.errs, err)
}

func (p *Parser) parseBracesScope(valids ...ast.NodeType) *ast.ScopeNode {
	return p.parseEnclosedScope(lexer.TknOpenBrace, lexer.TknCloseBrace, valids...)
}

func (p *Parser) parseEnclosedScope(opener, closer lexer.TokenType, valids ...ast.NodeType) *ast.ScopeNode {
	p.parseRequired(opener)
	scope := p.parseScope(closer, valids...)
	p.parseRequired(closer)
	return scope
}

func (p *Parser) parseScope(terminator lexer.TokenType, valids ...ast.NodeType) *ast.ScopeNode {
	scope := new(ast.ScopeNode)
	scope.Parent = p.scope
	p.scope = scope
	for p.hasTokens(1) {
		if p.current().Type == terminator {
			p.scope = scope.Parent
			return scope
		}
		found := false
		for _, c := range getPrimaryConstructs() {
			if c.is(p) {
				//fmt.Printf("FOUND: %s at index %d\n", c.name, p.index)
				c.parse(p)
				found = true
				break
			}
		}
		if !found {
			// try interpreting it as an expression
			saved := p.index
			expr := p.parseExpression()
			if expr == nil {
				p.index = saved
				//fmt.Printf("Unrecognised construct at index %d: %s\n", p.index, p.current().TokenString())
				p.addError(fmt.Sprintf("Unrecognised construct: %s", p.current().String()))
				p.next()
			} else {
				// ?

				p.scope.AddSequential(expr)
			}
		}
	}
	return scope
}
