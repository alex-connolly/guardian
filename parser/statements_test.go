package parser

import (
	"fmt"
	"testing"

	"github.com/end-r/guardian/ast"

	"github.com/end-r/goutil"
)

func TestParseAssignmentStatementSingleConstant(t *testing.T) {
	p, _ := ParseString("x = 6")
	goutil.AssertNow(t, p != nil, "scope is nil")
	goutil.AssertLength(t, len(p.Sequence), 1)
	first := p.Next()
	goutil.AssertNow(t, first.Type() == ast.AssignmentStatement, "wrong node type")
	assignmentStmt := first.(*ast.AssignmentStatementNode)
	goutil.AssertNow(t, len(assignmentStmt.Left) == 1, "wrong left length")
	goutil.AssertNow(t, len(assignmentStmt.Right) == 1, "wrong right length")
	left := assignmentStmt.Left[0]
	right := assignmentStmt.Right[0]
	goutil.AssertNow(t, left != nil, "left is nil")
	goutil.AssertNow(t, left.Type() == ast.Identifier, "wrong left type")
	goutil.AssertNow(t, right != nil, "right is nil")
	goutil.AssertNow(t, right.Type() == ast.Literal, "wrong right type")
}

func TestParseAssignmentStatementIncrement(t *testing.T) {
	_, errs := ParseString("x++")
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestParseIfStatement(t *testing.T) {
	p := createParser(`if x == 2 {

	}`)
	goutil.Assert(t, isIfStatement(p), "should detect if statement")
	parseIfStatement(p)
	goutil.AssertNow(t, len(p.errs) == 0, fmt.Sprintln(p.errs))
	first := p.scope.Next()
	goutil.Assert(t, first.Type() == ast.IfStatement, "wrong node type")
}

func TestParseIfStatementComplex(t *testing.T) {
	p := createParser(`if proposals[p].voteCount > winningVoteCount {

	}`)
	goutil.Assert(t, isIfStatement(p), "should detect if statement")
	parseIfStatement(p)
	goutil.AssertNow(t, len(p.errs) == 0, p.errs.Format())
	first := p.scope.Next()
	goutil.Assert(t, first.Type() == ast.IfStatement, "wrong node type")
}

func TestParseIfStatementInit(t *testing.T) {
	p := createParser(`if x = 0; x > 5 {

	}`)
	goutil.Assert(t, isIfStatement(p), "should detect if statement")
	parseIfStatement(p)
	goutil.AssertNow(t, len(p.errs) == 0, p.errs.Format())
	first := p.scope.Next()
	goutil.Assert(t, first.Type() == ast.IfStatement, "wrong node type")
}

func TestParseIfStatementElse(t *testing.T) {
	p := createParser(`if x == 2 {

	} else {

	}`)
	goutil.Assert(t, isIfStatement(p), "should detect if statement")
	parseIfStatement(p)
	goutil.AssertNow(t, len(p.errs) == 0, p.errs.Format())
	first := p.scope.Next()
	goutil.Assert(t, first.Type() == ast.IfStatement, "wrong node type")
	ifStat := first.(*ast.IfStatementNode)
	goutil.Assert(t, ifStat.Init == nil, "init should be nil")
}

func TestParseIfStatementInitElse(t *testing.T) {
	p := createParser(`if x = 0; x > 5 {

	} else {

	}`)
	goutil.Assert(t, isIfStatement(p), "should detect if statement")
	parseIfStatement(p)
	goutil.AssertNow(t, len(p.errs) == 0, p.errs.Format())
	first := p.scope.Next()
	goutil.Assert(t, first.Type() == ast.IfStatement, "wrong node type")
	ifStat := first.(*ast.IfStatementNode)
	goutil.Assert(t, ifStat.Init != nil, "init shouldn't be nil")
}

func TestParseIfStatementElifElse(t *testing.T) {
	p := createParser(`if x > 4 {

	} elif x < 4 {

	} else {

	}`)
	goutil.Assert(t, isIfStatement(p), "should detect if statement")
	parseIfStatement(p)
	goutil.AssertNow(t, len(p.errs) == 0, p.errs.Format())
	first := p.scope.Next()
	goutil.Assert(t, first.Type() == ast.IfStatement, "wrong node type")
	ifStat := first.(*ast.IfStatementNode)
	goutil.Assert(t, ifStat.Init == nil, "init should be nil")
	goutil.AssertNow(t, len(ifStat.Conditions) == 2, "should have two conditions")
	nextFirst := ifStat.Conditions[0]
	goutil.AssertNow(t, nextFirst.Condition.Type() == ast.BinaryExpression, "first binary expr not recognised")
	nextSecond := ifStat.Conditions[1]
	goutil.AssertNow(t, nextSecond.Condition.Type() == ast.BinaryExpression, "second binary expr not recognised")
	goutil.AssertNow(t, ifStat.Else != nil, "else should not be nil")
}

func TestParseForStatementCondition(t *testing.T) {
	p := createParser(`for x < 5 {}`)
	goutil.Assert(t, isForStatement(p), "should detect for statement")
	parseForStatement(p)
	goutil.Assert(t, len(p.errs) == 0, "should be error-free")

	first := p.scope.Next()
	goutil.Assert(t, first.Type() == ast.ForStatement, "wrong node type")
	forStat := first.(*ast.ForStatementNode)
	goutil.Assert(t, forStat.Init == nil, "init should be nil")
	goutil.Assert(t, forStat.Cond != nil, "cond should not be nil")
	goutil.Assert(t, forStat.Post == nil, "post should be nil")
}

func TestParseForStatementInitCondition(t *testing.T) {
	p := createParser(`for x = 0; x < 5 {}`)
	goutil.Assert(t, isForStatement(p), "should detect for statement")
	parseForStatement(p)
	goutil.AssertNow(t, len(p.errs) == 0, p.errs.Format())
	first := p.scope.Next()
	goutil.Assert(t, first.Type() == ast.ForStatement, "wrong node type")
	forStat := first.(*ast.ForStatementNode)
	goutil.Assert(t, forStat.Init != nil, "init should not be nil")
	goutil.Assert(t, forStat.Cond != nil, "cond should not be nil")
	goutil.Assert(t, forStat.Post == nil, "post should be nil")
}

func TestParseForStatementInitConditionStatement(t *testing.T) {
	p := createParser(`for x = 0; x < 5; x++ {}`)
	goutil.Assert(t, isForStatement(p), "should detect for statement")
	parseForStatement(p)
	goutil.Assert(t, len(p.errs) == 0, "should be error-free")
	goutil.AssertNow(t, p.scope != nil, "scope should not be nil")
	goutil.AssertNow(t, len(p.scope.Sequence) == 1, fmt.Sprintf("wrong sequence length: %d", len(p.scope.Sequence)))
	first := p.scope.Next()
	goutil.Assert(t, first.Type() == ast.ForStatement, "wrong node type")
	forStat := first.(*ast.ForStatementNode)
	goutil.Assert(t, forStat.Init != nil, "init should not be nil")
	goutil.Assert(t, forStat.Cond != nil, "cond should not be nil")
	goutil.Assert(t, forStat.Post != nil, "post should not be nil")
}

func TestParseSwitchStatement(t *testing.T) {
	p := createParser(`switch x {}`)
	goutil.AssertNow(t, isSwitchStatement(p), "should detect switch statement")
	parseSwitchStatement(p)
}

func TestParseSwitchStatementSingleCase(t *testing.T) {
	p := createParser(`switch x { case 5{}}`)
	goutil.Assert(t, isSwitchStatement(p), "should detect switch statement")
	parseSwitchStatement(p)
}

func TestParseSwitchStatementMultiCase(t *testing.T) {
	p := createParser(`switch x {
		case 5:
			x += 2
			break
		case 4:
			x *= 2
			break
	}`)
	goutil.Assert(t, isSwitchStatement(p), "should detect switch statement")
	parseSwitchStatement(p)
}

func TestParseSwitchStatementExclusive(t *testing.T) {
	p := createParser(`exclusive switch x {}
        `)
	goutil.Assert(t, isSwitchStatement(p), "should detect switch statement")
	parseSwitchStatement(p)

}

func TestParseCaseStatementSingle(t *testing.T) {
	p := createParser(`case 5:`)
	goutil.Assert(t, isCaseStatement(p), "should detect case statement")
	parseCaseStatement(p)

}

func TestParseCaseStatementMultiple(t *testing.T) {
	p := createParser(`case 5, 8, 9 {}`)
	goutil.Assert(t, isCaseStatement(p), "should detect case statement")
	parseCaseStatement(p)
}

func TestEmptyReturnStatement(t *testing.T) {
	p := createParser("return")
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)
}

