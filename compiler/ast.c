
#include "ast.h"

#define SET_BASE(node,type_tok,_meta)   node->base.tag = tagv; node->base.token = _tok; node->base.meta = _meta
#define SET_CONTEXT(ctxt)                     node->base.context = ctxt

#define NEW_NODE(name) struct name* node = malloc(sizeof(struct name))
#define RETURN_NODE return (struct node*) node

struct node* create_binary_expression_node(struct token op, struct node* left,
        struct node* right, struct node* context) {

	if (!left || !right) {
        return NULL;
	}

    NEW_NODE(node_binary_expression);

	SET_BASE(NODE_BINARY_EXPR, left->token, NULL);

    node->context = context;

	node->op = op;
	node->left = left;
	node->right = right;

	RETURN_NODE;
}

struct node* create_unary_expression_node(struct token op, struct node* operand){

    if (!operand){
        return NULL;
    }

    NEW_NODE(node_unary_expression);

    node->op = op;
    node->operand = operand;

    RETURN_NODE;
}

struct node* create_enum_declaration_node(struct token token, const char* identifier,
        struct table* table, struct node* context){

    NEW_NODE(node_enum_declaration);

    node->context = context;
    node->identifier = identifier;
    node->table = table;

    RETURN_NODE;
}

struct node* create_function_declaration_node(struct token token, const char* identifier,
        struct node** params, struct node* context){

    NEW_NODE(node_function_declaration);

    SET_CONTEXT(context);

    node->identifier = identifier;
    node->params = params;

    RETURN_NODE;
}

struct node* create_variable_declaration_node(struct token token, const char* identifier,
    struct node** declarations, struct node* context){

    NEW_NODE(node_variable_declaration);

    SET_CONTEXT(context);

    //node->type = type;
    node->declarations = declarations;

    RETURN_NODE;
}

struct node* create_struct_declaration_node(struct token token, const char* identifier){

}

CREATE_NODE_METHOD(loop_statement)(struct token tok, struct node* cond, struct node* stmt,
    struct node* expr, struct node* decl){

        NEW_NODE(node_loop_statement);

        SETBASE(node, NODE_LOOP_STAT, token, NULL);
        SETDECL(node, decl);
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

CREATE_NODE_METHOD(package_declaration)(struct token tok, const char* identifier, struct node** declarations, void* meta, struct node* context){

    NEW_NODE(package_declaration);
    node->identifier = identifier;
    node->context = context;
    node->declarations = declarations;
    RETURN_NODE;
}

CREATE_NODE_METHOD(enum_declaration)(struct token tok, const char* identifier, struct symtable* table, void* meta, struct node* context){
    NEW_NODE(enum_declaration);
    node->context = context;
    node->identifier = identifier;
	node->symtable = table;
    RETURN_NODE;
}

CREATE_NODE_METHOD(function_declaration)(struct token tok, const char* identifier, struct node** params,
    struct node_compound_statement* block, void* meta, struct node* context) {

    NEW_NODE(function_declaration);
    node->identifier = identifier;

	SETBASE(node, NODE_FUNCTION_DECL, token, meta);
    node->context = context;
	node->identifier = identifier;
	node->params = params;
	node->block = block;
	node->nlocals = 0;
	node->uplist = NULL;

	RETURN_NODE;
}

CREATE_NODE_METHOD(variable_declation)(gtoken_s token, gtoken_t type, gtoken_t access_specifier, gtoken_t storage_specifier, gnode_r *declarations, void *meta, gnode_t *decl) {

    NEW_NODE(variable_declaration);

	SETBASE(node, NODE_VARIABLE_DECL, token, meta);
    SETDECL(node, decl);
	node->type = type;
	node->access = access_specifier;
	node->storage = storage_specifier;
	node->decls = declarations;

	RETURN_NODE;
}

CREATE_NODE_METHOD(variable)(struct token tok, const char* identifier, const char* annotation_type, struct node* expr, struct node* context) {

    NEW_NODE(variable);

	SETBASE(node, NODE_VARIABLE, token, NULL);
    node->context = context;
	node->identifier = identifier;
	node->annotation_type = annotation_type;
	node->expr = expr;

	RETURN_NODE;
}
