# {C} Copyright 2018 Pensando Systems Inc. All rights reserved

include ${MKDEFS}/pre.mk
MODULE_TARGET   := penctl.submake
MODULE_PIPELINE := iris
ifeq ($(PIPELINE), iris)
MODULE_PREREQS = nmd_halproto.submake agent_irisproto.submake agent_halproto.submake
endif
include ${MKDEFS}/post.mk
