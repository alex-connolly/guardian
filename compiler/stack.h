#pragma once

#include <stdlib.h>


struct stack {
    void** data;
    size_t size;
    size_t element_size;
};

struct stack* stack_create(size_t element_size);
void* stack_pop(struct stack* stack);
void stack_push(struct stack* stack, void* element);
void stack_free(struct stack* stack);
