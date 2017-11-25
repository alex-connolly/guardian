package guardian

import (
	"fmt"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/lexer"
	"github.com/end-r/guardian/parser"
	"github.com/end-r/guardian/util"
	"github.com/end-r/guardian/validator"
	"github.com/end-r/vmgen"
)

// A VMTranslator is the mechanism through which all vm-specific features are applied
// to the Guardian AST: bytecode generation, type enforcement
type VM interface {
	Traverse(ast.Node) vmgen.Bytecode
	GetBuiltins() ast.Node
	Types() map[string]validator.Type
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

	errs = validator.Validate(vm, ast)

	if errs != nil {
		reportErrors("Type Valdidation", errs)
	}

	bytecode, errs := vm.Traverse(ast)

	if errs != nil {
		reportErrors("Bytecode Generation", errs)
	}
}

/* EVM ...
func EVM() Traverser {
	return evm.NewTraverser()
}

// FireVM ...
func FireVM() Traverser {
	return firevm.NewTraverser()
}*/
