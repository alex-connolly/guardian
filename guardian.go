package guardian

import (
	"github.com/end-r/guardian/compiler/ast"
	"github.com/end-r/guardian/compiler/parser"
	"github.com/end-r/guardian/compiler/validator"
	"github.com/end-r/guardian/vm/evm"
	"github.com/end-r/guardian/vm/firevm"
)

// Traverser ...
type Traverser interface {
	Traverse(ast.Node)
	AddBytecode(string, ...byte)
}

// CompileFile ...
func CompileFile(t Traverser, path string) []string {
	// generate AST
	p := parser.ParseFile(path)

	v := validator.ValidateScope(p.Scope)
	// Traverse AST
	t.Traverse(p.Scope)
	return nil
}

// CompileString ...
func CompileString(t Traverser, data string) []string {

	// generate AST
	p := parser.ParseString(data)

	v := validator.ValidateScope(p.Scope)
	// Traverse AST
	t.Traverse(p.Scope)
	return nil
}

// CompileBytes ...
func CompileBytes(t Traverser, bytes []byte) []string {
	// generate AST
	p := parser.ParseBytes(bytes)

	v := validator.ValidateScope(p.Scope)

	// Traverse AST
	t.Traverse(p.Scope)
	return nil
}

// EVM ...
func EVM() Traverser {
	return evm.NewTraverser()
}

// FireVM ...
func FireVM() Traverser {
	return firevm.NewTraverser()
}
