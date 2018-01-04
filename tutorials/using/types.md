
 ## Types

 Guardian has a robust and flexible type system.

 ### Plain Types

 ```go
alpha.beta

alpha
 ```

 ### Function Types

 Function types can have named parameters. Functions of this type are not required to share the same parameter names: they are merely documentary in nature.

 ```go
func(a, b, c int) (d, e, f string)
 ```

 They can also be left unnamed.

 ```go
 func(int, int, int) (string, string, string)
 ```

### Array Types

```go
[]string
```

### Map Types

```go
map[string]int
```

All of these types are capable of generic parametrization.
