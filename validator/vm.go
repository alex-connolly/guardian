package validator

import (
	"strconv"

	"github.com/end-r/guardian/lexer"

	"github.com/end-r/guardian/parser"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/vmgen"
)

// A VM is the mechanism through which all vm-specific features are applied
// to the Guardian AST: bytecode generation, type enforcement etc
type VM interface {
	Traverse(ast.Node) vmgen.Bytecode
	Builtins() *ast.ScopeNode
	Primitives() map[string]Type
	Literals() LiteralMap
	Operators() OperatorMap
}

type TestVM struct {
}

func NewTestVM() TestVM {
	return TestVM{}
}

func (v TestVM) Traverse(ast.Node) vmgen.Bytecode {
	return vmgen.Bytecode{}
}

func (v TestVM) Builtins() *ast.ScopeNode {
	ast, _ := parser.ParseString(`
    		type byte uint8
            type string []byte
            type address [20]byte
    	`)
	return ast
}

func (v TestVM) Literals() LiteralMap {
	return LiteralMap{
		lexer.TknString:  SimpleLiteral("string"),
		lexer.TknTrue:    SimpleLiteral("bool"),
		lexer.TknFalse:   SimpleLiteral("bool"),
		lexer.TknInteger: resolveIntegerLiteral,
		lexer.TknFloat:   resolveFloatLiteral,
	}
}

func resolveIntegerLiteral(v *Validator, data string) Type {
	i, err := strconv.ParseInt(data, 10, 64)
	if err != nil {

	}
	return v.smallestNumericType(int(i))
}

func resolveFloatLiteral(v *Validator, data string) Type {
	// convert to float
	return standards[Unknown]
}

func getIntegerTypes() map[string]Type {
	m := map[string]Type{}
	const maxSize = 256
	const increment = 8
	for i := increment; i <= maxSize; i += increment {
		m["uint"+string(i)] = NumericType{size: i, signed: false, integer: true}
		m["int"+string(i)] = NumericType{size: i, signed: true, integer: true}
	}
	m["int"] = NumericType{size: maxSize, signed: false, integer: true}
	m["uint"] = NumericType{size: maxSize, signed: true, integer: true}
	return m
}

func (v TestVM) Primitives() map[string]Type {
	it := getIntegerTypes()

	s := map[string]Type{
		"bool": BooleanType{},
	}

	for k, v := range it {
		s[k] = v
	}
	return s
}

func (v TestVM) Operators() (m OperatorMap) {
	m = OperatorMap{}
	m.Add(SimpleOperator("bool"), lexer.TknGeq, lexer.TknLeq,
		lexer.TknLss, lexer.TknNeq, lexer.TknEql, lexer.TknGtr)

	return m
}
