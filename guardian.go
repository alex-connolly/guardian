package guardian

import (
	"fmt"

	"github.com/end-r/guardian/util"
)

func reportErrors(category string, errs util.Errors) {
	msg := fmt.Sprintf("%s Errors\n", category)
	msg += errs.Format()
	fmt.Println(msg)
}

/*
// CompileBytes ...
func CompileBytes(vm validator.VM, bytes []byte) vmgen.Bytecode {
	l := lexer.LexString(string(bytes))

	if l.Errors != nil {
		reportErrors("Lexing", l.Errors)
	}

	p := parser.Parse(l)

	if p.Errors != nil {
		reportErrors("Parsing", p.Errors)
	}

	errs := validator.ValidateString()

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

func CompileFilesData(vm validator.VM, data [][]byte) (vmgen.Bytecode, util.Errors) {

}
*/

/* EVM ...
func EVM() Traverser {
	return evm.NewTraverser()
}

// FireVM ...
func FireVM() Traverser {
	return firevm.NewTraverser()
}*/
