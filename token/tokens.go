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

func distinctToken(name string, typ Type) ProtoToken {
	return ProtoToken{
		Name:    name,
		Type:    typ,
		Process: processFixed(len(name), typ),
	}
}

func fixedToken(name string, typ Type) ProtoToken {
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

	if isFloat(b) {
		return &ProtoToken{Name: "float", Type: Float, Process: processFloat}
	} else if isInteger(b) {
		return &ProtoToken{Name: "integer", Type: Integer, Process: processInteger}
	}

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
		if !isIdentifierByte(current(b)) {
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
	if isString(b) {
		return &ProtoToken{Name: "string", Type: String, Process: processString}
	} else if isCharacter(b) {
		return &ProtoToken{Name: "character", Type: Character, Process: processCharacter}
	} else if isWhitespace(b) {
		return &ProtoToken{Name: "ignored", Type: None, Process: processIgnored}
	} else if isNewLine(b) {
		return &ProtoToken{Name: "new line", Type: NewLine, Process: processNewLine}
	}

	if isIdentifier(b) {
		return &ProtoToken{Name: "identifier", Type: Identifier, Process: processIdentifier}
	}
	return nil
}

var distinct = map[string]ProtoToken{
	"contract":  distinctToken("contract", Contract),
	"class":     distinctToken("class", Class),
	"event":     distinctToken("event", Event),
	"enum":      distinctToken("enum", Enum),
	"interface": distinctToken("interface", Interface),
	"inherits":  distinctToken("inherits", Inherits),

	"run":   distinctToken("run", Run),
	"defer": distinctToken("defer", Defer),

	"switch":      distinctToken("switch", Switch),
	"case":        distinctToken("case", Case),
	"exclusive":   distinctToken("exclusive", Exclusive),
	"default":     distinctToken("default", Default),
	"fallthrough": distinctToken("fallthrough", Fallthrough),
	"break":       distinctToken("break", Break),
	"continue":    distinctToken("continue", Continue),

	"constructor": distinctToken("constructor", Constructor),
	"destructor":  distinctToken("destructor", Destructor),

	"if":   distinctToken("if", If),
	"elif": distinctToken("elif", ElseIf),
	"else": distinctToken("else", Else),

	"for":    distinctToken("for", For),
	"func":   distinctToken("func", Func),
	"goto":   distinctToken("goto", Goto),
	"import": distinctToken("import", Import),
	"is":     distinctToken("is", Is),
	"as":     distinctToken("as", As),
	"typeof": distinctToken("typeof", TypeOf),
	"type":   distinctToken("type", KWType),

	"in":    distinctToken("in", In),
	"map":   distinctToken("map", Map),
	"macro": distinctToken("macro", Macro),

	"package": distinctToken("package", Package),
	"return":  distinctToken("return", Return),

	"true":  distinctToken("true", True),
	"false": distinctToken("false", False),

	"var":   distinctToken("var", Var),
	"const": distinctToken("const", Const),

	"test":     distinctToken("test", Test),
	"fallback": distinctToken("fallback", Fallback),
	"version":  distinctToken("version", Version),
}

var fixed = map[string]ProtoToken{
	":":  fixedToken(":", Colon),
	"/*": fixedToken("/*", CommentOpen),
	"*/": fixedToken("*/", CommentClose),
	"//": fixedToken("//", LineComment),

	"+=":  fixedToken("+=", AddAssign),
	"++":  fixedToken("++", Increment),
	"+":   fixedToken("+", Add),
	"-=":  fixedToken("-=", SubAssign),
	"--":  fixedToken("--", Decrement),
	"-":   fixedToken("-", Sub),
	"/=":  fixedToken("/=", DivAssign),
	"/":   fixedToken("/", Div),
	"**=": fixedToken("**=", ExpAssign),
	"**":  fixedToken("**", Exp),
	"*=":  fixedToken("*=", MulAssign),
	"*":   fixedToken("*", Mul),
	"%=":  fixedToken("%=", ModAssign),
	"%":   fixedToken("%", Mod),
	"<<=": fixedToken("<<=", ShlAssign),
	"<<":  fixedToken("<<", Shl),
	">>=": fixedToken(">>=", ShrAssign),
	">>":  fixedToken(">>", Shr),

	"&":   fixedToken("", And),
	"&=":  fixedToken("&=", AndAssign),
	"|":   fixedToken("|", Or),
	"|=":  fixedToken("|=", OrAssign),
	"^":   fixedToken("^", Xor),
	"^=":  fixedToken("^=", XorAssign),
	"==":  fixedToken("==", Eql),
	"!=":  fixedToken("!=", Neq),
	"!":   fixedToken("!", Not),
	">=":  fixedToken(">=", Geq),
	"<=":  fixedToken("<=", Leq),
	":=":  fixedToken(":=", Define),
	"...": fixedToken("...", Ellipsis),

	"{": fixedToken("}", OpenBrace),
	"}": fixedToken("}", CloseBrace),
	"<": fixedToken("<", Lss),
	">": fixedToken(">", Gtr),
	"[": fixedToken("[", OpenSquare),
	"]": fixedToken("]", CloseSquare),
	"(": fixedToken("(", OpenBracket),
	")": fixedToken(")", CloseBracket),

	"?": fixedToken("?", Ternary),
	";": fixedToken(";", Semicolon),
	".": fixedToken(".", Dot),
	",": fixedToken(",", Comma),
	"=": fixedToken("=", Assign),
	"@": fixedToken("@", At),
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
	Increment    // ++
	Decrement    // --
	AddAssign    // +=
	SubAssign    // -=
	MulAssign    // *=
	ExpAssign    // **=
	DivAssign    // /=
	ModAssign    // %=
	AndAssign    // &=
	OrAssign     // |=
	XorAssign    // ^=
	ShlAssign    // <<=
	ShrAssign    // >>=
	LogicalAnd   // and
	LogicalOr    // or
	ArrowLeft    // <-
	ArrowRight   // ->
	Inc          // ++
	Dec          // --
	Eql          // ==
	Lss          // <
	Gtr          // >
	Not          // !
	Neq          // !=
	Leq          // <=
	Geq          // >=
	Define       // :=
	Ellipsis     // ...
	Dot          // .
	Semicolon    // ;
	Ternary      // ?
	Break
	Continue
	Contract
	Class
	Event
	Enum
	Interface
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
	Proto *ProtoToken
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
