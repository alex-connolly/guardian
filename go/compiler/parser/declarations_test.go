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

func TestParseInterfaceDeclarationSingleInheritance(t *testing.T) {
	p := createParser(`interface Wagable inherits Visible {}`)
	goutil.Assert(t, isInterfaceDeclaration(p), "should detect interface decl")
	parseInterfaceDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("interface")) == 1, "wrong node count")
}

func TestParseInterfaceDeclarationMultipleInheritance(t *testing.T) {
	p := createParser(`interface Wagable inherits Visible, Movable {}`)
	goutil.Assert(t, isInterfaceDeclaration(p), "should detect interface decl")
	parseInterfaceDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("interface")) == 1, "wrong node count")
}

func TestParseInterfaceDeclarationAbstract(t *testing.T) {
	p := createParser(`abstract interface Wagable {}`)
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

func TestParseClassDeclarationSingleInterface(t *testing.T) {
	p := createParser(`class Wagable is Visible {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect interface decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("class")) == 1, "wrong node count")
}

func TestParseClassDeclarationMultipleInterfaces(t *testing.T) {
	p := createParser(`class Wagable is Visible, Movable {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("class")) == 1, "wrong node count")
}

func TestParseClassDeclarationSingleInterfaceSingleInheritance(t *testing.T) {
	p := createParser(`class Wagable is Visible inherits Object {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect interface decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("class")) == 1, "wrong node count")
}

func TestParseClassDeclarationMultipleInterfaceMultipleInheritance(t *testing.T) {
	p := createParser(`class Wagable inherits A,B is Visible, Movable  {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("class")) == 1, "wrong node count")
}

func TestParseClassDeclarationSingleInheritance(t *testing.T) {
	p := createParser(`class Wagable inherits Visible {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect interface decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("class")) == 1, "wrong node count")
}

func TestParseClassDeclarationMultipleInheritance(t *testing.T) {
	p := createParser(`class Wagable inherits Visible, Movable {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("class")) == 1, "wrong node count")
}

func TestParseClassDeclarationSingleInheritanceMultipleInterface(t *testing.T) {
	p := createParser(`class Wagable inherits Visible {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect interface decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("class")) == 1, "wrong node count")
}

func TestParseClassDeclarationSingleInheritanceSingleInterface(t *testing.T) {
	p := createParser(`class Wagable inherits Object is Visible {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect interface decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("class")) == 1, "wrong node count")
}

func TestParseClassDeclarationMultipleInheritanceSingleInterface(t *testing.T) {
	p := createParser(`class Wagable is Visible, Movable inherits A {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("class")) == 1, "wrong node count")
}

func TestParseClassDeclarationMultipleInheritanceMultipleInterface(t *testing.T) {
	p := createParser(`class Wagable is Visible, Movable inherits A,B {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("class")) == 1, "wrong node count")
}

func TestParseClassDeclarationAbstract(t *testing.T) {
	p := createParser(`abstract class Wagable {}`)
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

func TestParseFuncNoParameters(t *testing.T) {
	p := createParser(`func foo(){}`)
	goutil.Assert(t, isFuncDeclaration(p), "should detect func decl")
	parseFuncDeclaration(p)
}

func TestParseFuncOneParameter(t *testing.T) {
	p := createParser(`func foo(a int){}`)
	goutil.Assert(t, isFuncDeclaration(p), "should detect func decl")
	parseFuncDeclaration(p)
}

func TestParseFuncParameters(t *testing.T) {
	p := createParser(`func foo(a int, b string){}`)
	goutil.Assert(t, isFuncDeclaration(p), "should detect func decl")
	parseFuncDeclaration(p)
}

func TestParseFuncMultiplePerType(t *testing.T) {
	p := createParser(`func foo(a, b int){}`)
	goutil.Assert(t, isFuncDeclaration(p), "should detect func decl")
	parseFuncDeclaration(p)
}

func TestParseFuncMultiplePerTypeExtra(t *testing.T) {
	p := createParser(`func foo(a, b int, c string){}`)
	goutil.Assert(t, isFuncDeclaration(p), "should detect func decl")
	parseFuncDeclaration(p)
}

func TestParseConstructorNoParameters(t *testing.T) {
	p := createParser(`constructor(){}`)
	goutil.Assert(t, isConstructorDeclaration(p), "should detect Constructor decl")
	parseConstructorDeclaration(p)
}

func TestParseConstructorOneParameter(t *testing.T) {
	p := createParser(`constructor(a int){}`)
	goutil.Assert(t, isConstructorDeclaration(p), "should detect Constructor decl")
	parseConstructorDeclaration(p)
}

func TestParseConstructorParameters(t *testing.T) {
	p := createParser(`constructor(a int, b string){}`)
	goutil.Assert(t, isConstructorDeclaration(p), "should detect Constructor decl")
	parseConstructorDeclaration(p)
}

func TestParseConstructorMultiplePerType(t *testing.T) {
	p := createParser(`constructor(a, b int){}`)
	goutil.Assert(t, isConstructorDeclaration(p), "should detect Constructor decl")
	parseConstructorDeclaration(p)
}

func TestParseConstructorMultiplePerTypeExtra(t *testing.T) {
	p := createParser(`constructor(a, b int, c string){}`)
	goutil.Assert(t, isConstructorDeclaration(p), "should detect Constructor decl")
	parseConstructorDeclaration(p)
}
