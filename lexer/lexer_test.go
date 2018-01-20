package lexer

import (
	"testing"

	"github.com/end-r/guardian/token"

	"github.com/end-r/goutil"
)

func TestLexFileNonExistent(t *testing.T) {
	l := LexFile("fake_file.grd")
	goutil.AssertNow(t, l.Errors != nil, "should error")
}

func TestLexUnrecognisedToken(t *testing.T) {
	LexString("~")
}

func TestLexerTokenLength(t *testing.T) {
	l := LexString("hi this is me")
	goutil.AssertNow(t, len(l.Tokens) == 4, "wrong number of tokens")
}

func TestLexerAssignmentOperators(t *testing.T) {
	l := LexString("hi += 5")
	goutil.AssertNow(t, len(l.Tokens) == 3, "1 wrong number of tokens")
	goutil.AssertNow(t, l.Tokens[1].Type == token.AddAssign, "wrong operator type")
	l = LexString("hi -= 5")
	goutil.AssertNow(t, len(l.Tokens) == 3, "2 wrong number of tokens")
	goutil.AssertNow(t, l.Tokens[1].Type == token.SubAssign, "wrong operator type")
	l = LexString("hi *= 5")
	goutil.AssertNow(t, len(l.Tokens) == 3, "3 wrong number of tokens")
	goutil.AssertNow(t, l.Tokens[1].Type == token.MulAssign, "wrong operator type")
	l = LexString("hi /= 5")
	goutil.AssertNow(t, len(l.Tokens) == 3, "wrong number of tokens")
	goutil.AssertNow(t, l.Tokens[1].Type == token.DivAssign, "wrong operator type")
	l = LexString("hi %= 5")
	goutil.AssertNow(t, len(l.Tokens) == 3, "wrong number of tokens")
	goutil.AssertNow(t, l.Tokens[1].Type == token.ModAssign, "wrong operator type")
}

func TestLexerLiterals(t *testing.T) {
	l := LexString(`x = "hello we are dog"`)
	goutil.AssertNow(t, len(l.Tokens) == 3, "wrong number of tokens")
	goutil.AssertNow(t, l.Tokens[2].Type == token.String, "wrong string literal type")
	l = LexString(`x = 'a'`)
	goutil.AssertNow(t, len(l.Tokens) == 3, "wrong number of tokens")
	goutil.AssertNow(t, l.Tokens[2].Type == token.Character, "wrong character literal type")
	// test length
	l = LexString(`x = 6`)
	goutil.AssertNow(t, len(l.Tokens) == 3, "wrong number of tokens")
	goutil.AssertNow(t, l.Tokens[2].Type == token.Integer, "wrong integer literal type")
	l = LexString(`x = 5.5`)
	goutil.AssertNow(t, len(l.Tokens) == 3, "wrong number of tokens")
}

