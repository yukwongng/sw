# {C} Copyright 2018 Pensando Systems Inc. All rights reserved
include ${MKDEFS}/pre.mk

MODULE_TARGET       = delphi.proto
MODULE_GEN_TYPES    = CC PY
MODULE_INCS         = ${MODULE_DIR} \
                      ${TOPDIR}/nic \
                      ${TOPDIR}/nic/hal/third-party/google/include \
                      /usr/local/include
MODULES_PREREQS = protoc-gen-delphi.gobin
MODULE_LDLIBS       = pthread
MODULE_POSTGEN_MK   = module_protolib.mk

include ${MKDEFS}/post.mk
