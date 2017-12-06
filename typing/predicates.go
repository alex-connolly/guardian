package typing

func assignableTo(left, right Type) bool {
	// assignable if the two types are equal
	if left.Compare(right) {
		return true
	}
	// assignable if o implements t
	if left.implements(right) {
		return true
	}
	// assignable if t is a superclass of o
	if left.inherits(right) {
		return true
	}

	return false
}
