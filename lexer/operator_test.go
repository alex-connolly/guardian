package lexer

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestIsUnaryOperator(t *testing.T) {
	x := TknNot
	goutil.Assert(t, x.IsUnaryOperator(), "not should be unary")
	x = TknAdd
	goutil.Assert(t, !x.IsUnaryOperator(), "add should not be unary")
}

func TestIsBinaryOperator(t *testing.T) {
	x := TknAdd
	goutil.Assert(t, x.IsBinaryOperator(), "add should be binary")
	x = TknNot
	goutil.Assert(t, !x.IsBinaryOperator(), "not should not be binary")
}

func TestIsIdentifierByte(t *testing.T) {
	goutil.Assert(t, isIdentifierByte('A'), "upper letter not id byte")
	goutil.Assert(t, isIdentifierByte('t'), "lower letter not id byte")
	goutil.Assert(t, isIdentifierByte('2'), "number not id byte")
	goutil.Assert(t, isIdentifierByte('_'), "underscore not id byte")
	goutil.Assert(t, !isIdentifierByte(' '), "space should not be id byte")
}

func TestDistinguishKeywordsConst(t *testing.T) {
	l := LexString("constant")
	checkTokens(t, l.Tokens, []TokenType{TknIdentifier})
	l = LexString("const (")
	checkTokens(t, l.Tokens, []TokenType{TknConst, TknOpenBracket})
	l = LexString("const(")
	checkTokens(t, l.Tokens, []TokenType{TknConst, TknOpenBracket})
}

func TestDistinguishKeywordsInt(t *testing.T) {
	l := LexString("int")
	checkTokens(t, l.Tokens, []TokenType{TknIdentifier})
	l = LexString("int (")
	checkTokens(t, l.Tokens, []TokenType{TknIdentifier, TknOpenBracket})
	l = LexString("int(")
	checkTokens(t, l.Tokens, []TokenType{TknIdentifier, TknOpenBracket})
}

func TestDistinguishKeywordsInterface(t *testing.T) {
	l := LexString("interface")
	checkTokens(t, l.Tokens, []TokenType{TknInterface})
	l = LexString("interface (")
	checkTokens(t, l.Tokens, []TokenType{TknInterface, TknOpenBracket})
	l = LexString("interface(")
	checkTokens(t, l.Tokens, []TokenType{TknInterface, TknOpenBracket})
}

func TestDistinguishDots(t *testing.T) {
	l := LexString("...")
	checkTokens(t, l.Tokens, []TokenType{TknEllipsis})
	l = LexString(".")
	checkTokens(t, l.Tokens, []TokenType{TknDot})
	l = LexString("....")
	checkTokens(t, l.Tokens, []TokenType{TknEllipsis, TknDot})
}
