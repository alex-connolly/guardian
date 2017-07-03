#include "stack.h"

struct stack* stack_create(size_t element_size){
    struct stack* s = malloc(sizeof(struct stack));
    s->element_size = element_size;
    return s;
}

void* stack_pop(struct stack* stack){
    return stack->data[stack->size--];
}

void stack_push(struct stack* stack, void* element){
    if (stack->size == 0){
        stack->data = malloc(stack->element_size);
    } else if (stack->size == (stack->size & (~stack->size + 1))) {
        stack->data = realloc(stack->data, stack->size * 2 * stack->element_size);
    }
    stack->data[stack->size++] = element;
}

void stack_free(struct stack* stack){
    free(stack);
}
