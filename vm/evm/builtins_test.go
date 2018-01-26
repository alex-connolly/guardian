package evm

import (
	"testing"

	"github.com/end-r/guardian/validator"

	"github.com/end-r/goutil"
)

func TestBuiltinRequire(t *testing.T) {
	e := NewVM()
	a, errs := validator.ValidateExpression(e, "require(5 > 3)")
	goutil.AssertNow(t, len(errs) == 0, errs.Format())

	code := e.traverseExpression(a)
	expected := []string{
		"PUSH1",
		"PUSH1",
		"LT",
		"PUSH1",
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

func TestBuiltinArrayLength(t *testing.T) {
	e := NewVM()
	a, errs := validator.ValidateExpression(e, `len("hello")`)
	goutil.AssertNow(t, errs == nil, errs.Format())
	code := e.traverseExpression(a)
	expected := []string{
		"PUSH7",
		"MLOAD",
	}
	goutil.Assert(t, code.CompareMnemonics(expected), code.Format())
}

func TestBuiltinAddmod(t *testing.T) {
	e := NewVM()
	a, errs := validator.ValidateExpression(e, "addmod(uint(1), uint(2), uint(3))")
	goutil.AssertNow(t, errs == nil, errs.Format())
	code := e.traverseExpression(a)
	expected := []string{
		"PUSH32",
		"PUSH32",
		"PUSH32",
		"ADDMOD",
	}
	goutil.Assert(t, code.CompareMnemonics(expected), code.Format())
}

func TestBuiltinMulmod(t *testing.T) {
	e := NewVM()
	a, errs := validator.ValidateExpression(e, "mulmod(uint(1), uint(2), uint(3))")
	goutil.AssertNow(t, errs == nil, errs.Format())
	code := e.traverseExpression(a)
	expected := []string{
		"PUSH32",
		"PUSH32",
		"PUSH32",
		"ADDMOD",
	}
	goutil.Assert(t, code.CompareMnemonics(expected), code.Format())
}

func TestBuiltinCall(t *testing.T) {
	e := NewVM()
	a, errs := validator.ValidateExpression(e, "call(0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF)")
	goutil.AssertNow(t, errs == nil, errs.Format())
	code := e.traverseExpression(a)
	expected := []string{
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"CALL",
	}
	goutil.Assert(t, code.CompareMnemonics(expected), code.Format())
}

func TestBuiltinDelegateCall(t *testing.T) {
	e := NewVM()
	a, errs := validator.ValidateExpression(e, "delegate(0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF)")
	goutil.AssertNow(t, errs == nil, errs.Format())
	code := e.traverseExpression(a)
	expected := []string{
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"CALL",
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
	a, errs := validator.ValidateExpression(e, `revert()`)
	goutil.AssertNow(t, errs == nil, errs.Format())
	code := e.traverseExpression(a)
	expected := []string{
		"REVERT",
	}
	goutil.Assert(t, code.CompareMnemonics(expected), code.Format())
}

func TestCalldata(t *testing.T) {
	e := new(GuardianEVM)
	expr, _ := validator.ValidateExpression(e, "msg.data")
	bytecode := e.traverseExpression(expr)
	expected := []string{"CALLDATA"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestGas(t *testing.T) {
	e := new(GuardianEVM)
	expr, _ := validator.ValidateExpression(e, "msg.gas")
	bytecode := e.traverseExpression(expr)
	expected := []string{"GAS"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestCaller(t *testing.T) {
	e := new(GuardianEVM)
	expr, _ := validator.ValidateExpression(e, "msg.sender")
	bytecode := e.traverseExpression(expr)
	expected := []string{"CALLER"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestSignature(t *testing.T) {
	e := new(GuardianEVM)
	expr, _ := validator.ValidateExpression(e, "msg.sig")
	bytecode := e.traverseExpression(expr)
	// should get first 4 bytes of calldata
	expected := []string{"CALLDATA"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTimestamp(t *testing.T) {
	e := new(GuardianEVM)
	expr, _ := validator.ValidateExpression(e, "block.timestamp")
	bytecode := e.traverseExpression(expr)
	// should get first 4 bytes of calldata
	expected := []string{"TIMESTAMP"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestNumber(t *testing.T) {
	e := new(GuardianEVM)
	expr, _ := validator.ValidateExpression(e, "block.timestamp")
	bytecode := e.traverseExpression(expr)
	// should get first 4 bytes of calldata
	expected := []string{"NUMBER"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestCoinbase(t *testing.T) {
	e := new(GuardianEVM)
	expr, _ := validator.ValidateExpression(e, "block.timestamp")
	bytecode := e.traverseExpression(expr)
	// should get first 4 bytes of calldata
	expected := []string{"COINBASE"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestGasLimit(t *testing.T) {
	e := new(GuardianEVM)
	expr, _ := validator.ValidateExpression(e, "block.gasLimit")
	bytecode := e.traverseExpression(expr)
	// should get first 4 bytes of calldata
	expected := []string{"GASLIMIT"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestBlockhash(t *testing.T) {
	e := new(GuardianEVM)
	expr, _ := validator.ValidateExpression(e, "block.gasLimit")
	bytecode := e.traverseExpression(expr)
	// should get first 4 bytes of calldata
	expected := []string{"GASLIMIT"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestGasPrice(t *testing.T) {
	e := new(GuardianEVM)
	expr, _ := validator.ValidateExpression(e, "tx.gasPrice")
	bytecode := e.traverseExpression(expr)
	// should get first 4 bytes of calldata
	expected := []string{"GASPRICE"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestOrigin(t *testing.T) {
	e := new(GuardianEVM)
	expr, _ := validator.ValidateExpression(e, "tx.origin")
	bytecode := e.traverseExpression(expr)
	// should get first 4 bytes of calldata
	expected := []string{"ORIGIN"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTransfer(t *testing.T) {
	e := new(GuardianEVM)
	expr, _ := validator.ValidateExpression(e, `
		transfer(0x123f681646d4a755815f9cb19e1acc8565a0c2ac, 1000)
	`)
	bytecode := e.traverseExpression(expr)
	// should get first 4 bytes of calldata
	expected := []string{
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"CALL",
	}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestCall(t *testing.T) {
	e := new(GuardianEVM)
	expr, _ := validator.ValidateExpression(e, `
		Call(0x123f681646d4a755815f9cb19e1acc8565a0c2ac).gas(1000).value(1000).sig("aaaaa").call()
	`)
	bytecode := e.traverseExpression(expr)
	// should get first 4 bytes of calldata
	expected := []string{
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"CALL",
	}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}
