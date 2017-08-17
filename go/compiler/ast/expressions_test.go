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
	goutil.AssertNow(t, len(vm.Instructions) == 3, "wrong num of instructions")
	expected := []string{"PUSH", "PUSH", "ADD"}
	for index, i := range vm.Instructions {
		goutil.Assert(t, i.Mnemonic == expected[index], "wrong mnemonic")
	}
}
