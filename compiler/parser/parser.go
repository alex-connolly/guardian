package parser

import (
	"fmt"

	"github.com/end-r/guardian/compiler/ast"
	"github.com/end-r/guardian/compiler/lexer"
)

// Parser ...
type Parser struct {
	Scope      *ast.ScopeNode
	Expression ast.ExpressionNode
	lexer      *lexer.Lexer
	index      int
	Errs       []string
	line       int
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

func (p *Parser) parseOptional(t lexer.TokenType) bool {
	if !p.hasTokens(1) {
		return false
	}
	if p.current().Type == t {
		p.next()
		return true
	}
	return false
}

// TODO: clarify what this actually returns
func (p *Parser) parseRequired(types ...lexer.TokenType) lexer.TokenType {
	if !p.hasTokens(1) {
		p.addError(fmt.Sprintf("Required %s, found nothing", "x"))
		return p.current().Type
	}
	for _, t := range types {
		if p.current().Type == t {
			p.next()
			return t
		}
	}
	p.addError(fmt.Sprintf("Required %s, found %d", "x", p.current().Type))
	return p.current().Type
}

func (p *Parser) parseIdentifier() string {
	if !p.hasTokens(1) {
		p.addError("Required indentifier, nothing found")
		return ""
	}
	if p.current().Type != lexer.TknIdentifier {
		p.addError("Required identifier, found {add token type}")
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

func (p *Parser) addError(err string) {
	p.Errs = append(p.Errs, err)
}

func (p *Parser) parseEnclosedScope(valids ...ast.NodeType) *ast.ScopeNode {
	p.parseRequired(lexer.TknOpenBrace)
	scope := p.parseScope(valids...)
	p.parseRequired(lexer.TknCloseBrace)
	return scope
}

func (p *Parser) parseScope(valids ...ast.NodeType) *ast.ScopeNode {
	scope := new(ast.ScopeNode)
	scope.Parent = p.Scope
	p.Scope = scope
	for p.hasTokens(1) {
		if p.current().Type == lexer.TknCloseBrace {
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
			// try interpreting it as a call expression
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
