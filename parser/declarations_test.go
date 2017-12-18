package parser

import (
	"fmt"
	"testing"

	"github.com/end-r/guardian/ast"

	"github.com/end-r/goutil"
)

func TestParseInterfaceDeclarationEmpty(t *testing.T) {
	p := createParser(`interface Wagable {}`)
	goutil.AssertNow(t, len(p.tokens) == 4, "wrong token length")
	goutil.Assert(t, isInterfaceDeclaration(p), "should detect interface decl")
	parseInterfaceDeclaration(p)
	goutil.AssertNow(t, p.scope != nil, "scope should not be nil")
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.InterfaceDeclaration, "wrong node type")
	i := n.(*ast.InterfaceDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, i.Supers == nil, "wrong supers")
}

func TestParseInterfaceDeclarationSingleInheritance(t *testing.T) {
	p := createParser(`interface Wagable inherits Visible {}`)
	goutil.Assert(t, isInterfaceDeclaration(p), "should detect interface decl")
	parseInterfaceDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.InterfaceDeclaration, "wrong node type")
	i := n.(*ast.InterfaceDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseInterfaceDeclarationMultipleInheritance(t *testing.T) {
	p := createParser(`interface Wagable inherits Visible, Movable {}`)
	goutil.Assert(t, isInterfaceDeclaration(p), "should detect interface decl")
	parseInterfaceDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.InterfaceDeclaration, "wrong node type")
	i := n.(*ast.InterfaceDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 2, "wrong supers length")
}

func TestParseInterfaceDeclarationAbstract(t *testing.T) {
	p := createParser(`abstract interface Wagable {}`)
	goutil.Assert(t, isModifier(p), "should detect modifier")
	parseModifiers(p)
	goutil.Assert(t, isInterfaceDeclaration(p), "should detect interface decl")
	parseInterfaceDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.InterfaceDeclaration, "wrong node type")
	i := n.(*ast.InterfaceDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 0, "wrong supers length")
}

func TestParseContractDeclarationEmpty(t *testing.T) {
	p := createParser(`contract Wagable {}`)
	goutil.AssertNow(t, len(p.tokens) == 4, fmt.Sprintf("wrong token length: %d", len(p.tokens)))
	goutil.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.ContractDeclaration, "wrong node type")
	i := n.(*ast.ContractDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 0, "wrong supers length")
}

