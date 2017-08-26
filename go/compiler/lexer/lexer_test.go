package lexer

import (
	"fmt"
	"testing"

	"github.com/end-r/goutil"
)

func TestLexerTokenLength(t *testing.T) {
	l := LexString("hi this is me")
	goutil.Assert(t, len(l.Tokens) == 4, "wrong number of tokens")
}

func TestLexerAssignmentOperators(t *testing.T) {
	l := LexString("hi += 5")
	goutil.Assert(t, len(l.Tokens) == 3, "wrong number of tokens")
	goutil.Assert(t, l.Tokens[1].Type == TknAddAssign, "wrong operator type")
	l = LexString("hi -= 5")
	goutil.Assert(t, len(l.Tokens) == 3, "wrong number of tokens")
	goutil.Assert(t, l.Tokens[1].Type == TknSubAssign, "wrong operator type")
	l = LexString("hi *= 5")
	goutil.Assert(t, len(l.Tokens) == 3, "wrong number of tokens")
	goutil.Assert(t, l.Tokens[1].Type == TknMulAssign, "wrong operator type")
	l = LexString("hi /= 5")
	goutil.Assert(t, len(l.Tokens) == 3, "wrong number of tokens")
	goutil.Assert(t, l.Tokens[1].Type == TknDivAssign, "wrong operator type")
	l = LexString("hi %= 5")
	goutil.Assert(t, len(l.Tokens) == 3, "wrong number of tokens")
	goutil.Assert(t, l.Tokens[1].Type == TknModAssign, "wrong operator type")
}

func TestLexerLiterals(t *testing.T) {
	l := LexString(`x := "hello this is dog"`)
	goutil.Assert(t, len(l.Tokens) == 3, "wrong number of tokens")
	goutil.Assert(t, l.Tokens[2].Type == TknString, "wrong string literal type")
	goutil.Assert(t, l.TokenStringAndTrim(l.Tokens[2]) == "hello this is dog",
		fmt.Sprintf("wrong string produced: %s", l.TokenStringAndTrim(l.Tokens[2])))
	l = LexString(`x := 'a'`)
	goutil.Assert(t, len(l.Tokens) == 3, "wrong number of tokens")
	goutil.Assert(t, l.Tokens[2].Type == TknCharacter, "wrong character literal type")
	// test length
	l = LexString(`x := 6`)
	goutil.Assert(t, len(l.Tokens) == 3, "wrong number of tokens")
	goutil.Assert(t, l.Tokens[2].Type == TknNumber, "wrong integer literal type")
	l = LexString(`x := 5.5`)
	goutil.Assert(t, len(l.Tokens) == 3, "wrong number of tokens")
}

func TestLexerFiles(t *testing.T) {
	l := LexFile("tests/constants.grd")
	goutil.Assert(t, len(l.Tokens) == 19, fmt.Sprintf("wrong number of tokens (%d)", len(l.Tokens)))
	expected := []TokenType{
		TknContract, TknIdentifier, TknOpenBrace,
		TknConst, TknOpenBracket,
		TknIdentifier, TknAssign, TknNumber,
		TknIdentifier, TknAssign, TknString,
		TknIdentifier, TknAssign, TknCharacter,
		TknIdentifier, TknAssign, TknNumber,
		TknCloseBracket,
		TknCloseBrace,
	}
	for i, tok := range l.Tokens {
		goutil.Assert(t, tok.Type == expected[i], fmt.Sprintf("token type %d didn't match", i))
	}
}

func TestLexerError(t *testing.T) {
	l := LexString("")
	goutil.Assert(t, len(l.errors) == 0, "error len should be zero")
	l.error("this is an error")
	goutil.Assert(t, len(l.errors) == 1, "error len should be 1")
}
