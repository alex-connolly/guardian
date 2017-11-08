package validator

import (
	"testing"

	"github.com/end-r/guardian/compiler/parser"

	"github.com/end-r/goutil"
)

func TestTypeScanValid(t *testing.T) {
	p := parser.ParseString(`
            a Dog
            type Dog int
        `)
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, p.Scope.Declarations.Length() == 2, "wrong decl length")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestTypeScanInvalid(t *testing.T) {
	p := parser.ParseString(`
            b Cat
            type Dog int
        `)
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, v.formatErrors())
}
