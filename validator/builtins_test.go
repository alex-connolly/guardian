package validator

import (
	"axia/guardian/parser"
	"fmt"
	"testing"

	"github.com/end-r/guardian/lexer"

	"github.com/end-r/goutil"
)

func TestAdd(t *testing.T) {
	m := OperatorMap{}

	goutil.Assert(t, len(m) == 0, "wrong initial length")
	// numericalOperator with floats/ints

	m.Add(BinaryNumericOperator, lexer.TknSub, lexer.TknMul, lexer.TknDiv)

	goutil.Assert(t, len(m) == 3, fmt.Sprintf("wrong added length: %d", len(m)))

	// integers only
	m.Add(BinaryIntegerOperator, lexer.TknShl, lexer.TknShr)

	goutil.Assert(t, len(m) == 5, fmt.Sprintf("wrong final length: %d", len(m)))
}

func TestParseBuiltinsVariables(t *testing.T) {
	v := new(Validator)
	ast, _ := parser.ParseString(`
		wei = 1
		kwei = 1000 * wei
		babbage = kwei
		mwei = 1000 * kwei
		lovelace = mwei
		gwei = 1000 * mwei
		shannon = gwei
		microether = 1000 * gwei
		szabo = microether
		milliether = 1000 * microether
		finney = milliether
		ether = 1000 * milliether
	`)
	tvm := NewTestVM()
	v.primitives = tvm.Primitives()
	v.builtinScope = ast
	v.parseBuiltins()
	goutil.AssertNow(t, v.builtinVariables != nil, "vars should not be nil")
	goutil.AssertNow(t, len(v.builtinVariables) == 12, "should be 12 vars")
}

func TestParseBuiltinsFunctions(t *testing.T) {
	v := new(Validator)
	ast, _ := parser.ParseString(`
		func sub(a, b int) int {
			return a - b
		}
	`)
	tvm := NewTestVM()
	v.primitives = tvm.Primitives()
	v.builtinScope = ast
	v.parseBuiltins()
	goutil.AssertNow(t, v.builtinVariables != nil, "vars should not be nil")
	goutil.AssertNow(t, len(v.builtinVariables) == 1, "should be 1 vars")
}

func TestParseBuiltinsPartialFunctions(t *testing.T) {
	v := new(Validator)
	ast, _ := parser.ParseString(`
		balance func(a address) uint256
		transfer func(a address, amount uint256) uint
		send func(a address, amount uint256) bool
		call func(a address) bool
		delegateCall func(a address)
	`)
	tvm := NewTestVM()
	v.primitives = tvm.Primitives()
	v.builtinScope = ast
	v.parseBuiltins()
	goutil.AssertNow(t, v.builtinVariables != nil, "vars should not be nil")
	goutil.AssertNow(t, len(v.builtinVariables) == 5, "should be 5 vars")
}
