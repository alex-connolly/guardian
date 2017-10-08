package validator

import (
	"bytes"

	"github.com/end-r/guardian/compiler/ast"
)

// There are 5 first-class guardian types:
// Literal: int, string etc.
// Array: arrays[Type]
// NOTE: array = golang's slice, there is no golang array equivalent
// Map: map[Type]Type
// Func: func(Tuple)Tuple

// There are 2 second-class guardian types:
// Tuple: (Type...)
// Aliased: string -> Type

type Type interface {
	write(*bytes.Buffer)
	compare(Type) bool
}

type BaseType int

const (
	Invalid BaseType = iota
	Int
	Int8
	Int16
	Int32
	Int64
	Int128
	Int256

	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uint128
	Uint256

	String
	Bool

	// aliases
	Byte = Uint8
)

type StandardType struct {
	name string
}

var standards = map[BaseType]StandardType{
	Invalid: StandardType{"invalid"},
	Int:     StandardType{"int"},
	Int8:    StandardType{"int8"},
	Int16:   StandardType{"int16"},
	Int32:   StandardType{"int32"},
	Int64:   StandardType{"int64"},
	Int128:  StandardType{"int128"},
	Int256:  StandardType{"int256"},

	Uint:    StandardType{"uint"},
	Uint8:   StandardType{"uint8"},
	Uint16:  StandardType{"uint16"},
	Uint32:  StandardType{"uint32"},
	Uint64:  StandardType{"uint64"},
	Uint128: StandardType{"uint128"},
	Uint256: StandardType{"uint256"},

	String: StandardType{"string"},
	Bool:   StandardType{"bool"},
	// just an alias
}

type Array struct {
	Value Type
}

func NewArray(value Type) Array {
	return Array{
		Value: value,
	}
}

type Map struct {
	Key   Type
	Value Type
}

func NewMap(key, value Type) Map {
	return Map{
		Key:   key,
		Value: value,
	}
}

type Func struct {
	Params  Tuple
	Results Tuple
}

func NewFunc(params, results Tuple) Func {
	return Func{
		Params:  params,
		Results: results,
	}
}

type Tuple struct {
	types []Type
}

func NewTuple(types ...Type) Tuple {
	return Tuple{
		types: types,
	}
}

func (v *Validator) ExpressionTuple(exprs []ast.ExpressionNode) Tuple {
	types := make([]Type, len(exprs))
	for i, expression := range exprs {
		types[i] = v.resolveExpression(expression)
	}
	return NewTuple(types...)
}

type Aliased struct {
	alias      string
	underlying Type
}

func NewAliased(alias string, underlying Type) Aliased {
	return Aliased{
		alias:      alias,
		underlying: underlying,
	}
}

type Class struct {
}
