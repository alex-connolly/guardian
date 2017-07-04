
#include "identifiers.h"

IDENTIFIER(builtin_operator){
    int c = LEXER_PEEK;
    return ((c == '+') || (c == '-') || (c == '*') || (c == '/') ||
			(c == '<') || (c == '>') || (c == '!') || (c == '=') ||
			(c == '|') || (c == '&') || (c == '^') || (c == '%') ||
			(c == '~') || (c == '.') || (c == ';') || (c == ':') ||
			(c == '?') || (c == ',') || (c == '{') || (c == '}') ||
			(c == '[') || (c == ']') || (c == '(') || (c == ')') );
}

IDENTIFIER(whitespace){
    int c = LEXER_PEEK;
    return ((c == ' ' || c == '\t'));
}

IDENTIFIER(newline){
    int c = LEXER_PEEK;
    return ((c == '\n'));
}
