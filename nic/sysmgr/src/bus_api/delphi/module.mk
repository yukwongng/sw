# {C} Copyright 2018 Pensando Systems Inc. All rights reserved

include ${MKDEFS}/pre.mk
MODULE_TARGET   = libsysmgrbus.so
MODULE_PIPELINE = iris
MODULE_SOLIBS   = delphisdk
MODULE_LDLIBS   = ${NIC_THIRDPARTY_GOOGLE_LDLIBS} rt ev pthread z dl
MODULE_ARLIBS   = sysmgrproto delphiproto
ALL_CC_FILES    = $(wildcard ${MODULE_SRC_DIR}/*.cpp)
ALL_TEST_FILES  = $(wildcard ${MODULE_SRC_DIR}/*_test.cpp)
MODULE_SRCS     = $(filter-out $(ALL_TEST_FILES), $(ALL_CC_FILES))
include ${MKDEFS}/post.mk