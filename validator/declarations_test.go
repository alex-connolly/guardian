package validator

import (
	"testing"

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
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 0, errs.Format())

}

func TestValidateEventDeclValidSingle(t *testing.T) {
	scope, _ := parser.ParseString("event Dog(a int)")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 0, errs.Format())

}

func TestValidateEventDeclValidMultiple(t *testing.T) {
	scope, _ := parser.ParseString("event Dog(a int, b string)")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateEventDeclInvalidSingle(t *testing.T) {
	scope, _ := parser.ParseString("event Dog(c Cat)")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateEventDeclInvalidMultiple(t *testing.T) {
	scope, _ := parser.ParseString("event Dog(c Cat, a Animal)")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 2, errs.Format())

}

func TestValidateEventDeclMixed(t *testing.T) {
	scope, _ := parser.ParseString("event Dog(a int, b Cat)")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateFuncDeclEmpty(t *testing.T) {
	scope, _ := parser.ParseString("func Dog() {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateFuncDeclValidSingle(t *testing.T) {
	scope, _ := parser.ParseString("func Dog(a int) {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateFuncDeclValidMultiple(t *testing.T) {
	scope, _ := parser.ParseString("func Dog(a int, b string) {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateFuncDeclInvalidSingle(t *testing.T) {
	scope, _ := parser.ParseString("func dog(a Cat) {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateFuncDeclInvalidMultiple(t *testing.T) {
	scope, _ := parser.ParseString("func Dog(a Cat, b Animal) {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 2, errs.Format())

}

func TestValidateFuncDeclMixed(t *testing.T) {
	scope, _ := parser.ParseString("func Dog(a int, b Cat) {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateConstructorDeclEmpty(t *testing.T) {
	scope, _ := parser.ParseString("constructor Dog() {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateConstructorDeclValidSingle(t *testing.T) {
	scope, _ := parser.ParseString("constructor Dog(a int) {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateConstructorDeclValidMultiple(t *testing.T) {
	scope, _ := parser.ParseString("constructor Dog(a int, b string) {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateConstructorDeclInvalidSingle(t *testing.T) {
	scope, _ := parser.ParseString("constructor(a Cat) {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateConstructorDeclInvalidMultiple(t *testing.T) {
	scope, _ := parser.ParseString("constructor(a Cat, b Animal) {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 2, errs.Format())

}

func TestValidateConstructorDeclMixed(t *testing.T) {
	scope, _ := parser.ParseString("constructor(a int, b Cat) {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateContractDeclEmpty(t *testing.T) {
	scope, _ := parser.ParseString("contract Dog {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateContractDeclValidSingle(t *testing.T) {
	scope, _ := parser.ParseString("contract Canine{} contract Dog inherits Canine {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateContractDeclValidMultiple(t *testing.T) {
	scope, _ := parser.ParseString("contract Canine {} contract Animal {} contract Dog inherits Canine, Animal {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateContractDeclInvalidSingle(t *testing.T) {
	scope, _ := parser.ParseString("contract Dog inherits Canine {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateContractDeclInvalidMultiple(t *testing.T) {
	scope, _ := parser.ParseString("contract Dog inherits Canine, Animal {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 2, errs.Format())

}

func TestValidateContractDeclMixed(t *testing.T) {
	scope, _ := parser.ParseString("contract Canine{} contract Dog inherits Canine, Animal {}")
	goutil.AssertNow(t, scope != nil, "scope should not be nil")
	goutil.AssertNow(t, scope.Declarations != nil, "declarations shouldn't be nil")
	errs := Validate(scope, NewTestVM())
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}
