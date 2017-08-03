package disasm

import (
	"axia/guardian/go/util"
	"testing"
)

func TestDisassembly(t *testing.T) {
	operations := DisasmString("600101")
	util.Assert(t, len(operations) == 1, "wrong number of operations")
}
