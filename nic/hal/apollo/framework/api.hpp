/**
 * Copyright (c) 2018 Pensando Systems, Inc.
 *
 * @file    api.hpp
 *
 * @brief   basic types for API processing
 */

#if !defined (__API_HPP__)
#define __API_HPP__

namespace api {

/**< forward declaration */
typedef struct api_ctxt_s api_ctxt_t;

/**
 * @brief    API operation
 */
typedef enum api_op_e {
    API_OP_NONE,
    API_OP_CREATE,
    API_OP_DELETE,
    API_OP_UPDATE,
    API_OP_GET,
} api_op_t;

/**
 * @brief    object identifiers
 */
typedef enum obj_id_e {
    OBJ_ID_NONE,
    OBJ_ID_VCN,
    OBJ_ID_SUBNET,
    OBJ_ID_VNIC,
    OBJ_ID_ROUTE,
    API_ID_SECURITY_RULES,
    OBJ_ID_TEP,
    OBJ_ID_SWITCHPORT,
    OBJ_ID_MAX,
} obj_id_t;

}    // namespace api

#endif    /** __API_HPP__ */

