# Credence

Credence is a formal verifier for the Guardian programming language.

## Invariants

Credence enforces invariant conditions. These conditions are implemented as guardian functions.

```go
contract Basic {

    internal name string

    external (
        name string
        day int
    )

    external (

        @Credence()
        func setName(n string){
            enforce(n != "Alex")
            name = n
        }

        func getName() string {
            return name
        }

    )

}
```
