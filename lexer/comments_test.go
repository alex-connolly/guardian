package lexer

/*
func TestComments(t *testing.T) {
	l := LexString(`

    `)
	checkTokens(t, l.Tokens, []token.Type{
		token.NewLine,
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
}*/
