# VMs

Example specifications for generating vm bytecode through the Guardian compiler.

Currently Supported:

Planned:

- EVM
- AVM
- FireVM

## Primitive Types

There a couple of validator types which make this easy:

```go
// a 64-bit signed int (int64)
a := validator.NumericType{size: 64, signed: true, integral: true}
```

Types which rely on primitive types (such as ```string```), or are aliased from primitive types (```byte``` from ```uint8```) should be included as builtin type declarations:

```go
type byte uint8
```

## Builtin Variables/Functions

The simplest way to add builtins is to parse another section of Guardian code into an AST.

```go
func (v MyVM) Builtins() ast.ScopeNode {
    scope, _ := parser.ParseString(`
        type byte uint8
        type address [20]byte
        class BMessage {
            var data []byte
            var sender address
            // unimplemented funcs should be variables
            // will be able to implement them at bytecode generation time
            var getSender func() address
            func implementedGetSender() address {
                return sender
            }
        }
        var msg BMessage
    `)
    return scope
}
```

Note that it is necessary to declare a class type to be able to declare a builtin variable of that type. It will not be possible to declare new variables of this type in the normal flow of the program.

It is not possible to assign to a builtin variable. This may change in the future, but I can't see any reasons why this behaviour should be allowed, particularly when it is a source of confusion for contract authors and a potential attack vector.

## Modifiers

Guardian has a set of default modifiers (see the docs for a full list). It is not possible to overwrite these modifiers, but additional modifiers can be added using the ```ModifierGroups``` interface.

```go
var (
    location = NewModifierGroup("memory", "storage")
    access = NewModifierGroup("external", "internal", "global")
)
```
