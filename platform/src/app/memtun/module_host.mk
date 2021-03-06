# {C} Copyright 2018 Pensando Systems Inc. All rights reserved
include ${MKDEFS}/pre.mk
MODULE_ARCH     := x86_64
MODULE_TARGET   := memtun_host.bin
MODULE_SRCS     := ${MODULE_SRC_DIR}/host.c \
                    ${MODULE_SRC_DIR}/pktloop.c \
                    ${MODULE_SRC_DIR}/queue.c \
                    ${MODULE_SRC_DIR}/tun.c \
                    ${MODULE_SRC_DIR}/util.c
MODULE_LDLIBS   := pthread
include ${MKDEFS}/post.mk
