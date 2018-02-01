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

// numeric operators:
// uint + uint = uint
// int + int = int
// float + any = float

func BinaryNumericOperator(v *Validator, types []typing.Type, exprs []ast.ExpressionNode) typing.Type {
	left := typing.ResolveUnderlying(types[0])
	right := typing.ResolveUnderlying(types[1])
	if na, ok := left.(*typing.NumericType); ok {
		if nb, ok := right.(*typing.NumericType); ok {
			// literals are handled differently
			// uint = uint + 1
			// should resolve to uint
			if lit, ok := exprs[0].(*ast.LiteralNode); ok {
				hasSign := (lit.Data[0] == '-')
				integral := (lit.LiteralType == token.Integer)
				if nb.AcceptsLiteral(na.BitSize, integral, hasSign) {
					return nb
				}
			}
			if lit, ok := exprs[1].(*ast.LiteralNode); ok {
				hasSign := (lit.Data[0] == '-')
				integral := (lit.LiteralType == token.Integer)
				if na.AcceptsLiteral(nb.BitSize, integral, hasSign) {
					return na
				}
			}

			if na.BitSize > nb.BitSize {
				if na.Integer && !nb.Integer {
					return v.SmallestFloat(na.BitSize)
				}
				if !na.Signed && nb.Signed {
					return v.SmallestInteger(na.BitSize, true)
				}
				return na
			}
			if nb.Integer && !na.Integer {
				return v.SmallestFloat(nb.BitSize)
			}
			if !nb.Signed && na.Signed {
				return v.SmallestInteger(nb.BitSize, true)
			}
			return nb
		}
	}
	return typing.Invalid()
}

func BinaryIntegerOperator(v *Validator, types []typing.Type, exprs []ast.ExpressionNode) typing.Type {
	if na, ok := types[0].(*typing.NumericType); ok && na.Integer {
		if nb, ok := types[1].(*typing.NumericType); ok && nb.Integer {
			if na.BitSize > nb.BitSize {
				if !na.Signed && nb.Signed {
					return v.SmallestInteger(na.BitSize, true)
				}
				return na
			}
			if !nb.Signed && na.Signed {
				return v.SmallestInteger(nb.BitSize, true)
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
func (v *Validator) SmallestFloat(bits int) typing.Type {
	smallest := -1
	smallestType := typing.Type(typing.Invalid())
	for _, typ := range v.primitives {
		n, ok := typ.(*typing.NumericType)
		if ok {
			if n.Integer {
				continue
			}
			if smallest == -1 || n.BitSize < smallest {
				if n.BitSize >= bits {
					smallest = n.BitSize
					smallestType = n
				}
			}
		}
	}
	return smallestType
}

// SmallestNumericType ...
func (v *Validator) SmallestInteger(bits int, isSigned bool) typing.Type {
	smallest := -1
	smallestType := typing.Type(typing.Unknown())
	for _, typ := range v.primitives {
		n, ok := typ.(*typing.NumericType)
		if ok {
			if !n.Integer {
				continue
			}
			// must be both
			if (n.Signed && !isSigned) || (!n.Signed && isSigned) {
				continue
			}
			if smallest == -1 || n.BitSize < smallest {
				if n.BitSize >= bits {
					smallest = n.BitSize
					smallestType = n
				}
			}
		}
	}
	return smallestType
}
