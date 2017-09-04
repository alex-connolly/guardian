package parser

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestIdentifierSafety(t *testing.T) {
	p := createParser("")
	// none of these should crash
	goutil.Assert(t, !isClassDeclaration(p), "class declaration not recognised")
	goutil.Assert(t, !isContractDeclaration(p), "contract declaration not recognised")
	goutil.Assert(t, !isInterfaceDeclaration(p), "interface declaration not recognised")
	goutil.Assert(t, !isEventDeclaration(p), "event declaration not recognised")
	goutil.Assert(t, !isFuncDeclaration(p), "func declaration not recognised")
}

func TestIsClassDeclaration(t *testing.T) {
	p := createParser("class Dog {")
	goutil.Assert(t, isClassDeclaration(p), "class declaration not recognised")
	p = createParser("abstract class Dog {")
	goutil.Assert(t, isClassDeclaration(p), "abstract class declaration not recognised")
}

func TestIsInterfaceDeclaration(t *testing.T) {
	p := createParser("interface Box {")
	goutil.Assert(t, isInterfaceDeclaration(p), "interface declaration not recognised")
	p = createParser("abstract interface Box {")
	goutil.Assert(t, isInterfaceDeclaration(p), "abstract interface declaration not recognised")
}

func TestIsContractDeclaration(t *testing.T) {
	p := createParser("contract Box {")
	goutil.Assert(t, isContractDeclaration(p), "contract declaration not recognised")
	p = createParser("abstract contract Box {")
	goutil.Assert(t, isContractDeclaration(p), "abstract contract declaration not recognised")
}

func TestIsFuncDeclaration(t *testing.T) {
	p := createParser("func main(){")
	goutil.Assert(t, isFuncDeclaration(p), "function declaration not recognised")
	p = createParser("abstract func main(){")
	goutil.Assert(t, isFuncDeclaration(p), "abstract function declaration not recognised")
	p = createParser("func main() int {")
	goutil.Assert(t, isFuncDeclaration(p), "returning function declaration not recognised")
	p = createParser("func main() (int, int) {")
	goutil.Assert(t, isFuncDeclaration(p), "tuple returning function declaration not recognised")
}

func TestIsTypeDeclaration(t *testing.T) {
	p := createParser("type Large int")
	goutil.Assert(t, isTypeDeclaration(p), "type declaration not recognised")
	p = createParser("type Large []int")
	goutil.Assert(t, isTypeDeclaration(p), "array type declaration not recognised")
	p = createParser("type Large map[int]string")
	goutil.Assert(t, isTypeDeclaration(p), "map type declaration not recognised")
}

func TestIsReturnStatement(t *testing.T) {
	p := createParser("return 0")
	goutil.Assert(t, isReturnStatement(p), "return statement not recognised")
	p = createParser("return (0, 0)")
	goutil.Assert(t, isReturnStatement(p), "tuple return statement not recognised")
}

func TestIsForStatement(t *testing.T) {
	p := createParser("for i := 0; i < 10; i++ {}")
	goutil.Assert(t, isForStatement(p), "for statement not recognised")
	p = createParser("for i < 10 {}")
	goutil.Assert(t, isForStatement(p), "cond only for statement not recognised")
	p = createParser("for i in 0...10{}")
	goutil.Assert(t, isForStatement(p), "range for statement not recognised")
	p = createParser("for i, _ in array")
	goutil.Assert(t, isForStatement(p), "")
}

func TestIsIfStatement(t *testing.T) {
	p := createParser("if x > 5 {}")
	goutil.Assert(t, isIfStatement(p), "if statement not recognised")
	p = createParser("if p := getData(); p < 5 {}")
	goutil.Assert(t, isIfStatement(p), "init if statement not recognised")
}

func TestIsSwitchStatement(t *testing.T) {
	p := createParser("switch x {}")
	goutil.Assert(t, isSwitchStatement(p), "switch statement not recognised")
	p = createParser("exclusive switch x {}")
	goutil.Assert(t, isSwitchStatement(p), "exclusive switch statement not recognised")
}

func TestIsCaseStatement(t *testing.T) {
	p := createParser("case 1, 2, 3: break")
	goutil.Assert(t, isCaseStatement(p), "multi case statement not recognised")
	p = createParser("case 1: break")
	goutil.Assert(t, isCaseStatement(p), "single case statement not recognised")
}

func TestIsEventDeclaration(t *testing.T) {
	p := createParser("event Notification()")
	goutil.Assert(t, isEventDeclaration(p), "empty event not recognised")
	p = createParser("event Notification(string)")
	goutil.Assert(t, isEventDeclaration(p), "single event not recognised")
	p = createParser("event Notification(string, dog.Dog)")
	goutil.Assert(t, isEventDeclaration(p), "multiple event not recognised")
}
