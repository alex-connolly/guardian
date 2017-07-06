
#include "parser.h"

#define PARSER_NEXT parser_next(parser)
#define REQUIRE(type) parser_require(parser, type)
#define REQUIRE_NEXT(type) REQUIRE(type); PARSER_NEXT

void parser_run(struct parser* parser){
    parser_next();
}

void parser_next(struct parser* parser){
    if (parser->offset = offset->size){
        return;
    }
    for (int i = 0; i < NUM_CONSTRUCTS; i++){
        if (constructs[i].identify(parser)){
            constructs[i].parse(parser);
            break;
        }
    }
    parser_next(parser);
}

void parser_require(struct parser* parser, enum token_type type){
    if (parser->current.type != type){

    }
}

PARSE(identifier){
    REQUIRE(TKN_IDENTIFIER);
    return CREATE_NODE(identifier)(extract_value(CURRENT_TOKEN));
}

PARSE(identifier_list){
    struct array* list = array_create();
    while (PARSER_CURRENT == TKN_COMMA){
        PARSER_NEXT;
        array_add(list, parse_identifier(parser));
    }
    return create_identifier_list_node(list);
}

PARSE(expression_list){
    struct array* list = array_create();
    while (PARSER_CURRENT.type == TKN_COMMA){
        PARSER_NEXT;
        array_add(list, check_expression())
    }
}

// <abstract>? class <name> {
PARSE(class){

    // parse the declaration
    bool abstract = false;
    if (CURRENT_TOKEN.type == TKN_ABSTRACT){
        abstract = true;
    }
    REQUIRE_NEXT(TKN_CLASS);
    REQUIRE(TKN_IDENTIFIER);
    const char* name = extract_value(CURRENT_TOKEN);
    PARSER_NEXT;

    const char* name = parse_super_class(parser);

    struct array* interfaces = parser_interfaces(parser);

    REQUIRE_NEXT(TKN_OPEN_BRACE);



    // start
    return CREATE_NODE(class)(name, super, interfaces);
}

PARSE(pointer){
    REQUIRE(TKN_ADDRESS);
    struct expr_node* expr = parse_type(parser);
    return CREATE_NODE(pointer)(expr);
}

PARSE(parameter_list){
    struct array* list = array_create(sizeof(struct expression_node));

}

PARSE(parameters){
    struct array* list = NULL;
    REQUIRE_NEXT(TKN_OPEN_PAREN);
    if (CURRENT_TOKEN != TKN_CLOSE_PAREN){
        list = parse_parameter_list(parser);
    }
    REQUIRE_NEXT(TKN_CLOSE_PAREN);
    return CREATE_NODE(field_list)(list);
}

PARSE(result){

}

PARSE(map){
    REQUIRE(TKN_MAP);
    REQUIRE(TKN_OPEN_SQUARE);
    struct type_node* key = parse_type(parser);
    REQUIRE(TKN_CLOSE_SQUARE);
    struct type_node* value = parse_type(parser);
    return CREATE_NODE(map)(key, value);
}

PARSE(channel){

}

// a block is a brace-enclosed series of statements
PARSE(block){
    REQUIRE(TKN_OPEN_BRACE);
    OPEN_SCOPE();
    struct array* list = parse_stat_list(parser);
    CLOSE_SCOPE();
    REQUIRE(TKN_CLOSE_BRACE);
    return block_stat_node_create(list);
}

PARSE(statement){
    struct array* left = parse_lhs_list(parser);

    switch (CURRENT_TOKEN){

    }

}

struct array* parse_superclasses(struct parser* parser){
    if (CURRENT_TOKEN.type == TKN_INHERITS){
        REQUIRE_NEXT(TKN_INHERITS);
        REQUIRE(TKN_IDENTIFIER);
        struct array* list = array_create(sizeof(char*));
        array_add(list, extract_value(CURRENT_TOKEN));
        while (CURRENT_TOKEN == TKN_COMMA){
            PARSER_NEXT;
            REQUIRE(TKN_IDENTIFIER);
            array_add(list, extract_value(CURRENT_TOKEN));
            PARSER_NEXT;
        }
    }
}

struct array* parse_interfaces(struct parser* parser){
    if (CURRENT_TOKEN.type == TKN_IS){
        REQUIRE_NEXT(TKN_IS);
        REQUIRE(TKN_IDENTIFIER);
         = array_create(sizeof(char*));
        array_add(list, extract_value(CURRENT_TOKEN));
        while (CURRENT_TOKEN == TKN_COMMA){
            PARSER_NEXT;
            REQUIRE(TKN_IDENTIFIER);
            array_add(list, extract_value(CURRENT_TOKEN));
            PARSER_NEXT;
        }
    }
}

