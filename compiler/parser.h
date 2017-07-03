
#include <stdint.h>
#include <stdlib.h>
#include "stack.h"
#include "lexer.h"

struct parser {
    struct stack* lexers;
    struct token** tokens;
    size_t num_tokens;
};

struct parser* parser_create(const char *source, size_t len, uint32_t fileid, bool is_static);
struct node* parser_run(struct parser* parser);
void parser_free(struct parser* parser);
