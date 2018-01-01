package evm

import (
	"fmt"
	"testing"

	"github.com/end-r/guardian/validator"

	"github.com/end-r/goutil"
)

func TestBuiltinRequire(t *testing.T) {
	e := NewVM()
	a, errs := validator.ValidateExpression(e, "require(5 > 3)")
	goutil.AssertNow(t, len(errs) == 0, errs.Format())

	fmt.Println("here")
	code := e.traverseExpression(a)
	expected := []string{
		"PUSH1",
		"PUSH1",
		"LT",
		"PUSH",
		"JUMPI",
		"REVERT",
	}
	goutil.Assert(t, code.CompareMnemonics(expected), code.Format())
}

func TestBuiltinAssert(t *testing.T) {
	e := NewVM()
	a, errs := validator.ValidateExpression(e, "assert(5 > 3)")
	goutil.AssertNow(t, errs == nil, errs.Format())
	code := e.traverseExpression(a)
	expected := []string{
		"PUSH1",
		"PUSH1",
		"LT",
		"PUSH",
		"JUMPI",
		"INVALID",
	}
	goutil.Assert(t, code.CompareMnemonics(expected), code.Format())
}

func TestBuiltinAddmod(t *testing.T) {
	e := NewVM()
	a, errs := validator.ValidateExpression(e, "addmod(1, 2, 3)")
	goutil.AssertNow(t, errs == nil, errs.Format())
	code := e.traverseExpression(a)
	expected := []string{
		"PUSH1",
		"PUSH1",
		"PUSH1",
		"ADDMOD",
	}
	goutil.Assert(t, code.CompareMnemonics(expected), code.Format())
}

func TestBuiltinMulmod(t *testing.T) {
	e := NewVM()
	a, errs := validator.ValidateExpression(e, "addmod(1, 2, 3)")
	goutil.AssertNow(t, errs == nil, errs.Format())
	code := e.traverseExpression(a)
	expected := []string{
		"PUSH1",
		"PUSH1",
		"PUSH1",
		"ADDMOD",
	}
	goutil.Assert(t, code.CompareMnemonics(expected), code.Format())
}

func TestBuiltinBalance(t *testing.T) {
	e := NewVM()
	a, errs := validator.ValidateExpression(e, "balance(0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF)")
	goutil.AssertNow(t, errs == nil, errs.Format())
	code := e.traverseExpression(a)
	expected := []string{
		"PUSH1",
		"BALANCE",
	}
	goutil.Assert(t, code.CompareMnemonics(expected), code.Format())
}

func TestBuiltinSha3(t *testing.T) {
	e := NewVM()
	a, errs := validator.ValidateExpression(e, `sha3("hello")`)
	goutil.AssertNow(t, errs == nil, errs.Format())
	code := e.traverseExpression(a)
	expected := []string{
		"PUSH",
		"SHA3",
	}
	goutil.Assert(t, code.CompareMnemonics(expected), code.Format())
}

func TestBuiltinRevert(t *testing.T) {
	e := NewVM()
	a, errs := validator.ValidateExpression(e, `sha3("hello")`)
	goutil.AssertNow(t, errs == nil, errs.Format())
	code := e.traverseExpression(a)
	expected := []string{
		"PUSH",
		"SHA3",
	}
	goutil.Assert(t, code.CompareMnemonics(expected), code.Format())
}
