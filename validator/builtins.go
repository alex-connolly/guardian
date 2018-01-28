package validator

import (
	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/typing"
)

func (v *Validator) validateBuiltinDeclarations(scope *ast.ScopeNode) {
	if scope.Declarations != nil {
		// order doesn't matter here
		for _, i := range scope.Declarations.Map() {
			// add in placeholders for all declarations
			v.validateDeclaration(i.(ast.Node))
		}
	}
}

func (v *Validator) validateBuiltinSequence(scope *ast.ScopeNode) {
	for _, node := range scope.Sequence {
		v.validate(node)
	}
}

func SimpleLiteral(typeName string) LiteralFunc {
	return func(v *Validator, data string) typing.Type {
		t, _ := v.isTypeVisible(typeName)
		return t
	}
}

func BooleanLiteral(v *Validator, data string) typing.Type {
	return typing.Boolean()
}

func (m OperatorMap) Add(function OperatorFunc, types ...token.Type) {
	for _, t := range types {

		m[t] = function
	}
}

func SimpleOperator(typeName string) OperatorFunc {
	return func(v *Validator, types []typing.Type, exprs []ast.ExpressionNode) typing.Type {
		t, _ := v.isTypeVisible(typeName)
		return t
	}
}
func BinaryNumericOperator(v *Validator, types []typing.Type, exprs []ast.ExpressionNode) typing.Type {
	left := typing.ResolveUnderlying(types[0])
	right := typing.ResolveUnderlying(types[1])
	if na, ok := left.(*typing.NumericType); ok {
		if nb, ok := right.(*typing.NumericType); ok {
			if na.BitSize > nb.BitSize {
				return v.SmallestNumericType(na.BitSize, true)
			}
			return v.SmallestNumericType(nb.BitSize, true)
		}
	}
	return typing.Invalid()
}

func BinaryIntegerOperator(v *Validator, types []typing.Type, exprs []ast.ExpressionNode) typing.Type {
	if na, ok := types[0].(*typing.NumericType); ok && na.Integer {
		if nb, ok := types[1].(*typing.NumericType); ok && nb.Integer {
			if na.BitSize > nb.BitSize {
				return na
			}
			return nb
		}
	}
	return typing.Invalid()
}

func CastOperator(v *Validator, types []typing.Type, exprs []ast.ExpressionNode) typing.Type {

	left := types[0]
	t := v.validateType(exprs[1])
	if t == typing.Unknown() || t == typing.Invalid() || t == nil {
		v.addError(exprs[1].Start(), errImpossibleCastToNonType)
		return left
	}

	if !typing.AssignableTo(left, t, false) {

		if exprs[0].Type() == ast.Literal {
			l := exprs[0].(*ast.LiteralNode)

			if l.LiteralType != token.String {

				num, ok := typing.ResolveUnderlying(t).(*typing.NumericType)
				if ok {
					if l.Data[0] == '-' {
						if num.Signed {
							return t
						}
					} else {
						// TODO: handle floats
						return t
					}
				}
			}
		}
		v.addError(exprs[1].Start(), errImpossibleCast, typing.WriteType(left), typing.WriteType(t))
		return t
	}
	return t
}

func BooleanOperator(v *Validator, types []typing.Type, exprs []ast.ExpressionNode) typing.Type {
	return typing.Boolean()
}

func (v *Validator) LargestNumericType(allowFloat bool) typing.Type {
	largest := -1
	largestType := typing.Type(typing.Unknown())
	for _, typ := range v.primitives {
		n, ok := typ.(*typing.NumericType)
		if ok {
			if !n.Integer && !allowFloat {
				continue
			}
			if largest == -1 || n.BitSize > largest {
				if n.Signed {
					// TODO: only select signed?
					largest = n.BitSize
					largestType = n
				}
			}
		}
	}
	return largestType
}

// SmallestNumericType ...
func (v *Validator) SmallestNumericType(bits int, allowFloat bool) typing.Type {
	smallest := -1
	smallestType := typing.Type(typing.Unknown())
	for _, typ := range v.primitives {
		n, ok := typ.(*typing.NumericType)
		if ok {
			if !n.Integer && !allowFloat {
				continue
			}
			if smallest == -1 || n.BitSize < smallest {
				if n.BitSize >= bits {
					if n.Signed {
						// TODO: only select signed?
						smallest = n.BitSize
						smallestType = n
					}

				}
			}
		}
	}
	return smallestType
}
