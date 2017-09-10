package parser

import "testing"

func TestParseNonExistentFile(t *testing.T) {
	ParseFile("tests/fake_contract.grd")
}

func TestParseString(t *testing.T) {
	ParseString("contract Dog {}")
}

func TestParseFile(t *testing.T) {
	ParseFile("tests/empty.grd")
}
