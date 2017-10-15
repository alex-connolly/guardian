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
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, fmt.Sprintf("wrong err length: %s", len(v.errors)))
}

func TestValidateEventDeclValidSingle(t *testing.T) {
	p := parser.ParseString("event Dog(int)")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, fmt.Sprintf("wrong err length: %s", len(v.errors)))
}

func TestValidateEventDeclValidMultiple(t *testing.T) {
	p := parser.ParseString("event Dog(int, string)")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 0, fmt.Sprintf("wrong err length: %s", len(v.errors)))
}

func TestValidateEventDeclInvalidSingle(t *testing.T) {
	p := parser.ParseString("event Dog(Cat)")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, fmt.Sprintf("wrong err length: %s", len(v.errors)))
}

func TestValidateEventDeclInvalidMultiple(t *testing.T) {
	p := parser.ParseString("event Dog(Cat, Animal)")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 2, fmt.Sprintf("wrong err length: %s", len(v.errors)))

}

func TestValidateEventDeclMixed(t *testing.T) {
	p := parser.ParseString("event Dog(int, Cat)")
	v := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(v.errors) == 1, fmt.Sprintf("wrong err length: %s", len(v.errors)))
}

func TestValidateFuncDecl(t *testing.T) {

}

func TestValidateConstructorDecl(t *testing.T) {

}

func TestValidateContractDecl(t *testing.T) {

}
