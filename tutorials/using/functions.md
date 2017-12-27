## Functions

Guardian functions use golang syntax. Note that types are declared after paramter names.

```go
func add(a int, b int) int {
    return a + b
}
```

Parameters of the same type can be grouped together as follows. Guardian functions can return an unlimited number of variables.

```go
func nothing(a, b int) (int, int) {
    return a, b
}

// This makes the following assignment possible
a, b, c, d = nothing(1, 2), nothing(3, 4)
```

Guardian also supports variable arguments:

```go
func extended(a ...int) {
    for x in a {

    }
}
```




To be considered:

- named return types
