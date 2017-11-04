package validator

import (
	"axia/guardian/compiler/parser"
	"testing"

	"github.com/end-r/goutil"
)

func TestValidateAssignmentValid(t *testing.T) {

	p := parser.ParseString(`
			a := 0
			a = 5
		`)
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestValidateAssignmentInvalid(t *testing.T) {
	p := parser.ParseString(`
			a := 0
			a = "hello world"
		`)
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, v.formatErrors())
}

func TestValidateForStatementValidCond(t *testing.T) {
	p := parser.ParseString("for a := 0; a < 5 {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(p.Scope.Sequence) == 1, "wrong sequence length")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestValidateForStatementInvalidCond(t *testing.T) {
	p := parser.ParseString("for a := 0; a + 5 {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(p.Scope.Sequence) == 1, "wrong sequence length")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, v.formatErrors())
}
