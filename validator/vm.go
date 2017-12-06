package validator

import (
	"strconv"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/lexer"
	"github.com/end-r/guardian/parser"
	"github.com/end-r/guardian/typing"
	"github.com/end-r/guardian/util"
	"github.com/end-r/vmgen"
)

// A VM is the mechanism through which all vm-specific features are applied
// to the Guardian AST: bytecode generation, type enforcement etc
type VM interface {
	Traverse(ast.Node) (vmgen.Bytecode, util.Errors)
	Builtins() *ast.ScopeNode
	Primitives() map[string]typing.Type
	Literals() LiteralMap
	Operators() OperatorMap
	BooleanName() string
}

type TestVM struct {
}

func NewTestVM() TestVM {
	return TestVM{}
}

func (v TestVM) BooleanName() string {
	return "bool"
}

func (v TestVM) Traverse(ast.Node) (vmgen.Bytecode, util.Errors) {
	return vmgen.Bytecode{}, nil
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
		lexer.TknTrue:    BooleanLiteral,
		lexer.TknFalse:   BooleanLiteral,
		lexer.TknInteger: resolveIntegerLiteral,
		lexer.TknFloat:   resolveFloatLiteral,
	}
}

func resolveIntegerLiteral(v *Validator, data string) typing.Type {
	x := typing.BitsNeeded(len(data))
	return v.SmallestNumericType(x, false)
}

func resolveFloatLiteral(v *Validator, data string) typing.Type {
	// convert to float
	return typing.Unknown()
}

func getIntegerTypes() map[string]typing.Type {
	m := map[string]typing.Type{}
	const maxSize = 256
	const increment = 8
	for i := increment; i <= maxSize; i += increment {
		intName := "int" + strconv.Itoa(i)
		uintName := "u" + intName
		m[uintName] = typing.NumericType{Name: uintName, BitSize: i, Signed: false, Integer: true}
		m[intName] = typing.NumericType{Name: intName, BitSize: i, Signed: true, Integer: true}
	}
	m["int"] = typing.NumericType{Name: "int", BitSize: maxSize, Signed: false, Integer: true}
	m["uint"] = typing.NumericType{Name: "int", BitSize: maxSize, Signed: true, Integer: true}
	return m
}

func (v TestVM) Primitives() map[string]typing.Type {
	return getIntegerTypes()
}

func (v TestVM) Operators() (m OperatorMap) {
	m = OperatorMap{}

	m.Add(BooleanOperator, lexer.TknGeq, lexer.TknLeq,
		lexer.TknLss, lexer.TknNeq, lexer.TknEql, lexer.TknGtr)

	m.Add(operatorAdd, lexer.TknAdd)
	m.Add(BooleanOperator, lexer.TknLogicalAnd, lexer.TknLogicalOr)

	// numericalOperator with floats/ints
	m.Add(BinaryNumericOperator, lexer.TknSub, lexer.TknMul, lexer.TknDiv)

	// integers only
	m.Add(BinaryIntegerOperator, lexer.TknShl, lexer.TknShr)

	m.Add(CastOperator, lexer.TknAs)

	return m
}

func operatorAdd(v *Validator, ts ...typing.Type) typing.Type {
	switch ts[0].(type) {
	case typing.NumericType:
		return BinaryNumericOperator(v, ts...)
	}
	strType := v.getNamedType("string")
	if ts[0].Compare(strType) {
		return strType
	}
	return typing.Invalid()
}
