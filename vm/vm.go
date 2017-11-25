package vm

import (
	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/validator"
	"github.com/end-r/vmgen"
)

// A VMImplementation is the mechanism through which all vm-specific features are applied
// to the Guardian AST: bytecode generation, type enforcement etc
type VMImplementation interface {
	Traverse(ast.Node) vmgen.Bytecode
	Builtins() ast.Node
	Types() map[string]validator.Type
}
