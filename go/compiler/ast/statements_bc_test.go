package ast

import (
	"testing"
)

func TestIncrementStatement(t *testing.T) {
	p := ParseString(`x++`)
	Traverse(vm, p.scope)
	checkMnemonics(t, vm.Instructions, []string{
		"PUSH", // push string data
		"PUSH", // push hash(x)
		"PUSH", // push offset (0)
		"SET",  // store result in memory at hash(x)[0]
	})
}

func TestAssignmentStatementLiteralDeclaration(t *testing.T) {
	p := ParseString(`x := "this is a string"`)
	vm := firevm.Create()
	Traverse(vm, p.Scope)
	checkMnemonics(t, vm.Instructions, []string{
		"PUSH", // push string data
		"PUSH", // push hash(x)
		"PUSH", // push offset (0)
		"SET",  // store result in memory at hash(x)[0]
	})
}

func TestAssignmentStatementBinaryExpressionDeclaration(t *testing.T) {
	p := ParseString("x := 1 + 2")
	vm := firevm.Create()
	Traverse(vm, p.Scope)
	checkMnemonics(t, vm.Instructions, []string{
		"PUSH", // push string data
		"PUSH", // push hash(x)
		"PUSH", // push offset (0)
		"SET",  // store result in memory at hash(x)[0]
	})
}

func TestAssignmentStatementReferencingDeclaration(t *testing.T) {
	p := ParseString(`
		x := 5
		y := x
		`)
	vm := firevm.Create()
	Traverse(vm, p.Scope)
	checkMnemonics(t, vm.Instructions, []string{
		"PUSH", // push data
		"PUSH", // push x
		"SET",  // set x
		"PUSH", // push x
		"GET",  // get x
		"PUSH", // get y
		"SET",  // set y
	})
}
