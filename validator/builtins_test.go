package validator

import (
	"fmt"
	"testing"

	"github.com/end-r/guardian/lexer"

	"github.com/end-r/goutil"
)

func TestAdd(t *testing.T) {
	m := OperatorMap{}

	goutil.Assert(t, len(m) == 0, "wrong initial length")
	// numericalOperator with floats/ints

	m.Add(BinaryNumericOperator, lexer.TknSub, lexer.TknMul, lexer.TknDiv)

	goutil.Assert(t, len(m) == 3, fmt.Sprintf("wrong added length: %d", len(m)))

	// integers only
	m.Add(BinaryIntegerOperator, lexer.TknShl, lexer.TknShr)

	goutil.Assert(t, len(m) == 5, fmt.Sprintf("wrong final length: %d", len(m)))
}