func TestSingleLiteralReturnStatement(t *testing.T) {
	p := createParser(`return "twenty"`)
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)

	u := p.scope.Next()
	goutil.AssertNow(t, u.Type() == ast.ReturnStatement, "wrong return type")
	r := u.(*ast.ReturnStatementNode)
	goutil.AssertNow(t, len(r.Results) == 1, "wrong result length")
	goutil.AssertNow(t, r.Results[0].Type() == ast.Literal, "wrong literal type")
}

func TestMultipleLiteralReturnStatement(t *testing.T) {
	p := createParser(`return "twenty", "thirty"`)
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)

	u := p.scope.Next()
	goutil.AssertNow(t, u.Type() == ast.ReturnStatement, "wrong return type")
	r := u.(*ast.ReturnStatementNode)
	goutil.AssertNow(t, len(r.Results) == 2, "wrong result length")
	goutil.AssertNow(t, r.Results[0].Type() == ast.Literal, "wrong result 0 type")
	goutil.AssertNow(t, r.Results[1].Type() == ast.Literal, "wrong result 1 type")
}

func TestSingleReferenceReturnStatement(t *testing.T) {
	p := createParser(`return twenty`)
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)

	u := p.scope.Next()
	goutil.AssertNow(t, u.Type() == ast.ReturnStatement, "wrong return type")
	r := u.(*ast.ReturnStatementNode)
	goutil.AssertNow(t, len(r.Results) == 1, "wrong result length")
	goutil.AssertNow(t, r.Results[0].Type() == ast.Identifier, "wrong result 0 type")
}

