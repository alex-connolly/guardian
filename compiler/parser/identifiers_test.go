package parser

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestIdentifierSafety(t *testing.T) {
	p := createParser("")
	// none of these should crash
	for _, c := range getPrimaryConstructs() {
		c.is(p)
	}
}

func TestIsClassDeclaration(t *testing.T) {
	p := createParser("class Dog {")
	goutil.Assert(t, isClassDeclaration(p), "class declaration not recognised")
	p = createParser("abstract class Dog {")
	goutil.Assert(t, isClassDeclaration(p), "abstract class declaration not recognised")
	p = createParser("external class Box {")
	goutil.Assert(t, isClassDeclaration(p), "modifier statement not recognised")
	p = createParser("abstract external class Box {")
	goutil.Assert(t, isClassDeclaration(p), "double modifier statement not recognised")
}

func TestIsInterfaceDeclaration(t *testing.T) {
	p := createParser("interface Box {")
	goutil.Assert(t, isInterfaceDeclaration(p), "interface declaration not recognised")
	p = createParser("abstract interface Box {")
	goutil.Assert(t, isInterfaceDeclaration(p), "abstract interface declaration not recognised")
	p = createParser("external interface Box {")
	goutil.Assert(t, isInterfaceDeclaration(p), "modifier statement not recognised")
	p = createParser("abstract external interface Box {")
	goutil.Assert(t, isInterfaceDeclaration(p), "double modifier statement not recognised")
}

func TestIsContractDeclaration(t *testing.T) {
	p := createParser("contract Box {")
	goutil.Assert(t, isContractDeclaration(p), "contract declaration not recognised")
	p = createParser("abstract contract Box {")
	goutil.Assert(t, isContractDeclaration(p), "abstract contract declaration not recognised")
	p = createParser("external func main() (int, int) {")
	goutil.Assert(t, isFuncDeclaration(p), "modifier statement not recognised")
	p = createParser("abstract external func main() (int, int) {")
	goutil.Assert(t, isFuncDeclaration(p), "double modifier statement not recognised")
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
	p = createParser("external func main() (int, int) {")
	goutil.Assert(t, isFuncDeclaration(p), "modifier statement not recognised")
	p = createParser("abstract external func main() (int, int) {")
	goutil.Assert(t, isFuncDeclaration(p), "double modifier statement not recognised")
}

func TestIsTypeDeclaration(t *testing.T) {
	p := createParser("type Large int")
	goutil.Assert(t, isTypeDeclaration(p), "type declaration not recognised")
	p = createParser("type Large []int")
	goutil.Assert(t, isTypeDeclaration(p), "array type declaration not recognised")
	p = createParser("type Large map[int]string")
	goutil.Assert(t, isTypeDeclaration(p), "map type declaration not recognised")
	p = createParser("external type Large int")
	goutil.Assert(t, isTypeDeclaration(p), "modifier statement not recognised")
	p = createParser("abstract external type Large int")
	goutil.Assert(t, isTypeDeclaration(p), "double modifier statement not recognised")
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

func TestIsExplicitVarDeclaration(t *testing.T) {
	p := createParser("x string")
	goutil.Assert(t, isExplicitVarDeclaration(p), "var expvar statement not recognised")
	p = createParser("x, a string")
	goutil.Assert(t, isExplicitVarDeclaration(p), "multiple var expvar statement not recognised")
	p = createParser("x map[string]string")
	goutil.Assert(t, isExplicitVarDeclaration(p), "map expvar statement not recognised")
	p = createParser("x [string]")
	goutil.Assert(t, isExplicitVarDeclaration(p), "array expvar statement not recognised")
	p = createParser("external x string")
	goutil.Assert(t, isExplicitVarDeclaration(p), "modifier statement not recognised")
	p = createParser("abstract external x string")
	goutil.Assert(t, isExplicitVarDeclaration(p), "double modifier statement not recognised")
	p = createParser("x = 5")
	goutil.Assert(t, !isExplicitVarDeclaration(p), "should not recognise simple assignment")
	p = createParser("a[b] = 5")
	goutil.Assert(t, !isExplicitVarDeclaration(p), "should not recognise index assignment")
	p = createParser("a[b].c = 5")
	goutil.Assert(t, !isExplicitVarDeclaration(p), "should not recognise reference assignment")
}

func TestIsSwitchStatement(t *testing.T) {
	p := createParser("switch x {}")
	goutil.Assert(t, isSwitchStatement(p), "switch statement not recognised")
	p = createParser("exclusive switch x {}")
	goutil.Assert(t, isSwitchStatement(p), "exclusive switch statement not recognised")
}

func TestIsCaseStatement(t *testing.T) {
	p := createParser("case 1, 2, 3 { break }")
	goutil.Assert(t, isCaseStatement(p), "multi case statement not recognised")
	p = createParser("case 1 { break }")
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

func TestIsAssignmentEdgeCases(t *testing.T) {
	p := createParser("5")
	goutil.Assert(t, !isAssignmentStatement(p), "expression should not be recognised")
	p = createParser("")
	goutil.Assert(t, !isAssignmentStatement(p), "blank statement should not be recognised")
	p = createParser("{}")
	goutil.Assert(t, !isAssignmentStatement(p), "braces should not be recognised")
	p = createParser("{")
	goutil.Assert(t, !isAssignmentStatement(p), "open brace should not be recognised")
	p = createParser("}")
	goutil.Assert(t, !isAssignmentStatement(p), "close brace should not be recognised")
}

func TestIsAssignmentStatementReferenceLiteral(t *testing.T) {
	p := createParser("x = 5")
	goutil.Assert(t, isAssignmentStatement(p), "simple assignment not recognised")
	p = createParser("x, y = 5")
	goutil.Assert(t, isAssignmentStatement(p), "multiple left assignment not recognised")
	p = createParser("x, y = 5, 3")
	goutil.Assert(t, isAssignmentStatement(p), "multiple l/r not recognised")
	p = createParser("x := 5")
	goutil.Assert(t, isAssignmentStatement(p), "simple definition not recognised")
	p = createParser("x > 5")
	goutil.Assert(t, !isAssignmentStatement(p), "comparison should not be recognised")
	p = createParser("proposals[p].voteCount > winningVoteCount")
	goutil.Assert(t, !isAssignmentStatement(p), "complex comparison shoudl not be recognised")
}

func TestIsAssignmentStatementIncrementDecrement(t *testing.T) {
	p := createParser("x++")
	goutil.Assert(t, isAssignmentStatement(p), "simple increment not recognised")
	p = createParser("y--")
	goutil.Assert(t, isAssignmentStatement(p), "simple decrement not recognised")

}

func TestIsNextAssignmentStatement(t *testing.T) {
	p := createParser("++")
	goutil.Assert(t, p.isNextTokenAssignment(), "simple increment not recognised")
	p = createParser("--")
	goutil.Assert(t, p.isNextTokenAssignment(), "simple decrement not recognised")

}