const char* identifier = parse_identifier(struct parser* parser){
    REQUIRE(TKN_IDENTIFIER);
    const char* name = extract_value(CURRENT_TOKEN);
    PARSER_NEXT;
    return name;
}

PARSE(class_sig){
    // class signature syntax:
    // abstract? contract? class <name> inherits? <supers...> is <interfaces...> {
    bool is_abstract = parse_optional_token(TKN_ABSTRACT);

    bool is_contract = parse_optional_token(TKN_CONTRACT);

    REQUIRE_NEXT(TKN_CLASS);

    const char* identifier = parse_identifier(parser);

    struct array* supers = parse_superclasses(parser);

    struct array* interfaces = parse_interfaces(parser);

    class_decl_node_create(name, supers, interfaces, is_abstract, is_contract);
}

bool parse_optional_token(enum token_type type){
    bool is = (CURRENT_TOKEN.type == type);
    if (is) {
        PARSER_NEXT;
    }
    return is;
}

PARSE(interface_decl){
    // interfaces signature syntax:
    // abstract? interface <name> inherits? <supers...> {
    bool is_abstract = parse_optional_token(TKN_ABSTRACT);

    REQUIRE_NEXT(TKN_INTERFACE);

    const char* identifier = parse_identifier(parser);

    struct array* supers = parse_superclasses(parser);

    interface_decl_node_create(identifier, supers, is_abstract);
}

PARSE(interface){

    struct node* sig = parse_interface_sig(parser);

    REQUIRE_NEXT(TKN_OPEN_BRACE);
    struct node* body = parse_body(parser);
    REQUIRE_NEXT(TKN_CLOSE_BRACE);

    interface_node_create(sig, body);
}

PARSE(func){

    // functions are composed of:
    // sig { body }

    struct node* sig = parse_func_sig(parser);

    REQUIRE_NEXT(TKN_OPEN_BRACE);
    struct node* body = parse_body(parser);
    REQUIRE_NEXT(TKN_CLOSE_BRACE);

    func_node_create(sig, body);
}

PARSE(func_sig){

    bool is_abstract = parse_optional_token(TKN_ABSTRACT);

    const char* identifier = parse_identifier(parser);

    struct array* parameters = parse_parameters(parser);

    struct array* results = parse_results(parser);

    func_sig_node_create(identifier, parameters, results, is_abstract);
}

PARSE(switch_stat){

    // switch statement syntax
    // exclusive? switch (condition)?
    bool is_exclusive = parse_optional_token(TKN_EXCLUSIVE);

    REQUIRE_NEXT(TKN_SWITCH);

    struct node* switching_on = NULL;
    if (CURRENT_TOKEN.type != TKN_OPEN_BRACE){
        switching_on = parse_expression(parser);
    }

    switch_stat_node_create(switching_on, is_exclusive);
}

PARSE(for_stat){

    REQUIRE_NEXT(TKN_FOR);

    struct statement_node* init = parse_statement(parser);
    if (init){
        REQUIRE_NEXT(TKN_SEMICOLON);
    }

    // mandatory condition
    struct expression_node* cond = parse_expression(parser);

    struct statement_node* post = NULL;
    if (CURRENT_TOKEN.type != TKN_OPEN_BRACE){
        REQUIRE_NEXT(TKN_SEMICOLON);
        post = parse_statement(parser);
    }

    struct body_stat_node* block = parse_body_stat(parser);

    for_stat_node_create(init, cond, post, block);
}

PARSE(if_stat){

    REQUIRE_NEXT(TKN_IF);

    struct statement_node* init = parse_statement(parser);
    if (init){
        REQUIRE_NEXT(TKN_SEMICOLON);
    }

    struct expression_node* cond = parse_expression(parser);

    struct body_stat_node* block = parse_body_stat(parser);

}

PARSE(return_stat){

    REQUIRE_NEXT(TKN_RETURN);
    struct array* list = NULL;
    if (CURRENT_TOKEN.type != TKN_CLOSE_BRACE){
        list = array_create(sizeof(struct node*));
        if (CURRENT_TOKEN.type == TKN_OPEN_PAREN){
            PARSER_NEXT;
        }
        array_add(list, parse_expression(parser);
        for (CURRENT_TOKEN.type != TKN_COMMA){
            PARSER_NEXT;
            array_add(list, parse_expression(parser));
            PARSER_NEXT;
        }
    }
    return_stat_node_create(list);
}
