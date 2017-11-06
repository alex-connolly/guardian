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
	inherits(Type) bool
	implements(Type) bool
}

type BaseType int

const (
	Invalid BaseType = iota
	Unknown

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
	Float

	// aliases
	Byte = Uint8
)

type StandardType struct {
	name string
}

var standards = map[BaseType]StandardType{
	Invalid: StandardType{"invalid"},
	Unknown: StandardType{"unknown"},

	Int:    StandardType{"int"},
	Int8:   StandardType{"int8"},
	Int16:  StandardType{"int16"},
	Int32:  StandardType{"int32"},
	Int64:  StandardType{"int64"},
	Int128: StandardType{"int128"},
	Int256: StandardType{"int256"},

	Uint:    StandardType{"uint"},
	Uint8:   StandardType{"uint8"},
	Uint16:  StandardType{"uint16"},
	Uint32:  StandardType{"uint32"},
	Uint64:  StandardType{"uint64"},
	Uint128: StandardType{"uint128"},
	Uint256: StandardType{"uint256"},

	Float:  StandardType{"float"},
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
	name    string
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
	var types []Type
	for _, expression := range exprs {
		typ := v.resolveExpression(expression)
		// expression tuples force inner tuples to just be lists of types
		// ((int, string)) --> (int, string)
		// ((int), string) --> (int, string)
		// this is to facilitate assignment comparisons
		if tuple, ok := typ.(Tuple); ok {
			types = append(types, tuple.types...)
		} else {
			types = append(types, typ)
		}
	}
	t := NewTuple(types...)
	return t
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
	Name       string
	Supers     []Class
	Properties map[string]Type
	Interfaces []Interface
}

func NewClass(name string, properties map[string]Type, interfaces []Interface, supers []Class) Class {
	return Class{
		Name:       name,
		Supers:     supers,
		Properties: properties,
		Interfaces: interfaces,
	}
}

type Enum struct {
	Name   string
	Supers []Enum
	Items  map[string]bool
}

func NewEnum(name string, supers []Enum) Enum {
	return Enum{
		Name:   name,
		Supers: supers,
	}
}

type Interface struct {
	Name   string
	Supers []Interface
	Funcs  map[string]Func
}

func NewInterface(name string, funcs map[string]Func, supers []Interface) Interface {
	return Interface{
		Name:   name,
		Supers: supers,
		Funcs:  funcs,
	}
}

type Contract struct {
	Name       string
	Supers     []Contract
	Interfaces []Interface
}

func NewContract(name string, supers []Contract, interfaces []Interface) Contract {
	return Contract{
		Name:       name,
		Supers:     supers,
		Interfaces: interfaces,
	}
}

type Event struct {
	Name       string
	Parameters Tuple
}

func NewEvent(name string, params Tuple) Event {
	return Event{
		Name:       name,
		Parameters: params,
	}
}
