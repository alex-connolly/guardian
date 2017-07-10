package lexer

import (
	"axia/guardian/go/util"
	"testing"
)

func TestIsString(t *testing.T) {
	l := lex([]byte(":"))
	util.Assert(t, is(":")(l), "failed is string test for operator")
}
