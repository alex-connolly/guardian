package evm

import (
	"testing"

	"github.com/end-r/guardian"
)

func TestBinaryExpressionBytecodeLiterals(t *testing.T) {
	e := new(EVMTraverser)
	guardian.CompileString(e, "1 + 5")
	checkMnemonics(t, e.VM.Instructions, []string{
		"PUSH", // push 1
		"PUSH", // push 5
		"ADD",  // perform operation
	})
	checkStack(t, e.VM.Stack, [][]byte{
		[]byte{byte(6)},
	})
}
