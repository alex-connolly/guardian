package ast

import (
	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/typing"

	"github.com/blang/semver"
)

type ImportStatementNode struct {
	Begin, Final uint
	Path         string
	Alias        string
}

func (n *ImportStatementNode) Type() NodeType { return ImportStatement }

type PackageStatementNode struct {
	Begin, Final uint
	Name         string
	Version      semver.Version
}

func (n *PackageStatementNode) Type() NodeType { return PackageStatement }

type AssignmentStatementNode struct {
	Begin, Final uint
	Modifiers    typing.Modifiers
	Left         []ExpressionNode
	Operator     token.Type
	Right        []ExpressionNode
}

func (n *AssignmentStatementNode) Type() NodeType { return AssignmentStatement }

type ReturnStatementNode struct {
	Begin, Final uint
	Results      []ExpressionNode
}

func (n *ReturnStatementNode) Type() NodeType { return ReturnStatement }

type ConditionNode struct {
	Begin, Final uint
	Condition    ExpressionNode
	Body         *ScopeNode
}

func (n *ConditionNode) Type() NodeType { return IfStatement }

type IfStatementNode struct {
	Begin, Final uint
	Init         Node
	Conditions   []*ConditionNode
	Else         *ScopeNode
}

func (n *IfStatementNode) Type() NodeType { return IfStatement }

type SwitchStatementNode struct {
	Begin, Final uint
	Target       ExpressionNode
	Cases        *ScopeNode
	IsExclusive  bool
}

func (n *SwitchStatementNode) Start() uint    { return n.Begin }
func (n *SwitchStatementNode) End() uint      { return n.Final }
func (n *SwitchStatementNode) Type() NodeType { return SwitchStatement }

type CaseStatementNode struct {
	Begin, Final uint
	Expressions  []ExpressionNode
	Block        *ScopeNode
}

func (n *CaseStatementNode) Type() NodeType { return CaseStatement }
func (n *CaseStatementNode) Start() uint    { return n.Begin }
func (n *CaseStatementNode) End() uint      { return n.Final }

type ForStatementNode struct {
	Begin, Final uint
	Init         *AssignmentStatementNode
	Cond         ExpressionNode
	Post         StatementNode
	Block        *ScopeNode
}

func (n *ForStatementNode) Type() NodeType { return ForStatement }
func (n *ForStatementNode) Start() uint    { return n.Begin }
func (n *ForStatementNode) End() uint      { return n.Final }

type ForEachStatementNode struct {
	Begin, Final uint
	Variables    []string
	Producer     ExpressionNode
	Block        *ScopeNode
	ResolvedType typing.Type
}

func (n *ForEachStatementNode) Start() uint    { return n.Begin }
func (n *ForEachStatementNode) End() uint      { return n.Final }
func (n *ForEachStatementNode) Type() NodeType { return ForEachStatement }

type FlowStatementNode struct {
	Begin, Final uint
	Token        token.Type
}

func (n *FlowStatementNode) Start() uint    { return n.Begin }
func (n *FlowStatementNode) End() uint      { return n.Final }
func (n *FlowStatementNode) Type() NodeType { return FlowStatement }
