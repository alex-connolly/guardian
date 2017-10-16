package validator

import (
	"axia/guardian/compiler/parser"
	"fmt"
	"testing"

	"github.com/end-r/goutil"
)

func TestValidateAssignmentValid(t *testing.T) {
	v := NewValidator()
	p := parser.ParseString(`
			a := 0
			a = 5
		`)
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	n := p.Scope.Next()
	v.validate(n)
	goutil.AssertNow(t, len(v.errors) == 0, fmt.Sprintf("wrong err length: %d", len(v.errors)))
}

func TestValidateAssignmentInvalid(t *testing.T) {
	v := NewValidator()
	p := parser.ParseString(`
			a := 0
			a = "hello world"
		`)
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	n := p.Scope.Next()
	v.validate(n)
	goutil.AssertNow(t, len(v.errors) == 1, fmt.Sprintf("wrong err length: %d", len(v.errors)))
}

func TestValidateForStatementValidCond(t *testing.T) {
	v := NewValidator()
	p := parser.ParseString("for a := 0; a < 5 {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(p.Scope.Sequence) == 1, "wrong sequence length")
	n := p.Scope.Next()
	v.validate(n)
	goutil.AssertNow(t, len(v.errors) == 1, fmt.Sprintf("wrong err length: %d", len(v.errors)))
}

/*func TestValidateForStatementInvalidCond(t *testing.T) {
	v := NewValidator()
	p := parser.ParseString("for a := 0; a + 5 {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(p.Scope.Sequence) == 1, "wrong sequence length")
	n := p.Scope.Next()
	v.validate(n)
	goutil.AssertNow(t, len(v.errors) == 1, fmt.Sprintf("wrong err length: %d", len(v.errors)))
}*/
