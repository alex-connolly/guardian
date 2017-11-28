package validator

import (
	"fmt"

	"github.com/end-r/guardian/util"

	"github.com/end-r/guardian/ast"
)

// Validate...
func Validate(scope *ast.ScopeNode, vm VM) util.Errors {
	v := new(Validator)
	v.scope = &TypeScope{
		parent: nil,
		scope:  scope,
	}

	v.importVM(vm)

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
	builtinScope      *ast.ScopeNode
	scope             *TypeScope
	errs              util.Errors
	isParsingBuiltins bool
	literals          LiteralMap
	operators         OperatorMap
	primitives        map[string]Type
	builtinVariables  map[string]Type
}

// TypeScope ...
type TypeScope struct {
	parent    *TypeScope
	scope     *ast.ScopeNode
	variables map[string]Type
	types     map[string]Type
}

func (v *Validator) importVM(vm VM) {
	v.literals = vm.Literals()
	v.operators = vm.Operators()
	v.primitives = vm.Primitives()
	v.builtinScope = vm.Builtins()
	if v.builtinScope != nil {
		v.isParsingBuiltins = true
		if v.builtinScope.Declarations != nil {
			for _, i := range v.builtinScope.Declarations.Map() {
				v.validateDeclaration(i.(ast.Node))
			}
		}
		if v.builtinScope.Sequence != nil {
			for _, node := range v.builtinScope.Sequence {
				v.validate(node)
			}
		}
		v.isParsingBuiltins = false
	}
}

// NewValidator creates a new validator
func NewValidator(vm VM) *Validator {
	v := Validator{
		scope:     new(TypeScope),
		literals:  vm.Literals(),
		operators: vm.Operators(),
	}
	v.scope.scope = new(ast.ScopeNode)

	v.importVM(vm)

	return &v
}

func (v *Validator) addError(err string, data ...interface{}) {
	v.errs = append(v.errs, util.Error{
		LineNumber: 12345,
		Message:    fmt.Sprintf(err, data...),
	})
}
