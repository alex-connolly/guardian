package parser

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/compiler/lexer"
)

type parser struct {
	scope ast.Node
	lexer *lexer.Lexer
	index int
	errs  []string
}

func createParser(data string) *parser {
	p := new(parser)
	p.lexer = lexer.LexString(data)
	p.scope = ast.FileNode{}
	return p
}

func (p *parser) run() {
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

func (p *parser) current() lexer.Token {
	return p.token(0)
}

func (p *parser) next() {
	p.index++
}

func (p *parser) token(offset int) lexer.Token {
	return p.lexer.Tokens[p.index+offset]
}

func (p *parser) parseOptional(t lexer.TokenType) bool {
	if p.current().Type == t {
		p.next()
		return true
	}
	return false
}

func (p *parser) parseRequired(t lexer.TokenType) {
	if p.lexer.Tokens[p.index].Type != t {
		p.addError("Required x, found y")
	}
	p.next()
}

func (p *parser) parseIdentifier() string {
	if p.lexer.Tokens[p.index].Type != lexer.TknIdentifier {
		p.addError("Required indentifier, found y")
		return ""
	}
	s := p.lexer.TokenString(p.lexer.Tokens[p.index])
	p.next()
	return s
}

func (p *parser) validate(t ast.NodeType) {
	if p.scope != nil {
		if !p.scope.Validate(t) {
			p.addError("Invalid declaration in scope")
		}
	}
}

func (p *parser) parseType() ast.Node {
	return ast.TypeDeclarationNode{}
}

func (p *parser) addError(err string) {
	p.errs = append(p.errs, err)
}
