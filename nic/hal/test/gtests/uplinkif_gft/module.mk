# {C} Copyright 2018 Pensando Systems Inc. All rights reserved
include ${MKDEFS}/pre.mk
MODULE_TARGET   = uplinkif_gft_test.gtest
MODULE_PIPELINE = gft
MODULE_SOLIBS   = ${NIC_HAL_GTEST_SOLIBS}
MODULE_LDLIBS   = ${NIC_HAL_GTEST_LDLIBS}
MODULE_ARCH     = x86_64
include ${MKDEFS}/post.mk
