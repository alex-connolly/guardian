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
	case ast.ContractDeclaration:
		a.traverseContract(node.(ast.ContractDeclarationNode))
		break
	case ast.FuncDeclaration:
		a.traverseFunc(node.(ast.FuncDeclarationNode))
		break
	}
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
