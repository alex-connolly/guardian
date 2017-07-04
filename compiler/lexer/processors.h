
#include "lexer.h"

#define PROCESS(name) struct token process_##name(struct lexer* lexer)

#define NEW_TOKEN struct token token; token.start = lexer->offset; token.end = lexer->offset;
#define SET_TYPE(t) token.type = t;
#define INC_LENGTH(num) token.end += num

#define SET_TYPE_RETURN(type) SET_TYPE(type); return token;

#define IF_NEXT_TYPE(character, type) if (c2 == character){ \
    INC_OFFSET(1); INC_LENGTH(1); SET_TYPE(type); \
}

// return type1 if the next character is '=' else type2
#define IF_ASSIGNMENT_ELSE(type1, type2) if (c2 == '='){ \
    SET_TYPE_RETURN(type1); \
} else { \
    SET_TYPE_RETURN(type2); \
}

PROCESS(builtin_operator){
    NEW_TOKEN;

    INC_LENGTH(1);
    int c = LEXER_NEXT;
    int c2 = LEXER_CURRENT;

    switch (c){

        // mathematical operators
        case '+': IF_ASSIGNMENT_ELSE(TKN_ADD_ASSIGN, TKN_ADD)
        case '-': IF_ASSIGNMENT_ELSE(TKN_SUB_ASSIGN, TKN_SUB)
        case '*': IF_ASSIGNMENT_ELSE(TKN_MUL_ASSIGN, TKN_MUL)
        case '/': IF_ASSIGNMENT_ELSE(TKN_DIV_ASSIGN, TKN_DIV)
        case '%': IF_ASSIGNMENT_ELSE(TKN_MOD_ASSIGN, TKN_MOD)

        // comparison operators
        case '!': IF_ASSIGNMENT_ELSE(TKN_NEQ, TKN_NOT)
        case '=': IF_ASSIGNMENT_ELSE(TKN_EQL, TKN_ASSIGN)
        case '<':
            if (c2 == '<'){
                c2 = LEXER_NEXT;
                IF_ASSIGNMENT_ELSE(TKN_SHL_ASSIGN, TKN_SHL)
            }
            IF_ASSIGNMENT_ELSE(TKN_LEQ, TKN_LSS)
        case '>':
            if (c2 == '>'){
                c2 = LEXER_NEXT;
                IF_ASSIGNMENT_ELSE(TKN_SHR_ASSIGN, TKN_SHR)
            }
            IF_ASSIGNMENT_ELSE(TKN_GEQ, TKN_GTR)

        // bitwise operators
        case '&':
            IF_NEXT_TYPE('&', TKN_LOGICAL_AND)
            IF_ASSIGNMENT_ELSE(TKN_AND_ASSIGN TKN_AND)
        case '|':
            IF_NEXT_TYPE('|', TKN_OR)
            IF_ASSIGNMENT_ELSE(TKN_BIT_OR_ASSIGN, TKN_BIT_OR)
        case '^': IF_ASSIGNMENT_ELSE(TKN_XOR_ASSIGN, TKN_XOR)
        case '~': IF_ASSIGNMENT_ELSE(TKN_LNO TKN_BIT_NOT)

        // miscellaneous operators
        case ',': SET_TYPE_RETURN(TKN_COMMA);
        case ':': SET_TYPE_RETURN(TKN_COLON);
        case '?': SET_TYPE_RETURN(TKN_TERNARY);

        // braces and brackets
        case '{': SET_TYPE_RETURN(TKN_BRACE_OPEN);
        case '}': SET_TYPE_RETURN(TKN_BRACE_CLOSE);
        case '[': SET_TYPE_RETURN(TKN_SQUARE_OPEN);
        case ']': SET_TYPE_RETURN(TKN_SQUARE_CLOSE);
        case '(': SET_TYPE_RETURN(TKN_BRACKET_OPEN);
        case ')': SET_TYPE_RETURN(TKN_BRACKET_CLOSE);
    }
}
