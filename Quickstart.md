# Guardian Quickstart

## First contract

Let's start by defining a simple contract:

```go
contract Test {

    lastPersonGreeted string

    external sayHi(name string){
        lastPersonGreeted = name
    }

}
```

This contract creates a public (```exported```) function interface ```sayHi```, which anyone can call with one provided ```string``` parameter. Whoever last called this function will have their ```name``` stored as the ```lastPersonGreeted```.
