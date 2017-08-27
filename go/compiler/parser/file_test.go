package parser

import (
	"log"
	"testing"

	"github.com/end-r/goutil"
	"github.com/end-r/guardian/go/compiler/ast"
)

func TestEmptyFile(t *testing.T) {
	p := ParseFile("tests/empty_contract.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, p.Errs == nil, "parser should not have errors")
	goutil.Assert(t, p.Scope.Type() == ast.File, "outer Scope should be file")
	log.Println(p.Scope.Type())
}

func TestConstructorContract(t *testing.T) {
	p := ParseFile("tests/constructor_contract.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, p.Errs == nil, "parser should not have errors")
	log.Println(p.Errs)
	goutil.Assert(t, p.Scope.Type() == ast.File, "outer Scope should be file")
}

func TestMacroContract(t *testing.T) {
	p := ParseFile("tests/macro_contract.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, p.Errs == nil, "parser should not have errors")
	goutil.Assert(t, p.Scope.Type() == ast.File, "outer Scope should be file")
}
