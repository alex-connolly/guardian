package guardian

import (
	"github.com/end-r/guardian/compiler/parser"

	"github.com/end-r/guardian/compiler/ast"
)

// Traverser ...
type Traverser interface {
	Traverse(ast.Node)
}

// CompileFile ...
func CompileFile(t Traverser, path string) []string {
	// generate AST
	p := parser.ParseFile(path)
	// Traverse AST
	t.Traverse(p.Scope)
	return nil
}

// CompileString ...
func CompileString(t Traverser, data string) []string {

	// generate AST
	p := parser.ParseString(data)
	// Traverse AST
	t.Traverse(p.Scope)
	return nil
}

// CompileBytes ...
func CompileBytes(t Traverser, bytes []byte) []string {
	// generate AST
	p := parser.ParseBytes(bytes)
	// Traverse AST
	t.Traverse(p.Scope)
	return nil
}
