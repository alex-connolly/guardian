# Validator

VM implementations must conform to the following interface:

```go
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
```

Where ```LiteralMap``` is an alias for ```map[token.Type]func(*Validator, string) Type``` and ```OperatorMap``` is an alias for ```map[token.Type]func(*Validator, ...Type) Type```.

This interface is how language-specific features are implemented on top of the Guardian core systems.

## Operators

To allow VM implementors to pick and choose the operators defined for their language, all operators must be of the form:

```
func operator(v *validator.Validator, ...validator.Type) validator.Type {

}
```

Where the returned ```Type``` is the type produced by the operator in the given context (the operand types provided).

To make this simple, Guardian provides a few helper functions:

```go
// does a simple type lookup using 'name'
func SimpleOperator(name string) OperatorFunc
// returns the smallest available numeric type
func NumericalOperator() OperatorFunc
// returns the smallest available integer type
func IntegerOperator() OperatorFunc
// returns the smallest available fixed point type
func DecimalOperator() OperatorFunc
```

## Literals

Guardian supports custom functions for converting lexer tokens to Types. While you can only use tokens which are defined in the lexer (see the lexer docs for a full list), each lexer token can be mapped to a function which can produce contextual types (dependent on the content of the string, for example).

All Literal functions must be of the following form:

```go
func literal(v *validator.Validator, data string) validator.Type {

}
```

## Primitives

Primitive types are the fundamental building block of any Guardian VM. Generally speaking, you should only specify numeric types in this map.

```go
var type = &typing.NumericType{Size: 16, Signed: true}
```

## Builtins



```go

```

## Castable and Assignable
