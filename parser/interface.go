package parser

import (
	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/lexer"
	"github.com/end-r/guardian/util"
)

// Parse ...
func Parse(lexer *lexer.Lexer) (*ast.ScopeNode, util.Errors) {
	p := new(Parser)
	if lexer.Errors != nil {
		return nil, lexer.Errors
	}
	p.lexer = lexer
	p.line = 1
	p.seenCastOperator = false
	p.parseScope([]token.Type{token.CloseBrace}, ast.ContractDeclaration)
	return p.scope, p.errs
}

// ParseExpression ...
func ParseExpression(expr string) ast.ExpressionNode {
	p := new(Parser)
	p.lexer = lexer.LexString(expr)
	return p.parseExpression()
}

// ParseType ...
func ParseType(t string) ast.Node {
	p := new(Parser)
	p.lexer = lexer.LexString(t)
	return p.parseType()
}

// ParseFile ...
func ParseFile(path string) (scope *ast.ScopeNode, errs util.Errors) {
	l := lexer.LexFile(path)
	return Parse(l)
}

// ParseString ...
func ParseString(data string) (scope *ast.ScopeNode, errs util.Errors) {
	return ParseBytes([]byte(data))
}

// ParseBytes ...
func ParseBytes(data []byte) (scope *ast.ScopeNode, errs util.Errors) {
	lexer := lexer.Lex("input", data)
	return Parse(lexer)
}
