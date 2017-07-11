package parser

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/compiler/lexer"
)

type parser struct {
	scope  ast.Node
	tokens []lexer.Token
	index  int
	errs   []string
}

func createParser(data string) *parser {
	p := new(parser)
	tokens, errs := lexer.LexString(data)
	p.tokens = append(p.tokens, tokens...)
	p.errs = append(p.errs, errs...)
	return p
}

func (p *parser) run() {
	if p.index >= len(p.tokens) {
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
		//p.addError(fmt.Sprintf(errUnrecognisedConstruct, p.lexer.tokenString(p.current())))
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
	return p.tokens[p.index+offset]
}

func (p *parser) parseOptional(t lexer.TokenType) bool {
	if p.current().Type == t {
		p.next()
		return true
	}
	return false
}

func (p *parser) parseRequired(t lexer.TokenType) {

}

func (p *parser) parseIdentifier() string {
	return "hi"
}

func (p *parser) validate(t ast.NodeType) {
	if !p.scope.Validate(t) {
		p.addError("Invalid declaration in scope")
	}
}

func (p *parser) parseType() ast.Node {
	return ast.TypeDeclarationNode{}
}

func (p *parser) addError(err string) {
	p.errs = append(p.errs, err)
}
