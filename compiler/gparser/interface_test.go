package gparser

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestParseNonExistentFile(t *testing.T) {
	ParseFile("tests/fake_contract.grd")
}

func TestParseString(t *testing.T) {
	p := ParseString("event Dog()")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
}

func TestParseStringDeclaration(t *testing.T) {
	p := ParseString("func hello() {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, p.Scope.Declarations != nil, "decls should not be nil")
}

func TestParseFile(t *testing.T) {
	ParseFile("tests/empty.grd")
}
