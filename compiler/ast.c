
#include "ast.h"

#define NEW_NODE(name) struct name* node = malloc(sizeof(struct name))
#define RETURN_NODE return (struct node*) node

CREATE_NODE_METHOD(binary_expression)(struct token op, struct node* left,
        struct node* right, struct node* context) {

	if (!left || !right) {
        return NULL;
	}

    NEW_NODE(binary_expression);

    node->context = context;

	node->op = op;
	node->left = left;
	node->right = right;

	RETURN_NODE;
}

CREATE_NODE_METHOD(unary_expression)(struct token op, struct node* operand){

    if (!operand){
        return NULL;
    }

    NEW_NODE(unary_expression);

    node->op = op;
    node->operand = operand;

    RETURN_NODE;
}

CREATE_NODE_METHOD(loop_statement)(struct token tok, struct node* cond, struct node* stmt,
    struct node* expr, struct node* context){

        NEW_NODE(node_loop_statement);

        node->context = context;
    	node->cond = cond;
    	node->stmt = stmt;
    	node->expr = expr;
    	node->nclose = UINT32_MAX;
    	RETURN_NODE;
}

CREATE_NODE_METHOD(class_declaration)(struct token tok, const char* identifier,
    struct node* super, struct node** interfaces, struct node** declarations, void* meta, struct node* context){

        NEW_NODE(class_declaration);
        node->identifier = identifier;
        node->context = context;
        node->super = super;
        node->interfaces = interfaces;
        node->declarations = declarations;
        RETURN_NODE;
}

CREATE_NODE_METHOD(package_declaration)(struct token tok, const char* identifier,
    struct node** declarations, void* meta, struct node* context){

    NEW_NODE(package_declaration);
    node->identifier = identifier;
    node->context = context;
    node->declarations = declarations;
    RETURN_NODE;
}

CREATE_NODE_METHOD(enum_declaration)(struct token tok, const char* identifier,
    struct symtable* table, void* meta, struct node* context){
    NEW_NODE(enum_declaration);
    node->context = context;
    node->identifier = identifier;
	node->symtable = table;
    RETURN_NODE;
}

CREATE_NODE_METHOD(function_declaration)(struct token tok, const char* identifier, struct node** params,
    struct node_compound_statement* block, void* meta, struct node* context) {

    NEW_NODE(function_declaration);

    node->context = context;
	node->identifier = identifier;
	node->params = params;
	node->block = block;
	node->nlocals = 0;
	node->uplist = NULL;

	RETURN_NODE;
}

CREATE_NODE_METHOD(variable_declation)(struct token token, struct token type,
    struct node** declarations, struct node* context) {

    NEW_NODE(variable_declaration);

    node->context = context;
	node->type = type;
	node->access = access_specifier;
	node->storage = storage_specifier;
	node->decls = declarations;

	RETURN_NODE;
}

CREATE_NODE_METHOD(variable)(struct token tok, const char* identifier,
    struct node* expr, struct node* context) {

    NEW_NODE(variable);

    node->context = context;
	node->identifier = identifier;
	node->annotation_type = annotation_type;
	node->expr = expr;

	RETURN_NODE;
}

CREATE_NODE_METHOD(contract_declaration)(struct token tok, const char* identifier,
    struct node* super, struct node** interfaces, struct node** declarations, void* meta, struct node* context){

    NEW_NODE(contract_declaration);

    node->context = context;
    node->identifier = identifier;
    node->super = super;
    node->interfaces = interfaces;
    node->declarations = declarations;

    RETURN_NODE;
}
