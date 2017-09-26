package typing

import "github.com/end-r/guardian/compiler/ast"

/*
Type system based on go's, with the following changes:

- prefer explicit typing to duck typing
*/

// Type interface used by all types
type Type interface {
	Underlying() Type
	Representation() string
}

type StandardType int

const (
	Int StandardType = iota
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

type Func struct {
	params  Tuple
	results Tuple
}

func (f Func) Representation() string {
	return nil
}

func (f Func) Underlying() Type {
	return nil
}

type Tuple struct {
	types []Type // order is important
}

func varsTuple(nodes []ast.Node) Tuple {
	types := make([]Type, 0)
	for _, n := range nodes {

	}
	return Tuple{
		types: types,
	}
}

func exprTuple(nodes []ast.ExpressionNode) Tuple {
	types := make([]Type, 0)
	for _, n := range nodes {
		types = append(types, ResolveExpression(e))
	}
	return Tuple{
		types: types,
	}
}

func (t *Tuple) equals(other Tuple) bool {
	if len(t.types) != len(other.types) {
		return false
	}
	for i, typ := range t.types {
		if typ != other.types[i] {
			return false
		}
	}
	return true
}

// Array type
type Array struct {
	length int64
	key    Type
}

// NewArray creates an array from the specified parameters
func NewArray(key Type, length int64) *Array {
	return &Array{
		length: length,
		key:    key,
	}
}

// Slice ...
type Slice struct {
	key Type
}

// NewSlice creates an array from the specified parameters
func NewSlice(key Type) *Slice {
	return &Slice{
		key: key,
	}
}

// Map ...
type Map struct {
	key, value Type
}

// NewMap creates an array from the specified parameters
func NewMap(key, value Type) *Map {
	return &Map{
		key:   key,
		value: value,
	}
}
