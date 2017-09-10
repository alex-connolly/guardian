package parser

import (
	"io/ioutil"

	"github.com/end-r/guardian/compiler/ast"
	"github.com/end-r/guardian/compiler/lexer"
)

// ParseFile ...
func ParseFile(path string) *Parser {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		//log.Println(err)
		return nil
	}
	p := new(Parser)
	p.lexer = lexer.LexBytes(bytes)
	p.Scope = &ast.ScopeNode{
		ValidTypes: []ast.NodeType{
			ast.ContractDeclaration,
		},
	}
	p.parseScope(p.Scope)
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
	p.Scope = &ast.ScopeNode{
		ValidTypes: []ast.NodeType{
			ast.ContractDeclaration,
		},
	}
	p.parseScope(p.Scope)
	return p
}
