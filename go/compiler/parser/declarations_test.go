package parser

import (
	"fmt"
	"testing"

	"github.com/end-r/goutil"
)

func TestParseInterfaceDeclarationEmpty(t *testing.T) {
	p := createParser(`interface Wagable {}`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 4, "wrong token length")
	goutil.Assert(t, isInterfaceDeclaration(p), "should detect interface decl")
	parseInterfaceDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("interface")) == 1, "wrong node count")
}

func TestParseContractDeclarationEmpty(t *testing.T) {
	p := createParser(`contract Wagable {}`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 4, fmt.Sprintf("wrong token length: %d", len(p.lexer.Tokens)))
	goutil.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("contract")) == 1, "wrong node count")
}

func TestParseClassDeclarationEmpty(t *testing.T) {
	p := createParser(`class Wagable {}`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 4, "wrong token length")
	goutil.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("class")) == 1, "wrong node count")
}

func TestParseTypeDeclaration(t *testing.T) {
	p := createParser(`type Wagable int`)
	fmt.Println(p.lexer.Tokens)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 3, fmt.Sprintf("wrong token length: %d", len(p.lexer.Tokens)))
	goutil.Assert(t, isTypeDeclaration(p), "should detect type decl")
	parseTypeDeclaration(p)
}

func TestParseExplicitVarDeclaration(t *testing.T) {
	p := createParser(`a string`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 2, "wrong token length")
	goutil.Assert(t, isExplicitVarDeclaration(p), "should detect expvar decl")
	parseExplicitVarDeclaration(p)
}

func TestParseEventDeclarationEmpty(t *testing.T) {
	p := createParser(`event Notification()`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 4, "wrong token length")
	goutil.Assert(t, isEventDeclaration(p), "should detect event decl")
	parseEventDeclaration(p)
}

func TestParseEventDeclarationSingle(t *testing.T) {
	p := createParser(`event Notification(string)`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 5, "wrong token length")
	goutil.Assert(t, isEventDeclaration(p), "should detect event decl")
	parseEventDeclaration(p)
}

func TestParseEventDeclarationMultiple(t *testing.T) {
	p := createParser(`event Notification(string, string)`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 7, "wrong token length")
	goutil.Assert(t, isEventDeclaration(p), "should detect event decl")
	parseEventDeclaration(p)
}

func TestParseConstructorEmpty(t *testing.T) {
	p := createParser(`constructor() {}`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 5, "wrong token length")
	goutil.Assert(t, isConstructorDeclaration(p), "should detect constructor decl")
	parseConstructorDeclaration(p)
}

func TestParseConstructorSingle(t *testing.T) {
	p := createParser(`constructor(a int) {}`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 7, "wrong token length")
	goutil.Assert(t, isConstructorDeclaration(p), "should detect constructor decl")
	parseConstructorDeclaration(p)
}

func TestParseConstructorMultiple(t *testing.T) {
	p := createParser(`constructor(a, b string) {}`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 9, "wrong token length")
	goutil.Assert(t, isConstructorDeclaration(p), "should detect constructor decl")
	parseConstructorDeclaration(p)
}

func TestParseEnum(t *testing.T) {
	p := createParser(`enum Weekday {}`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 4, "wrong token length")
	goutil.Assert(t, isEnumDeclaration(p), "should detect enum decl")
	parseEnumDeclaration(p)
}

func TestParseEnumInheritsSingle(t *testing.T) {
	p := createParser(`enum Day inherits Weekday {}`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 6, "wrong token length")
	goutil.Assert(t, isEnumDeclaration(p), "should detect enum decl")
	parseEnumDeclaration(p)
}

func TestParseEnumInheritsMultiple(t *testing.T) {
	p := createParser(`enum Day inherits Weekday, Weekend {}`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 8, "wrong token length")
	goutil.Assert(t, isEnumDeclaration(p), "should detect enum decl")
	parseEnumDeclaration(p)
}
