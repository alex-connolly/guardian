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
	return b.Offset() >= len(b.Bytes())
}

func hasBytes(b Byterable, l int) bool {
	return b.Offset()+l <= len(b.Bytes())
}

func current(b Byterable) byte {
	return b.Bytes()[b.Offset()]
}

func next(b Byterable) byte {
	a := current(b)
	b.SetOffset(b.Offset() + 1)
	return a
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
		As, And, Or, Eql, Xor, Is, Shl, Shr, LogicalAnd, LogicalOr,
	}
}

// GetDeclarations ...
func GetDeclarations() []Type {
	return []Type{
		Class, Interface, Enum, KWType, Contract, Func,
	}
}

// GetLifecycles ....
func GetLifecycles() []Type {
	return []Type{Constructor, Destructor, Fallback}
}

// GetAssignments ...
func GetAssignments() []Type {
	return []Type{Assign, AddAssign, SubAssign, MulAssign,
		DivAssign, ShrAssign, ShlAssign, ModAssign, AndAssign,
		OrAssign, XorAssign, Define}
}

func Distinct(name string, typ Type) ProtoToken {
	return ProtoToken{
		Name:    name,
		Type:    typ,
		Process: processFixed(len(name), typ),
	}
}

func Fixed(name string, typ Type) ProtoToken {
	return ProtoToken{
		Name:    name,
		Type:    typ,
		Process: processFixed(len(name), typ),
	}
}

func getNextString(b Byterable, len int) string {
	return string(b.Bytes()[b.Offset() : b.Offset()+len])
}

// NextProtoToken ...
func NextProtoToken(b Byterable) *ProtoToken {
	longest := 3 // longest fixed token
	available := longest
	if available > len(b.Bytes())-b.Offset() {
		available = len(b.Bytes()) - b.Offset()
	}
	for i := available; i >= 0; i-- {
		key := getNextString(b, i)
		f, ok := fixed[key]
		if ok {
			return &f
		}
	}
	start := b.Offset()
	for hasBytes(b, 1) {
		if isWhitespace(b) || isNewLine(b) {
			break
		}
		next(b)
	}
	end := b.Offset()
	b.SetOffset(start)
	key := getNextString(b, end-start)
	r, ok := distinct[key]
	if ok {
		return &r
	}
	b.SetOffset(start)
	// special cases
	if isFloat(b) {
		return &ProtoToken{Name: "float", Type: Float, Process: processFloat}
	} else if isInteger(b) {
		return &ProtoToken{Name: "integer", Type: Integer, Process: processInteger}
	} else if isString(b) {
		return &ProtoToken{Name: "string", Type: String, Process: processString}
	} else if isCharacter(b) {
		return &ProtoToken{Name: "character", Type: Character, Process: processCharacter}
	} else if isWhitespace(b) {
		return &ProtoToken{Name: "ignored", Type: None, Process: processIgnored}
	} else if isNewLine(b) {
		return &ProtoToken{Name: "new line", Type: NewLine, Process: processNewLine}
	} else if isIdentifier(b) {
		return &ProtoToken{Name: "identifier", Type: Identifier, Process: processIdentifier}
	}
	return nil
}

