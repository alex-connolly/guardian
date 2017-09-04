package parser

import (
	"testing"

	"github.com/end-r/guardian/go/compiler/ast"

	"github.com/end-r/goutil"
)

func TestParseReturnStatementSingleConstant(t *testing.T) {
	p := createParser("return 6")
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)
	goutil.Assert(t, true, "wrong scope type")
}

func TestParseAssignmentStatementSingleConstant(t *testing.T) {
	p := createParser("x = 6")
	goutil.Assert(t, isAssignmentStatement(p), "should detect assignment statement")
	parseAssignmentStatement(p)
	goutil.Assert(t, p.Scope.Type() == ast.AssignmentStatement, "wrong node type")
}

func TestParseIfStatement(t *testing.T) {
	p := createParser(`
        x = 6
        if x > 4 {

        } elif x < 4 {

        } else {

        }
        `)
	goutil.Assert(t, isIfStatement(p), "should detect if statement")
	parseIfStatement(p)
	goutil.Assert(t, p.Scope.Type() == ast.IfStatement, "wrong node type")
}

func TestParseForStatementCondition(t *testing.T) {
	p := createParser(`
        for x < 5 {

        }
        `)
	goutil.Assert(t, isForStatement(p), "should detect for statement")
	parseForStatement(p)
	goutil.Assert(t, p.Scope.Type() == ast.ForStatement, "wrong node type")
}

func TestParseForStatementInitCondition(t *testing.T) {
	p := createParser(`
        for x := 0; x < 5 {

        }
        `)
	goutil.Assert(t, isForStatement(p), "should detect for statement")
	parseForStatement(p)
	goutil.Assert(t, p.Scope.Type() == ast.ForStatement, "wrong node type")
}

func TestParseForStatementInitConditionStatement(t *testing.T) {
	p := createParser(`
        for x := 0; x < 5; x++ {

        }
        `)
	goutil.Assert(t, isForStatement(p), "should detect for statement")
	parseForStatement(p)
	goutil.Assert(t, p.Scope.Type() == ast.ForStatement, "wrong node type")
}

func TestParseSwitchStatement(t *testing.T) {
	p := createParser(`
        switch x {

        }
        `)
	goutil.Assert(t, isSwitchStatement(p), "should detect switch statement")
	parseSwitchStatement(p)
	goutil.Assert(t, p.Scope.Type() == ast.SwitchStatement, "wrong node type")
}

func TestParseSwitchStatementSingleCase(t *testing.T) {
	p := createParser(`
        switch x {
        case 5:

        }
        `)
	goutil.Assert(t, isSwitchStatement(p), "should detect switch statement")
	parseSwitchStatement(p)
	goutil.Assert(t, p.Scope.Type() == ast.SwitchStatement, "wrong node type")
}

func TestParseSwitchStatementMultiCase(t *testing.T) {
	p := createParser(`
        switch x {
        case 5:
            x += 2
            break
        case 4:
            x *= 2
            break
        }
        `)
	goutil.Assert(t, isSwitchStatement(p), "should detect switch statement")
	parseSwitchStatement(p)
	goutil.Assert(t, p.Scope.Type() == ast.SwitchStatement, "wrong node type")
}

func TestParseSwitchStatementExclusive(t *testing.T) {
	p := createParser(`
        exclusive switch x {

        }
        `)
	goutil.Assert(t, isSwitchStatement(p), "should detect switch statement")
	parseSwitchStatement(p)
	goutil.Assert(t, p.Scope.Type() == ast.SwitchStatement, "wrong node type")
}

func TestParseCaseStatementSingle(t *testing.T) {
	p := createParser(`
        case 5 {

		}
        `)
	goutil.Assert(t, isCaseStatement(p), "should detect case statement")
	parseCaseStatement(p)
	goutil.Assert(t, p.Scope.Type() == ast.CaseStatement, "wrong node type")
}

func TestParseCaseStatementMultiple(t *testing.T) {
	p := createParser(`case 5, 8, 9 {
		
		}`)
	goutil.Assert(t, isCaseStatement(p), "should detect case statement")
	parseCaseStatement(p)
	goutil.Assert(t, p.Scope.Type() == ast.CaseStatement, "wrong node type")
}
