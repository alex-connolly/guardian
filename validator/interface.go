package validator

import (
	"fmt"

	"github.com/end-r/guardian/util"

	"github.com/end-r/guardian/ast"
)

// Validate...
func Validate(scope *ast.ScopeNode, vm VM) util.Errors {
	v := new(Validator)
	ts := &TypeScope{
		parent: nil,
		scope:  scope,
	}

	for name, typ := range vm.Primitives() {
		v.DeclareType(name, typ)
	}

	builtins := vm.Builtins()

	v.isParsingBuiltins = true

	if builtins.Declarations != nil {
		for _, i := range builtins.Declarations.Map() {
			// add in placeholders for all declarations
			v.validateDeclaration(i.(ast.Node))
		}
	}

	if builtins.Sequence != nil {
		for _, node := range builtins.Sequence {
			v.validate(node)
		}
	}

	v.isParsingBuiltins = false

	v.scope = ts

	v.validateScope(scope)

	return v.errs
}

func (v *Validator) validateScope(scope *ast.ScopeNode) (map[string]Type, map[string]Type) {

	ts := &TypeScope{
		parent: v.scope,
		scope:  scope,
	}

	v.scope = ts

	// just in case
	v.isParsingBuiltins = false

	v.validateDeclarations(scope)

	v.validateSequence(scope)

	types := v.scope.types
	properties := v.scope.variables

	v.scope = v.scope.parent

	return types, properties
}

func (v *Validator) validateDeclarations(scope *ast.ScopeNode) {
	if scope.Declarations != nil {
		// order doesn't matter here
		for _, i := range scope.Declarations.Map() {
			// add in placeholders for all declarations
			v.validateDeclaration(i.(ast.Node))
		}
	}
}

func (v *Validator) validateSequence(scope *ast.ScopeNode) {
	for _, node := range scope.Sequence {
		v.validate(node)
	}
}

func (v *Validator) validate(node ast.Node) {
	if node.Type() == ast.CallExpression {
		v.validateCallExpression(node.(ast.CallExpressionNode))
	} else {
		v.validateStatement(node)
	}
}

// Validator ...
type Validator struct {
	scope             *TypeScope
	errs              util.Errors
	isParsingBuiltins bool
}

// TypeScope ...
type TypeScope struct {
	parent           *TypeScope
	scope            *ast.ScopeNode
	variables        map[string]Type
	types            map[string]Type
	builtinVariables map[string]Type
	builtinTypes     map[string]Type
}

// NewValidator creates a new validator
func NewValidator() *Validator {
	return &Validator{
		scope: new(TypeScope),
	}
}

func (v *Validator) addError(err string, data ...interface{}) {
	v.errs = append(v.errs, util.Error{
		LineNumber: 12345,
		Message:    fmt.Sprintf(err, data...),
	})
}
