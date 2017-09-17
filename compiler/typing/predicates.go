package typing

type Predicate func(Type) bool

func isBoolean(t Type) bool {
	return false
}

func isString(t Type) bool {
	return false
}

func isNumeric(t Type) bool {
	return false
}

func isInteger(t Type) bool {
	return false
}

func isArrayIdentical(x, y Type) bool {

}

func isSliceIdentical(x, y Type) bool {
	if y, ok := y.(*Slice); ok {
		return identical(x.key, y.key)
	}
	return false
}

func isInterfaceIdentical(x, y Type) bool {
	return false
}

func isMapIdentical(x, y Type) bool {
	if y, ok := y.(*Map); ok {
		return identical(x.key, y.key) && identical(x.value, y.value)
	}
	return false
}

func isStructIdentical(x, y Type) bool {
	return false
}

func identical(x, y Type) bool {
	if x == y {
		return true
	}
	switch x := x.(type) {
	case *Array:
		return isArrayIdentical(x, y)
	case *Slice:
		return isSliceIdentical(x, y)
	case *Struct:
		return isStructIdentical(x, y)
	case *Interface:
		return isInterfaceIdentical(x, y)
	case *Map:
		return isMapIdentical(x, y)
	}
	return false
}
