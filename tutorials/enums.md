
## Enums

```go
enum Weekday {
    Monday, Tuesday, Wednesday, Thursday, Friday
}

enum Weekend {
    Saturday, Sunday
}
```

Enums can inherit from parents:

```go
enum DayOfWeek inherits Weekday, Weekend {
    
}
```
