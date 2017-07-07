
// "firevm"

func add(firevm.Instance vm) {
    size := vm.Current.Parameters["size"]
    x := vm.Stack.Pop(size)
    y := vm.Stack.Pop(size)
    x.add(y)
    vm.Stack.Push(size, x)
}
