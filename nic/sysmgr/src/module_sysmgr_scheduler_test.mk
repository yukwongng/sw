# {C} Copyright 2018 Pensando Systems Inc. All rights reserved
include ${MKDEFS}/pre.mk

MODULE_TARGET   = sysmgr_scheduler_test.gtest
MODULE_ARCH     = x86_64
MODULE_SOLIBS   = delphisdk
MODULE_LDLIBS   = ${NIC_THIRDPARTY_GOOGLE_LDLIBS} rt ev
MODULE_ARLIBS   = sysmgrproto delphiproto
INCLUDE_FILES   = $(wildcard ${MODULE_SRC_DIR}/*.cpp)
EXCLUDE_FILES   = ${MODULE_SRC_DIR}/main.cpp $(wildcard ${MODULE_SRC_DIR}/*_test.cpp) 
MODULE_SRCS     = ${MODULE_SRC_DIR}/scheduler_test.cpp \
	              $(filter-out $(EXCLUDE_FILES), $(INCLUDE_FILES))

include ${MKDEFS}/post.mk
