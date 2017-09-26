package typing

import "github.com/end-r/guardian/compiler/ast"

func ValidateScope(scope ast.ScopeNode) {
	// Valid scopes have the following properties:
	// all type assignments are valid
	// all references are defined before they are called

	// Process:
	// declare ever

}

// validates whether a call expression matches the function signature in scope
func ValidateCallExpression(scope ast.ScopeNode, ce ast.CallExpressionNode) {

	fnType := ResolveExpression(ce.Call)

	fn := fnType.(Func)

	if !fn.params.equals(exprTuple(ce.Arguments)) {
		// TODO: report error
	}
}

// validates whether an index expression has:
// maps: comparable index, map indexpression
// arrays: int-resolvable index, array expression
func ValidateIndexExpression(scope ast.ScopeNode, index ast.IndexExpressionNode) {

}

// validates that a slice expression has:
// an int-resolvable low expression (or nothing)
// an int-resolvable hgih expression (or nothing)
// an array-resolvable base expression
func ValidateSliceExpression(scope ast.ScopeNode, index ast.SliceExpressionNode) {

}

// AssertableTo reports whether a value of type V can be asserted to have type T.
func AssertableTo(v *Interface, t Type) bool {
	m, _ := assertableTo(v, t)
	return m == nil
}

// AssignableTo reports whether a value of type V is assignable to a variable of type T.
func AssignableTo(v, t Type) bool {
	x := operand{mode: value, typ: V}
	return x.assignableTo(nil, T, nil) // config not needed for non-constant x
}

// ConvertibleTo reports whether a value of type V is convertible to a value of type T.
func ConvertibleTo(V, T Type) bool {
	x := operand{mode: value, typ: V}
	return x.convertibleTo(nil, T) // config not needed for non-constant x
}

// Implements reports whether type V implements interface T.
func Implements(V Type, T *Interface) bool {
	// TODO: check whether all methods are implemented
	return false
}

func ExpectType(e ast.ExpressionNode, t Type) bool {
	return evaluateType(e) == t
}
