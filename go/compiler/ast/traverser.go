package ast

import (
	"axia/guardian/go/compiler/lexer"

	"github.com/end-r/firevm"
)

func Traverse(node Node) {
	node.Traverse()
}

const operators = map[lexer.TknType]string{
	lexer.TknAdd: "ADD",
	lexer.TknSub: "SUB",
	lexer.TknMul: "MUL",
	lexer.TknDiv: "DIV",
	lexer.TknMod: "MOD",
	lexer.TknShr: "SHR",
	lexer.TknShl: "SHL",
}

// principles for mapping assignment operators
// simply expand and do as normal
// e.g. x += 5 --> x = x + 5

// TraverseOperator traverses an operator
func TraverseOperator(vm *firevm.FireVM, tkn lexer.TknType) {
	vm.AddInstruction(operators[tkn])
}

func traverseExpression(node Node) {

}
