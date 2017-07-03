#pragma once

#include <stdbool.h>
#include <stdlib.h>
#include "utils.h"

struct lexer {
    size_t size;
    unsigned int offset;
    unsigned int column;
    unsigned int line;
    char* buffer;
};

struct token {
    int start;
    int end;
};

struct proto_token {
    bool (*is)(struct lexer* lexer);
    void (*process)(struct lexer* lexer);
};

#define DECLARE_TOKEN(name) \
bool is_##name(struct lexer* lexer); \
void process_##name(struct lexer* lexer);

DECLARE_TOKEN(string)
DECLARE_TOKEN(number)
DECLARE_TOKEN(whitespace)
DECLARE_TOKEN(builtin_operator)

#define NUM_TOKENS 10

#define TOKEN(name) (struct proto_token){is_##name, process_##name}

static struct proto_token proto_tokens [NUM_TOKENS] = {
    TOKEN(string),
    TOKEN(number),
    TOKEN(whitespace),
    TOKEN(builtin_operator),
};


struct lexer* lexer_create(char* buffer, size_t size);
void lexer_free(struct lexer* lexer);
struct token* lexer_run(struct lexer* lexer, int* num_tokens);
