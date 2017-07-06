
#include <stdlib.h>
#include <stdbool.h>
#include "../lexer/token.h"
#include "../lexer/types.h"

#define CREATE_NODE(name) struct node* name_node_create

CREATE_NODE(comment)();
CREATE_NODE(field)();
CREATE_NODE(field_list)();

struct node {
    enum node_type type;
};

struct field_node {
    struct node* base;
    const char** names;
    struct expression_node* type;
    const char* tag;
};

struct field_list_node {
    struct node* base;
    struct field_node** list;
};
