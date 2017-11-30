package validator

import (
	"fmt"
	"testing"

	"github.com/end-r/guardian/lexer"

	"github.com/end-r/goutil"
)

func TestConvertToBits(t *testing.T) {
	goutil.Assert(t, BitsNeeded(0) == 1, "wrong 0")
	goutil.Assert(t, BitsNeeded(1) == 1, "wrong 1")
	goutil.Assert(t, BitsNeeded(2) == 2, "wrong 2")
	goutil.Assert(t, BitsNeeded(10) == 4, "wrong 10")
}

func TestAdd(t *testing.T) {
	m := OperatorMap{}

	goutil.Assert(t, len(m) == 0, "wrong initial length")
	// numericalOperator with floats/ints

	m.Add(BinaryNumericOperator(), lexer.TknSub, lexer.TknMul, lexer.TknDiv)

	goutil.Assert(t, len(m) == 3, fmt.Sprintf("wrong added length: %d", len(m)))

	// integers only
	m.Add(BinaryIntegerOperator(), lexer.TknShl, lexer.TknShr)

	goutil.Assert(t, len(m) == 5, fmt.Sprintf("wrong final length: %d", len(m)))
}
