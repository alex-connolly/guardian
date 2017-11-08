# Credence

Credence is a formal verifier for the Guardian programming language.

## Invariants

Credence enforces invariant conditions. These conditions are implemented as guardian functions.

@Credence()

```go
contract Basic {

    name string

    external func setName(n string){
        enforce(n != "Alex")
        name = n
    }
}
```
