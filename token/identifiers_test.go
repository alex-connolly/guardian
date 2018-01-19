package token

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestIsUnaryOperator(t *testing.T) {
	x := Not
	goutil.Assert(t, x.IsUnaryOperator(), "not should be unary")
	x = Add
	goutil.Assert(t, !x.IsUnaryOperator(), "add should not be unary")
}

func TestIsBinaryOperator(t *testing.T) {
	x := Add
	goutil.Assert(t, x.IsBinaryOperator(), "add should be binary")
	x = Not
	goutil.Assert(t, !x.IsBinaryOperator(), "not should not be binary")
}

func TestIsIdentifierByte(t *testing.T) {
	goutil.Assert(t, isIdentifierByte('A'), "upper letter not id byte")
	goutil.Assert(t, isIdentifierByte('t'), "lower letter not id byte")
	goutil.Assert(t, isIdentifierByte('2'), "number not id byte")
	goutil.Assert(t, isIdentifierByte('_'), "underscore not id byte")
	goutil.Assert(t, !isIdentifierByte(' '), "space should not be id byte")
}

func TestIsNumber(t *testing.T) {
	byt := []byte(`9`)
	b := &bytecode{bytes: byt}
	goutil.Assert(t, isInteger(b), "positive integer")
	byt = []byte(`-9`)
	b = &bytecode{bytes: byt}
	goutil.Assert(t, isInteger(b), "negative integer")
	byt = []byte(`9.0`)
	b = &bytecode{bytes: byt}
	goutil.Assert(t, isFloat(b), "positive float")
	byt = []byte(`-9.0`)
	b = &bytecode{bytes: byt}
	goutil.Assert(t, isFloat(b), "negative float")
	byt = []byte(`-.9`)
	b = &bytecode{bytes: byt}
	goutil.Assert(t, isFloat(b), "negative float")
}

func TestIsWhitespace(t *testing.T) {
	byt := []byte(` `)
	b := &bytecode{bytes: byt}
	goutil.Assert(t, isWhitespace(b), "space")

	byt = []byte(`	`)
	b = &bytecode{bytes: byt}
	goutil.Assert(t, isWhitespace(b), "tab")
}

func TestIsAssignment(t *testing.T) {
	a := Assign
	goutil.AssertNow(t, a.IsAssignment(), "assign not assignment")
}

func TestIsDeclaration(t *testing.T) {
	a := Event
	goutil.AssertNow(t, a.IsDeclaration(), "event not decl")
}
