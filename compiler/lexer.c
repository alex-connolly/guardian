#include "lexer.h"

//
struct lexer* lexer_create(char* buffer, size_t size){
    struct lexer* lexer = malloc(sizeof(lexer));
    lexer->size = size;
    lexer->offset = 0;
    lexer->column = 0;
    lexer->line = 1;
    lexer->buffer = buffer;
    return lexer;
}

//
void lexer_free(struct lexer* lexer){
    free(lexer);
}

//
struct token* lexer_run(struct lexer* lexer, size_t* num_tokens){
    struct tokens* tokens;
    while (*lexer->buffer){
        for (int i = 0; i < NUM_TOKENS; i++){
            if (proto_tokens[i].is(lexer)){
                struct token* token = proto_tokens[i].process(lexer);
                array_add(tokens, num_tokens, token, sizeof(struct token));
            }
        }
    }
    return tokens;
}


#include "proto_tokens.h"

#define IS(name) bool is_##name(struct lexer* lexer)
#define PROCESS(name) void process_##name(struct lexer* lexer)

#define NEW_TOKEN struct token token; token.start = lexer->offset; token.end = lexer->offset;
#define SET_TYPE(t) token.type = t;
#define INC_LENGTH(num) token.end += num

IS(string){
    char c = *lexer->buffer;
    return (c == '"') || (c == '`');
}

PROCESS(string){

}

IS(number){
    char c = *lexer->buffer;
    return true;
}

PROCESS(hex_number){

}

PROCESS(binary_number){

    NEW_TOKEN();

    // process 0b prefix
    INC_LENGTH(2);

    bool dot_found = false;

    while ((c == CURRENT) != EOF){
        switch (c) {
            case '.':
                if (dot_found){
                    return lexer_error("Only one dot allowed in binary number.");
                }
                dot_found = true;
                INC_LENGTH(1);
                break;
            case '0', '1':
                INC_LENGTH(1);
                break;
            default:
                return lexer_error("Invalid character in binary number.")
        }
    }
    return token;
}

PROCESS(octal_number){

}

PROCESS(decimal_number){

}

PROCESS(number){
    if (CURRENT == '0'){
        switch (PEEK_CURRENT) {
            case 'x', 'X':
                return process_hex_number(lexer);
            case 'b', 'B':
                return process_binary_number(lexer);
            case 'o', 'O':
                return process_octal_number(lexer);
            default:
                return lexer_error();
        }
    }
    return process_decimal_number(lexer);
}


IS(builtin_operator){
    return ((c == '+') || (c == '-') || (c == '*') || (c == '/') ||
			(c == '<') || (c == '>') || (c == '!') || (c == '=') ||
			(c == '|') || (c == '&') || (c == '^') || (c == '%') ||
			(c == '~') || (c == '.') || (c == ';') || (c == ':') ||
			(c == '?') || (c == ',') || (c == '{') || (c == '}') ||
			(c == '[') || (c == ']') || (c == '(') || (c == ')') );
}

#define SET_TYPE_RETURN(type) SET_TYPE(type); return token;

#define IF_NEXT_TYPE(character, type) if (c2 == character){ \
    INC_OFFSET(1); INC_LENGTH(1); SET_TYPE(type); \
}

// return type1 if the next character is '=' else type2
#define IF_ASSIGNMENT_ELSE(type1, type2) if (c2 == '='){ \
    SET_TYPE_RETURN(type1);
} else { \
    SET_TYPE_RETURN(type2); \
}

PROCESS(builtin_operator){
    NEW_TOKEN;

    INC_LENGTH(1);
    int c = NEXT;
    int c2 = CURRENT;

    switch (c){

        // mathematical operators
        case '+': IF_ASSIGNMENT_ELSE(TKN_ADD_ASSIGN, TKN_ADD)
        case '-': IF_ASSIGNMENT_ELSE(TKN_SUB_ASSIGN, TKN_SUB)
        case '*': IF_ASSIGNMENT_ELSE(TKN_MUL_ASSIGN, TKN_MUL)
        case '/': IF_ASSIGNMENT_ELSE(TKN_DIV_ASSIGN, TKN_DIV)
        case '%': IF_ASSIGNMENT_ELSE(TKN_MOD_ASSIGN, TKN_MOD)

        // comparison operators
        case '!': IF_ASSIGNMENT_ELSE(TKN_INQUALITY, TKN_NOT)
        case '=': IF_ASSIGNMENT_ELSE(TKN_EQUALITY, TKN_ASSIGN)
        case '<':
            if (c2 == '<'){
                c2 = NEXT;
                IF_ASSIGNMENT_ELSE(TKN_LEFT_SHIFT_ASSIGN, TKN_LEFT_SHIFT)
            }
            IF_ASSIGNMENT_ELSE(TKN_)
        case '>':
            if (c2 == '>'){
                c2 = NEXT;
                IF_ASSIGNMENT_ELSE(TKN_RIGHT_SHIFT_ASSIGN, TKN_RIGHT_SHIFT)
            }
            IF_ASSIGNMENT_ELSE(TKN_GR_EQ, TKN_GR)

        // bitwise operators
        case '&':
            IF_NEXT_TYPE('&', TKN_AND)
            IF_ASSIGNMENT_ELSE(TKN_BIT_AND_ASSIGN, TKN_BIT_AND)
        case '|':
            IF_NEXT_TYPE('|', TKN_OR)
            IF_ASSIGNMENT_ELSE(TKN_BIT_OR_ASSIGN, TKN_BIT_OR)
        case '^': IF_ASSIGNMENT_ELSE(TKN_XOR_ASSIGN, TKN_XOR)
        case '~': IF_ASSIGNMENT_ELSE(TKN_BIT_NOT_ASSIGN, TKN_BIT_NOT)

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
