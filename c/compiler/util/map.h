
#include <stdlib.h>

struct map {
    size_t size;
    void** data;
    int (*hash)(void* key);
};

struct map* map_create(size_t key_size, size_t value_size);
void map_add(struct map* map, void* key, void* value);
void map_remove(struct map* map, void* key);
void* map_get(struct map* map, void* key);
void map_free(struct map* map);
