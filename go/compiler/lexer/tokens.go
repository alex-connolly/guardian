
type TokenType int

// lexer copied over in part from my efp

const (
	TknValue = iota
	TknNumber
	TknString
	TknComment
	TknAlias
	TknAssign       // =
	TknComma        // ,
	TknOpenBrace    // {
	TknCloseBrace   // }
	TknOpenSquare   // [
	TknCloseSquare  // 10
	TknOpenBracket  // (
	TknCloseBracket // )
	TknOpenCorner   // <
	TknCloseCorner  // >
	TknColon        // :
)

func getProtoTokens() []protoToken {
	return []protoToken{
		protoToken{"Open Square Bracket", is('['), processOperator(TknOpenSquare)},
		protoToken{"Close Square Bracket", is(']'), processOperator(TknCloseSquare)},
		protoToken{"Open Bracket", is('('), processOperator(TknOpenBracket)},
		protoToken{"Close Bracket", is(')'), processOperator(TknCloseBracket)},
		protoToken{"Open Brace", is('{'), processOperator(TknOpenBrace)},
		protoToken{"Close Brace", is('}'), processOperator(TknCloseBrace)},
		protoToken{"Open Corner", is('<'), processOperator(TknOpenCorner)},
		protoToken{"Close Corner", is('>'), processOperator(TknCloseCorner)},
		protoToken{"Assignment Operator", is('='), processOperator(TknAssign)},
		protoToken{"Required Operator", is('!'), processOperator(TknRequired)},
		protoToken{"Comma", is(','), processOperator(TknComma)},
		protoToken{"Colon", is(':'), processOperator(TknColon)},
		protoToken{"Or", is('|'), processOperator(TknOr)},
		protoToken{"New Line", isNewLine, processNewLine},
		protoToken{"Whitespace", isWhitespace, processIgnored},
		protoToken{"String", isString, processString},
		protoToken{"Number", isNumber, processNumber},
		protoToken{"Identifier", isIdentifier, processIdentifier},
	}
}

func processOperator(tkn tokenType) processorFunc {
	return func(l *lexer) (t token) {
		t.start = l.offset
		t.end = l.offset
		t.tkntype = tkn
		l.offset++
		return t
	}
}

func isIdentifier(b byte) bool {
	return ('A' <= b && b <= 'Z') || ('a' <= b && b <= 'z') || ('0' <= b && b <= '9') || (b == '_')
}

func isNumber(b byte) bool {
	return ('0' <= b && b <= '9')
}

func isString(b byte) bool {
	return ((b == '"') || (b == '\''))
}

func isWhitespace(b byte) bool {
	return (b == ' ') || (b == '\t')
}

func isNewLine(b byte) bool {
	return (b == '\n')
}

func is(a byte) isFunc {
	return func(b byte) bool {
		return b == a
	}
}

type isFunc func(byte) bool
type processorFunc func(*lexer) token

type protoToken struct {
	name       string // for debugging
	identifier isFunc
	process    processorFunc
}

type token struct {
	tkntype tokenType
	start   int
	end     int
}
