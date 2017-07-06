
#include "expressions.h"

CREATE_NODE(literal_expr)(struct token token, enum literal_type type){
    NEW_NODE(node, literal_expr);

    node->token = token;
    node->type = type;
    SET_BASE(node, NODE_LITERAL);

    RETURN_NODE(node);
}

CREATE_NODE(function_literal_expr)(struct func_sig_node* sig){

    NEW_NODE(node, function_literal_expr);

    node->sig = sig;
    SET_BASE(node, NODE_FUNCTION_LITERAL);

    RETURN_NODE(node);
}

CREATE_NODE(composite_literal_expr)(struct expression_node* type, struct array* elements){
    NEW_NODE(node, composite_literal_expr);

    node->type = type;
    node->elements = elements;
    SET_BASE(node, NODE_COMPOSITE_LITERAL);

    RETURN_NODE(node);
}

CREATE_NODE(index_expr)(struct expression_node* expression, struct expression_node* index){
    NEW_NODE(node, index_expr);

    node->expression = expression;
    node->index = index;
    SET_BASE(node, NODE_INDEX_EXPR);

    RETURN_NODE(node);
}

CREATE_NODE(generic_expr)(struct expression_node* generic){
    NEW_NODE(node, generic_expr);

    node->generic = generic;
    SET_BASE(node, NODE_GENERIC_EXPR);

    RETURN_NODE(node);
}

CREATE_NODE(slice_expr)(struct expression_node* expression, struct expression_node* low,
    struct expression_node* high, struct expression_node* max, bool slice_three){
    NEW_NODE(node, slice_expr);

    node->expression = expression;
    node->low = low;
    node->high = high;
    node->max = max;
    node->slice_three = slice_three;
    SET_BASE(node, NODE_SLICE_EXPR);

    RETURN_NODE(node);
}

CREATE_NODE(binary_expr)(struct expression_node* left, struct token op, struct expression_node* right){
    NEW_NODE(node, binary_expr);

    node->left = left;
    node->op = op;
    node->right = right;
    SET_BASE(node, NODE_BINARY_EXPR);

    RETURN_NODE(node);
}

CREATE_NODE(unary_expr)(struct expression_node* operand, struct token op){
    NEW_NODE(node, unary_expr);

    node->left = left;
    node->op = op;
    node->right = right;
    SET_BASE(node, NODE_UNARY_EXPR);

    RETURN_NODE(node);
}
