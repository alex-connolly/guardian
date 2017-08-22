package ast

import (
	"testing"

	"github.com/end-r/firevm"
	"github.com/end-r/guardian/go/compiler/parser"
)

func TestIncrementStatement(t *testing.T) {
	p := parser.ParseString(`x++`)
	vm := firevm.NewVM()
	p.Scope.Traverse(vm)
	checkMnemonics(t, vm.Instructions, []string{
		"PUSH", // push string data
		"PUSH", // push hash(x)
		"PUSH", // push offset (0)
		"SET",  // store result in memory at hash(x)[0]
	})
}

func TestAssignmentStatementLiteralDeclaration(t *testing.T) {
	p := parser.ParseString(`x := "this is a string"`)
	vm := firevm.NewVM()
	p.Scope.Traverse(vm)
	checkMnemonics(t, vm.Instructions, []string{
		"PUSH", // push string data
		"PUSH", // push hash(x)
		"PUSH", // push offset (0)
		"SET",  // store result in memory at hash(x)[0]
	})
}

func TestAssignmentStatementBinaryExpressionDeclaration(t *testing.T) {
	p := parser.ParseString("x := 1 + 2")
	vm := firevm.NewVM()
	p.Scope.Traverse(vm)
	checkMnemonics(t, vm.Instructions, []string{
		"PUSH", // push string data
		"PUSH", // push hash(x)
		"PUSH", // push offset (0)
		"SET",  // store result in memory at hash(x)[0]
	})
}

func TestAssignmentStatementReferencingDeclaration(t *testing.T) {
	p := parser.ParseString(`
		x := 5
		y := x
		`)
	vm := firevm.NewVM()
	p.Scope.Traverse(vm)
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

func TestSwitchStatement(t *testing.T) {
	p := parser.ParseString(`
		x := 1 + 5
		switch x {
		case 4:
			break
		case 6:
			break
		}
		`)
	vm := firevm.NewVM()
	p.Scope.Traverse(vm)
	checkMnemonics(t, vm.Instructions, []string{
		"PUSH", // push data
		"PUSH", // push x
		"SET",  // set x
		"PUSH", // push x
		"GET",  // get x
		"PUSH", // get y
		"SET",  /// set y
	})
}
