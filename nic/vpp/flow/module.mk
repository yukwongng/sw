# {C} Copyright 2019 Pensando Systems Inc. All rights reserved

include ${MKDEFS}/pre.mk
MODULE_TARGET    = flow_plugin.lib
MODULE_PIPELINE  = apollo artemis apulu
MODULE_PREREQS   = vpp_pkg.export
MODULE_SOLIBS    = pdsvpp_impl pdsvpp_api pdsvpp_ipc pdsvpp_cfg
MODULE_LDLIBS    = ${SDK_THIRD_PARTY_VPP_LIBS} ${NIC_THIRDPARTY_GOOGLE_LDLIBS}
MODULE_INCS      = ${VPP_PLUGINS_INCS} ${BLD_PROTOGEN_DIR}
MODULE_DEFS      = ${VPP_DEFINES_${ARCH}}
MODULE_FLAGS     = ${VPP_FLAGS_${ARCH}}
include ${MKDEFS}/post.mk
