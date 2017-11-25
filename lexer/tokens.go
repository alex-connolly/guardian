package lexer

import "strings"

// Lexer copied over in part from my efp

// TokenType denotes the type of a token
type TokenType int

// IsBinaryOperator ...
func (t TokenType) IsBinaryOperator() bool {
	return t.isToken(GetBinaryOperators())
}

// IsUnaryOperator ...
func (t TokenType) IsUnaryOperator() bool {
	switch t {
	case TknNot:
		return true
	}
	return false
}

// IsModifier reports whether a token is a modifier
func (t TokenType) IsModifier() bool {
	return t.isToken(GetModifiers())
}

// IsAssignment reports whether a token is an assignment operator
func (t TokenType) IsAssignment() bool {
	return t.isToken(GetAssignments())
}

func (t TokenType) isToken(list []TokenType) bool {
	for _, m := range list {
		if t == m {
			return true
		}
	}
	return false
}

// GetBinaryOperators ...
func GetBinaryOperators() []TokenType {
	return []TokenType{
		TknAdd, TknSub, TknMul, TknDiv, TknGtr, TknLss, TknGeq, TknLeq,
		TknAs, TknAnd, TknOr, TknEql, TknXor, TknIs, TknShl, TknShr,
	}
}

// GetModifiers ....
func GetModifiers() []TokenType {
	return []TokenType{TknConst, TknInternal, TknExternal, TknPublic, TknPrivate,
		TknProtected, TknStatic, TknAbstract, TknStorage, TknTest}
}

// GetLifecycles ....
func GetLifecycles() []TokenType {
	return []TokenType{TknConstructor, TknDestructor}
}

// GetAssignments ...
func GetAssignments() []TokenType {
	return []TokenType{TknAssign, TknAddAssign, TknSubAssign, TknMulAssign,
		TknDivAssign, TknShrAssign, TknShlAssign, TknModAssign, TknAndAssign,
		TknOrAssign, TknXorAssign, TknDefine}
}

// TokenType
const (
	TknInvalid TokenType = iota
	TknIdentifier
	TknInteger
	TknFloat
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
	TknIncrement    // ++
	TknDecrement    // --
	TknAddAssign    // +=
	TknSubAssign    // -=
	TknMulAssign    // *=
	TknDivAssign    // /=
	TknModAssign    // %=
	TknAndAssign    // &=
	TknOrAssign     // |=
	TknXorAssign    // ^=
	TknShlAssign    // <<=
	TknShrAssign    // >>=
	TknLogicalAnd   // and
	TknLogicalOr    // or
	TknArrowLeft    // <-
	TknArrowRight   // ->
	TknInc          // ++
	TknDec          // --
	TknEql          // ==
	TknLss          // <
	TknGtr          // >
	TknNot          // !
	TknNeq          // !=
	TknLeq          // <=
	TknGeq          // >=
	TknDefine       // :=
	TknEllipsis     // ...
	TknDot          // .
	TknSemicolon    // ;
	TknTernary      // ?
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
	TknElif
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
	TknTrue
	TknFalse
	TknExternal
	TknInternal
	TknPublic
	TknPrivate
	TknProtected
	TknStatic
	TknStorage
	TknNewLine
	TknLineComment
	TknCommentOpen
	TknCommentClose
	TknTest
)

// TODO: protoToken{"Operator", isOperator, processOperator}
// TODO: benchmark recognising that a token is an op and then sorting from there

