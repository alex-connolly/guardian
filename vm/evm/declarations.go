package evm

import "github.com/end-r/guardian/compiler/ast"

func (e *Traverser) traverseType(n ast.TypeDeclarationNode) {

}

func (e *Traverser) traverseClass(n ast.ClassDeclarationNode) {
	// create constructor hooks
	// create function hooks
}

func (e *Traverser) traverseInterface(n ast.InterfaceDeclarationNode) {
	// don't need to be interacted with
	// all interfaces are dealt with by the type system
}

func (e *Traverser) traverseEnum(n ast.EnumDeclarationNode) {
	// create hook
}

func (e *Traverser) traverseContract(n ast.ContractDeclarationNode) {
	// create hooks for functions
	// create hooks for constructors
	// create hooks for events
}

func (e *Traverser) addHook(name string) {
	hook := hook{
		name: name,
	}
	e.hooks = append(e.hooks)
}

func (e *Traverser) traverseEvent(n ast.EventDeclarationNode) {
	e.AddBytecode(EncodeSignature())
}

func EncodeSignature() string {
	return ""
}

func (e *Traverser) traverseFunc(n ast.FuncDeclarationNode) {
	e.AddBytecode("JUMPDEST")
	e.Traverse(n.Body)
	// TODO: add something to prevent further execution

	// all evm functions create a hook at the start of the contract
	// when executing, will jump to one of these functions
	e.AddBytecode("CALLDATA")
	// the ABI defines this as being the first 4 bytes of the SHA-3 hash of the function signature
	// as guardian signatures are not stringified quite as easily
	// have to do something clever
	if hasModifier(n, "external") {
		e.AddBytecode(EncodeSignature())
		e.AddBytecode("EQL")
		e.AddBytecode("JMPI")
	}
}
