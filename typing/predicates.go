package typing

// AssignableTo checks whether a value of type 'right' can be assigned to a variable of type 'left'
func AssignableTo(left, right Type, allowUnknown bool) bool {

	if left == Unknown() && allowUnknown {
		return true
	}

	if rg, ok := right.(*Generic); ok {
		if rg.Accepts(left) {
			return true
		}
	}

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

	// ints --> larger ints
	// uints --> larger uints
	// uints --> larger ints
	if l, ok := ResolveUnderlying(left).(*NumericType); ok {
		if r, ok := ResolveUnderlying(right).(*NumericType); ok {
			if !r.Signed {
				if !l.Signed {
					// uints --> larger uints
					if l.BitSize >= r.BitSize {
						return true
					}
				} else {
					// uints --> larger ints
					if l.BitSize >= r.BitSize+1 {
						return true
					}
				}
			} else {
				if l.Signed {
					if l.BitSize >= r.BitSize {
						return true
					}
				}
			}
		}
	}

	return false
}