var distinct = map[string]ProtoToken{
	"contract":  Distinct("contract", Contract),
	"class":     Distinct("class", Class),
	"event":     Distinct("event", Event),
	"enum":      Distinct("enum", Enum),
	"interface": Distinct("interface", Interface),
	"inherits":  Distinct("inherits", Inherits),

	"run":   Distinct("run", Run),
	"defer": Distinct("defer", Defer),

	"switch":      Distinct("switch", Switch),
	"case":        Distinct("case", Case),
	"exclusive":   Distinct("exclusive", Exclusive),
	"default":     Distinct("default", Default),
	"fallthrough": Distinct("fallthrough", Fallthrough),
	"break":       Distinct("break", Break),
	"continue":    Distinct("continue", Continue),

	"constructor": Distinct("constructor", Constructor),
	"destructor":  Distinct("destructor", Destructor),

	"if":      Distinct("if", If),
	"else if": Distinct("else if", ElseIf),
	"else":    Distinct("else", Else),

	"for":    Distinct("for", For),
	"func":   Distinct("func", Func),
	"goto":   Distinct("goto", Goto),
	"import": Distinct("import", Import),
	"is":     Distinct("is", Is),
	"as":     Distinct("as", As),
	"typeof": Distinct("typeof", TypeOf),
	"type":   Distinct("type", KWType),

	"in":    Distinct("in", In),
	"map":   Distinct("map", Map),
	"macro": Distinct("macro", Macro),

	"package": Distinct("package", Package),
	"return":  Distinct("return", Return),

	"true":  Distinct("true", True),
	"false": Distinct("false", False),

	"test":     Distinct("test", Test),
	"fallback": Distinct("fallback", Fallback),
	"version":  Distinct("version", Version),
}

var fixed = map[string]ProtoToken{
	":":  Fixed(":", Colon),
	"/*": Fixed("/*", CommentOpen),
	"*/": Fixed("*/", CommentClose),
	"//": Fixed("//", LineComment),

	"+=":  Fixed("+=", AddAssign),
	"++":  Fixed("++", Increment),
	"+":   Fixed("+", Add),
	"-=":  Fixed("-=", SubAssign),
	"--":  Fixed("--", Decrement),
	"-":   Fixed("-", Sub),
	"/=":  Fixed("/=", DivAssign),
	"/":   Fixed("/", Div),
	"**=": Fixed("**=", ExpAssign),
	"**":  Fixed("**", Exp),
	"*=":  Fixed("*=", MulAssign),
	"*":   Fixed("*", Mul),
	"%=":  Fixed("%=", ModAssign),
	"%":   Fixed("%", Mod),
	"<<=": Fixed("<<=", ShlAssign),
	"<<":  Fixed("<<", Shl),
	">>=": Fixed(">>=", ShrAssign),
	">>":  Fixed(">>", Shr),

	"&":   Fixed("", And),
	"|":   Fixed("|", Or),
	"==":  Fixed("==", Eql),
	"!=":  Fixed("!=", Neq),
	"!":   Fixed("!", Not),
	">=":  Fixed(">=", Geq),
	"<=":  Fixed("<=", Leq),
	":=":  Fixed(":=", Define),
	"...": Fixed("...", Ellipsis),

	"{": Fixed("}", OpenBrace),
	"}": Fixed("}", CloseBrace),
	"<": Fixed("<", Lss),
	">": Fixed(">", Gtr),
	"[": Fixed("[", OpenSquare),
	"]": Fixed("]", CloseSquare),
	"(": Fixed("(", OpenBracket),
	")": Fixed(")", CloseBracket),

	"?": Fixed("?", Ternary),
	";": Fixed(";", Semicolon),
	".": Fixed(".", Dot),
	",": Fixed(",", Comma),
	"=": Fixed("=", Assign),
	"@": Fixed("@", At),
}

// Type
const (
	Invalid Type = iota
	Custom
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
	Colon        // :
	And          // &
	Or           // |
	Xor          // ^
	Shl          // <<
	Shr          // >>
	Add          // +
	Sub          // -
	Mul          // *
	Exp          // **
	Div          // /
	Mod          // %

	Increment  // ++
	Decrement  // --
	AddAssign  // +=
	SubAssign  // -=
	MulAssign  // *=
	ExpAssign  // **=
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
	ElseIf
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
	Public
	Private
	Protected
	Static
	NewLine
	LineComment
	CommentOpen
	CommentClose
	Test
	Fallback
	Version
)

type processorFunc func(Byterable) Token

// ProtoToken ...
type ProtoToken struct {
	Name    string // for debugging
	Type    Type
	Process processorFunc
}

// Name returns the name of a token
func (t Token) Name() string {
	return t.Proto.Name
}

// Finalise ...
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
