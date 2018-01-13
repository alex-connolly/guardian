package ast

import (
	"github.com/end-r/guardian/token"
	"github.com/end-r/guardian/typing"
	"github.com/end-r/guardian/util"

	"github.com/blang/semver"
)

type ImportStatementNode struct {
	Begin, Final util.Location
	Path         string
	Alias        string
}

func (n *ImportStatementNode) Start() util.Location { return n.Begin }
func (n *ImportStatementNode) End() util.Location   { return n.Final }
func (n *ImportStatementNode) Type() NodeType       { return ImportStatement }

type PackageStatementNode struct {
	Begin, Final util.Location
	Name         string
	Version      semver.Version
}

func (n *PackageStatementNode) Start() util.Location { return n.Begin }
func (n *PackageStatementNode) End() util.Location   { return n.Final }
func (n *PackageStatementNode) Type() NodeType       { return PackageStatement }

type AssignmentStatementNode struct {
	Begin, Final util.Location
	Modifiers    typing.Modifiers
	Left         []ExpressionNode
	Operator     token.Type
	Right        []ExpressionNode
}

func (n *AssignmentStatementNode) Start() util.Location { return n.Begin }
func (n *AssignmentStatementNode) End() util.Location   { return n.Final }
func (n *AssignmentStatementNode) Type() NodeType       { return AssignmentStatement }

type ReturnStatementNode struct {
	Begin, Final util.Location
	Results      []ExpressionNode
}

func (n *ReturnStatementNode) Start() util.Location { return n.Begin }
func (n *ReturnStatementNode) End() util.Location   { return n.Final }
func (n *ReturnStatementNode) Type() NodeType       { return ReturnStatement }

type ConditionNode struct {
	Begin, Final util.Location
	Condition    ExpressionNode
	Body         *ScopeNode
}

func (n *ConditionNode) Start() util.Location { return n.Begin }
func (n *ConditionNode) End() util.Location   { return n.Final }
func (n *ConditionNode) Type() NodeType       { return IfStatement }

type IfStatementNode struct {
	Begin, Final util.Location
	Init         Node
	Conditions   []*ConditionNode
	Else         *ScopeNode
}

func (n *IfStatementNode) Start() util.Location { return n.Begin }
func (n *IfStatementNode) End() util.Location   { return n.Final }
func (n *IfStatementNode) Type() NodeType       { return IfStatement }

type SwitchStatementNode struct {
	Begin, Final util.Location
	Target       ExpressionNode
	Cases        *ScopeNode
	IsExclusive  bool
}

func (n *SwitchStatementNode) Start() util.Location { return n.Begin }
func (n *SwitchStatementNode) End() util.Location   { return n.Final }
func (n *SwitchStatementNode) Type() NodeType       { return SwitchStatement }

type CaseStatementNode struct {
	Begin, Final util.Location
	Expressions  []ExpressionNode
	Block        *ScopeNode
}

func (n *CaseStatementNode) Type() NodeType       { return CaseStatement }
func (n *CaseStatementNode) Start() util.Location { return n.Begin }
func (n *CaseStatementNode) End() util.Location   { return n.Final }

type ForStatementNode struct {
	Begin, Final util.Location
	Init         *AssignmentStatementNode
	Cond         ExpressionNode
	Post         StatementNode
	Block        *ScopeNode
}

func (n *ForStatementNode) Type() NodeType       { return ForStatement }
func (n *ForStatementNode) Start() util.Location { return n.Begin }
func (n *ForStatementNode) End() util.Location   { return n.Final }

type ForEachStatementNode struct {
	Begin, Final util.Location
	Variables    []string
	Producer     ExpressionNode
	Block        *ScopeNode
	ResolvedType typing.Type
}

func (n *ForEachStatementNode) Start() util.Location { return n.Begin }
func (n *ForEachStatementNode) End() util.Location   { return n.Final }
func (n *ForEachStatementNode) Type() NodeType       { return ForEachStatement }

type FlowStatementNode struct {
	Begin, Final util.Location
	Token        token.Type
}

func (n *FlowStatementNode) Start() util.Location { return n.Begin }
func (n *FlowStatementNode) End() util.Location   { return n.Final }
func (n *FlowStatementNode) Type() NodeType       { return FlowStatement }
