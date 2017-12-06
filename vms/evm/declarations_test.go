package evm

import (
	"testing"

	"github.com/end-r/guardian/validator"

	"github.com/end-r/goutil"
	"github.com/end-r/guardian/parser"
)

func TestTraverseExplicitVariableDeclaration(t *testing.T) {
	ast, _ := parser.ParseString(`name uint8`)
	e := NewVM()
	types, _ := validator.Validate(ast, e)
	e.Traverse(*ast)
	goutil.Assert(t, len(e.storage) == 1, "didn't allocate a block")
}
