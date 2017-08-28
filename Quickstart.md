# Guardian Quickstart

## First contract

Let's start by defining a simple contract:

```go
contract Test {

    lastGreeted string

    export sayHi(name string){
        this.lastGreeted = name
    }

}
```
