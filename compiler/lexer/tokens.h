
enum token_type {

    // Special tokens
	TKN_EOF,
	TKN_COMMENT,

	// Identifiers and basic type literals
	// (these tokens stand for classes of literals)
	TKN_IDENTIFIER,  // main
	TKN_INT, // 12345
	TKN_FLOAT,  // 123.45
	TKN_IMAG, // 123.45i
	TKN_CHAR, // 'a'
	TKN_STRING, // "abc"

	// Operators and delimiters
	TKN_ADD, // +
	TKN_SUB, // -
	TKN_MUL, // *
	TKN_DIV, // /
	TKN_MOD, // %

	TKN_AND, // &
	TKN_OR, // |
	TKN_XOR, // ^
	TKN_SHL, // <<
	TKN_SHR, // >>
	TKN_AND_NOT, // &^

	TKN_ADD_ASSIGN, // +=
	TKN_SUB_ASSIGN, // -=
	TKN_MUL_ASSIGN, // *=
	TKN_DIV_ASSIGN, // /=
	TKN_MOD_ASSIGN, // %=

	TKN_AND_ASSIGN, // &=
	TKN_OR_ASSIGN, // |=
	TKN_XOR_ASSIGN, // ^=
	TKN_SHL_ASSIGN, // <<=
	TKN_SHR_ASSIGN, // >>=
	TKN_AND_NOT_ASSIGN, // &^=

	TKN_LAND,  // &&
	TKN_LOR, // ||
	TKN_ARROW, // <-
	TKN_INC, // ++
	TKN_DEC, // --

	TKN_EQL, // ==
	TKN_LSS, // <
	TKN_GTR, // >
	TKN_ASSIGN, // =
	TKN_NOT, // !

	TKN_NEQ, // !=
	TKN_LEQ, // <=
	TKN_GEQ, // >=
	TKN_DEFINE, // :=
	TKN_ELLIPSIS, // ...

	TKN_BRACKET_OPEN, // (
    TKN_BRACKET_CLOSE, // )
	TKN_SQUARE_OPEN, // [
    TKN_SQUARE_CLOSE, // ]
	TKN_BRACE_OPEN, // {
    TKN_BRACE_CLOSE, // }

	TKN_COMMA,  // ,
	TKN_DOT, // .
	TKN_SEMICOLON, // ;
	TKN_COLON, // :
    TKN_TERNARY, // ?

	// Keywords
	TKN_BREAK,
    TKN_CONTINUE,
	TKN_CASE,
	TKN_CHAN,
	TKN_CONST,


	TKN_DEFAULT,
	TKN_DEFER,
	TKN_ELSE,
	TKN_FALLTHROUGH,
	TKN_FOR,

	TKN_FUNC,
	TKN_GOTO,
	TKN_IF,
	TKN_IMPORT,

	TKN_INTERFACE,
    TKN_CLASS,

    TKN_IS,
    TKN_AS,

	TKN_MAP,
	TKN_PACKAGE,
	TKN_RANGE,
	TKN_RETURN,
    TKN_RUN,

	TKN_SWITCH,
	TKN_TYPE,
	TKN_VAR
};

// precedence order taken from Swift
// number are immaterial --> only order is of concern
// no overloaded operators allowed for now (probably ever)
int get_precedence(enum token_type type){
    switch (type){
        // assignments
        case TKN_ADD_ASSIGN:
        case TKN_SUB_ASSIGN:
        case TKN_MUL_ASSIGN:
        case TKN_DIV_ASSIGN:
        case TKN_AND_ASSIGN:
        case TKN_AND_NOT_ASSIGN:
        case TKN_OR_ASSIGN:
        case TKN_XOR_ASSIGN:
        case TKN_SHR_ASSIGN:
        case TKN_SHL_ASSIGN:
            return 1;
        // ternary
        case TKN_TERNARY: return 2;
        // logical or
        case TKN_OR: return 3;
        // logical and
        case TKN_AND: return 4;
        // comparsion
        case TKN_EQL:
        case TKN_NEQ:
        case TKN_LEQ:
        case TKN_GEQ:
        case TKN_LSS:
        case TKN_GTR:
            return 5;
        // typing
        case TKN_IS:
        case TKN_AS:
            return 6;
        // range
        case TKN_RANGE:
            return 7;
        // factors
        case TKN_MUL:
        case TKN_DIV:
        case TKN_MOD:
            return 8;
        // shifts
        case TKN_SHL:
        case TKN_SHR:
            return 9;
        case TKN_DOT:
        case TKN_BRACKET_OPEN:
        case TKN_SQUARE_OPEN:
            return 10;
        default:
            return 0;
    }
}
