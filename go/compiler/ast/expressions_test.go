package ast

import (
	"testing"

	"github.com/end-r/firevm"
	"github.com/end-r/goutil"
)

func TestBinaryExpressionBytecodeLiterals(t *testing.T) {
	p := ParseString("1 + 5")
	vm := firevm.Create()
	Traverse(vm, p.Scope)
	checkMnemonics(t, vm.Instructions, []string{
		"PUSH", // push string data
		"PUSH", // push hash(x)
		"ADD"
	})
}

func TestBinaryExpressionBytecodeReferences(t *testing.T) {
	p := ParseString("a + b")
	vm := firevm.Create()
	Traverse(vm, p.Scope)
	checkMnemonics(t, vm.Instructions, []string{
		"PUSH", // push string data
		"PUSH", // push hash(x)
		"ADD"
	})
}

func TestBinaryExpressionBytecodeStringLiteralConcat(t *testing.T) {
	p := ParseString(`"my name is" + " who knows tbh"`)
	vm := firevm.Create()
	Traverse(vm, p.Scope)
	checkMnemonics(t, vm.Instructions, []string{
		"PUSH", // push string data
		"PUSH", // push hash(x)
		"ADD"
	})
}

func TestBinaryExpressionBytecodeStringReferenceConcat(t *testing.T) {
	p := ParseString(`
		var (
			a = "hello"
			b = "world"
		)
		a + b
		`)
	vm := firevm.Create()
	Traverse(vm, p.Scope)
	checkMnemonics(t, vm.Instructions, []string{
		"PUSH", // push string data
		"PUSH", // push hash(x)
		"ADD"
	})
}

func TestUnaryExpressionBytecodeLiteral(t *testing.T) {
	p := ParseString("!1")
	vm := firevm.Create()
	Traverse(vm, p.Scope)
	checkMnemonics(t, vm.Instructions, []string{
		"PUSH", // push string data
		"NOT"
	})
}

func TestCallExpressionBytecodeLiteral(t *testing.T){
	p := ParseString(`
		doSomething("data")
		`)
	vm := firevm.Create()
	Traverse(vm, p.Scope)
	checkMnemonics(t, vm.Instructions, []string{
		"PUSH", // push string data
		"NOT"
	})
}

func TestCallExpressionBytecodeUseResult(t *testing.T){
	p := ParseString(`
		s := doSomething("data")
		`)
	vm := firevm.Create()
	Traverse(vm, p.Scope)
	checkMnemonics(t, vm.Instructions, []string{
		"PUSH", // push string data
		"NOT"
	})
}

func TestCallExpressionBytecodeUseMultipleResults(t *testing.T){
	p := ParseString(`
		s, a, p := doSomething("data")
		`)
	vm := firevm.Create()
	Traverse(vm, p.Scope)
	checkMnemonics(t, vm.Instructions, []string{
		"PUSH", // push string data
		// this is where the function code would go
		"PUSH", // push p
		"SET", // set p
		"PUSH", // push a
		"SET", // set a
		"PUSH", // push s
		"SET", // set s
	})
}

func TestCallExpressionBytecodeIgnoredResult(t *testing.T){
	p := ParseString(`
		s, _, p := doSomething("data")
		`)
	vm := firevm.Create()
	Traverse(vm, p.Scope)
	checkMnemonics(t, vm.Instructions, []string{
		"PUSH", // push string data
		// this is where the function code would go
		// stack ends: {R3, R2, R1}
		"PUSH", // push p
		"SET", // set p
		"POP", // ignore second return value
		"PUSH", // push s
		"SET", // set s
	})
}

func TestCallExpressionBytecodeNestedCall(t *testing.T){
	p := ParseString(`
		err := saySomething(doSomething("data"))
		`)
	vm := firevm.Create()
	Traverse(vm, p.Scope)
	checkMnemonics(t, vm.Instructions, []string{
		"PUSH", // push string data
		// this is where the function code would go
		// stack ends: {R1}
		// second function code goes here
		"PUSH", // push err
		"SET", // set err
	})
}
