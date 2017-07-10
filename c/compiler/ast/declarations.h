
#include "ast.h"

CREATE_NODE(type_decl)(const char* identifier, struct expression_node* type);
CREATE_NODE(variable_decl)(struct array* vars, struct expression_node* type, struct array* values);
CREATE_NODE(import)(const char* tag, const char* path);
CREATE_NODE(interface_decl)(const char* identifier, struct block_stat_node* body, bool is_abstract);
CREATE_NODE(func_decl)(struct func_sig_node* type, struct block_stat_node* body);
CREATE_NODE(class_decl)(const char* identifier, struct block_stat_node* body, bool is_abstract, bool is_contract);

struct type_declnode {
    const char* identifier;
    struct expression_node* type;
};

struct variable_decl_node {
    struct array* names;
    struct expression_node* type;
    struct array* values;
};

struct import_decl_node {
    const char* tag;
    const char* path;
};

struct interface_decl_node {
    const char* identifier;
    struct block_stat_node* body; // func_sig_nodes
};

struct func_decl_node {
    struct func_sig_node* type;
    struct block_stat_node* body;
};

struct class_decl_node {
    const char* identifier;
    bool is_abstract;
    bool is_contract;
    struct block_stat_node* body; // methods/variables etc
};
