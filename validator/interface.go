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
	// have to resolve as well so that bytecode generators can process
	v.resolveExpression(expr)
	return expr, v.errs
}

// ValidateFile ...
func ValidateFile(vm VM, name string) (*ast.ScopeNode, util.Errors) {
	a, errs := parser.ParseFile(name)
	if errs != nil {
		return a, errs
	}
	es := Validate(a, vm)
	return a, es
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

	v.isParsingBuiltins = false

	v.validateScope(nil, scope)

	return v.errs
}

func (v *Validator) validateScope(context ast.Node, scope *ast.ScopeNode) (types map[string]typing.Type, properties map[string]typing.Type, lifecycles typing.LifecycleMap) {

	ts := &TypeScope{
		context: context,
		parent:  v.scope,
		scope:   scope,
	}

	v.scope = ts

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
	context           typing.Type
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
	context    ast.Node
	scope      *ast.ScopeNode
	lifecycles typing.LifecycleMap
	variables  map[string]typing.Type
	types      map[string]typing.Type
}

func (v *Validator) importVM(vm VM) {
	v.literals = vm.Literals()
	v.operators = operators()
	v.primitives = vm.Primitives()
	v.builtinScope = vm.Builtins()
	v.modifierGroups = defaultGroups
	v.modifierGroups = append(v.modifierGroups, vm.Modifiers()...)

	v.DeclareBuiltinType(vm.BooleanName(), typing.Boolean())

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
		scope: new(TypeScope),
	}
	v.scope.scope = new(ast.ScopeNode)

	v.importVM(vm)

	return &v
}

func (v *Validator) addError(line int, err string, data ...interface{}) {
	v.errs = append(v.errs, util.Error{
		LineNumber: line,
		Message:    fmt.Sprintf(err, data...),
	})
}
