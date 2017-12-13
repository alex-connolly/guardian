package token

import (
	"strings"
)

// TODO: make the token parser faster using a map or something

// Byterable ...
type Byterable interface {
	Bytes() []byte
	Offset() int
	SetOffset(int)
}

func isEnd(b Byterable) bool {
	return !hasBytes(b, 1)
}

func hasBytes(b Byterable, l int) bool {
	return b.Offset()+l < len(b.Bytes())
}

func current(b Byterable) byte {
	return b.Bytes()[b.Offset()]
}

func next(b Byterable) byte {
	b.SetOffset(b.Offset() + 1)
	return current(b)
}

// Type denotes the type of a token
type Type int

// IsBinaryOperator ...
func (t Type) IsBinaryOperator() bool {
	return t.isToken(GetBinaryOperators())
}

// IsUnaryOperator ...
func (t Type) IsUnaryOperator() bool {
	switch t {
	case Not:
		return true
	}
	return false
}

// IsModifier reports whether a token is a modifier
func (t Type) IsModifier() bool {
	return t.isToken(GetModifiers())
}

// IsDeclaration reports whether a token is a modifier
func (t Type) IsDeclaration() bool {
	return t.isToken(GetDeclarations())
}

// IsAssignment reports whether a token is an assignment operator
func (t Type) IsAssignment() bool {
	return t.isToken(GetAssignments())
}

func (t Type) isToken(list []Type) bool {
	for _, m := range list {
		if t == m {
			return true
		}
	}
	return false
}

// GetBinaryOperators ...
func GetBinaryOperators() []Type {
	return []Type{
		Add, Sub, Mul, Div, Gtr, Lss, Geq, Leq,
		As, And, Or, Eql, Xor, Is, Shl, Shr,
	}
}

// GetDeclarations ...
func GetDeclarations() []Type {
	return []Type{
		Class, Interface, Enum, KWType, Contract, Func,
	}
}

// GetModifiers ....
func GetModifiers() []Type {
	return []Type{Const, Internal, External, Public, Private,
		Protected, Static, Abstract, Storage, Test}
}

// GetLifecycles ....
func GetLifecycles() []Type {
	return []Type{Constructor, Destructor}
}

// GetAssignments ...
func GetAssignments() []Type {
	return []Type{Assign, AddAssign, SubAssign, MulAssign,
		DivAssign, ShrAssign, ShlAssign, ModAssign, AndAssign,
		OrAssign, XorAssign, Define}
}

// Type
const (
	Invalid Type = iota
	Identifier
	Integer
	Float
	String
	Character
	Comment
	Alias

	At           // at
	Assign       // =
	Comma        // ,
	OpenBrace    // {
	CloseBrace   // }
	OpenSquare   // [
	CloseSquare  // ]
	OpenBracket  // (
	CloseBracket // )
	OpenCorner   // <
	CloseCorner  // >
	Colon        // :
	And          // &
	Or           // |
	Xor          // ^
	Shl          // <<
	Shr          // >>
	Add          // +
	Sub          // -
	Mul          // *
	Div          // /
	Mod          // %

	Increment  // ++
	Decrement  // --
	AddAssign  // +=
	SubAssign  // -=
	MulAssign  // *=
	DivAssign  // /=
	ModAssign  // %=
	AndAssign  // &=
	OrAssign   // |=
	XorAssign  // ^=
	ShlAssign  // <<=
	ShrAssign  // >>=
	LogicalAnd // and
	LogicalOr  // or
	ArrowLeft  // <-
	ArrowRight // ->
	Inc        // ++
	Dec        // --
	Eql        // ==
	Lss        // <
	Gtr        // >
	Not        // !
	Neq        // !=
	Leq        // <=
	Geq        // >=
	Define     // :=
	Ellipsis   // ...
	Dot        // .
	Semicolon  // ;
	Ternary    // ?
	Break
	Continue
	Contract
	Class
	Event
	Enum
	Interface
	Abstract
	Constructor
	Destructor
	Const
	Var
	Run
	Defer
	If
	Elif
	Else
	Switch
	Case
	Exclusive
	Default
	Fallthrough
	For
	Func
	Goto
	Import
	Is
	As
	Inherits
	KWType
	TypeOf
	In
	Map
	Macro
	Package
	Return
	None
	True
	False
	External
	Internal
	Public
	Private
	Protected
	Static
	Storage
	NewLine
	LineComment
	CommentOpen
	CommentClose
	Test
	Fallback
)

