package validator

import (
	"testing"

	"github.com/end-r/guardian/compiler/parser"

	"github.com/end-r/goutil"
)

func TestCallExpressionValid(t *testing.T) {
	p := parser.ParseString(`
        func call(a, b int) int {
            if a == 0 or b == 0 {
                return 0
            }
            return call(a - 1, b - 1)
        }

        call(5, 5)
        `)
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestCallExpressionInvalid(t *testing.T) {
	p := parser.ParseString(`
        interface Open {

        }

        Open(5, 5)
        `)
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, v.formatErrors())
}

func TestCallExpressionEmptyConstructorValid(t *testing.T) {
	p := parser.ParseString(`
        class Dog {

        }

        d = Dog()
        `)
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestCallExpressionSingleArgumentConstructorValid(t *testing.T) {
	p := parser.ParseString(`
        class Dog {

            yearsOld int

            constructor(age int){
                yearsOld = age
            }
        }

        d = Dog(10)
        `)
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestCallExpressionMultipleArgumentConstructorValid(t *testing.T) {
	p := parser.ParseString(`
        class Dog {

            yearsOld int
            fullName string

            constructor(name string, age int){
                fullName = name
                yearsOld = age
            }
        }

        d = Dog("alan", 10)
        `)
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestCallExpressionConstructorInvalid(t *testing.T) {
	p := parser.ParseString(`
        class Dog {

        }

        d = Dog(6, 6)
        `)
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, v.formatErrors())
}
