//
// {C} Copyright 2020 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// This file contains athena_app test utility API
///
//----------------------------------------------------------------------------

#ifndef __ATHENA_APP_TEST_UTILS_HPP__
#define __ATHENA_APP_TEST_UTILS_HPP__

#include <unistd.h>
#include <getopt.h>
#include <limits.h>
#include <stdlib.h>
#include <stdint.h>
#include <string>
#include <iostream>
#include <map>
#include <vector>
#include <stdarg.h>
#include <sys/stat.h>
#include <sys/time.h>
#include <netinet/in.h>
#include "nic/sdk/third-party/zmq/include/zmq.h"
#include "nic/sdk/include/sdk/base.hpp"
#include "nic/apollo/test/api/utils/base.hpp"
#include "nic/apollo/api/include/athena/pds_base.h"
#include "script_parser.hpp"
#include "nic/include/ftl_dev_if.hpp"
#include "nic/apollo/api/include/athena/pds_conntrack.h"
#include "nic/apollo/api/include/athena/pds_flow_session_info.h"

#ifndef ARRAY_SIZE
#define ARRAY_SIZE(x)   (sizeof(x) / sizeof((x)[0]))
#endif

#define TEST_LOG_INFO(args...)                      \
    do {                                            \
        printf(args);                               \
        fflush(stdout);                             \
    } while (false)

#define TEST_LOG_ERR(args...)                       \
    do {                                            \
         printf("ERROR: " args);                    \
         fflush(stdout);                             \
     } while (false)


namespace test {
namespace athena_app {

/**
 * Generic test function and parameters
 */
class test_param_t
{
public:
    test_param_t(uint32_t num) : 
        type(TOKEN_TYPE_NUM),
        num_(num)
    {
    }
    test_param_t(const std::string& str) : 
        type(TOKEN_TYPE_STR),
        str_(str)
    {
    }

    test_param_t(const std::vector<test_param_t>& t) : 
        type(TOKEN_TYPE_TUPLE_BEGIN)
    {
        tuple_ = t;
    }

    bool is_num(void) const { return type == TOKEN_TYPE_NUM;  }
    bool is_str(void) const { return type == TOKEN_TYPE_STR;  }
    bool is_tuple(void) const { return type == TOKEN_TYPE_TUPLE_BEGIN;  }

    pds_ret_t  num(uint32_t *ret_num,
                   bool suppress_err_log = false) const;
    pds_ret_t  str(std::string *ret_str,
                   bool suppress_err_log = false) const;
    pds_ret_t  bool_val(bool *ret_bool,
                        bool suppress_err_log = false) const;
    pds_ret_t  flowtype(pds_flow_type_t *ret_flowtype,
                        bool suppress_err_log = false) const;
    pds_ret_t  flowstate(pds_flow_state_t *ret_flowstate,
                         bool suppress_err_log = false) const;
    pds_ret_t  proto(uint32_t *ret_proto,
                     bool suppress_err_log = false) const;
    pds_ret_t  tuple(std::vector<test_param_t> *ret_tuple,
                     bool suppress_err_log = false) const;

    friend class test_vparam_t;

private:
    token_type_t                type;
    uint32_t                    num_;
    std::string                 str_;
    std::vector<test_param_t>   tuple_;
};

typedef std::vector<test_param_t> test_param_tuple_t;

/**
 * Vector parameters
 */
class test_vparam_t
{
public:
    test_vparam_t() {}
    ~test_vparam_t() { clear(); }

    void push_back(test_param_t param) { vparam.push_back(param); }

    void clear(void) { vparam.clear(); }

    uint32_t size(void) const { return vparam.size(); }

    pds_ret_t num(uint32_t idx,
                  uint32_t *ret_num,
                  bool suppress_err_log = false) const;
    pds_ret_t str(uint32_t idx,
                  std::string *ret_str,
                  bool suppress_err_log = false) const;
    pds_ret_t flowtype(uint32_t idx,
                       pds_flow_type_t *ret_flowtype,
                       bool suppress_err_log = false) const;
    pds_ret_t flowstate(uint32_t idx,
                        pds_flow_state_t *ret_flowstate,
                        bool suppress_err_log = false) const;
    pds_ret_t proto(uint32_t idx,
                    uint32_t *ret_proto,
                    bool suppress_err_log = false) const;
    pds_ret_t tuple(uint32_t idx,
                    test_param_tuple_t *ret_tuple,
                    bool suppress_err_log = false) const;
    bool expected_bool(bool dflt = false) const;
    uint32_t expected_num(uint32_t dflt = 0) const;
    const std::string& expected_str(void) const;

private:
    std::vector<test_param_t>   vparam;
};

typedef const test_vparam_t&    test_vparam_ref_t;
typedef bool (*test_fn_t)(test_vparam_ref_t vparam);

#define APP_TEST_NAME2STR(name)                     \
    #name
    
#define APP_TEST_NAME2FN_MAP_ENTRY(name)            \
    {APP_TEST_NAME2STR(name), name}

/*
 * Tuple evaluation helper
 */
class tuple_eval_t
{
public:
    tuple_eval_t() :
        fail_count(0) {}

