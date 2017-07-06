
#include "base.h"

CREATE_NODE(basic_literal)();
CREATE_NODE(paren_expr)();
CREATE_NODE(function_literal)();
CREATE_NODE(composite_literal)();
CREATE_NODE(index_expr)();
CREATE_NODE(generic_param)();
CREATE_NODE(slice_expr)();

struct basic_literal_node {
    struct node base;
    struct token token;
};

struct paren_expr_node {
    struct node base;
    struct expression_node* expression;
};

struct function_literal_node {
    struct node base;
    struct func_type_node* type;
    struct block_stat_node* body;
};

struct composite_literal_node {
    struct node base;
    struct expression_node* type;
    struct expression_node** elements;
};

struct index_expr_node {
    struct node base;
    struct expression_node* expression;
    struct expression_node* index;
};

struct generic_param_node {
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

struct bracket_expr_node {

};
