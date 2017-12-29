package parser

import (
	"fmt"
	"testing"

	"github.com/end-r/goutil"
)

// mini-parser tests belong here

func TestHasTokens(t *testing.T) {
	p := createParser("this is data")
	goutil.Assert(t, len(p.tokens) == 3, "wrong length of tokens")
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

func TestParseSingleLineComment(t *testing.T) {
	p := createParser("// this is a comment")
	goutil.AssertNow(t, p.index == 0, "should start at 0")
	parseSingleLineComment(p)
	goutil.AssertNow(t, p.index == 5, "should finish at 5")

	p = createParser(`// this is a comment
		`)
	goutil.AssertNow(t, p.index == 0, "should start at 0")
	parseSingleLineComment(p)
	goutil.AssertNow(t, p.index == 6, "should finish at 6")
}

func TestParseMultiLineComment(t *testing.T) {
	p := createParser("/* this is a comment */")
	goutil.AssertNow(t, p.index == 0, "should start at 0")
	parseMultiLineComment(p)
	goutil.AssertNow(t, p.index == len(p.tokens), "should finish at end")

	p = createParser(`/* this is a comment

		*/`)
	goutil.AssertNow(t, p.index == 0, "should start at 0")
	parseMultiLineComment(p)
	goutil.AssertNow(t, p.index == len(p.tokens), "should finish at 7 end")

	p = createParser(`/* this is a comment
		fadnlkdlf a,s ds'


		d'sad
		sd
		ss

		dsdd
		a

		dsd
		s
		d
		class Dog {
			name string
		}
		*/`)
	goutil.AssertNow(t, p.index == 0, "should start at 0")
	parseMultiLineComment(p)
	goutil.AssertNow(t, p.index == len(p.tokens), "should finish at end")
}
