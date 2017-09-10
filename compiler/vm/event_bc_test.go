package vm

import (
	"testing"

	"github.com/end-r/firevm"
	"github.com/end-r/guardian/compiler/parser"
)

func TestParametrizedEvent(t *testing.T) {
	p := parser.ParseString(`
            contract Dog {

                event NameEvent(string)

                var name = "Buffy"

                constructor(){
                    NameEvent(name)
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
