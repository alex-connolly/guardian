
#include "../util/array.h"
#include <stdbool.h>
#include "../lexer/token.h"
#include "../ast/ast.h"

struct parser {
    struct token current;
};

struct parser* parser_create();
void parser_run(struct parser* parser);
void parser_free(struct parser* parser);

#define PARSE(name) parse_##name(struct parser* parser)

#define CONSTRUCT(name)

struct construct {
    bool (*identify)(struct parser* parser);
    void (*parse)(struct parser* parser);
};
