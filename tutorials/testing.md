## Testing

Guardian provides the capability for built-in unit tests.

```go
import (
    "bastion"
    "calculator"
)

contract BTester inherits Bastion {

    func TestAddition(){
        res = calculator.Add(5, 10)
        tester.Assert(res == 15, "incorrect addition value")
    }
}
```

These tests can be run using ```guardian test```.
