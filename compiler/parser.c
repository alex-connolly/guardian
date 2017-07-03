#include "parser.h"

#define CURRENT_TOKEN parser->tokens[parser->index];
#define NEXT_TOKEN parser->index++;

struct node* parse_collection_literal(struct parser* parser){

    /*
    The inbuilt Guardian collections are maps and arrays.
    Arrays: {1, 2, 3}
    Maps: {"a": 10, "b": 20, "c": 30}
    */

    parse_required(parser, TKN_BRACE_OPEN);

    struct node* first_expression = parse_expression(parser);

    if (CURRENT_TOKEN.type == TKN_COLON){
        // the collection is a map


    } else {
        // the collection is an array
        struct token* current;
        while ((current = CURRENT_TOKEN).type != TKN_BRACE_CLOSE){

            struct node* expression = parse_expression(parser);

            NEXT_TOKEN;
        }
    }




}

struct parser* parser_create(const char* source, size_t len, uint32_t fileid, bool is_static){

    struct parser* p = malloc(sizeof(struct parser));
    if (!p){
        return NULL;
    }

    struct lexer* l = lexer_create(source, len);
    if (!l){
        free(p);
        return NULL;
    }

    p->lexers = stack_create(sizeof(struct lexer));
    stack_push(p->lexers, l);

    return p;
}

struct node* parser_run(struct parser* parser){

    while (parser->lexers->size > 0){
        struct lexer* l = stack_pop(parser->lexers);
        size_t num_tokens;
        parser->tokens = lexer_run(l, &num_tokens);
        parser->num_tokens = num_tokens;

        while (parser->index < parser->num_tokens){
            for (int i = 0; i < )
        }

        free_lexer(l);
    }
}

void parser_free(struct parser* parser){
    if (parser->lexers){
        while(parser->lexers->size > 0){
            struct lexer* lexer = stack_pop(parser->lexers);
            lexer_free(lexer);
        }
        stack_free(parser->lexers);
    }
    if (parser->declarations){
        free(parser->declarations);
    }
    free(parser);
}
