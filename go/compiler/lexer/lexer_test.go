package lexer

import (
	"axia/guardian/go/util"
	"fmt"
	"testing"
)

func TestLexerTokenLength(t *testing.T) {
	n, e := LexString("")
	util.Assert(t, len(n) == 0 && len(e) == 0, "empty string failed")
	n, e = LexString(`6`)
	util.Assert(t, len(n) == 1 && len(e) == 0, "number failed")
	n, e = LexString(`hi this is`)
	util.Assert(t, len(n) == 3 && len(e) == 0, "ids failed")
	n, e = LexString(`"hello`)
	util.Assert(t, len(n) == 1 && len(e) == 1, "unclosed string failed")
	n, e = LexString(`contract`)
	util.Assert(t, len(n) == 1 && len(e) == 0, "contract decl 1 failed")
	n, e = LexString(`contract Hello`)
	fmt.Println(n)
	util.Assert(t, len(n) == 2 && len(e) == 0, "contract decl 2 failed")
	n, e = LexString(`contract Hello {`)
	util.Assert(t, len(n) == 3 && len(e) == 0, "contract decl 3 failed")
}
