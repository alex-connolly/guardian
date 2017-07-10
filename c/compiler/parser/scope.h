
#include "../util/map.h"
#include "../ast/ast.h"
#include "<stdlib.h>"

struct scope {
    struct map* map;
};

struct scope* scope_create();
void scope_add(struct scope* scope, const char* identifier, struct node* node);
void scope_free(struct scope* scope);
