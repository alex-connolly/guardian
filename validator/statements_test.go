package validator

import (
	"fmt"
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
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateAssignmentToFuncValid(t *testing.T) {

	scope, _ := parser.ParseString(`
			func x() int8 {
				return 3
			}
			a = 0
			a = 5
			a = 5 + 6
			a = x()
		`)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
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
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateAssignmentToFuncLiteralValid(t *testing.T) {

	scope, _ := parser.ParseString(`
			var x func(int, int) string
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
			a = func(c int, b int) string {
				return "hello"
			}
			x = a
			b = func(q, c int) string {
				return "hello"
			}
			x = b
		`)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateAssignmentMultipleLeft(t *testing.T) {

	scope, _ := parser.ParseString(`
			a = 0
			b = 5
			a, b = 1, 2
			a, b = 2
		`)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateAssignmentMultipleLeftMixedTuple(t *testing.T) {

	scope, _ := parser.ParseString(`
			func x() (int8, int8){
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
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateAssignmentInvalid(t *testing.T) {
	scope, _ := parser.ParseString(`
			a = 0
			a = "hello world"
			a = 5 > 6
		`)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 2, errs.Format())
}

func TestValidateForStatementValidCond(t *testing.T) {
	scope, _ := parser.ParseString("for a = 0; a < 5 {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(scope.Sequence) == 1, "wrong sequence length")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateForStatementInvalidCond(t *testing.T) {
	scope, _ := parser.ParseString("for a = 0; a + 5 {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(scope.Sequence) == 1, "wrong sequence length")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

/*
func TestValidateIfStatementValidInit(t *testing.T) {
	scope, _ := parser.ParseString("if x = 0; x < 5 {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(scope.Sequence) == 1, "wrong sequence length")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}*/

func TestValidateIfStatementValidCondition(t *testing.T) {
	scope, _ := parser.ParseString("if true {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(scope.Sequence) == 1, "wrong sequence length")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateIfStatementValidElse(t *testing.T) {
	scope, _ := parser.ParseString(`
		if true {

		} else {

		}
	`)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(scope.Sequence) == 1, "wrong sequence length")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateSwitchStatementValidEmpty(t *testing.T) {
	scope, _ := parser.ParseString(`
		x = 5
		switch x {

		}
	`)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(scope.Sequence) == 2, "wrong sequence length")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateSwitchStatementValidCases(t *testing.T) {
	scope, _ := parser.ParseString(`
		x = 5
		switch x {
		case 4:
		case 3:

		}
	`)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(scope.Sequence) == 2, "wrong sequence length")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateClassAssignmentStatement(t *testing.T) {
	scope, _ := parser.ParseString(`
		class Dog {
			var name string
		}

		d = Dog{
			name: "Fido",
		}

		x = d.name
	`)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(scope.Sequence) == 2, "wrong sequence length")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateSuperClassAssignmentStatement(t *testing.T) {
	scope, _ := parser.ParseString(`
		class Animal {
			var legs int
		}

		class Dog inherits Animal {
			var name string
		}

		d = Dog{
			name: "Fido",
		}

		x = d.name
		y = d.legs
	`)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(scope.Sequence) == 3, "wrong sequence length")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateInterfaceAssignmentStatement(t *testing.T) {
	scope, _ := parser.ParseString(`
		interface Dog {
			name() string
		}
		var x Dog

		y = x.name()
	`)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(scope.Sequence) == 1, "wrong sequence length")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())

}

func TestValidateInterfaceWrongTypeAssignmentStatement(t *testing.T) {
	scope, _ := parser.ParseString(`
		interface Dog {
			name()
		}
		var x Dog

		y = x.name()
	`)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(scope.Sequence) == 1, "wrong sequence length")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateClassAssignmentStatementInvalid(t *testing.T) {
	scope, _ := parser.ParseString(`
		class Dog {
			var name string
		}

		d = Dog {
			name: "Fido",
		}

		x = d.wrongName
	`)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(scope.Sequence) == 2, "wrong sequence length")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateForEachStatementValid(t *testing.T) {
	scope, _ := parser.ParseString(`
		a = []string{"a", "b"}
		for x, y in a {

		}
	`)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(scope.Sequence) == 2, fmt.Sprintf("wrong sequence length: %d", len(scope.Sequence)))
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestTripleReferenceIdentifiers(t *testing.T) {
	scope, _ := parser.ParseString(`
		class A {
			var b string
		}
		class C {
			var a A
		}
		var x string
		var c C
		x = c.a.b
	`)
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestTripleReferenceCalls(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class B {
			func c() string {
				return "hi"
			}
		}
		class A {
			func b() B {
				return B{}
			}
		}
		var x string
		var a A
		x = a.b().c()
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestPrivateVariableAccess(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class A {
			private var all string
		}
		var min A
		x = min.all
	`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestPrivateVariableAccessFromFunction(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class A {
			private var a string
		}
		class B inherits A {
			func getA() string {
				return a
			}
		}
	`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestProtectedVariableAccessFromFunction(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class A {
			protected var a string
		}
		class B inherits A {
			func getA() string {
				return a
			}
		}
		class C {
			var a A
			func getA() string {
				return a.a
			}
		}
	`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestFunctionSingleReturnTypeValid(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		func a() string {
			return "hi"
		}
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestFunctionSingleReturnTypeInvalid(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		func a() string {
			return 6
		}
	`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestDanglingReturn(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		return 6
	`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestDuplicatePackageStatement(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		package dog guardian 0.0.1
		package car guardian 0.0.1
	`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestImportBeforePackage(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		import "dog"
		package car guardian 0.0.1
	`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestDividedImports(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		package car varsion 0.0.1
		import "dog"
		x = 1
		import "cat"
	`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestLastSlash(t *testing.T) {
	path := "dog/a"
	goutil.AssertNow(t, trimPath(path) == "a", "wrong path")
	path = "x/dog/a"
	goutil.AssertNow(t, trimPath(path) == "a", "wrong path")
}

func TestLiteralValidAssignment(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		func main(){
			var a int
			a = 5
			var b uint
			b = 5
			var c int8
			c = 5
			var d uint8
			d = 5
		}
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestGenericCallExpression(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		func main(){
			var a []int
			a = append(a, 0)
		}
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}
