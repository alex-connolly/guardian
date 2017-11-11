package lexer

import "testing"

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
	checkTokens(t, l.Tokens, []TokenType{
		TknNewLine,
		TknLineComment, TknIdentifier, TknNewLine,
		TknNewLine,
		TknCommentOpen, TknNewLine,
		TknIdentifier, TknNewLine,
		TknCommentClose, TknNewLine,
		TknNewLine,
		TknCommentOpen, TknIdentifier, TknCommentClose, TknNewLine,
		TknNewLine,
		TknCommentOpen, TknNewLine,
		TknLineComment, TknIdentifier, TknNewLine,
		TknCommentClose, TknNewLine,
	})
}
