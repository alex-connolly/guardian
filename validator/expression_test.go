package validator

import (
	"testing"

	"github.com/end-r/guardian/parser"

	"github.com/end-r/goutil"
)

func TestCallExpressionValid(t *testing.T) {
	scope, _ := parser.ParseString(`
        func call(a, b int8) int8 {
			if a == 0 or b == 0 {
                return 0
            }
            return call(a - 1, b - 1)
        }
    `)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestCallExpressionInvalid(t *testing.T) {
	scope, _ := parser.ParseString(`
        interface Open {

        }

        Open(5, 5)
    `)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestCallExpressionEmptyConstructorValid(t *testing.T) {
	scope, _ := parser.ParseString(`
        class Dog {

        }

        d = new Dog()
        `)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestCallExpressionSingleArgumentConstructorValid(t *testing.T) {
	scope, _ := parser.ParseString(`
        class Dog {

            var yearsOld int8

            constructor(age int8){
                yearsOld = age
            }
        }

        d = new Dog(10)
        `)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestCallExpressionMultipleArgumentConstructorValid(t *testing.T) {
	scope, _ := parser.ParseString(`
        class Dog {

            var yearsOld int8
            var fullName string

            constructor(name string, age int8){

            }
        }

        d = new Dog("alan", 10 as int8)
        `)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestCallExpressionConstructorInvalid(t *testing.T) {
	scope, _ := parser.ParseString(`
        class Dog {

        }

        d = new Dog(6, 6)
        `)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestOnlyExpressionResolved(t *testing.T) {
	_, errs := ValidateExpression(NewTestVM(), "5 < 4")
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}
