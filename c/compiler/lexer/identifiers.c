
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

// numbers can have the following forms:
// 0xAA, 0b0101, 0o55, 55, 5.5,
IDENTIFIER(number){
    int c = LEXER_PEEK;
    return ((c >= 0) && (c <= 9));
}
