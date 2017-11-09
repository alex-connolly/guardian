package validator

import (
	"fmt"
	"testing"

	"github.com/end-r/guardian/compiler/parser"

	"github.com/end-r/goutil"
)

func TestTypeValidateValid(t *testing.T) {
	p := parser.ParseString(`
            a Dog
            type Dog int
        `)
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	le := p.Scope.Declarations.Length()
	goutil.AssertNow(t, le == 2, fmt.Sprintf("wrong decl length: %d", le))
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestTypeValidateInvalid(t *testing.T) {
	p := parser.ParseString(`
            b Cat
            type Dog int
        `)
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	le := p.Scope.Declarations.Length()
	goutil.AssertNow(t, le == 2, fmt.Sprintf("wrong decl length: %d", le))
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, v.formatErrors())
}
