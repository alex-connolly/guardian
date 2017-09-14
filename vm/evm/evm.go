package evm

import (
	"github.com/end-r/firevm"

	"github.com/end-r/vmgen"

	"github.com/end-r/guardian/compiler/ast"
)

// EVMTraverser burns down trees
type EVMTraverser struct {
	VM *vmgen.VM
}

func NewTraverser() EVMTraverser {
	return EVMTraverser{}
}

// Traverse ...
func (e *EVMTraverser) Traverse(node ast.Node) {
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
	case ast.ConstructorDeclaration:
		e.traverseConstructor(node.(ast.ConstructorDeclarationNode))
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
