package typing

/*
Type system based on go's, with the following changes:

- prefer explicit typing to duck typing
*/

// Type interface used by all types
type Type interface {
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
