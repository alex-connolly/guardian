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

func TestExclusiveSwitchStatement(t *testing.T) {
	p := parser.ParseString(`
		x := 1 + 5
		exclusive switch x {
		case 4, 5:
			x *= 2
		case 6:
			x += 5
		}
		`)
	vm := firevm.NewVM()
	p.Scope.Traverse(vm)
	checkMnemonics(t, vm.Instructions, []string{
		"PUSH", // push 6 --> constant evaluation should be already done
		"PUSH", // push 4
		"EQL",  // check for equality
		"JMPI", // conditional jump
		"PUSH", // push 5
		"EQL",  // check for equality
		"PUSH", // push 6
		"EQL",  // check for equality
	})
}

func TestSwitchStatement(t *testing.T) {
	p := parser.ParseString(`
		x := 1 + 5
		switch x {
		case 4, 5:
			break
		case 6:
			break
		}
		`)
	vm := firevm.NewVM()
	p.Scope.Traverse(vm)
	checkMnemonics(t, vm.Instructions, []string{
		"PUSH", // push 6 --> constant evaluation should be already done
		"PUSH", // push 4
		"EQL",  // check for equality
		"JMPI", // conditional jump
		"PUSH", // push 5
		"EQL",  // check for equality
		"PUSH", // push 6
		"EQL",  // check for equality
	})
}
