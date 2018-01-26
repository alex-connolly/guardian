package validator

import (
	"fmt"
	"testing"

	"github.com/end-r/guardian/typing"

	"github.com/end-r/guardian/token"

	"github.com/end-r/goutil"
)

func TestAdd(t *testing.T) {
	m := OperatorMap{}

	goutil.Assert(t, len(m) == 0, "wrong initial length")
	// numericalOperator with floats/ints

	m.Add(BinaryNumericOperator, token.Sub, token.Mul, token.Div)

	goutil.Assert(t, len(m) == 3, fmt.Sprintf("wrong added length: %d", len(m)))

	// integers only
	m.Add(BinaryIntegerOperator, token.Shl, token.Shr)

	goutil.Assert(t, len(m) == 5, fmt.Sprintf("wrong final length: %d", len(m)))
}

func TestImportVM(t *testing.T) {
	v := new(Validator)
	tvm := NewTestVM()
	v.importVM(tvm)
	goutil.AssertNow(t, len(v.errs) == 0, v.errs.Format())
	goutil.AssertNow(t, v.primitives != nil, "primitives should not be nil")
	typ, _ := v.isTypeVisible("address")
	goutil.AssertNow(t, typ != typing.Unknown(), "addr unrecognised")
}

func TestCastingValidToUnsigned(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), "x = uint(5)")
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestCastingInvalidToUnsigned(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), "x = uint(-5)")
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}
