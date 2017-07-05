
#include "base.h"

CREATE_NODE(basic_literal)();
CREATE_NODE(paren_expr)();
CREATE_NODE(function_literal)();
CREATE_NODE(composite_literal)();
CREATE_NODE(index_expr)();
CREATE_NODE(generic_param)();
CREATE_NODE(slice_expr)();

struct basic_literal_node {
    size_t position;
    struct token token;
};

struct paren_expr_node {
    size_t left_pos;
    struct expression_node* expression;
    size_t right_pos;
};

struct function_literal_node {
    struct func_type_node* type;
    struct block_stat_node* body;
};

struct composite_literal_node {
    struct expression_node* type;
    size_t left_pos;
    struct expression_node** elements;
    size_t right_pos;
};

struct index_expr_node {
    struct expression_node* expression;
    size_t left_pos;
    struct expression_node* index;
    size_t right_pos;
};

struct generic_param_node {
    size_t left_pos; // position of <
    size_t right_pos; // position of >
    struct expression_node* generic;
};

struct slice_expr_node {
    struct expression_node* expression;
    size_t left_pos;
    struct expression_node* low;
    struct expression_node* high;
    struct expression_node* max;
    bool slice_three;
    size_t right_pos;
};

struct bracket_expr_node {

};
