There are several levels of parser constructs:

Primary:

Class, Interface, Contract, Assignment, If, For

Secondary:

The following contract should be converted into an AST as follows:

```go
contract Empty {

    class Dog{

    }

    start(x, y int){

    }

}
```

ContractDeclarationNode
ClassDeclarationNode, FuncDeclarationNode(ParametersNode(ParameterNode))
