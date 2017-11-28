package validator

import (
	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/lexer"
)

// NumericType... used to pass types into the s
type NumericType struct {
	size    int
	name    string
	signed  bool
	integer bool
}

type BooleanType struct {
}

func (v *Validator) validateBuiltinDeclarations(scope *ast.ScopeNode) {
	if scope.Declarations != nil {
		// order doesn't matter here
		for _, i := range scope.Declarations.Map() {
			// add in placeholders for all declarations
			v.validateDeclaration(i.(ast.Node))
		}
	}
}

func (v *Validator) validateBuiltinSequence(scope *ast.ScopeNode) {
	for _, node := range scope.Sequence {
		v.validate(node)
	}
}

type LiteralFunc func(*Validator) Type
type LiteralMap map[lexer.TokenType]LiteralFunc

func simpleLiteral(typeName string) LiteralFunc {
	return func(v *Validator) Type {
		return v.findVariable(typeName)
	}
}

type OperatorFunc func(*Validator, ...Type) Type
type OperatorMap map[lexer.TokenType]OperatorFunc

func (m OperatorMap) Add(function OperatorFunc, types ...lexer.TokenType) {
	for _, t := range types {
		m[t] = function
	}
}

func simpleOperator(typeName string) OperatorFunc {
	return func(v *Validator, types ...Type) Type {
		return v.findVariable(typeName)
	}
}