func TestMultipleReferenceReturnStatement(t *testing.T) {
	p := createParser(`return twenty, thirty`)
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)

	u := p.scope.Next()
	goutil.AssertNow(t, u.Type() == ast.ReturnStatement, "wrong return type")
	r := u.(*ast.ReturnStatementNode)
	goutil.AssertNow(t, len(r.Results) == 2, "wrong result length")
	goutil.AssertNow(t, r.Results[0].Type() == ast.Identifier, "wrong result 0 type")
	goutil.AssertNow(t, r.Results[1].Type() == ast.Identifier, "wrong result 1 type")
}

func TestSingleCallReturnStatement(t *testing.T) {
	p := createParser(`return param()`)
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)

	u := p.scope.Next()
	goutil.AssertNow(t, u.Type() == ast.ReturnStatement, "wrong return type")
	r := u.(*ast.ReturnStatementNode)
	goutil.AssertNow(t, len(r.Results) == 1, "wrong result length")
	goutil.AssertNow(t, r.Results[0].Type() == ast.CallExpression, "wrong result 0 type")
}

func TestMultipleCallReturnStatement(t *testing.T) {
	p := createParser(`return a(param, "param"), b()`)
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)

	u := p.scope.Next()
	goutil.AssertNow(t, u.Type() == ast.ReturnStatement, "wrong return type")
	r := u.(*ast.ReturnStatementNode)
	goutil.AssertNow(t, len(r.Results) == 2, fmt.Sprintf("wrong result length: %d", len(r.Results)))
	goutil.AssertNow(t, r.Results[0].Type() == ast.CallExpression, "wrong result 0 type")
	c1 := r.Results[0].(*ast.CallExpressionNode)
	goutil.AssertNow(t, len(c1.Arguments) == 2, fmt.Sprintf("wrong c1 args length: %d", len(c1.Arguments)))
	goutil.AssertNow(t, r.Results[1].Type() == ast.CallExpression, "wrong result 1 type")
}

