package validator

import (
	"fmt"

	"github.com/end-r/guardian/parser"

	"github.com/end-r/guardian/typing"

	"github.com/end-r/guardian/util"

	"github.com/end-r/guardian/ast"
)

// ValidateExpression ...
func ValidateExpression(vm VM, text string) (ast.ExpressionNode, util.Errors) {
	expr := parser.ParseExpression(text)
	v := NewValidator(vm)
	v.validateExpression(expr)
	return expr, v.errs
}

// ValidateString ...
func ValidateString(vm VM, text string) (*ast.ScopeNode, util.Errors) {
	a, errs := parser.ParseString(text)
	es := Validate(a, vm)
	es = append(es, errs...)
	return a, es
}

// Validate ...
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

func (v *Validator) validateScope(scope *ast.ScopeNode) (types map[string]typing.Type, properties map[string]typing.Type, lifecycles typing.LifecycleMap) {

	ts := &TypeScope{
		parent: v.scope,
		scope:  scope,
	}

	v.scope = ts

	// just in case
	v.isParsingBuiltins = false

	v.validateDeclarations(scope)

	v.validateSequence(scope)

	types = v.scope.types
	properties = v.scope.variables
	lifecycles = v.scope.lifecycles

	v.scope = v.scope.parent

	return types, properties, lifecycles
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
		v.validateCallExpression(node.(*ast.CallExpressionNode))
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
	primitives        map[string]typing.Type
	builtinVariables  map[string]typing.Type
	modifierGroups    []*ModifierGroup
}

// TypeScope ...
type TypeScope struct {
	parent     *TypeScope
	scope      *ast.ScopeNode
	lifecycles typing.LifecycleMap
	variables  map[string]typing.Type
	types      map[string]typing.Type
}

func (v *Validator) importVM(vm VM) {
	v.literals = vm.Literals()
	v.operators = vm.Operators()
	v.primitives = vm.Primitives()
	v.builtinScope = vm.Builtins()

	v.modifierGroups = defaultGroups
	v.modifierGroups = append(v.modifierGroups, vm.Modifiers()...)

	v.DeclareBuiltinType(vm.BooleanName(), typing.BooleanType{})

	v.parseBuiltins()
}

func (v *Validator) parseBuiltins() {
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
