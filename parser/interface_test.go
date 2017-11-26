package parser

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestParseNonExistentFile(t *testing.T) {
	ParseFile("tests/fake_contract.grd")
}

func TestParseString(t *testing.T) {
	ast, _ := ParseString("event Dog()")
	goutil.AssertNow(t, ast != nil, "scope should not be nil")
}

func TestParseStringDeclaration(t *testing.T) {
	ast, _ := ParseString("func hello() {}")
	goutil.AssertNow(t, ast != nil, "scope should not be nil")
	goutil.AssertNow(t, ast.Declarations != nil, "decls should not be nil")
}

func TestParseFile(t *testing.T) {
	ParseFile("tests/empty.grd")
}
