package parser

import (
	"fmt"
	"testing"

	"github.com/end-r/guardian/compiler/ast"

	"github.com/end-r/goutil"
)

func TestParseInterfaceDeclarationEmpty(t *testing.T) {
	p := createParser(`interface Wagable {}`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 4, "wrong token length")
	goutil.Assert(t, isInterfaceDeclaration(p), "should detect interface decl")
	parseInterfaceDeclaration(p)
	goutil.AssertNow(t, len(p.Scope.Nodes("interface")) == 1, "wrong node count")
	n := p.Scope.Nodes(interfaceKey)[0]
	goutil.AssertNow(t, n.Type() == ast.InterfaceDeclaration, "wrong node type")
	i := n.(ast.InterfaceDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, i.Supers == nil, "wrong supers")
}

func TestParseInterfaceDeclarationSingleInheritance(t *testing.T) {
	p := createParser(`interface Wagable inherits Visible {}`)
	goutil.Assert(t, isInterfaceDeclaration(p), "should detect interface decl")
	parseInterfaceDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("interface")) == 1, "wrong node count")
	n := p.Scope.Nodes(interfaceKey)[0]
	goutil.AssertNow(t, n.Type() == ast.InterfaceDeclaration, "wrong node type")
	i := n.(ast.InterfaceDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
	goutil.AssertNow(t, i.Supers[0].Names[0] == "Visible", "wrong supers 0 name")
}

func TestParseInterfaceDeclarationMultipleInheritance(t *testing.T) {
	p := createParser(`interface Wagable inherits Visible, Movable {}`)
	goutil.Assert(t, isInterfaceDeclaration(p), "should detect interface decl")
	parseInterfaceDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("interface")) == 1, "wrong node count")
	n := p.Scope.Nodes(interfaceKey)[0]
	goutil.AssertNow(t, n.Type() == ast.InterfaceDeclaration, "wrong node type")
	i := n.(ast.InterfaceDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 2, "wrong supers length")
	goutil.AssertNow(t, i.Supers[0].Names[0] == "Visible", "wrong supers 0 name")
	goutil.AssertNow(t, i.Supers[1].Names[0] == "Movable", "wrong supers 1 name")
}

func TestParseInterfaceDeclarationAbstract(t *testing.T) {
	p := createParser(`abstract interface Wagable {}`)
	goutil.Assert(t, isInterfaceDeclaration(p), "should detect interface decl")
	parseInterfaceDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("interface")) == 1, "wrong node count")
	n := p.Scope.Nodes(interfaceKey)[0]
	goutil.AssertNow(t, n.Type() == ast.InterfaceDeclaration, "wrong node type")
	i := n.(ast.InterfaceDeclarationNode)
	goutil.AssertNow(t, i.IsAbstract, "wrong abstract")
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 0, "wrong supers length")
}

