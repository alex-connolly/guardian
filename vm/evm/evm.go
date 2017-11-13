package evm

import (
	"github.com/end-r/firevm"

	"github.com/end-r/vmgen"

	"github.com/end-r/guardian/compiler/ast"
)

type bytecode []byte

// Traverser
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

func (t Traverser) AddBytecode(op string, params ...byte) {
	t.VM.AddBytecode(op, params...)
}

func NewTraverser() Traverser {
	return Traverser{}
}

// A hook conditionally jumps the code to a particular point
//

// Traverse ...
func (e Traverser) Traverse(node ast.Node) bytecode {
	// do pre-processing/hooks etc
	e.traverse(node)
	// generate the bytecode
	// finalise the bytecode
	e.finalise()
}

func (e Traverser) finalise() {
	// number of instructions =
	//
	for _, hook := range e.hooks {
		e.VM.AddBytecode("POP")

		e.VM.AddBytecode("EQL")
		e.VM.AddBytecode("JMPI")
	}
	// if the data matches none of the function hooks
	e.VM.AddBytecode("STOP")
	for _, callable := range e.callables {
		// add function bytecode
	}
}

func (e Traverser) traverse(node ast.Node) bytecode {
	// initialise the vm
	if e.VM == nil {
		e.VM = firevm.NewVM()
	}
	switch node.Type() {
	case ast.ClassDeclaration:
		e.traverseClass(node.(ast.ClassDeclarationNode))
		break
	case ast.InterfaceDeclaration:
		e.traverseInterface(node.(ast.InterfaceDeclarationNode))
		break
	case ast.EnumDeclaration:
		e.traverseEnum(node.(ast.EnumDeclarationNode))
		break
	case ast.EventDeclaration:
		e.traverseEvent(node.(ast.EventDeclarationNode))
		break
	case ast.TypeDeclaration:
		e.traverseType(node.(ast.TypeDeclarationNode))
		break
	case ast.ContractDeclaration:
		e.traverseContract(node.(ast.ContractDeclarationNode))
		break
	case ast.FuncDeclaration:
		e.traverseFunc(node.(ast.FuncDeclarationNode))
		break
	case ast.IndexExpression:
		e.traverseIndex(node.(ast.IndexExpressionNode))
		break
	case ast.CallExpression:
		e.traverseCallExpr(node.(ast.CallExpressionNode))
		break
	case ast.BinaryExpression:
		e.traverseBinaryExpr(node.(ast.BinaryExpressionNode))
		break
	case ast.UnaryExpression:
		e.traverseUnaryExpr(node.(ast.UnaryExpressionNode))
		break
	case ast.Literal:
		e.traverseLiteral(node.(ast.LiteralNode))
		break
	case ast.CompositeLiteral:
		e.traverseCompositeLiteral(node.(ast.CompositeLiteralNode))
		break
	case ast.SliceExpression:
		e.traverseSliceExpression(node.(ast.SliceExpressionNode))
		break
	case ast.ArrayLiteral:
		e.traverseArrayLiteral(node.(ast.ArrayLiteralNode))
		break
	case ast.MapLiteral:
		e.traverseMapLiteral(node.(ast.MapLiteralNode))
		break
	case ast.ForStatement:
		e.traverseForStatement(node.(ast.ForStatementNode))
		break
	case ast.AssignmentStatement:
		e.traverseAssignmentStatement(node.(ast.AssignmentStatementNode))
		break
	case ast.CaseStatement:
		e.traverseCaseStatement(node.(ast.CaseStatementNode))
		break
	case ast.ReturnStatement:
		e.traverseReturnStatement(node.(ast.ReturnStatementNode))
		break
	case ast.IfStatement:
		e.traverseIfStatement(node.(ast.IfStatementNode))
		break
	case ast.SwitchStatement:
		e.traverseSwitchStatement(node.(ast.SwitchStatementNode))
		break
	}
}

func (e *Traverser) inStorage() bool {
	return false
}
