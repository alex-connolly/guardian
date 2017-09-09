package bytecode

import "github.com/end-r/guardian/go/ast"

// Traverser of a Guardian AST
type Traverser interface {
	Traverse(ast.Node)
}
