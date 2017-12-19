package lexer

import (
	"fmt"
	"testing"

	"github.com/end-r/guardian/token"

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
	goutil.Assert(t, tokens[1].Type == token.AddAssign, "wrong operator type")
	tokens, _ = LexString("hi -= 5")
	goutil.Assert(t, len(tokens) == 3, "wrong number of tokens")
	goutil.Assert(t, tokens[1].Type == token.SubAssign, "wrong operator type")
	tokens, _ = LexString("hi *= 5")
	goutil.Assert(t, len(tokens) == 3, "wrong number of tokens")
	goutil.Assert(t, tokens[1].Type == token.MulAssign, "wrong operator type")
	tokens, _ = LexString("hi /= 5")
	goutil.Assert(t, len(tokens) == 3, "wrong number of tokens")
	goutil.Assert(t, tokens[1].Type == token.DivAssign, "wrong operator type")
	tokens, _ = LexString("hi %= 5")
	goutil.Assert(t, len(tokens) == 3, "wrong number of tokens")
	goutil.Assert(t, tokens[1].Type == token.ModAssign, "wrong operator type")
}

func TestLexerLiterals(t *testing.T) {
	tokens, _ := LexString(`x := "hello this is dog"`)
	goutil.Assert(t, len(tokens) == 3, "wrong number of tokens")
	goutil.Assert(t, tokens[2].Type == token.String, "wrong string literal type")
	goutil.Assert(t, tokens[2].TrimmedString() == "hello this is dog",
		fmt.Sprintf("wrong string produced: %s", tokens[2].TrimmedString()))
	tokens, _ = LexString(`x := 'a'`)
	goutil.Assert(t, len(tokens) == 3, "wrong number of tokens")
	goutil.Assert(t, tokens[2].Type == token.Character, "wrong character literal type")
	// test length
	tokens, _ = LexString(`x := 6`)
	goutil.Assert(t, len(tokens) == 3, "wrong number of tokens")
	goutil.Assert(t, tokens[2].Type == token.Integer, "wrong integer literal type")
	tokens, _ = LexString(`x := 5.5`)
	goutil.Assert(t, len(tokens) == 3, "wrong number of tokens")
}

func TestLexerFileConstants(t *testing.T) {
	tokens, _ := LexFile("tests/constants.grd")
	checkTokens(t, tokens, []token.Type{
		token.Contract, token.Identifier, token.OpenBrace, token.NewLine,
		token.NewLine,
		token.Const, token.OpenBracket, token.NewLine,
		token.Identifier, token.Assign, token.Integer, token.NewLine,
		token.Identifier, token.Assign, token.String, token.NewLine,
		token.Identifier, token.Assign, token.Character, token.NewLine,
		token.Identifier, token.Assign, token.Float, token.NewLine,
		token.CloseBracket, token.NewLine,
		token.NewLine,
		token.CloseBrace, token.NewLine,
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
	checkTokens(t, tokens, []token.Type{
		token.NewLine,
		token.Contract, token.Identifier, token.OpenBrace, token.NewLine,
		token.Class, token.Identifier, token.OpenBrace, token.NewLine,
		token.NewLine,
		token.CloseBrace, token.NewLine,
		token.Interface, token.Identifier, token.OpenBrace, token.NewLine,
		token.NewLine,
		token.CloseBrace, token.NewLine,
		token.Event, token.Identifier, token.OpenBracket, token.Identifier, token.CloseBracket, token.NewLine,
		token.CloseBrace, token.NewLine})
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
	checkTokens(t, tokens, []token.Type{token.KWType, token.Identifier, token.Identifier})
}

func TestLexerIncrement(t *testing.T) {
	tokens, _ := LexString("x++")
	checkTokens(t, tokens, []token.Type{token.Identifier, token.Increment})
	tokens, _ = LexString("x--")
	checkTokens(t, tokens, []token.Type{token.Identifier, token.Decrement})
}

func TestLexerComparators(t *testing.T) {
	tokens, _ := LexString("> < >= <= ==")
	checkTokens(t, tokens, []token.Type{token.Gtr, token.Lss, token.Geq, token.Leq, token.Eql})
}

func TestHasByte(t *testing.T) {
	text := "interface"
	l := new(Lexer)
	l.buffer = []byte(text)
	l.byteOffset = len(text)
}

func TestLexerModifiers(t *testing.T) {
	tokens, _ := LexString("protected abstract external internal")
	checkTokens(t, tokens, []token.Type{token.Protected, token.Abstract, token.External, token.Internal})
}

func TestLexerExpVar(t *testing.T) {
	tokens, _ := LexString("x string")
	checkTokens(t, tokens, []token.Type{token.Identifier, token.Identifier})
}

func TestLexerIntegers(t *testing.T) {
	tokens, _ := LexString("x = 555")
	checkTokens(t, tokens, []token.Type{token.Identifier, token.Assign, token.Integer})
}

func TestLexerFloats(t *testing.T) {
	tokens, _ := LexString("x = 5.55")
	checkTokens(t, tokens, []token.Type{token.Identifier, token.Assign, token.Float})
	tokens, _ = LexString("x = .55")
	checkTokens(t, tokens, []token.Type{token.Identifier, token.Assign, token.Float})
}

func TextLexerGeneric(t *testing.T) {
	tokens, _ := LexString("<T|S|R>")
	checkTokens(t, tokens, []token.Type{
		token.Lss,
		token.Identifier, token.Or,
		token.Identifier, token.Or,
		token.Identifier,
		token.Gtr,
	})
}
