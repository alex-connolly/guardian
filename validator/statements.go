package validator

import (
	"github.com/end-r/guardian/typing"

	"github.com/end-r/guardian/ast"
)

func (v *Validator) validateStatement(node ast.Node) {
	switch node.Type() {
	case ast.AssignmentStatement:
		v.validateAssignment(node.(*ast.AssignmentStatementNode))
		break
	case ast.ForStatement:
		v.validateForStatement(node.(*ast.ForStatementNode))
		break
	case ast.IfStatement:
		v.validateIfStatement(node.(*ast.IfStatementNode))
		break
	case ast.ReturnStatement:
		v.validateReturnStatement(node.(*ast.ReturnStatementNode))
		break
	case ast.SwitchStatement:
		v.validateSwitchStatement(node.(*ast.SwitchStatementNode))
		break
		// TODO: handle call statement
	}
}

func (v *Validator) validateAssignment(node *ast.AssignmentStatementNode) {
	// TODO: fundamental: define operator or not
	// valid assignments must have
	// 1. valid left hand expression (cannot be a call, literal, slice)
	// 2. types of left assignable to right
	// 3. right all valid expressions

	for _, r := range node.Right {
		v.validateExpression(r)
	}

	for _, l := range node.Left {
		switch l.Type() {
		case ast.CallExpression, ast.Literal, ast.MapLiteral,
			ast.ArrayLiteral, ast.SliceExpression, ast.FuncLiteral:
			v.addError(errInvalidExpressionLeft)
		}
	}

	leftTuple := v.ExpressionTuple(node.Left)
	rightTuple := v.ExpressionTuple(node.Right)

	if len(leftTuple.Types) > len(rightTuple.Types) && len(rightTuple.Types) == 1 {
		right := rightTuple.Types[0]
		for _, left := range leftTuple.Types {
			if !typing.AssignableTo(right, left) {
				v.addError(errInvalidAssignment, typing.WriteType(left), typing.WriteType(right))
			}
		}

		for i, left := range node.Left {
			if leftTuple.Types[i] == typing.Unknown() {
				if id, ok := left.(*ast.IdentifierNode); ok {
					id.Resolved = rightTuple.Types[0]
					v.declareContextualVar(id.Name, rightTuple.Types[0])
				}
			}
		}

	} else {
		if !rightTuple.Compare(leftTuple) {
			v.addError(errInvalidAssignment, typing.WriteType(leftTuple), typing.WriteType(rightTuple))
		}

		// length of left tuple should always equal length of left
		// this is because tuples are not first class types
		// cannot assign to tuple expressions
		if len(node.Left) == len(leftTuple.Types) {
			for i, left := range node.Left {
				if leftTuple.Types[i] == typing.Unknown() {
					if id, ok := left.(*ast.IdentifierNode); ok {
						id.Resolved = rightTuple.Types[i]
						//fmt.Printf("Declaring %s as %s\n", id.Name,typing.WriteType(rightTuple.Types[i]))
						v.declareContextualVar(id.Name, rightTuple.Types[i])
					}
				}
			}
		}
	}
}

func (v *Validator) validateIfStatement(node *ast.IfStatementNode) {

	if node.Init != nil {
		v.validateStatement(node.Init)
	}

	for _, cond := range node.Conditions {
		// condition must be of type bool
		v.requireType(typing.Boolean(), v.resolveExpression(cond.Condition))
		v.validateScope(cond.Body)
	}

	if node.Else != nil {
		v.validateScope(node.Else)
	}
}

func (v *Validator) validateSwitchStatement(node *ast.SwitchStatementNode) {

	switchType := v.resolveExpression(node.Target)
	// target must be matched by all cases
	for _, node := range node.Cases.Sequence {
		if node.Type() == ast.CaseStatement {
			v.validateCaseStatement(switchType, node.(*ast.CaseStatementNode))
		}
	}

}

func (v *Validator) validateCaseStatement(switchType typing.Type, clause *ast.CaseStatementNode) {
	for _, expr := range clause.Expressions {
		v.requireType(switchType, v.resolveExpression(expr))
	}
	v.validateScope(clause.Block)
}

func (v *Validator) validateReturnStatement(node *ast.ReturnStatementNode) {
	// must be in the context of a function (checked during ast generation)
	// resolved tuple must match function expression

	// scope is now a func declaration
}

func (v *Validator) validateForEachStatement(node *ast.ForEachStatementNode) {
	// get type of
	gen := v.resolveExpression(node.Producer)
	var req int
	switch a := gen.(type) {
	case typing.Map:
		// maps must handle k, v in MAP
		req = 2
		if len(node.Variables) != req {
			v.addError(errInvalidForEachVariables, len(node.Variables), req)
		} else {
			v.declareContextualVar(node.Variables[0], a.Key)
			v.declareContextualVar(node.Variables[1], a.Value)
		}
		break
	case typing.Array:
		// arrays must handle i, v in ARRAY
		req = 2
		if len(node.Variables) != req {
			v.addError(errInvalidForEachVariables, len(node.Variables), req)
		} else {
			v.declareContextualVar(node.Variables[0], v.LargestNumericType(false))
			v.declareContextualVar(node.Variables[1], a.Value)
		}
		break
	default:
		v.addError(errInvalidForEachType, typing.WriteType(gen))
		// prevent double errors
		req = len(node.Variables)
	}

}

func (v *Validator) validateForStatement(node *ast.ForStatementNode) {
	// init statement must be valid
	// if it exists
	if node.Init != nil {
		v.validateAssignment(node.Init)
	}

	// cond statement must be a boolean
	v.requireType(typing.Boolean(), v.resolveExpression(node.Cond))

	// post statement must be valid
	if node.Post != nil {
		v.validateStatement(node.Post)
	}

	v.validateScope(node.Block)
}
