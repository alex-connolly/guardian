package validator

import (
	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/lexer"
)

// NumericType... used to pass types into the s
type NumericType struct {
	size    int
	name    string
	signed  bool
	integer bool
}

type BooleanType struct {
}

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

type LiteralFunc func(*Validator, string) Type
type LiteralMap map[lexer.TokenType]LiteralFunc

func SimpleLiteral(typeName string) LiteralFunc {
	return func(v *Validator, data string) Type {
		return v.getNamedType(typeName)
	}
}

type OperatorFunc func(*Validator, ...Type) Type
type OperatorMap map[lexer.TokenType]OperatorFunc

func (m OperatorMap) Add(function OperatorFunc, types ...lexer.TokenType) {
	for _, t := range types {
		m[t] = function
	}
}

func SimpleOperator(typeName string) OperatorFunc {
	return func(v *Validator, types ...Type) Type {
		return v.getNamedType(typeName)
	}
}

func (v *Validator) smallestNumericType(largerThan int) Type {
	// largerThan is the full size, not the size in bits
	// probably need to check the int64 conversion
	minBits := convertToBits(largerThan)
	smallest := -1
	smallestType := Type(standards[Unknown])
	for _, typ := range v.primitives {
		n, ok := typ.(NumericType)
		if ok {
			if smallest == -1 || n.size < smallest {
				if smallest > minBits {
					smallest = n.size
					smallestType = n
				}
			}
		}
	}
	return smallestType
}

// can probs use logs here
func convertToBits(x int) int {
	count := 0
	for x > 0 {
		count++
		x = x >> 1
	}
	return count
}
