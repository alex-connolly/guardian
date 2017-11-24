package evm

import (
	"axia/guardian/validator"

	"github.com/end-r/guardian/ast"
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

func (e Traverser) GetTypes() map[string]validator.Type {
	return map[string]validator.Type{
		"uint":    validator.NumericType{size: 256, signed: false, integer: true},
		"int":     validator.NumericType{size: 256, signed: true, integer: true},
		"uint256": validator.NumericType{size: 256, signed: false, integer: true},
		"int256":  validator.NumericType{size: 256, signed: true, integer: true},
		"uint8":   validator.NumericType{size: 8, signed: false, integer: true},
		"int8":    validator.NumericType{size: 8, signed: true, integer: true},
		"uint16":  validator.NumericType{size: 16, signed: false, integer: true},
		"int16":   validator.NumericType{size: 16, signed: true, integer: true},
		"byte":    validator.NumericType{size: 8, signed: true, integer: true},
		"string":  validator.ArrayType{key: "byte"},
		"address": validator.ArrayType{key: "byte", length: 20},
		"bool":    validator.BooleanType{},
	}
}

func (e Traverser) GetBuiltins() string {
	return `

		func add(a, b int) int {
			return a + b
		}

		// partial declarations mean you implement them yourself

		balance func(a address) uint256
		transfer func(a address, amount uint256) uint
		send func(a address, amount uint256) bool
		call func(a address) bool
		delegateCall func(a address)

		// variable declarations follow the following format
		// these variables have no set values
		class BuiltinMessage {
			data []byte
			gas uint
			sender address
			sig [4]byte
		}

		class BuiltinBlock {
			timestamp uint
			number uint
			coinbase address
			gaslimit uint
			blockhash func(blockNumber uint) [32]byte
		}

		class BuiltinTransaction {
			gasprice uint
			origin address
		}

		block BuiltinBlock
		msg BuiltinMessage
		tx BuiltinTransaction
	`
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
