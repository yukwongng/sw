CUR_DIR:=$(PWD)
STORAGE_DIR:=$(shell dirname ${CUR_DIR})
SW_DIR:=$(shell dirname ${STORAGE_DIR})
CUR_USER:=$(shell whoami)
CUR_TIME:=$(shell date +%Y-%m-%d_%H.%M.%S)
CONTAINER_NAME:=${CUR_USER}_${CUR_TIME}

# get a shell with the dependencies image loaded, with the host filesystem mounted.
shell: build-runtime-image
	mkdir -p ${SW_DIR}/bazel-cache;
	docker run -it --rm --sysctl net.ipv6.conf.all.disable_ipv6=0 --privileged --name ${CONTAINER_NAME} -v $(SW_DIR):/sw -v $(SW_DIR)/bazel-cache:/root/.cache -v /lib/modules:/lib/modules -v /usr/src/kernels:/usr/src/kernels -w /sw/storage/offload pensando/offload bash

REGISTRY = registry.test.pensando.io:5000

build-runtime-image: install_box
	if [ "x${NO_PULL}" = "x" ]; then docker pull $(REGISTRY)/pensando/nic:1.35; fi
	cd ../.. && BOX_INCLUDE_ENV="NO_COPY" NO_COPY=1 box -t pensando/offload storage/offload/box.rb

install_box:
	@if [ ! -x /usr/local/bin/box ]; then echo "Installing box, sudo is required"; curl -sSL box-builder.sh | sudo bash; fi

SONICSIM = sonic-sim
SONICSIM_INCDIRS = include src/osal src/osal/test src/sim third-party/freebsd-crypto/src

$(SONICSIM)-y = src/sim/pnso_sim_api.o
$(SONICSIM)-y += src/sim/lzrw1-a.o
$(SONICSIM)-y += src/sim/sim_algo.o
$(SONICSIM)-y += src/sim/sim_chain.o
$(SONICSIM)-y += src/sim/sim_crypto.o
$(SONICSIM)-y += src/sim/sim_util.o
$(SONICSIM)-y += src/sim/sim_worker.o
$(SONICSIM)-y += third-party/freebsd-crypto/src/crc32.o
$(SONICSIM)-y += third-party/freebsd-crypto/src/rijndael-alg-fst.o
$(SONICSIM)-y += third-party/freebsd-crypto/src/rijndael-api.o
$(SONICSIM)-y += third-party/freebsd-crypto/src/rijndael-api-fst.o
$(SONICSIM)-y += third-party/freebsd-crypto/src/sha256c.o
$(SONICSIM)-y += third-party/freebsd-crypto/src/sha512c.o
$(SONICSIM)-y += third-party/freebsd-crypto/src/xform_aes_xts.o

$(SONICSIM)-y += src/osal/linux/kernel/osal_logger.o 
$(SONICSIM)-y += src/osal/linux/kernel/osal_sys.o 
$(SONICSIM)-y += src/osal/linux/kernel/osal_mem.o 
$(SONICSIM)-y += src/osal/linux/kernel/osal_thread.o 
$(SONICSIM)-y += src/osal/linux/kernel/osal_atomic.o 
ifdef ENABLE_PNSO_SONIC_TEST
$(SONICSIM)-y += src/osal/test/osal_test.o
endif

SONIC = sonic
SONIC_INCDIRS = include src/osal src/osal/test src/common src/drv src/drv/linux/kernel src/drv/test

$(SONIC)-y = src/drv/pnso_mpool.o
$(SONIC)-y += src/drv/pnso_api.o
$(SONIC)-y += src/common/pnso_svc.o
$(SONIC)-y += src/drv/pnso_req.o
$(SONIC)-y += src/drv/pnso_init.o
$(SONIC)-y += src/drv/pnso_batch.o
$(SONIC)-y += src/drv/pnso_chain.o
$(SONIC)-y += src/drv/pnso_cp.o
$(SONIC)-y += src/drv/pnso_dc.o
$(SONIC)-y += src/drv/pnso_chksum.o
$(SONIC)-y += src/drv/pnso_hash.o
$(SONIC)-y += src/drv/pnso_cpdc.o
$(SONIC)-y += src/drv/pnso_cpdc_cmn.o
$(SONIC)-y += src/drv/pnso_crypto.o
$(SONIC)-y += src/drv/pnso_crypto_ops.o
$(SONIC)-y += src/drv/pnso_crypto_cmn.o
$(SONIC)-y += src/drv/pnso_decompact.o
$(SONIC)-y += src/drv/pnso_seq.o
$(SONIC)-y += src/drv/pnso_seq_hw.o
$(SONIC)-y += src/drv/pnso_utils.o
$(SONIC)-y += src/drv/pnso_stats.o
$(SONIC)-y += src/drv/linux/kernel/sonic_bus_pci.o
$(SONIC)-y += src/drv/linux/kernel/sonic_dev.o
$(SONIC)-y += src/drv/linux/kernel/sonic_lif.o
$(SONIC)-y += src/drv/linux/kernel/sonic_interrupt.o
$(SONIC)-y += src/drv/linux/kernel/sonic_main.o
$(SONIC)-y += src/drv/linux/kernel/sonic_api_int.o

