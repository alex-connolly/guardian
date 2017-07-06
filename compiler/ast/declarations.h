
#include "ast.h"

CREATE_NODE(type)();
CREATE_NODE(value)();
CREATE_NODE(import)();
CREATE_NODE(gen_decl)();
CREATE_NODE(func_decl)();

struct type_node {
    const char* identifier;
    struct expression_node* type;
};

struct value_node {
    const char** names;
    struct expression_node* type;
    struct expression_node** values;
};

struct import_node {
    const char* name;
    const char* path;
};

struct gen_decl_node {
    struct token token;
    struct spec_node** specs;
};

struct func_decl_node {
    struct func_sig_node* type;
    struct block_stat_node* body;
};
