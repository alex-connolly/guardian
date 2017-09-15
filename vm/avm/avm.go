package avm

import (
	"github.com/end-r/guardian/compiler/ast"
)

// Traverser ...
type Traverser struct {
}

// NewTraverser ...
func NewTraverser() Traverser {
	return Traverser{}
}

// Traverse ...
func (a Traverser) Traverse(node ast.Node) {

}
