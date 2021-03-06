// {C} Copyright 2018 Pensando Systems Inc. All rights reserved.

// CGO wrappers for accessing delphi metrics

#include <stdbool.h>
#include <stdio.h>
#include <cstdint>
#include <map>
#include <iostream>
#include <string>

#include "nic/delphi/shm/delphi_metrics_cgo.h"
#include "nic/delphi/shm/delphi_metrics.hpp"
#ifdef __x86_64__
#elif __aarch64__
#include "nic/sdk/lib/pal/pal.hpp"
#endif

using namespace delphi::shm;
using std::string;

// IteratorCtx is a temporary struct used to hold the iterator context
class IteratorCtx {
public:
    delphi::shm::TableMgrUptr     tbl_;
    delphi::shm::TableIterator    iter_;
    char                          *currVal_ = NULL;
};

void DelphiMetricsInit_cgo() {
    // initialize pal
#ifdef __x86_64__
#elif __aarch64__
    assert(sdk::lib::pal_init(platform_type_t::PLATFORM_TYPE_HAPS) == sdk::lib::PAL_RET_OK);
#endif

}

DelphiMetricsIterator_cgo NewMetricsIterator_cgo(const char *kind) {
    // get the shared memory object
    delphi::shm::DelphiShmPtr shm = delphi::metrics::DelphiMetrics::GetDelphiShm();
    assert(shm != NULL);

    // create the table if it doesnt exist already
    MetricsCreateTable(kind);

    // create temporary context
    IteratorCtx *metr = new IteratorCtx();
    assert(metr->currVal_ == NULL);

    // find the table
    metr->tbl_ = std::move(shm->Kvstore()->Table(kind));
    assert(metr->tbl_ != NULL);

    // create the iterator
    metr->iter_ = metr->tbl_->Iterator();

    return (void *)metr;
}

bool MetricsIteratorIsNotNil_cgo(DelphiMetricsIterator_cgo metr) {
    return ((IteratorCtx *)metr)->iter_.IsNotNil();
}

bool MetricsIteratorIsNil_cgo(DelphiMetricsIterator_cgo metr) {
    return ((IteratorCtx *)metr)->iter_.IsNil();
}

DelphiMetrics_cgo MetricsIteratorNext_cgo(DelphiMetricsIterator_cgo metr) {
     // release old value
     if (((IteratorCtx *)metr)->currVal_ != NULL) {
         ((IteratorCtx *)metr)->iter_.ReleaseValue(((IteratorCtx *)metr)->currVal_);
     }

    // get current value
    ((IteratorCtx *)metr)->currVal_ = ((IteratorCtx *)metr)->iter_.AcquireValue();
    
    // move the iterator
    ((IteratorCtx *)metr)->iter_.Next();

    // return current value
    return (void *)((IteratorCtx *)metr)->currVal_;
}

DelphiMetrics_cgo MetricsIteratorFind_cgo(DelphiMetricsIterator_cgo metr, char *key, int keylen) {
    // return the current value
    auto *mptr = ((IteratorCtx *)metr)->iter_.Find(key, keylen);
    return (void *)mptr;
}

void MetricsIteratorDelete_cgo(DelphiMetricsIterator_cgo metr, char *key, int keylen) {
    ((IteratorCtx *)metr)->iter_.Delete(key, keylen);
}

DelphiMetrics_cgo MetricsIteratorGet_cgo(DelphiMetricsIterator_cgo metr) {
    // return the current value
    auto *mptr = ((IteratorCtx *)metr)->iter_.Value();
    return (void *)mptr;
}

// MetricsIteratorFree_cgo frees the memory associated with iterator
void MetricsIteratorFree_cgo(DelphiMetricsIterator_cgo metr) {
    IteratorCtx *mitr = ((IteratorCtx *)metr);

     // release old value
     if (mitr->currVal_ != NULL) {
         mitr->iter_.ReleaseValue(mitr->currVal_);
     }

    mitr->tbl_ = nullptr;
    delete mitr;
}

const char *MetricsEntryKey(DelphiMetrics_cgo mtr) {
    return KEY_PTR_FROM_HASH_ENTRY(HASH_ENTRY_FROM_VAL_PTR(mtr, delphi::metrics::DelphiMetrics::GetDelphiShm()));
}

