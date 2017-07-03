#pragma once

#include <stdint.h>
#include "lexer.h"

enum node_type {
	NODE_LIST_STAT, NODE_COMPOUND_STAT, NODE_LABEL_STAT, NODE_FLOW_STAT,
	NODE_JUMP_STAT, NODE_LOOP_STAT, NODE_EMPTY_STAT,

	NODE_ENUM_DECL, NODE_FUNCTION_DECL, NODE_VARIABLE_DECL, NODE_CLASS_DECL,
	NODE_MODULE_DECL, NODE_VARIABLE, NODE_CONTRACT_DECL,


	NODE_BINARY_EXPR, NODE_UNARY_EXPR, NODE_FILE_EXPR, NODE_LIST_EXPR,
	NODE_LITERAL_EXPR, NODE_IDENTIFIER_EXPR,
	NODE_POSTFIX_EXPR, NODE_KEYWORD_EXPR,

	NODE_CALL_EXPR, NODE_SUBSCRIPT_EXPR, NODE_ACCESS_EXPR
};

struct node {
    enum node_type type;
    uint32_t ref_count;
};

#define CREATE_NODE_METHOD(name) struct node* ##name_node_create

CREATE_NODE_METHOD(struct_declaration)(const char* identifier);
CREATE_NODE_METHOD(function_declaration)(const char* identifier);
CREATE_NODE_METHOD(interface_declaration)(const char* identifier);
CREATE_NODE_METHOD(enum_declaration)(const char* identifier);
CREATE_NODE_METHOD(variable_declaration)();
CREATE_NODE_METHOD(package_declaration)(struct token tok, const char* identifier);
CREATE_NODE_METHOD(contract_declaration)(struct token tok, const char* identifier);

CREATE_NODE_METHOD(binary_expression_node)(struct token op, struct node* left, struct node* right);
CREATE_NODE_METHOD(unary_expression_node)(struct token op, struct node* operand);

CREATE_NODE_METHOD(literal_string)();
CREATE_NODE_METHOD(literal_int)();
CREATE_NODE_METHOD(literal_bool)();
CREATE_NODE_METHOD(literal_float)();

CREATE_NODE_METHOD(loop_statement)(struct token tok, struct node* cond, struct node* stmt, struct node* expr, struct node* decl);
CREATE_NODE_METHOD(label_statement)(struct token tok, struct node* expr, struct node* stmt, struct node* decl);

struct binary_expression_node {
    struct node base;
    struct token op;
    struct node* left;
    struct node* right;
};

struct unary_expression_node {
    struct node base;
    struct token op;
    struct node* operand;
};

struct literal_expression_node {
    struct node base;
    // some marker for the type
    uint32_t length;
    union {
        char* string;
        double d;
        int64_t n64;
    } value;
};

struct class_declaration_node {
    struct node base;
    struct node* context;
    // something about access modifiers
    const char* identifier;
    struct node* super_class;
    struct table* symtable;
    uint32_t num_instance_vars;
    uint32_t num_static_vars;
};

struct enum_declaration_node {
    struct node base;
    struct node* context;
    // something about access modifiers
    struct table* symtable;
    const char* identifier;
};

struct variable_declaration_node {
    struct node base;
    struct node* context;
    // something about access modifiers
    struct node** declarations;
};

struct function_declaration_node {
    struct node base;
    struct node* context;
    // something about access modifiers
    struct table* symtable;
    const char* identifier;
    struct node** params;
    uint16_t num_locals;
    uint16_t num_params;
};

struct loop_statement_node {
    struct node base;
    struct node* condition;
    struct node* statement;
    uint32_t num_close;
};

struct contract_declaration_node {
	struct node base;
	struct node* context;
	const char* identifier;
	struct node** interfaces;
	struct node* super;
	struct node** declarations;
}
