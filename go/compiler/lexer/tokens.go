package lexer

import "fmt"

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

func getProtoTokens() []protoToken {
	return []protoToken{
		protoToken{"Open Square Bracket", is("["), processFixed(TknOpenSquare)},
		protoToken{"Close Square Bracket", is("]"), processFixed(TknCloseSquare)},
		protoToken{"Open Bracket", is("("), processFixed(TknOpenBracket)},
		protoToken{"Close Bracket", is(")"), processFixed(TknCloseBracket)},
		protoToken{"Open Brace", is("{"), processFixed(TknOpenBrace)},
		protoToken{"Close Brace", is("}"), processFixed(TknCloseBrace)},
		protoToken{"Open Corner", is("<"), processFixed(TknOpenCorner)},
		protoToken{"Close Corner", is(">"), processFixed(TknCloseCorner)},
		protoToken{"Assignment", is("="), processFixed(TknAssign)},
		protoToken{"Add", is("+"), processFixed(TknAdd)},
		protoToken{"Sub", is("-"), processFixed(TknSub)},
		protoToken{"Star", is("*"), processFixed(TknStar)},
		protoToken{"Div", is("/"), processFixed(TknDiv)},
		protoToken{"Mod", is("%"), processFixed(TknMod)},
		protoToken{"Not", is("!"), processFixed(TknNot)},
		protoToken{"Comma", is(","), processFixed(TknComma)},
		protoToken{"Colon", is(":"), processFixed(TknColon)},
		protoToken{"Or", is("|"), processFixed(TknOr)},
		protoToken{"Address", is("@"), processFixed(TknAddress)},

		protoToken{"KW: Contract", is("contract"), processFixed(TknContract)},
		protoToken{"KW: Class", is("class"), processFixed(TknClass)},
		protoToken{"KW: Interface", is("interface"), processFixed(TknInterface)},
		protoToken{"KW: Abstract", is("abstract"), processFixed(TknAbstract)},
		protoToken{"KW: Inherits", is("inherits"), processFixed(TknInherits)},

		protoToken{"KW: Const", is("const"), processFixed(TknConst)},
		protoToken{"KW: Var", is("var"), processFixed(TknVar)},

		protoToken{"KW: Run", is("run"), processFixed(TknRun)},
		protoToken{"KW: Defer", is("defer"), processFixed(TknDefer)},

		protoToken{"KW: Switch", is("switch"), processFixed(TknSwitch)},
		protoToken{"KW: Case", is("case"), processFixed(TknCase)},
		protoToken{"KW: Exclusive", is("exclusive"), processFixed(TknExclusive)},
		protoToken{"KW: Default", is("default"), processFixed(TknDefault)},
		protoToken{"KW: Fallthrough", is("fallthrough"), processFixed(TknFallthrough)},
		protoToken{"KW: Break", is("break"), processFixed(TknBreak)},
		protoToken{"KW: Continue", is("continue"), processFixed(TknContinue)},

		protoToken{"KW: If", is("if"), processFixed(TknIf)},
		protoToken{"KW: Else If", is("else if"), processFixed(TknElseIf)},
		protoToken{"KW: Else", is("default"), processFixed(TknElse)},

		protoToken{"KW: For", is("for"), processFixed(TknFor)},
		protoToken{"KW: Func", is("func"), processFixed(TknFunc)},
		protoToken{"KW: Goto", is("goto"), processFixed(TknGoto)},
		protoToken{"KW: Import", is("import"), processFixed(TknImport)},
		protoToken{"KW: Is", is("is"), processFixed(TknIs)},
		protoToken{"KW: As", is("as"), processFixed(TknAs)},
		protoToken{"KW: Type", is("type"), processFixed(TknType)},
		protoToken{"KW: Typeof", is("typeof"), processFixed(TknTypeOf)},

		protoToken{"KW: In", is("in"), processFixed(TknIn)},
		protoToken{"KW: Map", is("map"), processFixed(TknMap)},

		protoToken{"KW: Package", is("package"), processFixed(TknPackage)},
		protoToken{"KW: Return", is("return"), processFixed(TknReturn)},

		protoToken{"New Line", isNewLine, processNewLine},
		protoToken{"Whitespace", isWhitespace, processIgnored},
		protoToken{"String", isString, processString},
		protoToken{"Number", isNumber, processNumber},
		protoToken{"Identifier", isIdentifier, processIdentifier},
		protoToken{"Character", isCharacter, processCharacter},
	}
}

func processFixed(tkn TokenType) processorFunc {
	return func(l *lexer) (t Token) {
		// start and end don't matter
		t.Type = tkn
		l.offset++
		return t
	}
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
		fmt.Printf("cmp %s to %s\n", string(l.buffer[l.offset:l.offset+len(a)]), a)
		return string(l.buffer[l.offset:l.offset+len(a)]) == a
	}
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
