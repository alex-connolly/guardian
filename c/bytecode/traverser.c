
#include "traverser.h"

TRAVERSE(for_stat){

}



TRAVERSE(class_decl){

    if (node->superclass){

    }
}

TRAVERSE(variable_decl){

}

TRAVERSE(unary_expr){
    // possible unary expressions: ~, !, @, *, -

}
