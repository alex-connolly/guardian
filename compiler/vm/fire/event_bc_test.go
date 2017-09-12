package vm

import (
	"axia/guardian"
	"testing"
)

func TestParametrizedEvent(t *testing.T) {
	a := new(Arsonist)
	guardian.CompileString(a, `
            contract Dog {

                event NameEvent(string)

                var name = "Buffy"

                constructor(){
                    NameEvent(name)
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