    void reset(test_vparam_ref_t vparam,
               uint32_t vparam_idx);
    uint32_t num(uint32_t idx);
    std::string str(uint32_t idx);
    pds_flow_type_t flowtype(uint32_t idx);
    pds_flow_state_t flowstate(uint32_t idx);

    uint32_t size(void) { return tuple.size(); }
    bool zero_failures(void)
    { 
        if (fail_count) {
            TEST_LOG_ERR("Tuple evaluation failed\n");
            return false;
        }
        return true;
    }

private:
    test_param_tuple_t          tuple;
    uint32_t                    fail_count;
};

/*
 * Map of created session/conntrack IDs, for the purpose of
 * cross checking when IDs are aged out.
 */
class id_map_t
{
public:

    ~id_map_t() { clear(); }

    bool insert(uint32_t id)
    {
        std::pair<std::map<uint32_t,uint32_t>::iterator,bool> ret;
        ret = id_map.insert(std::make_pair(id, 0));
        return ret.second;
    }

    void erase(uint32_t id) { id_map.erase(id); }

    bool find(uint32_t id)
    {
        auto iter = id_map.find(id);
        return iter != id_map.end();
    }

    bool find_erase(uint32_t id)
    {
        if (find(id)) {
            erase(id);
            return true;
        }
        return false;
    }

    void clear(void) { id_map.clear(); }

    uint32_t size(void) { return (uint32_t)id_map.size(); }

private:
    std::map<uint32_t,uint32_t> id_map;
};

/*
 * Aging timeout config failure counts for various operations
 */
class tmo_cfg_fail_count_t
{
public:

    tmo_cfg_fail_count_t() { clear(); }

    void clear(void) { counters = {0}; }

    uint32_t total(void) { return counters.get + counters.set; }

    friend class aging_tmo_cfg_t;

private:

    struct {
        uint32_t                get;
        uint32_t                set;;
    } counters;
};

/*
 * Aging timeout config
 */
class aging_tmo_cfg_t
{
public:
    aging_tmo_cfg_t(bool is_accel_tmo) :
        is_accel_tmo(is_accel_tmo),
        max_tmo(SCANNER_SESSION_TMO_DFLT)
    {
    }

    void reset(void);
    void session_tmo_set(uint32_t tmo_val);
    uint32_t session_tmo_get(void) const { return tmo_rec.session_tmo; }
    uint32_t max_tmo_get(void) const { return max_tmo; }

    void conntrack_tmo_set(pds_flow_type_t flowtype,
                           pds_flow_state_t flowstate,
                           uint32_t tmo_val);
    uint32_t conntrack_tmo_get(pds_flow_type_t flowtype,
                               pds_flow_state_t flowstate);

    uint32_t fail_count(void) { return failures.total(); }

private:
    void tmo_set(void);

    pds_flow_age_timeouts_t     tmo_rec;
    tmo_cfg_fail_count_t        failures;
    bool                        is_accel_tmo;
    uint32_t                    max_tmo;    // largest timeout overall
};

/*
 * Tolerance failure counts for various operations
 */
class tolerance_fail_count_t
{
public:

    tolerance_fail_count_t() { clear(); }

    void clear(void) { counters = {0}; }

    uint32_t total(void)
    {
        return counters.accel_control   +
               counters.info_read       +
               counters.under_age       +
               counters.over_age        +
               counters.create_add      +
               counters.create_erase    +
               counters.empty_check     +
               counters.delete_errors;
    }

    friend class aging_tolerance_t;

private:

    struct {
        uint32_t                accel_control;
        uint32_t                info_read;
        uint32_t                under_age;
        uint32_t                over_age;
        uint32_t                create_add;
        uint32_t                create_erase;
        uint32_t                empty_check;
        uint32_t                delete_errors;
    } counters;
};

/*
 * Aging result, with tolerance
 */

/*
 * Two modes of usage of create_id_map inside aging_tolerance_t:
 * 1) Storage of IDs to keep track of all ID creation and aging removal, or
 * 2) Count only, suitable for large scale testing where it may not
 *    be memory efficient to store hundreds of thousands of IDs.
 */
#define AGING_TOLERANCE_STORE_ID_THRESH     16384

class aging_tolerance_t
{
public:
    aging_tolerance_t(uint32_t num_ids_max = AGING_TOLERANCE_STORE_ID_THRESH) :
        normal_tmo(false),
        accel_tmo(true),
        curr_tmo(&normal_tmo),
        num_ids_max(num_ids_max),
        tolerance_secs(0),
        over_age_min_(UINT32_MAX),
        over_age_max_(0),
        using_fte_indices_(false),
        create_count_(0),
        erase_count_(0),
        expiry_count_(0) {}

