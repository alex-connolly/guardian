
## Guardian --> EVM

Experimental implementation of a compiler from Guardian to EVM bytecode.

### Encoding

Encoding variables and byte signatures. As Guardian enforces the fact that no variables or functions have the same name.

Solidity uses the left-most 4 bytes of the SHA3 hash of the function signature (including parameters).

Guardian currently uses the left-most 4 bytes of the SHA3 hash of the function/variable name. 

The general structure of a file is as follows:

### Functions

All functions must be able to be called by reference to a pa. This includes:

- ```external``` functions
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

```

## Loops

Consider the following example:

```go
for i = 0; i < 10; i++ {

}
```

This simple loop is translated as follows:

```
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
