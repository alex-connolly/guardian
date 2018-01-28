package validator

import (
	"testing"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/typing"

	"github.com/end-r/guardian/parser"

	"github.com/end-r/goutil"
)

func TestValidateClassDecl(t *testing.T) {

}

func TestValidateInterfaceDecl(t *testing.T) {

}

func TestValidateEnumDecl(t *testing.T) {

}

func TestValidateEventDeclEmpty(t *testing.T) {
	scope, _ := parser.ParseString("event Dog()")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())

}

func TestValidateEventDeclValidSingle(t *testing.T) {
	scope, _ := parser.ParseString("event Dog(a int)")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())

}

func TestValidateEventDeclValidMultiple(t *testing.T) {
	scope, _ := parser.ParseString("event Dog(a int, b string)")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateEventDeclInvalidSingle(t *testing.T) {
	scope, _ := parser.ParseString("event Dog(c Cat)")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateEventDeclInvalidMultiple(t *testing.T) {
	scope, _ := parser.ParseString("event Dog(c Cat, a Animal)")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 2, errs.Format())

}

func TestValidateEventDeclMixed(t *testing.T) {
	scope, _ := parser.ParseString("event Dog(a int, b Cat)")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateFuncDeclEmpty(t *testing.T) {
	scope, _ := parser.ParseString("func Dog() {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateFuncDeclValidSingle(t *testing.T) {
	scope, _ := parser.ParseString("func Dog(a int) {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateFuncDeclValidMultiple(t *testing.T) {
	scope, _ := parser.ParseString("func Dog(a int, b string) {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateFuncDeclInvalidSingle(t *testing.T) {
	scope, _ := parser.ParseString("func dog(a Cat) {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateFuncDeclInvalidMultiple(t *testing.T) {
	scope, _ := parser.ParseString("func Dog(a Cat, b Animal) {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 2, errs.Format())

}

func TestValidateFuncDeclMixed(t *testing.T) {
	scope, _ := parser.ParseString("func Dog(a int, b Cat) {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateConstructorDeclEmpty(t *testing.T) {
	scope, _ := parser.ParseString("constructor() {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateConstructorDeclValidSingle(t *testing.T) {
	scope, _ := parser.ParseString("constructor(a int) {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateConstructorDeclValidMultiple(t *testing.T) {
	scope, _ := parser.ParseString("constructor(a int, b string) {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateConstructorDeclInvalidSingle(t *testing.T) {
	scope, _ := parser.ParseString("constructor(a Cat) {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateConstructorDeclInvalidMultiple(t *testing.T) {
	scope, _ := parser.ParseString("constructor(a Cat, b Animal) {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 2, errs.Format())

}

func TestValidateConstructorDeclMixed(t *testing.T) {
	scope, _ := parser.ParseString("constructor(a int, b Cat) {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateContractDeclEmpty(t *testing.T) {
	scope, _ := parser.ParseString("contract Dog {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateContractDeclValidSingle(t *testing.T) {
	scope, _ := parser.ParseString("contract Canine{} contract Dog inherits Canine {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateContractDeclValidMultiple(t *testing.T) {
	scope, _ := parser.ParseString("contract Canine {} contract Animal {} contract Dog inherits Canine, Animal {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateContractDeclInvalidSingle(t *testing.T) {
	scope, _ := parser.ParseString("contract Dog inherits Canine {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateContractDeclInvalidMultiple(t *testing.T) {
	scope, _ := parser.ParseString("contract Dog inherits Canine, Animal {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 2, errs.Format())

}

func TestValidateContractDeclMixed(t *testing.T) {
	scope, _ := parser.ParseString("contract Canine{} contract Dog inherits Canine, Animal {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateExplicitVarDecl(t *testing.T) {
	scope, _ := parser.ParseString("var hi uint8")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(NewTestVM(), scope, nil)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateModifiersValidAccess(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), "public var name string")
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateModifiersDuplicateAccess(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), "public public var name string")
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateModifiersMEAccess(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), "public private var name string")
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateModifiersInvalidNodeType(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), "test var name string")
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateModifiersInvalidUnrecognised(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), "elephant var name string")
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestInterfaceParents(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		interface Switchable{}
		interface Deletable{}
		interface Light inherits Switchable, Deletable {}
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestInterfaceParentsImplemented(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		interface Switchable{
			on()
			off()
		}
		interface Deletable{}
		interface Light inherits Switchable, Deletable {

		}
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestMapTypeValid(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		var a map[int]int
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestMapTypeInvalid(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		var a map[zoop]doop
	`)
	goutil.AssertNow(t, len(errs) == 2, errs.Format())
}

func TestEnumValid(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		enum Weekday {
			Mon, Tue, Wed, Thurs, Fri
		}
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestEnumInheritanceValid(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		enum Weekday {
			Mon, Tue, Wed, Thurs, Fri
		}

		enum Weekend {
			Sat, Sun
		}

		enum Day inherits Weekday, Weekend {

		}
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestEnumInheritanceInvalid(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		enum Weekday {
			Mon, Tue, Wed, Thurs, Fri
		}

		interface Weekend {

		}

		enum Day inherits Weekday, Weekend {

		}
	`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestClassInheritanceValid(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class Dog {}
		class Mastiff inherits Dog {}
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestClassInheritanceInvalid(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		interface Dog {}
		class Mastiff inherits Dog {}
	`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestClassImplementationValid(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		interface Dog {}
		class Mastiff is Dog {}
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestClassImplementationInvalid(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class Dog {}
		class Mastiff is Dog {}
	`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestContractImplementationValid(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		interface Dog {}
		contract Mastiff is Dog {}
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestContractImplementationInvalidWrongType(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		contract Dog {}
		contract Mastiff is Dog {}
	`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestClassImplementationInvalidWrongType(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class Dog {}
		class Mastiff is Dog {}
	`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestContractImplementationInvalidMissingMethods(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		interface Dog {
			woof()
		}
		contract Mastiff is Dog {}
	`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestClassImplementationInvalidMissingMethods(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		interface Dog {
			woof()
		}
		class Mastiff is Dog {}
	`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestContractImplementationInvalidWrongMethodParameters(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		interface Dog {
			woof()
		}
		contract Mastiff is Dog {
			func woof(m string){

			}
		}
	`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestClassImplementationInvalidWrongMethodParameters(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		interface Dog {
			woof()
		}
		class Mastiff is Dog {
			func woof(m string){

			}
		}
	`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestContractImplementationValidThroughParent(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		interface Dog {
			woof()
		}

		contract X {
			func woof(){

			}
		}

		contract Mastiff is Dog inherits X {

		}
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestClassImplementationValidThroughParent(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		interface Dog {
			woof()
		}

		class X {
			func woof(){

			}
		}

		class Mastiff is Dog inherits X {

		}
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestContractImplementationInvalidMissingParentMethods(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		interface Dog {
			woof()
		}

		interface Cat inherits Dog {

		}

		contract Mastiff is Cat {}
	`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestClassImplementationInvalidMissingParentMethods(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		interface Dog {
			woof()
		}

		interface Cat inherits Dog {

		}

		class Mastiff is Cat {}
	`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestInterfaceMethods(t *testing.T) {
	a, errs := ValidateString(NewTestVM(), `
		interface Calculator {
			add(a, b int) int
			sub(a, b int) int
		}
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
	goutil.AssertNow(t, a.Declarations != nil, "nil declarations")
	n := a.Declarations.Next().(*ast.InterfaceDeclarationNode)
	i := n.Resolved.(*typing.Interface)
	goutil.AssertLength(t, len(i.Funcs), 2)
	add := i.Funcs["add"]
	goutil.AssertNow(t, add != nil, "add is nil")
	sub := i.Funcs["sub"]
	goutil.AssertNow(t, sub != nil, "sub is nil")
}

func TestCancellationEnums(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		enum Weekday { Mon, Tue, Wed, Thu, Fri }
		enum WeekdayTwo { Mon, Tue, Wed, Thu, Fri }
		enum Cancelled inherits Weekday, WeekdayTwo {}
		x = Cancelled.Mon
	`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestCancellationClass(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class Lion { var name string }
		class Tiger { var name string }
		class Liger inherits Lion, Tiger {
			func getName() {
				x = name
			}
		}
	`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestCancellationContract(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		contract Lion { var name string }
		contract Tiger { var name string }
		contract Liger inherits Lion, Tiger {
			func getName() {
				x = name
			}
		}
	`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestFuncDeclarationSingleReturn(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		func hi() string {
			return "hi"
		}
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestFuncDeclarationMultipleReturn(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		func hi() (int, int) {
			return 6, 6
		}
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestFuncDeclarationInvalidMultipleReturn(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		func hi() (int, int) {
			return 6, "hi"
		}
	`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestFuncDeclarationVoidEmptyReturn(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		func hi() {
			return
		}
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestFuncDeclarationVoidSingleInvalidReturn(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		func hi() {
			return a
		}
	`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestAccessBaseContractProperty(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		contract A {
			var b uint256
			constructor(){
				this.balance = this.b
			}
		}
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}
