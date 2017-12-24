
## Guardian --> EVM

Experimental implementation of a compiler from a Guardian AST to EVM bytecode.

### Encoding/ABI

Solidity uses the left-most 4 bytes of the SHA3 hash of the function signature (including parameters).

Guardian currently uses the left-most 4 bytes of the SHA3 hash of the function/variable name.

TODO: does this necessitate some kind of ABI

Solidity's way is (in my opinion) very restrictive because it enforces a more complex standard of.

## Access Modifiers

Solidity uses four function access modifiers, which have the following meanings.

| Name | Restriction |
|:-----------:|:----|
| external | Can only be called from outside the contract |
| internal | Can only be called by this contract and derived contracts |
| public   | Can be called from anywhere |
| private  | Can only be called by this contract |

In Guardian, the ```public``` and ```private``` keywords derive their regular OOP meanings, and so function restrictions must be defined at a VM-level using different keywords.

| Name | Restriction |
|:-----------:|:----|
| external | Can only be called from outside the contract |
| internal | Can only be called from inside this contract |
| global   | Can be called from anywhere |

Note that this means that contracts can be both ```internal``` and ```private``` or ```public```.

Further, there is no default access level for functions. It must be specified by the writer of the contract. This is a deliberate design choice to encourage secure and gas-efficient practices (as reading calldata is cheaper than reading memory) and therefore ```external``` functions are generally preferable to ```global``` functions where available.

## Builtin Functions/Variables

These builtins are passed to the Guardian compiler along with the traverser and prevent the compilation of.

These builtins are largely the same as Solidity, to promote consistency.

Block/Transaction Properties:

| Name | Type | EVM Opcode | Description |
|:-----------:|:----:|:----------:|:----|
| msg.data | string | CALLDATA | complete calldata |
| msg.gas | uint | GAS | remaining gas |
| msg.sender | address | CALLER | sender of the message (current call) |
| msg.sig | bytes4 | N/A | First 4 bytes of msg.data |
| block.timestamp | uint | TIMESTAMP | current block timestamp as seconds since unix epoch |
| block.number | uint | BLOCKNUMBER | current block number |
| block.blockhash(blockNumber uint) | bytes32 | BLOCKHASH | hash of the given block - only works for 256 most recent blocks excluding current |
| block.gaslimit | uint | GASLIMIT | current block's gas limit |
| block.coinbase | address | COINBASE | current block's miner address |
| tx.gasprice | uint | GASPRICE | gas price of the transaction |
| tx.origin | address | ORIGIN | sender of the transaction (full call chain) |

Mathematical/Cryptographic Functions:

| Name | Type | EVM Opcode | Description |
|:-----------:|:----:|:----------:|:----|
| addmod(x, y, k uint) | uint | ADDMOD | compute (x + y) % k where the addition is performed with arbitrary precision and does not wrap around at 2**256 |
| mulmod(x, y, k uint) | uint | MULMOD | compute (x * y) % k where the multiplication is performed with arbitrary precision and does not wrap around at 2**256 |
| keccak256(...) | bytes32 |
| sha256 | | | |
| sha3 | | | alias to keccak256 |
| ripemd160 | | | |
| ecreover(v uint8, hash, r, s bytes32) | address | | |


Address Related:

Guardian treats these as functions rather than properties, to avoid confusion with syntax like ```this.setX()```.

| Name | Type | EVM Opcode | Description |
|:-----------:|:----:|:----------:|:----|
| balance(address) | uint256 | BALANCE |
| transfer(address, amount uint256) | uint | MULMOD |
| send(address, amount uint256) | bool |
| call(address) | bool |
| delegatecall(address) | bool | DELEGATECALL  |

Contract Related:

| Name | Type | EVM Opcode | Description |
|:-----------:|:----:|:----------:|:----|
| this | address | ADDRESS | the current contract address
| selfdestruct(recipient address) | uint256 | SELFDESTRUCT | destroy the current contract, sending its funds to the given Address |


The general structure of a file is as follows:

### Functions

All functions must be able to be called by reference to a pa. This includes:

