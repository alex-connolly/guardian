#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#include "commands.h"

#define MAX_FILES 10000
const char* suffix = ".grd";

int main(int argc, char** argv){
    if (argc == 1){
        execute_help(argc, argv);
        return -1;
    }
    for (int i = 0; i < NUM_COMMANDS; i++){
        if (strcmp(commands[i].name, argv[1]) == 0){
            if (commands[i].validate(argc, argv)){
                commands[i].execute(argc, argv);
            }
            return 0;
        }
    }
    // none of the commands were matched, try compiling
    if (validate_compile(argc, argv)){
        execute_compile(argc, argv);
    }
}

/*bool find_file(char* name){
    return false;
}

bool validate_file_name(char* name){
    int length = strlen(name);
    int s_len = strlen(suffix);
    if (length <= s_len){
        return false;
    }
    // compare last s_len characters of the file name to our suffix
    return strncmp(&name[length-s_len-1], suffix, s_len);
}

bool validate_arguments(int argc, char** argv){
    if (argc == 1){
        // no arguments provided to program
        printf("At least one %s file must be provided.\n", suffix);
        return false;
    } else if (argc > MAX_FILES + 1){
        // too many arguments provided to program
        printf("Guardian can only process %d files at once.\n", MAX_FILES);
        return false;
    } else {
        for (int i = 1; i < argc; i++){
            if (!validate_file_name(argv[i])){
                printf("%s is not a %s file.\n", argv[i], suffix);
                return false;
            } else if (!find_file(argv[i])){
                printf("Could not locate file %s.\n", argv[i]);
                return false;
            }
        }
    }
}*/
