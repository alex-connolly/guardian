
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

CREATE_NODE(if_stat)(struct statement_node* init, struct expression_node* cond,
    struct block_stat_node* body, struct block_stat_node* else_branch){

    NEW_NODE(node, if_stat);

    node->init = init;
    node->cond = cond;
    node->body = body;
    node->else_branch = else_branch;
    SET_BASE(node, NODE_IF_STAT);

    RETURN_NODE(node);
}

CREATE_NODE(branch_stat)(const char* identifier){
    NEW_NODE(node, branch_stat);

    node->identifier = identifier;
    SET_BASE(node, NODE_BRANCH_STAT);

    RETURN_NODE(node);
}

CREATE_NODE(case_clause)(struct array* cases, struct block_stat_node* block){

    NEW_NODE(node, block_stat);

    node->cases = cases;
    node->block = block;

    SET_BASE(node, NODE_CASE_CLAUSE);

    RETURN_NODE(node);
}
