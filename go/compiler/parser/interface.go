package parser

import (
	"axia/guardian/go/compiler/ast"
	"axia/guardian/go/compiler/lexer"
	"io/ioutil"
	"log"
)

// ParseFile ...
func ParseFile(path string) (p *Parser) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
		return nil
	}
	p.lexer = lexer.LexBytes(bytes)
	p.scope = ast.FileNode{}
	p.run()
	return p
}

// ParseString ...
func ParseString(data string) *Parser {
	return ParseBytes([]byte(data))
}

// ParseBytes ...
func ParseBytes(data []byte) (p *Parser) {
	p.lexer = lexer.LexBytes(data)
	p.run()
	return p
}
