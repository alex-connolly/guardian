package ast

import (
	"github.com/end-r/guardian/go/compiler/lexer"

	"github.com/end-r/vmgen"
)

type AssignmentStatementNode struct {
	Left     []ExpressionNode
	Operator lexer.TokenType
	Right    []ExpressionNode
}

func (n AssignmentStatementNode) Type() NodeType { return AssignmentStatement }

func (n AssignmentStatementNode) Traverse(vm *vmgen.VM) {
	if len(n.Left) != len(n.Right) {
		if len(n.Right) == 1 {
			n.Right[0].Traverse(vm)
			vm.AddBytecode("PUSH")
		}
	} else {
		/*for i, l := range n.Left {

			vm.AddBytecode("PUSH")
			vm.AddBytecode("PUSH")
			// then do a memset
			vm.AddBytecode("SET")
		}*/
	}

}

type ReturnStatementNode struct {
	Results []ExpressionNode
}

func (n ReturnStatementNode) Type() NodeType { return ReturnStatement }

func (n ReturnStatementNode) Traverse(vm *vmgen.VM) {

}

type ConditionNode struct {
	Condition ExpressionNode
	Body      ScopeNode
}

func (n ConditionNode) Type() NodeType { return IfStatement }

func (n ConditionNode) Traverse(vm *vmgen.VM) {
	//position := vm.Position()
	//ops := n.Body.Traverse(vm)
	n.Condition.Traverse(vm)
	//vm.AddBytecode("JUMP", position+ops)
}

type IfStatementNode struct {
	Init       Node
	Conditions []ConditionNode
	Else       ScopeNode
}

func (n IfStatementNode) Type() NodeType { return IfStatement }

func (n IfStatementNode) Traverse(vm *vmgen.VM) {
	n.Init.Traverse(vm)
}

type SwitchStatementNode struct {
	Target      ExpressionNode
	Clauses     []CaseStatementNode
	Default     ScopeNode
	IsExclusive bool
}

func (n SwitchStatementNode) Type() NodeType { return SwitchStatement }

func (n SwitchStatementNode) Traverse(vm *vmgen.VM) {
	// traverse Expression
	n.Target.Traverse(vm)
	for _, clause := range n.Clauses {
		clause.Traverse(vm)
		if n.IsExclusive {
			vm.AddBytecode("")
		}
	}

}

type CaseStatementNode struct {
	Expressions []ExpressionNode
	Block       ScopeNode
}

func (n CaseStatementNode) Type() NodeType { return CaseStatement }

func (n CaseStatementNode) Traverse(vm *vmgen.VM) {
	for _, expr := range n.Expressions {
		expr.Traverse(vm)
	}

}

type ForStatementNode struct {
	Init  AssignmentStatementNode
	Cond  ExpressionNode
	Post  StatementNode
	Block ScopeNode
}

func (n ForStatementNode) Type() NodeType { return ForStatement }

func (n ForStatementNode) Traverse(vm *vmgen.VM) {
	n.Init.Traverse(vm)
	vm.AddBytecode("JUMPDEST")
	//position := vm.Position()
	n.Cond.Traverse(vm)
	vm.AddBytecode("JUMPI")
	n.Block.Traverse(vm)
	n.Post.Traverse(vm)
	//vm.AddBytecode("JUMP", position)

}
