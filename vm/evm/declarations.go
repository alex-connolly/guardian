package evm

import "github.com/end-r/guardian/compiler/ast"

func (e *Traverser) traverseType(n ast.TypeDeclarationNode) {

}

func (e *Traverser) traverseClass(n ast.ClassDeclarationNode) {

}

func (e *Traverser) traverseInterface(n ast.InterfaceDeclarationNode) {

}

func (e *Traverser) traverseEnum(n ast.EnumDeclarationNode) {

}

func (e *Traverser) traverseContract(n ast.ContractDeclarationNode) {

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
