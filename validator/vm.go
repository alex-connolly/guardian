package validator

import (
	"github.com/end-r/guardian/ast"
	"github.com/end-r/vmgen"
)

// A VM is the mechanism through which all vm-specific features are applied
// to the Guardian AST: bytecode generation, type enforcement etc
type VM interface {
	Traverse(ast.Node) vmgen.Bytecode
	Builtins() *ast.ScopeNode
	Primitives() map[string]Type
}

type TestVM struct {
}

func NewTestVM() TestVM {
	return TestVM{}
}

func (v TestVM) Traverse(ast.Node) vmgen.Bytecode {
	return vmgen.Bytecode{}
}

func (v TestVM) Builtins() *ast.ScopeNode {
	return nil
}

func (v TestVM) Primitives() map[string]Type {
	return nil
}
