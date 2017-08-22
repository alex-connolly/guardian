package ast

import (
	"testing"

	"github.com/end-r/guardian/go/compiler/parser"

	"github.com/end-r/firevm"
)

func TestBytecodeClassDeclaration(t *testing.T) {
	p := parser.ParseString(
		`contract Tester {
            var x = 5
            const y = 10
        }`)
	vm := firevm.NewVM()
	p.Scope.Traverse(vm)
	checkMnemonics(t, vm.Instructions, []string{
		"PUSH", // push string data
		"PUSH", // push hash(x)
		"PUSH", // push offset (0)
		"SET",  // store result in memory at hash(x)[0]
	})
}

func TestBytecodeFuncDeclaration(t *testing.T) {
	p := parser.ParseString(
		`contract Tester {
            add(a, b int) int {
                return a + b
            }
        }`)
	vm := firevm.NewVM()
	p.Scope.Traverse(vm)
	checkMnemonics(t, vm.Instructions, []string{
		"PUSH",   // push string data
		"PUSH",   // push hash(x)
		"ADD",    // push offset (0)
		"RETURN", // store result in memory at hash(x)[0]
	})
}

func TestBytecodeInterfaceDeclaration(t *testing.T) {

}
