package evm

import (
	"fmt"
	"testing"

	"github.com/end-r/guardian/parser"

	"github.com/end-r/goutil"
)

func TestTraverseSimpleIdentifierExpression(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("hello")
	bytes := e.traverseExpression(expr)
	//goutil.Assert(t, bytes.compare(expected), invalidBytecodeMessage(bytes, expected))
	fmt.Println(bytes.Format())
	goutil.Assert(t, bytes.Length() == 2, "wrong bc length")
}

func TestTraverseLiteralsBinaryExpression(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("1 + 2")
	bytes := e.traverseExpression(expr)
	//goutil.Assert(t, bytes.compare(expected), invalidBytecodeMessage(bytes, expected))
	fmt.Println(bytes.Format())
	goutil.Assert(t, bytes.Length() == 3, "wrong bc length")
}

func TestTraverseIdentifierBinaryExpression(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("a + b")
	bytes := e.traverseExpression(expr)
	//goutil.Assert(t, bytes.compare(expected), invalidBytecodeMessage(bytes, expected))
	fmt.Println(bytes.Format())
	// should be two commands for each id
	goutil.Assert(t, bytes.Length() == 5, "wrong bc length")
}

func TestTraverseCallBinaryExpression(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("a() + b()")
	bytes := e.traverseExpression(expr)
	//goutil.Assert(t, bytes.compare(expected), invalidBytecodeMessage(bytes, expected))
	fmt.Println(bytes.Format())
	// should be two commands for each id
	goutil.Assert(t, bytes.Length() == 5, "wrong bc length")
}

func TestTraverseIndexExpressionIdentifierLiteral(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("a[1]")
	bytes := e.traverseExpression(expr)
	//goutil.Assert(t, bytes.compare(expected), invalidBytecodeMessage(bytes, expected))
	fmt.Println(bytes.Format())
	// should be two commands for each id
	goutil.Assert(t, bytes.Length() == 5, "wrong bc length")
}

func TestTraverseIndexExpressionIdentifierIdentifier(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("a[b]")
	bytes := e.traverseExpression(expr)
	//goutil.Assert(t, bytes.compare(expected), invalidBytecodeMessage(bytes, expected))
	fmt.Println(bytes.Format())
	// should be two commands for each id
	goutil.Assert(t, bytes.Length() == 5, "wrong bc length")
}

func TestTraverseIndexExpressionIdentifierCall(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("a[b()]")
	bytes := e.traverseExpression(expr)
	//goutil.Assert(t, bytes.compare(expected), invalidBytecodeMessage(bytes, expected))
	fmt.Println(bytes.Format())
	// should be two commands for each id
	goutil.Assert(t, bytes.Length() == 5, "wrong bc length")
}

func TestTraverseIndexExpressionIdentifierIndex(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("a[b[c]]")
	bytes := e.traverseExpression(expr)
	//goutil.Assert(t, bytes.compare(expected), invalidBytecodeMessage(bytes, expected))
	fmt.Println(bytes.Format())
	// should be two commands for each id
	goutil.Assert(t, bytes.Length() == 5, "wrong bc length")
}

func TestTraverseIndexExpressionCallIdentifier(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("a()[b]")
	bytes := e.traverseExpression(expr)
	//goutil.Assert(t, bytes.compare(expected), invalidBytecodeMessage(bytes, expected))
	fmt.Println(bytes.Format())
	// should be two commands for each id
	goutil.Assert(t, bytes.Length() == 5, "wrong bc length")
}

func TestTraverseIndexExpressionCallLiteral(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("a()[1]")
	bytes := e.traverseExpression(expr)
	//goutil.Assert(t, bytes.compare(expected), invalidBytecodeMessage(bytes, expected))
	fmt.Println(bytes.Format())
	// should be two commands for each id
	goutil.Assert(t, bytes.Length() == 5, "wrong bc length")
}

func TestTraverseIndexExpressionCallCall(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("a()[b()]")
	bytes := e.traverseExpression(expr)
	//goutil.Assert(t, bytes.compare(expected), invalidBytecodeMessage(bytes, expected))
	fmt.Println(bytes.Format())
	// should be two commands for each id
	goutil.Assert(t, bytes.Length() == 5, "wrong bc length")
}

func TestTraverseReferenceCallEmpty(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("math.Pow()")
	bytes := e.traverseExpression(expr)
	//goutil.Assert(t, bytes.compare(expected), invalidBytecodeMessage(bytes, expected))
	fmt.Println(bytes.Format())
	// should be two commands for each id
	goutil.Assert(t, bytes.Length() == 5, "wrong bc length")
}

func TestTraverseReferenceCallArgs(t *testing.T) {
	e := new(GuardianEVM)
	expr := parser.ParseExpression("math.Pow(2, 2)")
	bytes := e.traverseExpression(expr)
	//goutil.Assert(t, bytes.compare(expected), invalidBytecodeMessage(bytes, expected))
	fmt.Println(bytes.Format())
	// should be two commands for each id
	goutil.Assert(t, bytes.Length() == 5, "wrong bc length")
}
