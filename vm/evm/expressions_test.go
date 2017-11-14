package evm

import (
	"axia/guardian/compiler/parser"
	"testing"

	"github.com/end-r/goutil"
	"github.com/end-r/guardian"
)

func TestBinaryExpressionBytecodeLiterals(t *testing.T) {
	e := new(Traverser)
	guardian.CompileString(e, "x = 1 + 5")
	checkMnemonics(t, e.VM.Instructions, []string{
		"PUSH", // push 1
		"PUSH", // push 5
		"ADD",  // perform operation
	})
	checkStack(t, e.VM.Stack, [][]byte{
		[]byte{byte(6)},
	})
}

func TraverseSimpleIdentifierExpression(t *testing.T) {
	e := new(Traverser)
	expr := parser.ParseExpression("hello")
	bytes := e.traverse(expr)
	expected := bytecodeString("s")
	goutil.Assert(t, bytes.compare(expected), invalidBytecodeMessage(bytes, expected))
}
