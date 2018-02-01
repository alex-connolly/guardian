package validator

import (
	"fmt"
	"testing"

	"github.com/end-r/guardian/util"

	"github.com/end-r/guardian/typing"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/parser"

	"github.com/end-r/goutil"
)

func TestResolveLiteralExpressionBoolean(t *testing.T) {
	v := NewValidator(NewTestVM())
	p := parser.ParseExpression("true")
	goutil.AssertNow(t, p != nil, "expr should not be nil")
	goutil.AssertNow(t, p.Type() == ast.Literal, "wrong expression type")
	a := p.(*ast.LiteralNode)
	goutil.Assert(t, v.resolveExpression(a).Compare(typing.Boolean()), "wrong true expression type")
}

func TestResolveCallExpression(t *testing.T) {
	v := NewValidator(NewTestVM())
	fn := new(typing.Func)
	fn.Params = typing.NewTuple(typing.Boolean(), typing.Boolean())
	fn.Results = typing.NewTuple(typing.Boolean())
	v.declareVar(util.Location{}, "hello", &typing.Func{
		Params:  typing.NewTuple(),
		Results: typing.NewTuple(typing.Boolean()),
	})
	p := parser.ParseExpression("hello(5, 5)")
	goutil.AssertNow(t, p.Type() == ast.CallExpression, "wrong expression type")
	a := p.(*ast.CallExpressionNode)
	resolved, ok := v.resolveExpression(a).(*typing.Tuple)
	goutil.Assert(t, ok, "wrong base type")
	goutil.Assert(t, len(resolved.Types) == 1, "wrong type length")
	goutil.Assert(t, fn.Results.Compare(resolved), "should be equal")
}

func TestResolveArrayLiteralExpression(t *testing.T) {
	v := NewValidator(NewTestVM())
	v.declareType(util.Location{}, "dog", typing.Unknown())
	p := parser.ParseExpression("[]dog{}")
	goutil.AssertNow(t, p.Type() == ast.ArrayLiteral, "wrong expression type")
	a := p.(*ast.ArrayLiteralNode)
	_, ok := v.resolveExpression(a).(*typing.Array)
	goutil.Assert(t, ok, "wrong base type")
}

func TestResolveArrayLiteralSliceExpressionCopy(t *testing.T) {
	v := NewValidator(NewTestVM())
	v.declareType(util.Location{}, "dog", typing.Unknown())
	p := parser.ParseExpression("[]dog{}[:]")
	goutil.AssertNow(t, p.Type() == ast.SliceExpression, "wrong expression type")
	_, ok := v.resolveExpression(p).(*typing.Array)
	goutil.Assert(t, ok, "wrong base type")
}

func TestResolveArrayLiteralSliceExpressionLower(t *testing.T) {
	v := NewValidator(NewTestVM())
	v.declareType(util.Location{}, "dog", typing.Unknown())
	p := parser.ParseExpression("[]dog{}[6:]")
	goutil.AssertNow(t, p.Type() == ast.SliceExpression, "wrong expression type")
	_, ok := v.resolveExpression(p).(*typing.Array)
	goutil.Assert(t, ok, "wrong base type")
}

func TestResolveArrayLiteralSliceExpressionUpper(t *testing.T) {
	v := NewValidator(NewTestVM())
	v.declareType(util.Location{}, "dog", typing.Unknown())
	p := parser.ParseExpression("[]dog{}[:10]")
	goutil.AssertNow(t, p.Type() == ast.SliceExpression, "wrong expression type")
	_, ok := v.resolveExpression(p).(*typing.Array)
	goutil.Assert(t, ok, "wrong base type")
}

func TestResolveArrayLiteralSliceExpressionBoth(t *testing.T) {
	v := NewValidator(NewTestVM())
	v.declareType(util.Location{}, "dog", typing.Unknown())
	p := parser.ParseExpression("[]dog{}[6:10]")
	goutil.AssertNow(t, p.Type() == ast.SliceExpression, "wrong expression type")
	_, ok := v.resolveExpression(p).(*typing.Array)
	goutil.Assert(t, ok, "wrong base type")
}

func TestResolveMapLiteralExpression(t *testing.T) {
	v := NewValidator(NewTestVM())
	v.declareType(util.Location{}, "dog", typing.Unknown())
	v.declareType(util.Location{}, "cat", typing.Unknown())
	p := parser.ParseExpression("map[dog]cat{}")
	goutil.AssertNow(t, p.Type() == ast.MapLiteral, "wrong expression type")
	m, ok := v.resolveExpression(p).(*typing.Map)
	goutil.AssertNow(t, ok, "wrong base type")
	goutil.Assert(t, m.Key.Compare(typing.Unknown()), fmt.Sprintf("wrong key: %s", typing.WriteType(m.Key)))
	goutil.Assert(t, m.Value.Compare(typing.Unknown()), fmt.Sprintf("wrong val: %s", typing.WriteType(m.Value)))
}

