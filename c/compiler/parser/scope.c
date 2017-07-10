
#include "scope.h"

struct scope* scope_create(){
    struct scope* s = malloc(sizeof(struct scope));
    s->map = create_map(sizeof(char*), sizeof(struct node*));
    return s;
}
void scope_add(struct scope* scope, const char* identifier, struct node* node){
    map_add(scope->map, identifier, node);
}

void scope_free(struct scope* scope){
    map_free(scope->map);
    free(scope);
}
