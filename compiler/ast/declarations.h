
#include "base.h"

CREATE_NODE(type)();
CREATE_NODE(value)();
CREATE_NODE(import)();
CREATE_NODE(gen_decl)();
CREATE_NODE(func_decl)();

struct type_node {
    struct comment_node* doc;
    const char* identifier;
    size_t assign_position;
    struct expression_node* type;
    struct comment_node* comment;
};

struct value_node {
    struct comment_node* doc;
    const char** names;
    struct expression_node* type;
    struct expression_node** values;
    struct comment_node* comment;
};

struct import_node {
    struct comment_node* doc;
    const char* name;
    const char* path;
    struct comment_node* comment;
    size_t end_position;
};

struct gen_decl_node {
    struct comment_node* doc;
    size_t token_position;
    struct token token;
    size_t left_paren_position;
    struct spec_node** specs;
    size_t right_parent_position;
};

struct func_decl_node {
    struct comment_node* doc;
    const char* identifier;
    struct func_type_node* type;
    struct block_stat_node* body;
};