func TestResolveIndexExpressionArrayLiteral(t *testing.T) {
	v := NewValidator(NewTestVM())
	v.declareVar(util.Location{}, "cat", typing.Boolean())
	p := parser.ParseExpression("[]cat{}[0]")
	goutil.AssertNow(t, p.Type() == ast.IndexExpression, "wrong expression type")
	b := p.(*ast.IndexExpressionNode)
	resolved := v.resolveExpression(b)
	typ, _ := v.isTypeVisible("cat")
	goutil.AssertNow(t, resolved.Compare(typ), "wrong expression type")
}

func TestResolveIndexExpressionMapLiteral(t *testing.T) {
	v := NewValidator(NewTestVM())
	v.declareType(util.Location{}, "dog", typing.Unknown())
	v.declareType(util.Location{}, "cat", typing.Unknown())
	p := parser.ParseExpression(`map[dog]cat{}["hi"]`)
	goutil.AssertNow(t, p.Type() == ast.IndexExpression, "wrong expression type")
	typ, _ := v.isTypeVisible("cat")
	ok := v.resolveExpression(p).Compare(typ)
	goutil.AssertNow(t, ok, "wrong type returned")

}

func TestResolveBinaryExpressionSimpleNumeric(t *testing.T) {
	p := parser.ParseExpression("5 + 5")
	goutil.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong expression type")
	b := p.(*ast.BinaryExpressionNode)
	v := NewValidator(NewTestVM())
	resolved := v.resolveExpression(b)
	goutil.AssertNow(t, resolved.Compare(v.SmallestInteger(256, false)), fmt.Sprintf("wrong expression type: %s", typing.WriteType(resolved)))
}

func TestResolveBinaryExpressionConcatenation(t *testing.T) {
	p := parser.ParseExpression(`"a" + "b"`)
	goutil.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong expression type")
	b := p.(*ast.BinaryExpressionNode)
	v := NewValidator(NewTestVM())
	resolved := v.resolveExpression(b)
	typ, _ := v.isTypeVisible("string")
	goutil.AssertNow(t, resolved.Compare(typ), "wrong expression type")
}

func TestResolveBinaryExpressionEql(t *testing.T) {
	p := parser.ParseExpression("5 == 5")
	goutil.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong expression type")
	b := p.(*ast.BinaryExpressionNode)
	v := NewValidator(NewTestVM())
	resolved := v.resolveExpression(b)
	goutil.AssertNow(t, resolved.Compare(typing.Boolean()), "wrong expression type")
}

func TestResolveBinaryExpressionGeq(t *testing.T) {
	p := parser.ParseExpression("5 >= 5")
	goutil.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong expression type")
	b := p.(*ast.BinaryExpressionNode)
	v := NewValidator(NewTestVM())
	resolved := v.resolveExpression(b)
	goutil.AssertNow(t, resolved.Compare(typing.Boolean()), "wrong expression type")
}

func TestResolveBinaryExpressionLeq(t *testing.T) {
	p := parser.ParseExpression("5 <= 5")
	goutil.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong expression type")
	b := p.(*ast.BinaryExpressionNode)
	v := NewValidator(NewTestVM())
	resolved := v.resolveExpression(b)
	goutil.AssertNow(t, resolved.Compare(typing.Boolean()), "wrong expression type")
}

func TestResolveBinaryExpressionLss(t *testing.T) {
	p := parser.ParseExpression("5 < 5")
	goutil.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong expression type")
	b := p.(*ast.BinaryExpressionNode)
	v := NewValidator(NewTestVM())
	resolved := v.resolveExpression(b)
	goutil.AssertNow(t, resolved.Compare(typing.Boolean()), "wrong expression type")
}

func TestResolveBinaryExpressionGtr(t *testing.T) {
	p := parser.ParseExpression("5 > 5")
	goutil.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong b expression type")
	b := p.(*ast.BinaryExpressionNode)
	v := NewValidator(NewTestVM())
	resolved := v.resolveExpression(b)
	goutil.AssertNow(t, resolved.Compare(typing.Boolean()), fmt.Sprintf("wrong resolved expression type: %s", typing.WriteType(resolved)))
}

