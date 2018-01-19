package parser

import (
	"fmt"
	"testing"

	"github.com/end-r/goutil"
)

// mini-parser tests belong here

func TestHasTokens(t *testing.T) {
	p := createParser("this is data")
	goutil.Assert(t, len(p.lexer.Tokens) == 3, "wrong length of tokens")
	goutil.Assert(t, p.hasTokens(0), "should have 0 tokens")
	goutil.Assert(t, p.hasTokens(1), "should have 1 tokens")
	goutil.Assert(t, p.hasTokens(2), "should have 2 tokens")
	goutil.Assert(t, p.hasTokens(3), "should have 3 tokens")
	goutil.Assert(t, !p.hasTokens(4), "should not have 4 tokens")
}

func TestParseIdentifier(t *testing.T) {
	p := createParser("identifier")
	goutil.Assert(t, p.parseIdentifier() == "identifier", "wrong identifier")
	p = createParser("")
	goutil.Assert(t, p.parseIdentifier() == "", "empty should be nil")
	p = createParser("{")
	goutil.Assert(t, p.parseIdentifier() == "", "wrong token should be nil")
}

func TestParserNumDeclarations(t *testing.T) {
	a, _ := ParseString(`
		var b int
		var a string
	`)
	goutil.AssertNow(t, a != nil, "scope should not be nil")
	goutil.AssertNow(t, a.Declarations != nil, "scope declarations should not be nil")
	le := a.Declarations.Length()
	goutil.AssertNow(t, le == 2, fmt.Sprintf("wrong decl length: %d", le))
}

func TestParseSimpleStringAnnotationValid(t *testing.T) {
	_, errs := ParseString(`@Builtin("hello")`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestParseMultipleStringAnnotationValid(t *testing.T) {
	_, errs := ParseString(`@Builtin("hello", "world")`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestParseSingleIntegerAnnotationInvalid(t *testing.T) {
	_, errs := ParseString(`@Builtin(6)`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestParseMultipleIntegerAnnotationInvalid(t *testing.T) {
	_, errs := ParseString(`@Builtin(6, 6)`)
	goutil.AssertNow(t, len(errs) == 2, errs.Format())
}
