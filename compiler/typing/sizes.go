package typing

// map types to their size in bytes
var defaultSizes = map[StandardType]int{
	Bool:   1,
	Int8:   1,
	Int16:  2,
	Int32:  4,
	Int64:  8,
	Uint8:  1,
	Uint16: 2,
	Uint32: 4,
	Uint64: 8,
}
