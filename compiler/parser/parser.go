package parser

import (
	"fmt"

	"github.com/end-r/guardian/compiler/ast"
	"github.com/end-r/guardian/compiler/lexer"
)

const (
	classKey       = "class"
	contractKey    = "contract"
	enumKey        = "enum"
	flowKey        = "flow"
	interfaceKey   = "interface"
	varKey         = "var"
	eventKey       = "event"
	constructorKey = "constructor"
	funcKey        = "func"
	typeKey        = "type"
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

func (p *Parser) parseRequired(types ...lexer.TokenType) bool {
	if !p.hasTokens(1) {
		return false
	}
	for _, t := range types {
		if p.current().Type == t {
			p.next()
			return true
		}
	}
	p.addError(fmt.Sprintf("Required %s, found %d", "x", p.current().Type))
	return false
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

func (p *Parser) parseEnclosedScope(scope *ast.ScopeNode) {
	p.parseRequired(lexer.TknOpenBrace)
	p.parseScope(scope)
	p.parseRequired(lexer.TknCloseBrace)
}

func (p *Parser) parseScope(scope *ast.ScopeNode) {

	scope.Parent = p.Scope
	p.Scope = scope

	for p.hasTokens(1) {
		if p.current().Type == lexer.TknCloseBrace {
			return
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
			//fmt.Printf("Unrecognised construct at index %d: %s\n", p.index, p.lexer.TokenString(p.current()))
			p.addError(fmt.Sprintf("unrecognised construct: %s", p.lexer.TokenString(p.current())))
			p.next()
		}
	}
}
