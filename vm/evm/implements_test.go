package evm

import (
	"axia/guardian/vm"
	"testing"

	"github.com/end-r/goutil"
)

func TestImplements(t *testing.T) {
	var v vm.VM
	var e GuardianEVM
	v = e
	goutil.Assert(t, v.BooleanName() == e.BooleanName(), "this doesn't matter")
}
