## Loops

Guardian provides a number of mechanisms for looping over data types.

### Arrays

```go
x = []string{"hi", "hello", "g'day"}
for text in x {

}
for index, text in x {

}
```

### Maps

Guardian's maps are deterministically iterable - the ```in``` iterator will return the values in order of insertion.

```go
x = map[string]int

for key, value in x {

}
```