func TestSingleArrayLiteralReturnStatement(t *testing.T) {
	p := createParser(`return []int{}`)
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)

	u := p.scope.Next()
	goutil.AssertNow(t, u.Type() == ast.ReturnStatement, "wrong return type")
	r := u.(*ast.ReturnStatementNode)
	goutil.AssertNow(t, len(r.Results) == 1, "wrong result length")
	goutil.AssertNow(t, r.Results[0].Type() == ast.ArrayLiteral, "wrong result 0 type")
}

func TestMultipleArrayLiteralReturnStatement(t *testing.T) {
	p := createParser(`return []string{"one", "two"}, []Dog{}`)
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)

	u := p.scope.Next()
	goutil.AssertNow(t, u.Type() == ast.ReturnStatement, "wrong return type")
	r := u.(*ast.ReturnStatementNode)
	goutil.AssertNow(t, len(r.Results) == 2, "wrong result length")
	goutil.AssertNow(t, r.Results[0].Type() == ast.ArrayLiteral, "wrong result 0 type")
	goutil.AssertNow(t, r.Results[1].Type() == ast.ArrayLiteral, "wrong result 1 type")
}

func TestSingleMapLiteralReturnStatement(t *testing.T) {
	p := createParser(`return map[string]int{"one":2, "two":3}`)
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)

	u := p.scope.Next()
	goutil.AssertNow(t, u.Type() == ast.ReturnStatement, "wrong return type")
	r := u.(*ast.ReturnStatementNode)
	goutil.AssertNow(t, len(r.Results) == 1, "wrong result length")
	goutil.AssertNow(t, r.Results[0].Type() == ast.MapLiteral, "wrong result 0 type")
}

func TestMultipleMapLiteralReturnStatement(t *testing.T) {
	p := createParser(`return map[string]int{"one":2, "two":3}, map[int]Dog{}`)
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)

	u := p.scope.Next()
	goutil.AssertNow(t, u.Type() == ast.ReturnStatement, "wrong return type")
	r := u.(*ast.ReturnStatementNode)
	goutil.AssertNow(t, len(r.Results) == 2, "wrong result length")
	goutil.AssertNow(t, r.Results[0].Type() == ast.MapLiteral, "wrong result 0 type")
	goutil.AssertNow(t, r.Results[1].Type() == ast.MapLiteral, "wrong result 1 type")
}

func TestSingleCompositeLiteralReturnStatement(t *testing.T) {
	p := createParser(`return Dog{}`)
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)

	u := p.scope.Next()
	goutil.AssertNow(t, u.Type() == ast.ReturnStatement, "wrong return type")
	r := u.(*ast.ReturnStatementNode)
	goutil.AssertNow(t, len(r.Results) == 1, "wrong result length")
	goutil.AssertNow(t, r.Results[0].Type() == ast.CompositeLiteral, "wrong result 0 type")
}

