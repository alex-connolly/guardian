package lexer

// Lexer copied over in part from my efp

// TokenType denotes the type of a token
type TokenType int

// IsBinaryOperator ...
func (t TokenType) IsBinaryOperator() bool {
	switch t {
	case TknAdd, TknSub, TknMul, TknDiv:
		return true
	}
	return false
}

// IsUnaryOperator ...
func (t TokenType) IsUnaryOperator() bool {
	switch t {
	case TknNot:
		return true
	}
	return false
}

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
	TknMul          // *
	TknDiv          // /
	TknMod          // %

	TknIncrement // ++
	TknDecrement // --

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

	TknLogicalAnd // and
	TknLogicalOr  // or
	TknArrowLeft  // <-
	TknArrowRight // ->
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

	TknNewLine

	TknBreak
	TknContinue

	TknContract
	TknClass
	TknEvent
	TknEnum
	TknInterface
	TknAbstract

	TknConstructor
	TknDestructor

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
	TknMacro

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
		createFixed("event", TknEvent),
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
		createFixed("typeof", TknTypeOf),
		createFixed("type", TknType),

		createFixed("in", TknIn),
		createFixed("map", TknMap),
		createFixed("macro", TknMacro),

		createFixed("package", TknPackage),
		createFixed("return", TknReturn),

		createFixed("+=", TknAddAssign),
		createFixed("++", TknIncrement),
		createFixed("+", TknAdd),
		createFixed("-=", TknSubAssign),
		createFixed("--", TknDecrement),
		createFixed("-", TknSub),
		createFixed("/=", TknDivAssign),
		createFixed("/", TknDiv),
		createFixed("*=", TknMulAssign),
		createFixed("*", TknMul),
		createFixed("%=", TknModAssign),
		createFixed("%", TknMod),
		createFixed("<<=", TknShlAssign),
		createFixed("<<", TknShl),
		createFixed(">>=", TknShrAssign),
		createFixed(">>", TknShr),

		createFixed("<-", TknArrowLeft),
		createFixed(">-", TknArrowRight),
		createFixed("and", TknLogicalAnd),
		createFixed("&", TknAnd),
		createFixed("or", TknLogicalOr),
		createFixed("|", TknOr),
		createFixed("==", TknEql),
		createFixed("!=", TknNeq),
		createFixed("!", TknNot),
		createFixed(">=", TknGeq),
		createFixed("<=", TknLeq),
		createFixed(":=", TknDefine),
		createFixed("...", TknEllipsis),

		createFixed("{", TknOpenBrace),
		createFixed("}", TknCloseBrace),
		createFixed("<", TknOpenCorner),
		createFixed(">", TknCloseCorner),
		createFixed("[", TknOpenSquare),
		createFixed("]", TknCloseSquare),
		createFixed("(", TknOpenBracket),
		createFixed(")", TknCloseBracket),

		createFixed(":", TknColon),
		createFixed("?", TknTernary),
		createFixed(";", TknSemicolon),
		createFixed(".", TknDot),
		createFixed(",", TknComma),
		createFixed("=", TknAssign),

		protoToken{"New Line", isNewLine, processNewLine},
		protoToken{"Whitespace", isWhitespace, processIgnored},
		protoToken{"String", isString, processString},
		protoToken{"Number", isNumber, processNumber},
		protoToken{"Identifier", isIdentifier, processIdentifier},
		protoToken{"Character", isCharacter, processCharacter},
	}
}

func processFixed(len int, tkn TokenType) processorFunc {
	return func(l *Lexer) (t Token) {
		// start and end don't matter
		t.Type = tkn
		l.byteOffset += len
		return t
	}
}

func isIdentifier(l *Lexer) bool {
	return ('A' <= l.current() && l.current() <= 'Z') ||
		('a' <= l.current() && l.current() <= 'z') ||
		('0' <= l.current() && l.current() <= '9') ||
		(l.current() == '_')
}

func isNumber(l *Lexer) bool {
	return ('0' <= l.current() && l.current() <= '9')
}

func isString(l *Lexer) bool {
	return (l.current() == '"')
}

func isWhitespace(l *Lexer) bool {
	return (l.current() == ' ') || (l.current() == '\t')
}

func isNewLine(l *Lexer) bool {
	return (l.current() == '\n')
}

func isCharacter(l *Lexer) bool {
	return (l.current() == '\'')
}

func is(a string) isFunc {
	return func(l *Lexer) bool {
		if l.byteOffset+len(a) > len(l.buffer) {
			return false
		}
		//	fmt.Printf("cmp %s to %s\n", string(l.buffer[l.offset:l.offset+len(a)]), a)
		return string(l.buffer[l.byteOffset:l.byteOffset+len(a)]) == a
	}
}

func createFixed(kw string, tkn TokenType) protoToken {
	return protoToken{"KW: " + kw, is(kw), processFixed(len(kw), tkn)}
}

type isFunc func(*Lexer) bool
type processorFunc func(*Lexer) Token

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
