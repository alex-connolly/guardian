package parser

import (
	"io/ioutil"

	"github.com/end-r/guardian/compiler/ast"
	"github.com/end-r/guardian/compiler/lexer"
)

// ParseFile ...
func ParseFile(path string) *Parser {
	p := new(Parser)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		p.addError("Unable to read file")
		return nil
	}
	p.lexer = lexer.LexBytes(bytes)
	p.parseScope(ast.ContractDeclaration)
	return p
}

// ParseString ...
func ParseString(data string) *Parser {
	return ParseBytes([]byte(data))
}

// ParseBytes ...
func ParseBytes(data []byte) *Parser {
	p := new(Parser)
	p.Scope = new(ast.ScopeNode)
	p.lexer = lexer.LexBytes(data)
	p.Scope = p.parseScope(ast.ContractDeclaration)
	return p
}
