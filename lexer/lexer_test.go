package lexer

import (
	"fmt"
	"testing"

	"github.com/end-r/goutil"
)

func TestLexFileNonExistent(t *testing.T) {
	_, errs := LexFile("fake_file.grd")
	goutil.Assert(t, errs != nil, "should error")
}

func TestLexUnrecognisedToken(t *testing.T) {
	LexString("~")
}

func TestLexerTokenLength(t *testing.T) {
	tokens, _ := LexString("hi this is me")
	goutil.Assert(t, len(tokens) == 4, "wrong number of tokens")
}

func TestLexerAssignmentOperators(t *testing.T) {
	tokens, _ := LexString("hi += 5")
	goutil.Assert(t, len(tokens) == 3, "wrong number of tokens")
	goutil.Assert(t, tokens[1].Type == TknAddAssign, "wrong operator type")
	tokens, _ = LexString("hi -= 5")
	goutil.Assert(t, len(tokens) == 3, "wrong number of tokens")
	goutil.Assert(t, tokens[1].Type == TknSubAssign, "wrong operator type")
	tokens, _ = LexString("hi *= 5")
	goutil.Assert(t, len(tokens) == 3, "wrong number of tokens")
	goutil.Assert(t, tokens[1].Type == TknMulAssign, "wrong operator type")
	tokens, _ = LexString("hi /= 5")
	goutil.Assert(t, len(tokens) == 3, "wrong number of tokens")
	goutil.Assert(t, tokens[1].Type == TknDivAssign, "wrong operator type")
	tokens, _ = LexString("hi %= 5")
	goutil.Assert(t, len(tokens) == 3, "wrong number of tokens")
	goutil.Assert(t, tokens[1].Type == TknModAssign, "wrong operator type")
}

func TestLexerLiterals(t *testing.T) {
	tokens, _ := LexString(`x := "hello this is dog"`)
	goutil.Assert(t, len(tokens) == 3, "wrong number of tokens")
	goutil.Assert(t, tokens[2].Type == TknString, "wrong string literal type")
	goutil.Assert(t, tokens[2].TrimTokenString() == "hello this is dog",
		fmt.Sprintf("wrong string produced: %s", tokens[2].TrimTokenString()))
	tokens, _ = LexString(`x := 'a'`)
	goutil.Assert(t, len(tokens) == 3, "wrong number of tokens")
	goutil.Assert(t, tokens[2].Type == TknCharacter, "wrong character literal type")
	// test length
	tokens, _ = LexString(`x := 6`)
	goutil.Assert(t, len(tokens) == 3, "wrong number of tokens")
	goutil.Assert(t, tokens[2].Type == TknInteger, "wrong integer literal type")
	tokens, _ = LexString(`x := 5.5`)
	goutil.Assert(t, len(tokens) == 3, "wrong number of tokens")
}

func TestLexerFileConstants(t *testing.T) {
	tokens, _ := LexFile("tests/constants.grd")
	checkTokens(t, tokens, []TokenType{
		TknContract, TknIdentifier, TknOpenBrace, TknNewLine,
		TknNewLine,
		TknConst, TknOpenBracket, TknNewLine,
		TknIdentifier, TknAssign, TknInteger, TknNewLine,
		TknIdentifier, TknAssign, TknString, TknNewLine,
		TknIdentifier, TknAssign, TknCharacter, TknNewLine,
		TknIdentifier, TknAssign, TknFloat, TknNewLine,
		TknCloseBracket, TknNewLine,
		TknNewLine,
		TknCloseBrace, TknNewLine,
	})
}

func TestLexerFileDeclarations(t *testing.T) {
	tokens, _ := LexString(`
		contract Dog {
		    class Hi {

		    }
		    interface XX {

		    }
		    event Hello(string)
		}
	`)
	checkTokens(t, tokens, []TokenType{
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
	tokens, _ := LexString("hello")
	goutil.AssertNow(t, len(tokens) == 1, "wrong token length")
	tokens, _ = LexString("hello.dog")
	goutil.AssertNow(t, len(tokens) == 3, "wrong token length")
	tokens, _ = LexString("hello.dog.cat")
	goutil.AssertNow(t, len(tokens) == 5, "wrong token length")
}

func TestLexerError(t *testing.T) {
	_, errs := LexString("")
	goutil.Assert(t, len(errs) == 0, "error len should be zero")
}

func TestLexerType(t *testing.T) {
	tokens, _ := LexString("type Number int")
	checkTokens(t, tokens, []TokenType{TknType, TknIdentifier, TknIdentifier})
}

func TestLexerIncrement(t *testing.T) {
	tokens, _ := LexString("x++")
	checkTokens(t, tokens, []TokenType{TknIdentifier, TknIncrement})
	tokens, _ = LexString("x--")
	checkTokens(t, tokens, []TokenType{TknIdentifier, TknDecrement})
}

func TestLexerComparators(t *testing.T) {
	tokens, _ := LexString("> < >= <= ==")
	checkTokens(t, tokens, []TokenType{TknGtr, TknLss, TknGeq, TknLeq, TknEql})
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

func TestLexerModifiers(t *testing.T) {
	tokens, _ := LexString("protected abstract external internal")
	checkTokens(t, tokens, []TokenType{TknProtected, TknAbstract, TknExternal, TknInternal})
}

func TestLexerExpVar(t *testing.T) {
	tokens, _ := LexString("x string")
	checkTokens(t, tokens, []TokenType{TknIdentifier, TknIdentifier})
}

func TestLexerIntegers(t *testing.T) {
	tokens, _ := LexString("x = 555")
	checkTokens(t, tokens, []TokenType{TknIdentifier, TknAssign, TknInteger})
}

func TestLexerFloats(t *testing.T) {
	tokens, _ := LexString("x = 5.55")
	checkTokens(t, tokens, []TokenType{TknIdentifier, TknAssign, TknFloat})
	tokens, _ = LexString("x = .55")
	checkTokens(t, tokens, []TokenType{TknIdentifier, TknAssign, TknFloat})
}

func TestLexerIsFloat(t *testing.T) {
	l := Lexer{buffer: []byte("5.55")}
	goutil.Assert(t, isFloat(&l), "full float isn't recognised")
	l = Lexer{buffer: []byte(".55")}
	goutil.Assert(t, isFloat(&l), "no start float isn't recognised")
}
