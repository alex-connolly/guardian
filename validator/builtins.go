package validator

import (
	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/lexer"
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

type LiteralFunc func(*Validator, string) typing.Type
type LiteralMap map[lexer.TokenType]LiteralFunc

func SimpleLiteral(typeName string) LiteralFunc {
	return func(v *Validator, data string) typing.Type {
		return v.getNamedType(typeName)
	}
}

func BooleanLiteral(v *Validator, data string) typing.Type {
	return typing.Boolean()
}

type OperatorFunc func(*Validator, ...typing.Type) typing.Type
type OperatorMap map[lexer.TokenType]OperatorFunc

func (m OperatorMap) Add(function OperatorFunc, types ...lexer.TokenType) {
	for _, t := range types {

		m[t] = function
	}
}

func SimpleOperator(typeName string) OperatorFunc {
	return func(v *Validator, types ...typing.Type) typing.Type {
		return v.getNamedType(typeName)
	}
}
func BinaryNumericOperator(v *Validator, ts ...typing.Type) typing.Type {
	left := typing.ResolveUnderlying(ts[0])
	right := typing.ResolveUnderlying(ts[1])
	if na, ok := left.(typing.NumericType); ok {
		if nb, ok := right.(typing.NumericType); ok {
			if na.BitSize > nb.BitSize {
				return v.SmallestNumericType(na.BitSize, true)
			}
			return v.SmallestNumericType(nb.BitSize, true)
		}
	}
	return typing.Invalid()
}

func BinaryIntegerOperator(v *Validator, ts ...typing.Type) typing.Type {
	if na, ok := ts[0].(typing.NumericType); ok && na.Integer {
		if nb, ok := ts[1].(typing.NumericType); ok && nb.Integer {
			if na.BitSize > nb.BitSize {
				return v.SmallestNumericType(na.BitSize, false)
			}
			return v.SmallestNumericType(nb.BitSize, false)
		}
	}
	return typing.Invalid()
}

func CastOperator(v *Validator, ts ...typing.Type) typing.Type {
	// pretend it's a valid type
	left := ts[0]
	right := ts[1]
	if !typing.AssignableTo(left, right) {
		v.addError(errImpossibleCast, typing.WriteType(left), typing.WriteType(right))
		return left
	}
	return right
}

func BooleanOperator(v *Validator, ts ...typing.Type) typing.Type {
	if len(ts) != 2 {

	}
	return typing.Boolean()
}

func (v *Validator) LargestNumericType(allowFloat bool) typing.Type {
	largest := -1
	largestType := typing.Type(typing.Unknown())
	for _, typ := range v.primitives {
		n, ok := typ.(typing.NumericType)
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

func (v *Validator) SmallestNumericType(bits int, allowFloat bool) typing.Type {
	smallest := -1
	smallestType := typing.Type(typing.Unknown())
	for _, typ := range v.primitives {
		n, ok := typ.(typing.NumericType)
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
