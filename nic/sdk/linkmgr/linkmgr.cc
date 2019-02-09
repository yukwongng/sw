// {C} Copyright 2017 Pensando Systems Inc. All rights reserved

#include "lib/thread/thread.hpp"
#include "lib/periodic/periodic.hpp"
#include "nic/sdk/platform/drivers/xcvr.hpp"
#include "linkmgr.hpp"
#include "linkmgr_state.hpp"
#include "linkmgr_internal.hpp"
#include "port.hpp"
#include "timer_cb.hpp"

namespace sdk {
namespace linkmgr {

// global log buffer
char log_buf[MAX_LOG_SIZE];

// global sdk-linkmgr state
linkmgr_state *g_linkmgr_state;

// global sdk-linkmgr config
linkmgr_cfg_t g_linkmgr_cfg;

// sdk-linkmgr threads
sdk::lib::thread *g_linkmgr_threads[LINKMGR_THREAD_ID_MAX];

// xcvr poll timer handle
void *xcvr_poll_timer_handle;

// link down poll list
void *port_link_poll_timer_handle;
port *link_poll_timer_list[MAX_UPLINK_PORTS];

// per producer request queues
linkmgr_queue_t g_linkmgr_workq[LINKMGR_THREAD_ID_MAX];

// Global setting for link status poll. Default is enabled.
bool port_link_poll_en = true;

bool
port_link_poll_enabled (void)
{
    return port_link_poll_en;
}

void
linkmgr_set_link_poll_enable (bool enable)
{
    port_link_poll_en = enable;
}

uint32_t
glbl_mode_mgmt (mac_mode_t mac_mode)
{
    return g_linkmgr_cfg.catalog->glbl_mode_mgmt(mac_mode);
}

uint32_t
ch_mode_mgmt (mac_mode_t mac_mode, uint32_t ch)
{
    return g_linkmgr_cfg.catalog->ch_mode_mgmt(mac_mode, ch);
}

uint32_t
glbl_mode (mac_mode_t mac_mode)
{
    return g_linkmgr_cfg.catalog->glbl_mode(mac_mode);
}

uint32_t
ch_mode (mac_mode_t mac_mode, uint32_t ch)
{
    return g_linkmgr_cfg.catalog->ch_mode(mac_mode, ch);
}

uint32_t
num_asic_ports(uint32_t asic)
{
    return g_linkmgr_cfg.catalog->num_asic_ports(asic);
}

uint32_t
sbus_addr_asic_port(uint32_t asic, uint32_t asic_port)
{
    return g_linkmgr_cfg.catalog->sbus_addr_asic_port(asic, asic_port);
}

uint32_t
jtag_id(void)
{
    return g_linkmgr_cfg.catalog->jtag_id();
}

uint8_t
sbm_clk_div(void)
{
    return g_linkmgr_cfg.catalog->sbm_clk_div();
}

uint8_t
num_sbus_rings(void)
{
    return g_linkmgr_cfg.catalog->num_sbus_rings();
}

bool
aacs_server_en(void)
{
    uint8_t en = g_linkmgr_cfg.catalog->aacs_server_en();
    if (en == 1) {
        return true;
    }

    return false;
}

bool
aacs_connect(void)
{
    uint8_t connect = g_linkmgr_cfg.catalog->aacs_connect();
    if (connect == 1) {
        return true;
    }

    return false;
}

uint32_t
aacs_server_port(void)
{
    return g_linkmgr_cfg.catalog->aacs_server_port();
}

std::string
aacs_server_ip(void)
{
    return g_linkmgr_cfg.catalog->aacs_server_ip();
}

serdes_info_t*
serdes_info_get(uint32_t sbus_addr,
                uint32_t port_speed,
                uint32_t cable_type)
{
    return g_linkmgr_cfg.catalog->serdes_info_get(sbus_addr,
                                               port_speed,
                                               cable_type);
}

int
serdes_build_id (void)
{
    return g_linkmgr_cfg.catalog->serdes_build_id();
}

int
serdes_rev_id (void)
{
    return g_linkmgr_cfg.catalog->serdes_rev_id();
}

std::string
serdes_fw_file (void)
{
    return g_linkmgr_cfg.catalog->serdes_fw_file();
}

sdk_ret_t
port_link_poll_timer_add(port *port)
{
    link_poll_timer_list[port->port_num() - 1] = port;
    return SDK_RET_OK;
}

sdk_ret_t
port_link_poll_timer_delete(port *port)
{
    link_poll_timer_list[port->port_num() - 1] = NULL;
    return SDK_RET_OK;
}

sdk::lib::thread *
current_thread (void)
{
    return sdk::lib::thread::current_thread() ?
                       sdk::lib::thread::current_thread() :
                       g_linkmgr_threads[LINKMGR_THREAD_ID_CFG];
}

bool
is_linkmgr_ctrl_thread()
{
    sdk::lib::thread *curr_thread = current_thread();
    sdk::lib::thread *ctrl_thread = g_linkmgr_threads[LINKMGR_THREAD_ID_CTRL];

    // TODO assert once the thread store is fixed
    if (curr_thread == NULL) {
        return false;
    }

    // if curr_thread/ctrl_thread is NULL, then init has failed or not invoked
    if (ctrl_thread == NULL) {
        assert(0);
    }

    // if ctrl_thread is not running, then linkmgr_event_wait hasn't been invoked
    // and no one is waiting to handle incoming msgs
    if (!ctrl_thread->is_running()) {
        return true;
    }

    if (curr_thread->thread_id() == ctrl_thread->thread_id()) {
        return true;
    }

    return false;
}

//------------------------------------------------------------------------------
// linkmgr thread notification by other threads
//------------------------------------------------------------------------------
sdk_ret_t
linkmgr_notify (uint8_t operation, linkmgr_entry_data_t *data,
                q_notify_mode_t mode)
{
    uint16_t            pindx = 0;
    sdk::lib::thread    *curr_thread = current_thread();
    uint32_t            curr_tid = curr_thread->thread_id();
    linkmgr_entry_t     *rw_entry;
    sdk_ret_t           ret = SDK_RET_OK;

    if (g_linkmgr_workq[curr_tid].nentries >= LINKMGR_CONTROL_Q_SIZE) {
        SDK_TRACE_ERR("Error: operation %d for thread %s, tid %d full",
                      operation, curr_thread->name(), curr_tid);
        return SDK_RET_ERR;
    }

    // SDK_TRACE_DEBUG("Thread: %s, Notify op %d",
    //                 current_thread()->name(), operation);

    while (!SDK_ATOMIC_COMPARE_EXCHANGE_WEAK(
               &g_linkmgr_workq[curr_tid].pindx,
               &pindx, (pindx+1) % LINKMGR_CONTROL_Q_SIZE));

    rw_entry = &g_linkmgr_workq[curr_tid].entries[pindx];
    rw_entry->opn = operation;
    rw_entry->status = SDK_RET_ERR;

    memcpy(&rw_entry->data, data, sizeof(linkmgr_entry_data_t));

    SDK_ATOMIC_STORE_BOOL(&rw_entry->done, false);

    SDK_ATOMIC_FETCH_ADD(&g_linkmgr_workq[curr_tid].nentries, 1);

    if (mode == q_notify_mode_t::Q_NOTIFY_MODE_BLOCKING) {
        while (SDK_ATOMIC_LOAD_BOOL(&rw_entry->done) == false) {
            if (curr_thread->can_yield()) {
                pthread_yield();
            }
        }
        ret = rw_entry->status;
    } else {
        ret = SDK_RET_OK;
    }

    return ret;
}

void
linkmgr_start (void)
{
    int thread_id = LINKMGR_THREAD_ID_CTRL;
    g_linkmgr_threads[thread_id]->start(g_linkmgr_threads[thread_id]);
}

static sdk_ret_t
port_event_enable (linkmgr_entry_data_t *data)
{
    port *port_p = (port *)data->ctxt;
    return port_p->port_enable();
}

static sdk_ret_t
port_event_disable (linkmgr_entry_data_t *data)
{
    port *port_p = (port *)data->ctxt;
    return port_p->port_disable();
}

static sdk_ret_t
port_bringup_timer (linkmgr_entry_data_t *data)
{
    uint32_t  retries = 0;
    sdk_ret_t ret     = SDK_RET_OK;
    port      *port_p = (port *)data->ctxt;

    // if the timer is called back with current timer handle, reset it
    if (port_p->link_bring_up_timer() == data->timer) {
        port_p->set_link_bring_up_timer(NULL);
    }

    // if the bringup timer has expired, reset and try again
    if (port_p->bringup_timer_expired() == true) {
        retries = port_p->num_retries();

        // port disable resets num_retries
        ret = port_p->port_disable();

        // Enable port only if max retries is not reached
        if (1 || retries < MAX_PORT_LINKUP_RETRIES) {
            port_p->set_num_retries(retries + 1);
            ret = port_p->port_enable();
        } else {
            // TODO notify upper layers?
            SDK_TRACE_DEBUG("Max retries: %d reached for port: %d",
                            MAX_PORT_LINKUP_RETRIES, port_p->port_num());
        }
    } else {
        ret = port_p->port_link_sm_process();
    }

    return ret;
}

static sdk_ret_t
port_debounce_timer (linkmgr_entry_data_t *data)
{
    port *port_p = (port *)data->ctxt;

    // if the timer is called back with current timer handle, reset it
    if (port_p->link_debounce_timer() == data->timer) {
        port_p->set_link_debounce_timer(NULL);
    }

    return port_p->port_debounce_timer();
}

sdk_ret_t
xcvr_poll_timer_wrapper (linkmgr_entry_data_t *data)
{
    sdk::platform::xcvr_poll_timer();

    // Reschedule the xcvr timer
    xcvr_poll_timer_handle =
        sdk::lib::timer_schedule(
            SDK_TIMER_ID_XCVR_POLL, XCVR_POLL_TIME, NULL,
            (sdk::lib::twheel_cb_t)xcvr_poll_timer_cb,
            false);

    return SDK_RET_OK;
}

sdk_ret_t
port_link_poll_timer (linkmgr_entry_data_t *data)
{
    if (port_link_poll_enabled() == true) {
        for (int i = 0; i < MAX_UPLINK_PORTS; ++i) {
            port *port_p = link_poll_timer_list[i];

            if (port_p != NULL) {
                if(port_p->port_link_status() == false) {
                    SDK_TRACE_DEBUG("%d: Link DOWN", port_p->port_num());
                    port_p->port_link_dn_handler();
                }
            }
        }
    }

    // reschedule the poll timer
    port_link_poll_timer_handle =
        sdk::lib::timer_schedule(
            SDK_TIMER_ID_LINK_POLL, LINKMGR_LINK_POLL_TIME, NULL,
            (sdk::lib::twheel_cb_t)port_link_poll_timer_cb,
            false);

    return SDK_RET_OK;
}

//------------------------------------------------------------------------------
// linkmgr's forever loop
//------------------------------------------------------------------------------
void*
linkmgr_event_loop (void* ctxt)
{
    uint32_t         qid         = 0;
    uint16_t         cindx       = 0;
    bool             work_done   = false;
    bool             rv          = true;
    linkmgr_entry_t  *rw_entry   = NULL;

    SDK_THREAD_INIT(ctxt);

    while (TRUE) {
        work_done = false;
        for (qid = 0; qid < LINKMGR_THREAD_ID_MAX; qid++) {
            if (!g_linkmgr_workq[qid].nentries) {
                // no read/write requests
                continue;
            }

            // found a read/write request to serve
            cindx = g_linkmgr_workq[qid].cindx;
            rw_entry = &g_linkmgr_workq[qid].entries[cindx];

            // SDK_TRACE_ERR("Thread: %s, op: %d",
            //               current_thread()->name(), rw_entry->opn);

            switch (rw_entry->opn) {
            case LINKMGR_OPERATION_PORT_ENABLE:
                port_event_enable(&rw_entry->data);
                break;

            case LINKMGR_OPERATION_PORT_DISABLE:
                port_event_disable(&rw_entry->data);
                break;

            case LINKMGR_OPERATION_PORT_BRINGUP_TIMER:
                port_bringup_timer(&rw_entry->data);
                break;

            case LINKMGR_OPERATION_PORT_DEBOUNCE_TIMER:
                port_debounce_timer(&rw_entry->data);
                break;

            case LINKMGR_OPERATION_PORT_LINK_POLL_TIMER:
                port_link_poll_timer(&rw_entry->data);
                break;

            case LINKMGR_OPERATION_XCVR_POLL_TIMER:
                xcvr_poll_timer_wrapper(&rw_entry->data);
                break;

            default:
                SDK_TRACE_ERR("Invalid operation %d", rw_entry->opn);
                rv = false;
                break;
            }

            // populate the results
            rw_entry->status =  rv ? SDK_RET_OK : SDK_RET_ERR;
            SDK_ATOMIC_STORE_BOOL(&rw_entry->done, true);

            // advance to next entry in the queue
            g_linkmgr_workq[qid].cindx++;
            if (g_linkmgr_workq[qid].cindx >= LINKMGR_CONTROL_Q_SIZE) {
                g_linkmgr_workq[qid].cindx = 0;
            }

            SDK_ATOMIC_FETCH_SUB(&g_linkmgr_workq[qid].nentries, 1);

            work_done = true;
        }

        // all queues scanned once, check if any work was found
        if (!work_done) {
            // didn't find any work, yield and give chance to other threads
            pthread_yield();
        }
    }
}

static sdk_ret_t
thread_init (linkmgr_cfg_t *cfg)
{
    int thread_prio  = 0;
    int thread_id    = 0;
    int sched_policy = SCHED_OTHER;

    thread_prio = sched_get_priority_max(sched_policy);
    if (thread_prio < 0) {
        return SDK_RET_ERR;
    }

    // wait until the periodic thread is ready
    while (!sdk::lib::periodic_thread_is_running()) {
        pthread_yield();
    }

    // start the poll timer
    port_link_poll_timer_handle =
        sdk::lib::timer_schedule(
            SDK_TIMER_ID_LINK_POLL, LINKMGR_LINK_POLL_TIME, NULL,
            (sdk::lib::twheel_cb_t)port_link_poll_timer_cb,
            false);

    // init the control thread
    thread_id = LINKMGR_THREAD_ID_CTRL;
    g_linkmgr_threads[thread_id] =
        sdk::lib::thread::factory(std::string("linkmgr-ctrl").c_str(),
                                  thread_id,
                                  sdk::lib::THREAD_ROLE_CONTROL,
                                  0x0 /* use all control cores */,
                                  linkmgr_event_loop,
                                  thread_prio,
                                  sched_policy,
                                  true);

    // TODO
    // Since Linkmgr's current_thread doesn't have access to HAL's CFG thread,
    // init the CFG thread as a work around for now.
    // This thread wont be started, but the thread structure is used only to
    // get the current_thread's ID during config push.
    if (cfg->process_mode == false) {
        // init the config thread
        thread_id = LINKMGR_THREAD_ID_CFG;
        g_linkmgr_threads[thread_id] =
            sdk::lib::thread::factory(std::string("linkmgr-cfg").c_str(),
                                      thread_id,
                                      sdk::lib::THREAD_ROLE_CONTROL,
                                      0x0 /* use all control cores */,
                                      linkmgr_event_loop /* TODO NULL? */,
                                      thread_prio,
                                      sched_policy,
                                      true);
        }

    return SDK_RET_OK;
}

static void
linkmgr_workq_init (void)
{
    uint32_t qid = 0;
    for (qid = 0; qid < LINKMGR_THREAD_ID_MAX; qid++) {
        g_linkmgr_workq[qid].nentries = 0;
    }
}

sdk_ret_t
xcvr_poll_init (linkmgr_cfg_t *cfg)
{
    sdk::platform::xcvr_init(cfg->xcvr_event_cb);

    xcvr_poll_timer_handle =
        sdk::lib::timer_schedule(
            SDK_TIMER_ID_XCVR_POLL, XCVR_POLL_TIME, NULL,
            (sdk::lib::twheel_cb_t)xcvr_poll_timer_cb,
            false);

    return SDK_RET_OK;
}

sdk_ret_t
linkmgr_init (linkmgr_cfg_t *cfg)
{
    sdk_ret_t    ret = SDK_RET_OK;

    g_linkmgr_cfg = *cfg;

    linkmgr_workq_init();

    if ((ret = thread_init(cfg)) != SDK_RET_OK) {
        SDK_TRACE_ERR("linkmgr thread init failed");
        return ret;
    }

    g_linkmgr_state = linkmgr_state::factory();
    if (NULL == g_linkmgr_state) {
        SDK_TRACE_ERR("linkmgr init failed");
        return SDK_RET_ERR;
    }

    // initialize the port mac and serdes functions
    port::port_init(cfg);

    // initialize xcvr polling
    xcvr_poll_init(cfg);

    return SDK_RET_OK;
}

//-----------------------------------------------------------------------------
// Validate speed against number of lanes
//-----------------------------------------------------------------------------
static bool
validate_speed_lanes (port_speed_t speed, uint32_t num_lanes)
{
    bool valid = true;

    // speed = NONE implies no update
    switch (num_lanes) {
    case 4:
        switch (speed) {
            case port_speed_t::PORT_SPEED_100G:
            case port_speed_t::PORT_SPEED_40G:
            case port_speed_t::PORT_SPEED_NONE:
                break;

            default:
                valid = false;
                break;
        }
        break;

    case 2:
        switch (speed) {
            case port_speed_t::PORT_SPEED_50G:
            case port_speed_t::PORT_SPEED_NONE:
                break;

            default:
                valid = false;
                break;
        }
        break;

    case 1:
        switch (speed) {
            case port_speed_t::PORT_SPEED_25G:
            case port_speed_t::PORT_SPEED_10G:
            case port_speed_t::PORT_SPEED_1G:
            case port_speed_t::PORT_SPEED_NONE:
                break;

            default:
                valid = false;
                break;
        }
        break;

    default:
        valid = false;
        break;
    }

    if (valid == false) {
        SDK_TRACE_ERR("Invalid speed and lanes config."
                      " num_lanes: %d, speed: %d",
                      num_lanes, speed);
    }

    return valid;
}

//-----------------------------------------------------------------------------
// set port args given transceiver state
//-----------------------------------------------------------------------------
sdk_ret_t
port_args_set_by_xcvr_state (port_args_t *port_args)
{
    // if xcvr is inserted:
    //     set the cable type
    //     set AN? TODO
    // else:
    //     set admin_state as ADMIN_DOWN
    //     (only admin_state is down. user_admin_state as per request msg)

    int xcvr_port =
        sdk::lib::catalog::port_num_to_qsfp_port(port_args->port_num);

    port::xcvr_port_mac_addr(xcvr_port, port_args->mac_addr);

    // default cable type is CU
    port_args->cable_type = cable_type_t::CABLE_TYPE_CU;

    if (xcvr_port != -1) {
        if (sdk::platform::xcvr_valid(xcvr_port-1) == true) {
            port_args->cable_type = sdk::platform::cable_type(xcvr_port-1);

            // for older boards, cable type returned is NONE.
            // Set it to CU
            if (port_args->cable_type == cable_type_t::CABLE_TYPE_NONE) {
                port_args->cable_type = cable_type_t::CABLE_TYPE_CU;
            }

            port_args->port_an_args =
                            sdk::platform::xcvr_get_an_args(xcvr_port - 1);

           SDK_TRACE_DEBUG("port : %u, xcvr_port : %u, user_cap : %u "
                              "fec ability : %u, fec request : %u, "
                              "cable_type : %u", port_args->port_num,
                              xcvr_port, port_args->port_an_args->user_cap,
                              port_args->port_an_args->fec_ability,
                              port_args->port_an_args->fec_request,
                              port_args->cable_type);

            // If AN=true and
            // (1) fiber cable is inserted OR
            // (2) QSA is inserted,
            // set the speed, fec, num_lanes based on cable
            if (port_args->auto_neg_enable == true) {
                if ((port_args->cable_type == cable_type_t::CABLE_TYPE_FIBER) ||
                    (sdk::platform::xcvr_type(xcvr_port-1) ==
                                            xcvr_type_t::XCVR_TYPE_SFP)) {
                    port_args->auto_neg_enable = false;
                    port_args->port_speed =
                        sdk::platform::cable_speed(xcvr_port-1);

                    // TODO FEC assumed based on cable speed
                    switch (port_args->port_speed) {
                        case port_speed_t::PORT_SPEED_100G:
                            port_args->fec_type = port_fec_type_t::PORT_FEC_TYPE_RS;
                            port_args->num_lanes = 4;
                            break;

                        case port_speed_t::PORT_SPEED_40G:
                            port_args->fec_type = port_fec_type_t::PORT_FEC_TYPE_NONE;
                            port_args->num_lanes = 4;
                            break;

                        case port_speed_t::PORT_SPEED_25G:
                            port_args->fec_type = port_fec_type_t::PORT_FEC_TYPE_FC;
                            port_args->num_lanes = 1;
                            break;

                        case port_speed_t::PORT_SPEED_10G:
                            port_args->fec_type = port_fec_type_t::PORT_FEC_TYPE_NONE;
                            port_args->num_lanes = 1;
                            break;

                        default:
                            port_args->fec_type = port_fec_type_t::PORT_FEC_TYPE_NONE;
                            break;
                    }
                }
            }
        } else {
            port_args->admin_state = port_admin_state_t::PORT_ADMIN_STATE_DOWN;
        }
    }
    return SDK_RET_OK;
}

//-----------------------------------------------------------------------------
// Validate port create config
//-----------------------------------------------------------------------------
static bool
validate_port_create (port_args_t *args)
{
    if (args->port_type == port_type_t::PORT_TYPE_NONE) {
        SDK_TRACE_ERR("Invalid port type for port: %d", args->port_num);
        return false;
    }

    if (args->port_speed == port_speed_t::PORT_SPEED_NONE) {
        SDK_TRACE_ERR("Invalid port speed for port: %d", args->port_num);
        return false;
    }

    // validate speed against num_lanes only if AN=false
    if (args->auto_neg_enable == false) {
        return validate_speed_lanes(args->port_speed, args->num_lanes);
    }

    return true;
}

static void
port_init_defaults (port_args_t *args)
{
    if (args->mtu == 0) {
        args->mtu = 9216;   // TODO define?
    }
}

//-----------------------------------------------------------------------------
// PD If Create
//-----------------------------------------------------------------------------
void *
port_create (port_args_t *args)
{
    sdk_ret_t    ret = SDK_RET_OK;
    port         *port_p = NULL;

    if (validate_port_create (args) == false) {
        // TODO return codes
        return NULL;
    }

    port_init_defaults(args);

    port_p = (port *)g_linkmgr_state->port_slab()->alloc();
    port_p->set_port_num(args->port_num);
    port_p->set_port_type(args->port_type);
    port_p->set_port_speed(args->port_speed);
    port_p->set_mac_id(args->mac_id);
    port_p->set_mac_ch(args->mac_ch);
    port_p->set_num_lanes(args->num_lanes);
    port_p->set_debounce_time(args->debounce_time);
    port_p->set_fec_type(args->fec_type);
    port_p->set_auto_neg_enable(args->auto_neg_enable);
    port_p->set_mtu(args->mtu);
    port_p->set_pause(args->pause);
    port_p->set_cable_type(args->cable_type);
    port_p->set_loopback_mode(args->loopback_mode);

    // store the user configured admin state
    port_p->set_user_admin_state(args->user_admin_state);

    // store the user configured AN
    port_p->set_auto_neg_cfg(args->auto_neg_cfg);

    // store the user configured num_lanes
    port_p->set_num_lanes_cfg(args->num_lanes_cfg);

    port_p->set_mac_fns(&mac_fns);
    port_p->set_serdes_fns(&serdes_fns);

    if(args->port_type == sdk::types::port_type_t::PORT_TYPE_MGMT) {
        port_p->set_mac_fns(&mac_mgmt_fns);
    }

    for (uint32_t i = 0; i < args->num_lanes; ++i) {
        port_p->sbus_addr_set(i, args->sbus_addr[i]);
    }

    // set the source mac addr for pause frames
    port_p->port_mac_set_pause_src_addr(args->mac_addr);

    if (args->port_an_args != NULL) {
        port_p->set_user_cap(args->port_an_args->user_cap);
        port_p->set_fec_ability(args->port_an_args->fec_ability);
        port_p->set_fec_request(args->port_an_args->fec_request);
    }

    // disable a port to invoke soft reset
    ret = port::port_disable(port_p);
    if (ret != SDK_RET_OK) {
        SDK_TRACE_ERR("port disable failed");
    }

    // if admin up is set, enable the port
    if (args->admin_state == port_admin_state_t::PORT_ADMIN_STATE_UP) {
        ret = port::port_enable(port_p);
        if (ret != SDK_RET_OK) {
            SDK_TRACE_ERR("port enable failed");
        }
    }

    return port_p;
}

//-----------------------------------------------------------------------------
// Validate port update config
//-----------------------------------------------------------------------------
static bool
validate_port_update (port *port_p, port_args_t *args)
{
    if (args->port_type != port_p->port_type()) {
        SDK_TRACE_ERR("port_type update not supported for port: %d",
                        args->port_num);
        return false;
    }

    // TODO skip num_lanes check to support auto speed set for QSA
#if 0
    if (args->num_lanes != port_p->num_lanes()) {
        SDK_TRACE_ERR("num_lanes update not supported for port: %d",
                        args->port_num);
        return false;
    }
#endif

    // validate speed against num_lanes only if AN=false
    if (args->auto_neg_enable == false) {
        return validate_speed_lanes(args->port_speed, args->num_lanes);
    }

    return true;
}

//-----------------------------------------------------------------------------
// update given port
//-----------------------------------------------------------------------------
sdk_ret_t
port_update (void *pd_p, port_args_t *args)
{
    sdk_ret_t          ret           = SDK_RET_OK;
    bool               configured    = false;
    port               *port_p       = (port *)pd_p;
    port_admin_state_t prev_admin_st = port_p->admin_state();

    if (validate_port_update (port_p, args) == false) {
        // TODO return codes
        return SDK_RET_ERR;
    }

    // check if any properties have changed

    if (args->port_speed != port_speed_t::PORT_SPEED_NONE &&
        args->port_speed != port_p->port_speed()) {
        SDK_TRACE_DEBUG("speed updated. new: %d, old: %d",
                        args->port_speed, port_p->port_speed());

        // TODO when to reset mac config?
        if (port_p->port_speed() == port_speed_t::PORT_SPEED_100G ||
            port_p->port_speed() == port_speed_t::PORT_SPEED_40G) {
            port_p->port_deinit();
        }

        port_p->set_port_speed(args->port_speed);
        configured = true;
    }

    if (args->admin_state != port_admin_state_t::PORT_ADMIN_STATE_NONE &&
        args->admin_state != port_p->admin_state()) {
        SDK_TRACE_DEBUG("admin_state updated. new: %d, old: %d",
                        args->admin_state, port_p->admin_state());

        // Do not update admin_state here
        // port_disable following configured=true will update it
        configured = true;
    }

    if (args->user_admin_state != port_admin_state_t::PORT_ADMIN_STATE_NONE &&
        args->user_admin_state != port_p->user_admin_state()) {
        SDK_TRACE_DEBUG("user_admin_state updated. new: %d, old: %d",
                        args->user_admin_state, port_p->user_admin_state());
        port_p->set_user_admin_state(args->user_admin_state);
    }

    if (args->cable_type != cable_type_t::CABLE_TYPE_NONE &&
        args->cable_type != port_p->cable_type()) {
        SDK_TRACE_DEBUG("cable_type updated. new: %d, old: %d",
                        args->cable_type, port_p->cable_type());
        port_p->set_cable_type(args->cable_type);
        configured = true;
    }

    // FEC_TYPE_NONE is valid
    if (args->fec_type != port_p->fec_type()) {
        SDK_TRACE_DEBUG("fec updated. new: %d, old: %d",
                        args->fec_type, port_p->fec_type());
        port_p->set_fec_type(args->fec_type);
        configured = true;
    }

    // MTU 0 is invalid
    if (args->mtu != 0 &&
        args->mtu != port_p->mtu()) {
        SDK_TRACE_DEBUG("mtu updated. new: %d, old: %d",
                        args->mtu, port_p->mtu());
        port_p->set_mtu(args->mtu);
        configured = true;
    }

    if (args->debounce_time != port_p->debounce_time()) {
        SDK_TRACE_DEBUG("Debounce updated. new: %d, old: %d",
                        args->debounce_time, port_p->debounce_time());
        port_p->set_debounce_time(args->debounce_time);
        configured = true;
    }

    if (args->auto_neg_enable != port_p->auto_neg_enable()) {
        SDK_TRACE_DEBUG("AN updated. new: %d, old: %d",
                        args->auto_neg_enable, port_p->auto_neg_enable());
        port_p->set_auto_neg_enable(args->auto_neg_enable);
        configured = true;
    }

    if (args->auto_neg_cfg != port_p->auto_neg_cfg()) {
        SDK_TRACE_DEBUG("AN cfg updated. new: %d, old: %d",
                        args->auto_neg_cfg, port_p->auto_neg_cfg());
        port_p->set_auto_neg_cfg(args->auto_neg_cfg);
        configured = true;
    }

    // TODO update num_lanes to support auto speed set for QSA
    if (args->num_lanes != port_p->num_lanes()) {
        SDK_TRACE_DEBUG("num_lanes updated. new: %d, old: %d",
                        args->num_lanes, port_p->num_lanes());
        port_p->set_num_lanes(args->num_lanes);
        configured = true;
    }

    if (args->num_lanes_cfg != port_p->num_lanes_cfg()) {
        SDK_TRACE_DEBUG("num_lanes_cfg updated. new: %d, old: %d",
                        args->num_lanes_cfg, port_p->num_lanes_cfg());
        port_p->set_num_lanes_cfg(args->num_lanes_cfg);
        configured = true;
    }

    if (args->loopback_mode != port_p->loopback_mode()) {
        SDK_TRACE_DEBUG(
            "loopback updated. new: %s, old: %s",
            args->loopback_mode ==
                port_loopback_mode_t::PORT_LOOPBACK_MODE_MAC? "mac" :
            args->loopback_mode ==
                port_loopback_mode_t::PORT_LOOPBACK_MODE_PHY? "phy" : "none",
            port_p->loopback_mode() ==
                port_loopback_mode_t::PORT_LOOPBACK_MODE_MAC? "mac" :
            port_p->loopback_mode() ==
                port_loopback_mode_t::PORT_LOOPBACK_MODE_PHY? "phy" : "none");

        port_p->set_loopback_mode(args->loopback_mode);
        configured = true;
    }

    if (args->pause != port_p->pause()) {
        SDK_TRACE_DEBUG("Pause updated. new: %d, old: %d",
                        args->pause, port_p->pause());
        port_p->set_pause(args->pause);
        configured = true;
    }

    // TODO need check?
    if (args->mac_stats_reset == true) {
        SDK_TRACE_DEBUG("Resetting MAC stats");
        port_p->port_mac_stats_reset(args->mac_stats_reset);
        port_p->port_mac_stats_reset(false);
    }

    if (args->port_an_args != NULL) {
        port_p->set_user_cap(args->port_an_args->user_cap);
        port_p->set_fec_ability(args->port_an_args->fec_ability);
        port_p->set_fec_request(args->port_an_args->fec_request);
    }

    // Disable the port if any config has changed
    if (configured == true) {
        ret = port::port_disable(port_p);
    }

    // Enable the port if -
    //      admin-up state is set in request msg OR
    //      admin state is not set in request msg, but port was admin up
    switch(args->admin_state) {
    case port_admin_state_t::PORT_ADMIN_STATE_NONE:
        if (prev_admin_st == port_admin_state_t::PORT_ADMIN_STATE_UP) {
            ret = port::port_enable(port_p);
        }
        break;

    case port_admin_state_t::PORT_ADMIN_STATE_DOWN:
        ret = port::port_disable(port_p);
        break;

    case port_admin_state_t::PORT_ADMIN_STATE_UP:
        ret = port::port_enable(port_p);
        break;

    default:
        break;
    }

    return ret;
}

//-----------------------------------------------------------------------------
// delete given port by disabling the port and then deleting the port instance
//-----------------------------------------------------------------------------
sdk_ret_t
port_delete (void *pd_p)
{
    sdk_ret_t    ret = SDK_RET_OK;
    port         *port_p = (port *)pd_p;

    SDK_TRACE_DEBUG("port delete");

    port_p->port_deinit();

    ret = port::port_disable(port_p);
    g_linkmgr_state->port_slab()->free(port_p);

    return ret;
}

//-----------------------------------------------------------------------------
// PD Port get
//-----------------------------------------------------------------------------
sdk_ret_t
port_get (void *pd_p, port_args_t *args)
{
    port      *port_p = (port *)pd_p;

    args->port_type   = port_p->port_type();
    args->port_speed  = port_p->port_speed();
    args->admin_state = port_p->admin_state();
    args->mac_id      = port_p->mac_id();
    args->mac_ch      = port_p->mac_ch();
    args->num_lanes   = port_p->num_lanes();
    args->fec_type    = port_p->fec_type();
    args->mtu         = port_p->mtu();
    args->debounce_time    = port_p->debounce_time();
    args->auto_neg_enable  = port_p->auto_neg_enable();
    args->auto_neg_cfg     = port_p->auto_neg_cfg();
    args->user_admin_state = port_p->user_admin_state();
    args->num_lanes_cfg    = port_p->num_lanes_cfg();
    args->pause       = port_p->pause();
    args->link_sm     = port_p->port_link_sm();
    args->loopback_mode = port_p->loopback_mode();
    args->num_link_down = port_p->num_link_down();

    if (args->link_sm == port_link_sm_t::PORT_LINK_SM_AN_CFG &&
        port_p->auto_neg_enable() == true) {
        args->link_sm = port_p->port_link_an_sm();
    } else if (args->link_sm == port_link_sm_t::PORT_LINK_SM_DFE_TUNING) {
        args->link_sm = port_p->port_link_dfe_sm();
    }

    // TODO send live link status until poll timer is reduced
    // args->oper_status = port_p->oper_status();
    if (port_p->port_link_status() == true) {
        args->oper_status = port_oper_status_t::PORT_OPER_STATUS_UP;
    } else {
        args->oper_status = port_oper_status_t::PORT_OPER_STATUS_DOWN;
    }

    if (NULL != args->stats_data) {
        port_p->port_mac_stats_get (args->stats_data);
    }

    return SDK_RET_OK;
}

#if 0
void *
port_make_clone (void *pd_orig_p)
{
    port  *port_p;
    port  *pd_new_clone_p;

    SDK_TRACE_DEBUG("port clone");
    port_p = (port *)pd_orig_p;
    pd_new_clone_p =
        (port *)g_linkmgr_state->port_slab()->alloc();
    // populate cloned pd instance from existing pd instance
    pd_new_clone_p->set_oper_status(port_p->oper_status());
    pd_new_clone_p->set_port_speed(port_p->port_speed());
    pd_new_clone_p->set_port_type(port_p->port_type());
    pd_new_clone_p->set_admin_state(port_p->admin_state());

    pd_new_clone_p->set_port_link_sm(port_p->port_link_sm());
    pd_new_clone_p->set_link_bring_up_timer(port_p->link_bring_up_timer());

    pd_new_clone_p->set_mac_id(port_p->mac_id());
    pd_new_clone_p->set_mac_ch(port_p->mac_ch());
    pd_new_clone_p->set_num_lanes(port_p->num_lanes());

    return pd_new_clone_p;
}
#endif

//-----------------------------------------------------------------------------
// PD Port mem free
//-----------------------------------------------------------------------------
sdk_ret_t
port_mem_free (port_args_t *args)
{
    sdk_ret_t    ret = SDK_RET_OK;
    port         *port_p = (port *)args->port_p;

    g_linkmgr_state->port_slab()->free(port_p);

    return ret;
}

bool
port_has_speed_changed (port_args_t *args)
{
    port *port_p = (port *)args->port_p;
    return (args->port_speed != port_p->port_speed());
}

bool
port_has_admin_state_changed (port_args_t *args)
{
    port *port_p = (port *)args->port_p;
    return (args->admin_state != port_p->admin_state());
}

}    // namespace linkmgr
}    // namespace sdk
