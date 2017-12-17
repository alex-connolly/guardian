package evm

import (
	"fmt"
	"testing"

	"github.com/end-r/guardian/validator"

	"github.com/end-r/guardian/parser"

	"github.com/end-r/goutil"
)

func TestTraverseSimpleIdentifierExpression(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("hello")
	bytes := e.traverseExpression(expr)
	//goutil.Assert(t, bytes.Compare(expected), invalidBytecodeMessage(bytes, expected))
	fmt.Println(bytes.Format())
	goutil.Assert(t, bytes.Length() == 2, "wrong bc length")
}

func TestTraverseLiteralsBinaryExpression(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("1 + 2")
	bytes := e.traverseExpression(expr)
	//goutil.Assert(t, bytes.Compare(expected), invalidBytecodeMessage(bytes, expected))
	fmt.Println(bytes.Format())
	goutil.Assert(t, bytes.Length() == 3, "wrong bc length")
}

func TestTraverseIdentifierBinaryExpression(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("a + b")
	bytes := e.traverseExpression(expr)
	//goutil.Assert(t, bytes.Compare(expected), invalidBytecodeMessage(bytes, expected))
	fmt.Println(bytes.Format())
	// should be two commands for each id
	goutil.Assert(t, bytes.Length() == 5, "wrong bc length")
}

func TestTraverseCallBinaryExpression(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("a() + b()")
	bytes := e.traverseExpression(expr)
	//goutil.Assert(t, bytes.Compare(expected), invalidBytecodeMessage(bytes, expected))
	fmt.Println(bytes.Format())
	// should be two commands for each id
	goutil.Assert(t, bytes.Length() == 5, "wrong bc length")
}

func TestTraverseIndexExpressionIdentifierLiteral(t *testing.T) {
	a, _ := parser.ParseString(`
		a [5]int
		x = a[1]
	`)
	e := NewVM()
	validator.Validate(a, NewVM())
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
	a, _ := validator.ValidateString(vm, `
		b [5][5]int
		x = b[2][3]
	`)
	bytecode, _ := NewVM().Traverse(a)
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
	a, _ := validator.ValidateString(vm, `
		b [5]int
		x = b[a]
	`)
	bytecode, _ := NewVM().Traverse(a)
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
	a, _ := validator.ValidateString(vm, `
		b [5]int
		x = b[a()]
	`)
	bytecode, _ := NewVM().Traverse(a)
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
	a, _ := validator.ValidateString(vm, `
		b [5]int
		a [4]int
		c = 3
		x = b[a[c]]
	`)
	bytecode, _ := NewVM().Traverse(a)
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
	a, _ := validator.ValidateString(vm, `
		x = a()[b]
	`)
	bytecode, _ := NewVM().Traverse(a)
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
	a, _ := validator.ValidateString(vm, `
		x = a()[1]
	`)
	bytecode, _ := NewVM().Traverse(a)
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
	a, _ := validator.ValidateString(vm, `
		x = a()[b()]
	`)
	bytecode, _ := NewVM().Traverse(a)
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
