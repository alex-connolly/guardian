
#include <stdlib.h>
#include <stdbool.h>
#include "../lexer/token.h"
#include "../lexer/types.h"

#define CREATE_NODE(name) struct node* name_node_create

CREATE_NODE(field)();
CREATE_NODE(field_list)();

enum node_type {

    // expressions
    NODE_LITERAL,
    NODE_FUNCTION_LITERAL,
    NODE_COMPOSITE_LITERAL,
    NODE_INDEX_EXPR,
    NODE_SLICE_EXPR,
    NODE_GENERIC_PARAM,

    // declarations
    NODE_CLASS_DECL,
    NODE_INTERFACE_DECL,
    NODE_FUNCTION_DECL,

    // statements
    NODE_FOR_STAT,
    NODE_RETURN_STAT,
    NODE_IF_STAT,
    NODE_BRANCH_STAT,
};

struct node {
    enum node_type type;
};

struct field_node {
    struct node base;
    const char** names;
    struct expression_node* type;
    const char* tag;
};

struct field_list_node {
    struct node base;
    struct field_node** list;
};
