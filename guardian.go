package guardian

import (
	"fmt"

	"github.com/end-r/guardian/validator"
	"github.com/end-r/guardian/ast"
	"github.com/end-r/vmgen"
)

// A VMTranslator is the mechanism through which all vm-specific features are applied
// to the Guardian AST: bytecode generation, type enforcement
type VM interface {
	Traverse(ast.Node) vmgen.Bytecode
	Builtins() map[string]Builtin
	Types() map[string]Type
	Permitted() map[ast.NodeType]bool
}

func reportErrors(category string, errs []util.Error) {
	msg := fmt.Sprintf("%s Errors\n", category)
	for _, e := range errs {
		msg += fmt.Sprintf("%s", e.msg)
	}
	fmt.Println(msg)
}

func CompileBytes(vm VM, bytes []byte) {
	tokens, errs := lexer.Lex(bytes)

	if errs != nil {
		reportErrors("Lexing", errs)
	}

	ast, errs := parser.Parse(tokens)

	if errs != nil {
		reportErrors("Parsing", errs)
	}

	errs := validator.Validate(vm, ast)

	if errs != nil {
		reportErrors("Type Valdidation", errs)
	}

	bytecode, errs := vm.Traverse(ast)

	if errs != nil {
		reportErrors("Bytecode Generation", errs)
	}
}

func CompileFile(vm VM path string){

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
