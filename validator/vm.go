package validator

import (
	"strconv"

	"github.com/end-r/guardian/ast"

	"github.com/end-r/guardian/token"

	"github.com/end-r/guardian/parser"
	"github.com/end-r/guardian/typing"
	"github.com/end-r/guardian/util"
	"github.com/end-r/vmgen"
)

// SimpleInstruction returns a neat anon func
func SimpleInstruction(name string) BytecodeGenerator {
	return func(vm VM) (a vmgen.Bytecode) {
		a.Add(name)
		return a
	}
}

type ModifierGroup struct {
	Name       string
	Modifiers  []string
	RequiredOn []ast.NodeType
	AllowedOn  []ast.NodeType
	selected   []string
	Maximum    int
}

func (mg *ModifierGroup) allowedOn(t ast.NodeType) bool {
	if mg.AllowedOn == nil {
		return false
	}
	for _, r := range mg.AllowedOn {
		if t == r {
			return true
		}
	}
	return false
}

func (mg *ModifierGroup) requiredOn(t ast.NodeType) bool {
	if mg.RequiredOn == nil {
		return false
	}
	for _, r := range mg.RequiredOn {
		if t == r {
			return true
		}
	}
	return false
}

func (mg *ModifierGroup) reset() {
	mg.selected = nil
}

func (mg *ModifierGroup) has(mod string) bool {
	for _, m := range mg.Modifiers {
		if m == mod {
			return true
		}
	}
	return false
}

var defaultAnnotations = []*typing.Annotation{
	ParseAnnotation("Bytecode", handleBytecode, 1),
}

func ParseAnnotation(name string, f AnnotationFunction, required int) *typing.Annotation {
	a := new(typing.Annotation)
	a.Name = name
	a.Required = required
	return a
}

type AnnotationFunction func(vm VM, params BuiltinParams, a *typing.Annotation)

type BytecodeGenerator func(vm VM) vmgen.Bytecode

type BuiltinParams struct {
	Bytecode *vmgen.Bytecode
}

func handleBytecode(vm VM, params BuiltinParams, a *typing.Annotation) {
	// TODO: check if it's there?
	bg := vm.BytecodeGenerators()[a.Parameters[0]]
	params.Bytecode.Concat(bg(vm))
}

var defaultGroups = []*ModifierGroup{
	&ModifierGroup{
		Name:       "Access",
		Modifiers:  []string{"public", "private", "protected"},
		RequiredOn: nil,
		AllowedOn:  ast.AllDeclarations,
		Maximum:    1,
	},
	&ModifierGroup{
		Name:       "Concreteness",
		Modifiers:  []string{"abstract"},
		RequiredOn: nil,
		AllowedOn:  ast.AllDeclarations,
		Maximum:    1,
	},
	&ModifierGroup{
		Name:       "Instantiability",
		Modifiers:  []string{"static"},
		RequiredOn: nil,
		AllowedOn:  []ast.NodeType{ast.FuncDeclaration, ast.ClassDeclaration},
		Maximum:    1,
	},
	&ModifierGroup{
		Name:       "Testing",
		Modifiers:  []string{"test"},
		RequiredOn: nil,
		AllowedOn:  []ast.NodeType{ast.FuncDeclaration},
		Maximum:    1,
	},
}

// A VM is the mechanism through which all vm-specific features are applied
// to the Guardian typing: bytecode generation, type enforcement etc
type OperatorFunc func(*Validator, []typing.Type, []ast.ExpressionNode) typing.Type
type OperatorMap map[token.Type]OperatorFunc

type LiteralFunc func(*Validator, string) typing.Type
type LiteralMap map[token.Type]LiteralFunc

// A VM is the mechanism through which all vm-specific features are applied
// to the Guardian typing: bytecode generation, type enforcement etc
type VM interface {
	Traverse(ast.Node) (vmgen.Bytecode, util.Errors)
	Builtins() *ast.ScopeNode
	BaseContract() (*ast.ContractDeclarationNode, util.Errors)
	Primitives() map[string]typing.Type
	Literals() LiteralMap
	BooleanName() string
	ValidExpressions() []ast.NodeType
	ValidStatements() []ast.NodeType
	ValidDeclarations() []ast.NodeType
	Modifiers() []*ModifierGroup
	Annotations() []*typing.Annotation
	BytecodeGenerators() map[string]BytecodeGenerator
	Castable(val *Validator, to, from typing.Type, fromExpression ast.ExpressionNode) bool
	Assignable(val *Validator, to, from typing.Type, fromExpression ast.ExpressionNode) bool
}

type TestVM struct {
}

func NewTestVM() TestVM {
	return TestVM{}
}

func LiteralAssignable(left, right typing.Type, fromExpression ast.ExpressionNode) bool {
	if t, ok := typing.ResolveUnderlying(left).(*typing.NumericType); ok {
		if li, ok := fromExpression.(*ast.LiteralNode); ok {
			if li.LiteralType != token.Integer && li.LiteralType != token.Float {
				return false
			}
			hasSign := (li.Data[0] == '-')
			bitLen := len(li.Data)
			if hasSign {
				bitLen--
			}
			integer := li.LiteralType == token.Integer
			if t.AcceptsLiteral(typing.BitsNeeded(bitLen), integer, hasSign) {
				return true
			}
		}
	}
	return false
}

