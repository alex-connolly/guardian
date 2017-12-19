package parser

import (
	"fmt"
	"testing"

	"github.com/end-r/guardian/ast"

	"github.com/end-r/goutil"
)

func TestParseSimpleClassGeneric(t *testing.T) {
	p := createParser("class List<T> {}")
	parseClassDeclaration(p)
	goutil.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.ClassDeclarationNode)
	goutil.AssertNow(t, c.Generics != nil, "nil generics")
	goutil.AssertNow(t, len(c.Generics) == 1, fmt.Sprintf("wrong gen len: %d", len(c.Generics)))

}

func TestParseExtendingClassGeneric(t *testing.T) {
	p := createParser(`class List<T inherits Item> {}`)
	parseClassDeclaration(p)
	goutil.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.ClassDeclarationNode)
	goutil.AssertNow(t, c.Generics != nil, "nil generics")
	goutil.AssertNow(t, len(c.Generics) == 1, fmt.Sprintf("wrong gen len: %d", len(c.Generics)))

}

func TestParseImplementingClassGeneric(t *testing.T) {
	p := createParser(`class List<T inherits Item> {}`)
	parseClassDeclaration(p)
	goutil.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.ClassDeclarationNode)
	goutil.AssertNow(t, c.Generics != nil, "nil generics")
	goutil.AssertNow(t, len(c.Generics) == 1, fmt.Sprintf("wrong gen len: %d", len(c.Generics)))

}

func TestParseExtendsImplementsClassGeneric(t *testing.T) {
	p := createParser(`class List<T inherits Item is Comparable> {}`)
	parseClassDeclaration(p)
	goutil.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.ClassDeclarationNode)
	goutil.AssertNow(t, c.Generics != nil, "nil generics")
	goutil.AssertNow(t, len(c.Generics) == 1, fmt.Sprintf("wrong gen len: %d", len(c.Generics)))

}

func TestParseImplementsExtendsClassGeneric(t *testing.T) {
	p := createParser(`class List<T is Comparable inherits Item> {}`)
	parseClassDeclaration(p)
	goutil.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.ClassDeclarationNode)
	goutil.AssertNow(t, c.Generics != nil, "nil generics")
	goutil.AssertNow(t, len(c.Generics) == 1, fmt.Sprintf("wrong gen len: %d", len(c.Generics)))

}

func TestParseMultipleExtendsandImplementsClassGeneric(t *testing.T) {
	p := createParser(`class List<T is Comparable, Real inherits Item, Dog> {}`)
	parseClassDeclaration(p)
	goutil.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.ClassDeclarationNode)
	goutil.AssertNow(t, c.Generics != nil, "nil generics")
	goutil.AssertNow(t, len(c.Generics) == 1, fmt.Sprintf("wrong gen len: %d", len(c.Generics)))

}

func TestParseSimpleClassMultipleGeneric(t *testing.T) {
	p := createParser("class List<T|S|R> {}")
	parseClassDeclaration(p)
	goutil.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.ClassDeclarationNode)
	goutil.AssertNow(t, c.Generics != nil, "nil generics")
	goutil.AssertNow(t, len(c.Generics) == 3, fmt.Sprintf("wrong gen len: %d", len(c.Generics)))

}

func TestParseExtendingClassMultipleGeneric(t *testing.T) {
	p := createParser(`class List<T inherits Item | S inherits Dog> {}`)
	parseClassDeclaration(p)
	goutil.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.ClassDeclarationNode)
	goutil.AssertNow(t, c.Generics != nil, "nil generics")
	goutil.AssertNow(t, len(c.Generics) == 2, fmt.Sprintf("wrong gen len: %d", len(c.Generics)))

}

func TestParseImplementingClassMultipleGeneric(t *testing.T) {
	p := createParser(`class List<T is Item | S is Dog> {}`)
	parseClassDeclaration(p)
	goutil.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.ClassDeclarationNode)
	goutil.AssertNow(t, c.Generics != nil, "nil generics")
	goutil.AssertNow(t, len(c.Generics) == 2, fmt.Sprintf("wrong gen len: %d", len(c.Generics)))
}

