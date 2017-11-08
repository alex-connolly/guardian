package validator

import (
	"testing"

	"github.com/end-r/guardian/compiler/parser"

	"github.com/end-r/goutil"
)

func TestValidateClassDecl(t *testing.T) {

}

func TestValidateInterfaceDecl(t *testing.T) {

}

func TestValidateEnumDecl(t *testing.T) {

}

func TestValidateEventDeclEmpty(t *testing.T) {
	p := parser.ParseString("event Dog()")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())

}

func TestValidateEventDeclValidSingle(t *testing.T) {
	p := parser.ParseString("event Dog(a int)")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())

}

func TestValidateEventDeclValidMultiple(t *testing.T) {
	p := parser.ParseString("event Dog(a int, b string)")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestValidateEventDeclInvalidSingle(t *testing.T) {
	p := parser.ParseString("event Dog(c Cat)")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, v.formatErrors())
}

func TestValidateEventDeclInvalidMultiple(t *testing.T) {
	p := parser.ParseString("event Dog(c Cat, a Animal)")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 2, v.formatErrors())

}

func TestValidateEventDeclMixed(t *testing.T) {
	p := parser.ParseString("event Dog(a int, b Cat)")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, v.formatErrors())
}

func TestValidateFuncDeclEmpty(t *testing.T) {
	p := parser.ParseString("func Dog() {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, p.Scope.Declarations != nil, "declarations shouldn't be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestValidateFuncDeclValidSingle(t *testing.T) {
	p := parser.ParseString("func Dog(a int) {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, p.Scope.Declarations != nil, "declarations shouldn't be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestValidateFuncDeclValidMultiple(t *testing.T) {
	p := parser.ParseString("func Dog(a int, b string) {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, p.Scope.Declarations != nil, "declarations shouldn't be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestValidateFuncDeclInvalidSingle(t *testing.T) {
	p := parser.ParseString("func dog(a Cat) {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, p.Scope.Declarations != nil, "declarations shouldn't be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, v.formatErrors())
}

func TestValidateFuncDeclInvalidMultiple(t *testing.T) {
	p := parser.ParseString("func Dog(a Cat, b Animal) {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, p.Scope.Declarations != nil, "declarations shouldn't be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 2, v.formatErrors())

}

func TestValidateFuncDeclMixed(t *testing.T) {
	p := parser.ParseString("func Dog(a int, b Cat) {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, p.Scope.Declarations != nil, "declarations shouldn't be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, v.formatErrors())
}

func TestValidateConstructorDeclEmpty(t *testing.T) {
	p := parser.ParseString("constructor Dog() {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, p.Scope.Declarations != nil, "declarations shouldn't be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestValidateConstructorDeclValidSingle(t *testing.T) {
	p := parser.ParseString("constructor Dog(a int) {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, p.Scope.Declarations != nil, "declarations shouldn't be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestValidateConstructorDeclValidMultiple(t *testing.T) {
	p := parser.ParseString("constructor Dog(a int, b string) {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, p.Scope.Declarations != nil, "declarations shouldn't be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestValidateConstructorDeclInvalidSingle(t *testing.T) {
	p := parser.ParseString("constructor(a Cat) {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, p.Scope.Declarations != nil, "declarations shouldn't be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, v.formatErrors())
}

func TestValidateConstructorDeclInvalidMultiple(t *testing.T) {
	p := parser.ParseString("constructor(a Cat, b Animal) {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, p.Scope.Declarations != nil, "declarations shouldn't be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 2, v.formatErrors())

}

func TestValidateConstructorDeclMixed(t *testing.T) {
	p := parser.ParseString("constructor(a int, b Cat) {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, p.Scope.Declarations != nil, "declarations shouldn't be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, v.formatErrors())
}

func TestValidateContractDeclEmpty(t *testing.T) {
	p := parser.ParseString("contract Dog {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, p.Scope.Declarations != nil, "declarations shouldn't be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestValidateContractDeclValidSingle(t *testing.T) {
	p := parser.ParseString("contract Canine{} contract Dog inherits Canine {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, p.Scope.Declarations != nil, "declarations shouldn't be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestValidateContractDeclValidMultiple(t *testing.T) {
	p := parser.ParseString("contract Canine {} contract Animal {} contract Dog inherits Canine, Animal {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, p.Scope.Declarations != nil, "declarations shouldn't be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, v.formatErrors())
}

func TestValidateContractDeclInvalidSingle(t *testing.T) {
	p := parser.ParseString("contract Dog inherits Canine {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, p.Scope.Declarations != nil, "declarations shouldn't be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, v.formatErrors())
}

func TestValidateContractDeclInvalidMultiple(t *testing.T) {
	p := parser.ParseString("contract Dog inherits Canine, Animal {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, p.Scope.Declarations != nil, "declarations shouldn't be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 2, v.formatErrors())

}

func TestValidateContractDeclMixed(t *testing.T) {
	p := parser.ParseString("contract Canine{} contract Dog inherits Canine, Animal {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	goutil.AssertNow(t, p.Scope.Declarations != nil, "declarations shouldn't be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, v.formatErrors())
}
