package evm

import (
	"testing"

	"github.com/end-r/goutil"
	"github.com/end-r/guardian"
)

func TestIncrement(t *testing.T) {

}

func TestForStatement(t *testing.T) {
	bytecode, errs := guardian.CompileString(NewVM(), `
        for i := 0; i < 5; i++ {

        }
    `)
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
	goutil.Assert(t, bytecode.CompareMnemonics(expected), "wrong mnemonics")
}

func TestReturnStatement(t *testing.T) {

}
