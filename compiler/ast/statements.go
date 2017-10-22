package ast

import (
	"github.com/end-r/guardian/compiler/lexer"
)

type AssignmentStatementNode struct {
	Modifiers []lexer.TokenType
	Left      []ExpressionNode
	Operator  lexer.TokenType
	Right     []ExpressionNode
}

func (n AssignmentStatementNode) Type() NodeType { return AssignmentStatement }

type ReturnStatementNode struct {
	Results []ExpressionNode
}

func (n ReturnStatementNode) Type() NodeType { return ReturnStatement }

type ConditionNode struct {
	Condition ExpressionNode
	Body      *ScopeNode
}

func (n ConditionNode) Type() NodeType { return IfStatement }

type IfStatementNode struct {
	Init       Node
	Conditions []ConditionNode
	Else       *ScopeNode
}

func (n IfStatementNode) Type() NodeType { return IfStatement }

type SwitchStatementNode struct {
	Target      ExpressionNode
	Cases       *ScopeNode
	IsExclusive bool
}

func (n SwitchStatementNode) Type() NodeType { return SwitchStatement }

type CaseStatementNode struct {
	Expressions []ExpressionNode
	Block       *ScopeNode
}

func (n CaseStatementNode) Type() NodeType { return CaseStatement }

type ForStatementNode struct {
	Init  *AssignmentStatementNode
	Cond  ExpressionNode
	Post  StatementNode
	Block *ScopeNode
}

func (n ForStatementNode) Type() NodeType { return ForStatement }
