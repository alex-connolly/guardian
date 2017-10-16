package validator

import "github.com/end-r/guardian/compiler/ast"

func (v *Validator) validateCallExpression(call ast.CallExpressionNode) {
	exprType := v.resolveExpression(call.Call)
	args := v.ExpressionTuple(call.Arguments)
	switch exprType.(type) {
	case Func:
		fn := exprType.(Func)
		if !fn.Params.compare(args) {
			v.addError(errInvalidCall, WriteType(fn), WriteType(args))
		}
		break
		// TODO: also handle lifecycle calls
	}
}
