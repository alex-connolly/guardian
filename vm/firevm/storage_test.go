package firevm

import (
	"testing"

	"github.com/end-r/guardian"
)

func TestStoreVariable(t *testing.T) {
	a := new(Arsonist)
	guardian.CompileString(a, `
            contract Dog {
                var name = "Buffy"
            }
        `)
	checkMnemonics(t, a.VM.Instructions, []string{
		"PUSH", // push string data
		"PUSH", // push hash(name)
		"SET",  // store result in memory
	})
}

func TestStoreAndLoadVariable(t *testing.T) {
	a := new(Arsonist)
	guardian.CompileString(a, `
            contract Dog {
                var name = "Buffy"

                constructor(){
                    log(name)
                }

            }
        `)
	checkMnemonics(t, a.VM.Instructions, []string{
		"PUSH",  // push string data
		"PUSH",  // push hash(name)
		"STORE", // store result in memory
		"PUSH",  // push hash(name)
		"LOAD",  // load the data at name
		"LOG",   // expose that data
	})
}
