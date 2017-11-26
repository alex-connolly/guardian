package guardian

import (
	"axia/vmgen"
	"fmt"

	"github.com/end-r/guardian/lexer"
	"github.com/end-r/guardian/parser"
	"github.com/end-r/guardian/util"
	"github.com/end-r/guardian/validator"
	"github.com/end-r/guardian/vm"
)

func reportErrors(category string, errs []util.Errors) {
	msg := fmt.Sprintf("%s Errors\n", category)
	msg += errs.format()
	fmt.Println(msg)
}

func CompileBytes(vm vm.VM, bytes []byte) {
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

func CompileBytes(vm vm.VM, bytes []byte) vmgen.Bytecode {
	tokens, errs := lexer.Lex(bytes)

	if errs != nil {
		reportErrors("Lexing", errs)
		return nil
	}

	ast, errs := parser.Parse(tokens)

	if errs != nil {
		reportErrors("Parsing", errs)
		return nil
	}

	errs = validator.Validate(vm, ast)

	if errs != nil {
		reportErrors("Type Valdidation", errs)
		return nil
	}

	bytecode, errs := vm.Traverse(ast)

	if errs != nil {
		reportErrors("Bytecode Generation", errs)
		return nil
	}

	return bytecode
}

func CompileString(vm vm.VM, data string) vmgen.Bytecode {
	return CompileBytes(vm, []byte(data))
}

/* EVM ...
func EVM() Traverser {
	return evm.NewTraverser()
}

// FireVM ...
func FireVM() Traverser {
	return firevm.NewTraverser()
}*/
