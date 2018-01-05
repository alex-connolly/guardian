package evm

import (
	"testing"

	"github.com/end-r/guardian/validator"

	"github.com/end-r/guardian/ast"

	"github.com/end-r/goutil"
)

func TestIncrement(t *testing.T) {

}

func TestSimpleAssignmentStatement(t *testing.T) {
	e := NewVM()
	scope, _ := validator.ValidateString(e, `
        i = 0
    `)
	f := scope.Sequence[0].(*ast.AssignmentStatementNode)
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

func TestIndexAssignmentStatement(t *testing.T) {
	e := NewVM()
	scope, _ := validator.ValidateString(e, `
        nums [5]int
        nums[3] = 0
    `)

	bytecode := e.traverse(scope)
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

func TestIfStatement(t *testing.T) {
	e := NewVM()
	scope, _ := validator.ValidateString(e, `
        if x = 0; x > 5 {

        }
    `)
	f := scope.Sequence[0].(*ast.IfStatementNode)
	bytecode := e.traverseIfStatement(f)
	expected := []string{
		// init
		"PUSH", "PUSH", "MSTORE",
		// top of loop
		"PUSH", "PUSH", "GT",
		// jumper
		"PUSH", "JUMPI",
		// loop body
		"JUMPDEST",
	}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

/*
func TestElseIfStatement(t *testing.T) {
	e := NewVM()
	scope, _ := validator.ValidateString(e, `
        if x = 0; x > 5 {
            x = 1
        } else if x < 3 {
            x = 2
        }
    `)
	f := scope.Sequence[0].(*ast.IfStatementNode)
	bytecode := e.traverseIfStatement(f)
	expected := []string{
		// init
		"PUSH", "PUSH", "MSTORE",
		// top of loop
		"PUSH", "PUSH", "GT",
		// jumper
		"PUSH", "JUMPI",
		// if body
		"PUSH", "PUSH", "MSTORE",
		// else if condition
		"PUSH", "PUSH", "LT",
		// jumper
		"PUSH", "JUMPI",
		// else if body
		"PUSH", "PUSH", "MSTORE",
		"JUMPDEST",
	}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}
*/
func TestElseStatement(t *testing.T) {
	e := NewVM()

	validator.ValidateString(e, `
        if x = 0; x > 5 {

        } else {

        }
    `)
	/*f := scope.Sequence[0].(*ast.IfStatementNode)
	bytecode := e.traverseIfStatement(f)
	expected := []string{
		// init
		"PUSH", "PUSH", "MSTORE",
		// top of loop
		"PUSH", "PUSH", "GT",
		// jumper
		"PUSH", "JUMPI",
		// if body
		// else
		"JUMPDEST",
	}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())*/
}

/*
func TestForStatement(t *testing.T) {
	e := NewVM()
	scope, _ := validator.ValidateString(e, `
        for i = 0; i < 5; i++ {

        }
    `)
	f := scope.Sequence[0].(*ast.ForStatementNode)
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
*/
func TestReturnStatement(t *testing.T) {

}

func TestBreakStatement(t *testing.T) {
	e := NewVM()
	scope, _ := validator.ValidateString(e, `
        for x = 0; x < 5; x++ {
			if x == 3 {
				break
			}
		}
    `)
	f := scope.Sequence[0].(*ast.ForStatementNode)
	bytecode := e.traverseForStatement(f)
	expected := []string{
		// init
		"PUSH", "PUSH", "MSTORE",
		// top of loop
		"PUSH", "PUSH", "GT",
		// jumper
		"PUSH", "JUMPI",
		// loop body
		// if statement
		"PUSH", "MLOAD", "PUSH", "EQ", "ISZERO", "JUMPI",
		"JUMP",
		"JUMPDEST",
	}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func Test2DLoopBreakStatement(t *testing.T) {
	e := NewVM()
	scope, _ := validator.ValidateString(e, `
        for x = 0; x < 5; x++ {
			for y = 0; y < 5; y++ {
				if x + y == 3 {
					break
				}
			}
		}
    `)
	f := scope.Sequence[0].(*ast.ForStatementNode)
	bytecode := e.traverseForStatement(f)
	expected := []string{
		// init
		"PUSH", "PUSH", "MSTORE",
		// top of loop
		"PUSH", "PUSH", "GT",
		// jumper
		"PUSH", "JUMPI",
		// loop body
		// init
		"PUSH", "PUSH", "MSTORE",
		// top of loop
		"PUSH", "PUSH", "GT",
		// jumper
		"PUSH", "JUMPI",
		// if statement
		"PUSH", "MLOAD", "PUSH", "MLOAD,", "ADD", "PUSH", "EQ", "ISZERO", "JUMPI",
		// break statement
		"JUMP",
		"JUMPDEST",
		"JUMPDEST",
	}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}
