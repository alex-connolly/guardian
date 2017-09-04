package lexer

import (
	"fmt"
	"testing"

	"github.com/end-r/goutil"
)

func TestLexFileNonExistent(t *testing.T) {
	l := LexFile("fake_file.grd")
	goutil.Assert(t, l == nil, "lexer should be nil")
}

func TestLexUnrecognisedToken(t *testing.T) {
	LexString("~")
}

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

func TestLexerFileConstants(t *testing.T) {
	l := LexFile("tests/constants.grd")
	checkTokens(t, l.Tokens, []TokenType{
		TknContract, TknIdentifier, TknOpenBrace, TknNewLine,
		TknNewLine,
		TknConst, TknOpenBracket, TknNewLine,
		TknIdentifier, TknAssign, TknNumber, TknNewLine,
		TknIdentifier, TknAssign, TknString, TknNewLine,
		TknIdentifier, TknAssign, TknCharacter, TknNewLine,
		TknIdentifier, TknAssign, TknNumber, TknNewLine,
		TknCloseBracket, TknNewLine,
		TknNewLine,
		TknCloseBrace, TknNewLine,
	})
}

func TestLexerFileDeclarations(t *testing.T) {
	l := LexString(`
		contract Dog {
		    class Hi {

		    }
		    interface XX {

		    }
		    event Hello(string)
		}
	`)
	checkTokens(t, l.Tokens, []TokenType{
		TknNewLine,
		TknContract, TknIdentifier, TknOpenBrace, TknNewLine,
		TknClass, TknIdentifier, TknOpenBrace, TknNewLine,
		TknNewLine,
		TknCloseBrace, TknNewLine,
		TknInterface, TknIdentifier, TknOpenBrace, TknNewLine,
		TknNewLine,
		TknCloseBrace, TknNewLine,
		TknEvent, TknIdentifier, TknOpenBracket, TknIdentifier, TknCloseBracket, TknNewLine,
		TknCloseBrace, TknNewLine})
}

func TestLexerReference(t *testing.T) {
	l := LexString("hello")
	goutil.AssertNow(t, len(l.Tokens) == 1, "wrong token length")
	l = LexString("hello.dog")
	goutil.AssertNow(t, len(l.Tokens) == 3, "wrong token length")
	l = LexString("hello.dog.cat")
	goutil.AssertNow(t, len(l.Tokens) == 5, "wrong token length")
}

func TestLexerFloat(t *testing.T) {
	l := LexString("5.5")
	goutil.AssertNow(t, len(l.Tokens) == 1, "wrong token length")
	l = LexString("5.5.5")
	goutil.AssertNow(t, len(l.Tokens) == 3, "wrong token length")
}

func TestLexerError(t *testing.T) {
	l := LexString("")
	goutil.Assert(t, len(l.errors) == 0, "error len should be zero")
	l.error("this is an error")
	goutil.Assert(t, len(l.errors) == 1, "error len should be 1")
}

func TestLexerType(t *testing.T) {
	l := LexString("type Number int")
	checkTokens(t, l.Tokens, []TokenType{TknType, TknIdentifier, TknIdentifier})
}

func TestHasByte(t *testing.T) {
	text := "interface"
	l := new(Lexer)
	l.buffer = []byte(text)
	l.byteOffset = len(text)
	goutil.Assert(t, !l.hasBytes(1), "end should not have bytes")
	l.byteOffset = 0
	goutil.Assert(t, l.hasBytes(1), "start should have bytes")
}
