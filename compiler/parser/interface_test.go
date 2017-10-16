package parser

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

func TestParseFile(t *testing.T) {
	ParseFile("tests/empty.grd")
}
