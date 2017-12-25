package evm

import (
	"strconv"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/util"
	"github.com/end-r/guardian/validator"
	"github.com/end-r/vmgen"
)

// GuardianEVM
type GuardianEVM struct {
	VM                 *vmgen.VM
	hooks              []hook
	lastSlot           uint
	lastOffset         uint
	storage            map[string]*storageBlock
	freedMemory        []*memoryBlock
	memoryCursor       uint
	memory             map[string]*memoryBlock
	currentlyAssigning string
	internalHooks      map[string]hook
	externalHooks      map[string]hook
	globalHooks        map[string]hook
	eventHooks         map[string]hook
	lifecycleHooks     map[string]hook
	inStorage          bool
}

func push(data []byte) (code vmgen.Bytecode) {
	if len(data) > 32 {
		// TODO: error
	}
	m := "PUSH" + strconv.Itoa(len(data))
	code.Add(m, data...)
	return code
}

// support all offsets which can be stored in a 64 bit integer
func pushMarker(offset int) (code vmgen.Bytecode) {
	//TODO: fix
	code.AddMarker("PUSH"+strconv.Itoa(8), offset)
	return code
}

var (
	builtinScope *ast.ScopeNode
	litMap       validator.LiteralMap
	opMap        validator.OperatorMap
)

func (evm GuardianEVM) Traverse(node ast.Node) (vmgen.Bytecode, util.Errors) {
	// do pre-processing/hooks etc
	code := evm.traverse(node)
	// generate the bytecode
	// finalise the bytecode
	//evm.finalise()
	return code, nil
}

// NewGuardianEVM ...
func NewVM() GuardianEVM {
	return GuardianEVM{}
}

// A hook conditionally jumps the code to a particular point
//

func (e *GuardianEVM) finalise() {

	// add external functions
	// add internal functions
	// add events

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

//

// can be called from outside or inside the contract
func (e *GuardianEVM) hookPublicFunc(h *hook) {

}

// can be
func (e *GuardianEVM) hookPrivateFunc(h *hook) {

}

func (e GuardianEVM) traverse(n ast.Node) (code vmgen.Bytecode) {
	/* initialise the vm
	if e.VM == nil {
		e.VM = firevm.NewVM()
	}*/
	switch node := n.(type) {
	case *ast.ScopeNode:
		return e.traverseScope(node)
	case *ast.ClassDeclarationNode:
		return e.traverseClass(node)
	case *ast.InterfaceDeclarationNode:
		return e.traverseInterface(node)
	case *ast.EnumDeclarationNode:
		return e.traverseEnum(node)
	case *ast.EventDeclarationNode:
		return e.traverseEvent(node)
	case *ast.ExplicitVarDeclarationNode:
		return e.traverseExplicitVarDecl(node)
	case *ast.TypeDeclarationNode:
		return e.traverseType(node)
	case *ast.ContractDeclarationNode:
		return e.traverseContract(node)
	case *ast.FuncDeclarationNode:
		return e.traverseFunc(node)
	case *ast.ForStatementNode:
		return e.traverseForStatement(node)
	case *ast.AssignmentStatementNode:
		return e.traverseAssignmentStatement(node)
	case *ast.CaseStatementNode:
		return e.traverseCaseStatement(node)
	case *ast.ReturnStatementNode:
		return e.traverseReturnStatement(node)
	case *ast.IfStatementNode:
		return e.traverseIfStatement(node)
	case *ast.SwitchStatementNode:
		return e.traverseSwitchStatement(node)
	}
	return code
}

func (e *GuardianEVM) traverseScope(s *ast.ScopeNode) (code vmgen.Bytecode) {
	if s == nil {
		return code
	}
	if s.Declarations != nil {
		for _, d := range s.Declarations.Array() {
			code.Concat(e.traverse(d.(ast.Node)))
		}
	}

	return code
}
