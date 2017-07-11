package parser

import (
	"axia/guardian/go/util"
	"testing"
)

func TestIsMapLiteral(t *testing.T) {
	p := createParser("map[key]value{}")
	util.Assert(t, p.isMapLiteral(), "map literal not recognised")
	// errors but should still be recognised
	p = createParser("map[]{}")
	util.Assert(t, p.isMapLiteral(), "map literal no key not recognised")
	p = createParser("map{}")
	util.Assert(t, p.isMapLiteral(), "map literal no key no squares not recognised")
}

func TestIsArrayLiteral(t *testing.T) {
	p := createParser("[]int{1, 2, 3}")
	util.Assert(t, p.isArrayLiteral(), "array literal not recognised")
	p = createParser("[]int{}")
	util.Assert(t, p.isArrayLiteral(), "empty array literal not recognised")
}

func TestIsLiteral(t *testing.T) {
	p := createParser(`"hello"`)
	util.Assert(t, p.isLiteral(), "string literal not recognised")
	p = createParser(`6`)
	util.Assert(t, p.isLiteral(), "int literal not recognised")
	p = createParser(`'c'`)
	util.Assert(t, p.isLiteral(), "character literal not recognised")
	p = createParser(`7.4`)
	util.Assert(t, p.isLiteral(), "float literal not recognised")
}

func TestIsCompositeLiteral(t *testing.T) {
	p := createParser("Dog{4, 4}")
	util.Assert(t, p.isCompositeLiteral(), "composite literal not recognised")
	p = createParser("D{}")
	util.Assert(t, p.isCompositeLiteral(), "empty composite literal not recognised")
}

func TestIsCallExpression(t *testing.T) {
	p := createParser("hi()")
	util.Assert(t, p.isCallExpression(), "call expression not recognised")
	p = createParser("hi(2, 4, 5)")
	util.Assert(t, p.isCallExpression(), "arguments call expression not recognised")
}

func TestIsClassDeclaration(t *testing.T) {
	p := createParser("class Dog {")
	util.Assert(t, isClassDeclaration(p), "class declaration not recognised")
	p = createParser("abstract class Dog {")
	util.Assert(t, isClassDeclaration(p), "abstract class declaration not recognised")
}

func TestIsInterfaceDeclaration(t *testing.T) {
	p := createParser("interface Box {")
	util.Assert(t, isInterfaceDeclaration(p), "interface declaration not recognised")
	p = createParser("abstract interface Box {")
	util.Assert(t, isInterfaceDeclaration(p), "abstract interface declaration not recognised")
}

func TestIsContractDeclaration(t *testing.T) {
	p := createParser("contract Box {")
	util.Assert(t, isContractDeclaration(p), "contract declaration not recognised")
	p = createParser("abstract contract Box {")
	util.Assert(t, isContractDeclaration(p), "abstract contract declaration not recognised")
}

func TestIsFuncDeclaration(t *testing.T) {
	p := createParser("main(){")
	util.Assert(t, isFuncDeclaration(p), "function declaration not recognised")
	p = createParser("abstract main(){")
	util.Assert(t, isFuncDeclaration(p), "abstract function declaration not recognised")
	p = createParser("main() int {")
	util.Assert(t, isFuncDeclaration(p), "returning function declaration not recognised")
	p = createParser("main() (int, int) {")
	util.Assert(t, isFuncDeclaration(p), "tuple returning function declaration not recognised")
}

func TestIsTypeDeclaration(t *testing.T) {
	p := createParser("type Large int")
	util.Assert(t, isTypeDeclaration(p), "type declaration not recognised")
	p = createParser("type Large []int")
	util.Assert(t, isTypeDeclaration(p), "array type declaration not recognised")
	p = createParser("type Large map[int]string")
	util.Assert(t, isTypeDeclaration(p), "map type declaration not recognised")
}

func TestIsReturnStatement(t *testing.T) {
	p := createParser("return 0")
	util.Assert(t, isReturnStatement(p), "return statement not recognised")
	p = createParser("return (0, 0)")
	util.Assert(t, isReturnStatement(p), "tuple return statement not recognised")
}

func TestIsForStatement(t *testing.T) {
	p := createParser("for i := 0; i < 10; i++ {}")
	util.Assert(t, isForStatement(p), "for statement not recognised")
	p = createParser("for i < 10 {}")
	util.Assert(t, isForStatement(p), "cond only for statement not recognised")
	p = createParser("for i in 0...10{}")
	util.Assert(t, isForStatement(p), "range for statement not recognised")
	p = createParser("for i, _ in array")
	util.Assert(t, isForStatement(p), "")
}

func TestIsIfStatement(t *testing.T) {
	p := createParser("if x > 5 {}")
	util.Assert(t, isIfStatement(p), "if statement not recognised")
	p = createParser("if p := getData(); p < 5 {}")
	util.Assert(t, isIfStatement(p), "init if statement not recognised")
}

func TestIsSwitchStatement(t *testing.T) {
	p := createParser("switch x {}")
	util.Assert(t, isSwitchStatement(p), "switch statement not recognised")
	p = createParser("exclusive switch x {}")
	util.Assert(t, isSwitchStatement(p), "exclusive switch statement not recognised")
}

func TestIsCaseStatement(t *testing.T) {
	p := createParser("case 1, 2, 3: break")
	util.Assert(t, isCaseStatement(p), "multi case statement not recognised")
	p = createParser("case 1: break")
	util.Assert(t, isCaseStatement(p), "single case statement not recognised")
}
