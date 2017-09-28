package validator

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestFindReference(t *testing.T) {
	v := NewValidator()
	v.DeclareType("a", standards[Int])
	goutil.Assert(t, v.findReference("a") == standards[Int], "did not find reference")
}