func TestResolveBinaryExpressionCast(t *testing.T) {
	p := parser.ParseExpression("uint8(5)")
	goutil.AssertNow(t, p.Type() == ast.CallExpression, "wrong expression type")
	v := NewValidator(NewTestVM())
	resolved := v.resolveExpression(p)
	goutil.AssertNow(t, len(v.errs) == 0, v.errs.Format())
	typ, _ := v.isTypeVisible("uint8")
	goutil.AssertNow(t, resolved.Compare(typ), fmt.Sprintf("wrong resolved expression type: %s", typing.WriteType(resolved)))
}

func TestReferenceCallResolution(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class A {
			func b() string {

			}
		}
		var x string
		var a A
		x = a.b()
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestReferenceIdentifierResolution(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class A {
			var b string
		}
		var x string
		var a A
		x = a.b
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestReferenceIndexResolution(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class A {
			var b []int
		}
		var x int
		var a A
		x = a.b[2]
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestReferenceSliceResolution(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class A {
			var b []int
		}
		var x []int
		var a A
		x = a.b[:]
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestAliasedReferenceCallResolution(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class A {
			func c() string {

			}
		}
		type B A
		var x string
		var b B
		x = b.c()
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestAliasedReferenceIdentifierResolution(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class A {
			var c string
		}
		type B A
		var x string
		var b B
		x = b.c
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestAliasedReferenceIndexResolution(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class A {
			var c []int
		}
		type B A
		var x int
		var b B
		x = b.c[2]
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestAliasedReferenceSliceResolution(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class A {
			var c []int
		}
		type B A
		var x []int
		var b B
		x = b.c[:]
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestTripleReferenceCallResolution(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class B {
			func c() string {

			}
		}
		class A {
			var b B
		}
		var x string
		var a A
		x = a.b.c()
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestTripleReferenceIdentifierResolution(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class B {
			var c string
		}
		class A {
			var b B
		}
		var x string
		var a A
		x = a.b.c
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestTripleReferenceIndexResolution(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class B {
			var c []int
		}
		class A {
			var b B
		}
		var x int
		var a A
		x = a.b.c[2]
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestTripleReferenceSliceResolution(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class B {
			var c []int
		}
		class A {
			var b B
		}
		var x []int
		var a A
		x = a.b.c[:]
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestTripleAliasedReferenceCallResolution(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class B {
			func c() string {

			}
		}
		class A {
			var b B
		}
		type Z A
		var x string
		var z Z
		x = z.b.c()
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestTripleAliasedReferenceIdentifierResolution(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class B {
			var c string
		}
		class A {
			var b B
		}
		type Z A
		var x string
		var z Z
		x = z.b.c
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestTripleAliasedReferenceIndexResolution(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class B {
			var c []int
		}
		class A {
			var b B
		}
		type Z A
		var x int
		var z Z
		x = z.b.c[2]
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestTripleAliasedReferenceSliceResolution(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class B {
			var c []int
		}
		class A {
			var b B
		}
		type Z A
		var x []int
		var z Z
		x = z.b.c[:]
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestInvalidThis(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		this
	`)
	goutil.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidThisClassConstructor(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		class Dog {

			var name string

			constructor(n string){
				this.name = n
			}
		}
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidThisContractConstructor(t *testing.T) {
	_, errs := ValidateString(NewTestVM(), `
		contract Dog {

			var name string

			constructor(n string){
				this.name = n
			}
		}
	`)
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestIsValidMapKey(t *testing.T) {
	a := typing.Boolean()
	goutil.AssertNow(t, !isValidMapKey(a), "boolean valid")
	a = typing.Unknown()
	goutil.AssertNow(t, !isValidMapKey(a), "unknown valid")
	a = typing.Invalid()
	goutil.AssertNow(t, !isValidMapKey(a), "invalid valid")
	a = &typing.NumericType{BitSize: 8, Signed: true}
	goutil.AssertNow(t, isValidMapKey(a), "int8 invalid")
	a = &typing.NumericType{BitSize: 8, Signed: false}
	goutil.AssertNow(t, isValidMapKey(a), "uint8 invalid")
	a = &typing.Array{Value: &typing.NumericType{BitSize: 8, Signed: false}}
	goutil.AssertNow(t, isValidMapKey(a), "uint8 array invalid")
	a = &typing.Array{Value: &typing.NumericType{BitSize: 8, Signed: true}}
	goutil.AssertNow(t, isValidMapKey(a), "int8 array invalid")
	b := &typing.Aliased{Alias: "string", Underlying: a}
	goutil.AssertNow(t, isValidMapKey(b), "aliased int8 array invalid")
}

func TestValidateSimpleCast(t *testing.T) {
	exp, errs := ValidateExpression(NewTestVM(), `uint(0)`)
	goutil.AssertNow(t, exp != nil, "exp isn't nil")
	goutil.AssertNow(t, len(errs) == 0, errs.Format())
}
