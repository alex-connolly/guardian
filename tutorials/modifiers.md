## Modifiers

Guardian modifiers are more restricted than those of Solidity, but still provide a crucial mechanism for controlling access to contract state, or defining the properties of declarations. They cannot be created by users, and the preferred translation of Solidity modifiers is as ``enforce``-able functions which are called at the start of the function to which the modifier was previously attached.

Modifiers come in several mutually exclusive sets:

### ```external```/```internal```

The most crucial of the Guardian modifiers, ```external``` defines when a function or variable property on a contract should be exposed to the world. If a function is ```external```, it may be called by any person at any time (but can still be safeguarded using contract logic). It should be noted from the outset that the default state of a Guardian function is ```internal```, so ```internal``` declarations are only used for clarity.

```go
contract Dog {

    internal woofCount int
    internal volume int

    external func woof(){
        woofCount++
    }

    external func volume(){
        volume++
    }

}
```

Of course, having to write out these modifiers might become tedious, so Guardian modifiers can be grouped using golang syntax.

```go
contract Dog {

    internal (
        woofCount int
        volume int
    )

    external (

        func woof(){
            woofCount++
        }

        func volume(){
            volume++
        }
    )

}
```

### ```public```/```private```/```protected```

### ```abstract```/```static```
