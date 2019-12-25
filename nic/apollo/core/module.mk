# {C} Copyright 2018 Pensando Systems Inc. All rights reserved

include ${MKDEFS}/pre.mk
MODULE_TARGET   = libpdscore.so
MODULE_PIPELINE = apollo artemis apulu athena
MODULE_SOLIBS   = pdsnicmgr pdspciemgr pciemgrd pcieport event_thread \
                  pdslearn pdsfte
include ${MKDEFS}/post.mk
