package parser

import (
	"fmt"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/lexer"
)

// Parser ...
type Parser struct {
	Scope      *ast.ScopeNode
	Expression ast.ExpressionNode
	lexer      *lexer.Lexer
	modifiers  []lexer.TokenType
	index      int
	Errs       []Error
	line       int
	simple     bool
}

// An Error is
type Error struct {
	lineNumber int
	message    string
}

func createParser(data string) *Parser {
	p := new(Parser)
	p.lexer = lexer.LexString(data)
	p.Scope = &ast.ScopeNode{
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
	return p.lexer.Tokens[p.index+offset]
}

// index 0, tknlength 1
// hasTokens(0) yes always
// hasTokens(1) yes
// hasTokens(2) no
func (p *Parser) hasTokens(offset int) bool {
	return p.index+offset <= len(p.lexer.Tokens)
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
	s := p.lexer.TokenString(p.current())
	p.next()
	return s
}

func (p *Parser) validate(t ast.NodeType) {
	if p.Scope != nil {
		if !p.Scope.IsValid(t) {
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

func (p *Parser) formatErrors() string {
	whole := ""
	whole += fmt.Sprintf("%d errors\n", len(p.Errs))
	for _, e := range p.Errs {
		whole += fmt.Sprintf("Line %d: %s\n", e.lineNumber, e.message)
	}
	return whole
}

func (p *Parser) addError(message string) {
	err := Error{
		message:    message,
		lineNumber: p.line,
	}
	p.Errs = append(p.Errs, err)
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
	scope.Parent = p.Scope
	p.Scope = scope
	for p.hasTokens(1) {
		if p.current().Type == terminator {
			p.Scope = scope.Parent
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
				//fmt.Printf("Unrecognised construct at index %d: %s\n", p.index, p.lexer.TokenString(p.current()))
				p.addError(fmt.Sprintf("Unrecognised construct: %s", p.lexer.TokenString(p.current())))
				p.next()
			} else {
				// ?
				p.Scope.AddSequential(expr)
			}
		}
	}
	return scope
}
