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
		// all others are simply to assist in testing etc.
	}
}

func (v *Validator) validateCallExpression(call *ast.CallExpressionNode) {
	exprType := v.resolveExpression(call.Call)
	fullType := v.resolveExpression(call)
	args := v.ExpressionTuple(call.Arguments)
	switch a := exprType.(type) {
	case *typing.Func:
		if !typing.AssignableTo(args, a.Params) {
			v.addError(errInvalidFuncCall, typing.WriteType(args), typing.WriteType(a))
		}
		break
	case *typing.StandardType:
		if a, ok := fullType.(*typing.Class); ok {
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
	default:
		v.addError(errInvalidCall, typing.WriteType(exprType))
	}

}
