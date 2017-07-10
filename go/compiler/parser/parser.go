package parser

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/compiler/lexer"
)

type parser struct {
	constructs []construct
	scope      ast.Node
	tokens     []lexer.token
	index      int
	errs       []string
}

func (p *parser) current() lexer.Token {
	return p.tokens[p.index]
}

func (p *parser) parseOptional(t lexer.TknType) bool {
	if p.current().Type == t {
		p.next()
		return true
	}
	return false
}

func (p *parser) parseRequired(t lexer.TknType) {

}

func (p *parser) parseIdentifier() string {

}

func (p.parser) error() {

}
