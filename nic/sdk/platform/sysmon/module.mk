# {C} Copyright 2019 Pensando Systems Inc. All rights reserved

include ${MKDEFS}/pre.mk
MODULE_TARGET   = libsysmon.so
MODULE_PIPELINE = iris
MODULE_SRCS     = $(wildcard ${MODULE_SRC_DIR}/*.cc)
MODULE_FLAGS    = -DCAPRI_SW ${NIC_CSR_FLAGS}
include ${MKDEFS}/post.mk
