package bytecode

import "axia/guardian/go/compiler/ast"

func Traverse(node ast.Node) {
	node.Traverse()
}
