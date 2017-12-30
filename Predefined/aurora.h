#ifndef AURORA_H
#define AURORA_H

#define FALSE 0
#define TRUE 1

#define TYPE(type) struct ar_##type
#define OPT_TYPE(type) struct ar_opt_##type

#define NAME(type) ar_##type
#define OPT_NAME(type) ar_opt_##type
#define GCPTR_NAME(type) ar_gcptr_##type

#define OPT_DEFINE(type) OPT_TYPE(type) { TYPE(type)* data; char exists; }
#define OPT_CREATE(name, type, value) TYPE(type) __opt_val_##name = value; OPT_TYPE(type) name = { .data = &__opt_val_##name, exists = TRUE }
#define OPT_CREATE_EMPTY(name, type) OPT_TYPE(type) name = { exists = FALSE }

#define OBJ_CREATE(name, type, value) TYPE(type) name = value
#define OBJ_DEFINE(type) TYPE(type); OPT_DEFINE(type); TYPE(type)
#define OBJ_FUNC(name, type) ar_##type_##name(TYPE(type)* this
#define OBJ_FUNC_CALL(inst, name, type) ar_##type_##name(inst

#define FUNC_RET(type, value) TYPE(type) __ret = value; return __ret

#define INT_IM(value) { .i = (value) }
#define INT_MATH(left, op, right) { .i = (left.i op right.i) } 

#define STR_IM(value) { .cstr = value, .length = sizeof(value) - 1 }

#endif