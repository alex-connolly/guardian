package lexer

import (
	"testing"

	"github.com/end-r/guardian/token"
)

func TestComments(t *testing.T) {
	l := LexString(`
        // this

        /*
        this
        */

        /* comment */

        /*
        // comment
        */
    `)
	checkTokens(t, l.Tokens, []token.Type{
		token.NewLine,
		token.LineComment, token.Identifier, token.NewLine,
		token.NewLine,
		token.CommentOpen, token.NewLine,
		token.Identifier, token.NewLine,
		token.CommentClose, token.NewLine,
		token.NewLine,
		token.CommentOpen, token.Identifier, token.CommentClose, token.NewLine,
		token.NewLine,
		token.CommentOpen, token.NewLine,
		token.LineComment, token.Identifier, token.NewLine,
		token.CommentClose, token.NewLine,
	})
}
