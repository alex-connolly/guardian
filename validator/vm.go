package validator

import (
	"math/big"
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
	n := new(big.Int)
	n, ok := n.SetString("314159265358979323846264338327950288419716939937510582097494459", 10)
	if !ok {

	}
	return v.smallestNumericType(n.BitLen())
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
		intName := "int" + strconv.Itoa(i)
		uintName := "u" + intName
		m[uintName] = NumericType{name: uintName, size: i, signed: false, integer: true}
		m[intName] = NumericType{name: intName, size: i, signed: true, integer: true}
	}
	m["int"] = NumericType{name: "int", size: maxSize, signed: false, integer: true}
	m["uint"] = NumericType{name: "int", size: maxSize, signed: true, integer: true}
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
	m.Add(operatorAdd, lexer.TknAdd)
	m.Add(booleanOperator, lexer.TknAnd, lexer.TknOr)

	// numericalOperator with floats/ints
	m.Add(BinaryNumericOperator(), lexer.TknSub, lexer.TknMul, lexer.TknDiv)

	// integers only
	m.Add(BinaryIntegerOperator(), lexer.TknShl, lexer.TknShr)

	m.Add(castOperator, lexer.TknAs)

	return m
}

func castOperator(v *Validator, ts ...Type) Type {
	// pretend it's a valid type
	left := ts[0]
	right := ts[1]
	if !assignableTo(left, right) {
		v.addError(errImpossibleCast, WriteType(left), WriteType(right))
		return left
	}
	return right
}

func booleanOperator(v *Validator, ts ...Type) Type {
	if len(ts) != 2 {

	}
	return standards[Bool]
}

func operatorAdd(v *Validator, ts ...Type) Type {
	switch ts[0].(type) {
	case NumericType:
		// TODO: make this meaningful
		return v.smallestNumericType(0)
	}
	strType := v.getNamedType("string")
	if ts[0].compare(strType) {
		return strType
	}
	return standards[Invalid]
}
