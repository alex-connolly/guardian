#define NEW_MPZ(name) mpz_t name; mpz_init()
#define POP()
#define PUSH()

OP(add){
    POP(x);
    POP(y);
    NEW_MPZ(z);
    mpz_add(z, x, y);
    PUSH(z);
}

OP(sub){
    POP(x);
    POP(y);
    NEW_MPZ(z);
    mpz_sub(z, x, y);
    PUSH(z);
}

OP(mul){
    POP(x);
    POP(y);
    NEW_MPZ(z);
    mpz_mul(z, x, y);
    PUSH(z);
}

OP(div){
    POP(x);
    POP(y);
    NEW_MPZ(z);
    mpz_div(z, x, y);
    PUSH(z);
}
