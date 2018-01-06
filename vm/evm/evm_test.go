package evm

import (
	"fmt"
	"testing"

	"github.com/end-r/goutil"
)

func TestBytesRequired(t *testing.T) {
	goutil.Assert(t, bytesRequired(1) == 1, fmt.Sprintf("wrong 1: %d", bytesRequired(1)))
	goutil.Assert(t, bytesRequired(8) == 1, fmt.Sprintf("wrong 8: %d", bytesRequired(8)))
	goutil.Assert(t, bytesRequired(257) == 2, fmt.Sprintf("wrong 257: %d", bytesRequired(257)))
}
