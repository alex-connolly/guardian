package ast

type AssignmentStatementNode struct {
	Left  []Node
	Right []Node
}

func (n *AssignmentStatementNode) Type() NodeType { return AssignmentStatement }

type ReturnStatementNode struct {
	Results []Node
}

func (n *ReturnStatementNode) Type() NodeType { return ReturnStatement }

type BranchStatementNode struct {
	identifier string
}

func (n *BranchStatementNode) Type() NodeType { return BranchStatement }

type IfStatementNode struct {
	Init Node
	Cond Node
	Body BlockStatementNode
	Else BlockStatementNode
}

func (n *IfStatementNode) Type() NodeType { return IfStatement }

type SwitchStatementNode struct {
	Target  Node
	Clauses []CaseClause
	Default BlockStatementNode
}

func (n *SwitchStatementNode) Type() NodeType { return SwitchStatement }

type CaseClauseNode struct {
	Clauses []Node
	Body    BlockStatementNode
}

func (n *CaseClauseNode) Type() NodeType { return CaseClause }

type BlockStatementNode struct {
	Body []Node
}

func (n *BlockStatementNode) Type() NodeType { return BlockStatement }

type ForStatementNode struct {
	Init  Node
	Cond  Node
	Post  Node
	Block BlockStatementNode
}

func (n *ForStatementNode) Type() NodeType { return ForStatement }