func TestMultipleCompositeLiteralReturnStatement(t *testing.T) {
	p := createParser(`return Cat{name:"Doggo", age:"Five"}, Dog{name:"Katter"}`)
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)

	u := p.scope.Next()
	goutil.AssertNow(t, u.Type() == ast.ReturnStatement, "wrong return type")
	r := u.(*ast.ReturnStatementNode)
	goutil.AssertNow(t, len(r.Results) == 2, fmt.Sprintf("wrong result length: %d", len(r.Results)))
	goutil.AssertNow(t, r.Results[0].Type() == ast.CompositeLiteral, "wrong result 0 type")
	goutil.AssertNow(t, r.Results[1].Type() == ast.CompositeLiteral, "wrong result 1 type")
}

func TestSimpleLiteralAssignmentStatement(t *testing.T) {
	p, _ := ParseString("x = 5")
	n := p.Next()
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	a := n.(*ast.AssignmentStatementNode)
	goutil.AssertNow(t, len(a.Left) == 1, "should be one left value")
	goutil.AssertNow(t, a.Left[0].Type() == ast.Identifier, "wrong left type")
}

func TestIncrementReferenceAssignmentStatement(t *testing.T) {
	p, _ := ParseString("x++")
	n := p.Next()
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	a := n.(*ast.AssignmentStatementNode)
	goutil.AssertNow(t, len(a.Left) == 1, "should be one left value")
	goutil.AssertNow(t, a.Left[0].Type() == ast.Identifier, "wrong left type")
}

func TestMultiToSingleLiteralAssignmentStatement(t *testing.T) {
	p, _ := ParseString("x, y = 5")

	n := p.Next()
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	a := n.(*ast.AssignmentStatementNode)
	goutil.AssertNow(t, len(a.Left) == 2, "should be two left values")
	goutil.AssertNow(t, a.Left[0].Type() == ast.Identifier, "wrong left type")
}

func TestIndexReferenceAssignmentStatement(t *testing.T) {
	a, _ := ParseString("voters[chairperson].weight = 1")
	goutil.AssertNow(t, a != nil, "ast should not be nil")
	n := a.Next()
	goutil.AssertNow(t, len(a.Sequence) == 1, fmt.Sprintf("wrong sequence length: %d", len(a.Sequence)))
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong type")
}

func TestCompositeLiteralAssignmentStatement(t *testing.T) {
	a, _ := ParseString(`proposals = append(proposals, Proposal{
		name: proposalNames[i],
		voteCount: 0,
	})`)
	n := a.Next()
	goutil.AssertNow(t, len(a.Sequence) == 1, fmt.Sprintf("wrong sequence length: %d\n", len(a.Sequence)))
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong type")
}

func TestMultiLiteralAssignmentStatement(t *testing.T) {
	p, _ := ParseString("x, y = 5, 3")

	n := p.Next()
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	a := n.(*ast.AssignmentStatementNode)
	goutil.AssertNow(t, len(a.Left) == 2, "should be two left values")
}

func TestSimpleReferenceAssignmentStatement(t *testing.T) {
	p, _ := ParseString("x = a")

	n := p.Next()
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	a := n.(*ast.AssignmentStatementNode)
	goutil.AssertNow(t, len(a.Left) == 1, "should be 1 left value")
}

func TestMultiToSingleReferenceAssignmentStatement(t *testing.T) {
	p, _ := ParseString("x, y = a")

	n := p.Next()
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	a := n.(*ast.AssignmentStatementNode)
	goutil.AssertNow(t, len(a.Left) == 2, "should be two left values")
}

func TestMultiReferenceAssignmentStatement(t *testing.T) {
	p, _ := ParseString("x, y = a, b")

	n := p.Next()
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	a := n.(*ast.AssignmentStatementNode)
	goutil.AssertNow(t, len(a.Left) == 2, "should be two left values")
}

func TestSimpleCallAssignmentStatement(t *testing.T) {
	p, _ := ParseString("x = a()")

	n := p.Next()
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	a := n.(*ast.AssignmentStatementNode)
	goutil.AssertNow(t, len(a.Left) == 1, "should be one left values")
}

