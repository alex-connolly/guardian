package validator

// There are 5 first-class guardian types:
// Literal: int, string etc.
// Array: arrays[Type]
// NOTE: array = golang's slice, there is no golang array equivalent
// Map: map[Type]Type
// Func: func(Tuple)Tuple

// There is 1 second-class guardian type:
// Tuple: (Type...)

type Type interface {
	Base() BaseType
}

type StandardType int

const (
	Invalid StandardType = iota
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

func (s StandardType) Base() BaseType { return StandardBase }

type BaseType int

const (
	ArrayBase BaseType = iota
	FuncBase
	MapBase
	TupleBase
	StandardBase
)

type Array struct {
	Value Type
}

func (a Array) Base() BaseType { return ArrayBase }

type Map struct {
	Key   Type
	Value Type
}

func (m Map) Base() BaseType { return MapBase }

type Func struct {
	Params  Tuple
	Results Tuple
}

func (f Func) Base() BaseType { return FuncBase }

type Tuple struct {
	types []Type
}

func (t Tuple) Base() BaseType { return TupleBase }
