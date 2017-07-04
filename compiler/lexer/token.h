
#include <stdlib.h>
#include "types.h"

struct token {
    enum token_type type;
    size_t start;
    size_t end;
};
