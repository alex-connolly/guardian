package parser

import (
	"io/ioutil"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/lexer"
	"github.com/end-r/guardian/util"
)

func Parse(tokens []lexer.Token) (scope *ast.ScopeNode, errs util.Errors) {
	p := new(Parser)
	p.parseScope(lexer.TknCloseBrace, ast.ContractDeclaration)
	return p.Scope, errs
}

// ParseExpression ...
func ParseExpression(expr string) ast.ExpressionNode {
	p := new(Parser)
	p.tokens, _ = lexer.LexString(expr)
	return p.parseExpression()
}

// ParseFile ...
func ParseFile(path string) (scope *ast.ScopeNode, errs []util.Error) {
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
func ParseString(data string) (scope *ast.ScopeNode, errs []util.Error) {
	return ParseBytes([]byte(data))
}

// ParseBytes ...
func ParseBytes(data []byte) (scope *ast.ScopeNode, errs []util.Error) {
	tokens, errs := lexer.Lex(data)
	if errs != nil {
		return nil, errs
	}
	return Parse(tokens)
}
