package parser

import (
	"testing"

	"github.com/end-r/guardian/ast"

	"github.com/end-r/goutil"
)

func TestClassSimpleRecognisedModifiers(t *testing.T) {
	a, errs := ParseString(`
		public static class Dog {

		}
	`)
	goutil.AssertNow(t, a != nil, "ast is nil")
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
	goutil.AssertLength(t, a.Declarations.Length(), 1)
	n := a.Declarations.Next()
	c := n.(*ast.ClassDeclarationNode)
	goutil.AssertLength(t, len(c.Modifiers.Modifiers), 2)
}

func TestClassSimpleUnrecognisedModifiers(t *testing.T) {
	a, errs := ParseString(`
		relaxed class Dog {

		}
	`)
	goutil.AssertNow(t, a != nil, "ast is nil")
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
	goutil.AssertLength(t, a.Declarations.Length(), 1)
	n := a.Declarations.Next()
	c := n.(*ast.ClassDeclarationNode)
	goutil.AssertLength(t, len(c.Modifiers.Modifiers), 0)
}

func TestGroupedClassSimpleRecognisedModifiers(t *testing.T) {
	a, errs := ParseString(`
		public static (
			class Dog {

			}

			class Cat {

			}
		)
	`)
	goutil.AssertNow(t, a != nil, "ast is nil")
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
	goutil.AssertNow(t, a.Declarations != nil, "nil declarations")
	goutil.AssertLength(t, a.Declarations.Length(), 2)
	n := a.Declarations.Next()
	c := n.(*ast.ClassDeclarationNode)
	goutil.AssertLength(t, len(c.Modifiers.Modifiers), 2)
}

func TestGroupSimpleUnrecognisedModifiers(t *testing.T) {
	a, errs := ParseString(`
		relaxed (
			Dog {

			}
		)
	`)
	goutil.AssertNow(t, a != nil, "ast is nil")
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
	goutil.AssertNow(t, a.Declarations != nil, "nil declarations")
	goutil.AssertLength(t, a.Declarations.Length(), 1)
	n := a.Declarations.Next()
	c := n.(*ast.ClassDeclarationNode)
	goutil.AssertLength(t, len(c.Modifiers.Modifiers), 0)
}

func TestGroupedClassMultiLevelRecognisedModifiers(t *testing.T) {
	a, errs := ParseString(`
		public static (
			class Dog {
				name string
			}

			class Cat {
				name string
			}
		)
	`)
	goutil.AssertNow(t, a != nil, "ast is nil")
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
	goutil.AssertNow(t, a.Declarations != nil, "nil declarations")
	goutil.AssertLength(t, a.Declarations.Length(), 2)
	n := a.Declarations.Next()
	c := n.(*ast.ClassDeclarationNode)
	goutil.AssertLength(t, len(c.Modifiers.Modifiers), 2)
}