func TestMultiToSingleCallAssignmentStatement(t *testing.T) {
	p, _ := ParseString("x, y = ab()")

	n := p.Next()
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	a := n.(*ast.AssignmentStatementNode)
	goutil.AssertNow(t, len(a.Left) == 2, "should be two left values")
}

func TestMultiCallAssignmentStatement(t *testing.T) {
	p, _ := ParseString("x, y = a(), b()")
	n := p.Next()
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	a := n.(*ast.AssignmentStatementNode)
	goutil.AssertNow(t, len(a.Left) == 2, "should be two left values")
}

func TestSimpleCompositeLiteralAssignmentStatement(t *testing.T) {
	p, _ := ParseString("x = Dog{}")

	n := p.Next()
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	a := n.(*ast.AssignmentStatementNode)
	goutil.AssertNow(t, len(a.Left) == 1, "should be two left values")
}

func TestMultiToSingleCompositeLiteralAssignmentStatement(t *testing.T) {
	p, _ := ParseString("x, y = Dog{}")

	n := p.Next()
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	a := n.(*ast.AssignmentStatementNode)
	goutil.AssertNow(t, len(a.Left) == 2, "should be two left values")
}

func TestMultiCompositeLiteralAssignmentStatement(t *testing.T) {
	p, _ := ParseString("x, y = Dog{}, Cat{}")

	n := p.Next()
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	a := n.(*ast.AssignmentStatementNode)
	goutil.AssertNow(t, len(a.Left) == 2, "should be two left values")
}

func TestSimpleArrayLiteralAssignmentStatement(t *testing.T) {
	p, _ := ParseString("x = []int{3, 5}")

	n := p.Next()
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	a := n.(*ast.AssignmentStatementNode)
	goutil.AssertNow(t, len(a.Left) == 1, "should be 1 left values")
}

func TestMultiToSingleArrayLiteralAssignmentStatement(t *testing.T) {
	p, _ := ParseString("x, y = []int{3, 5}")

	n := p.Next()
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	a := n.(*ast.AssignmentStatementNode)
	goutil.AssertNow(t, len(a.Left) == 2, "should be two left values")
}

func TestMultiArrayLiteralAssignmentStatement(t *testing.T) {
	p, _ := ParseString("x, y = []int{1, 2}, [int]{}")

	n := p.Next()
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	a := n.(*ast.AssignmentStatementNode)
	goutil.AssertNow(t, len(a.Left) == 2, "should be two left values")
}

func TestSimpleMapLiteralAssignmentStatement(t *testing.T) {
	p, _ := ParseString("x = []int{3, 5}")

	n := p.Next()
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	a := n.(*ast.AssignmentStatementNode)
	goutil.AssertNow(t, len(a.Left) == 1, "should be 1 left values")
}

func TestMultiToSingleMapLiteralAssignmentStatement(t *testing.T) {
	p, _ := ParseString("x, y = []int{3, 5}")

	n := p.Next()
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	a := n.(*ast.AssignmentStatementNode)
	goutil.AssertNow(t, len(a.Left) == 2, "should be two left values")
}

func TestMultiMapLiteralAssignmentStatement(t *testing.T) {
	p, _ := ParseString(`x, y = map[int]string{1:"A", 2:"B"}, map[string]int{"A":3, "B": 4}`)

	n := p.Next()
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	a := n.(*ast.AssignmentStatementNode)
	goutil.AssertNow(t, len(a.Left) == 2, "should be two left values")
}

func TestAssignmentStatementSingleAdd(t *testing.T) {
	p, _ := ParseString(`x += 5`)

	n := p.Next()
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	a := n.(*ast.AssignmentStatementNode)
	goutil.AssertNow(t, len(a.Left) == 1, "should be one left value")
}

