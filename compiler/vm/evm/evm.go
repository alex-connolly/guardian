package vm

import (
	"github.com/end-r/vmgen"

	"github.com/end-r/guardian/compiler/ast"
)

// EVMTraverser burns down trees
type EVMTraverser struct {
	VM *vmgen.VM
}

// Traverse ...
func (e *EVMTraverser) Traverse(node ast.Node) {

}
