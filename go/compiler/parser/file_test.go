package parser

import (
	"testing"

	"github.com/end-r/guardian/go/compiler/ast"
	"github.com/end-r/guardian/go/util"
)

func TestEmptyFile(t *testing.T) {
	p := ParseFile("tests/empty_contract.grd")
	util.Assert(t, p != nil, "parser should not be nil")
	util.Assert(t, p.Errs == nil, "parser should not have errors")
	util.Assert(t, p.Scope.Type() == ast.File, "outer Scope should be file")
}

func TestConstructorContract(t *testing.T) {
	p := ParseFile("tests/constructor_contract.grd")
	util.Assert(t, p != nil, "parser should not be nil")
	util.Assert(t, p.Errs == nil, "parser should not have errors")
	util.Assert(t, p.Scope.Type() == ast.File, "outer Scope should be file")
}

func TestMacroContract(t *testing.T) {
	p := ParseFile("tests/macro_contract.grd")
	util.Assert(t, p != nil, "parser should not be nil")
	util.Assert(t, p.Errs == nil, "parser should not have errors")
	util.Assert(t, p.Scope.Type() == ast.File, "outer Scope should be file")
}
