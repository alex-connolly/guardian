package validator

import (
	"testing"

	"github.com/end-r/guardian/typing"

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
	v := NewValidator(NewTestVM())
	goutil.Assert(t, v.requireType(typing.Boolean(), typing.Boolean()), "direct should be equal")
	v.DeclareType("a", typing.Boolean())
	goutil.Assert(t, v.requireType(typing.Boolean(), v.getNamedType("a")), "indirect should be equal")
}

func TestRequireTypeUnmatched(t *testing.T) {
	v := NewValidator(NewTestVM())
	goutil.Assert(t, !v.requireType(typing.Boolean(), typing.Unknown()), "direct should not be equal")
	v.DeclareType("a", typing.Unknown())
	goutil.Assert(t, !v.requireType(typing.Boolean(), v.getNamedType("a")), "indirect should not be equal")
}