func TestParseContractDeclarationSingleInterface(t *testing.T) {
	p := createParser(`contract Wagable is Visible {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect interface decl")
	parseContractDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.ContractDeclaration, "wrong node type")
	i := n.(*ast.ContractDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 0, "wrong supers length")
}

func TestParseContractDeclarationMultipleInterfaces(t *testing.T) {
	p := createParser(`contract Wagable is Visible, Movable {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.ContractDeclaration, "wrong node type")
	i := n.(*ast.ContractDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 0, "wrong supers length")
}

func TestParseContractDeclarationSingleInterfaceSingleInheritance(t *testing.T) {
	p := createParser(`contract Wagable is Visible inherits Object {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect interface decl")
	parseContractDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.ContractDeclaration, "wrong node type")
	i := n.(*ast.ContractDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Interfaces) == 1, "wrong interfaces length")
	goutil.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseContractDeclarationMultipleInterfaceMultipleInheritance(t *testing.T) {
	p := createParser(`contract Wagable inherits A,B is Visible, Movable  {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.ContractDeclaration, "wrong node type")
	i := n.(*ast.ContractDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 2, "wrong supers length")
}

func TestParseContractDeclarationSingleInheritance(t *testing.T) {
	p := createParser(`contract Wagable inherits Visible {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.ContractDeclaration, "wrong node type")
	i := n.(*ast.ContractDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseContractDeclarationMultipleInheritance(t *testing.T) {
	p := createParser(`contract Wagable inherits Visible, Movable {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.ContractDeclaration, "wrong node type")
	i := n.(*ast.ContractDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 2, "wrong supers length")
}

func TestParseContractDeclarationSingleInheritanceMultipleInterface(t *testing.T) {
	p := createParser(`contract Wagable inherits Visible is A, B {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect interface decl")
	parseContractDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.ContractDeclaration, "wrong node type")
	i := n.(*ast.ContractDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseContractDeclarationSingleInheritanceSingleInterface(t *testing.T) {
	p := createParser(`contract Wagable inherits Object is Visible {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect interface decl")
	parseContractDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.ContractDeclaration, "wrong node type")
	i := n.(*ast.ContractDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseContractDeclarationMultipleInterfaceSingleInheritance(t *testing.T) {
	p := createParser(`contract Wagable is Visible, Movable inherits A {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.ContractDeclaration, "wrong node type")
	i := n.(*ast.ContractDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
	goutil.AssertNow(t, len(i.Interfaces) == 2, "wrong interfaces length")
}

func TestParseContractDeclarationMultipleInheritanceMultipleInterface(t *testing.T) {
	p := createParser(`contract Wagable is Visible, Movable inherits A,B {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.ContractDeclaration, "wrong node type")
	i := n.(*ast.ContractDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 2, "wrong supers length")
}

func TestParseContractDeclarationAbstract(t *testing.T) {
	p := createParser(`abstract contract Wagable {}`)
	goutil.Assert(t, isModifier(p), "should detect modifier")
	parseModifiers(p)
	goutil.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.ContractDeclaration, "wrong node type")
	i := n.(*ast.ContractDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 0, "wrong supers length")
}

func TestParseClassDeclarationSingleInterface(t *testing.T) {
	p := createParser(`class Wagable is Visible {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect interface decl")
	parseClassDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.ClassDeclaration, "wrong node type")
	i := n.(*ast.ClassDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 0, "wrong supers length")
}

func TestParseClassDeclarationMultipleInterfaces(t *testing.T) {
	p := createParser(`class Wagable is Visible, Movable {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.ClassDeclaration, "wrong node type")
	i := n.(*ast.ClassDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 0, "wrong supers length")
}

func TestParseClassDeclarationSingleInterfaceSingleInheritance(t *testing.T) {
	p := createParser(`class Wagable is Visible inherits Object {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect interface decl")
	parseClassDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.ClassDeclaration, "wrong node type")
	i := n.(*ast.ClassDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseClassDeclarationMultipleInterfaceMultipleInheritance(t *testing.T) {
	p := createParser(`class Wagable inherits A,B is Visible, Movable  {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.ClassDeclaration, "wrong node type")
	i := n.(*ast.ClassDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 2, "wrong supers length")
}

func TestParseClassDeclarationSingleInheritance(t *testing.T) {
	p := createParser(`class Wagable inherits Visible {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect interface decl")
	parseClassDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.ClassDeclaration, "wrong node type")
	i := n.(*ast.ClassDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseClassDeclarationMultipleInheritance(t *testing.T) {
	p := createParser(`class Wagable inherits Visible, Movable {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.ClassDeclaration, "wrong node type")
	i := n.(*ast.ClassDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 2, "wrong supers length")
}

func TestParseClassDeclarationSingleInheritanceMultipleInterface(t *testing.T) {
	p := createParser(`class Wagable inherits Visible {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect interface decl")
	parseClassDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.ClassDeclaration, "wrong node type")
	i := n.(*ast.ClassDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseClassDeclarationSingleInheritanceSingleInterface(t *testing.T) {
	p := createParser(`class Wagable inherits Object is Visible {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect interface decl")
	parseClassDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.ClassDeclaration, "wrong node type")
	i := n.(*ast.ClassDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseClassDeclarationMultipleInheritanceSingleInterface(t *testing.T) {
	p := createParser(`class Wagable is Visible, Movable inherits A {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.ClassDeclaration, "wrong node type")
	i := n.(*ast.ClassDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseClassDeclarationMultipleInheritanceMultipleInterface(t *testing.T) {
	p := createParser(`class Wagable is Visible, Movable inherits A,B {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.ClassDeclaration, "wrong node type")
	i := n.(*ast.ClassDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 2, "wrong supers length")
}

func TestParseClassDeclarationAbstract(t *testing.T) {
	p := createParser(`abstract class Wagable {}`)
	goutil.Assert(t, isModifier(p), "should detect modifier")
	parseModifiers(p)
	goutil.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.ClassDeclaration, "wrong node type")
	i := n.(*ast.ClassDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
}

func TestParseTypeDeclaration(t *testing.T) {
	p := createParser(`type Wagable int`)
	goutil.AssertNow(t, len(p.tokens) == 3, fmt.Sprintf("wrong token length: %d", len(p.tokens)))
	goutil.Assert(t, isTypeDeclaration(p), "should detect type decl")
	parseTypeDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.TypeDeclaration, "wrong node type")
	e := n.(*ast.TypeDeclarationNode)
	goutil.AssertNow(t, e.Identifier == "Wagable", "wrong type name")

}

func TestParseExplicitVarDeclaration(t *testing.T) {
	p := createParser(`a string`)
	goutil.Assert(t, isExplicitVarDeclaration(p), "should detect expvar decl")
	parseExplicitVarDeclaration(p)
}

func TestParseEventDeclarationEmpty(t *testing.T) {
	p := createParser(`event Notification()`)
	goutil.AssertNow(t, len(p.tokens) == 4, "wrong token length")
	goutil.Assert(t, isEventDeclaration(p), "should detect event decl")
	parseEventDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.EventDeclaration, "wrong node type")
	e := n.(*ast.EventDeclarationNode)
	goutil.AssertNow(t, len(e.Parameters) == 0, "wrong param length")
}

func TestParseEventDeclarationSingle(t *testing.T) {
	p := createParser(`event Notification(a string)`)
	goutil.Assert(t, isEventDeclaration(p), "should detect event decl")
	parseEventDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.EventDeclaration, "wrong node type")
	e := n.(*ast.EventDeclarationNode)
	goutil.AssertNow(t, len(e.Parameters) == 1, "wrong param length")
}

func TestParseEventDeclarationMultiple(t *testing.T) {
	p := createParser(`event Notification(a string, b string)`)
	goutil.Assert(t, isEventDeclaration(p), "should detect event decl")
	parseEventDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.EventDeclaration, "wrong node type")
	e := n.(*ast.EventDeclarationNode)
	goutil.AssertNow(t, len(e.Parameters) == 2, "wrong param length")
}

func TestParseEnum(t *testing.T) {
	p := createParser(`enum Weekday {}`)
	goutil.AssertNow(t, len(p.tokens) == 4, "wrong token length")
	goutil.Assert(t, isEnumDeclaration(p), "should detect enum decl")
	parseEnumDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.EnumDeclaration, "wrong node type")
	e := n.(*ast.EnumDeclarationNode)
	goutil.AssertNow(t, e.Identifier == "Weekday", "wrong identifier")
}

func TestParseEnumInheritsSingle(t *testing.T) {
	p := createParser(`enum Day inherits Weekday {}`)
	goutil.AssertNow(t, len(p.tokens) == 6, "wrong token length")
	goutil.Assert(t, isEnumDeclaration(p), "should detect enum decl")
	parseEnumDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.EnumDeclaration, "wrong node type")
	e := n.(*ast.EnumDeclarationNode)
	goutil.AssertNow(t, e.Identifier == "Day", "wrong identifier")
}

func TestParseEnumInheritsMultiple(t *testing.T) {
	p := createParser(`enum Day inherits Weekday, Weekend {}`)
	goutil.AssertNow(t, len(p.tokens) == 8, "wrong token length")
	goutil.Assert(t, isEnumDeclaration(p), "should detect enum decl")
	parseEnumDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.EnumDeclaration, "wrong node type")
	e := n.(*ast.EnumDeclarationNode)
	goutil.AssertNow(t, e.Identifier == "Day", "wrong identifier")
}

func TestParseVarDeclarationSimple(t *testing.T) {
	p := createParser("a int")
	d := p.parseVarDeclaration()
	goutil.AssertNow(t, len(d.Identifiers) == 1, "wrong id length")
	goutil.AssertNow(t, d.Identifiers[0] == "a", "wrong id 0 value")
	dt := d.DeclaredType
	goutil.AssertNow(t, dt.Type() == ast.PlainType, "wrong node type")
}

func TestParseVarDeclarationMultiple(t *testing.T) {
	p := createParser("a, b int")
	d := p.parseVarDeclaration()
	goutil.AssertNow(t, len(d.Identifiers) == 2, "wrong id length")
	goutil.AssertNow(t, d.Identifiers[0] == "a", "wrong id 0 value")
	goutil.AssertNow(t, d.Identifiers[1] == "b", "wrong id 1 value")
}

func TestParseVarDeclarationMultipleExternal(t *testing.T) {
	p := createParser("a, b pkg.Type")
	d := p.parseVarDeclaration()
	goutil.AssertNow(t, len(d.Identifiers) == 2, "wrong id length")
	goutil.AssertNow(t, d.Identifiers[0] == "a", "wrong id 0 value")
	goutil.AssertNow(t, d.Identifiers[1] == "b", "wrong id 1 value")
	dt := d.DeclaredType
	goutil.AssertNow(t, dt.Type() == ast.PlainType, "wrong node type")
}

func TestParseVarDeclarationMap(t *testing.T) {
	p := createParser("a, b map[string]string")
	d := p.parseVarDeclaration()
	goutil.AssertNow(t, len(d.Identifiers) == 2, "wrong id length")
	goutil.AssertNow(t, d.Identifiers[0] == "a", "wrong id 0 value")
	goutil.AssertNow(t, d.Identifiers[1] == "b", "wrong id 1 value")
	dt := d.DeclaredType
	goutil.AssertNow(t, dt.Type() == ast.MapType, "wrong node type")
	m := dt.(*ast.MapTypeNode)
	goutil.AssertNow(t, m.Key.Type() == ast.PlainType, "wrong key type")
	goutil.AssertNow(t, m.Value.Type() == ast.PlainType, "wrong value type")
}

func TestParseVarDeclarationArray(t *testing.T) {
	p := createParser("a, b []string")
	d := p.parseVarDeclaration()
	goutil.AssertNow(t, len(d.Identifiers) == 2, "wrong id length")
	goutil.AssertNow(t, d.Identifiers[0] == "a", "wrong id 0 value")
	goutil.AssertNow(t, d.Identifiers[1] == "b", "wrong id 1 value")
	dt := d.DeclaredType
	goutil.AssertNow(t, dt.Type() == ast.ArrayType, "wrong node type")
	m := dt.(*ast.ArrayTypeNode)
	goutil.AssertNow(t, m.Value.Type() == ast.PlainType, "wrong key type")
}

func TestParseFuncNoParameters(t *testing.T) {
	p := createParser(`func foo(){}`)
	goutil.Assert(t, isFuncDeclaration(p), "should detect func decl")
	parseFuncDeclaration(p)

	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.FuncDeclaration, "wrong node type")
	f := n.(*ast.FuncDeclarationNode)
	goutil.AssertNow(t, len(f.Signature.Parameters) == 0,
		fmt.Sprintf("wrong param length: %d", len(f.Signature.Parameters)))
}

func TestParseFuncOneParameter(t *testing.T) {
	p := createParser(`func foo(a int){}`)
	goutil.Assert(t, isFuncDeclaration(p), "should detect func decl")
	parseFuncDeclaration(p)

	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.FuncDeclaration, "wrong node type")
	f := n.(*ast.FuncDeclarationNode)
	goutil.AssertNow(t, len(f.Signature.Parameters) == 1, "wrong param length")
}

func TestParseFuncParameters(t *testing.T) {
	p := createParser(`func foo(a int, b string){}`)
	goutil.Assert(t, isFuncDeclaration(p), "should detect func decl")
	parseFuncDeclaration(p)

	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.FuncDeclaration, "wrong node type")
	f := n.(*ast.FuncDeclarationNode)
	goutil.AssertNow(t, len(f.Signature.Parameters) == 2, "wrong param length")
}

func TestParseFuncMultiplePerType(t *testing.T) {
	p := createParser(`func foo(a, b int){}`)
	goutil.Assert(t, isFuncDeclaration(p), "should detect func decl")
	parseFuncDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.FuncDeclaration, "wrong node type")
	f := n.(*ast.FuncDeclarationNode)
	goutil.AssertNow(t, len(f.Signature.Parameters) == 1, "wrong param length")
}

func TestParseFuncMultiplePerTypeExtra(t *testing.T) {
	p := createParser(`func foo(a, b int, c string){}`)
	goutil.Assert(t, isFuncDeclaration(p), "should detect func decl")
	parseFuncDeclaration(p)

	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.FuncDeclaration, "wrong node type")
	f := n.(*ast.FuncDeclarationNode)
	goutil.AssertNow(t, len(f.Signature.Parameters) == 2, "wrong param length")
}

func TestParseConstructorNoParameters(t *testing.T) {
	p := createParser(`constructor(){}`)
	goutil.Assert(t, isLifecycleDeclaration(p), "should detect Constructor decl")
	parseLifecycleDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.LifecycleDeclaration, "wrong node type")
	c := n.(*ast.LifecycleDeclarationNode)
	goutil.AssertNow(t, len(c.Parameters) == 0, "wrong param length")
}

