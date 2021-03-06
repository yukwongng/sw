TOPDIR = ../../../..
GEN_DIR = $(TOPDIR)/gen
CUSTOM = $(GEN_DIR)/.tcp_rx_asm
CUSTOM_CLEAN = myclean

include $(TOPDIR)/make/Makefile.inc
include $(TOPDIR)/make/Make.inc

PROJ = tcp_proxy_rxdma
OBJ_DIR = $(TOPDIR)/obj/p4plus_bin
GEN_DIR = $(TOPDIR)/gen
ASM_HDR_DIR = $(GEN_DIR)/$(PROJ)/asm_out/
ALT_ASM_HDR_DIR = $(GEN_DIR)/$(PROJ)/alt_asm_out/
P4_DIR = $(TOPDIR)/p4/proxy
COMMON_P4_DIR = $(TOPDIR)/p4/common-p4+
SRC_DIR = .

CAPSIM_BIN = ../../../../asic/capri/model/capsim-gen/bin/capsim

SRC_FILES := \
    $(SRC_DIR)/tcp-ack.s \
    $(SRC_DIR)/tcp-ack-slow.s \
    $(SRC_DIR)/tcp-rdesc-alloc.s \
    $(SRC_DIR)/tcp-read-rnmdr-alloc-idx.s \
    $(SRC_DIR)/tcp-read-rnmdr-fc.s \
    $(SRC_DIR)/tcp-read-tx2rx-shared.s \
    $(SRC_DIR)/tcp-read-tx2rx-shared-ext.s \
    $(SRC_DIR)/tcp-rpage-alloc.s \
    $(SRC_DIR)/tcp-rtt.s \
    $(SRC_DIR)/tcp-rx.s \
    $(SRC_DIR)/tcp-cc.s \
    $(SRC_DIR)/tcp-cc-newreno.s \
    $(SRC_DIR)/tcp-fc.s \
    $(SRC_DIR)/tcp-rx-stats.s \
    $(SRC_DIR)/tcp-slow-write-serq.s \
    $(SRC_DIR)/tcp-write-serq.s \
    $(SRC_DIR)/tcp-write-serq2.s \
    $(SRC_DIR)/tcp-write-serq3.s \
    $(SRC_DIR)/tcp-write-arq.s \

BIN_FILES := $(patsubst %.s, %.bin, $(SRC_FILES))

FILES = $(basename $(notdir $(BIN_FILES)))
COMP_BIN_FILES = $(FILES:%=$(OBJ_DIR)/%.bin)

$(GEN_DIR)/.tcp_rx_asm: $(OBJ_DIR)/.dir $(COMP_BIN_FILES)
	touch $@

DEPS_DIRS = $(subst -I,, $(INC_DIRS))
DEPS = $(shell find $(DEPS_DIRS) -name '*.h')
$(OBJ_DIR)/%.bin: %.s $(DEPS) $(CAPAS_BIN)
	$(CAPAS_BIN) $(CAPAS_COV) $(INC_DIRS) $(AS_DEFINES) $< -o $@

TEST_FILES := \
    $(SRC_DIR)/ctl/tcp-rx.ctl \
    $(SRC_DIR)/ctl/tcp-write-serq.ctl \
    $(SRC_DIR)/ctl/tcp-cc.ctl \
    $(SRC_DIR)/ctl/tcp-rtt-first.ctl \
    $(SRC_DIR)/ctl/tcp-rtt-non-first.ctl \
    $(SRC_DIR)/ctl/tcp-fastretx-alert-cwr-end.ctl

TMP_FILES  := $(patsubst %.ctl, %.x, $(TEST_FILES))

tcp-rx_ARGS :=
tcp-write-serq_ARGS :=
tcp-cc_ARGS :=
tcp-rtt_ARGS :=
tcp-rtt-non-first_ARGS :=
tcp-fastretx-alert_cwr_end_ARGS :=

$(TMP_FILES): %.x: %.ctl
	@echo Testing ...
	@echo $<
	$(CAPSIM_BIN) -e ${$(basename ${<})_ARGS} $(INC_DIRS) $< $(OBJ_DIR)/$(patsubst %.x,%.bin,$@)

test: $(TMP_FILES)

run: $(COMP_BIN_FILES)
	$(CAPSIM_BIN) -c $(INC_DIRS) $(SRC_DIR)/ctl/tcp-rx.ctl $(OBJ_DIR)/tcp-rx.bin \
		-P tcp_rx_rtt_stage2_start=0x100000 \
		-P tcp_rx_read_rnmdr_stage2_start=0x200000 \
		-P tcp_rx_read_rnpdr_stage2_start=0x300000 \
		-P tcp_rx_read_serq_stage2_start=0x400000 \
		-P tcp_rx_sack_stage2_start=0x500000 \
		-P tcp_rx_read_rnmpr_stage2_start=0x600000 \
		> $(SRC_DIR)/run/tcp-rx.run
	$(CAPSIM_BIN) -c $(INC_DIRS) $(SRC_DIR)/ctl/tcp-write-serq.ctl $(OBJ_DIR)/tcp-write-serq.bin > $(SRC_DIR)/run/tcp-write-serq.run
	$(CAPSIM_BIN) -c $(INC_DIRS) $(SRC_DIR)/ctl/tcp-cc.ctl $(OBJ_DIR)/tcp-cc.bin \
	    	-P tcp_rx_fc_stage5_start=0x100000 > $(SRC_DIR)/run/tcp-cc.run
	$(CAPSIM_BIN) -c $(INC_DIRS) $(SRC_DIR)/ctl/tcp-rtt-first.ctl $(OBJ_DIR)/tcp-rtt.bin \
	    	-P tcp_rx_fra_stage3_start=0x100000 > $(SRC_DIR)/run/tcp-rtt-first.run
	$(CAPSIM_BIN) -c $(INC_DIRS) $(SRC_DIR)/ctl/tcp-rtt-non-first.ctl $(OBJ_DIR)/tcp-rtt.bin \
	  	-P tcp_rx_fra_stage3_start=0x100000 > $(SRC_DIR)/run/tcp-rtt-non-first.run
	$(CAPSIM_BIN) -c $(INC_DIRS) $(SRC_DIR)/ctl/tcp-fastretx-alert-cwr-end.ctl $(OBJ_DIR)/tcp-fastretx-alert.bin \
	    	-P tcp_rx_cc_stage4_start=0x100000 > $(SRC_DIR)/run/tcp-fastretx-alert-cwr-end.run

capsimrun:
	$(CAPSIM_BIN) -c $(INC_DIRS) $(SRC_DIR)/ctl/tcp-rtt-first.ctl $(OBJ_DIR)/tcp-rtt.bin \
	    	> $(SRC_DIR)/run/tcp-rtt-first.run
	$(CAPSIM_BIN) -c $(INC_DIRS) $(SRC_DIR)/ctl/tcp-rtt-non-first.ctl $(OBJ_DIR)/tcp-rtt.bin \
	    	> $(SRC_DIR)/run/tcp-rtt-non-first.run



asm_tests:
	(cd test; py.test -v)

myclean:
	@rm -fR $(OBJ_DIR)/*.bin
	@rm -fR $(GEN_DIR)/.tcp_rx_asm
