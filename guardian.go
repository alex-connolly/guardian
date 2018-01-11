package guardian

import (
	"fmt"

	"github.com/end-r/vmgen"

	"github.com/end-r/guardian/lexer"
	"github.com/end-r/guardian/parser"
	"github.com/end-r/guardian/util"
	"github.com/end-r/guardian/validator"
)

func reportErrors(category string, errs util.Errors) {
	msg := fmt.Sprintf("%s Errors\n", category)
	msg += errs.Format()
	fmt.Println(msg)
}

// CompileBytes ...
func CompileBytes(vm validator.VM, bytes []byte) vmgen.Bytecode {
	tokens, errs := lexer.Lex(bytes)

	if errs != nil {
		reportErrors("Lexing", errs)
	}

	ast, errs := parser.Parse(tokens, errs)

	if errs != nil {
		reportErrors("Parsing", errs)
	}

	errs = validator.Validate(ast, vm)

	if errs != nil {
		reportErrors("Type Validation", errs)
	}

	bytecode, errs := vm.Traverse(ast)

	if errs != nil {
		reportErrors("Bytecode Generation", errs)
	}
	return bytecode
}

// CompileString ...
func CompileString(vm validator.VM, data string) vmgen.Bytecode {
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
