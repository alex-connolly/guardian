
#include "statements.h"

#define NEW_NODE(var_name, name) struct name_node* var_name = malloc(sizeof(struct name_node))
#define RETURN_NODE(var_name) (struct node*) var_name;
#define SET_BASE(var_name, t) node->base = (struct node*){.type = t}

CREATE_NODE(for_stat)(struct statement_node* init, struct expression_node* cond,
    struct expression_node* post, struct block_stat_node block){

    NEW_NODE(node, for_stat);

    node->init = init;
    node->cond = cond;
    node->post = post;
    node->block = block;
    SET_BASE(node, NODE_FOR_STAT);

    RETURN_NODE(node);
}

CREATE_NODE(return_stat)(struct expression_node** results){

    NEW_NODE(node, return_stat);

    node->results = results;
    SET_BASE(node, NODE_RETURN_STAT);

    RETURN_NODE(node);
}
