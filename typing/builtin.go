package typing

// NumericType ... used to pass types into the s
type NumericType struct {
	Mods    *Modifiers
	BitSize int
	Name    string
	Signed  bool
	Integer bool
}

func (nt *NumericType) AcceptsLiteral(length int, integer, hasSign bool) bool {
	if !nt.Signed && hasSign {
		return false
	}
	if nt.Integer && !integer {
		return false
	}
	if nt.BitSize < length {
		return false
	}
	return true
}

type VoidType struct {
	Mods *Modifiers
}

type BooleanType struct {
	Mods *Modifiers
}

// can probs use logs here
func BitsNeeded(x int) int {
	if x == 0 {
		return 1
	}
	count := 0
	for x > 0 {
		count++
		x = x >> 1
	}
	return count
}
