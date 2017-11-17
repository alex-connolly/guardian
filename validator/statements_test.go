package validator

import (
	"testing"

	"github.com/end-r/guardian/compiler/parser"

	"github.com/end-r/goutil"
)

func TestValidateAssignmentValid(t *testing.T) {

	p := parser.ParseString(`
			a = 0
			a = 5
			a = 5 + 6
		`)
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestValidateAssignmentToFuncValid(t *testing.T) {

	p := parser.ParseString(`
			func x() int {
				return 3
			}
			a = 0
			a = 5
			a = 5 + 6
			a = x()
		`)
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestValidateAssignmentToFuncInvalid(t *testing.T) {

	p := parser.ParseString(`
			func x() string {
				return "hi"
			}
			a = 0
			a = x()
		`)
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, v.formatErrors())
}

func TestValidateAssignmentToFuncLiteralValid(t *testing.T) {

	p := parser.ParseString(`
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
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestValidateAssignmentMultipleLeft(t *testing.T) {

	p := parser.ParseString(`
			a = 0
			b = 5
			a, b = 1, 2
			a, b = 2
		`)
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestValidateAssignmentMultipleLeftMixedTuple(t *testing.T) {

	p := parser.ParseString(`
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
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestValidateAssignmentInvalid(t *testing.T) {
	p := parser.ParseString(`
			a = 0
			a = "hello world"
			a = 5 > 6
		`)
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 2, v.formatErrors())
}

func TestValidateForStatementValidCond(t *testing.T) {
	p := parser.ParseString("for a = 0; a < 5 {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(p.Scope.Sequence) == 1, "wrong sequence length")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestValidateForStatementInvalidCond(t *testing.T) {
	p := parser.ParseString("for a = 0; a + 5 {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(p.Scope.Sequence) == 1, "wrong sequence length")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, v.formatErrors())
}

/*
func TestValidateIfStatementValidInit(t *testing.T) {
	p := parser.ParseString("if x = 0; x < 5 {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(p.Scope.Sequence) == 1, "wrong sequence length")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}*/

func TestValidateIfStatementValidCondition(t *testing.T) {
	p := parser.ParseString("if true {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(p.Scope.Sequence) == 1, "wrong sequence length")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestValidateIfStatementValidElse(t *testing.T) {
	p := parser.ParseString(`
		if true {

		} else {

		}
	`)
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(p.Scope.Sequence) == 1, "wrong sequence length")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestValidateSwitchStatementValidEmpty(t *testing.T) {
	p := parser.ParseString(`
		x = 5
		switch x {

		}
	`)
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(p.Scope.Sequence) == 2, "wrong sequence length")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestValidateSwitchStatementValidCases(t *testing.T) {
	p := parser.ParseString(`
		x = 5
		switch x {
			case 4 {

			}
			case 3 {

			}
		}
	`)
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(p.Scope.Sequence) == 2, "wrong sequence length")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestValidateClassAssignmentStatement(t *testing.T) {
	p := parser.ParseString(`
		class Dog {
			name string
		}

		d = Dog{
			name: "Fido",
		}

		x = d.name
	`)
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(p.Scope.Sequence) == 2, "wrong sequence length")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestValidateClassAssignmentStatementInvalid(t *testing.T) {
	p := parser.ParseString(`
		class Dog {
			name string
		}

		d = Dog{
			name: "Fido",
		}

		x = d.wrongName
	`)
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(p.Scope.Sequence) == 2, "wrong sequence length")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, v.formatErrors())
}
