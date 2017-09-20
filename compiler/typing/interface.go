package typing

import "github.com/end-r/guardian/compiler/ast"

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
