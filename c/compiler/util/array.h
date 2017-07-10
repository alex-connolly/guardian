
#include <stdlib.h>

struct array {
    size_t element_size;
    size_t size;
    void** data;
};

struct array* array_create(size_t element_size);
void array_add(struct array* array, void* element);
