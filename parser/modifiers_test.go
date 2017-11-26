package parser

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestExplicitVarDeclarationModifiers(t *testing.T) {
	p := createParser("private a string")
	parseExplicitVarDeclaration(p)
	goutil.AssertNow(t, p.scope != nil, "")
	x, _ := ParseString(`
        private a string
        private b string
        `)
	goutil.AssertNow(t, x != nil, "")
}

func TestClassModifiers(t *testing.T) {
	p := createParser("private class Dog {}")
	parseClassDeclaration(p)
	goutil.AssertNow(t, p.scope != nil, "")
}

func TestInterfaceModifiers(t *testing.T) {
	p := createParser("private interface Dog {}")
	parseInterfaceDeclaration(p)
	goutil.AssertNow(t, p.scope != nil, "")
}

func TestEventModifiers(t *testing.T) {
	p := createParser("protected event Dog()")
	parseEventDeclaration(p)
	goutil.AssertNow(t, p.scope != nil, "")
}
