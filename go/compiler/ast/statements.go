package ast

type AssignmentStatementNode struct {
	Left  []Node
	Right []Node
}

func (n AssignmentStatementNode) Type() NodeType { return AssignmentStatement }

func (n AssignmentStatementNode) Validate(t NodeType) bool {
	return true
}

func (n AssignmentStatementNode) Declare(key string, node Node) {

}

func (n AssignmentStatementNode) Traverse() {

}

type ReturnStatementNode struct {
	Results []Node
}

func (n ReturnStatementNode) Type() NodeType { return ReturnStatement }

func (n ReturnStatementNode) Validate(t NodeType) bool {
	return true
}

func (n ReturnStatementNode) Declare(key string, node Node) {

}

func (n ReturnStatementNode) Traverse() {

}

type BranchStatementNode struct {
	Identifier string
}

func (n BranchStatementNode) Type() NodeType { return BranchStatement }

func (n BranchStatementNode) Validate(t NodeType) bool {
	return true
}

func (n BranchStatementNode) Declare(key string, node Node) {

}

func (n BranchStatementNode) Traverse() {

}

type IfStatementNode struct {
	Init Node
	Cond Node
	Body BlockStatementNode
	Else BlockStatementNode
}

func (n IfStatementNode) Type() NodeType { return IfStatement }

func (n IfStatementNode) Validate(t NodeType) bool {
	return true
}

func (n IfStatementNode) Declare(key string, node Node) {

}

func (n IfStatementNode) Traverse() {

}

type SwitchStatementNode struct {
	Target  Node
	Clauses []CaseStatementNode
	Default BlockStatementNode
}

func (n SwitchStatementNode) Type() NodeType { return SwitchStatement }

func (n SwitchStatementNode) Validate(t NodeType) bool {
	return true
}

func (n SwitchStatementNode) Declare(key string, node Node) {

}

func (n SwitchStatementNode) Traverse() {

}

type CaseStatementNode struct {
	Clauses []Node
	Body    BlockStatementNode
}

func (n CaseStatementNode) Type() NodeType { return CaseStatement }

func (n CaseStatementNode) Validate(t NodeType) bool {
	return true
}

func (n CaseStatementNode) Declare(key string, node Node) {

}

func (n CaseStatementNode) Traverse() {

}

type BlockStatementNode struct {
	Body []Node
}

func (n BlockStatementNode) Type() NodeType { return BlockStatement }

func (n BlockStatementNode) Validate(t NodeType) bool {
	return true
}

func (n BlockStatementNode) Declare(key string, node Node) {

}

func (n BlockStatementNode) Traverse() {

}

type ForStatementNode struct {
	Init  Node
	Cond  Node
	Post  Node
	Block BlockStatementNode
}

func (n ForStatementNode) Type() NodeType { return ForStatement }

func (n ForStatementNode) Validate(t NodeType) bool {
	return true
}

func (n ForStatementNode) Declare(key string, node Node) {

}

func (n ForStatementNode) Traverse() {

}