func TestParseContractDeclarationEmpty(t *testing.T) {
	p := createParser(`contract Wagable {}`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 4, fmt.Sprintf("wrong token length: %d", len(p.lexer.Tokens)))
	goutil.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(contractKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(contractKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ContractDeclaration, "wrong node type")
	i := n.(ast.ContractDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 0, "wrong supers length")
}

func TestParseContractDeclarationSingleInterface(t *testing.T) {
	p := createParser(`contract Wagable is Visible {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect interface decl")
	parseContractDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(contractKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(contractKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ContractDeclaration, "wrong node type")
	i := n.(ast.ContractDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 0, "wrong supers length")
}

func TestParseContractDeclarationMultipleInterfaces(t *testing.T) {
	p := createParser(`contract Wagable is Visible, Movable {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(contractKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(contractKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ContractDeclaration, "wrong node type")
	i := n.(ast.ContractDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 0, "wrong supers length")
}

func TestParseContractDeclarationSingleInterfaceSingleInheritance(t *testing.T) {
	p := createParser(`contract Wagable is Visible inherits Object {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect interface decl")
	parseContractDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(contractKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(contractKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ContractDeclaration, "wrong node type")
	i := n.(ast.ContractDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Interfaces) == 1, "wrong interfaces length")
	goutil.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseContractDeclarationMultipleInterfaceMultipleInheritance(t *testing.T) {
	p := createParser(`contract Wagable inherits A,B is Visible, Movable  {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("contract")) == 1, "wrong node count")
	goutil.Assert(t, len(p.Scope.Nodes(contractKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(contractKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ContractDeclaration, "wrong node type")
	i := n.(ast.ContractDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 2, "wrong supers length")
}

func TestParseContractDeclarationSingleInheritance(t *testing.T) {
	p := createParser(`contract Wagable inherits Visible {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(contractKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(contractKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ContractDeclaration, "wrong node type")
	i := n.(ast.ContractDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseContractDeclarationMultipleInheritance(t *testing.T) {
	p := createParser(`contract Wagable inherits Visible, Movable {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(contractKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(contractKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ContractDeclaration, "wrong node type")
	i := n.(ast.ContractDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 2, "wrong supers length")
}

func TestParseContractDeclarationSingleInheritanceMultipleInterface(t *testing.T) {
	p := createParser(`contract Wagable inherits Visible is A, B {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect interface decl")
	parseContractDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(contractKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(contractKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ContractDeclaration, "wrong node type")
	i := n.(ast.ContractDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseContractDeclarationSingleInheritanceSingleInterface(t *testing.T) {
	p := createParser(`contract Wagable inherits Object is Visible {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect interface decl")
	parseContractDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(contractKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(contractKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ContractDeclaration, "wrong node type")
	i := n.(ast.ContractDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseContractDeclarationMultipleInterfaceSingleInheritance(t *testing.T) {
	p := createParser(`contract Wagable is Visible, Movable inherits A {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("contract")) == 1, "wrong node count")
	goutil.Assert(t, len(p.Scope.Nodes(contractKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(contractKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ContractDeclaration, "wrong node type")
	i := n.(ast.ContractDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
	goutil.AssertNow(t, len(i.Interfaces) == 2, "wrong interfaces length")
}

func TestParseContractDeclarationMultipleInheritanceMultipleInterface(t *testing.T) {
	p := createParser(`contract Wagable is Visible, Movable inherits A,B {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(contractKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(contractKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ContractDeclaration, "wrong node type")
	i := n.(ast.ContractDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 2, "wrong supers length")
}

func TestParseContractDeclarationAbstract(t *testing.T) {
	p := createParser(`abstract contract Wagable {}`)
	goutil.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(contractKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(contractKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ContractDeclaration, "wrong node type")
	i := n.(ast.ContractDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 0, "wrong supers length")
}

func TestParseClassDeclarationSingleInterface(t *testing.T) {
	p := createParser(`class Wagable is Visible {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect interface decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(classKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(classKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ClassDeclaration, "wrong node type")
	i := n.(ast.ClassDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 0, "wrong supers length")
}

func TestParseClassDeclarationMultipleInterfaces(t *testing.T) {
	p := createParser(`class Wagable is Visible, Movable {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(classKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(classKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ClassDeclaration, "wrong node type")
	i := n.(ast.ClassDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 0, "wrong supers length")
}

func TestParseClassDeclarationSingleInterfaceSingleInheritance(t *testing.T) {
	p := createParser(`class Wagable is Visible inherits Object {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect interface decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(classKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(classKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ClassDeclaration, "wrong node type")
	i := n.(ast.ClassDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseClassDeclarationMultipleInterfaceMultipleInheritance(t *testing.T) {
	p := createParser(`class Wagable inherits A,B is Visible, Movable  {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(classKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(classKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ClassDeclaration, "wrong node type")
	i := n.(ast.ClassDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 2, "wrong supers length")
}

func TestParseClassDeclarationSingleInheritance(t *testing.T) {
	p := createParser(`class Wagable inherits Visible {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect interface decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(classKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(classKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ClassDeclaration, "wrong node type")
	i := n.(ast.ClassDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseClassDeclarationMultipleInheritance(t *testing.T) {
	p := createParser(`class Wagable inherits Visible, Movable {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(classKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(classKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ClassDeclaration, "wrong node type")
	i := n.(ast.ClassDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 2, "wrong supers length")
}

func TestParseClassDeclarationSingleInheritanceMultipleInterface(t *testing.T) {
	p := createParser(`class Wagable inherits Visible {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect interface decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(classKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(classKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ClassDeclaration, "wrong node type")
	i := n.(ast.ClassDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseClassDeclarationSingleInheritanceSingleInterface(t *testing.T) {
	p := createParser(`class Wagable inherits Object is Visible {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect interface decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(classKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(classKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ClassDeclaration, "wrong node type")
	i := n.(ast.ClassDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseClassDeclarationMultipleInheritanceSingleInterface(t *testing.T) {
	p := createParser(`class Wagable is Visible, Movable inherits A {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(classKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(classKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ClassDeclaration, "wrong node type")
	i := n.(ast.ClassDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseClassDeclarationMultipleInheritanceMultipleInterface(t *testing.T) {
	p := createParser(`class Wagable is Visible, Movable inherits A,B {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(classKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(classKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ClassDeclaration, "wrong node type")
	i := n.(ast.ClassDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, len(i.Supers) == 2, "wrong supers length")
}

func TestParseClassDeclarationAbstract(t *testing.T) {
	p := createParser(`abstract class Wagable {}`)
	goutil.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes("class")) == 1, "wrong node count")
	goutil.Assert(t, len(p.Scope.Nodes(classKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(classKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ClassDeclaration, "wrong node type")
	i := n.(ast.ClassDeclarationNode)
	goutil.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	goutil.AssertNow(t, i.IsAbstract, "wrong abstract")
}

func TestParseTypeDeclaration(t *testing.T) {
	p := createParser(`type Wagable int`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 3, fmt.Sprintf("wrong token length: %d", len(p.lexer.Tokens)))
	goutil.Assert(t, isTypeDeclaration(p), "should detect type decl")
	parseTypeDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(typeKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(typeKey)[0]
	goutil.AssertNow(t, n.Type() == ast.TypeDeclaration, "wrong node type")
	e := n.(ast.TypeDeclarationNode)
	goutil.AssertNow(t, e.Identifier == "Wagable", "wrong type name")

}

func TestParseExplicitVarDeclaration(t *testing.T) {
	p := createParser(`a string`)
	goutil.Assert(t, isExplicitVarDeclaration(p), "should detect expvar decl")
	parseExplicitVarDeclaration(p)
}

func TestParseEventDeclarationEmpty(t *testing.T) {
	p := createParser(`event Notification()`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 4, "wrong token length")
	goutil.Assert(t, isEventDeclaration(p), "should detect event decl")
	parseEventDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(eventKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(eventKey)[0]
	goutil.AssertNow(t, n.Type() == ast.EventDeclaration, "wrong node type")
	e := n.(ast.EventDeclarationNode)
	goutil.AssertNow(t, len(e.Parameters) == 0, "wrong param length")
}

func TestParseEventDeclarationSingle(t *testing.T) {
	p := createParser(`event Notification(string)`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 5, "wrong token length")
	goutil.Assert(t, isEventDeclaration(p), "should detect event decl")
	parseEventDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(eventKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(eventKey)[0]
	goutil.AssertNow(t, n.Type() == ast.EventDeclaration, "wrong node type")
	e := n.(ast.EventDeclarationNode)
	goutil.AssertNow(t, len(e.Parameters) == 1, "wrong param length")
}

func TestParseEventDeclarationMultiple(t *testing.T) {
	p := createParser(`event Notification(string, string)`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 7, "wrong token length")
	goutil.Assert(t, isEventDeclaration(p), "should detect event decl")
	parseEventDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(eventKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(eventKey)[0]
	goutil.AssertNow(t, n.Type() == ast.EventDeclaration, "wrong node type")
	e := n.(ast.EventDeclarationNode)
	goutil.AssertNow(t, len(e.Parameters) == 2, "wrong param length")
}

func TestParseEnum(t *testing.T) {
	p := createParser(`enum Weekday {}`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 4, "wrong token length")
	goutil.Assert(t, isEnumDeclaration(p), "should detect enum decl")
	parseEnumDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(enumKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(enumKey)[0]
	goutil.AssertNow(t, n.Type() == ast.EnumDeclaration, "wrong node type")
	e := n.(ast.EnumDeclarationNode)
	goutil.AssertNow(t, e.Identifier == "Weekday", "wrong identifier")
}

func TestParseEnumInheritsSingle(t *testing.T) {
	p := createParser(`enum Day inherits Weekday {}`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 6, "wrong token length")
	goutil.Assert(t, isEnumDeclaration(p), "should detect enum decl")
	parseEnumDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(enumKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(enumKey)[0]
	goutil.AssertNow(t, n.Type() == ast.EnumDeclaration, "wrong node type")
	e := n.(ast.EnumDeclarationNode)
	goutil.AssertNow(t, e.Identifier == "Day", "wrong identifier")
}

func TestParseEnumInheritsMultiple(t *testing.T) {
	p := createParser(`enum Day inherits Weekday, Weekend {}`)
	goutil.AssertNow(t, len(p.lexer.Tokens) == 8, "wrong token length")
	goutil.Assert(t, isEnumDeclaration(p), "should detect enum decl")
	parseEnumDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(enumKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(enumKey)[0]
	goutil.AssertNow(t, n.Type() == ast.EnumDeclaration, "wrong node type")
	e := n.(ast.EnumDeclarationNode)
	goutil.AssertNow(t, e.Identifier == "Day", "wrong identifier")
}

func TestParseVarDeclarationSimple(t *testing.T) {
	p := createParser("a int")
	d := p.parseVarDeclaration()
	goutil.AssertNow(t, len(d.Identifiers) == 1, "wrong id length")
	goutil.AssertNow(t, d.Identifiers[0] == "a", "wrong id 0 value")
	dt := d.DeclaredType
	goutil.AssertNow(t, dt.Type() == ast.Reference, "wrong node type")
	r := dt.(ast.ReferenceNode)
	goutil.AssertNow(t, len(r.Names) == 1, "wrong ref length")
	goutil.AssertNow(t, r.Names[0] == "int", "wrong ref 0 value")
}

func TestParseVarDeclarationMultiple(t *testing.T) {
	p := createParser("a, b int")
	d := p.parseVarDeclaration()
	goutil.AssertNow(t, len(d.Identifiers) == 2, "wrong id length")
	goutil.AssertNow(t, d.Identifiers[0] == "a", "wrong id 0 value")
	goutil.AssertNow(t, d.Identifiers[1] == "b", "wrong id 1 value")
	dt := d.DeclaredType
	goutil.AssertNow(t, dt.Type() == ast.Reference, "wrong node type")
	r := dt.(ast.ReferenceNode)
	goutil.AssertNow(t, len(r.Names) == 1, "wrong ref length")
	goutil.AssertNow(t, r.Names[0] == "int", "wrong ref 0 value")
}

func TestParseVarDeclarationMultipleExternal(t *testing.T) {
	p := createParser("a, b pkg.Type")
	d := p.parseVarDeclaration()
	goutil.AssertNow(t, len(d.Identifiers) == 2, "wrong id length")
	goutil.AssertNow(t, d.Identifiers[0] == "a", "wrong id 0 value")
	goutil.AssertNow(t, d.Identifiers[1] == "b", "wrong id 1 value")
	dt := d.DeclaredType
	goutil.AssertNow(t, dt.Type() == ast.Reference, "wrong node type")
	r := dt.(ast.ReferenceNode)
	goutil.AssertNow(t, len(r.Names) == 2, fmt.Sprintf("wrong ref length: %d", len(r.Names)))
	goutil.AssertNow(t, r.Names[0] == "pkg", "wrong ref 0 value")
	goutil.AssertNow(t, r.Names[1] == "Type", "wrong ref 1 value")
}

func TestParseVarDeclarationMap(t *testing.T) {
	p := createParser("a, b map[string]string")
	d := p.parseVarDeclaration()
	goutil.AssertNow(t, len(d.Identifiers) == 2, "wrong id length")
	goutil.AssertNow(t, d.Identifiers[0] == "a", "wrong id 0 value")
	goutil.AssertNow(t, d.Identifiers[1] == "b", "wrong id 1 value")
	dt := d.DeclaredType
	goutil.AssertNow(t, dt.Type() == ast.MapType, "wrong node type")
	m := dt.(ast.MapTypeNode)
	goutil.AssertNow(t, m.Key.Type() == ast.Reference, "wrong key type")
	goutil.AssertNow(t, m.Value.Type() == ast.Reference, "wrong value type")
}

func TestParseVarDeclarationArray(t *testing.T) {
	p := createParser("a, b [string]")
	d := p.parseVarDeclaration()
	goutil.AssertNow(t, len(d.Identifiers) == 2, "wrong id length")
	goutil.AssertNow(t, d.Identifiers[0] == "a", "wrong id 0 value")
	goutil.AssertNow(t, d.Identifiers[1] == "b", "wrong id 1 value")
	dt := d.DeclaredType
	goutil.AssertNow(t, dt.Type() == ast.ArrayType, "wrong node type")
	m := dt.(ast.ArrayTypeNode)
	goutil.AssertNow(t, m.Value.Type() == ast.Reference, "wrong key type")
}

func TestParseFuncNoParameters(t *testing.T) {
	p := createParser(`func foo(){}`)
	goutil.Assert(t, isFuncDeclaration(p), "should detect func decl")
	parseFuncDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(funcKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(funcKey)[0]
	goutil.AssertNow(t, n.Type() == ast.FuncDeclaration, "wrong node type")
	f := n.(ast.FuncDeclarationNode)
	goutil.AssertNow(t, len(f.Parameters) == 0,
		fmt.Sprintf("wrong param length: %d", len(f.Parameters)))
}

func TestParseFuncOneParameter(t *testing.T) {
	p := createParser(`func foo(a int){}`)
	goutil.Assert(t, isFuncDeclaration(p), "should detect func decl")
	parseFuncDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(funcKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(funcKey)[0]
	goutil.AssertNow(t, n.Type() == ast.FuncDeclaration, "wrong node type")
	f := n.(ast.FuncDeclarationNode)
	goutil.AssertNow(t, len(f.Parameters) == 1, "wrong param length")
}

func TestParseFuncParameters(t *testing.T) {
	p := createParser(`func foo(a int, b string){}`)
	goutil.Assert(t, isFuncDeclaration(p), "should detect func decl")
	parseFuncDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(funcKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(funcKey)[0]
	goutil.AssertNow(t, n.Type() == ast.FuncDeclaration, "wrong node type")
	f := n.(ast.FuncDeclarationNode)
	goutil.AssertNow(t, len(f.Parameters) == 2, "wrong param length")
}

func TestParseFuncMultiplePerType(t *testing.T) {
	p := createParser(`func foo(a, b int){}`)
	goutil.Assert(t, isFuncDeclaration(p), "should detect func decl")
	parseFuncDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(funcKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(funcKey)[0]
	goutil.AssertNow(t, n.Type() == ast.FuncDeclaration, "wrong node type")
	f := n.(ast.FuncDeclarationNode)
	goutil.AssertNow(t, len(f.Parameters) == 1, "wrong param length")
}

func TestParseFuncMultiplePerTypeExtra(t *testing.T) {
	p := createParser(`func foo(a, b int, c string){}`)
	goutil.Assert(t, isFuncDeclaration(p), "should detect func decl")
	parseFuncDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(funcKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(funcKey)[0]
	goutil.AssertNow(t, n.Type() == ast.FuncDeclaration, "wrong node type")
	f := n.(ast.FuncDeclarationNode)
	goutil.AssertNow(t, len(f.Parameters) == 2, "wrong param length")
}

func TestParseConstructorNoParameters(t *testing.T) {
	p := createParser(`constructor(){}`)
	goutil.Assert(t, isConstructorDeclaration(p), "should detect Constructor decl")
	parseConstructorDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(constructorKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(constructorKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ConstructorDeclaration, "wrong node type")
	c := n.(ast.ConstructorDeclarationNode)
	goutil.AssertNow(t, len(c.Parameters) == 0, "wrong param length")
}

func TestParseConstructorOneParameter(t *testing.T) {
	p := createParser(`constructor(a int){}`)
	goutil.Assert(t, isConstructorDeclaration(p), "should detect Constructor decl")
	parseConstructorDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(constructorKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(constructorKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ConstructorDeclaration, "wrong node type")
	c := n.(ast.ConstructorDeclarationNode)
	goutil.AssertNow(t, len(c.Parameters) == 1, "wrong param length")
}

func TestParseConstructorParameters(t *testing.T) {
	p := createParser(`constructor(a int, b string){}`)
	goutil.Assert(t, isConstructorDeclaration(p), "should detect Constructor decl")
	parseConstructorDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(constructorKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(constructorKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ConstructorDeclaration, "wrong node type")
	c := n.(ast.ConstructorDeclarationNode)
	goutil.AssertNow(t, len(c.Parameters) == 2, "wrong param length")
}

func TestParseConstructorMultiplePerType(t *testing.T) {
	p := createParser(`constructor(a, b int){}`)
	goutil.Assert(t, isConstructorDeclaration(p), "should detect Constructor decl")
	parseConstructorDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(constructorKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(constructorKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ConstructorDeclaration, "wrong node type")
	c := n.(ast.ConstructorDeclarationNode)
	goutil.AssertNow(t, len(c.Parameters) == 1, "wrong param length")
}

func TestParseConstructorMultiplePerTypeExtra(t *testing.T) {
	p := createParser(`constructor(a, b int, c string){}`)
	goutil.Assert(t, isConstructorDeclaration(p), "should detect Constructor decl")
	parseConstructorDeclaration(p)
	goutil.Assert(t, len(p.Scope.Nodes(constructorKey)) == 1, "wrong node count")
	n := p.Scope.Nodes(constructorKey)[0]
	goutil.AssertNow(t, n.Type() == ast.ConstructorDeclaration, "wrong node type")
	c := n.(ast.ConstructorDeclarationNode)
	goutil.AssertNow(t, len(c.Parameters) == 2, "wrong param length")
}
