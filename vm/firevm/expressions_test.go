package firevm

import (
	"testing"

	"github.com/end-r/goutil"
	"github.com/end-r/guardian"
)

func TestBinaryExpressionBytecodeLiterals(t *testing.T) {
	a := new(Arsonist)
	guardian.CompileString(a, "x = 1 + 5")
	goutil.AssertNow(t, a.VM != nil, "vm shouldn't be nil")
	checkMnemonics(t, a.VM.Instructions, []string{
		"PUSH", // push string data
		"PUSH", // push hash(x)
		"ADD",
	})
	checkStack(t, a.VM.Stack, [][]byte{
		[]byte{byte(6)},
	})
}

func TestBinaryExpressionBytecodeReferences(t *testing.T) {
	a := new(Arsonist)
	guardian.CompileString(a, "a + b")
	goutil.AssertNow(t, a.VM != nil, "vm shouldn't be nil")
	checkMnemonics(t, a.VM.Instructions, []string{
		"PUSH", // push string data
		"PUSH", // push hash(x)
		"ADD",
	})
}

func TestBinaryExpressionBytecodeStringLiteralConcat(t *testing.T) {
	a := new(Arsonist)
	guardian.CompileString(a, `"my name is" + " who knows tbh"`)
	goutil.AssertNow(t, a.VM != nil, "vm shouldn't be nil")
	checkMnemonics(t, a.VM.Instructions, []string{
		"PUSH",   // push string data
		"PUSH",   // push hash(x)
		"CONCAT", // concatenate bytes
	})
}

func TestBinaryExpressionBytecodeStringReferenceConcat(t *testing.T) {
	a := new(Arsonist)
	guardian.CompileString(a, `
		var (
			a = "hello"
			b = "world"
			c = a + b
		)
		`)
	goutil.AssertNow(t, a.VM != nil, "vm shouldn't be nil")
	checkMnemonics(t, a.VM.Instructions, []string{
		"PUSH",   // push string data
		"PUSH",   // push a reference
		"STORE",  // store at a
		"PUSH",   // push string data
		"PUSH",   // push b reference
		"PUSH",   // push a reference
		"LOAD",   // load a data
		"PUSH",   // push b reference
		"LOAD",   // load b data
		"CONCAT", //
	})
}

func TesExtendedBinaryExpressionBytecodeStringReferenceConcat(t *testing.T) {
	a := new(Arsonist)
	goutil.AssertNow(t, a.VM != nil, "vm shouldn't be nil")
	guardian.CompileString(a, `
		var (
			a = "hello"
			b = "world"
			c = "www"
			d = "comma"
		)
		a + b + c + d
		`)
	checkMnemonics(t, a.VM.Instructions, []string{
		"PUSH",   // push hello
		"PUSH",   // push world
		"CONCAT", //
		"PUSH",   // push www
		"CONCAT", //
		"PUSH",   // push comma
		"CONCAT", //
	})
}

func TestUnaryExpressionBytecodeLiteral(t *testing.T) {
	a := new(Arsonist)
	guardian.CompileString(a, "!1")
	checkMnemonics(t, a.VM.Instructions, []string{
		"PUSH", // push data
		"NOT",
	})
}

func TestCallExpressionBytecodeLiteral(t *testing.T) {
	a := new(Arsonist)
	guardian.CompileString(a, `
		doSomething("data")
		`)
	checkMnemonics(t, a.VM.Instructions, []string{
		"PUSH", // push string data
		"NOT",
	})
}

func TestCallExpressionBytecodeUseResult(t *testing.T) {
	a := new(Arsonist)
	guardian.CompileString(a, `
		s := doSomething("data")
		`)
	checkMnemonics(t, a.VM.Instructions, []string{
		"PUSH", // push string data
		"NOT",
	})
}

func TestCallExpressionBytecodeUseMultipleResults(t *testing.T) {
	a := new(Arsonist)
	guardian.CompileString(a, `
		s, a, p := doSomething("data")
		`)
	checkMnemonics(t, a.VM.Instructions, []string{
		"PUSH", // push string data
		// this is where the function code would go
		"PUSH", // push p
		"SET",  // set p
		"PUSH", // push a
		"SET",  // set a
		"PUSH", // push s
		"SET",  // set s
	})
}

func TestCallExpressionBytecodeIgnoredResult(t *testing.T) {
	a := new(Arsonist)
	guardian.CompileString(a, `
		s, _, p := doSomething("data")
		`)
	checkMnemonics(t, a.VM.Instructions, []string{
		"PUSH", // push string data
		// this is where the function code would go
		// stack ends: {R3, R2, R1}
		"PUSH", // push p
		"SET",  // set p
		"POP",  // ignore second return value
		"PUSH", // push s
		"SET",  // set s
	})
}

func TestCallExpressionBytecodeNestedCall(t *testing.T) {
	a := new(Arsonist)
	guardian.CompileString(a, `
		err := saySomething(doSomething("data"))
		`)
	checkMnemonics(t, a.VM.Instructions, []string{
		"PUSH", // push string data
		// this is where the function code would go
		// stack ends: {R1}
		// second function code goes here
		"PUSH", // push err
		"SET",  // set err
	})
}
