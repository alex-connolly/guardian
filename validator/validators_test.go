package validator

import (
	"fmt"
	"testing"

	"github.com/end-r/guardian/parser"

	"github.com/end-r/goutil"
)

func TestTypeValidateValid(t *testing.T) {
	scope, _ := parser.ParseString(`
            a Dog
            type Dog int
        `)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	le := scope.Declarations.Length()
	goutil.AssertNow(t, le == 2, fmt.Sprintf("wrong decl length: %d", le))
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestTypeValidateInvalid(t *testing.T) {
	scope, _ := parser.ParseString(`
            b Cat
            type Dog int
        `)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	le := scope.Declarations.Length()
	goutil.AssertNow(t, le == 2, fmt.Sprintf("wrong decl length: %d", le))
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}
