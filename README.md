# Guardian

Guardian is a statically typed, object-oriented programming language for decentralised blockchain applications. Its syntax is primarily derived from [Go](https://golang.org), [Java]() and [Python](), with many of the blockchain-specific constructs drawn (at least in part) from [Solidity](https://github.com/ethereum/solidity) and [Viper](https://github.com/ethereum/viper).

Significantly, Guardian is virtual machine agnostic - the same syntax can be compiled into radically different bytecode for different virtual machines.


| Name | Status/Release |
|:-----------:|:----|
| Guardian Core | In Progress |
| EVM | In Progress |
| NEO VM | Upcoming |
| FireVM | Upcoming |

## Aims

In no particular order, Guardian strives to:

- Be executionally deterministic
- Successfully balance legibility and safety
- Have a rich feature set reminiscent of full OOP languages
- Be easy to learn
- Be capable of supporting bytecode generation for arbitrary stack-based VMs

These aims should be considered not only in the design and implementation language itself, but also by all Guardian tooling and documentation.

## Contracts

A very simple Guardian contract is presented below:

```go
contract Greeter {

    var name string

    external func getName() string {
        return name
    }

    external func setName(n string) {
        this.name = n
    }

}
```

In Guardian, there is no need for library contracts - top-level functions can be imported directly from  packages.

## Packaging and Version Declarations

Guardian uses go-style packaging and importing, so that related constructs can be grouped into logical modules. There is no requirement that each file contain a ```contract```, or that it contain only one ```contract```.

In order for future versions of Guardian to include potentially backwards-incompatible changes, each Guardian file must include a version declaration appended to the package declaration:

```go
package calculator guardian 0.0.1
```

## Importing packages

Guardian packages may be imported using the following syntax:

```go

import "guard"

import (
    "a"
    "b"
    "c"
)

import (
    "d" as dalias
    "e" as ealias
    "f" as falias
)

contract Watcher {

    doThing(){
        // this is a function from the guard package
        guard.watch()
        // this is a function from the 'a' package
        a.sayHi()
        // this is an event from the 'e' package
        ealias.LogEvent()
    }
}

```

## Typing

Guardian is strongly and statically typed, and code which does not conform to these standards will not compile. However, in order to promote brevity, types may be omitted when declaring variables (so long as the compiler can infer them from context).

```go
x = 5      // x will have type int
y = x      // y will have type int
x = "hello" // will not compile (x has type int)
y = 5.5 //  will not compile (y has type int)
```

Common types are declared as follows:
```go
var a int
// type inside a package
var b pkg.Dog
var c map[string]int
var d []string
var e func(string) string
```

### Inheritance

Guardian allows for multiple inheritence, such that the following is a valid class:

```go
class Liger inherits Lion, Tiger {

}
```

In cases where a class inherits two methods with identical names and parameters, the methods will 'cancel' and neither will be available in the subclass.

### Interfaces

Guardian uses java-style interface syntax.

```go
interface Walkable inherits Moveable {
    walk(distance int)
}

class Liger inherits Lion, Tiger is Walkable {

}
```

All types which explicitly implement an interface through the ```is``` keyword may be referenced as that interface type. However, there is no go-like implicit implementation, as it can be confusing and serves no particular purpose in a class-based language.

## Key Features

### Constructors and Destructors

Guardian uses ```constructor``` and ```destructor``` keywords. Each class may contain an arbitrary number of constructors and destructors, provided they have unique signatures.

```go
contract Test {

    constructor(name string){

    }

    destructor(){

    }

}
```

By default, the no-args constructor and destructor will be called.

### Generics

Generics can be specified using Java syntax:

```go
// Purchases can only be related to things which are Sellable
// this will be checked at compile time
contract Purchase<T is Sellable> {

    var item T
    var quantity int

    constructor(item T, quantity int){
        this.item = item
        this.quantity = quantity
    }
}
```

To declare several generics at once, use the ```|``` character:

```
class Dog <T|S|R> {

}
```

Generics can also be restricted:

```
class Dog<T inherits Tiger, Lion | S is Cat> {

}
```




### Iteration

Many languages (such as Go) only provide for randomised map iteration. Clearly, this is not deterministic, as demonstrated by the following example:


```go
myMap["hi"] = 2
myMap["bye"] = 3
count = 0
sum = 0
for k, v in myMap {
    sum += v * count
    count++
}
```

The value of sum will be radically different based on the iteration order of the map elements. In a blockchain context, this means that every node may reach different conclusion about the state of the contract. Guardian maps resolve this issue by guaranteeing that maps will be iterated over in order of insertion.

In Guardian, the above sequence of statements will always produce a result of 3.

### Modifers

Solidity uses access modifiers to control method access. In my opinion, access modifiers can and should be substituted for standard ```require``` statements, or for functions if a condition must be duplicated over several methods.

Solidity:

```go
modifier local(Location _loc) {
    require(_loc == loc);
    _;
}

chat(Location loc, string msg) local {

}
```

Guardian:

```go
enforceLocal(loc Location){
    require(location == loc)
}

chat(loc Location, msg string){
    enforceLocal(loc)
}
```

### Groups

Like Go, Guardian allows for the creation of keyword groups, to simplify declaring variables or types with similar characteristics, and to group logically-related declarations together.

These declarations can be any top-level type, or another group.
```go
class (
    Dog {

    }

    Cat {

    }
)
```

```go
public (
    a string
    b int
    c string
)
```

Groups apply to the first level of declarations reached, so nested groups are possible:

```go
internal (

    func (

        add(a, b int) int {
            return a + b
        }

        sub(a, b int) int {
            return a - b
        }

    )
)
```
