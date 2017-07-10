package main

func (vm* VM) add(size byte){
    a := vm.Stack.Pop(size)
    b := vm.Stack.Pop(size)
    c := a + b
    vm.Stack.Push(c)
}

func (vm* VM) sub(size byte){
    a := vm.Stack.Pop(size)
    b := vm.Stack.Pop(size)
    c := a - b
    vm.Stack.Push(c)
}

func (vm* VM) mul(size byte){
    a := vm.Stack.Pop(size)
    b := vm.Stack.Pop(size)
    c := a * b
    vm.Stack.Push(c)
}

func (vm* VM) div(size byte){
    a := vm.Stack.Pop(size)
    b := vm.Stack.Pop(size)
    if b == 0 {
        vm.Stack.Push(b)
        return
    }
    c := a / b
    vm.Stack.Push(c)
}

func (vm* VM) push(data []byte){
    vm.Stack.Push(data)
    vm.ProgramCounter += 2
}

func (vm* VM) pop(size byte){
    vm.Stack.Pop(size)
}

func (vm *VM) jump(to byte){

}

func (vm *VM) return() {

}

func (vm *VM) mstore(data []byte){

}

func (vm *VM) mload(size byte){

}

func (vm *VM) params(offset byte){

}
