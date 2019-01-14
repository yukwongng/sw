//-----------------------------------------------------------------------------
// {C} Copyright 2017 Pensando Systems Inc. All rights reserved
//-----------------------------------------------------------------------------
#ifndef __MEM_HASH_TYPES_HPP__
#define __MEM_HASH_TYPES_HPP__

#include "include/sdk/base.hpp"
#include "lib/p4/p4_api.hpp"

#include "mem_hash.hpp"

using sdk::table::mem_hash_factory_params_t;
using sdk::table::mem_hash_properties_t;

namespace sdk {
namespace table {
namespace memhash {

#define HINT_SLOT_IS_INVALID(_slot) \
        ((_slot) == mem_hash_api_context::hint_slot::HINT_SLOT_INVALID)
#define HINT_SLOT_IS_VALID(_slot) \
        ((_slot) != mem_hash_api_context::hint_slot::HINT_SLOT_INVALID)
#define HINT_SLOT_SET_INVALID(_slot) \
        ((_slot) = mem_hash_api_context::hint_slot::HINT_SLOT_INVALID)
#define HINT_SLOT_IS_MORE(_slot) \
        ((_slot) == mem_hash_api_context::hint_slot::HINT_SLOT_MORE)
#define HINT_SLOT_SET_MORE(_slot) \
        ((_slot) = mem_hash_api_context::hint_slot::HINT_SLOT_MORE)
#define HINT_SLOT_IS_MATCH(_ctx) \
        (HINT_SLOT_IS_FREE((_ctx)->hint_slot) &&\
         HINT_SLOT_IS_MORE((_ctx)->hint_slot))
#define HINT_IS_VALID(_hint) \
        ((_hint) != mem_hash_api_context::hint_index::HINT_INDEX_INVALID)
#define HINT_SET_INVALID(_hint) \
        ((_hint) = mem_hash_api_context::hint_index::HINT_INDEX_INVALID)

#define PRINT_API_CTX_HW_DATA(_ctx) {\
    if ((_ctx)->key2str) { \
        SDK_TRACE_DEBUG("- Key:[%s]", (_ctx)->key2str((_ctx)->hwkey));\
    } \
    if ((_ctx)->appdata2str) { \
        SDK_TRACE_DEBUG("- AppData:[%s]", (_ctx)->appdata2str((_ctx)->hwdata)); \
    } \
    SDK_TRACE_DEBUG("- HwData:[%s]", (_ctx)->hwdata2str()); \
}

#define PRINT_API_CTX(_name, _ctx) {\
    SDK_TRACE_DEBUG("%s: %s, [%s]", _name, (_ctx)->idstr(), (_ctx)->metastr()); \
}

class mem_hash_api_context {
public:
    enum match_type{
        MATCH_TYPE_NONE     = 0,
        MATCH_TYPE_EXM      = 1,
        MATCH_TYPE_HINT     = 2,
    };

    enum hint_index {
        // Hint index 0 is reserved
        HINT_INDEX_INVALID = 0,
    };

    enum hint_slot {
        // Slot 0 is invalid. NCC always generates from Slot 1-to-N.
        HINT_SLOT_INVALID  = 0,

        // Slot 1
        // ....
        // Slot N are all valid slots, auto-generated by NCC.

        // Slot 0xFF is reserved for MoreHints.
        HINT_SLOT_MORE  = 0xFF,
    };


private:
    static uint32_t numctx_;
    static mem_hash_api_context* alloc_(uint32_t hwkey_len, uint32_t hwdata_len) {
        mem_hash_api_context *ctx = NULL;
        void *mem = NULL;

        mem = (mem_hash_api_context *) SDK_CALLOC(SDK_MEM_ALLOC_MEM_HASH_API_CTX,
                                                  sizeof(mem_hash_api_context));
        if (!mem) {
            SDK_TRACE_ERR("Failed to alloc api context.");
            return NULL;
        }

        ctx = new (mem) mem_hash_api_context();
        ctx->hwkey_len = hwkey_len;
        if (ctx->hwkey_len) {
            ctx->hwkey = (uint8_t *) SDK_CALLOC(SDK_MEM_ALLOC_MEM_HASH_HWKEY, hwkey_len);
        }

        if (!ctx->hwkey) {
            SDK_TRACE_ERR("Failed to alloc hwkey:%p", ctx->hwkey);
            destroy(ctx);
            return NULL;
        }

        ctx->hwdata_len = hwdata_len;
        if (ctx->hwdata_len) {
            ctx->hwdata = (uint8_t *) SDK_CALLOC(SDK_MEM_ALLOC_MEM_HASH_HWDATA, hwdata_len);
            if (!ctx->hwdata) {
                SDK_TRACE_ERR("Failed to alloc hwdata:%p.", ctx->hwdata);
                destroy(ctx);
                return NULL;
            }
        }

        numctx_++;
        SDK_TRACE_DEBUG("Number of API Contexts = %d", numctx_);
        return ctx;
    }

public:
    // Input params
    void *key;
    void *appdata;

    uint32_t hwkey_len;
    uint32_t hwdata_len;
    uint32_t appdata_len;
    bool hash_valid;
    uint32_t hash_32b;
    void *cookie; // Callback cookie for iteratin
    key2str_t key2str;
    appdata2str_t appdata2str;
    //mem_hash_iterate_func_t iterfunc;

    // Derived fields from input
    uint32_t hash_msbits;

    // Placeholders for HW read/write operations
    bool hw_valid;
    uint8_t *hwkey;
    uint8_t *hwdata;

