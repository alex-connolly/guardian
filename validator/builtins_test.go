package validator

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestConvertToBits(t *testing.T) {
	goutil.Assert(t, convertToBits(0) == 1, "wrong 0")
	goutil.Assert(t, convertToBits(1) == 1, "wrong 1")
	goutil.Assert(t, convertToBits(2) == 2, "wrong 2")
	goutil.Assert(t, convertToBits(10) == 4, "wrong 10")
}

func TestSmallestNumericType(t *testing.T) {

}
