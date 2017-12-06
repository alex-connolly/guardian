package evm

import (
	"fmt"
	"testing"

	"github.com/end-r/guardian/validator"

	"github.com/end-r/goutil"
	"github.com/end-r/guardian/parser"
)

func TestTraverseExplicitVariableDeclaration(t *testing.T) {
	ast, errs := parser.ParseString(`name uint8`)
	goutil.AssertNow(t, errs == nil, "errs should be nil")
	goutil.AssertNow(t, ast != nil, "ast shouldn't be nil")
	goutil.AssertNow(t, ast.Declarations != nil, "ast decls shouldn't be nil")
	e := NewVM()
	fmt.Println("$$")
	errs = validator.Validate(ast, e)
	goutil.Assert(t, errs == nil, "errs should be nil")
	fmt.Println("$$$")
	e.Traverse(ast)
	goutil.Assert(t, len(e.storage) == 1, "didn't allocate a block")
}