func (v TestVM) Assignable(val *Validator, left, right typing.Type, fromExpression ast.ExpressionNode) bool {
	t, _ := val.isTypeVisible("address")
	if t.Compare(right) {
		switch left.(type) {
		case *typing.Contract:
			return true
		case *typing.Interface:
			return true
		}
	}
	if !typing.AssignableTo(left, right, true) {
		if LiteralAssignable(left, right, fromExpression) {
			return true
		}
		return false
	}
	return true
}

func (v TestVM) Castable(val *Validator, to, from typing.Type, fromExpression ast.ExpressionNode) bool {
	// can cast all addresses to all contracts
	t, _ := val.isTypeVisible("address")
	if t.Compare(from) {
		switch to.(type) {
		case *typing.Contract:
			return true
		case *typing.Interface:
			return true
		}
	}
	// can cast all uints to addresses
	if t.Compare(to) {
		switch a := typing.ResolveUnderlying(from).(type) {
		case *typing.NumericType:
			if !a.Signed {
				// check size
				return true
			}
			if l, ok := fromExpression.(*ast.LiteralNode); ok {
				hasSign := (l.Data[0] == '-')
				if !hasSign {
					return true
				}
			}
		}
	}
	if typing.AssignableTo(to, from, false) {
		return true
	}
	if LiteralAssignable(to, from, fromExpression) {
		return true
	}
	return false
}

func (v TestVM) BaseContract() (*ast.ContractDeclarationNode, util.Errors) {
	s, errs := parser.ParseString(`
		contract Base {
		    var balance uint
			var address address
		}
	`)
	c := s.Declarations.Next().(*ast.ContractDeclarationNode)
	return c, errs
}

func (v TestVM) Annotations() []*typing.Annotation {
	return nil
}

func (v TestVM) BooleanName() string {
	return "bool"
}

func (v TestVM) Traverse(ast.Node) (vmgen.Bytecode, util.Errors) {
	return vmgen.Bytecode{}, nil
}

func (v TestVM) Builtins() *ast.ScopeNode {
	a, _ := parser.ParseFile("test_builtins.grd")
	return a
}

func (v TestVM) Literals() LiteralMap {
	return LiteralMap{
		token.String:  SimpleLiteral("string"),
		token.True:    BooleanLiteral,
		token.False:   BooleanLiteral,
		token.Integer: resolveIntegerLiteral,
		token.Float:   resolveFloatLiteral,
	}
}

func resolveIntegerLiteral(v *Validator, data string) typing.Type {
	x := typing.BitsNeeded(len(data))
	return v.SmallestInteger(x, true)
}

func resolveFloatLiteral(v *Validator, data string) typing.Type {
	// convert to float
	return typing.Unknown()
}

func getIntegerTypes() map[string]typing.Type {
	m := map[string]typing.Type{}
	const maxSize = 256
	const increment = 8
	for i := increment; i <= maxSize; i += increment {
		intName := "int" + strconv.Itoa(i)
		uintName := "u" + intName
		m[uintName] = &typing.NumericType{Name: uintName, BitSize: i, Signed: false, Integer: true}
		m[intName] = &typing.NumericType{Name: intName, BitSize: i, Signed: true, Integer: true}
	}
	return m
}

var mods = []*ModifierGroup{
	&ModifierGroup{
		Name:      "Visibility",
		Modifiers: []string{"external", "internal", "global"},
		AllowedOn: []ast.NodeType{ast.FuncDeclaration},
		Maximum:   1,
	},
	&ModifierGroup{
		Name:      "Indexed",
		Modifiers: []string{"indexed"},
		AllowedOn: []ast.NodeType{ast.EventDeclaration, ast.ExplicitVarDeclaration},
		Maximum:   1,
	},
}

func (v TestVM) Modifiers() []*ModifierGroup {
	return mods
}

func (v TestVM) Primitives() map[string]typing.Type {
	return getIntegerTypes()
}

var m OperatorMap

func operators() OperatorMap {

	if m != nil {
		return m
	} else {
		m = make(OperatorMap)
	}
	m.Add(BooleanOperator, token.Geq, token.Leq,
		token.Lss, token.Neq, token.Eql, token.Gtr)

	m.Add(operatorAdd, token.Add)
	m.Add(BooleanOperator, token.LogicalAnd, token.LogicalOr)

	// numericalOperator with floats/ints
	m.Add(BinaryNumericOperator, token.Sub, token.Mul, token.Div)

	// integers only
	m.Add(BinaryIntegerOperator, token.Shl, token.Shr)

	m.Add(CastOperator, token.As)

	return m
}

func operatorAdd(v *Validator, types []typing.Type, expressions []ast.ExpressionNode) typing.Type {
	switch types[0].(type) {
	case *typing.NumericType:
		return BinaryNumericOperator(v, types, expressions)
	}
	strType, _ := v.isTypeVisible("string")
	if types[0].Compare(strType) {
		return strType
	}
	return typing.Invalid()
}

func (v TestVM) ValidExpressions() []ast.NodeType {
	return ast.AllExpressions
}

func (v TestVM) ValidStatements() []ast.NodeType {
	return ast.AllStatements
}

func (v TestVM) ValidDeclarations() []ast.NodeType {
	return ast.AllDeclarations
}

func (v TestVM) BytecodeGenerators() map[string]BytecodeGenerator {
	return nil
}
