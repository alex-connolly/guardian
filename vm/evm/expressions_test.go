package evm

import (
	"testing"

	"github.com/end-r/guardian/validator"

	"github.com/end-r/guardian/parser"

	"github.com/end-r/goutil"
)

func TestTraverseSimpleIdentifierExpression(t *testing.T) {
	e := NewVM()
	a, _ := validator.ValidateString(e, `
		hello = 5
		x = hello
	`)
	bytecode, _ := e.Traverse(a)
	expected := []string{
		// push x
		"PUSH",
		// push a
		"PUSH",
		// push index
		"PUSH",
		// push size of int
		"PUSH",
		// calculate offset
		"MUL",
		// create final position
		"ADD",
		"SLOAD",
	}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseLiteralsBinaryExpression(t *testing.T) {
	e := NewVM()
	a, _ := validator.ValidateString(e, `
		x = 1 + 2
	`)
	bytecode, _ := e.Traverse(a)
	expected := []string{
		// push x
		"PUSH",
		// push a
		"PUSH",
		// push index
		"PUSH",
		// push size of int
		"PUSH",
		// calculate offset
		"MUL",
		// create final position
		"ADD",
		"SLOAD",
	}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseIdentifierBinaryExpression(t *testing.T) {
	e := NewVM()
	a, _ := validator.ValidateString(e, `
		a = 1
		b = 2
		x = a + b
	`)
	bytecode, _ := e.Traverse(a)
	expected := []string{
		// push x
		"PUSH",
		// push a
		"PUSH",
		// push index
		"PUSH",
		// push size of int
		"PUSH",
		// calculate offset
		"MUL",
		// create final position
		"ADD",
		"SLOAD",
	}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseCallBinaryExpression(t *testing.T) {
	e := NewVM()
	a, _ := validator.ValidateString(e, `

		func a() int {
			return 1
		}

		func b() int {
			return 2
		}

		x = a() + b()
	`)
	bytecode, _ := e.Traverse(a)
	expected := []string{
		// push x
		"PUSH",
		// push a
		"PUSH",
		// push index
		"PUSH",
		// push size of int
		"PUSH",
		// calculate offset
		"MUL",
		// create final position
		"ADD",
		"SLOAD",
	}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseIndexExpressionIdentifierLiteral(t *testing.T) {
	e := NewVM()
	a, _ := validator.ValidateString(e, `
		var b [5]int
		x = b[1]
	`)
	bytecode, _ := e.Traverse(a)
	expected := []string{
		// push x
		"PUSH",
		// push a
		"PUSH",
		// push index
		"PUSH",
		// push size of int
		"PUSH",
		// calculate offset
		"MUL",
		// create final position
		"ADD",
		"SLOAD",
	}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseTwoDimensionalArray(t *testing.T) {
	e := NewVM()
	a, _ := validator.ValidateString(e, `
		b [5][5]int
		x = b[2][3]
	`)
	bytecode, _ := e.Traverse(a)
	expected := []string{
		// push x
		"PUSH",
		// push b
		"PUSH",
		// push index (2)
		"PUSH",
		// push size of type
		"PUSH",
	}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseIndexExpressionIdentifierIdentifier(t *testing.T) {
	e := NewVM()
	a, _ := validator.ValidateString(e, `
		b [5]int
		x = b[a]
	`)
	bytecode, _ := e.Traverse(a)
	expected := []string{
		// push x
		"PUSH",
		// push b
		"PUSH",
		// push index (2)
		"PUSH",
		// push size of type
		"PUSH",
	}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseIndexExpressionIdentifierCall(t *testing.T) {
	e := NewVM()
	a, _ := validator.ValidateString(e, `
		b [5]int
		x = b[a()]
	`)
	bytecode, _ := e.Traverse(a)
	expected := []string{
		// push x
		"PUSH",
		// push b
		"PUSH",
		// push index (2)
		"PUSH",
		// push size of type
		"PUSH",
	}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseIndexExpressionIdentifierIndex(t *testing.T) {
	e := NewVM()
	a, _ := validator.ValidateString(e, `
		b [5]int
		a [4]int
		c = 3
		x = b[a[c]]
	`)
	bytecode, _ := e.Traverse(a)
	expected := []string{
		// push x
		"PUSH",
		// push b
		"PUSH",
		// push index (2)
		"PUSH",
		// push size of type
		"PUSH",
	}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseIndexExpressionCallIdentifier(t *testing.T) {
	e := NewVM()
	a, _ := validator.ValidateString(e, `
		x = a()[b]
	`)
	bytecode, _ := e.Traverse(a)
	expected := []string{
		// push x
		"PUSH",
		// push b
		"PUSH",
		// push index (2)
		"PUSH",
		// push size of type
		"PUSH",
	}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseIndexExpressionCallLiteral(t *testing.T) {
	e := NewVM()
	a, _ := validator.ValidateString(e, `
		x = a()[1]
	`)
	bytecode, _ := e.Traverse(a)
	expected := []string{
		// push x
		"PUSH",
		// push b
		"PUSH",
		// push index (2)
		"PUSH",
		// push size of type
		"PUSH",
	}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseIndexExpressionCallCall(t *testing.T) {
	e := NewVM()
	a, _ := validator.ValidateString(e, `
		x = a()[b()]
	`)
	bytecode, _ := e.Traverse(a)
	expected := []string{
		// push x
		"PUSH",
		// push b
		"PUSH",
		// push index (2)
		"PUSH",
		// push size of type
		"PUSH",
	}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseLiteral(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("0")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseLiteralTwoBytes(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("256")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH2"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseLiteralThirtyTwoBytes(t *testing.T) {
	e := new(GuardianEVM)
	// 2^256
	expr := parser.ParseExpression("115792089237316195423570985008687907853269984665640564039457584007913129639936")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH32"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestSimpleBinarySignedLess(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("3 < 4")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "SLT"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestSimpleBinarySignedLessEqual(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("3 <= 4")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "SGT", "NOT"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestSimpleBinarySignedGreater(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("3 < 4")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "SLT"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestSimpleBinarySignedGreaterEqual(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("3 >= 4")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "SLT", "NOT"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestSimpleBinaryEqual(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("3 == 4")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "EQ"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestSimpleBinaryNotEqual(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("3 == 4")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "EQ", "NOT"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestSimpleBinaryAnd(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("3 & 4")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "AND"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestSimpleBinaryOr(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("3 | 4")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "OR"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestSimpleBinaryXor(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("3 ^ 4")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "XOR"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestSimpleBinaryLogicalAnd(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("true and false")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "OR"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestSimpleBinaryLogicalOr(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("true or false")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "OR"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestSimpleBinaryAddition(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("3 + 5")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "ADD"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestSimpleBinarySubtraction(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("3 - 5")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "SUB"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestSimpleBinaryMultiplication(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("3 * 5")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "MUL"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestSimpleBinarySignedDivision(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("4 / 2")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "DIV"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestSimpleBinaryDivision(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("4 as uint / 2 as uint")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "DIV"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestSimpleBinarySignedMod(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("4 % 2")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "SMOD"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestSimpleBinaryMod(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("4 as uint % 2 as uint")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "MOD"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestSimpleBinaryExp(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("2 ** 4")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "EXP"}
	goutil.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}
