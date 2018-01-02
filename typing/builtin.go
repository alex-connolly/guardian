package typing

// NumericType ... used to pass types into the s
type NumericType struct {
	BitSize int
	Name    string
	Signed  bool
	Integer bool
}

type VoidType struct {
}

type BooleanType struct {
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
