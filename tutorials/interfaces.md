
## Interfaces

Guardian interface types. They should generally be adjectives or participles. Note that the

```
interface Floating {
    getHeight() int
    getPosition() LatLng
}
```

Interfaces can inherit from other interfaces.

```
interface Hovering inherits Floating {
    duration() int
}
```

Guardian does not use duck-typing or implicit implementation for reasons of clarity. To implement an interface, use the ```is``` keyword instead.

```go
// Compiler will error: interface not implemented
class Balloon is Floating {

}

// Compiler will not error: interface fully implemented
class Cloud is Floating {

    func getHeight() int {
        return 1000
    }


    func getPosition() LatLng {
        return LatLng(63.4, 64.4)
    }
}
```
