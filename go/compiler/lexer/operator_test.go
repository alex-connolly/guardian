package lexer

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestIsUnaryOperator(t *testing.T) {
	x := TknNot
	goutil.Assert(t, x.IsUnaryOperator(), "not should be unary")
	x = TknAdd
	goutil.Assert(t, !x.IsUnaryOperator(), "add should not be unary")
}

func TestIsBinaryOperator(t *testing.T) {
	x := TknAdd
	goutil.Assert(t, x.IsBinaryOperator(), "add should be binary")
	x = TknNot
	goutil.Assert(t, !x.IsBinaryOperator(), "not should not be binary")
}