const char *MetricsEntryValue(DelphiMetrics_cgo mtr) {
    return (char *)mtr;
}

int MetricsEntryKeylen(DelphiMetrics_cgo mtr) {
    return (int)(HASH_ENTRY_FROM_VAL_PTR(mtr, delphi::metrics::DelphiMetrics::GetDelphiShm())->key_len);
}

int MetricsEntryVallen(DelphiMetrics_cgo mtr) {
    return (int)(HASH_ENTRY_FROM_VAL_PTR(mtr, delphi::metrics::DelphiMetrics::GetDelphiShm())->val_len);
}

void MetricsCreateTable(const char *kind) {
    // get the shared memory object
    delphi::shm::DelphiShmPtr shm = delphi::metrics::DelphiMetrics::GetDelphiShm();
    assert(shm != NULL);

    // create the table in shared memory
    delphi::shm::TableMgrUptr tbl = shm->Kvstore()->Table(kind);
    if (tbl == NULL) {
        tbl = shm->Kvstore()->CreateTable(kind, DEFAULT_METRIC_TBL_SIZE);
        assert(tbl != NULL);
    }

    return;
}

DelphiMetrics_cgo MetricsCreateEntry(const char *kind, char *key, int keylen, int vallen) {
    // get the shared memory object
    delphi::shm::DelphiShmPtr shm = delphi::metrics::DelphiMetrics::GetDelphiShm();
    assert(shm != NULL);

    // get the table
    delphi::shm::TableMgrUptr tbl = shm->Kvstore()->Table(kind);
    assert(tbl != NULL);

    // create an entry
    char *valptr = (char *)tbl->Create(key, keylen, vallen);
    assert(valptr != NULL);
    assert(((int64_t)valptr & 0x07) == 0);

    return (void *)valptr;
}

long long GetCounter(DelphiMetrics_cgo mtr, int offset) {
    auto entry = HASH_ENTRY_FROM_VAL_PTR(mtr, delphi::metrics::DelphiMetrics::GetDelphiShm());
    // check if this is a dpstats coming from PAL memory
    if (entry->flags & HT_ENTRY_FLAG_DPSTATS) {
        long long data = 0;
#ifdef __x86_64__
#elif __aarch64__
        uint64_t pal_addr = *(uint64_t *)mtr + offset;
        assert((pal_addr & 0x07) == 0);
        auto rc = sdk::lib::pal_reg_read(pal_addr, (uint32_t *)&data, 2);
        if (rc != sdk::lib::PAL_RET_OK) {
            LogError("Error reading from PAL. Err: %d\n", rc);
            return 0;
        }
#endif

        return data;
    }

    // return from shared memory
    long long *ptr = (long long *)((intptr_t)mtr + offset);
    assert(((int64_t)ptr & 0x07) == 0);
    return *ptr;
}

void SetCounter(DelphiMetrics_cgo mtr, long long val, int offset) {
    long long *ptr = (long long *)((intptr_t)mtr + offset);
    *ptr = val;
}

double GetGauge(DelphiMetrics_cgo mtr, int offset) {
    auto entry = HASH_ENTRY_FROM_VAL_PTR(mtr, delphi::metrics::DelphiMetrics::GetDelphiShm());
    // check if this is a dpstats coming from PAL memory
    if (entry->flags & HT_ENTRY_FLAG_DPSTATS) {
        double data = 0;
#ifdef __x86_64__
#elif __aarch64__
        uint64_t pal_addr = *(uint64_t *)mtr + offset;
        assert((pal_addr & 0x07) == 0);
        auto rc = sdk::lib::pal_reg_read(pal_addr, (uint32_t *)&data, 2);
        if (rc != sdk::lib::PAL_RET_OK) {
            LogError("Error reading from PAL. Err: %d\n", rc);
            return 0;
        }
#endif

        return data;
    }

    // return value from shared memory
    void *ptr = (void *)((intptr_t)mtr + offset);
    assert(((int64_t)ptr & 0x07) == 0);
    return *(double *)ptr;
}

void SetGauge(DelphiMetrics_cgo mtr, double val, int offset) {
    void *ptr = (void *)((intptr_t)mtr + offset);
     *(double *)ptr = val;
}
