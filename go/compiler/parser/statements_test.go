package parser

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestParseReturnStatementSingleConstant(t *testing.T) {
	p := createParser("return 6")
	goutil.Assert(t, isReturnStatement(p), "should detect return statement")
	parseReturnStatement(p)
}

func TestParseAssignmentStatementSingleConstant(t *testing.T) {
	p := createParser("x = 6")
	goutil.Assert(t, isAssignmentStatement(p), "should detect assignment statement")
	parseAssignmentStatement(p)
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
}

func TestParseForStatementCondition(t *testing.T) {
	p := createParser(`
        for x < 5 {

        }
        `)
	goutil.Assert(t, isForStatement(p), "should detect for statement")
	parseForStatement(p)
}

func TestParseForStatementInitCondition(t *testing.T) {
	p := createParser(`
        for x := 0; x < 5 {

        }
        `)
	goutil.Assert(t, isForStatement(p), "should detect for statement")
	parseForStatement(p)
}

func TestParseForStatementInitConditionStatement(t *testing.T) {
	p := createParser(`
        for x := 0; x < 5; x++ {

        }
        `)
	goutil.Assert(t, isForStatement(p), "should detect for statement")
	parseForStatement(p)
}

func TestParseSwitchStatement(t *testing.T) {
	p := createParser(`
        switch x {

        }
        `)
	goutil.Assert(t, isSwitchStatement(p), "should detect switch statement")
	parseSwitchStatement(p)
}

func TestParseSwitchStatementSingleCase(t *testing.T) {
	p := createParser(`
        switch x {
        case 5:

        }
        `)
	goutil.Assert(t, isSwitchStatement(p), "should detect switch statement")
	parseSwitchStatement(p)
}

func TestParseSwitchStatementMultiCase(t *testing.T) {
	p := createParser(`
        switch x {
	        case 5 {
				x += 2
	            break
			}
	        case 4{
				x *= 2
	            break
			}
        }
        `)
	goutil.Assert(t, isSwitchStatement(p), "should detect switch statement")
	parseSwitchStatement(p)
}

func TestParseSwitchStatementExclusive(t *testing.T) {
	p := createParser(`
        exclusive switch x {

        }
        `)
	goutil.Assert(t, isSwitchStatement(p), "should detect switch statement")
	parseSwitchStatement(p)
}

func TestParseCaseStatementSingle(t *testing.T) {
	p := createParser(`
        case 5 {

		}
        `)
	goutil.Assert(t, isCaseStatement(p), "should detect case statement")
	parseCaseStatement(p)
}

func TestParseCaseStatementMultiple(t *testing.T) {
	p := createParser(`case 5, 8, 9 {

		}`)
	goutil.Assert(t, isCaseStatement(p), "should detect case statement")
	parseCaseStatement(p)
}
