package evm

import (
	"github.com/end-r/guardian/validator"

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

func getIntegerTypes() map[string]validator.Type {
	m := map[string]validator.Type{}
	const maxSize = 256
	const increment = 8
	for i := increment; i <= maxSize; i += increment {
		m["uint"+string(i)] = validator.NumericType{size: i, signed: false, integer: true}
		m["int"+string(i)] = validator.NumericType{size: i, signed: true, integer: true}
	}
	m["int"] = validator.NumericType{size: maxSize, signed: false, integer: true}
	m["uint"] = validator.NumericType{size: maxSize, signed: true, integer: true}
	return m
}

func (e Traverser) GetTypes() map[string]validator.Type {
	it := getIntegerTypes()

	s := map[string]validator.Type{
		"byte":    validator.NumericType{size: 8, signed: true, integer: true},
		"string":  validator.ArrayType{key: "byte"},
		"address": validator.ArrayType{key: "byte", length: 20},
		"bool":    validator.BooleanType{},
	}

	for k, v := range it {
		s[k] = v
	}
	return s
}

func (e Traverser) GetBuiltins() string {
	return `

		wei = 1
		kwei = 1000 * wei
		babbage = kwei
		mwei = 1000 * kwei
		lovelace = mwei
		gwei = 1000 * mwei
		shannon = gwei
		microether = 1000 * gwei
		szabo = microether
		milliether = 1000 * microether
		finney = milliether
		ether = 1000 * milliether

		balance func(a address) uint256
		transfer func(a address, amount uint256) uint
		send func(a address, amount uint256) bool
		call func(a address) bool
		delegateCall func(a address)

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