func TestAssignmentStatementArrayLiteral(t *testing.T) {
	p, errs := ParseString(`x = []string{"a"}`)

	n := p.Next()
	goutil.AssertNow(t, n.Type() == ast.AssignmentStatement, "wrong assignment type")
	a := n.(*ast.AssignmentStatementNode)
	goutil.AssertNow(t, len(a.Left) == 1, "should be one left value")
	goutil.AssertNow(t, errs == nil, errs.Format())
}

func TestImportStatementPath(t *testing.T) {
	p := createParser(`import "dog"`)
	goutil.Assert(t, isImportStatement(p), "should detect import statement")
	parseImportStatement(p)

	n := p.scope.Next()
	goutil.AssertNow(t, n.Type() == ast.ImportStatement, "wrong import type")
	a := n.(*ast.ImportStatementNode)
	goutil.AssertNow(t, a.Path == "dog", "wrong import path value")
}

func TestImportStatementAlias(t *testing.T) {
	p := createParser(`import "dog" as d`)
	goutil.Assert(t, isImportStatement(p), "should detect import statement")
	parseImportStatement(p)

	n := p.scope.Next()
	goutil.AssertNow(t, n.Type() == ast.ImportStatement, "wrong import type")
	a := n.(*ast.ImportStatementNode)
	goutil.AssertNow(t, a.Path == "dog", "wrong import path value")
	goutil.AssertNow(t, a.Alias == "d", "wrong import alias value")
}

func TestPackageStatement(t *testing.T) {
	p := createParser(`package dog at 0.0.1`)
	goutil.Assert(t, isPackageStatement(p), "should detect package statement")
	parsePackageStatement(p)

	n := p.scope.Next()
	goutil.AssertNow(t, n.Type() == ast.PackageStatement, "wrong package type")
	a := n.(*ast.PackageStatementNode)
	goutil.AssertNow(t, a.Name == "dog", "wrong package name value")
}

func TestForEachStatement(t *testing.T) {
	p := createParser(`for x, y in a {}
	`)
	goutil.Assert(t, isForEachStatement(p), "should detect for statement")
	parseForEachStatement(p)
	goutil.Assert(t, p.errs == nil, p.errs.Format())
	n := p.scope.Next()
	goutil.AssertNow(t, len(p.scope.Sequence) == 1, fmt.Sprintf("wrong sequence len: %d", len(p.scope.Sequence)))
	goutil.AssertNow(t, n.Type() == ast.ForEachStatement, "wrong node type")
	a := n.(*ast.ForEachStatementNode)
	goutil.AssertNow(t, len(a.Variables) == 2, "wrong var length")
}

func TestDeclaredForEachStatement(t *testing.T) {
	a, errs := ParseString(`
		a = []string{"a", "b"}

		for x, y in a {}
	`)

	goutil.Assert(t, errs == nil, errs.Format())
	goutil.AssertNow(t, len(a.Sequence) == 2, fmt.Sprintf("wrong sequence len: %d", len(a.Sequence)))
	a.Next()
	n := a.Next()

	goutil.AssertNow(t, n.Type() == ast.ForEachStatement, "wrong node type")
	p := n.(*ast.ForEachStatementNode)
	goutil.AssertNow(t, len(p.Variables) == 2, "wrong var length")
	goutil.AssertNow(t, p.Producer.Type() == ast.Identifier, "wrong producer")
}

