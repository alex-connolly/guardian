
## Annotations

Guardian annotations can be used to specify additional properties of declarations. They are dissimilar to Java in that they are not able to be created in the form of classes, but are added as part of the VM implementation.

There is only one native annotation, ```@Builtin(string)```, which allows for the mapping of builtin properties onto bytecode-generating functions. To create an annotation, all that is required is the following:

```go
func (vm VM) Annotations() []Annotation {
    return []Annotation {
        ParseAnnotation("@Builtin(string)", handleBuiltin),
        ParseAnnotation("@Hello(int, func(int, int))", handleHello)
    }
}

func handleBuiltin(i {}interface, a []*Annotation) {

}

func handleHello(i {}interface, a []*Annotation) {

}

```

Currently, annotations can only accept string parameters, but this is a flexible limitation and is open to revision in future versions.
