package validator

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestMakeName(t *testing.T) {
	// single name
	names := []string{"hi"}
	goutil.Assert(t, makeName(names) == "hi", "wrong single make name")
	names = []string{"hi", "you"}
	goutil.Assert(t, makeName(names) == "hi.you", "wrong multiple make name")
}

func TestRequireTypeMatched(t *testing.T) {
	v := NewValidator()
	goutil.Assert(t, v.requireType(standards[Bool], standards[Bool]), "direct should be equal")
	v.DeclareType("a", standards[Bool])
	goutil.Assert(t, v.requireType(standards[Bool], v.findReference("a")), "indirect should be equal")
}

func TestRequireTypeUnmatched(t *testing.T) {
	v := NewValidator()
	goutil.Assert(t, !v.requireType(standards[Bool], standards[String]), "direct should not be equal")
	v.DeclareType("a", standards[String])
	goutil.Assert(t, !v.requireType(standards[Bool], v.findReference("a")), "indirect should not be equal")
}
