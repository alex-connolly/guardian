package parser

import (
	"io/ioutil"

	"github.com/end-r/guardian/compiler/ast"
	"github.com/end-r/guardian/compiler/lexer"
)

// ParseExpression ...
func ParseExpression(expr string) ast.ExpressionNode {
	p := new(Parser)
	p.lexer = lexer.LexString(expr)
	return p.parseExpression()
}

// ParseFile ...
func ParseFile(path string) *Parser {
	p := new(Parser)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		p.addError("Unable to read file")
		return nil
	}
	p.lexer = lexer.LexBytes(bytes)
	// TODO: update the need for the scope to have this
	p.parseScope(lexer.TknCloseBrace, ast.ContractDeclaration)
	return p
}

// ParseString ...
func ParseString(data string) *Parser {
	return ParseBytes([]byte(data))
}

// ParseBytes ...
func ParseBytes(data []byte) *Parser {
	p := new(Parser)
	p.lexer = lexer.LexBytes(data)
	p.Scope = p.parseScope(lexer.TknCloseBrace, ast.ContractDeclaration)
	return p
}
