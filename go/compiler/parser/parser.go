package parser

import (
	"github.com/end-r/guardian/go/compiler/ast"
	"github.com/end-r/guardian/go/compiler/lexer"
)

// Parser ...
type Parser struct {
	Scope      ast.Node
	Expression ast.ExpressionNode
	parent     ast.Node
	lexer      *lexer.Lexer
	index      int
	Errs       []string
}

func createParser(data string) *Parser {
	p := new(Parser)
	p.lexer = lexer.LexString(data)
	p.Scope = ast.FileNode{}
	return p
}

func (p *Parser) run() {
	if p.index >= len(p.lexer.Tokens) {
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
		//p.addError(fmt.Sprintf(errUnrecognisedConstruct, p.lexer.Tokenstring(p.current())))
		p.next()
	}
	p.run()
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

func (p *Parser) parseOptional(t lexer.TokenType) bool {
	if p.current().Type == t {
		p.next()
		return true
	}
	return false
}

func (p *Parser) parseRequired(t lexer.TokenType) {
	if p.lexer.Tokens[p.index].Type != t {
		p.addError("Required x, found y")
	}
	p.next()
}

func (p *Parser) parseIdentifier() string {
	if p.lexer.Tokens[p.index].Type != lexer.TknIdentifier {
		p.addError("Required indentifier, found y")
		return ""
	}
	s := p.lexer.TokenString(p.lexer.Tokens[p.index])
	p.next()
	return s
}

func (p *Parser) validate(t ast.NodeType) {
	if p.Scope != nil {
		if !p.Scope.Validate(t) {
			p.addError("Invalid declaration in Scope")
		}
	}
}

func (p *Parser) parseType() ast.Node {
	return ast.TypeDeclarationNode{}
}

func (p *Parser) addError(err string) {
	p.Errs = append(p.Errs, err)
}
