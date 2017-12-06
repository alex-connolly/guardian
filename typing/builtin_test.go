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
