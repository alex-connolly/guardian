package evm

import (
	"github.com/end-r/guardian/compiler/ast"
	"github.com/end-r/vmgen"
)

// Traverser ...
type Traverser struct {
	VM        *vmgen.VM
	hooks     []hook
	callables []callable
}

type hook struct {
	name     string
	position int
	bytecode []byte
}

type callable struct {
	name     string
	position int
	bytecode []byte
}

// NewTraverser ...
func NewTraverser() Traverser {
	return Traverser{}
}

// A hook conditionally jumps the code to a particular point
//

// Traverse ...
func (e Traverser) Traverse(node ast.Node) {
	// do pre-processing/hooks etc
	e.traverse(node)
	// generate the bytecode
	// finalise the bytecode
	e.finalise()
}

func (e Traverser) finalise() {
	// number of instructions =
	/*
		for _, hook := range e.hooks {
			e.VM.AddBytecode("POP")

			e.VM.AddBytecode("EQL")
			e.VM.AddBytecode("JMPI")
		}
		// if the data matches none of the function hooks
		e.VM.AddBytecode("STOP")
		for _, callable := range e.callables {
			// add function bytecode
		}*/
}

func (e Traverser) traverse(n ast.Node) (code vmgen.Bytecode) {
	/* initialise the vm
	if e.VM == nil {
		e.VM = firevm.NewVM()
	}*/
	switch node := n.(type) {
	case ast.ClassDeclarationNode:
		return e.traverseClass(node)
	case ast.InterfaceDeclarationNode:
		return e.traverseInterface(node)
	case ast.EnumDeclarationNode:
		return e.traverseEnum(node)
	case ast.EventDeclarationNode:
		return e.traverseEvent(node)
	case ast.TypeDeclarationNode:
		return e.traverseType(node)
	case ast.ContractDeclarationNode:
		return e.traverseContract(node)
	case ast.FuncDeclarationNode:
		return e.traverseFunc(node)
	case ast.ForStatementNode:
		return e.traverseForStatement(node)
	case ast.AssignmentStatementNode:
		return e.traverseAssignmentStatement(node)
	case ast.CaseStatementNode:
		return e.traverseCaseStatement(node)
	case ast.ReturnStatementNode:
		return e.traverseReturnStatement(node)
	case ast.IfStatementNode:
		return e.traverseIfStatement(node)
	case ast.SwitchStatementNode:
		return e.traverseSwitchStatement(node)
	}
	return code
}

func (e *Traverser) inStorage() bool {
	return false
}
