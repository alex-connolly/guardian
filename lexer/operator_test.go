package lexer

import (
	"testing"

	"github.com/end-r/guardian/token"
)

func TestDistinguishKeywordsConst(t *testing.T) {
	tokens, _ := LexString("constant")
	checkTokens(t, tokens, []token.Type{token.Identifier})
	tokens, _ = LexString("const (")
	checkTokens(t, tokens, []token.Type{token.Const, token.OpenBracket})
	tokens, _ = LexString("const(")
	checkTokens(t, tokens, []token.Type{token.Const, token.OpenBracket})
}

func TestDistinguishKeywordsInt(t *testing.T) {
	tokens, _ := LexString("int")
	checkTokens(t, tokens, []token.Type{token.Identifier})
	tokens, _ = LexString("int (")
	checkTokens(t, tokens, []token.Type{token.Identifier, token.OpenBracket})
	tokens, _ = LexString("int(")
	checkTokens(t, tokens, []token.Type{token.Identifier, token.OpenBracket})
}

func TestDistinguishKeywordsInterface(t *testing.T) {
	tokens, _ := LexString("interface")
	checkTokens(t, tokens, []token.Type{token.Interface})
	tokens, _ = LexString("interface (")
	checkTokens(t, tokens, []token.Type{token.Interface, token.OpenBracket})
	tokens, _ = LexString("interface(")
	checkTokens(t, tokens, []token.Type{token.Interface, token.OpenBracket})
}

func TestDistinguishDots(t *testing.T) {
	tokens, _ := LexString("...")
	checkTokens(t, tokens, []token.Type{token.Ellipsis})
	tokens, _ = LexString(".")
	checkTokens(t, tokens, []token.Type{token.Dot})
	tokens, _ = LexString("....")
	checkTokens(t, tokens, []token.Type{token.Ellipsis, token.Dot})
}
