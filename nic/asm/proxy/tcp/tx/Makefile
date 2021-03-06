TOPDIR = ../../../..
GEN_DIR = $(TOPDIR)/gen
CUSTOM = $(GEN_DIR)/.tcp_tx_asm
CUSTOM_CLEAN = my_clean

include $(TOPDIR)/make/Makefile.inc
include $(TOPDIR)/make/Make.inc

PROJ = tcp_proxy_txdma
OBJ_DIR = $(TOPDIR)/obj/p4plus_bin
GEN_DIR = $(TOPDIR)/gen
ASM_HDR_DIR = $(GEN_DIR)/$(PROJ)/asm_out/
ALT_ASM_HDR_DIR = $(GEN_DIR)/$(PROJ)/alt_asm_out/
P4_DIR = $(TOPDIR)/p4/proxy
COMMON_P4_DIR = $(TOPDIR)/p4/common-p4+
INC_DIRS = -I. -I$(ALT_ASM_HDR_DIR) -I$(ASM_HDR_DIR) -I../include -I../../../common-p4+/include \
           -I../../../cpu-p4plus/include -I$(TOPDIR)/include \
		   -I../../include -I$(TOPDIR)/p4/common/
SRC_DIR = .

CAPSIM_BIN = ../../../../asic/capri/model/capsim-gen/bin/capsim

SRC_FILES := \
    $(SRC_DIR)/tcp-read-rx2tx-shared.s \
    $(SRC_DIR)/tcp-read-rx2tx-shared-extra.s \
    $(SRC_DIR)/tcp-read-sesq-ci.s \
    $(SRC_DIR)/tcp-read-sesq-retx-ci.s \
    $(SRC_DIR)/tcp-s2-bubble.s \
    $(SRC_DIR)/tcp-read-descr.s \
    $(SRC_DIR)/tcp-read-tcp-flags.s \
    $(SRC_DIR)/tcp-retx.s \
    $(SRC_DIR)/tcp-clean-retx.s \
    $(SRC_DIR)/tcp-xmit.s \
    $(SRC_DIR)/tcp-xmit-ack.s \
    $(SRC_DIR)/tcp-tx-read-gc-nmdr-idx.s \
    $(SRC_DIR)/tcp-tso.s \
    $(SRC_DIR)/tcp-tx-stats.s

BIN_FILES := $(patsubst %.s, %.bin, $(SRC_FILES))

FILES = $(basename $(notdir $(BIN_FILES)))
COMP_BIN_FILES = $(FILES:%=$(OBJ_DIR)/%.bin)

$(GEN_DIR)/.tcp_tx_asm: $(OBJ_DIR)/.dir $(COMP_BIN_FILES)
	touch $@

DEPS_DIRS = $(subst -I,, $(INC_DIRS))
DEPS = $(shell find $(DEPS_DIRS) -name '*.h')
$(OBJ_DIR)/%.bin: %.s $(DEPS) $(CAPAS_BIN)
	$(CAPAS_BIN) $(CAPAS_COV) $(INC_DIRS) $(AS_DEFINES) $< -o $@

TEST_FILES := \
    $(SRC_DIR)/ctl/tcp-fast-timer.ctl \
    $(SRC_DIR)/ctl/tcp-slow-timer-keepalive.ctl \
    $(SRC_DIR)/ctl/tcp-slow-timer-pmtu.ctl \
    $(SRC_DIR)/ctl/tcp-tso.ctl \
    $(SRC_DIR)/ctl/tcp-tx.ctl

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

run: $(BIN_FILES)
	$(CAPSIM_BIN) -c $(INC_DIRS) $(SRC_DIR)/ctl/tcp-fast-timer.ctl $(OBJ_DIR)/tcp-fast-timer.bin > $(SRC_DIR)/run/tcp-fast-timer.run
	$(CAPSIM_BIN) -c $(INC_DIRS) $(SRC_DIR)/ctl/tcp-slow-timer-keepalive.ctl $(OBJ_DIR)/tcp-slow-timer.bin > $(SRC_DIR)/run/tcp-slow-timer-keepalive.run
	$(CAPSIM_BIN) -c $(INC_DIRS) $(SRC_DIR)/ctl/tcp-slow-timer-pmtu.ctl $(OBJ_DIR)/tcp-slow-timer.bin > $(SRC_DIR)/run/tcp-slow-timer-pmtu.run
	$(CAPSIM_BIN) -c $(INC_DIRS) $(SRC_DIR)/ctl/tcp-tso.ctl $(OBJ_DIR)/tcp-tso.bin > $(SRC_DIR)/run/tcp-tso.run
	$(CAPSIM_BIN) -c $(INC_DIRS) $(SRC_DIR)/ctl/tcp-tx.ctl $(OBJ_DIR)/tcp-tx.bin > $(SRC_DIR)/run/tcp-tx.run

my_clean:
	@rm -fR $(OBJ_DIR)/*.bin
	@rm -fR $(GEN_DIR)/.tcp_tx_asm