func TestParseConstructorOneParameter(t *testing.T) {
	p := createParser(`constructor(a int){}`)
	goutil.Assert(t, isLifecycleDeclaration(p), "should detect Constructor decl")
	parseLifecycleDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.LifecycleDeclaration, "wrong node type")
	c := n.(*ast.LifecycleDeclarationNode)
	goutil.AssertNow(t, len(c.Parameters) == 1, "wrong param length")
}

func TestParseConstructorParameters(t *testing.T) {
	p := createParser(`constructor(a int, b string){}`)
	goutil.Assert(t, isLifecycleDeclaration(p), "should detect Constructor decl")
	parseLifecycleDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.LifecycleDeclaration, "wrong node type")
	c := n.(*ast.LifecycleDeclarationNode)
	goutil.AssertNow(t, len(c.Parameters) == 2, "wrong param length")
}

func TestParseConstructorMultiplePerType(t *testing.T) {
	p := createParser(`constructor(a, b int){}`)
	goutil.Assert(t, isLifecycleDeclaration(p), "should detect Constructor decl")
	parseLifecycleDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.LifecycleDeclaration, "wrong node type")
	c := n.(*ast.LifecycleDeclarationNode)
	goutil.AssertNow(t, len(c.Parameters) == 1, "wrong param length")
}

