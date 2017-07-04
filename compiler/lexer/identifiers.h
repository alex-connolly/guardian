
#include "lexer.h"

#define IDENTIFIER(name) bool identifier_##name(struct lexer* lexer)

IDENTIFIER(builtin_operator);
