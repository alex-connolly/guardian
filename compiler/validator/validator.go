package validator

import (
	"github.com/end-r/guardian/compiler/ast"
)

func (v *Validator) validateType(node ast.Node) Type {
	if node == nil {
		return standards[Invalid]
	} else {
		switch node.Type() {
		case ast.PlainType:
			typ := node.(ast.PlainTypeNode)
			return v.requireVisibleType(typ.Names...)
		case ast.MapType:
			ref := node.(ast.MapTypeNode)
			key := v.validateType(ref.Key)
			val := v.validateType(ref.Value)
			return NewMap(key, val)
		case ast.ArrayType:
			ref := node.(ast.ArrayTypeNode)
			val := v.validateType(ref.Value)
			return NewArray(val)
		case ast.FuncType:
			ref := node.(ast.FuncTypeNode)
			var params []Type
			for _, p := range ref.Parameters {
				params = append(params, v.validateType(p))
			}
			var results []Type
			for _, r := range ref.Results {
				results = append(results, v.validateType(r))
			}
			return NewFunc(NewTuple(params...), NewTuple(results...))
		}
	}
	return standards[Invalid]
}