func getProtoTokens() []protoToken {
	return []protoToken{
		createDistinct("contract", TknContract),
		createDistinct("class", TknClass),
		createDistinct("event", TknEvent),
		createDistinct("enum", TknEnum),
		createDistinct("interface", TknInterface),
		createDistinct("inherits", TknInherits),

		createDistinct("const", TknConst),
		createDistinct("var", TknVar),

		createDistinct("run", TknRun),
		createDistinct("defer", TknDefer),

		createDistinct("switch", TknSwitch),
		createDistinct("case", TknCase),
		createDistinct("exclusive", TknExclusive),
		createDistinct("default", TknDefault),
		createDistinct("fallthrough", TknFallthrough),
		createDistinct("break", TknBreak),
		createDistinct("continue", TknContinue),

		createDistinct("constructor", TknConstructor),
		createDistinct("destructor", TknDestructor),

		createDistinct("if", TknIf),
		createDistinct("elif", TknElif),
		createDistinct("else", TknElse),

		createDistinct("for", TknFor),
		createDistinct("func", TknFunc),
		createDistinct("goto", TknGoto),
		createDistinct("import", TknImport),
		createDistinct("is", TknIs),
		createDistinct("as", TknAs),
		createDistinct("typeof", TknTypeOf),
		createDistinct("type", TknType),

		createDistinct("external", TknExternal),
		createDistinct("internal", TknInternal),
		createDistinct("public", TknPublic),
		createDistinct("private", TknPrivate),
		createDistinct("protected", TknProtected),
		createDistinct("abstract", TknAbstract),
		createDistinct("static", TknStatic),
		createDistinct("storage", TknStorage),

		createDistinct("in", TknIn),
		createDistinct("map", TknMap),
		createDistinct("macro", TknMacro),

		createDistinct("package", TknPackage),
		createDistinct("return", TknReturn),

		createDistinct("true", TknTrue),
		createDistinct("false", TknFalse),

		createDistinct("test", TknTest),

		protoToken{"Float", isFloat, processFloat},
		// must check float first
		protoToken{"Integer", isInteger, processInteger},

		createFixed("/*", TknCommentOpen),
		createFixed("*/", TknCommentClose),
		createFixed("//", TknLineComment),

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

		createDistinct("and", TknLogicalAnd),
		createFixed("&", TknAnd),
		createDistinct("or", TknLogicalOr),
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
		createFixed("<", TknLss),
		createFixed(">", TknGtr),
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

		protoToken{"Identifier", isIdentifier, processIdentifier},
		protoToken{"Character", isCharacter, processCharacter},
	}
}

func isIdentifierByte(b byte) bool {
	return ('A' <= b && b <= 'Z') ||
		('a' <= b && b <= 'z') ||
		('0' <= b && b <= '9') ||
		(b == '_')
}

func isIdentifier(l *Lexer) bool {
	return isIdentifierByte(l.current())
}

func isInteger(l *Lexer) bool {
	return ('0' <= l.current() && l.current() <= '9')
}

func isFloat(l *Lexer) bool {
	saved := l.byteOffset
	for l.hasBytes(1) && '0' <= l.current() && l.current() <= '9' {
		l.byteOffset++
	}
	if !l.hasBytes(1) || l.current() != '.' {
		l.byteOffset = saved
		return false
	}
	l.byteOffset++
	if !l.hasBytes(1) || !('0' <= l.current() && l.current() <= '9') {
		l.byteOffset = saved
		return false
	}
	l.byteOffset = saved
	return true
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

func isDistinct(a string) isFunc {
	return func(l *Lexer) bool {
		if !l.hasBytes(len(a)) {
			return false
		}
		end := l.byteOffset + len(a)
		if string(l.buffer[l.byteOffset:end]) != a {
			return false
		}
		if !l.hasBytes(len(a) + 1) {
			return true
		}
		return !isIdentifierByte(l.buffer[end])
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

func createFixed(kw string, tkn TokenType) protoToken {
	return protoToken{"KW: " + kw, is(kw), processFixed(len(kw), tkn)}
}

func createDistinct(kw string, tkn TokenType) protoToken {
	return protoToken{kw, isDistinct(kw), processFixed(len(kw), tkn)}
}

type isFunc func(*Lexer) bool
type processorFunc func(*Lexer) Token

type protoToken struct {
	name       string // for debugging
	identifier isFunc
	process    processorFunc
}

// Name returns the name of a token
func (t Token) Name() string {
	return t.proto.name
}

// Token ...
type Token struct {
	Type  TokenType
	proto protoToken
	start int
	end   int
	data  []byte
}

// TokenString creates a new string from the Token's value
// TODO: escaped characters
func (t Token) TokenString() string {
	return string(t.data)
}

// TokenStringAndTrim ...
func (t Token) TrimTokenString() string {
	s := t.TokenString()
	if strings.HasPrefix(s, "\"") {
		s = strings.TrimPrefix(s, "\"")
		s = strings.TrimSuffix(s, "\"")
	}
	return s
}
