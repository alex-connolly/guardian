
#include "map.h"

int basic_hash(void* key){
    return 0;
}

void map_free(struct map* map){
    free(map);
}

struct map* map_create(size_t key_size, size_t value_size){
    struct map* map = malloc(sizeof(struct map));
    map->size = 0;
    map->hash = basic_hash;
    return map;
}

void map_add(struct map* map, void* key, void* value){
    map->data[map->hash(key)] = value;
    map->size++;
}

void map_remove(struct map* map, void* key){
    int index = map->hash(key);
    if (map->data[index]){
        map->data[index] = NULL;
        map->size--;
    }
}

void* map_get(struct map* map, void* key){
    return map->data[map->hash(key)];
}
