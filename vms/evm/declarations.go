package evm

import (
	"github.com/end-r/vmgen"

	"github.com/end-r/guardian/lexer"

	"github.com/end-r/guardian/ast"
)

func (e *GuardianEVM) traverseType(n ast.TypeDeclarationNode) (code vmgen.Bytecode) {
	return code
}

func (e *GuardianEVM) traverseClass(n ast.ClassDeclarationNode) (code vmgen.Bytecode) {
	// create constructor hooks
	// create function hooks
	return code
}

func (e *GuardianEVM) traverseInterface(n ast.InterfaceDeclarationNode) (code vmgen.Bytecode) {
	// don't need to be interacted with
	// all interfaces are dealt with by the type system
	return code
}

func (e *GuardianEVM) traverseEnum(n ast.EnumDeclarationNode) (code vmgen.Bytecode) {
	// create hook
	return code
}

func (e *GuardianEVM) traverseContract(n ast.ContractDeclarationNode) (code vmgen.Bytecode) {
	// create hooks for functions
	// create hooks for constructors
	// create hooks for events
	return code
}

func (e *GuardianEVM) addHook(name string) {
	h := hook{
		name: name,
	}
	if e.hooks == nil {
		e.hooks = make([]hook, 0)
	}
	e.hooks = append(e.hooks, h)
}

func (e *GuardianEVM) traverseEvent(n ast.EventDeclarationNode) (code vmgen.Bytecode) {
	return code
}

func (e *GuardianEVM) traverseFunc(n ast.FuncDeclarationNode) (code vmgen.Bytecode) {

	hook := EncodeName(n.Identifier)

	e.addHook(string(hook))

	code.Add("JUMPDEST")
	e.Traverse(n.Body)
	// TODO: add something to prevent further execution

	// all evm functions create a hook at the start of the contract
	// when executing, will jump to one of these functions
	code.Add("CALLDATA")
	// the ABI defines this as being the first 4 bytes of the SHA-3 hash of the function signature
	// as guardian signatures are not stringified quite as easily
	// have to do something clever
	if hasModifier(n, lexer.TknExternal) {
		//code.Add(EncodeSignature())
		code.Add("EQL")
		code.Add("JMPI")
	}
	return code
}