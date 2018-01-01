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

func TestValidateString(t *testing.T) {
	scope, errs := ValidateString(NewTestVM(), `
		if x = 0; x > 5 {

		} else {

		}
	`)

	goutil.AssertNow(t, errs == nil, errs.Format())
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
}

func TestValidateExpression(t *testing.T) {
	expr, errs := ValidateExpression(NewTestVM(), "5")

	goutil.AssertNow(t, errs == nil, errs.Format())
	goutil.AssertNow(t, expr != nil, "expr should not be nil")
}

func TestNewValidator(t *testing.T) {
	te := NewTestVM()
	v := NewValidator(te)
	goutil.AssertLength(t, len(v.operators), len(te.Operators()))
	goutil.AssertLength(t, len(v.literals), len(te.Literals()))
	goutil.AssertNow(t, len(v.primitives) > 0, "no primitives")
}
