# Guardian

Guardian is a statically typed, object-oriented programming language, with a particular focus on providing a new and more complete set of tools for blockchain applications.
## Aims

In no particular order, the Guardian programming language strives to be:

- Executionally deterministic
- Memory safe
- Fully supportive of concurrency
- Simple and easy to learn

## Examples

A full list of example contracts is in the samples folder.

```go
contract SimpleStorage {

    uint storedData

    set(uint x){
        storedData = x
    }

    get() uint {
        return storedData
    }
}
```

## Imports and Exports

Guardian does not use ```public``` or ```private``` modifiers. Functions and Variables may be exported, as in Go, through giving their identifier. However, Guardian does not require exported declarations to be commented (comment where needed and avoid redundancy).

## Typing

Guardian is strongly and statically typed, and code which does not conform to these standards will not compile. However, in order to promote brevity, types may be omitted when declaring variables (so long as the compiler can infer them from context).

```go
x := 5      // x will have type int
y := x      // y will have type int
x = "hello" // will not compile
```

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

###  Concurrency

Concurrency is difficult to implement and acheive in a deterministic language. However, there are certain programs which process data in an arbitrary order, but

### Contracts

In the event of a ```terminate(args)``` call, the contract will be, and the state of all variables will be reset to their input values. This is tantamount to each operation being performed within a transaction.