$(SONIC)-y += src/osal/linux/kernel/osal_logger.o 
$(SONIC)-y += src/osal/linux/kernel/osal_sys.o 
$(SONIC)-y += src/osal/linux/kernel/osal_atomic.o
$(SONIC)-y += src/osal/linux/kernel/osal_mem.o
$(SONIC)-y += src/osal/linux/kernel/osal_rmem.o
$(SONIC)-y += src/osal/linux/kernel/osal_thread.o

ifdef ENABLE_PNSO_SONIC_TEST
$(SONIC)-y += src/drv/test/pnso_sonic_test.o
endif
$(SONIC)-y += src/common/pnso_pbuf.o


PENCAKE = pencake
PENCAKE_INCDIRS = include src/osal src/test src/drv/linux/kernel src/common third-party/libyaml/include

$(PENCAKE)-y = src/test/pnso_test_module.o
$(PENCAKE)-y += src/test/pnso_test_sysfs.o
$(PENCAKE)-y += src/test/pnso_test_parse.o
$(PENCAKE)-y += src/test/pnso_test_util.o
$(PENCAKE)-y += src/test/pnso_test.o
$(PENCAKE)-y += src/common/pnso_pbuf.o

$(PENCAKE)-y += src/osal/linux/kernel/osal_logger.o
$(PENCAKE)-y += src/osal/linux/kernel/osal_sys.o
$(PENCAKE)-y += src/osal/linux/kernel/osal_atomic.o
$(PENCAKE)-y += src/osal/linux/kernel/osal_thread.o
$(PENCAKE)-y += src/osal/linux/kernel/osal_mem.o
$(PENCAKE)-y += src/osal/linux/kernel/osal_random.o

$(PENCAKE)-y += third-party/libyaml/src/api.o
$(PENCAKE)-y += third-party/libyaml/src/loader.o
$(PENCAKE)-y += third-party/libyaml/src/parser.o
$(PENCAKE)-y += third-party/libyaml/src/reader.o
$(PENCAKE)-y += third-party/libyaml/src/scanner.o
$(PENCAKE)-y += third-party/libyaml/src/writer.o

KERN_DIR ?= /lib/modules/$(shell uname -r)/build
KERN_SRC ?= /lib/modules/$(shell uname -r)/source

sonic-sim-var: $(eval obj-m += $(SONICSIM).o)
sonic-sim-var: $(eval ccflags-y += $(SONICSIM_INCDIRS:%=-I$(PWD)/%))
ifdef ENABLE_PNSO_SONIC_TEST
sonic-sim-var: $(eval ccflags-y += -DENABLE_PNSO_SONIC_TEST)
endif
sonic-sim: sonic-sim-var modules-compile

sonic-var: $(eval obj-m += $(SONIC).o) 
sonic-var: $(eval ccflags-y += $(SONIC_INCDIRS:%=-I$(PWD)/%))
sonic-var: $(eval ccflags-y += -I$(KERN_DIR)/include/linux)
sonic-var: $(eval ccflags-y += -DNDEBUG)
sonic-var: $(eval ccflags-y += -Werror)
sonic-var: $(eval ccflags-y += -Wall)
ifdef ENABLE_PNSO_SONIC_TEST
sonic-var: $(eval ccflags-y += -DENABLE_PNSO_SONIC_TEST)
endif
sonic: sonic-var modules-compile

pencake-var: $(eval obj-m += $(PENCAKE).o)
pencake-var: $(eval ccflags-y += $(PENCAKE_INCDIRS:%=-I$(PWD)/%))
pencake-var: $(eval ccflags-y += -I$(KERN_DIR)/include/linux)
pencake-var: $(eval ccflags-y += -I$(KERN_SRC)/include/linux)
pencake-var: $(eval ccflags-y += -DHAVE_CONFIG_H=0)
pencake: pencake-var modules-compile

modules-compile: 
	make -C $(KERN_DIR) M=$(PWD) modules

modules: sonic-var sonic-sim-var pencake-var modules-compile

modules-clean:
	make -C $(KERN_DIR) M=$(PWD) clean

