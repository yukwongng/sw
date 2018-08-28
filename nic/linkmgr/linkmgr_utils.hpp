// {C} Copyright 2017 Pensando Systems Inc. All rights reserved

#ifndef __LINKMGR_LINKMGR_UTILS_HPP__
#define __LINKMGR_LINKMGR_UTILS_HPP__

#include "nic/gen/proto/hal/types.pb.h"
#include "nic/include/ip.hpp"
#include "nic/gen/proto/hal/port.pb.h"
#include "sdk/types.hpp"

using types::ApiStatus;

namespace linkmgr {

port_speed_t port_speed_spec_to_sdk_port_speed(::port::PortSpeed speed);
::port::PortSpeed sdk_port_speed_to_port_speed_spec(port_speed_t speed);

port_type_t port_type_spec_to_sdk_port_type(::port::PortType type);
::port::PortType sdk_port_type_to_port_type_spec(port_type_t type);

::port::PortAdminState
sdk_port_admin_st_to_port_admin_st_spec(port_admin_state_t admin_st);
port_admin_state_t
port_admin_st_spec_to_sdk_port_admin_st(::port::PortAdminState admin_st);

::port::PortFecType
sdk_port_fec_type_to_port_fec_type_spec(port_fec_type_t fec_type);
port_fec_type_t
port_fec_type_spec_to_sdk_port_fec_type(::port::PortFecType fec_type);

::port::PortOperStatus
sdk_port_oper_st_to_port_oper_st_spec(port_oper_status_t oper_st);

::port::PortBreakoutMode
sdk_port_breakout_mode_to_port_breakout_mode_spec(
                            port_breakout_mode_t breakout_mode);

}    // namespace linkmgr

#endif    // __LINKMGR_LINKMGR_UTILS_HPP__

