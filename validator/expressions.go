package validator

import (
	"github.com/end-r/guardian/util"

	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/ast"
	"github.com/end-r/guardian/typing"
)

func (v *Validator) validateExpression(node ast.ExpressionNode) {
	switch n := node.(type) {
	case *ast.CallExpressionNode:
		v.validateCallExpression(n)
		break
	case *ast.MapLiteralNode:
		v.validateMapLiteral(n)
		break
	case *ast.ArrayLiteralNode:
		v.validateArrayLiteral(n)
		break
	case *ast.FuncLiteralNode:
		v.validateFuncLiteral(n)
		break
	case *ast.KeywordNode:
		v.validateKeywordNode(n)
		break
		// all others are simply to assist in testing etc.
	}
}

func (v *Validator) validateArrayLiteral(n *ast.ArrayLiteralNode) {
	value := v.validateType(n.Signature.Value)
	for _, val := range n.Data {
		valueType := v.validateType(val)
		if typing.AssignableTo(value, valueType, false) {
			v.addError(val.Start(), errInvalidArrayLiteralValue, typing.WriteType(valueType), typing.WriteType(value))
		}
	}
}

func (v *Validator) validateMapLiteral(n *ast.MapLiteralNode) {
	key := v.validateType(n.Signature.Key)
	value := v.validateType(n.Signature.Value)
	for k, val := range n.Data {
		keyType := v.validateType(k)
		valueType := v.validateType(val)
		if typing.AssignableTo(key, keyType, false) {
			v.addError(val.Start(), errInvalidMapLiteralKey, typing.WriteType(valueType), typing.WriteType(value))
		}
		if typing.AssignableTo(value, valueType, false) {
			v.addError(val.Start(), errInvalidMapLiteralValue, typing.WriteType(valueType), typing.WriteType(value))
		}
	}
}

func (v *Validator) validateFuncLiteral(n *ast.FuncLiteralNode) {
	toDeclare := make(typing.TypeMap)
	atLocs := make(map[string]util.Location)
	for _, p := range n.Parameters {
		typ := v.validateType(p.DeclaredType)
		for _, i := range p.Identifiers {
			toDeclare[i] = typ
			atLocs[i] = p.Start()
		}
	}
	v.validateScope(n, n.Scope, toDeclare, atLocs)
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
		v.addError(kw.Start(), errInvalidConstructorCall, typing.WriteType(a), typing.WriteType(args))
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
		v.addError(kw.Start(), errInvalidConstructorCall, typing.WriteType(a), typing.WriteType(args))
		break
	}
}

func (v *Validator) validateCallExpression(call *ast.CallExpressionNode) {
	exprType := v.resolveExpression(call.Call)
	args := v.ExpressionTuple(call.Arguments)
	switch a := exprType.(type) {
	case *typing.Func:
		if !typing.AssignableTo(a.Params, args, false) {
			v.addError(call.Start(), errInvalidFuncCall, typing.WriteType(args), typing.WriteType(a))
		}
		break
	case *typing.Event:
		if !typing.AssignableTo(a.Parameters, args, false) {
			v.addError(call.Start(), errInvalidFuncCall, typing.WriteType(args), typing.WriteType(a))
		}
	default:
		v.addError(call.Start(), errInvalidCall, typing.WriteType(exprType))
	}

}
