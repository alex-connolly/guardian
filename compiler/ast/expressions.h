
#include "ast.h"

CREATE_NODE(literal)(struct token token, enum literal_type type);
CREATE_NODE(function_literal)(struct func_sig_node* sig);
CREATE_NODE(composite_literal)(struct expression_node* type, struct array* elements);
CREATE_NODE(index_expr)(struct expression_node* expression, struct expression_node* index);
CREATE_NODE(generic_expr)(struct expression_node* generic);
CREATE_NODE(slice_expr)(struct expression_node* expression, struct expression_node* low,
    struct expression_node* high, struct expression_node* max, bool slice_three);

enum literal_type {
    LITERAL_STRING,
    LITERAL_CHAR,
    LITERAL_NUMBER
};

struct literal_node {
    struct node base;
    struct token token;
};

struct function_literal_node {
    struct node base;
    struct func_sig_node* type;
};

struct composite_literal_node {
    struct node base;
    struct expression_node* type;
    struct array* elements;
};

struct index_expr_node {
    struct node base;
    struct expression_node* expression;
    struct expression_node* index;
};

struct generic_expr_node {
    struct node base;
    struct expression_node* generic;
};

struct slice_expr_node {
    struct node base;
    struct expression_node* expression;
    struct expression_node* low;
    struct expression_node* high;
    struct expression_node* max;
    bool slice_three;
};
