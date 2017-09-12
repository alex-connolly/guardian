package vm

import (
	"github.com/end-r/guardian/compiler/lexer"

	"github.com/end-r/firevm"
	"github.com/end-r/guardian/compiler/ast"
	"github.com/end-r/vmgen"
)

// An Arsonist burns down trees
type Arsonist struct {
	VM *vmgen.VM
}

// Traverse ...
func (a *Arsonist) Traverse(node ast.Node) {
	// initialise the vm
	if a.VM == nil {
		a.VM = firevm.NewVM()
	}
	switch node.Type() {
	case ast.ClassDeclaration:
		a.traverseClass(node.(ast.ClassDeclarationNode))
		break
	case ast.InterfaceDeclaration:
		a.traverseInterface(node.(ast.InterfaceDeclarationNode))
		break
	case ast.EnumDeclaration:
		a.traverseEnum(node.(ast.EnumDeclarationNode))
		break
	case ast.EventDeclaration:
		a.traverseEvent(node.(ast.EventDeclarationNode))
		break
	case ast.TypeDeclaration:
		a.traverseType(node.(ast.TypeDeclarationNode))
		break
	case ast.ContractDeclaration:
		a.traverseContract(node.(ast.ContractDeclarationNode))
		break
	case ast.FuncDeclaration:
		a.traverseFunc(node.(ast.FuncDeclarationNode))
		break
	case ast.ConstructorDeclaration:
		a.traverseConstructor(node.(ast.ConstructorDeclarationNode))
		break
	case ast.IndexExpression:
		a.traverseIndex(node.(ast.IndexExpressionNode))
		break
	case ast.CallExpression:
		a.traverseCallExpr(node.(ast.CallExpressionNode))
		break
	case ast.BinaryExpression:
		a.traverseBinaryExpr(node.(ast.BinaryExpressionNode))
		break
	case ast.UnaryExpression:
		a.traverseUnaryExpr(node.(ast.UnaryExpressionNode))
		break
	case ast.Literal:
		a.traverseLiteral(node.(ast.LiteralNode))
		break
	case ast.CompositeLiteral:
		a.traverseCompositeLiteral(node.(ast.CompositeLiteralNode))
		break
	case ast.SliceExpression:
		a.traverseSliceExpression(node.(ast.SliceExpressionNode))
		break
	case ast.ArrayLiteral:
		a.traverseArrayLiteral(node.(ast.ArrayLiteralNode))
		break
	case ast.MapLiteral:
		a.traverseMapLiteral(node.(ast.MapLiteralNode))
		break
	}
}

func (a *Arsonist) traverseArrayLiteral(n ast.ArrayLiteralNode) {

}

func (a *Arsonist) traverseSliceExpression(n ast.SliceExpressionNode) {

}

func (a *Arsonist) traverseCompositeLiteral(n ast.CompositeLiteralNode) {

}

func (a *Arsonist) traverseType(n ast.TypeDeclarationNode) {

}

func (a *Arsonist) traverseConstructor(n ast.ConstructorDeclarationNode) {

}

func (a *Arsonist) traverseClass(n ast.ClassDeclarationNode) {

}

func (a *Arsonist) traverseInterface(n ast.InterfaceDeclarationNode) {

}

func (a *Arsonist) traverseEnum(n ast.EnumDeclarationNode) {

}

func (a *Arsonist) traverseContract(n ast.ContractDeclarationNode) {

}

func (a *Arsonist) traverseEvent(n ast.EventDeclarationNode) {

}

func (a *Arsonist) traverseFunc(n ast.FuncDeclarationNode) {

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

func (a *Arsonist) traverseBinaryExpr(n ast.BinaryExpressionNode) {
	a.Traverse(n.Left)
	a.Traverse(n.Right)
	// operation
	a.VM.AddBytecode(binaryOps[n.Operator])
}

var unaryOps = map[lexer.TokenType]string{
	lexer.TknNot: "NOT",
}

func (a *Arsonist) traverseUnaryExpr(n ast.UnaryExpressionNode) {
	a.VM.AddBytecode(unaryOps[n.Operator])
	a.Traverse(n.Operand)
}

func (a *Arsonist) traverseCallExpr(n ast.CallExpressionNode) {
	for _, arg := range n.Arguments {
		a.Traverse(arg)
	}
	// parameters are at the top of the stack
	// jump to the top of the function
}

func (a *Arsonist) traverseLiteral(n ast.LiteralNode) {
	// Literal Nodes are directly converted to push instructions
	var parameters []byte
	bytes := n.GetBytes()
	parameters = append(parameters, byte(len(bytes)))
	parameters = append(parameters, bytes...)
	a.VM.AddBytecode("PUSH", parameters...)
}

func (a *Arsonist) traverseIndex(n ast.IndexExpressionNode) {
	// evaluate the index
	a.Traverse(n.Index)
	// then MLOAD it at the index offset
	a.Traverse(n.Expression)
	a.VM.AddBytecode("GET")
}

func (a *Arsonist) traverseMapLiteral(n ast.MapLiteralNode) {
	for k, v := range n.Data {
		a.Traverse(k)
		a.Traverse(v)
	}
	// push the size of the map
	a.VM.AddBytecode("PUSH", byte(1), byte(len(n.Data)))
	a.VM.AddBytecode("MAP")
}

func (a *Arsonist) traverseReference(n ast.ReferenceNode) {
	// reference e.g. dog.tail.wag()
	// get the object
	// if in storage
	a.VM.AddBytecode("LOAD")
}