func TestLexerFileConstants(t *testing.T) {
	l := LexFile("tests/constants.grd")
	checkTokens(t, l.Tokens, []token.Type{
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
	l := LexString(`
		contract Dog {
		    class Hi {

		    }
		    interface XX {

		    }
		    event Hello(string)
		}
	`)
	checkTokens(t, l.Tokens, []token.Type{
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
	l := LexString("hello")
	goutil.AssertNow(t, len(l.Tokens) == 1, "wrong token length")
	l = LexString("hello.dog")
	goutil.AssertNow(t, len(l.Tokens) == 3, "wrong token length")
	l = LexString("hello.dog.cat")
	goutil.AssertNow(t, len(l.Tokens) == 5, "wrong token length")
}

func TestLexerError(t *testing.T) {
	l := LexString("")
	goutil.AssertNow(t, len(l.Errors) == 0, "error len should be zero")
}

func TestLexerType(t *testing.T) {
	l := LexString("type Number int")
	checkTokens(t, l.Tokens, []token.Type{token.KWType, token.Identifier, token.Identifier})
}

func TestLexerIncrement(t *testing.T) {
	l := LexString("x++")
	checkTokens(t, l.Tokens, []token.Type{token.Identifier, token.Increment})
	l = LexString("x--")
	checkTokens(t, l.Tokens, []token.Type{token.Identifier, token.Decrement})
}

func TestLexerComparators(t *testing.T) {
	l := LexString("> < >= <= ==")
	checkTokens(t, l.Tokens, []token.Type{token.Gtr, token.Lss, token.Geq, token.Leq, token.Eql})
}

func TestHasByte(t *testing.T) {
	text := "interface"
	l := new(Lexer)
	l.buffer = []byte(text)
	l.byteOffset = uint(len(text))
}

func TestLexerExpVar(t *testing.T) {
	l := LexString("x string")
	checkTokens(t, l.Tokens, []token.Type{token.Identifier, token.Identifier})
}

func TestLexerIntegers(t *testing.T) {
	l := LexString("x = 555")
	checkTokens(t, l.Tokens, []token.Type{token.Identifier, token.Assign, token.Integer})
}

func TestLexerFloats(t *testing.T) {
	l := LexString("x = 5.55")
	checkTokens(t, l.Tokens, []token.Type{token.Identifier, token.Assign, token.Float})
	l = LexString("x = .55")
	checkTokens(t, l.Tokens, []token.Type{token.Identifier, token.Assign, token.Float})
}

func TextLexerGeneric(t *testing.T) {
	l := LexString("<T|S|R>")
	checkTokens(t, l.Tokens, []token.Type{
		token.Lss,
		token.Identifier, token.Or,
		token.Identifier, token.Or,
		token.Identifier,
		token.Gtr,
	})
}

func TestConditionals(t *testing.T) {
	l := LexString("if elif else")
	checkTokens(t, l.Tokens, []token.Type{
		token.If,
		token.ElseIf,
		token.Else,
	})
}

func TestComplexFile(t *testing.T) {
	l := LexFile("tests/strings.grd")
	goutil.AssertLength(t, len(l.Errors), 0)
}

func TestInterfaceInheritance(t *testing.T) {
	l := LexString(`interface Switchable{}
		interface Deletable{}
		interface Light inherits Switchable, Deletable {}
	`)
	goutil.AssertNow(t, len(l.Errors) == 0, l.Errors.Format())
}

func TestAddAssign(t *testing.T) {
	l := LexString(`y += 5`)
	checkTokens(t, l.Tokens, []token.Type{
		token.Identifier,
		token.AddAssign,
		token.Integer,
	})
}

func TestLineComment(t *testing.T) {
	l := LexString(`// aa `)
	checkTokens(t, l.Tokens, []token.Type{token.LineComment})
}

func TestLineCommentAndMore(t *testing.T) {
	l := LexString(`// aa
		func`)
	checkTokens(t, l.Tokens, []token.Type{token.LineComment, token.Func})
}

func TestMultilineSingleLineComment(t *testing.T) {
	l := LexString(`/* aa */`)
	checkTokens(t, l.Tokens, []token.Type{token.MultilineComment})
}

func TestMultilineComment(t *testing.T) {
	l := LexString(`/* a

		a */`)
	checkTokens(t, l.Tokens, []token.Type{token.MultilineComment})
}

func TestMultilineCommentAndMore(t *testing.T) {
	l := LexString(`/* a

		a */func`)
	checkTokens(t, l.Tokens, []token.Type{token.MultilineComment, token.Func})
}

func TestAssigns(t *testing.T) {
	l := LexString(`x = 7 y = 6`)
	checkTokens(t, l.Tokens, []token.Type{
		token.Identifier, token.Assign, token.Integer,
		token.Identifier, token.Assign, token.Integer,
	})
}

func TestAnd(t *testing.T) {
	l := LexString(`!fake and deposit >= value`)
	checkTokens(t, l.Tokens, []token.Type{
		token.Not, token.Identifier,
		token.LogicalAnd,
		token.Identifier, token.Geq, token.Identifier,
	})
}
