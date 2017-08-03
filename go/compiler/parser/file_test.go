package parser

import (
	"axia/guardian/go/util"
	"testing"
)

func TestEmptyFile(t *testing.T) {
	p := ParseFile("tests/empty_contract.grd")
	util.Assert(t, p != nil, "parser should not be nil")
	util.Assert(t, p.errs == nil, "parser should not have errors")
	util.Assert(t, p.Type() == ast.FileNode, "outer scope should be file")
}

func TestConstructorContract(t *testing.T) {
	p := ParseFile("tests/constructor_contract.grd")
	util.Assert(t, p != nil, "parser should not be nil")
	util.Assert(t, p.errs == nil, "parser should not have errors")
	util.Assert(t, p.Type() == ast.FileNode, "outer scope should be file")
}
