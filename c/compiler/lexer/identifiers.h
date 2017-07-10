
#include "lexer.h"

#define IDENTIFIER(name) bool identifier_##name(struct lexer* lexer)

IDENTIFIER(builtin_operator);
IDENTIFIER(whitespace);
IDENTIFIER(newline);
IDENTIFIER(identifier);
IDENTIFIER(number);
IDENTIFIER(string);
