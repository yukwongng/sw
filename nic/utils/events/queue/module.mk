# {C} Copyright 2018 Pensando Systems Inc. All rights reserved
include ${MKDEFS}/pre.mk
MODULE_TARGET   = libevents_queue.so
MODULE_PIPELINE = iris gft
MODULE_SOLIBS   = ipc
include ${MKDEFS}/post.mk