func TestParseCaseStatementDouble(t *testing.T) {
	a, errs := ParseString(`
	case x > 5:
		break
	case x == 10:
		break
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
	goutil.AssertNow(t, a != nil, "nil scope")
	goutil.AssertLength(t, len(a.Sequence), 2)
	one := a.Sequence[0]
	goutil.AssertNow(t, one.Type() == ast.CaseStatement, "1 not case statement")
	c1 := one.(*ast.CaseStatementNode)
	goutil.AssertLength(t, len(c1.Block.Sequence), 1)
	two := a.Sequence[1]
	goutil.AssertNow(t, two.Type() == ast.CaseStatement, "2 not case statement")
	c2 := two.(*ast.CaseStatementNode)
	goutil.AssertLength(t, len(c2.Block.Sequence), 1)
}

func TestParseSimpleAssignmentStatement(t *testing.T) {
	a, errs := ParseString(`
		if x = 0; x > 5 {

		}
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
	goutil.AssertNow(t, a != nil, "nil scope")
	goutil.AssertLength(t, len(a.Sequence), 1)
	one := a.Sequence[0]
	goutil.AssertNow(t, one.Type() == ast.IfStatement, "1 not if statement")
	i := one.(*ast.IfStatementNode)
	goutil.AssertNow(t, i.Init != nil, "nil init")
	goutil.AssertNow(t, i.Init.Type() == ast.AssignmentStatement, "not assignment statement")
	as := i.Init.(*ast.AssignmentStatementNode)
	goutil.AssertLength(t, len(as.Left), 1)
	goutil.AssertLength(t, len(as.Right), 1)
	goutil.Assert(t, as.Left[0] != nil, "left is nil")
	goutil.Assert(t, as.Right[0] != nil, "right is nil")
}

func TestParseOrIfCondition(t *testing.T) {
	a, errs := ParseString(`
		if a == 0 or b == 0 {
			return 0
		}
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
	goutil.AssertNow(t, a != nil, "nil scope")
	goutil.AssertLength(t, len(a.Sequence), 1)
}

func TestGenericStatementPlainType(t *testing.T) {
	a, errs := ParseString(`
		x = new List<string>()
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
	goutil.AssertNow(t, a != nil, "nil scope")
	goutil.AssertLength(t, len(a.Sequence), 1)
}

func TestGenericStatementArrayType(t *testing.T) {
	a, errs := ParseString(`
		x = []List<string> {
			new List<string>([]string{"hi"}),
		}
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
	goutil.AssertNow(t, a != nil, "nil scope")
	goutil.AssertLength(t, len(a.Sequence), 1)
}

func TestMultiLineMapTypeAssignment(t *testing.T) {
	a, errs := ParseString(`
		x = map[int]int{
			5: 4,
			3: 1,
		}
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
	goutil.AssertNow(t, a != nil, "nil scope")
	goutil.AssertLength(t, len(a.Sequence), 1)
}

func TestParseForStatementMultipleSimples(t *testing.T) {
	a, errs := ParseString(`
		for i, j, k = 0, 0, 0; i < j < k; j++ {

		}
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
	goutil.AssertNow(t, a != nil, "nil scope")
	goutil.AssertLength(t, len(a.Sequence), 1)
}

func TestParseMultiplePostInvalid(t *testing.T) {
	a, errs := ParseString(`
		i, j++
	`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
	goutil.AssertNow(t, a != nil, "nil scope")
	goutil.AssertLength(t, len(a.Sequence), 1)
}

func TestParseDecrement(t *testing.T) {
	a, errs := ParseString(`
		i--
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
	goutil.AssertNow(t, a != nil, "nil scope")
	goutil.AssertLength(t, len(a.Sequence), 1)
}

func TestParseEmptyReturn(t *testing.T) {
	a, errs := ParseString(`
		return
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
	goutil.AssertNow(t, a != nil, "nil scope")
	goutil.AssertLength(t, len(a.Sequence), 1)
	r := a.Sequence[0].(*ast.ReturnStatementNode)
	goutil.AssertLength(t, len(r.Results), 0)
}

func TestImportGroup(t *testing.T) {
	_, errs := ParseString(`
		import (
			"dog" as d
			"cat" as c
		)
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestStrangeAssignmentSameLine(t *testing.T) {
	_, errs := ParseString(`x = 7  y = 6`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestStrangeAssignmentWithLineComment(t *testing.T) {
	_, errs := ParseString(`x = 7 // hi
		y = 6`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestStrangeAssignmentWithMultilineComment(t *testing.T) {
	_, errs := ParseString(`x = 7 /* hi */
		y = 6`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}
