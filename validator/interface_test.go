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
	goutil.AssertNow(t, expr.ResolvedType() != nil, "resolved is nil")
	_, ok := expr.ResolvedType().(*typing.NumericType)
	goutil.AssertNow(t, ok, "wrong type")
}

func TestNewValidator(t *testing.T) {
	te := NewTestVM()
	v := NewValidator(te)
	goutil.AssertLength(t, len(v.operators), len(operators()))
	goutil.AssertLength(t, len(v.literals), len(te.Literals()))
	goutil.AssertNow(t, len(v.primitives) > 0, "no primitives")
}

func TestDeclarationAndCall(t *testing.T) {
	te := NewTestVM()
	_, errs := ValidateString(te, `
		func hi(a bool){

		}
		hi(6 > 7)
		hi(false)
	`)
	goutil.AssertNow(t, errs == nil, errs.Format())
}

func TestValidatePackageFileData(t *testing.T) {
	a := `
	package x guardian 0.0.1
	class Animal {}
	`
	b := `
	package x guardian 0.0.1
	class Dog inherits Animal {}
	`
	errs := ValidateFileData(NewTestVM(), []string{a, b})
	goutil.AssertNow(t, errs == nil, errs.Format())
}
