package typing

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestConvertToBits(t *testing.T) {
	goutil.Assert(t, BitsNeeded(0) == 1, "wrong 0")
	goutil.Assert(t, BitsNeeded(1) == 1, "wrong 1")
	goutil.Assert(t, BitsNeeded(2) == 2, "wrong 2")
	goutil.Assert(t, BitsNeeded(10) == 4, "wrong 10")
}

func TestAcceptLiteral(t *testing.T) {
	a := &NumericType{BitSize: 8, Signed: true, Integer: true}
	goutil.AssertNow(t, a.AcceptsLiteral(8, true, true), "1 failed")
	goutil.AssertNow(t, a.AcceptsLiteral(8, true, false), "2 failed")
	goutil.AssertNow(t, a.AcceptsLiteral(7, true, true), "3 failed")
	b := &NumericType{BitSize: 256, Signed: false, Integer: true}
	goutil.AssertNow(t, b.AcceptsLiteral(8, true, false), "4 failed")
}
