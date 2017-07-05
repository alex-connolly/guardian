
#include "base.h"

CREATE_NODE(assignment)(struct expression_node** left, size_t token_position, struct expression_node** right);
CREATE_NODE(call)();
CREATE_NODE(return_stat)(struct expression_node** results);
CREATE_NODE(branch_stat)(const char* identifier);
CREATE_NODE(if_stat)(struct expression_node* init, struct expression_node* cond,
    struct block_stat_node* body, struct block_stat_node* else_branch);
CREATE_NODE(case_clause)();
CREATE_NODE(switch_stat)();
CREATE_NODE(block_stat)();
CREATE_NODE(for_stat)();

struct assignment_node {
    struct expresion_node** left;
    size_t token_position;
    struct token token;
    struct expression_node** right;
};

// keyword <call>
// defer <call>
// run <call>
struct call_node {
    size_t position;
    struct call_expr_node* call;
};

struct return_stat_node {
    size_t return_position;
    struct expression_node** results;
};

struct branch_stat_node {
    size_t position;
    struct token token;
    const char* identifier;
};

struct if_stat_node {
    size_t if_position;
    struct statement* init;
    struct expression_node* cond;
    struct block_stat_node* body;
    struct block_stat_node* else_branch;
};

struct case_clause_node {
    size_t case_position;
    struct expression_node* stat;
    size_t colon_position;
    struct statement_node** body;
};

struct switch_stat_node {
    size_t switch_position;
    struct body_stat_node* body;
};

struct block_stat_node {
    size_t open_brace_position;
    struct statement* list;
    size_t close_brace_position;
};

struct for_stat_node {
    size_t for_position;
    struct statement_node* init;
    struct expression_node* cond;
    struct statement_node* post;
    struct block_stat_node* block;
};