func TestParseExtendsImplementsClassMultipleGeneric(t *testing.T) {
	p := createParser(`class List<T inherits Item is Comparable | S is Comparable> {}`)
	parseClassDeclaration(p)
	goutil.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.ClassDeclarationNode)
	goutil.AssertNow(t, c.Generics != nil, "nil generics")
	goutil.AssertNow(t, len(c.Generics) == 2, fmt.Sprintf("wrong gen len: %d", len(c.Generics)))

}

func TestParseImplementsExtendsClassMultipleGeneric(t *testing.T) {
	p := createParser(`class List<T is Comparable inherits Item | S inherits Item> {}`)
	parseClassDeclaration(p)
	goutil.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.ClassDeclarationNode)
	goutil.AssertNow(t, c.Generics != nil, "nil generics")
	goutil.AssertNow(t, len(c.Generics) == 2, fmt.Sprintf("wrong gen len: %d", len(c.Generics)))

}

func TestParseMultipleExtendsandImplementsClassMultipleGeneric(t *testing.T) {
	p := createParser(`class List<T is Comparable, Real inherits Item, Dog | S | R is Comparable inherits Item, Dog> {}`)
	parseClassDeclaration(p)
	goutil.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.ClassDeclarationNode)
	goutil.AssertNow(t, c.Generics != nil, "nil generics")
	goutil.AssertNow(t, len(c.Generics) == 3, fmt.Sprintf("wrong gen len: %d", len(c.Generics)))

}

func TestParseSingleGenericFunction(t *testing.T) {
	p := createParser(`func <T> hello(){}`)
	parseFuncDeclaration(p)
	goutil.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.FuncDeclarationNode)
	goutil.AssertNow(t, c.Generics != nil, "nil generics")
	goutil.AssertNow(t, len(c.Generics) == 1, fmt.Sprintf("wrong gen len: %d", len(c.Generics)))
}

func TestParseMultipleGenericFunction(t *testing.T) {
	p := createParser(`func <T|S|R> hello(){}`)
	parseFuncDeclaration(p)
	goutil.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.FuncDeclarationNode)
	goutil.AssertNow(t, c.Generics != nil, "nil generics")
	goutil.AssertLength(t, len(c.Generics), 3)
}

func TestParseGenerics(t *testing.T) {
	p := createParser("<T>")
	gens := p.parseGenerics()
	goutil.AssertNow(t, gens != nil, "nil generics")
	goutil.AssertLength(t, len(gens), 1)
}

func TestParseGenericsDouble(t *testing.T) {
	p := createParser("<T|S>")
	gens := p.parseGenerics()
	goutil.AssertNow(t, gens != nil, "nil generics")
	goutil.AssertLength(t, len(gens), 2)
}

func TestParseGenericsTriple(t *testing.T) {
	p := createParser("<T|S|R>")
	gens := p.parseGenerics()
	goutil.AssertNow(t, gens != nil, "nil generics")
	goutil.AssertLength(t, len(gens), 3)
}

func TestParseGenericsInheritance(t *testing.T) {
	p := createParser("<T inherits A>")
	gens := p.parseGenerics()
	goutil.AssertNow(t, gens != nil, "nil generics")
	goutil.AssertLength(t, len(gens), 1)
	goutil.AssertLength(t, len(gens[0].Inherits), 1)
}

func TestParseGenericsImplementation(t *testing.T) {
	p := createParser("<T is A>")
	gens := p.parseGenerics()
	goutil.AssertNow(t, gens != nil, "nil generics")
	goutil.AssertLength(t, len(gens), 1)
	goutil.AssertLength(t, len(gens[0].Implements), 1)
}

func TestParseGenericsImplementationDouble(t *testing.T) {
	p := createParser("<T is A|S is B>")
	gens := p.parseGenerics()
	goutil.AssertNow(t, gens != nil, "nil generics")
	goutil.AssertLength(t, len(gens), 2)
	goutil.AssertLength(t, len(gens[0].Implements), 1)
	goutil.AssertLength(t, len(gens[1].Implements), 1)
}

func TestParseGenericComplex(t *testing.T) {
	text := "<T is Comparable, Real inherits Item, Dog | S | R is Comparable inherits Item, Dog>"
	p := createParser(text)
	goutil.AssertNow(t, p.index == 0, "wrong starting index")
	gens := p.parseGenerics()
	goutil.AssertNow(t, gens != nil, "nil generics")
	goutil.AssertLength(t, len(gens), 3)
	goutil.AssertNow(t, p.index == len(p.tokens), "wrong ending index")
}
