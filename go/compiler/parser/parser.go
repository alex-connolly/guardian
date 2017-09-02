package parser

import (
	"fmt"

	"github.com/end-r/guardian/go/compiler/ast"
	"github.com/end-r/guardian/go/compiler/lexer"
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

func createContractParser(data string) *Parser {
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

func (p *Parser) parseRequired(t lexer.TokenType) {
	if p.current().Type != t {
		p.addError("Required x, found y")
	}
	p.next()
}

func (p *Parser) parseIdentifier() string {
	fmt.Printf("index: %d, tknlength %d", len(p.lexer.Tokens), p.index)
	if p.current().Type != lexer.TknIdentifier {
		p.addError("Required indentifier, found {add token type}")
		return ""
	}
	s := p.lexer.TokenString(p.lexer.Tokens[p.index])
	p.next()
	return s
}

func parseScopeClosure(p *Parser) {
	p.Scope = p.Scope.Parent
	p.next()
}

func (p *Parser) validate(t ast.NodeType) {
	if p.Scope != nil {
		if !p.Scope.IsValid(t) {
			p.addError("Invalid declaration in Scope")
		}
	}
}

func parseNewLine(p *Parser) {
	p.line++
	p.next()
}

func (p *Parser) parseType() ast.Node {
	return ast.TypeDeclarationNode{}
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

	fmt.Println("PARSING SCOPE")

	scope.Parent = p.Scope
	p.Scope = scope

	for p.hasTokens(1) {
		found := false
		for _, c := range getPrimaryConstructs() {
			if c.is(p) {
				fmt.Printf("FOUND: %s at index %d\n", c.name, p.index)
				if c.name == "scope closure" {
					fmt.Println("CLOSING SCOPE")
					return
				}
				c.parse(p)
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("Unrecognised token at index %d\n", p.index)
			p.addError(fmt.Sprintf("unrecognised construct: %s", p.lexer.TokenString(p.current())))
			p.next()
		}
	}
}