    ~aging_tolerance_t() { create_id_map.clear(); }

    void reset(uint32_t ids_max = AGING_TOLERANCE_STORE_ID_THRESH);
    void tolerance_secs_set(uint32_t tolerance_secs);
    void age_accel_control(bool enable_sense);
    void session_tmo_tolerance_check(uint32_t id);
    void conntrack_tmo_tolerance_check(uint32_t id);

    void create_id_map_insert(uint32_t id);
    void create_id_map_find_erase(uint32_t id);
    uint32_t create_id_map_size(void);
    void create_id_map_empty_check(void);

    uint32_t curr_max_tmo(void) const { return curr_tmo->max_tmo_get(); }
    void expiry_count_inc(void) { expiry_count_++; }
    void delete_errors_inc(void) { failures.counters.delete_errors++; }
    uint32_t delete_errors(void) { return failures.counters.delete_errors; }
    uint32_t create_count(void) { return create_count_; }
    uint32_t expiry_count(void) { return expiry_count_; }
    uint32_t over_age_min(void)
    {
        return over_age_min_ == UINT32_MAX ? 0 : over_age_min_;
    }
    uint32_t over_age_max(void) { return over_age_max_; }
    void using_fte_indices(bool yesno) { using_fte_indices_ = yesno; }
    bool using_fte_indices(void) { return using_fte_indices_; }

    bool create_map_with_ids(void)
    {
        return num_ids_max <= AGING_TOLERANCE_STORE_ID_THRESH;
    }

    uint32_t fail_count(void)
    { 
        return failures.total()         + 
               normal_tmo.fail_count()  +
               accel_tmo.fail_count();
    }

    bool zero_failures(void) { return fail_count() == 0; }

    aging_tmo_cfg_t             normal_tmo;
    aging_tmo_cfg_t             accel_tmo;

private:

    void tmo_tolerance_check(uint32_t id,
                             uint32_t delta_secs,
                             uint32_t applic_tmo_secs);

    aging_tmo_cfg_t             *curr_tmo;
    uint32_t                    num_ids_max;
    uint32_t                    tolerance_secs;
    uint32_t                    over_age_min_;
    uint32_t                    over_age_max_;
    tolerance_fail_count_t      failures;
    bool                        using_fte_indices_;

    id_map_t                    create_id_map;
    uint32_t                    create_count_;
    uint32_t                    erase_count_;
    uint32_t                    expiry_count_;
};

/**
 * Flow Key field helper
 */
class flow_key_field_t
{
public:
    flow_key_field_t()
    {
        reset(0, 1);
    }

    void reset(uint32_t value,
               uint32_t count)
    {
        value_ = value;
        count_ = count ? count : 1;
        save_value_ = value_;
        save_count_ = count;
    }

    uint32_t count(void) { return count_; }
    uint32_t value(void) { return value_; }
    void next_value(void)
    {
        SDK_ASSERT(count_);
        ++value_;
        --count_;
    }

    void restart(void)
    {
        value_ = save_value_;
        count_ = save_count_;
    }

private:
    uint32_t                    value_;
    uint32_t                    count_;
    uint32_t                    save_value_;
    uint32_t                    save_count_;
};

/**
 * Metrics
 */
class aging_metrics_t
{
public:

    aging_metrics_t(ftl_dev_if::ftl_qtype qtype) :
        qtype(qtype)
    {
        base = {0};
    }

    pds_ret_t baseline(void);
    uint64_t delta_expired_entries(void) const;
    uint64_t delta_num_qfulls(void) const;

    /*
     * For informational purposes, check if HW metrics agreed with SW count
     */
    bool expiry_count_check(uint32_t sw_expiry_count)
    {
        uint64_t delta = delta_expired_entries();
        if (delta != (uint64_t)sw_expiry_count) {
            TEST_LOG_INFO("HW delta_expired_entries %" PRIu64 
                          " != SW expiry_count %u\n", delta, sw_expiry_count);
        }
        return delta == sw_expiry_count;
    }

    void show(bool latest = true) const;

private:
    pds_ret_t refresh(ftl_dev_if::lif_attr_metrics_t& m) const;

