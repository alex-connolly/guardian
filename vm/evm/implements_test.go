package evm

import (
	"testing"

	"github.com/end-r/guardian/validator"

	"github.com/end-r/goutil"
)

func TestImplements(t *testing.T) {
	var v validator.VM
	var e GuardianEVM
	v = e
	goutil.Assert(t, v.BooleanName() == e.BooleanName(), "this doesn't matter")
}
