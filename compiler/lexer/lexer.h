
#include <stdlib.h>
#include <stdbool.h>
#include "token.h"

#define LEXER_PEEK lexer->buffer[lexer->offset];
#define LEXER_NEXT lexer->buffer[lexer->offset++];

struct lexer {
    char* buffer;
    size_t offset;

    size_t column;
    size_t line;
};



struct proto_token {
    bool (*identifier)(struct lexer* lexer);
    void (*processor)(struct lexer* lexer);
};

#define NUM_PROTO_TOKENS 5

#define PROTOTOKEN(name) (struct proto_token){identifier_##name, process_##name}

struct proto_token proto_tokens[NUM_PROTO_TOKENS] = {
    PROTOTOKEN(builtin_operator)
};
