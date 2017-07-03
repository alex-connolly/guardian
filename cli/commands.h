#include <stdbool.h>
#include <stdlib.h>
#include <stdio.h>

#define ARGUMENTS int argc, char** argv

struct command {
    const char* name;
    bool (*validate)(ARGUMENTS);
    void (*execute)(ARGUMENTS);
};

#define NUM_COMMANDS 3

#define DEFINE_COMMAND(n) bool validate_##n(ARGUMENTS); void execute_##n(ARGUMENTS)

DEFINE_COMMAND(compile);

DEFINE_COMMAND(help);
DEFINE_COMMAND(version);

#define COMMAND(search, n) (struct command){.name = search, .validate = validate_##n, .execute = execute_##n}

// The NULL command name indicates that it is the default commmand
static struct command commands[NUM_COMMANDS] = {
    COMMAND("help", help),
    COMMAND("version", version)
};
