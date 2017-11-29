package validator

import (
	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/lexer"
)

func (v *Validator) validateExpression(node ast.ExpressionNode) {
	switch n := node.(type) {
	case ast.CallExpressionNode:
		v.validateCallExpression(n)
		break

	}
}

func (v *Validator) validateCallExpression(call ast.CallExpressionNode) {
	exprType := v.resolveExpression(call.Call)
	args := v.ExpressionTuple(call.Arguments)
	switch a := exprType.(type) {
	case Func:
		if !a.Params.compare(args) {
			v.addError(errInvalidFuncCall, WriteType(a), WriteType(args))
		}
		break
	case Class:
		constructors := a.Lifecycles[lexer.TknConstructor]
		if NewTuple().compare(args) && len(constructors) == 0 {
			return
		}
		for _, c := range constructors {
			if NewTuple(c.Parameters...).compare(args) {
				return
			}
		}
		v.addError(errInvalidConstructorCall, WriteType(a), WriteType(args))
		break
	default:
		v.addError(errInvalidCall, WriteType(exprType))
	}

}
