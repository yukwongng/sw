//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// This file contains the base class implementation of test classes
///
//----------------------------------------------------------------------------

#include <stdarg.h>
#include "base.hpp"
#include "nic/apollo/include/api/oci.hpp"
#include "nic/apollo/include/api/oci_init.hpp"

/// Callback invoked for debug traces
///
/// This is sample implementation, hence doesn't check whether
//  user enabled traces at what level, it always prints the traces
/// but with a simple header prepended that tells what level the
/// trace is spwed at ... in reality, you call your favorite logger here
static int
trace_cb (sdk_trace_level_e trace_level, const char *format, ...)
{
    char logbuf[1024];
    va_list args;

    va_start(args, format);
    vsnprintf(logbuf, sizeof(logbuf), format, args);
    va_end(args);

    switch (trace_level) {
    case sdk::lib::SDK_TRACE_LEVEL_NONE:
        return 0;
        break;
    case sdk::lib::SDK_TRACE_LEVEL_ERR:
        fprintf(stderr, "[E] %s\n", logbuf);
        break;
    case sdk::lib::SDK_TRACE_LEVEL_WARN:
        fprintf(stderr, "[W] %s\n", logbuf);
        break;
    case sdk::lib::SDK_TRACE_LEVEL_INFO:
        fprintf(stdout, "[I] %s\n", logbuf);
        break;
    case sdk::lib::SDK_TRACE_LEVEL_DEBUG:
        fprintf(stdout, "[D] %s\n", logbuf);
        break;
    case sdk::lib::SDK_TRACE_LEVEL_VERBOSE:
        // fprintf(stdout, "[V] %s\n", logbuf);
        break;
    default:
        break;
    }
    fflush(stdout);
    return 0;
}

/// Called at the beginning of all test cases in this class,
/// initialize OCI HAL
void
oci_test_base::SetUpTestCase(const char *cfg_file, bool enable_fte)
{
    oci_init_params_t init_params;

    memset(&init_params, 0, sizeof(init_params));
    init_params.init_mode = OCI_INIT_MODE_COLD_START;
    init_params.trace_cb  = trace_cb;
    init_params.pipeline  = "apollo";
    init_params.cfg_file  = std::string(cfg_file);
    oci_init(&init_params);
}

/// Called at the end of all test cases in this class,
/// cleanup OCI HAL and quit
void
oci_test_base::TearDownTestCase(void)
{
    // oci_teardown();
}