    // NOTE NOTE NOTE:
    // Some of the below fields are re-used by main table and hint table
    // DO NOT USE these to pass info between tables.
    uint8_t level;          // Chaining level
    uint8_t num_hints;
    uint8_t max_recircs;
    uint8_t table_id;
    uint32_t table_index;
    uint8_t hint_slot;
    uint32_t hint;
    bool more_hashes;
    bool write_pending;
    uint32_t match_type;
    void *bucket;
    char str[256];
    char str2[16];

    // Parent API Context: used for context nesting.
    // 1st level HintTable: pctx = MainTable context.
    // 2nd level HintTable: pctx = 1st level HintTable context.
    // and so on...
    mem_hash_api_context *pctx;

    // Constructor
    mem_hash_api_context() {
    }

    // Destructor
    ~mem_hash_api_context() {
    }

    bool is_exact_match() {
        SDK_ASSERT(match_type != MATCH_TYPE_NONE);
        return match_type == MATCH_TYPE_EXM;
    }

    bool is_hint_match() {
        SDK_ASSERT(match_type != MATCH_TYPE_NONE);
        return match_type == MATCH_TYPE_HINT;
    }

    void set_exact_match() {
        match_type = MATCH_TYPE_EXM;
        return;
    }

    void set_hint_match() {
        match_type = MATCH_TYPE_HINT;
        return;
    }

    bool is_hint_valid() {
        return hint != HINT_INDEX_INVALID;
    }

    char* inputstr() {
        sprintf(str, "key:%p,data:%p,hash_valid:%d,hash_32b:%#x,cookie:%p",
                key, appdata, hash_valid, hash_32b, cookie);
        return str;
    }

    // Debug string
    char* metastr() {
        sprintf(str, "id:%d,idx:%d,slot:%d,hint:%d,"
                "more:%d,pending:%d,hash_msbits:%d,match_type:%d",
                table_id, table_index, hint_slot,
                hint, more_hashes, write_pending, hash_msbits, match_type);
        return str;
    }

    char* hwdata2str()
    {
        uint32_t len = 0;
        uint32_t i = 0;

        len += sprintf(str, "Valid=%d,More=%d,MoreHints=%d,",
                       p4pd_mem_hash_entry_get_entry_valid(table_id, hwdata),
                       p4pd_mem_hash_entry_get_more_hashes(table_id, hwdata),
                       p4pd_mem_hash_entry_get_more_hints(table_id, hwdata));
        for (i = 1; i <= num_hints; i++) {
            uint32_t hintX = p4pd_mem_hash_entry_get_hint(table_id, hwdata, i);
            uint32_t hash_msbitsX = p4pd_mem_hash_entry_get_hash(table_id, hwdata, i);
            if (hintX) {
                len += sprintf(str + len, "HashMsbits%d=%#x,Hint%d=%d,",
                               i, hash_msbitsX, i, hintX);
            } else {
                // if hint is not valid, then hash must be zero.
                SDK_ASSERT(hash_msbitsX == 0);
            }
        }
        return str;
    }


    bool ismain() {
        return (level == 0);
    }

    const char* idstr() {
        sprintf(str2, "%s%d-L%d", ismain() ? "M" : "H", table_index, level);
        return str2;
    }

    static mem_hash_api_context* factory(mem_hash_api_context *pctx) {
        mem_hash_api_context *ctx = alloc_(pctx->hwkey_len, pctx->hwdata_len);
        if (!ctx) {
            return NULL;
        }
        // Copy the api params
        ctx->key = pctx->key;
        ctx->appdata = pctx->appdata;
        ctx->hash_valid = pctx->hash_valid;
        ctx->hash_32b = pctx->hash_32b;
        ctx->hash_msbits = pctx->hash_msbits;
        // Copy the required properties of the table to the context
        ctx->appdata_len = pctx->appdata_len;
        ctx->num_hints = pctx->num_hints;
        ctx->max_recircs = pctx->max_recircs;
        ctx->level = pctx->level + 1;
        ctx->key2str = pctx->key2str;
        ctx->appdata2str = pctx->appdata2str;

        ctx->pctx = pctx;
        return ctx;
    }

    static mem_hash_api_context* factory(mem_hash_api_params_t *params,
                                         mem_hash_properties_t *props) {
        mem_hash_api_context *ctx = alloc_(props->key_len, props->data_len);
        if (!ctx) {
            return NULL;
        }
        // Copy the api params
        ctx->key = params->key;
        ctx->appdata = params->appdata;
        ctx->hash_valid = params->hash_valid;
        ctx->hash_32b = params->hash_32b;
        ctx->hash_msbits = 0;
        // Copy the required properties of the table to the context
        ctx->appdata_len = props->appdata_len;
        ctx->num_hints = props->num_hints;
        ctx->max_recircs = props->max_recircs;
        ctx->level = 0;
        ctx->key2str = props->key2str;
        ctx->appdata2str = props->appdata2str;
        return ctx;
    }

    static void destroy(mem_hash_api_context* ctx) {
        SDK_TRACE_DEBUG("Destroying api context: %s", ctx->idstr());
        if (ctx->hwkey) {
            SDK_FREE(SDK_MEM_ALLOC_MEM_HASH_HWKEY, ctx->hwkey);
        }
        if (ctx->hwdata) {
            SDK_FREE(SDK_MEM_ALLOC_MEM_HASH_HWDATA, ctx->hwdata);
        }
        SDK_FREE(SDK_MEM_ALLOC_MEM_HASH_API_CTX, ctx);

        numctx_--;
        SDK_TRACE_DEBUG("Number of API Contexts = %d", numctx_);
        return;
    }

    static uint32_t count() {
        return numctx_;
    }

    bool is_max_recircs() {
        return (level >= max_recircs);
    }
};

} // namespace memhash
} // namespace table
} // namespace sdk

#endif // __MEM_HASH_TYPES_HPP__
