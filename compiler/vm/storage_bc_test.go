package vm

import (
	"testing"

	"github.com/end-r/firevm"
	"github.com/end-r/guardian/compiler/parser"
)

func TestStoreVariable(t *testing.T) {
	p := parser.ParseString(`
            contract Dog {
                var name = "Buffy"
            }
        `)
	vm := firevm.NewVM()
	p.Scope.Traverse(vm)
	checkMnemonics(t, vm.Instructions, []string{
		"PUSH", // push string data
		"PUSH", // push hash(name)
		"SET",  // store result in memory
	})
}

func TestStoreAndLoadVariable(t *testing.T) {
	p := parser.ParseString(`
            contract Dog {
                var name = "Buffy"

                constructor(){
                    log(name)
                }

            }
        `)
	vm := firevm.NewVM()
	p.Scope.Traverse(vm)
	checkMnemonics(t, vm.Instructions, []string{
		"PUSH",  // push string data
		"PUSH",  // push hash(name)
		"STORE", // store result in memory
		"PUSH",  // push hash(name)
		"LOAD",  // load the data at name
		"LOG",   // expose that data
	})
}