// GetProtoTokens ...
func GetProtoTokens() []ProtoToken {
	return []ProtoToken{
		createDistinct("contract", Contract),
		createDistinct("class", Class),
		createDistinct("event", Event),
		createDistinct("enum", Enum),
		createDistinct("interface", Interface),
		createDistinct("inherits", Inherits),

		createDistinct("const", Const),
		createDistinct("var", Var),

		createDistinct("run", Run),
		createDistinct("defer", Defer),

		createDistinct("switch", Switch),
		createDistinct("case", Case),
		createDistinct("exclusive", Exclusive),
		createDistinct("default", Default),
		createDistinct("fallthrough", Fallthrough),
		createDistinct("break", Break),
		createDistinct("continue", Continue),

		createDistinct("constructor", Constructor),
		createDistinct("destructor", Destructor),

		createDistinct("if", If),
		createDistinct("elif", Elif),
		createDistinct("else", Else),

		createDistinct("for", For),
		createDistinct("func", Func),
		createDistinct("goto", Goto),
		createDistinct("import", Import),
		createDistinct("is", Is),
		createDistinct("as", As),
		createDistinct("typeof", TypeOf),
		createDistinct("type", KWType),

		createDistinct("external", External),
		createDistinct("internal", Internal),
		createDistinct("public", Public),
		createDistinct("private", Private),
		createDistinct("protected", Protected),
		createDistinct("abstract", Abstract),
		createDistinct("static", Static),
		createDistinct("storage", Storage),

		createDistinct("in", In),
		createDistinct("map", Map),
		createDistinct("macro", Macro),

		createDistinct("package", Package),
		createDistinct("return", Return),

		createDistinct("true", True),
		createDistinct("false", False),

		createDistinct("test", Test),
		createDistinct("fallback", Fallback),
		createDistinct("at", At),

		ProtoToken{"Float", isFloat, processFloat},
		// must check float first
		ProtoToken{"Integer", isInteger, processInteger},

		createFixed("/*", CommentOpen),
		createFixed("*/", CommentClose),
		createFixed("//", LineComment),

		createFixed("+=", AddAssign),
		createFixed("++", Increment),
		createFixed("+", Add),
		createFixed("-=", SubAssign),
		createFixed("--", Decrement),
		createFixed("-", Sub),
		createFixed("/=", DivAssign),
		createFixed("/", Div),
		createFixed("*=", MulAssign),
		createFixed("*", Mul),
		createFixed("%=", ModAssign),
		createFixed("%", Mod),
		createFixed("<<=", ShlAssign),
		createFixed("<<", Shl),
		createFixed(">>=", ShrAssign),
		createFixed(">>", Shr),

		createDistinct("and", LogicalAnd),
		createFixed("&", And),
		createDistinct("or", LogicalOr),
		createFixed("|", Or),
		createFixed("==", Eql),
		createFixed("!=", Neq),
		createFixed("!", Not),
		createFixed(">=", Geq),
		createFixed("<=", Leq),
		createFixed(":=", Define),
		createFixed("...", Ellipsis),

		createFixed("{", OpenBrace),
		createFixed("}", CloseBrace),
		createFixed("<", Lss),
		createFixed(">", Gtr),
		createFixed("[", OpenSquare),
		createFixed("]", CloseSquare),
		createFixed("(", OpenBracket),
		createFixed(")", CloseBracket),

		createFixed(":", Colon),
		createFixed("?", Ternary),
		createFixed(";", Semicolon),
		createFixed(".", Dot),
		createFixed(",", Comma),
		createFixed("=", Assign),

		ProtoToken{"New Line", isNewLine, processNewLine},
		ProtoToken{"Whitespace", isWhitespace, processIgnored},
		ProtoToken{"String", isString, processString},

		ProtoToken{"Identifier", isIdentifier, processIdentifier},
		ProtoToken{"Character", isCharacter, processCharacter},
	}
}

func createFixed(kw string, tkn Type) ProtoToken {
	return ProtoToken{"KW: " + kw, is(kw), processFixed(len(kw), tkn)}
}

func createDistinct(kw string, tkn Type) ProtoToken {
	return ProtoToken{kw, isDistinct(kw), processFixed(len(kw), tkn)}
}

type isFunc func(Byterable) bool
type processorFunc func(Byterable) Token

// ProtoToken ...
type ProtoToken struct {
	Name       string // for debugging
	Identifier isFunc
	Process    processorFunc
}

// Name returns the name of a token
func (t Token) Name() string {
	return t.Proto.Name
}

func (t *Token) Finalise(b Byterable) {
	t.Data = make([]byte, t.End-t.Start)
	copy(t.Data, b.Bytes()[t.Start:t.End])
}

// Token ...
type Token struct {
	Type  Type
	Proto ProtoToken
	Start int
	End   int
	Data  []byte
}

// String creates a new string from the Token's value
// TODO: escaped characters?
func (t Token) String() string {
	return string(t.Data)
}

// TrimmedString ...
func (t Token) TrimmedString() string {
	s := t.String()
	if strings.HasPrefix(s, "\"") {
		s = strings.TrimPrefix(s, "\"")
		s = strings.TrimSuffix(s, "\"")
	}
	return s
}
