package validator

import (
	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/typing"
)

func (v *Validator) validateExpression(node ast.ExpressionNode) {
	switch n := node.(type) {
	case *ast.CallExpressionNode:
		v.validateCallExpression(n)
		break
	case *ast.KeywordNode:
		v.validateKeywordNode(n)
		break
		// all others are simply to assist in testing etc.
	}
}

func (v *Validator) validateKeywordNode(kw *ast.KeywordNode) {
	t := v.validateType(kw.TypeNode)
	args := v.ExpressionTuple(kw.Arguments)
	switch a := t.(type) {
	case *typing.Class:
		constructors := a.Lifecycles[token.Constructor]
		if typing.NewTuple().Compare(args) && len(constructors) == 0 {
			return
		}
		for _, c := range constructors {
			paramTuple := typing.NewTuple(c.Parameters...)
			if paramTuple.Compare(args) {
				return
			}
		}
		v.addError(errInvalidConstructorCall, typing.WriteType(a), typing.WriteType(args))
		break
	case *typing.Contract:
		constructors := a.Lifecycles[token.Constructor]
		if typing.NewTuple().Compare(args) && len(constructors) == 0 {
			return
		}
		for _, c := range constructors {
			paramTuple := typing.NewTuple(c.Parameters...)
			if paramTuple.Compare(args) {
				return
			}
		}
		v.addError(errInvalidConstructorCall, typing.WriteType(a), typing.WriteType(args))
		break
	}
}

func (v *Validator) validateCallExpression(call *ast.CallExpressionNode) {
	exprType := v.resolveExpression(call.Call)
	args := v.ExpressionTuple(call.Arguments)
	switch a := exprType.(type) {
	case *typing.Func:
		if !typing.AssignableTo(a.Params, args, false) {
			v.addError(errInvalidFuncCall, typing.WriteType(args), typing.WriteType(a))
		}
		break
	default:
		v.addError(errInvalidCall, typing.WriteType(exprType))
	}

}
