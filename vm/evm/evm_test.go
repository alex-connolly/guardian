package evm

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestBytesRequired(t *testing.T) {
	goutil.Assert(t, bytesRequired(1) == 1, "wrong 1")
	goutil.Assert(t, bytesRequired(8) == 1, "wrong 2")
	goutil.Assert(t, bytesRequired(9) == 2, "wrong 4")
	goutil.Assert(t, bytesRequired(16) == 2, "wrong 4")
}
