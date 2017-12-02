# Validator

VM implementations must conform to the following interface:

```go
type VM interface {
	Traverse(ast.Node) (vmgen.Bytecode, util.Errors)
	Builtins() *ast.ScopeNode
	Primitives() map[string]Type
	Literals() LiteralMap
	Operators() OperatorMap
}
```

Where ```LiteralMap``` is an alias for ```map[lexer.TokenType]func(*Validator, string) Type``` and ```OperatorMap``` is an alias for ```map[lexer.TokenType]func(*Validator, ...Type) Type```.

This interface is how language-specific features are implemented on top of the Guardian core systems.

## Operators

To allow VM implementors to pick and choose the operators defined for their language, all operators must be of the form:

```
func operator(v *validator.Validator, ...validator.Type) validator.Type {

}
```

Where the returned ```Type``` is the type produced by the operator.

To make this simple, Guardian provides a few helper functions:

```go
// does a simple type lookup
func SimpleOperator(name string) OperatorFunc
```
