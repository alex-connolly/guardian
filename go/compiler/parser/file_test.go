package parser

import (
	"log"
	"testing"

	"github.com/end-r/goutil"
	"github.com/end-r/guardian/go/compiler/ast"
)

func TestEmptyContract(t *testing.T) {
	p := ParseFile("tests/empty_contract.grd")
	goutil.Assert(t, len(p.lexer.Tokens) == 7, "wrong length of tokens")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, p.Errs == nil, "parser should not have errors")
	log.Println(p.Errs)
}

func TestConstructorContract(t *testing.T) {
	p := ParseFile("tests/constructor_contract.grd")
	goutil.Assert(t, p != nil, "parser should not be nil")
	goutil.Assert(t, p.Errs == nil, "parser should not have errors")
	log.Println(p.Errs)
	goutil.Assert(t, p.Scope.Type() == ast.File, "outer Scope should be file")
}
