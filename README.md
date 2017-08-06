# Guardian

Guardian is a statically typed, object-oriented programming language, with a particular focus on providing a new and more complete set of tools for blockchain applications. Its syntax is primarily derived from Java and Go, with many of the blockchain-specific constructs drawn (at least in part) from [Solidity(https://github.com/ethereum/solidity)].

## Aims

In no particular order, the Guardian programming language strives to:

- Be executionally deterministic
- Be memory and void safe
- Be supportive of deterministic concurrency
- Have a rich feature set reminiscent of full OOP languages
- Be simple and easy to learn

## Contracts

In simple terms, contracts may be understood. There are numerous parallels between . In Guardian, contracts are analogous to top-level classes, within which any number.

## Packaging

Guardian uses go-style packaging and importing.

## Version Declarations

In order for future versions of Guardian to include potentially backwards-incompatible changes, each Guardian contract must include a version declaration appended to the package declaration:

```go
package maths @ 0.0.1
```

If you select an older Guardian version to run your , nodes which do not have

## Exported Variables

Guardian does not use ```public``` or ```private``` modifiers. Functions and Variables may be exported, as in Go, through giving their identifier a capital letter. Other contracts may only interact with exported declarations.

```go
contract Test {

    class Dog{} // exported
    class cat{} // not exported
}
```

However, Guardian does not require exported declarations to be commented, or enforce any sort of syntax requirement on those comments (just comment where needed and avoid redundancy).

## Imports

Guardian contracts may be imported using the following syntax:

```go

import "guard"

contract Watcher {

    guard(){
        Guard() // this is a function from guard.grd
    }
}

```

## Typing

Guardian is strongly and statically typed, and code which does not conform to these standards will not compile. However, in order to promote brevity, types may be omitted when declaring variables (so long as the compiler can infer them from context).

```go
x := 5      // x will have type int
y := x      // y will have type int
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

Similarly to Solidity, data about the current block may be extracted from the environment using a series of builtins.

Builtin variables:

| Guardian | Solidity | EVM  | FireVM |
| ------------- |-------------| -----| ---|
| block.coinbase | block.coinbase | COINBASE | COINBASE|
| block.difficulty | block.difficulty | DIFFICULTY | DIFFICULTY|
| block.gasLimit | block.gaslimit | GASLIMIT | FUELLIMIT |
| block.number | block.number | BLOCKNUMBER | BLOCK |
| block.timestamp | block.timestamp | TIMESTAMP | TIMESTAMP  |
| call.data | msg.data | CALLDATA | CALLDATA |
| call.gas | msg.gas | GAS | FUEL |
| call.caller | msg.sender | CALLER | CALLER |
| call.sig | msg.sig | H | H |
| cn.gasPrice | tx.gasprice | GASPRICE | FUELPRICE |
| cn.origin | tx.origin | ORIGIN | ORIGIN |
| this | this | |


Builtin functions:

| Guardian | Solidity | EVM  | FireVM |
| ------------- |-------------| -----|---|
| +% | addmod() | ADDMOD | ADD, MOD |
| *% | mulmod() | MULMOD | MUL, MOD |
| keccak() | keccak() | | |
| sha3() | sha3() |  |  |
| sha256() | sha256() | |   |
| ripemd160() | ripemd160() | | |
| ecrecover() | ecrecover() |  |  |
| terminate() | selfdestruct() | SELFDESTRUCT | TERMINATE |

Address functions:

| Guardian | Solidity | EVM  | FireVM |
| ------------- |-------------| -----| ---|
| .balance() | .balance() | BALANCE | |
| .transfer() | .transfer() | | |
| .send() | .send() | | |
| .call() | .callcode() |  |  |
| sha256() | sha256() | |   |
| ripemd160() | ripemd160() | | |
| ecrecover() | ecrecover() |  |  |

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

### Macros and Text-Substitution

Guardian allows text-replacement macros.

```go
macro HI {
    HI
}
```

You can pass arguments to macros as follows:

```go
macro MAX(a, b){
    (a > b) ? a : b
}
```

Braces are only necessary for multi-line macros:

```go
macro MAX(a, b) (a > b) ? a : b
```

To access the value of macro arguments as a string literal, use ```#name```, and to paste components together use ```##name```.

### Randomisation

Randomisation is traditionally a very difficult area for cryptocurrency, due to its inherent conflict with the properties of determinism proposed in the Aims section. The definition of random, then, should be the selection of an object from a group, such that the selection could not be predicted at the moment of contract creation, or at any point prior to attempting to insert the contract into a block, but . Guardian uses an inbuilt random() function .

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

The value of sum will be radically different based on the order in which the map is . In a blockchain context, this means that every node will reach. Guardian maps resolve this issue by guaranteeing that maps will be iterated over in order of insertion.

In Guardian, the above sequence of statements will always produce a result of 3.

### ABI Specification

Both Ethereum and any future cryptocurrency built using fireVM use an Application Binary Interface (ABI) to specify the encoding of instructions.

Ethereum functions may be called in the following manner (example stolen from the Solidity documentation):

0xcdcd77c0: the Method ID. This is derived as the first 4 bytes of the Keccak hash of the ASCII form of the signature baz(uint32,bool).

0x0000000000000000000000000000000000000000000000000000000000000045: the first parameter, a uint32 value 69 padded to 32 bytes
0x0000000000000000000000000000000000000000000000000000000000000001: the second parameter - boolean true, padded to 32 bytes

The total call, therefore, is:

0xcdcd77c000000000000000000000000000000000000000000000000000000000000000450000000000000000000000000000000000000000000000000000000000000001

The same call in FireVM may be expressed as (in hexadeciaml):

01        | size of size bytes
cdcd77c0  | function signature
01        | size of the next parameter
45        | actual next parameter
01        | size of the next parameter
01        | actual next parameter

The representation has been compressed from 68 bytes to 9 bytes.

Of course, if the first n bytes of each contract was used to store the number of bytes used in each size byte, then the theoretical size of each parameter could be increased arbitrarily (but still by a capped amount). Currently, the FireVM uses only the first byte, meaning that each

### Modifers

Solidity uses access modifiers to control method access. In my opinion, access modifiers can and should be substituted for standard ```require``` statements, or for functions if a condition must be duplicated over several methods.

Solidity:

```go
modifier local(Location _loc) {
    require(_loc == loc);
    _;
}

chat(msg string) local {

}
```

Guardian:

```go
enforceLocal(loc Location){
    require(location == loc)
}

chat(msg string){
    enforceLocal(loc)
}
```
