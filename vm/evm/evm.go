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

func (e Traverser) GetTypes() map[string]guardian.Type{
	return map[string]guardian.Type {
		"uint": guardian.NumericType{size:256, signed: false, integer:true},
		"int": guardian.NumericType{size:256, signed: true, integer:true},
		"uint256": guardian.NumericType{size:256, signed: false, integer:true},
		"int256": guardian.NumericType{size:256, signed: true, integer:true},
		"uint8": guardian.NumericType{size:8, signed: false, integer:true},
		"int8": guardian.NumericType{size:8, signed: true, integer:true},
		"byte": guardian.NumericType{size:8, signed: true, integer:true},
		"string": guardian.ArrayType{key:"byte"},
		"address":
	}
}

func (e Traverser) GetBuiltins() map[string]guardian.Type {
	 return map[string]guardian.Builtin {
		 "msg": guardian.ParseBuiltin(`
			 	data string
				gas uint
				sender address
				func sig() []byte {
					return data[:4]
				}
				`),
			
		 }
	 }
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
