## Testing

Guardian provides the capability for built-in unit tests. Tests must be stored in files with the suffix ```_test.grd```.

```go
package calculator

test func TestAddition(){
    res = calculator.Add(5, 10)
    assert(res == 15, "incorrect addition value")
}

test func TestSubtraction(){
    res = calculator.Sub(10, 5)
    assert(res == 5, "incorrect subtraction value")
}

test func TestMultiplication(){
    res = calculator.Mul(10, 5)
    assert(res == 50, "incorrect multiplication value")
}

test func TestDivision(){
    res = calculator.Sub(10, 5)
    assert(res == 2, "incorrect division value")
}
```

These tests can be run using ```guardian test```.

Unlike Go, tests do not have to be named ```Testxxx```, and the ```test``` modifier is sufficient.
