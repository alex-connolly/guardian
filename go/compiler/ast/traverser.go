package ast

import (
	"github.com/end-r/vmgen"
)

func Traverse(vm *vmgen.VM, scope Node) {
	scope.Traverse(vm)
}
