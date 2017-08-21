package ast

import (
	"testing"
)

func TestBytecodeClassDeclaration(t *testing.T) {
	p := ParseString(
		`contract Tester {
            var x = 5
            const y = 10
        }`)
	vm := firevm.Create()
	Traverse(vm, p.Scope)
	checkMnemonics(t, vm.Instructions, []string{
		"PUSH", // push string data
		"PUSH", // push hash(x)
		"PUSH", // push offset (0)
		"SET",  // store result in memory at hash(x)[0]
	})
}

func TestBytecodeFuncDeclaration(t *testing.T) {
	p := ParseString(
		`contract Tester {
            add(a, b int) int {
                return a + b
            }
        }`)
	vm := firevm.Create()
	Traverse(vm, p.Scope)
	checkMnemonics(t, vm.Instructions, []string{
		"PUSH",   // push string data
		"PUSH",   // push hash(x)
		"ADD",    // push offset (0)
		"RETURN", // store result in memory at hash(x)[0]
	})
}

func TestBytecodeInterfaceDeclaration(t *testing.T) {

}
