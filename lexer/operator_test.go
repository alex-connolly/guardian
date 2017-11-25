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
	tokens, _ := LexString("constant")
	checkTokens(t, tokens, []TokenType{TknIdentifier})
	tokens, _ = LexString("const (")
	checkTokens(t, tokens, []TokenType{TknConst, TknOpenBracket})
	tokens, _ = LexString("const(")
	checkTokens(t, tokens, []TokenType{TknConst, TknOpenBracket})
}

func TestDistinguishKeywordsInt(t *testing.T) {
	tokens, _ := LexString("int")
	checkTokens(t, tokens, []TokenType{TknIdentifier})
	tokens, _ = LexString("int (")
	checkTokens(t, tokens, []TokenType{TknIdentifier, TknOpenBracket})
	tokens, _ = LexString("int(")
	checkTokens(t, tokens, []TokenType{TknIdentifier, TknOpenBracket})
}

func TestDistinguishKeywordsInterface(t *testing.T) {
	tokens, _ := LexString("interface")
	checkTokens(t, tokens, []TokenType{TknInterface})
	tokens, _ = LexString("interface (")
	checkTokens(t, tokens, []TokenType{TknInterface, TknOpenBracket})
	tokens, _ = LexString("interface(")
	checkTokens(t, tokens, []TokenType{TknInterface, TknOpenBracket})
}

func TestDistinguishDots(t *testing.T) {
	tokens, _ := LexString("...")
	checkTokens(t, tokens, []TokenType{TknEllipsis})
	tokens, _ = LexString(".")
	checkTokens(t, tokens, []TokenType{TknDot})
	tokens, _ = LexString("....")
	checkTokens(t, tokens, []TokenType{TknEllipsis, TknDot})
}
