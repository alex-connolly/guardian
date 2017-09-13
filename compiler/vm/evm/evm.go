package evm

import (
	"github.com/end-r/firevm"
	"github.com/end-r/guardian/compiler/lexer"

	"github.com/end-r/vmgen"

	"github.com/end-r/guardian/compiler/ast"
)

// EVMTraverser burns down trees
type EVMTraverser struct {
	VM *vmgen.VM
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

func (e *EVMTraverser) traverseSwitchStatement(n ast.SwitchStatementNode) {

}

func (e *EVMTraverser) traverseCaseStatement(n ast.CaseStatementNode) {

}

func (e *EVMTraverser) traverseForStatement(n ast.ForStatementNode) {

}

func (e *EVMTraverser) traverseReturnStatement(n ast.ReturnStatementNode) {

}

func (e *EVMTraverser) traverseIfStatement(n ast.IfStatementNode) {

}

func (e *EVMTraverser) traverseAssignmentStatement(n ast.AssignmentStatementNode) {

	if n.IsStorage {
		for i, r := range n.Right {
			e.Traverse(r)
			if len(n.Left) == 1 {
				e.Traverse(n.Left[0])
			} else {
				e.Traverse(n.Left[i])
			}
			e.VM.AddBytecode("STORE")
		}
	}
}

func (e *EVMTraverser) traverseArrayLiteral(n ast.ArrayLiteralNode) {

}

func (e *EVMTraverser) traverseSliceExpression(n ast.SliceExpressionNode) {

}

func (e *EVMTraverser) traverseCompositeLiteral(n ast.CompositeLiteralNode) {

}

func (e *EVMTraverser) traverseType(n ast.TypeDeclarationNode) {

}

func (e *EVMTraverser) traverseConstructor(n ast.ConstructorDeclarationNode) {

}

func (e *EVMTraverser) traverseClass(n ast.ClassDeclarationNode) {

}

func (e *EVMTraverser) traverseInterface(n ast.InterfaceDeclarationNode) {

}

func (e *EVMTraverser) traverseEnum(n ast.EnumDeclarationNode) {

}

func (e *EVMTraverser) traverseContract(n ast.ContractDeclarationNode) {

}

func (e *EVMTraverser) traverseEvent(n ast.EventDeclarationNode) {

}

func (e *EVMTraverser) traverseFunc(n ast.FuncDeclarationNode) {

}

var binaryOps = map[lexer.TokenType]string{
	lexer.TknAdd: "ADD",
	lexer.TknSub: "SUB",
	lexer.TknMul: "MUL",
	lexer.TknDiv: "DIV",
	lexer.TknMod: "MOD",
	lexer.TknShl: "SHL",
	lexer.TknShr: "SHR",
	lexer.TknAnd: "AND",
	lexer.TknOr:  "OR",
	lexer.TknXor: "XOR",
}

func (e *EVMTraverser) traverseBinaryExpr(n ast.BinaryExpressionNode) {
	e.Traverse(n.Left)
	e.Traverse(n.Right)
	// operation
	e.VM.AddBytecode(binaryOps[n.Operator])
}

var unaryOps = map[lexer.TokenType]string{
	lexer.TknNot: "NOT",
}

func (e *EVMTraverser) traverseUnaryExpr(n ast.UnaryExpressionNode) {
	e.VM.AddBytecode(unaryOps[n.Operator])
	e.Traverse(n.Operand)
}

func (e *EVMTraverser) traverseCallExpr(n ast.CallExpressionNode) {
	for _, arg := range n.Arguments {
		e.Traverse(arg)
	}
	// parameters are at the top of the stack
	// jump to the top of the function
}

func (e *EVMTraverser) traverseLiteral(n ast.LiteralNode) {
	// Literal Nodes are directly converted to push instructions
	var parameters []byte
	bytes := n.GetBytes()
	parameters = append(parameters, byte(len(bytes)))
	parameters = append(parameters, bytes...)
	e.VM.AddBytecode("PUSH", parameters...)
}

func (e *EVMTraverser) traverseIndex(n ast.IndexExpressionNode) {
	// evaluate the index
	e.Traverse(n.Index)
	// then MLOAD it at the index offset
	e.Traverse(n.Expression)
	e.VM.AddBytecode("GET")
}

func (e *EVMTraverser) traverseMapLiteral(n ast.MapLiteralNode) {
	for k, v := range n.Data {
		e.Traverse(k)
		e.Traverse(v)
	}
	// push the size of the map
	e.VM.AddBytecode("PUSH", byte(1), byte(len(n.Data)))
	e.VM.AddBytecode("MAP")
}

func (e *EVMTraverser) traverseReference(n ast.ReferenceNode) {

	// reference e.g. dog.tail.wag()
	// get the object
	/*if n.InStorage {
		// if in storage
		// only the top level name is accessible in storage
		// everything else is accessed
		e.VM.AddBytecode("PUSH", len(n.Names[0]), n.Names[0])
		e.VM.AddBytecode("LOAD")

		// now get the sub-references
		// e.VM.AddBytecode("", params)
	} else {
		e.VM.AddBytecode("GET")
	}*/
}
