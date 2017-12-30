#include "aurora.h"

#include <stdio.h>

// AURORA PREDEFINES

OBJ_DEFINE(i64) {
    long long i;
};

OBJ_DEFINE(i32) {
    long i;
};

OBJ_DEFINE(i16) {
    short i;
};

OBJ_DEFINE(i8) {
    char i;
};

OBJ_DEFINE(string) {
    char* cstr;
    long length;
};

TYPE(i32) OBJ_FUNC(Length, string)) {
    FUNC_RET(i32, INT_IM(this->length));
}

/// USER DEFINED OBJECS

int main() {
    // CODE
}