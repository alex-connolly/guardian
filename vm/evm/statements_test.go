package evm

import (
	"testing"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/parser"

	"github.com/end-r/goutil"
)

func TestIncrement(t *testing.T) {

}

func TestAssignmentStatement(t *testing.T) {
	scope, _ := parser.ParseString(`
        i = 0
    `)
	f := scope.Sequence[0].(*ast.AssignmentStatementNode)
	e := NewVM()
	bytecode := e.traverseAssignmentStatement(f)
	expected := []string{
		// push left
		"PUSH",
		// push right
		"PUSH",
		// store (default is memory)
		"MSTORE",
	}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestForStatement(t *testing.T) {
	scope, _ := parser.ParseString(`
        for i = 0; i < 5; i++ {

        }
    `)
	f := scope.Sequence[0].(*ast.ForStatementNode)
	e := NewVM()
	bytecode := e.traverseForStatement(f)
	expected := []string{
		// init
		"PUSH", "PUSH", "MSTORE",
		// top of loop
		"JUMPDEST",
		// condition
		"PUSH", "MLOAD", "PUSH", "LT", "PUSH", "JUMPI",
		// body
		// post
		"PUSH", "MLOAD", "PUSH", "ADD", "PUSH", "MSTORE",
		// jump back to top
		"JUMP",
		"JUMPDEST",
	}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestReturnStatement(t *testing.T) {

}
