package parser

import (
	"io/ioutil"
	"log"

	"github.com/end-r/guardian/go/compiler/ast"
	"github.com/end-r/guardian/go/compiler/lexer"
)

// ParseFile ...
func ParseFile(path string) *Parser {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
		return nil
	}
	p := new(Parser)
	p.lexer = lexer.LexBytes(bytes)
	p.Scope = &ast.ScopeNode{}
	p.run()
	return p
}

// ParseString ...
func ParseString(data string) *Parser {
	return ParseBytes([]byte(data))
}

// ParseBytes ...
func ParseBytes(data []byte) *Parser {
	p := new(Parser)
	p.Scope = &ast.ScopeNode{}
	p.lexer = lexer.LexBytes(data)
	p.run()
	return p
}
