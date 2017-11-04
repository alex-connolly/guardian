package validator

import (
	"github.com/end-r/guardian/compiler/ast"
)

func (v *Validator) validateStatement(node ast.Node) {
	switch node.Type() {
	case ast.AssignmentStatement:
		v.validateAssignment(node.(ast.AssignmentStatementNode))
		break
	case ast.ForStatement:
		v.validateForStatement(node.(ast.ForStatementNode))
		break
	case ast.IfStatement:
		v.validateIfStatement(node.(ast.IfStatementNode))
		break
	case ast.ReturnStatement:
		v.validateReturnStatement(node.(ast.ReturnStatementNode))
		break
	case ast.SwitchStatement:
		v.validateSwitchStatement(node.(ast.SwitchStatementNode))
		break
		// TODO: handle call statement
	}
}

func (v *Validator) validateAssignment(node ast.AssignmentStatementNode) {
	// TODO: fundamental: define operator or not
	// valid assignments must have
	// 1. valid left hand expression (cannot be a call, literal, slice)
	// 2. type of left == type of right

	// check step one first
	for _, l := range node.Left {
		switch l.Type() {
		case ast.CallExpression, ast.Literal, ast.MapLiteral,
			ast.ArrayLiteral, ast.SliceExpression, ast.FuncLiteral:
			v.addError("Cannot assign to expression")
		}
	}

	if len(node.Left) > len(node.Right) && len(node.Right) == 1 {
		rightType := v.resolveType(node.Right[0])
		for _, l := range node.Left {
			left := v.resolveType(l)
			if !assignableTo(rightType, left) {
				v.addError(errInvalidAssignment, WriteType(rightType), WriteType(left))
			}
		}
		return
	}

	if len(node.Right) != len(node.Left) {
		v.addError("Assignment count mismatch: %s, %s")
		return
	}

	leftTuple := v.ExpressionTuple(node.Left)
	rightTuple := v.ExpressionTuple(node.Right)

	if !leftTuple.compare(rightTuple) {
		v.addError(errInvalidAssignment, WriteType(leftTuple), WriteType(rightTuple))
	}

	for i, left := range node.Left {
		if leftTuple.types[i] == standards[Unknown] {
			if id, ok := left.(ast.IdentifierNode); ok {
				v.DeclareType(id.Name, rightTuple.types[i])
			}
		}
	}

}

func (v *Validator) validateIfStatement(node ast.IfStatementNode) {

	if node.Init != nil {
		v.validateStatement(node.Init)
	}

	for _, cond := range node.Conditions {
		// condition must be of type bool
		v.requireType(standards[Bool], v.resolveExpression(cond.Condition))
		v.validateScope(cond.Body)
	}
}

func (v *Validator) validateSwitchStatement(node ast.SwitchStatementNode) {

	switchType := v.resolveExpression(node.Target)
	// target must be matched by all cases
	for _, node := range node.Cases.Sequence {
		if node.Type() == ast.CaseStatement {
			v.validateCaseStatement(switchType, node.(ast.CaseStatementNode))
		}
	}

}

func (v *Validator) validateCaseStatement(switchType Type, clause ast.CaseStatementNode) {
	for _, expr := range clause.Expressions {
		v.requireType(switchType, v.resolveExpression(expr))
	}
	v.validateScope(clause.Block)
}

func (v *Validator) validateReturnStatement(node ast.ReturnStatementNode) {
	// must be in the context of a function (checked during ast generation)
	// resolved tuple must match function expression
	scope := v.scope
	for scope.parent.scope.Type() != ast.FuncDeclaration && scope != nil {
		scope = scope.parent
	}

	// scope is now a func declaration

}

func (v *Validator) validateForStatement(node ast.ForStatementNode) {
	// init statement must be valid
	// if it exists
	if node.Init != nil {
		v.validateAssignment(*node.Init)
	}

	// cond statement must be a boolean
	v.requireType(standards[Bool], v.resolveExpression(node.Cond))

	// post statement must be valid
	if node.Post != nil {
		v.validateStatement(node.Post)
	}

	v.validateScope(node.Block)
}
