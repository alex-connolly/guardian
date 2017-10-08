package validator

import "github.com/end-r/guardian/compiler/ast"

func validateAssignment(v *Validator, node *ast.AssignmentStatementNode) {
	// valid assignments must have
	// 1. valid left hand expression (cannot be a call, literal, slice)
	// 2. type of left == type of right

	// check step one first
	for _, l := range node.Left {
		switch l.Type() {
		case ast.CallExpression, ast.Literal, ast.MapLiteral,
			ast.ArrayLiteral, ast.SliceExpression:
			// TODO: add error here
			// break?
		}
	}

	// special case where right side is of length 1:

	// short circuit if different lengths
	if len(node.Left) != len(node.Right) {
		// TODO: add error here
	}

	// else, just do a tuple comparison
	leftTuple := v.ExpressionTuple(node.Left)
	rightTuple := v.ExpressionTuple(node.Right)

	if !leftTuple.compare(rightTuple) {
		// TODO: add error here
	}

}

func (v *Validator) validateIfStatement(node *ast.IfStatementNode) {
	//v.Validate(node.Init)
	for _, cond := range node.Conditions {
		// condition must be of type bool
		v.requireType(standards[Bool], v.resolveExpression(cond.Condition))
		//v.validateScope(cond.Body)
	}
}

func (v *Validator) validateSwitchStatement(node *ast.SwitchStatementNode) {

	switchType := v.resolveExpression(node.Target)
	// target must be matched by all cases
	for _, clause := range node.Clauses {
		for _, expr := range clause.Expressions {
			v.requireType(switchType, v.resolveExpression(expr))
		}
		//v.validateScope(clause.Block)
	}

}

func (v *Validator) validateReturnStatement(node *ast.ReturnStatementNode) {

}

func (v *Validator) validateForStatement(node *ast.ForStatementNode) {

}
