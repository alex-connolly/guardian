package validator

import (
	"testing"

	"github.com/end-r/guardian/parser"

	"github.com/end-r/goutil"
)

func TestCallExpressionValid(t *testing.T) {
	scope, _ := parser.ParseString(`
        func f(a, b int8) int8 {
			if a == 0 or b == 0 {
                return 0
            }
            return f(a - 1, b - 1)
        }
    `)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestCallExpressionInvalid(t *testing.T) {
	scope, _ := parser.ParseString(`
        interface Open {

        }

        Open(5, 5)
    `)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestCallExpressionEmptyConstructorValid(t *testing.T) {
	scope, _ := parser.ParseString(`
        class Dog {

        }

        d = new Dog()
        `)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(NewTestVM(), scope, nil)
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
	errs := Validate(NewTestVM(), scope, nil)
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

        d = new Dog("alan", int8(10))
        `)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestCallExpressionConstructorInvalid(t *testing.T) {
	scope, _ := parser.ParseString(`
        class Dog {

        }

        d = new Dog(6, 6)
        `)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestOnlyExpressionResolved(t *testing.T) {
	_, errs := ValidateExpression(NewTestVM(), "5 < 4")
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestMapLiteralInvalidKeysAndValues(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		var m map[string]string
		m = map[string]string {
			1: 2,
			3: 5,
		}
	`)
	goutil.AssertNow(t, len(errs) == 4, errs.Format())
}

func TestMapLiteralInvalidKeys(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		var m map[string]string
		m = map[string]string {
			1: "hi",
			3: "hi",
		}
	`)
	goutil.AssertNow(t, len(errs) == 2, errs.Format())
}

func TestMapLiteralInvalidValues(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		var m map[string]string
		m = map[string]string {
			"hi": 1,
			"hi": 3,
		}
	`)
	goutil.AssertNow(t, len(errs) == 2, errs.Format())
}

func TestMapLiteralValidValues(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		var m map[string]string
		m = map[string]string {
			"hi": "bye",
			"hi": "bye",
		}
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestArrayLiteralValidValues(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		var m []string
		m = []string { "hi", "bye" }
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}
