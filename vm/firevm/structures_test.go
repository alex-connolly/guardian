package firevm

import (
	"testing"

	"github.com/end-r/guardian"
)

func TestStorageArrayDeclaration(t *testing.T) {
	a := new(Arsonist)
	guardian.CompileString(a,
		`contract ArrayTest {
            animals = [string]{
                "Dog", "Cat"
            }
        }
    `)
	checkMnemonics(t, a.VM.Instructions, []string{
		"",
	})
}

func TestStorageMapDeclaration(t *testing.T) {
	a := new(Arsonist)
	guardian.CompileString(a,
		`contract ArrayTest {
            animals = map[string]string{
                "Dog":"canine", "Cat":"feline",
            }
        }
    `)
	checkMnemonics(t, a.VM.Instructions, []string{
		"",
	})
}

func TestMemoryArrayDeclaration(t *testing.T) {
	a := new(Arsonist)
	guardian.CompileString(a,
		`contract ArrayTest {

            func doThings(){
                animals = [string]{
                    "Dog", "Cat"
                }
            }
        }
    `)
	checkMnemonics(t, a.VM.Instructions, []string{
		"",
	})
}

func TestMemoryMapDeclaration(t *testing.T) {
	a := new(Arsonist)
	guardian.CompileString(a,
		`contract ArrayTest {

            func doThings(){
                animals = map[string]string{
                    "Dog":"canine", "Cat":"feline",
                }
            }
        }
    `)
	checkMnemonics(t, a.VM.Instructions, []string{
		"",
	})
}
