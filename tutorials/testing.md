## Testing

Guardian provides the capability for built-in unit tests. Tests must be stored in files with the suffix ```_test.grd```.

```go
package calculator

import "bastion"

func TestAddition(){
    res = calculator.Add(5, 10)
    bastion.Assert(res == 15, "incorrect addition value")
}

func TestSubtraction(){
    res = calculator.Sub(10, 5)
    bastion.Assert(res == 5, "incorrect subtraction value")
}

func TestMultiplication(){
    res = calculator.Mul(10, 5)
    bastion.Assert(res == 50, "incorrect multiplication value")
}

func TestDivision(){
    res = calculator.Sub(10, 5)
    bastion.Assert(res == 2, "incorrect division value")
}
```

These tests can be run using ```guardian test```.
