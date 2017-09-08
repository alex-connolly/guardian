package ast

import (
	firevm "axia/fireVM"
	"axia/guardian/go/compiler/parser"
	"testing"
)

func TestStorageArrayDeclaration(t *testing.T) {
	p := parser.ParseString(
		`contract ArrayTest {
            animals = [string]{
                "Dog", "Cat"
            }
        }
    `)
	vm := firevm.NewVM()
	p.Scope.Traverse(vm)
	checkMnemonics(t, vm.Instructions, []string{
		"",
	})
}

func TestStorageMapDeclaration(t *testing.T) {
	p := parser.ParseString(
		`contract ArrayTest {
            animals = map[string]string{
                "Dog":"canine", "Cat":"feline",
            }
        }
    `)
	vm := firevm.NewVM()
	p.Scope.Traverse(vm)
	checkMnemonics(t, vm.Instructions, []string{
		"",
	})
}

func TestMemoryArrayDeclaration(t *testing.T) {
	p := parser.ParseString(
		`contract ArrayTest {

            func doThings(){
                animals = [string]{
                    "Dog", "Cat"
                }
            }
        }
    `)
	vm := firevm.NewVM()
	p.Scope.Traverse(vm)
	checkMnemonics(t, vm.Instructions, []string{
		"",
	})
}

func TestMemoryMapDeclaration(t *testing.T) {
	p := parser.ParseString(
		`contract ArrayTest {

            func doThings(){
                animals = map[string]string{
                    "Dog":"canine", "Cat":"feline",
                }
            }
        }
    `)
	vm := firevm.NewVM()
	p.Scope.Traverse(vm)
	checkMnemonics(t, vm.Instructions, []string{
		"",
	})
}