- ```external``` functions
- ```global``` functions
- ```internal``` functions
- function literals

```go
sig = first 4 bytes of CALLDATA
if sig == 0x0000000...0
    goto function 0
if sig == 0x0000000...1
    goto function 1
STOP // marks the end of the external functions
JUMPDEST // marks the start of the internal functions

STOP
// here are the actual functions
func add(a, b int) int { return a + b }
JUMPDEST
PUSH "hash of a"
MSTORE
PUSH "hash of b"
MSTORE
PUSH "hash of a"
MLOAD
PUSH "hash of b"
MLOAD
ADD
JUMP
// always include a stop so that there can be no bleed between functions
STOP

```

## Loops

Consider the following example:

```go
for i = 0; i < 10; i++ {

}
```

This simple loop is translated as follows:

```go
// init statement
1 | PUSH "hash of i"
2 | PUSH 0
3 | MSTORE

// condition
4 | JUMPDEST
5 | PUSH "hash of i"
6 | MLOAD
7 | PUSH 10
8 | LT

// if condition failed, jump to end of loop
9 | JUMPI 17

// regular loop processes would occur here

// post statement
10 | PUSH "hash of i"
11 | MLOAD
12 | PUSH 1
13 | ADD
14 | PUSH "hash of i"
15 | MSTORE

// jump back to them top of the loop
16 | JUMP 4
17 | JUMPDEST

// continue after the loop

```

## Conditionals

Consider the following example:

```go
if x = 0; x > 5 {

} else if x == 3 {

} else {

}
```

```go
// init
1 | PUSH "hash of i"
2 | PUSH 0
3 | MSTORE

// first condition
4 | PUSH "hash of i"
5 | MLOAD
6 | PUSH 5
7 | LT
// jump to next condition
8 | JUMPI 10

// 2nd if block
9 | JUMP 17

// evaluate second condition
10 | JUMPDEST
11 | PUSH "hash of i"
12 | PUSH 3
13 | EQ
14 | JUMPI 16

// 2nd if block
15 | JUMP 17

// else block
16 | JUMPDEST

// flow continues
17 | JUMPDEST
```

## Switch Statements

General structure:

```
evaluate original expression
for each case:
    evaluate each expression, Compare to result
    execute from there
    break = JUMP to end of most recent break-able construct
```

Consider the following example:

```go
switch x {
case 3:
    break
case 5, 6:
    break
}
```

```go
// evaluate original expression
1 | PUSH "hash of x"
2 | MLOAD

// first case
3 | PUSH 3
4 | EQ
// conditional jump to the next case
5 | JUMPI
// first case code
6 | JUMP // optional break statement

// second case
7 | PUSH 5
8 | EQ
// condition failed, jump to end
9 | JUMPI 12
// 2.2
10 | PUSH 6
11 | EQ
// conditional jump to t
```

This also works for switch expressions without targets:

```go
switch {
case x > 5:
    break
case x == 3:
    break
case x < 2:
    break
}
```

```go

```

## Assignments

General structure:

```go
for i in range len(left):
    evaluate left[i]
    evaluate right[i]
    contextually store it
```

Consider the following examples:

```go
x[6] = 7
```

## Return Statements

Return statements must push all the returned values onto the stack.

```
for each return parameter
    push it onto the stack
terminate function execution
```

```go
return 5, "hi"
```

```go
1 | PUSH 5
2 | PUSH "hi"
// top of the stack
```

## Expressions

### Binary Expressions

```go
// add left expr bytecode
// add right expr bytecode
// add operation
```

```go
1 | PUSH 1
2 | PUSH 4
3 | ADD
```

### Unary Expressions

```go
// add expr bytecode
// add operation
```

```go
1 | PUSH 0x000001
2 | NOT
```

### Call Expressions

### Index Expressions

Find the size of the type of an e

### Slice Expressions

### Identifiers

The data referenced by an identifier is either in storage or in memory. To access a variable, push the

```go
PUSH "hash of i"
```

And then either ```SLOAD```/```MLOAD``` the data at that address.

### Reference Expressions
