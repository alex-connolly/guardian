package lexer

// lexer copied over in part from my efp

// TokenType denotes the type of a token
type TokenType int

const (
	TknIdentifier TokenType = iota
	TknNumber
	TknString
	TknCharacter
	TknComment
	TknAlias
	TknAssign       // =
	TknComma        // ,
	TknOpenBrace    // {
	TknCloseBrace   // }
	TknOpenSquare   // [
	TknCloseSquare  // ]
	TknOpenBracket  // (
	TknCloseBracket // )
	TknOpenCorner   // <
	TknCloseCorner  // >
	TknColon        // :
	TknAnd          // &
	TknOr           // |
	TknXor          // ^
	TknShl          // <<
	TknShr          // >>
	TknAdd          // +
	TknSub          // -
	TknStar         // *
	TknDiv          // /
	TknMod          // %
	TknAddress      // @

	TknAddAssign // +=
	TknSubAssign // -=
	TknMulAssign // *=
	TknDivAssign // /=
	TknModAssign // %=

	TknAndAssign // &=
	TknOrAssign  // |=
	TknXorAssign // ^=
	TknShlAssign // <<=
	TknShrAssign // >>=

	TknLogicalAnd // &&
	TknLogicalOr  // ||
	TknArrow      // <-
	TknInc        // ++
	TknDec        // --

	TknEql // ==
	TknLss // <
	TknGtr // >
	TknNot // !

	TknNeq      // !=
	TknLeq      // <=
	TknGeq      // >=
	TknDefine   // :=
	TknEllipsis // ...

	TknDot       // .
	TknSemicolon // ;
	TknTernary   // ?

	TknBreak
	TknContinue

	TknContract
	TknClass
	TknInterface
	TknAbstract

	TknConst
	TknVar

	TknRun
	TknDefer

	TknIf
	TknElseIf
	TknElse

	TknSwitch
	TknCase
	TknExclusive
	TknDefault
	TknFallthrough

	TknFor
	TknFunc
	TknGoto
	TknImport
	TknIs
	TknAs
	TknInherits
	TknType
	TknTypeOf

	TknIn
	TknMap

	TknPackage
	TknReturn

	TknNone
)

// TODO: protoToken{"Operator", isOperator, processOperator}
// TODO: benchmark recognising that a token is an op and then sorting from there

func getProtoTokens() []protoToken {
	return []protoToken{
		createFixed("contract", TknContract),
		createFixed("class", TknClass),
		createFixed("interface", TknInterface),
		createFixed("abstract", TknAbstract),
		createFixed("inherits", TknInherits),

		createFixed("const", TknConst),
		createFixed("var", TknVar),

		createFixed("run", TknRun),
		createFixed("defer", TknDefer),

		createFixed("switch", TknSwitch),
		createFixed("case", TknCase),
		createFixed("exclusive", TknExclusive),
		createFixed("default", TknDefault),
		createFixed("fallthrough", TknFallthrough),
		createFixed("break", TknBreak),
		createFixed("continue", TknContinue),

		createFixed("if", TknIf),
		createFixed("else if", TknElseIf),
		createFixed("default", TknElse),

		createFixed("for", TknFor),
		createFixed("func", TknFunc),
		createFixed("goto", TknGoto),
		createFixed("import", TknImport),
		createFixed("is", TknIs),
		createFixed("as", TknAs),
		createFixed("type", TknType),
		createFixed("typeof", TknTypeOf),

		createFixed("in", TknIn),
		createFixed("map", TknMap),

		createFixed("package", TknPackage),
		createFixed("return", TknReturn),

		createFixed("+", TknAdd),
		createFixed("-", TknSub),
		createFixed("/", TknDiv),
		createFixed("*", TknStar),
		createFixed("%", TknMod),
		createFixed("@", TknAddress),
		createFixed("{", TknOpenBrace),
		createFixed("}", TknCloseBrace),
		createFixed("<", TknOpenCorner),
		createFixed(">", TknCloseCorner),
		createFixed("[", TknOpenSquare),
		createFixed("]", TknCloseSquare),
		createFixed("(", TknOpenBracket),
		createFixed(")", TknCloseBracket),

		protoToken{"New Line", isNewLine, processNewLine},
		protoToken{"Whitespace", isWhitespace, processIgnored},
		protoToken{"String", isString, processString},
		protoToken{"Number", isNumber, processNumber},
		protoToken{"Identifier", isIdentifier, processIdentifier},
		protoToken{"Character", isCharacter, processCharacter},
	}
}

func processFixed(len int, tkn TokenType) processorFunc {
	return func(l *lexer) (t Token) {
		// start and end don't matter
		t.Type = tkn
		l.offset += len
		return t
	}
}

func isOperator(l *lexer) bool {
	switch l.buffer[l.offset] {
	case '+', '-', '*', '/':
		return true
	}
	return false
}

func isIdentifier(l *lexer) bool {
	return ('A' <= l.current() && l.current() <= 'Z') ||
		('a' <= l.current() && l.current() <= 'z') ||
		('0' <= l.current() && l.current() <= '9') ||
		(l.current() == '_')
}

func isNumber(l *lexer) bool {
	return ('0' <= l.current() && l.current() <= '9')
}

func isString(l *lexer) bool {
	return ((l.current() == '"') || (l.current() == '\''))
}

func isWhitespace(l *lexer) bool {
	return (l.current() == ' ') || (l.current() == '\t')
}

func isNewLine(l *lexer) bool {
	return (l.current() == '\n')
}

func isCharacter(l *lexer) bool {
	return (l.current() == '\'')
}

func is(a string) isFunc {
	return func(l *lexer) bool {
		if l.offset+len(a) > len(l.buffer) {
			return false
		}
		//	fmt.Printf("cmp %s to %s\n", string(l.buffer[l.offset:l.offset+len(a)]), a)
		return string(l.buffer[l.offset:l.offset+len(a)]) == a
	}
}

func createFixed(kw string, tkn TokenType) protoToken {
	return protoToken{"KW: " + kw, is(kw), processFixed(len(kw), tkn)}
}

type isFunc func(*lexer) bool
type processorFunc func(*lexer) Token

type protoToken struct {
	name       string // for debugging
	identifier isFunc
	process    processorFunc
}

// Token ...
type Token struct {
	Type  TokenType
	start int
	end   int
}
