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
	errs := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(errs) == 0, fmt.Sprintf("wrong err length: %s", len(errs)))
}

func TestValidateEventDeclValidSingle(t *testing.T) {
	p := parser.ParseString("event Dog(int)")
	errs := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(errs) == 0, fmt.Sprintf("wrong err length: %s", len(errs)))
}

func TestValidateEventDeclValidMultiple(t *testing.T) {
	p := parser.ParseString("event Dog(int, string)")
	errs := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(errs) == 0, fmt.Sprintf("wrong err length: %s", len(errs)))
}

func TestValidateEventDeclInvalidSingle(t *testing.T) {
	p := parser.ParseString("event Dog(Cat)")
	errs := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(errs) == 1, fmt.Sprintf("wrong err length: %s", len(errs)))
}

func TestValidateEventDeclInvalidMultiple(t *testing.T) {
	p := parser.ParseString("event Dog(Cat, Animal)")
	errs := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(errs) == 2, fmt.Sprintf("wrong err length: %s", len(errs)))

}

func TestValidateEventDeclMixed(t *testing.T) {
	p := parser.ParseString("event Dog(int, Cat)")
	errs := ValidateScope(p.Scope)
	goutil.AssertNow(t, len(errs) == 1, fmt.Sprintf("wrong err length: %s", len(errs)))
}

func TestValidateFuncDecl(t *testing.T) {

}

func TestValidateConstructorDecl(t *testing.T) {

}

func TestValidateContractDecl(t *testing.T) {

}
