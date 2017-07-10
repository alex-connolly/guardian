# Memory Management in Guardian

## Allocating Memory

## Garbage Collection

Guardian uses a standard reference-counting garbage collector. As one of the primary goals of Guardian is executionary determinism, the garbage collector is not probabilistic but instead maintains a fully up-to-date count of variables at all times.
