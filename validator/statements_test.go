package validator

import (
	"testing"

	"github.com/end-r/guardian/parser"

	"github.com/end-r/goutil"
)

func TestValidateAssignmentValid(t *testing.T) {

	scope, _ := parser.ParseString(`
			a = 0
			a = 5
			a = 5 + 6
		`)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(scope)
	goutil.AssertNow(t, len(errs) == 0, errs.format())
}

func TestValidateAssignmentToFuncValid(t *testing.T) {

	scope, _ := parser.ParseString(`
			func x() int {
				return 3
			}
			a = 0
			a = 5
			a = 5 + 6
			a = x()
		`)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(scope)
	goutil.AssertNow(t, len(errs) == 0, errs.format())
}

func TestValidateAssignmentToFuncInvalid(t *testing.T) {

	scope, _ := parser.ParseString(`
			func x() string {
				return "hi"
			}
			a = 0
			a = x()
		`)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(scope)
	goutil.AssertNow(t, len(errs) == 1, errs.format())
}

func TestValidateAssignmentToFuncLiteralValid(t *testing.T) {

	scope, _ := parser.ParseString(`
			x func(int, int) string
			x = func(a int, b int) string {
				return "hello"
			}
			x = func(a, b int) string {
				return "hello"
			}
			func y(a int, b int) string {
				return "hello"
			}
			x = y
			func z(a, b int) string {
				return "hello"
			}
			x = z
			a = func(a int, b int) string {
				return "hello"
			}
			x = a
			b = func(a, b int) string {
				return "hello"
			}
			x = b
		`)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(scope)
	goutil.AssertNow(t, len(errs) == 0, errs.format())
}

func TestValidateAssignmentMultipleLeft(t *testing.T) {

	scope, _ := parser.ParseString(`
			a = 0
			b = 5
			a, b = 1, 2
			a, b = 2
		`)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(scope)
	goutil.AssertNow(t, len(errs) == 0, errs.format())
}

func TestValidateAssignmentMultipleLeftMixedTuple(t *testing.T) {

	scope, _ := parser.ParseString(`
			func x() (int, int){
				return 0, 1
			}
			a = 0
			b = 5
			c = 2
			d = 3
			a, b = x()
			c, a, b, d = x(), x()
		`)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(scope)
	goutil.AssertNow(t, len(errs) == 0, errs.format())
}

func TestValidateAssignmentInvalid(t *testing.T) {
	scope, _ := parser.ParseString(`
			a = 0
			a = "hello world"
			a = 5 > 6
		`)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(scope)
	goutil.AssertNow(t, len(errs) == 2, errs.format())
}

func TestValidateForStatementValidCond(t *testing.T) {
	scope, _ := parser.ParseString("for a = 0; a < 5 {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(scope.Sequence) == 1, "wrong sequence length")
	errs := Validate(scope)
	goutil.AssertNow(t, len(errs) == 0, errs.format())
}

func TestValidateForStatementInvalidCond(t *testing.T) {
	scope, _ := parser.ParseString("for a = 0; a + 5 {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(scope.Sequence) == 1, "wrong sequence length")
	errs := Validate(scope)
	goutil.AssertNow(t, len(errs) == 1, errs.format())
}

/*
func TestValidateIfStatementValidInit(t *testing.T) {
	scope, _ := parser.ParseString("if x = 0; x < 5 {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(scope.Sequence) == 1, "wrong sequence length")
	errs := Validate(scope)
	goutil.AssertNow(t, len(errs) == 0, errs.format())
}*/

func TestValidateIfStatementValidCondition(t *testing.T) {
	scope, _ := parser.ParseString("if true {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(scope.Sequence) == 1, "wrong sequence length")
	errs := Validate(scope)
	goutil.AssertNow(t, len(errs) == 0, errs.format())
}

func TestValidateIfStatementValidElse(t *testing.T) {
	scope, _ := parser.ParseString(`
		if true {

		} else {

		}
	`)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(scope.Sequence) == 1, "wrong sequence length")
	errs := Validate(scope)
	goutil.AssertNow(t, len(errs) == 0, errs.format())
}

func TestValidateSwitchStatementValidEmpty(t *testing.T) {
	scope, _ := parser.ParseString(`
		x = 5
		switch x {

		}
	`)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(scope.Sequence) == 2, "wrong sequence length")
	errs := Validate(scope)
	goutil.AssertNow(t, len(errs) == 0, errs.format())
}

func TestValidateSwitchStatementValidCases(t *testing.T) {
	scope, _ := parser.ParseString(`
		x = 5
		switch x {
			case 4 {

			}
			case 3 {

			}
		}
	`)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(scope.Sequence) == 2, "wrong sequence length")
	errs := Validate(scope)
	goutil.AssertNow(t, len(errs) == 0, errs.format())
}

func TestValidateClassAssignmentStatement(t *testing.T) {
	scope, _ := parser.ParseString(`
		class Dog {
			name string
		}

		d = Dog{
			name: "Fido",
		}

		x = d.name
	`)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(scope.Sequence) == 2, "wrong sequence length")
	errs := Validate(scope)
	goutil.AssertNow(t, len(errs) == 0, errs.format())
}

func TestValidateClassAssignmentStatementInvalid(t *testing.T) {
	scope, _ := parser.ParseString(`
		class Dog {
			name string
		}

		d = Dog{
			name: "Fido",
		}

		x = d.wrongName
	`)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(scope.Sequence) == 2, "wrong sequence length")
	errs := Validate(scope)
	goutil.AssertNow(t, len(errs) == 1, errs.format())
}
