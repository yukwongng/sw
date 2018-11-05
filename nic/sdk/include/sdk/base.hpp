//------------------------------------------------------------------------------
// {C} Copyright 2017 Pensando Systems Inc. All rights reserved
//
// base SDK header file that goes into rest of the SDK
//------------------------------------------------------------------------------

#ifndef __SDK_BASE_HPP__
#define __SDK_BASE_HPP__

#include <ostream>
#include <stdint.h>
#include <assert.h>
#include <stdlib.h>
#include "include/sdk/logger.hpp"
#include "include/sdk/platform.hpp"

namespace sdk {

#define TRUE                                         1
#define FALSE                                        0

#define _In_
#define _Out_
#define _Inout_

#define likely(x)           __builtin_expect(!!(x), 1)
#define unlikely(x)         __builtin_expect(!!(x), 0)

#ifndef __PACK__
#define __PACK__            __attribute__((packed))
#endif

#define __ALIGN__           __attribute__((__aligned__(n)))
#define __CACHE_ALIGN__     __ALIGN__(CACHE_LINE_SIZE)
#define __ASSERT__(x)       assert(x)

#define SDK_ASSERT_RETURN(cond, rv)                        \
do {                                                       \
    if (!(cond)) {                                         \
        SDK_TRACE_ERR("ASSERT FAILURE(" #cond ")");        \
        __ASSERT__(FALSE);                                 \
        return rv;                                         \
    }                                                      \
} while (FALSE)

#define SDK_ASSERT_RETURN_VOID(cond)                       \
do {                                                       \
    if (!(cond)) {                                         \
        SDK_TRACE_ERR("ASSERT FAILURE(" #cond ")");        \
        __ASSERT__(FALSE);                                 \
        return;                                            \
    }                                                      \
} while (FALSE)

#define SDK_ASSERT_GOTO(cond, label)                       \
do {                                                       \
    if (!(cond)) {                                         \
        SDK_TRACE_ERR("ASSERT FAILURE(" #cond ")");        \
        __ASSERT__(FALSE);                                 \
        goto label;                                        \
    }                                                      \
} while (FALSE)

#define SDK_ASSERT(cond)        __ASSERT__(cond)

#define SDK_ABORT(cond)                                    \
do {                                                       \
    if (!(cond)) {                                         \
        abort();                                           \
    }                                                      \
} while (FALSE)

//-----------------------------------------------------------------------------
// X-Macro for defining enums (generates enum definition and string formatter)
//
// Example:
//
// #define SAMPLE_ENUM_ENTRIES(ENTRY)
//    ENTRY(OK, 0, "ok")
//    ENTRY(ERR, 1, "err")
//
// SDK_DEFINE_ENUM(sample_enum_t, SAMPLE_ENUM_ENTRIES)
//------------------------------------------------------------------------------
#define _SDK_ENUM_FIELD(_name, _val, _desc) _name = _val,
#define _SDK_ENUM_STR(_name, _val, _desc)                               \
            inline const char *_name ## _str() {return _desc; }
#define _SDK_ENUM_CASE(_name, _val, _desc) case _name: return os << #_desc;
#define _SDK_ENUM_CASE_STR(_name, _val, _desc) case _name: return #_name;

#define SDK_DEFINE_ENUM(_type, _entries)                                \
    typedef enum { _entries(_SDK_ENUM_FIELD) } _type;                   \
    inline std::ostream& operator<<(std::ostream& os, _type t) {        \
        switch (t) {                                                    \
            _entries(_SDK_ENUM_CASE);                                   \
        }                                                               \
        return os;                                                      \
    }                                                                   \
    _entries(_SDK_ENUM_STR)

#define SDK_DEFINE_ENUM_TO_STR(_typ, _entries)                          \
    inline const char *_entries##_str(_typ c)                           \
    {                                                                   \
        switch (c) {                                                    \
            _entries(_SDK_ENUM_CASE_STR);                               \
        }                                                               \
        return "uknown";                                                \
    }

//------------------------------------------------------------------------------
// atomic operations
//------------------------------------------------------------------------------
#define SDK_ATOMIC_INC_UINT32(ptr, val)     __atomic_add_fetch(ptr, val, __ATOMIC_SEQ_CST)
#define SDK_ATOMIC_DEC_UINT32(ptr, val)     __atomic_sub_fetch(ptr, val, __ATOMIC_SEQ_CST)
#define SDK_ATOMIC_STORE_UINT32(ptr, vptr)  __atomic_store(ptr, vptr, __ATOMIC_SEQ_CST)
#define SDK_ATOMIC_LOAD_UINT32(ptr, vptr)   __atomic_load(ptr, vptr, __ATOMIC_SEQ_CST)
#define SDK_ATOMIC_INC_UINT64(ptr, val)     __atomic_add_fetch(ptr, val, __ATOMIC_SEQ_CST)
#define SDK_ATOMIC_DEC_UINT64(ptr, val)     __atomic_sub_fetch(ptr, val, __ATOMIC_SEQ_CST)

#define SDK_ARRAY_SIZE(arr)                (sizeof((arr))/sizeof((arr)[0]))

//------------------------------------------------------------------------------
// SDK return codes
//------------------------------------------------------------------------------
#define SDK_RET_ENTRIES(ENTRY)                                                              \
    ENTRY(SDK_RET_OK,                         0,    "ok")                                   \
    ENTRY(SDK_RET_OOM,                        1,    "out of memory error")                  \
    ENTRY(SDK_RET_INVALID_ARG,                2,    "invalid arg")                          \
    ENTRY(SDK_RET_INVALID_OP,                 3,    "invalid operation")                    \
    ENTRY(SDK_RET_ENTRY_NOT_FOUND,            4,    "lookup failure")                       \
    ENTRY(SDK_RET_ENTRY_EXISTS,               5,    "already present")                      \
    ENTRY(SDK_RET_NO_RESOURCE,                6,    "resource exhaustion")                  \
    ENTRY(SDK_RET_TABLE_FULL,                 7,    "hw table full")                        \
    ENTRY(SDK_RET_DUPLICATE_INS,              8,    "duplicate insert fail")                \
    ENTRY(SDK_RET_OOB,                        9,    "out-of-bound error")                   \
    ENTRY(SDK_RET_HW_PROGRAM_ERR,             10,   "h/w programming error")                \
    ENTRY(SDK_RET_RETRY,                      11,   "retry the operation")                  \
    ENTRY(SDK_RET_NOOP,                       12,   "No operation performed")               \
    ENTRY(SDK_RET_DUPLICATE_FREE,             13,   "duplicate free")                       \
    ENTRY(SDK_RET_HBM_HASH_COLL,              14,   "hbm hash table collision")             \
    ENTRY(SDK_RET_HBM_HASH_MAX_RECIRC_EXCEED, 15,   "hbm hash maximum recirc exceeded")     \
    ENTRY(SDK_RET_MAX,                        16,   "max return value")                     \
    ENTRY(SDK_RET_ERR,                        255,  "catch all generic error")

SDK_DEFINE_ENUM(sdk_ret_t, SDK_RET_ENTRIES)
SDK_DEFINE_ENUM_TO_STR(sdk_ret_t, SDK_RET_ENTRIES)
#undef SDK_RET_ENTRIES

}    // namespace sdk

using sdk::sdk_ret_t;

#endif    // __SDK_BASE_HPP__

