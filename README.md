# Guardian

Guardian is a statically typed, object-oriented programming language for blockchain applications. Its syntax is primarily derived from [Go](https://golang.org), with many of the blockchain-specific constructs drawn (at least in part) from [Solidity](https://github.com/ethereum/solidity) and [Viper](https://github.com/ethereum/viper).

Significantly, Guardian is (to some degree) virtual machine agnostic - the same syntax can be compiled into radically different bytecode depending on the virtual machine which is being targeted.

## Aims

In no particular order, the Guardian programming language strives to:

- Be executionally deterministic
- Successfully balance legibility and safety
- Have a rich feature set reminiscent of full OOP languages
- Be easy to learn
- Be capable of supporting bytecode generation for arbitrary VMs

These aims should be considered not only in the design and implementation language itself, but also by all Guardian tooling and documentation.

## Contracts

In simple terms, contracts may be understood a. There are numerous parallels between . In Guardian, contracts are analogous to top-level classes, within which any number.

## Packaging and Version Declarations

Guardian uses go-style packaging and importing, so that related constructs can be grouped into logical modules. There is no requirement that each file contain a ```contract```, or that it contain only one ```contract```.

In order for future versions of Guardian to include potentially backwards-incompatible changes, each Guardian file must include a version declaration appended to the package declaration:

```go
package calculator @ 0.0.1
```


## Exported Variables



## Imports

Guardian contracts may be imported using the following syntax:

```go

import "guard"

contract Watcher {

    guard(){
        guard.Guard() // this is a function from guard.grd
    }
}

```

## Typing

Guardian is strongly and statically typed, and code which does not conform to these standards will not compile. However, in order to promote brevity, types may be omitted when declaring variables (so long as the compiler can infer them from context).

```go
x = 5      // x will have type int
y = x      // y will have type int
x = "hello" // will not compile (x has type int)
```

### Inheritance

Guardian allows for multiple inheritence, such that the following is a valid class:

```go
class Liger inherits Lion, Tiger {

}
```

All EXPORTED methods and fields will be available in the subclass. The introduction of a ```protected``` keyword, or some other mechanism for allowing inheritance of imported methods, is currently under consideration.

However, in cases where a class inherits two methods with identical names and parameters, the methods will 'cancel', and a warning will be raised. To ignore this warning, simply annotate the offending class with ```@Ignore("cancellation")``` This sidesteps the diamond inheritance problem.

### Interfaces

Guardian uses java-style interface syntax.

```go
interface Walkable inherits Moveable {
    walk(int distance)
}

class Liger inherits Lion, Tiger is Walkable {

}
```

All types which explicitly implement an interface through the ```is``` keyword may be referenced as that interface type. However, there is no go-like implicit implementation, as it can be confusing and serves no particular purpose in a class-based language.

## Builtins

Similarly to Solidity, data about the current block may be extracted from the environment using a series of builtins. However, Solidity uses syntax reminiscent of object fields (e.g. ```msg.sender```), which are misleading as to the true nature of these builtins, which are immutable. Guardian therefore replaces them with a series of builtin method calls.

## Key Features

### Constructors and Destructors

Guardian uses ```constructor``` and ```destructor``` keywords. Each class may contain an arbitrary number of constructors and destructors, provided they have unique signatures.

```go
contract Test {

    constructor(string name){

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

    T item
    int quantity

    constructor(T item, int quantity){
        this.item = item
        this.quantity = quantity
    }
}
```
### Iteration

Many languages (such as Go) only provide for randomised map iteration. Clearly, this is not deterministic, as demonstrated by the following example:


```go
myMap["hi"] = 2
myMap["bye"] = 3
count := 0
sum := 0
for k, v := range myMap {
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
