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
