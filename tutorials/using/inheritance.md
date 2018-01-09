
# Inheritance

## Cancellation

Guardian facilitates multiple inheritance through 'cancellation' - where a clash between two properties creates an ambiguity, the properties will both be ellided from the final type and will be inaccessible. Attempts to use cancelled properties will raise specific warnings, and child classes must re-implement any desired functionality. Note that cancellation will occur regardless of the depth of the property within the ancestory of each parent type.

## Classes

A simple example of property cancellation is presented below:

```go
class Lion { var name string }
class Tiger { var name string }
class Liger inherits Lion, Tiger {
    func getName() string {
        return name
    }
}
```

When the ```getName()``` function makes reference to the ``name``` property, the compiler has no way of determining whether it refers to ```Lion.name``` or ```Tiger.name```. In this case, an error will be raised (as the ```Liger``` type has no accessible and non-ambiguous ```name``` property). 

## Enums

Enums can inherit from multiple parents, provided those types are also enums In order to faciliate logical and consistent enumeration, the order in which the inherited types are set out are as follows: all inherited enums, from left to right, and then the properties of the current enum. For example, consider the following:

```go
enum Weekday { Mon, Tue, Wed, Thu, Fri }
enum Weekend { Sat, Sun }
enum Day inherits Weekday, Weekend {}
```

The order of ```Day```s, starting at index ```0```, will be: ```Mon, Tue, Wed, Thu, Fri, Sat, Sun```, with ```Sun``` being equivalent to ```6```. Incrementing past 6 will cause the enum to loop back to the start (7 == Monday). In more formal terms, all operations on enum are performed ```mod $```, where ```$``` is the total length of all properties in this enum and its ancestors.

Now consider an alteration of the final line in the above snippet:
```go
enum Day inherits Weekend, Weekday {}
```

The new order of ```Day```s will be ```Sat, Sun, Mon, Tue, Wed, Thu, Fri```

And as a final example, the following:

```go
enum Day inherits Weekend, Weekday { Holi }
```

would produce ```Sat, Sun, Mon, Tue, Wed, Thu, Fri, Holi```.

Where parent enums share the same property, a cancellation will occur.
