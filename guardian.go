package guardian

import (
	"github.com/end-r/guardian/compiler/ast"
	"github.com/end-r/vmgen"
)

// A Traverser is the mechanism through which all vm-specific features are applied
// to the Guardian AST: bytecode generation, type enforcement
type VM interface {
	Traverse(ast.Node) vmgen.Bytecode
	Builtins() map[string]Builtin
	Types() map[string]Type
	Permitted() map[ast.NodeType]bool
}

/*
// Traverser ...
type Traverser interface {
	Traverse(ast.Node)
	AddBytecode(string, ...byte)
}

// CompileFile ...
func CompileFile(t Traverser, path string) []string {
	// generate AST
	p, errs := parser.ParseFile(path)

	if errs != nil {
		// print errors
	}

	v := validator.ValidateScope(p.Scope)

	if errs != nil {
		// print errors
	}

	// Traverse AST
	bytecode, errs := t.Traverse(p.Scope)

	/*
		b, errs := bastion.RunTests()


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
}*/

/* EVM ...
func EVM() Traverser {
	return evm.NewTraverser()
}

// FireVM ...
func FireVM() Traverser {
	return firevm.NewTraverser()
}*/
