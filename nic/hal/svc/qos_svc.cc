//------------------------------------------------------------------------------
// qos service implementation
// Vasanth Kumar (Pensando Systems)
//------------------------------------------------------------------------------
//
#include "nic/include/base.h"
#include "nic/include/trace.hpp"
#include "nic/hal/svc/qos_svc.hpp"
#include "nic/hal/src/qos.hpp"

Status
QOSServiceImpl::BufPoolCreate(ServerContext *context,
                              const BufPoolRequestMsg *req,
                              BufPoolResponseMsg *rsp)
{
    uint32_t             i, nreqs = req->request_size();
    BufPoolResponse      *response;

    HAL_TRACE_DEBUG("Rcvd BufPool Create Request");
    if (nreqs == 0) {
        return Status(grpc::StatusCode::INVALID_ARGUMENT, "Empty Request");
    }

    hal::hal_cfg_db_open(hal::CFG_OP_WRITE);
    for (i = 0; i < nreqs; i++) {
        response = rsp->add_response();
        auto spec = req->request(i);
        hal::buf_pool_create(spec, response);
    }
    hal::hal_cfg_db_close();
    return Status::OK;
}

Status
QOSServiceImpl::BufPoolUpdate(ServerContext *context,
                              const BufPoolRequestMsg *req,
                              BufPoolResponseMsg *rsp)
{
    HAL_TRACE_DEBUG("Rcvd BufPool Update Request");
    return Status::OK;
}

Status
QOSServiceImpl::BufPoolDelete(ServerContext *context,
                              const BufPoolDeleteRequestMsg *req,
                              BufPoolDeleteResponseMsg *rsp)
{
    HAL_TRACE_DEBUG("Rcvd BufPool Update Request");
    return Status::OK;
}

Status
QOSServiceImpl::BufPoolGet(ServerContext *context,
                           const BufPoolGetRequestMsg *req,
                           BufPoolGetResponseMsg *rsp)
{
    HAL_TRACE_DEBUG("Rcvd BufPool Get Request");
    return Status::OK;
}

Status
QOSServiceImpl::QueueCreate(ServerContext *context,
                            const QueueRequestMsg *req,
                            QueueResponseMsg *rsp)
{
    uint32_t           i, nreqs = req->request_size();
    QueueResponse      *response;

    HAL_TRACE_DEBUG("Rcvd Queue Create Request");
    if (nreqs == 0) {
        return Status(grpc::StatusCode::INVALID_ARGUMENT, "Empty Request");
    }

    hal::hal_cfg_db_open(hal::CFG_OP_WRITE);
    for (i = 0; i < nreqs; i++) {
        response = rsp->add_response();
        auto spec = req->request(i);
        hal::queue_create(spec, response);
    }
    hal::hal_cfg_db_close();
    return Status::OK;
}

Status
QOSServiceImpl::QueueUpdate(ServerContext *context,
                              const QueueRequestMsg *req,
                              QueueResponseMsg *rsp)
{
    HAL_TRACE_DEBUG("Rcvd Queue Update Request");
    return Status::OK;
}

Status
QOSServiceImpl::QueueDelete(ServerContext *context,
                              const QueueDeleteRequestMsg *req,
                              QueueDeleteResponseMsg *rsp)
{
    HAL_TRACE_DEBUG("Rcvd Queue Update Request");
    return Status::OK;
}

Status
QOSServiceImpl::QueueGet(ServerContext *context,
                           const QueueGetRequestMsg *req,
                           QueueGetResponseMsg *rsp)
{
    HAL_TRACE_DEBUG("Rcvd Queue Get Request");
    return Status::OK;
}

Status
QOSServiceImpl::PolicerCreate(ServerContext *context,
                              const PolicerRequestMsg *req,
                              PolicerResponseMsg *rsp)
{
    uint32_t           i, nreqs = req->request_size();
    PolicerResponse    *response;

    HAL_TRACE_DEBUG("Rcvd Policer Create Request");
    if (nreqs == 0) {
        return Status(grpc::StatusCode::INVALID_ARGUMENT, "Empty Request");
    }

    hal::hal_cfg_db_open(hal::CFG_OP_WRITE);
    for (i = 0; i < nreqs; i++) {
        response = rsp->add_response();
        auto spec = req->request(i);
        hal::policer_create(spec, response);
    }
    hal::hal_cfg_db_close();
    return Status::OK;
}

Status
QOSServiceImpl::PolicerUpdate(ServerContext *context,
                              const PolicerRequestMsg *req,
                              PolicerResponseMsg *rsp)
{
    HAL_TRACE_DEBUG("Rcvd Policer Update Request");
    return Status::OK;
}

Status
QOSServiceImpl::PolicerDelete(ServerContext *context,
                              const PolicerDeleteRequestMsg *req,
                              PolicerDeleteResponseMsg *rsp)
{
    HAL_TRACE_DEBUG("Rcvd Policer Update Request");
    return Status::OK;
}

Status
QOSServiceImpl::PolicerGet(ServerContext *context,
                           const PolicerGetRequestMsg *req,
                           PolicerGetResponseMsg *rsp)
{
    HAL_TRACE_DEBUG("Rcvd Policer Get Request");
    return Status::OK;
}
