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
			_, es := ValidateFile(vm, pkgScope, fmt.Sprintf("%s/%s", path, name))
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

	v.vm = vm

	v.importVM(vm)

	v.scope = nil

	v.validateScope(nil, scope)

	return v.errs
}

func (v *Validator) validateBaseContract(node *ast.ContractDeclarationNode) *typing.Contract {
	c := new(typing.Contract)
	c.Name = node.Identifier

	v.validateModifiers(node, node.Modifiers.Modifiers)

	v.validateAnnotations(ast.ContractDeclaration, node.Modifiers.Annotations)

	v.openScope(nil, nil)

	generics := v.validateGenerics(node.Generics)

	var supers []*typing.Contract
	for _, super := range node.Supers {
		t := v.validatePlainType(super)
		if t != typing.Unknown() {
			if c, ok := t.(*typing.Contract); ok {
				supers = append(supers, c)
			} else {
				v.addError(super.Start(), errTypeRequired, makeName(super.Names), "contract")
			}
		}
	}

	var interfaces []*typing.Interface
	for _, ifc := range node.Interfaces {
		t := v.validatePlainType(ifc)
		if t != typing.Unknown() {
			if c, ok := t.(*typing.Interface); ok {
				interfaces = append(interfaces, c)
			} else {
				v.addError(ifc.Start(), errTypeRequired, makeName(ifc.Names), "interface")
			}
		}
	}

	c.Supers = supers
	c.Interfaces = interfaces
	c.Generics = generics
	c.Mods = &node.Modifiers

	node.Resolved = c

	v.validateContractsCancellation(c, supers)

	//v.declareTypeInParent(node.Start(), node.Identifier, contractType)

	c.Types, c.Properties, c.Lifecycles = v.validateScope(node, node.Body)

	v.closeScope()

	v.validateContractInterfaces(node, c)
	return c
}

func (v *Validator) openScope(context ast.Node, scope *ast.ScopeNode) {
	ts := &TypeScope{
		context: context,
		parent:  v.scope,
		scope:   scope,
	}
	v.scope = ts
}

func (v *Validator) closeScope() {
	if v.scope.parent != nil {
		v.scope = v.scope.parent
	}
}

func (v *Validator) validateScope(context ast.Node, scope *ast.ScopeNode) (types map[string]typing.Type, properties map[string]typing.Type, lifecycles typing.LifecycleMap) {

	v.openScope(context, scope)

	v.validateDeclarations(scope)

	v.validateSequence(scope)

	types = v.scope.types
	properties = v.scope.variables
	lifecycles = v.scope.lifecycles

	v.closeScope()

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
		v.resolveCallExpression(node.(*ast.CallExpressionNode))
	} else {
		v.validateStatement(node)
	}
}

// Validator ...
type Validator struct {
	inFile          bool
	packageName     string
	context         typing.Type
	builtinScope    *TypeScope
	scope           *TypeScope
	primitives      typing.TypeMap
	errs            util.Errors
	literals        LiteralMap
	operators       OperatorMap
	modifierGroups  []*ModifierGroup
	finishedImports bool
	baseContract    *typing.Contract
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
	variables  typing.TypeMap
	types      typing.TypeMap
}

func (v *Validator) importVM(vm VM) {
	v.literals = vm.Literals()
	v.operators = operators()
	v.primitives = vm.Primitives()
	v.modifierGroups = defaultGroups
	v.modifierGroups = append(v.modifierGroups, vm.Modifiers()...)

	if v.primitives == nil {
		v.primitives = make(typing.TypeMap)
	}

	v.primitives[vm.BooleanName()] = typing.Boolean()

	v.validateScope(nil, vm.Builtins())

	v.builtinScope = v.scope

	c, errs := vm.BaseContract()
	if errs != nil {
		v.errs = append(v.errs, errs...)
	}

	v.baseContract = v.validateBaseContract(c)

}

// NewValidator creates a new validator
func NewValidator(vm VM) *Validator {
	v := new(Validator)

	v.importVM(vm)

	v.scope = &TypeScope{
		scope: new(ast.ScopeNode),
	}

	return v
}

func (v *Validator) addError(loc util.Location, err string, data ...interface{}) {
	v.errs = append(v.errs, util.Error{
		Location: loc,
		Message:  fmt.Sprintf(err, data...),
	})
}
