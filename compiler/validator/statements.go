package validator

import "github.com/end-r/guardian/compiler/ast"

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
	case ast.CaseStatement:
		v.validateCaseStatement(node.(ast.CaseStatementNode))
		break
	case ast.ReturnStatement:
		v.validateReturnStatement(node.(ast.ReturnStatementNode))
		break
	case ast.SwitchStatement:
		v.validateSwitchStatement(node.(ast.SwitchStatementNode))
		break
	}
}

func (v *Validator) validateAssignment(node ast.AssignmentStatementNode) {
	// valid assignments must have
	// 1. valid left hand expression (cannot be a call, literal, slice)
	// 2. type of left == type of right

	// check step one first
	for _, l := range node.Left {
		switch l.Type() {
		case ast.CallExpression, ast.Literal, ast.MapLiteral,
			ast.ArrayLiteral, ast.SliceExpression:
			v.addError("Cannot assign to expression")
		}
	}

	// special case where right side is of length 1:
	// TODO:

	// else, just do a tuple comparison
	leftTuple := v.ExpressionTuple(node.Left)
	rightTuple := v.ExpressionTuple(node.Right)

	if !leftTuple.compare(rightTuple) {
		v.addError("Cannot assign %s to %s", WriteType(rightTuple), WriteType(leftTuple))
	}

}

func (v *Validator) validateIfStatement(node ast.IfStatementNode) {
	//v.Validate(node.Init)
	for _, cond := range node.Conditions {
		// condition must be of type bool
		v.requireType(standards[Bool], v.resolveExpression(cond.Condition))
		//v.validateScope(cond.Body)
	}
}

func (v *Validator) validateSwitchStatement(node ast.SwitchStatementNode) {

	switchType := v.resolveExpression(node.Target)
	// target must be matched by all cases
	for _, clause := range node.Cases.Sequence {

		/*for _, expr := range clause.Expressions {
			v.requireType(switchType, v.resolveExpression(expr))
		}*/
		//v.validateScope(clause.Block)
	}

}

func (v *Validator) validateCaseStatement(node ast.CaseStatementNode) {

}

func (v *Validator) validateReturnStatement(node ast.ReturnStatementNode) {
	// must be in the context of a function (checked during ast generation)
	// resolved tuple must match function expression
	scope := v.scope
	for scope.parent.scope.Type() != ast.FuncDeclaration {
		scope = scope.parent
	}
	//fd := scope.scope

}

func (v *Validator) validateForStatement(node ast.ForStatementNode) {
	// init statement must be valid
	v.validateStatement(node.Init)

	// cond statement must be a boolean
	v.requireType(standards[Bool], v.resolveExpression(node.Cond))

	// post statement must be valid
	// v.Validate(node.Init)

	// v.ValidateScope(node.Block)
}
