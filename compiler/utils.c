#include "utils.h"

/*
 * Adds an element to an array.
 */
void array_add(void** array, size_t* array_size, void* element, size_t element_size){
    if (*array_size == 0){
        array = malloc(element_size);
    } else if (*array_size == (*array_size & (~(*array_size) + 1))){
        array = realloc(array, 2 * (*array_size) * element_size);
    }
    array[*array_size++] = element;
}
