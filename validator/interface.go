package validator

import (
	"fmt"
	"log"
	"os"
	"strings"

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
func ValidateFile(vm VM, packageScope *TypeScope, name string) (*ast.ScopeNode, util.Errors) {
	if !isGuardianFile(name) {
		e := make(util.Errors, 0)
		e = append(e, util.Error{
			Location: util.Location{
				Filename: name,
			},
			Message: "Not a guardian file",
		})
	}
	a, errs := parser.ParseFile(name)
	if errs != nil {
		return a, errs
	}
	es := Validate(a, packageScope, vm)
	return a, es
}

// ValidatePackage ...
func ValidatePackage(vm VM, path string) (*TypeScope, util.Errors) {
	// open directory
	// for all files in directory
	// 1. enforce that they are from the s
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	list, _ := file.Readdirnames(0) // 0 to read all files and folders
	var pkgScope *TypeScope
	var errors util.Errors
	for _, name := range list {
		if isGuardianFile(name) {
			_, es := ValidateFile(vm, pkgScope, name)
			errors = append(errors, es...)
		}
	}
	return pkgScope, errors
}

func isGuardianFile(name string) bool {
	return strings.HasSuffix(name, ".grd")
}

// ValidateString ...
func ValidateString(vm VM, text string) (*ast.ScopeNode, util.Errors) {
	a, errs := parser.ParseString(text)
	ts := &TypeScope{parent: nil, scope: a}
	es := Validate(a, ts, vm)
	es = append(es, errs...)
	return a, es
}

// Validate ...
func Validate(scope *ast.ScopeNode, typeScope *TypeScope, vm VM) util.Errors {
	v := new(Validator)

	v.scope = typeScope
	v.vm = vm

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
	if v.inFile {
		if len(scope.Sequence) == 0 || scope.Sequence[0].Type() != ast.PackageStatement {
			v.addError(util.Location{Filename: ""}, errNoPackageStatement)
		}
	}
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
	inFile            bool
	packageName       string
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
	finishedImports   bool
	// for passing to imported files
	// don't access properties through this
	vm VM
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

	v.DeclareBuiltinType(util.Location{}, vm.BooleanName(), typing.Boolean())

	v.parseBuiltins()

}

func (v *Validator) parseBuiltins() {
	if v.builtinScope != nil {
		v.isParsingBuiltins = true
		if v.builtinScope.Declarations != nil {
			// order shouldn't matter
			for _, i := range v.builtinScope.Declarations.Array() {
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

func (v *Validator) addError(loc util.Location, err string, data ...interface{}) {
	v.errs = append(v.errs, util.Error{
		Location: loc,
		Message:  fmt.Sprintf(err, data...),
	})
}
