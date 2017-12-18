package parser

import "testing"

func TestParseSimpleClassGeneric(t *testing.T) {
	p := createParser("class List<T> {}")
	parseClassDeclaration(p)
}

func TestParseExtendingClassGeneric(t *testing.T) {
	p := createParser(`class List<T extends Item>`)
	parseClassDeclaration(p)
}

func TestParseImplementingClassGeneric(t *testing.T) {
	p := createParser(`class List<T extends Item>`)
	parseClassDeclaration(p)
}

func TestParseExtendsImplementsClassGeneric(t *testing.T) {
	p := createParser(`class List<T extends Item is Comparable>`)
	parseClassDeclaration(p)
}

func TestParseImplementsExtendsClassGeneric(t *testing.T) {
	p := createParser(`class List<T is Comparable extends Item>`)
	parseClassDeclaration(p)
}

func TestParseMultipleExtendsandImplementsClassGeneric(t *testing.T) {
	p := createParser(`class List<T is Comparable, Real extends Item, Dog>`)
	parseClassDeclaration(p)
}

func TestParseSimpleClassMultipleGeneric(t *testing.T) {
	p := createParser("class List<T, S, R> {}")
	parseClassDeclaration(p)
}

func TestParseExtendingClassMultipleGeneric(t *testing.T) {
	p := createParser(`class List<T extends Item, S extends Dog>`)
	parseClassDeclaration(p)
}

func TestParseImplementingClassMultipleGeneric(t *testing.T) {
	p := createParser(`class List<T is Item, S is Dog>`)
	parseClassDeclaration(p)
}

func TestParseExtendsImplementsClassMultipleGeneric(t *testing.T) {
	p := createParser(`class List<T extends Item is Comparable, S is Comparable>`)
	parseClassDeclaration(p)
}

func TestParseImplementsExtendsClassMultipleGeneric(t *testing.T) {
	p := createParser(`class List<T is Comparable extends Item, S extends Item>`)
	parseClassDeclaration(p)
}

func TestParseMultipleExtendsandImplementsClassMultipleGeneric(t *testing.T) {
	p := createParser(`class List<T is Comparable, Real extends Item, Dog, S, R is Comparable extends Item, Dog>`)
	parseClassDeclaration(p)
}
