
#include <stdlib.h>
#include <stdbool.h>
#include "../lexer/token.h"

#define CREATE_NODE(name) struct node* name_node_create

CREATE_NODE(comment)();
CREATE_NODE(field)();
CREATE_NODE(field_list)();


struct comment_node {
    size_t position;
    char* text;
};

struct field_node {
    struct comment_node* doc;
    const char** names;
    struct expression_node* type;
    const char* tag;
    struct comment_node* comment;
};

struct field_list_node {
    size_t opening;
    struct field_node** list;
    size_t closing;
};
