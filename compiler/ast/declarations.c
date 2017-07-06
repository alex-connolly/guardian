
#include "declarations.h"

CREATE_NODE(type_decl)(const char* identifier, struct expression_node* type){
    NEW_NODE(node, type_decl);

    node->identifier = identifier;
    node->type = type;
    SET_BASE(node, NODE_TYPE_DECL);

    RETURN_NODE(node);
}

CREATE_NODE(variable_decl)(struct array* vars, struct expression_node* type, struct array* values){
    NEW_NODE(node, variable_decl);

    node->vars = vars;
    node->type = type;
    node->values = values;
    SET_BASE(node, NODE_VALUE_DECL);

    RETURN_NODE(node);
}

CREATE_NODE(import_decl)(const char* tag, const char* path){
    NEW_NODE(node, import_decl);

    node->tag = tag;
    node->path = path;
    SET_BASE(node, NODE_IMPORT_DECL);

    RETURN_NODE(node);
}

CREATE_NODE(interface_decl)(const char* identifier, struct block_stat_node* body, bool is_abstract){
    NEW_NODE(node, interface_decl);

    node->identifier = identifier;
    node->body = body;
    node->is_abstract = is_abstract;
    SET_BASE(node, NODE_INTERFACE_DECL);

    RETURN_NODE(node);
}

CREATE_NODE(func_decl)(struct func_sig_node* type, struct block_stat_node* body){
    NEW_NODE(node, func_decl);

    node->type = type;
    node->body = body;
    SET_BASE(node, NODE_FUNCTION_DECL);

    RETURN_NODE(node);
}

CREATE_NODE(class_decl)(const char* identifier, struct block_stat_node* body, bool is_abstract, bool is_contract){
    NEW_NODE(node, import);

    node->identifier = identifier;
    node->body = body;
    node->is_abstract = is_abstract;
    node->is_contract = is_contract;
    SET_BASE(node, NODE_CLASS_DECL);

    RETURN_NODE(node);
}
