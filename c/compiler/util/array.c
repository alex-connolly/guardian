
#include "array.h"


struct array* array_create(size_t element_size){
    struct array* a = malloc(sizeof(struct array));
    a->element_size = element_size;
    return a;
}

void array_add(struct array* array, void* element){
    if (array->size == 0){
        array->data = malloc(array->element_size);
    } else if ((array->size & (~array->size + 1)) == array->size){
        array->data = realloc(array->data, array->element_size * array->size * 2);
    }
    array->data[array->size++] = element;
}
