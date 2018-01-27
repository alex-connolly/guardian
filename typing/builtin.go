package typing

// NumericType ... used to pass types into the s
type NumericType struct {
	Mods    *Modifiers
	BitSize int
	Name    string
	Signed  bool
	Integer bool
}

func (nt *NumericType) AcceptsLiteral(t Type) bool {
	if other, ok := t.(*NumericType); ok {
		if !nt.Signed && other.Signed {
			return false
		}
		if nt.Integer && !other.Integer {
			return false
		}
		if nt.BitSize < other.BitSize {
			return false
		}
		return true
	}
	return false
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