func TestParseConstructorMultiplePerTypeExtra(t *testing.T) {
	p := createParser(`constructor(a, b int, c []string){}`)
	goutil.Assert(t, isLifecycleDeclaration(p), "should detect Constructor decl")
	parseLifecycleDeclaration(p)
	n := p.scope.NextDeclaration()
	goutil.AssertNow(t, n.Type() == ast.LifecycleDeclaration, "wrong node type")
	c := n.(*ast.LifecycleDeclarationNode)
	goutil.AssertNow(t, len(c.Parameters) == 2, "wrong param length")
	first := c.Parameters[0]
	second := c.Parameters[1]
	goutil.AssertNow(t, first.DeclaredType != nil, "first dt shouldn't be nil")
	goutil.AssertNow(t, first.DeclaredType.Type() == ast.PlainType, "first dt should be plain")
	goutil.AssertNow(t, second.DeclaredType != nil, "second dt shouldn't be nil")
	goutil.AssertNow(t, second.DeclaredType.Type() == ast.ArrayType, "second dt should be array")
}

func TestParseParametersSingleVarSingleType(t *testing.T) {
	p := createParser(`(a string)`)
	exps := p.parseParameters()
	goutil.AssertNow(t, exps != nil, "params not nil")
	goutil.AssertNow(t, len(exps) == 1, "params of length 1")
	goutil.AssertNow(t, exps[0].DeclaredType != nil, "declared type shouldn't be nil")
	goutil.AssertNow(t, exps[0].DeclaredType.Type() == ast.PlainType, "wrong declared type")
	goutil.AssertNow(t, len(exps[0].Identifiers) == 1, "wrong parameter length")
}

