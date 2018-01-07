package evm

import (
	"testing"

	"github.com/end-r/guardian/validator"

	"github.com/end-r/guardian/parser"

	"github.com/end-r/goutil"
)

func TestBuiltins(t *testing.T) {
	expr, _ := parser.ParseFile("test/builtins.grd")
	errs := validator.Validate(expr, NewVM())
	goutil.Assert(t, expr != nil, "expr is nil")
	goutil.Assert(t, len(errs) == 0, "expr is nil")
}

func TestGreeter(t *testing.T) {
	expr, _ := parser.ParseFile("test/greeter.grd")
	errs := validator.Validate(expr, NewVM())
	goutil.Assert(t, expr != nil, "expr is nil")
	goutil.Assert(t, len(errs) == 0, "expr is nil")
}
