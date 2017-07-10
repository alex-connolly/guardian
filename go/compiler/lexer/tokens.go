package lexer

// lexer copied over in part from my efp

// TokenType denotes the type of a token
type TokenType int

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
	TknNOT // !

	TknNEQ      // !=
	TknLEQ      // <=
	TknGEQ      // >=
	TknDEFINE   // :=
	TknELLIPSIS // ...

	TknCOMMA     //
	TknDOT       // .
	TknSEMICOLON // ;
	TknCOLON     // :
	TknTERNARY   // ?

	TknBreak
	TknContinue

	TknContract
	TknClass
	TknInterface
	TknAbstract

	TknConst
	TknVar

	TknConst
	TknVar

	TknRun
	TknDefer

	TknSwitch
	TknCase
	TknExclusive
	TknDefault
	TknFallthrough

	TknExclusive
	TknExclusive
	TknDefault

	TknAbstract

	TknFor

	TknFor
	TknFunc
	TknGoto
	TknImport
	TknIs
	TknAs
	TknType
	TknTypeOf

	TknIn
	TknMap

	TknPackage
	TknReturn
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

		protoToken{"KW: If", is("if"), processFixed(TknExclusive)},
		protoToken{"KW: Else If", is("else if"), processFixed(TknExclusive)},
		protoToken{"KW: Else", is("default"), processFixed(TknDefault)},

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

func processFixed(tkn tokenType) processorFunc {
	return func(l *lexer) (t token) {
		// start and end don't matter
		t.tkntype = tkn
		l.offset++
		return t
	}
}

func isIdentifier(b []byte) bool {
	return ('A' <= b[0] && b[0] <= 'Z') || ('a' <= b[0] && b[0] <= 'z') || ('0' <= b[0] && b[0] <= '9') || (b[0] == '_')
}

func isNumber(b []byte) bool {
	return ('0' <= b[0] && b[0] <= '9')
}

func isString(b []byte) bool {
	return ((b[0] == '"') || (b[0] == '\''))
}

func isWhitespace(b []byte) bool {
	return (b[0] == ' ') || (b[0] == '\t')
}

func isNewLine(b []byte) bool {
	return (b[0] == '\n')
}

func is(a string) isFunc {
	return func(b []byte) bool {
		return string(b) == a
	}
}

type isFunc func([]byte) bool
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