func TestParseParametersMultipleVarSingleType(t *testing.T) {
	p := createParser(`(a, b string)`)
	exps := p.parseParameters()
	goutil.AssertNow(t, exps != nil, "params not nil")
	goutil.AssertNow(t, len(exps) == 1, "params of length 1")
	goutil.AssertNow(t, exps[0].DeclaredType != nil, "declared type shouldn't be nil")
	goutil.AssertNow(t, exps[0].DeclaredType.Type() == ast.PlainType, "wrong declared type")
	goutil.AssertNow(t, len(exps[0].Identifiers) == 2, "wrong parameter length")
}

func TestParseParametersSingleVarMultipleType(t *testing.T) {
	p := createParser(`(a string, b int)`)
	exps := p.parseParameters()
	goutil.AssertNow(t, exps != nil, "params not nil")
	goutil.AssertNow(t, len(exps) == 2, "params of length 1")
	goutil.AssertNow(t, exps[0].DeclaredType != nil, "declared type shouldn't be nil")
	goutil.AssertNow(t, exps[0].DeclaredType.Type() == ast.PlainType, "wrong declared type")
	goutil.AssertNow(t, exps[1].DeclaredType != nil, "declared type shouldn't be nil")
	goutil.AssertNow(t, exps[1].DeclaredType.Type() == ast.PlainType, "wrong declared type")
	goutil.AssertNow(t, len(exps[0].Identifiers) == 1, "wrong parameter length")
	goutil.AssertNow(t, len(exps[1].Identifiers) == 1, "wrong parameter length")
}
