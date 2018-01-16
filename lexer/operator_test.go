package lexer

import (
	"testing"

	"github.com/end-r/guardian/token"
)

func TestDistinguishKeywordsConst(t *testing.T) {
	l := LexString("constant")
	checkTokens(t, l.Tokens, []token.Type{token.Identifier})
	l = LexString("const (")
	checkTokens(t, l.Tokens, []token.Type{token.Const, token.OpenBracket})
	l = LexString("const(")
	checkTokens(t, l.Tokens, []token.Type{token.Const, token.OpenBracket})
}

func TestDistinguishKeywordsInt(t *testing.T) {
	l := LexString("int")
	checkTokens(t, l.Tokens, []token.Type{token.Identifier})
	l = LexString("int (")
	checkTokens(t, l.Tokens, []token.Type{token.Identifier, token.OpenBracket})
	l = LexString("int(")
	checkTokens(t, l.Tokens, []token.Type{token.Identifier, token.OpenBracket})
}

func TestDistinguishKeywordsInterface(t *testing.T) {
	l := LexString("interface")
	checkTokens(t, l.Tokens, []token.Type{token.Interface})
	l = LexString("interface (")
	checkTokens(t, l.Tokens, []token.Type{token.Interface, token.OpenBracket})
	l = LexString("interface(")
	checkTokens(t, l.Tokens, []token.Type{token.Interface, token.OpenBracket})
}

func TestDistinguishDots(t *testing.T) {
	l := LexString("...")
	checkTokens(t, l.Tokens, []token.Type{token.Ellipsis})
	l = LexString(".")
	checkTokens(t, l.Tokens, []token.Type{token.Dot})
	l = LexString("....")
	checkTokens(t, l.Tokens, []token.Type{token.Ellipsis, token.Dot})
}
