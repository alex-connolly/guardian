package ast

import (
	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/typing"

	"github.com/blang/semver"
)

type ImportStatementNode struct {
	Path  string
	Alias string
}

func (n *ImportStatementNode) Type() NodeType { return ImportStatement }

type PackageStatementNode struct {
	Name    string
	Version semver.Version
}

func (n *PackageStatementNode) Type() NodeType { return PackageStatement }

type AssignmentStatementNode struct {
	Modifiers typing.Modifiers
	Left      []ExpressionNode
	Operator  token.Type
	Right     []ExpressionNode
}

func (n *AssignmentStatementNode) Type() NodeType { return AssignmentStatement }

type ReturnStatementNode struct {
	Results []ExpressionNode
}

func (n *ReturnStatementNode) Type() NodeType { return ReturnStatement }

type ConditionNode struct {
	Condition ExpressionNode
	Body      *ScopeNode
}

func (n *ConditionNode) Type() NodeType { return IfStatement }

type IfStatementNode struct {
	Init       Node
	Conditions []*ConditionNode
	Else       *ScopeNode
}

func (n *IfStatementNode) Type() NodeType { return IfStatement }

type SwitchStatementNode struct {
	Target      ExpressionNode
	Cases       *ScopeNode
	IsExclusive bool
}

func (n *SwitchStatementNode) Type() NodeType { return SwitchStatement }

type CaseStatementNode struct {
	Expressions []ExpressionNode
	Block       *ScopeNode
}

func (n *CaseStatementNode) Type() NodeType { return CaseStatement }

type ForStatementNode struct {
	Init  *AssignmentStatementNode
	Cond  ExpressionNode
	Post  StatementNode
	Block *ScopeNode
}

func (n *ForStatementNode) Type() NodeType { return ForStatement }

type ForEachStatementNode struct {
	Variables    []string
	Producer     ExpressionNode
	Block        *ScopeNode
	ResolvedType typing.Type
}

func (n *ForEachStatementNode) Type() NodeType { return ForEachStatement }

type FlowStatementNode struct {
	Token token.Type
}

func (n *FlowStatementNode) Type() NodeType { return FlowStatement }
