package validator

import "github.com/end-r/guardian/compiler/ast"

// ValidateScope validates an ast...
func ValidateScope(scope ast.ScopeNode) {
	v := new(Validator)
	v.validateScope(scope)
}

func (v *Validator) validateScope(scope ast.ScopeNode) {
	validateTypeDeclarations(v)
}

type Validator struct {
}
