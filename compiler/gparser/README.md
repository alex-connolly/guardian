Note: all parsing must be contextual

```go
x = 5 + 4 // This is meaningless

contract Dog {
    x = 5 + 4 // Contract storage
}

contract Cat {

    constructor(){
        x = 5 + 4 // Contract memory
    }
}
```


There are several levels of gparser constructs:

Primary:

Class, Interface, Contract, Assignment, If, For

Secondary:

The following contract should be converted into an AST as follows:

```go
contract Empty {

    class Dog{

    }

    start(int x, y){

    }

}
```

ContractDeclarationNode
ClassDeclarationNode, FuncDeclarationNode(ParametersNode(ParameterNode))
