package ast

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestingAssignmentStatementDeclaration(t *testing.T) {
	p := ParseString("x := a + b")
	vm := firevm.Create()
	Traverse(vm, p.Scope)
    expected := []string{
        "PUSH",
        "PUSH",
        "ADD",
        "PUSH",
        "PUSH",
        "MSTORE"
    }
	goutil.AssertNow(t, len(vm.Instructions) == len(expected), "wrong num of instructions")

	for index, i := range vm.Instructions {
		goutil.Assert(t, i.Mnemonic == expected[index], "wrong mnemonic")
	}
}
