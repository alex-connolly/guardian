package parser

import (
	"testing"

	"github.com/end-r/guardian/ast"

	"github.com/end-r/goutil"
)

func TestClassSimpleModifiers(t *testing.T) {
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

func TestGroupedClassSimpleModifiers(t *testing.T) {
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

func TestGroupedClassMultiLevelModifiers(t *testing.T) {
	a, errs := ParseString(`
		public static (
			class Dog {
				var name string
			}

			class Cat {
				var name string
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
