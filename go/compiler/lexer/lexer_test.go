package lexer

import (
	"fmt"
	"testing"

	"github.com/end-r/guardian/go/util"
)

func TestLexerTokenLength(t *testing.T) {
	l := LexString("hi this is me")
	util.Assert(t, len(l.Tokens) == 4, "wrong number of tokens")
}

func TestLexerAssignmentOperators(t *testing.T) {
	l := LexString("hi += 5")
	util.Assert(t, len(l.Tokens) == 3, "wrong number of tokens")
	util.Assert(t, l.Tokens[1].Type == TknAddAssign, "wrong operator type")
	l = LexString("hi -= 5")
	util.Assert(t, len(l.Tokens) == 3, "wrong number of tokens")
	util.Assert(t, l.Tokens[1].Type == TknSubAssign, "wrong operator type")
	l = LexString("hi *= 5")
	util.Assert(t, len(l.Tokens) == 3, "wrong number of tokens")
	util.Assert(t, l.Tokens[1].Type == TknMulAssign, "wrong operator type")
	l = LexString("hi /= 5")
	util.Assert(t, len(l.Tokens) == 3, "wrong number of tokens")
	util.Assert(t, l.Tokens[1].Type == TknDivAssign, "wrong operator type")
	l = LexString("hi %= 5")
	util.Assert(t, len(l.Tokens) == 3, "wrong number of tokens")
	util.Assert(t, l.Tokens[1].Type == TknModAssign, "wrong operator type")
}

func TestLexerLiterals(t *testing.T) {
	l := LexString(`x := "hello this is dog"`)
	util.Assert(t, len(l.Tokens) == 3, "wrong number of tokens")
	util.Assert(t, l.Tokens[2].Type == TknString, "wrong string literal type")
	util.Assert(t, l.TokenStringAndTrim(l.Tokens[2]) == "hello this is dog",
		fmt.Sprintf("wrong string produced: %s", l.TokenStringAndTrim(l.Tokens[2])))
	l = LexString(`x := 'a'`)
	util.Assert(t, len(l.Tokens) == 3, "wrong number of tokens")
	util.Assert(t, l.Tokens[2].Type == TknCharacter, "wrong character literal type")
	// test length
	l = LexString(`x := 6`)
	util.Assert(t, len(l.Tokens) == 3, "wrong number of tokens")
	util.Assert(t, l.Tokens[2].Type == TknNumber, "wrong integer literal type")
	l = LexString(`x := 5.5`)
	util.Assert(t, len(l.Tokens) == 3, "wrong number of tokens")
}

func TestLexerFiles(t *testing.T) {
	l := LexFile("tests/constants.grd")
	util.Assert(t, len(l.Tokens) == 19, fmt.Sprintf("wrong number of tokens (%d)", len(l.Tokens)))
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
		util.Assert(t, tok.Type == expected[i], fmt.Sprintf("token type %d didn't match", i))
	}
}
