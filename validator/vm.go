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
	Primitives() map[string]typing.Type
	Literals() LiteralMap
	BooleanName() string
	ValidExpressions() []ast.NodeType
	ValidStatements() []ast.NodeType
	ValidDeclarations() []ast.NodeType
	Modifiers() []*ModifierGroup
	Annotations() []*typing.Annotation
	BytecodeGenerators() map[string]BytecodeGenerator
}

type TestVM struct {
}

func NewTestVM() TestVM {
	return TestVM{}
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
	a, _ := parser.ParseString(`
		type uint uint256
		type int int256
		type byte int8
		type string []byte
		type address [20]byte

		wei = 1
		kwei = 1000 * wei
		babbage = kwei
		mwei = 1000 * kwei
		lovelace = mwei
		gwei = 1000 * mwei
		shannon = gwei
		microether = 1000 * gwei
		szabo = microether
		milliether = 1000 * microether
		finney = milliether
		ether = 1000 * milliether

		// account functions
		@Builtin("balance") var balance func(a address) uint256
		@Builtin("transfer") var transfer func(a address, amount uint256) uint
		@Builtin("send") var send func(a address, amount uint256) bool
		@Builtin("call") var call func(a address) bool
		@Builtin("delegateCall") var delegateCall func(a address)

		// cryptographic functions
		@Builtin("addmod") var addmod func(x, y, k uint) uint
		@Builtin("mulmod") var mulmod func(x, y, k uint) uint
		@Builtin("keccak256") var keccak256 func()
		@Builtin("sha256") var sha256 func()
		@Builtin("sha3") var sha3 func()
		@Builtin("ripemd160") var ripemd160 func()
		@Builtin("ecrecover") var ecrecover func (v uint8, h, r, s string) address

		// contract functions
		// NO THIS KEYWORD: confusing for most programmers, unintentional bugs etc

		@Builtin("selfDestruct") var selfDestruct func(recipient address) uint256


		class BuiltinMessage {
			@Builtin("calldata") var data []byte
			@Builtin("gas") var gas uint
			@Builtin("caller") var sender address
			@Builtin("signature") var sig [4]byte
		}

		var msg BuiltinMessage

		class BuiltinBlock {
			@Builtin("timestamp") var timestamp uint
			@Builtin("number") var number uint
			@Builtin("coinbase") var coinbase address
			@Builtin("gasLimit") var gasLimit uint
			@Builtin("blockhash") var blockhash func(blockNumber uint) [32]byte
		}

		var block BuiltinBlock

		class BuiltinTransaction {
			@Builtin("gasPrice") var gasPrice uint
			@Builtin("origin") var origin address
		}

		var tx BuiltinTransaction
    `)
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
	return v.SmallestNumericType(x, false)
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

func (v TestVM) Modifiers() []*ModifierGroup {
	return nil
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
	strType := v.getNamedType("string")
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
