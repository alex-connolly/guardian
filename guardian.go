package guardian

import (
	"github.com/end-r/guardian/compiler/parser"

	"github.com/end-r/guardian/compiler/ast"
)

// Compiler ...
type Compiler struct {
	traverser Traverser
}

// Traverser ...
type Traverser interface {
	Traverse(ast.Node)
}

// New ...
func New(t Traverser) *Compiler {
	return &Compiler{
		traverser: t,
	}
}

// CompileFile ...
func (c *Compiler) CompileFile(path string) []string {
	// generate AST
	p := parser.ParseFile(path)
	// Traverse AST
	c.traverser.Traverse(p.Scope)
	return nil
}

// CompileString ...
func (c *Compiler) CompileString(data string) []string {
	// generate AST
	p := parser.ParseString(data)
	// Traverse AST
	c.traverser.Traverse(p.Scope)
	return nil
}

// CompileBytes ...
func (c *Compiler) CompileBytes(bytes []byte) []string {
	// generate AST
	p := parser.ParseBytes(bytes)
	// Traverse AST
	c.traverser.Traverse(p.Scope)
	return nil
}