    enum ftl_dev_if::ftl_qtype     qtype;
    ftl_dev_if::lif_attr_metrics_t base;
};

#ifndef USEC_PER_SEC
#define USEC_PER_SEC    1000000L
#endif

#ifndef MSEC_PER_SEC
#define MSEC_PER_SEC    1000L
#endif

#define APP_TIME_LIMIT_EXEC_SECS(s)         ((uint64_t)(s) * USEC_PER_SEC)
#ifdef __x86_64__
#define APP_TIME_LIMIT_EXEC_GRACE           300 /* seconds over max tmo */
#else
#define APP_TIME_LIMIT_EXEC_GRACE           60 /* seconds over max_tmo */
#endif

#define APP_TIME_LIMIT_USLEEP_DFLT          10000 /* 10ms */

static inline uint64_t
timestamp(void)
{
    struct timeval tv;

    gettimeofday(&tv, NULL);
    return (((uint64_t)tv.tv_sec * USEC_PER_SEC) +
            (uint64_t)tv.tv_usec);
}

void mpu_tmr_wheel_update(void);

/**
 * Generic timestamp and expiry interval
 */
typedef bool (*time_limit_exec_fn_t)(void *);

class test_timestamp_t
{
public:
    test_timestamp_t() :
        ts(0),
        expiry(0) {}

    void time_expiry_set(uint64_t val)
    {
        ts = timestamp();
        expiry = val;
    }

    bool time_expiry_check(void)
    {
        return (expiry == 0) ||
               ((timestamp() - ts) > expiry);
    }

    void time_limit_exec(time_limit_exec_fn_t fn,
                         void *user_ctx = nullptr,
                         uint32_t sleep_us = APP_TIME_LIMIT_USLEEP_DFLT)
    {
        while (!time_expiry_check()) {
            if ((*fn)(user_ctx)) {
                break;
            }

            mpu_tmr_wheel_update();

            /*
             * Always need a sleep to yield on SIM platform;
             * whereas it'll be up to the caller on HW platform.
             */
            if (hw()) {
                if (sleep_us) {
                    usleep(sleep_us);
                }
            } else {
                usleep(sleep_us > APP_TIME_LIMIT_USLEEP_DFLT ?
                       sleep_us : APP_TIME_LIMIT_USLEEP_DFLT);
            }
        }
    }

private:
    uint64_t                    ts;
    uint64_t                    expiry;
};

static inline void
randomize_seed(void)
{
    srand(timestamp());
}

static inline uint32_t
randomize_max(uint32_t val_max,
              bool zero_ok = false)
{
    uint32_t rand_num = rand() % (val_max + 1);
    return (rand_num ? rand_num : (zero_ok ? 0 : 1));
}

pds_ret_t script_exec(const std::string& scripts_dir,
                      const std::string& script_fname);
/*
 * Server message process functions
 */
typedef pds_ret_t (*server_msg_proc_fn_t)(zmq_msg_t *rx_msg,
                                          zmq_msg_t *tx_msg);
pds_ret_t script_exec_msg_process(zmq_msg_t *rx_msg,
                                  zmq_msg_t *tx_msg);

/*
 * Miscellaneous
 */
static inline void
flow_session_key_init(pds_flow_session_key_t *key)
{
    memset(key, 0, sizeof(*key));
    key->direction = HOST_TO_SWITCH | SWITCH_TO_HOST;
}

static inline void
flow_conntrack_key_init(pds_conntrack_key_t *key)
{
    memset(key, 0, sizeof(*key));
}

uint64_t mpu_timestamp(void);
uint32_t session_timestamp(uint64_t mpu_timestamp,
                           bool underage_adjust = false);
uint32_t session_timestamp_diff(uint32_t session_ts_end,
                                uint32_t session_ts_start);
uint32_t conntrack_timestamp(uint64_t mpu_timestamp,
                             bool underage_adjust = false);
uint32_t conntrack_timestamp_diff(uint32_t conntrack_ts_end,
                                  uint32_t conntrack_ts_start);

/*
 * On Capri, P4 updates flow timestamp using bits 47:23 of the MPU tick.
 * With a clock speed of 833MHz, this gives interval of 1.01E-02 (10.1ms).
 */
static inline uint32_t
session_timestamp2secs(uint32_t session_ts)
{
    return session_ts / 100;
}

static inline uint32_t
conntrack_timestamp2secs(uint32_t conntrack_ts)
{
    return session_timestamp2secs(conntrack_ts);
}

}    // namespace athena_app
}    // namespace test

/*
 * Special well-known app_test exit function
 */
#define APP_TEST_EXIT_FN        app_test_exit
#define APP_TEST_EXIT_FN_STR    APP_TEST_NAME2STR(app_test_exit)
bool app_test_exit(test::athena_app::test_vparam_ref_t vparam);
bool skip_fte_flow_prog_set(test::athena_app::test_vparam_ref_t vparam);
bool flow_cache_dump(test::athena_app::test_vparam_ref_t vparam);
bool session_info_dump(test::athena_app::test_vparam_ref_t vparam);

bool skip_fte_flow_prog(void);
bool skip_dpdk_init(void);
void program_prepare_exit(void);

#endif   // __ATHENA_APP_TEST_UTILS_HPP__
