package validator

import (
	"fmt"
	"testing"

	"github.com/end-r/goutil"
	"github.com/end-r/guardian/compiler/parser"
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
	goutil.AssertNow(t, len(v.errors) == 0, fmt.Sprintf("wrong err length: %s", len(v.errors)))
}

func TestValidateEventDeclValidSingle(t *testing.T) {
	p := parser.ParseString("event Dog(int)")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, fmt.Sprintf("wrong err length: %s", len(v.errors)))
}

func TestValidateEventDeclValidMultiple(t *testing.T) {
	p := parser.ParseString("event Dog(int, string)")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, fmt.Sprintf("wrong err length: %s", len(v.errors)))
}

func TestValidateEventDeclInvalidSingle(t *testing.T) {
	p := parser.ParseString("event Dog(Cat)")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, fmt.Sprintf("wrong err length: %s", len(v.errors)))
}

func TestValidateEventDeclInvalidMultiple(t *testing.T) {
	p := parser.ParseString("event Dog(Cat, Animal)")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 2, fmt.Sprintf("wrong err length: %s", len(v.errors)))

}

func TestValidateEventDeclMixed(t *testing.T) {
	p := parser.ParseString("event Dog(int, Cat)")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, fmt.Sprintf("wrong err length: %s", len(v.errors)))
}

func TestValidateFuncDeclEmpty(t *testing.T) {
	p := parser.ParseString("func Dog() {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, fmt.Sprintf("wrong err length: %s", len(v.errors)))
}

func TestValidateFuncDeclValidSingle(t *testing.T) {
	p := parser.ParseString("func Dog(a int) {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, fmt.Sprintf("wrong err length: %s", len(v.errors)))
}

func TestValidateFuncDeclValidMultiple(t *testing.T) {
	p := parser.ParseString("func Dog(a int, b string) {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, fmt.Sprintf("wrong err length: %s", len(v.errors)))
}

func TestValidateFuncDeclInvalidSingle(t *testing.T) {
	p := parser.ParseString("func Dog(a Cat) {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, fmt.Sprintf("wrong err length: %s", len(v.errors)))
}

func TestValidateFuncDeclInvalidMultiple(t *testing.T) {
	p := parser.ParseString("func Dog(a Cat, b Animal) {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 2, fmt.Sprintf("wrong err length: %s", len(v.errors)))

}

func TestValidateFuncDeclMixed(t *testing.T) {
	p := parser.ParseString("func Dog(a int, b Cat) {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, fmt.Sprintf("wrong err length: %s", len(v.errors)))
}

func TestValidateConstructorDeclEmpty(t *testing.T) {
	p := parser.ParseString("constructor Dog() {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, fmt.Sprintf("wrong err length: %s", len(v.errors)))
}

func TestValidateConstructorDeclValidSingle(t *testing.T) {
	p := parser.ParseString("constructor Dog(a int) {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, fmt.Sprintf("wrong err length: %s", len(v.errors)))
}

func TestValidateConstructorDeclValidMultiple(t *testing.T) {
	p := parser.ParseString("constructor Dog(a int, b string) {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, fmt.Sprintf("wrong err length: %s", len(v.errors)))
}

func TestValidateConstructorDeclInvalidSingle(t *testing.T) {
	p := parser.ParseString("constructor Dog(a Cat) {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, fmt.Sprintf("wrong err length: %s", len(v.errors)))
}

func TestValidateConstructorDeclInvalidMultiple(t *testing.T) {
	p := parser.ParseString("constructor Dog(a Cat, b Animal) {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 2, fmt.Sprintf("wrong err length: %s", len(v.errors)))

}

func TestValidateConstructorDeclMixed(t *testing.T) {
	p := parser.ParseString("constructor Dog(a int, b Cat) {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, fmt.Sprintf("wrong err length: %s", len(v.errors)))
}

func TestValidateContractDeclEmpty(t *testing.T) {
	p := parser.ParseString("contract Dog {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, fmt.Sprintf("wrong err length: %s", len(v.errors)))
}

func TestValidateContractDeclValidSingle(t *testing.T) {
	p := parser.ParseString("contract Canine{} contract Dog inherits Canine {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, fmt.Sprintf("wrong err length: %s", len(v.errors)))
}

func TestValidateContractDeclValidMultiple(t *testing.T) {
	p := parser.ParseString("contract Canine {} contract Animal {} contract Dog inherits Canine, Animal {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, fmt.Sprintf("wrong err length: %s", len(v.errors)))
}

func TestValidateContractDeclInvalidSingle(t *testing.T) {
	p := parser.ParseString("contract Dog inherits Canine {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, fmt.Sprintf("wrong err length: %s", len(v.errors)))
}

func TestValidateContractDeclInvalidMultiple(t *testing.T) {
	p := parser.ParseString("contract Dog inherits Canine, Animal {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 2, fmt.Sprintf("wrong err length: %s", len(v.errors)))

}

func TestValidateContractDeclMixed(t *testing.T) {
	p := parser.ParseString("contract Canine{} contract Dog inherits Canine, Animal {}")
	goutil.AssertNow(t, p.Scope != nil, "scope should not be nil")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, fmt.Sprintf("wrong err length: %s", len(v.errors)))
}
