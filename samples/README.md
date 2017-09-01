
How is Guardian code converted to fireVM bytecode?

Let's imagine a simple for loop:

```go
for i := 0; i < 4096; i++ {
    // do nothing
}
```

First, we set the initial condition:

```
1 | PUSH 1 00
```

Then we create a comparison loop:

```
2 | PUSH 2 1000
3 | DUP 2
4 | LT
5 | ISZERO
6 | PUSH 1 2
7 | PUSH 1 1
8 | ADD
9 |
```

Note that we do not need to mark the top of the loop as a ```JUMPDEST```, as we would in the EVM. These can be added by the translator if we are compiling to EVM bytecode.



Then

But shouldn't a clever compiler be able to work out that our loop doesn't actually do anything? You're right - it should, and Guardian will try its best to do that for you. If a for loop is provably 'pure' - it does not modify (or possibly modify) any state external to itself. Consider the following:

```go
for count, i := 0; i < 1000; i++ {
    count += 2
}
```

Both ```i``` and ```count``` are modified during each iteration over this loop, but are both declared at the start of the loop and do not alter the state outside of this scope. Therefore, the compiler will optimise all the bytecode above to exactly 0 instructions.

But what about this:

```go
count := 0
for i := 0; i < 1000; i++ {
    count += 2
}
// count is never used again
```

Obviously, there are two ways of optimizing such a loop. One could use the knowledge that the value of ```count``` is ultimately immaterial to optimize the entire loop away. However, as this is only observable after the initial generation of the bytecode, it necessitates either:

- a multi-stage compilation process, or
- storing the bytes associated with each loop and clearing them when it falls out of scope

I'll think about it.

The second method of optimisation is realising that a for loop of this type (which adds 2 to a number 1000 times), may be simply replaced with a multiplication (2 * 1000). Again, I'll think about it.
