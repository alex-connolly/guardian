#include "commands.h"

const size_t v1 = 0;
const size_t v2 = 0;
const size_t v3 = 1;

#define VALIDATE(name) bool validate_##name(ARGUMENTS)
#define EXECUTE(name) void execute_##name(ARGUMENTS)

VALIDATE(help){
    if (argc > 2){
        printf("usage: guardian help\n");
        return false;
    }
    return true;
}

EXECUTE(help){
    printf("version: display current version\n");
    printf("help: show this message\n");
    printf("<file_name>: compile/execute a file\n");
}

VALIDATE(version){
    if (argc > 2){
        printf("usage: guardian version\n");
        return false;
    }
    return true;
}

EXECUTE(version){
    printf("%zu.%zu.%zu", v1, v2, v3);
}

VALIDATE(compile){
    return false;
}

EXECUTE(compile){
    return;
}
