package parser

import (
	"io/ioutil"

	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/lexer"
	"github.com/end-r/guardian/util"
)

// Parse ...
func Parse(tokens []token.Token) (*ast.ScopeNode, util.Errors) {
	p := new(Parser)
	p.tokens = tokens
	p.line = 1
	p.parseScope(token.CloseBrace, ast.ContractDeclaration)
	return p.scope, p.errs
}

// ParseExpression ...
func ParseExpression(expr string) ast.ExpressionNode {
	p := new(Parser)
	p.tokens, _ = lexer.LexString(expr)
	return p.parseExpression()
}

// ParseFile ...
func ParseFile(path string) (scope *ast.ScopeNode, errs util.Errors) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		errs = append(errs, util.Error{
			Message: "Unable to read file",
		})
		return nil, errs
	}
	return ParseBytes(bytes)
}

// ParseString ...
func ParseString(data string) (scope *ast.ScopeNode, errs util.Errors) {
	return ParseBytes([]byte(data))
}

// ParseBytes ...
func ParseBytes(data []byte) (scope *ast.ScopeNode, errs util.Errors) {
	tokens, es := lexer.Lex(data)
	if es != nil {
		// do something
	}
	return Parse(tokens)
}
