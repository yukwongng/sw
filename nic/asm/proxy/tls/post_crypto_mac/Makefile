TOPDIR = ../../../..
GEN_DIR = $(TOPDIR)/gen
CUSTOM = $(GEN_DIR)/.tls_post_ce_asm
CUSTOM_CLEAN = clean

include $(TOPDIR)/make/Makefile.inc
include $(TOPDIR)/make/Make.inc

PROJ = tls_txdma_post_crypto_mac
OBJ_DIR = $(TOPDIR)/obj/p4plus_bin
GEN_DIR = $(TOPDIR)/gen
ASM_HDR_DIR = $(GEN_DIR)/$(PROJ)/asm_out/
P4_DIR = $(TOPDIR)/p4/proxy
COMMON_P4_DIR = $(TOPDIR)/p4/common-p4+
INC_DIRS = -I. -I$(ASM_HDR_DIR) -I./include -I../include -I../../../common-p4+/include -I../../../cpu-p4plus/include -I../../include -I../../tcp/include -I$(TOPDIR)/include
SRC_DIR = .

CAPSIM_BIN = $(TOPDIR)/asic/capri/model/capsim-gen/bin/capsim

SRC_FILES := \
    $(SRC_DIR)/tls-mac-bsq.s \
    $(SRC_DIR)/tls-mac-queue-enc-brq.s \
    $(SRC_DIR)/tls-mac-read-idesc.s \
    $(SRC_DIR)/tls-mac-read-aad.s \
    $(SRC_DIR)/tls-mac-rx-bsq-mac.s \
    $(SRC_DIR)/tls-mac-bsq-consume.s \
    $(SRC_DIR)/tls-mac-read-tnmdr-alloc-idx.s \
    $(SRC_DIR)/tls-mac-read-tnmpr-alloc-idx.s \
    $(SRC_DIR)/tls-mac-tdesc-alloc.s \
    $(SRC_DIR)/tls-mac-tpage-alloc.s \
    $(SRC_DIR)/tls-mac-read-random-iv.s \
    $(SRC_DIR)/tls-mac-post-crypto-stats.s \
    $(SRC_DIR)/tls-mac-bld-enc-barco-req.s

BIN_FILES := $(patsubst %.s, %.bin, $(SRC_FILES))

FILES = $(basename $(notdir $(BIN_FILES)))
COMP_BIN_FILES = $(FILES:%=$(OBJ_DIR)/%.bin)

$(GEN_DIR)/.tls_post_ce_asm: $(OBJ_DIR)/.dir $(COMP_BIN_FILES)
	touch $@

DEPS_DIRS = $(subst -I,, $(INC_DIRS))
DEPS = $(shell find $(DEPS_DIRS) -name '*.h')
$(OBJ_DIR)/%.bin: %.s $(DEPS) $(CAPAS_BIN)
	$(CAPAS_BIN) $(CAPAS_COV) $(INC_DIRS) $< -o $@

TEST_FILES := \
    $(SRC_DIR)/ctl/tls-alloc-brq-pi.ctl \
    $(SRC_DIR)/ctl/tls-alloc-rnmdr.ctl \
    $(SRC_DIR)/ctl/tls-alloc-rnmdr-split.ctl \
    $(SRC_DIR)/ctl/tls-alloc-rnmpr.ctl \
    $(SRC_DIR)/ctl/tls-queue-brq.ctl \
    $(SRC_DIR)/ctl/tls-read-desc.ctl \
    $(SRC_DIR)/ctl/tls-read-header.ctl \
    $(SRC_DIR)/ctl/tls-rx-serq.ctl \
    $(SRC_DIR)/ctl/tls-serq.ctl \
    $(SRC_DIR)/ctl/tls-update-page-ctl.ctl

TMP_FILES  := $(patsubst %.ctl, %.x, $(TEST_FILES))


$(TMP_FILES): %.x: %.ctl
	@echo Testing ...
	@echo $<
	$(CAPSIM_BIN) -e ${$(basename ${<})_ARGS} $(INC_DIRS) $< $(OBJ_DIR)/$(patsubst %.x,%.bin,$@)

test: $(TMP_FILES)

run: $(BIN_FILES)
	$(CAPSIM_BIN) -c $(INC_DIRS) $(SRC_DIR)/ctl/tls-serq.ctl $(OBJ_DIR)/tls-serq.bin > $(SRC_DIR)/run/tls-serq.run
	$(CAPSIM_BIN) -c $(INC_DIRS) $(SRC_DIR)/ctl/tls-read-desc.ctl $(OBJ_DIR)/tls-read-desc.bin > $(SRC_DIR)/run/tls-read-desc.run
	$(CAPSIM_BIN) -c $(INC_DIRS) $(SRC_DIR)/ctl/tls-rx-serq.ctl $(OBJ_DIR)/tls-rx-serq.bin > $(SRC_DIR)/run/tls-rx-serq.run
	$(CAPSIM_BIN) -c $(INC_DIRS) $(SRC_DIR)/ctl/tls-read-header.ctl $(OBJ_DIR)/tls-read-header.bin > $(SRC_DIR)/run/tls-read-header.run
	$(CAPSIM_BIN) -c $(INC_DIRS) $(SRC_DIR)/ctl/tls-alloc-rnmdr-split.ctl $(OBJ_DIR)/tls-alloc-rnmdr-split.bin > $(SRC_DIR)/run/tls-alloc-rnmdr-split.run
	$(CAPSIM_BIN) -c $(INC_DIRS) $(SRC_DIR)/ctl/tls-update-page-ctl.ctl $(OBJ_DIR)/tls-update-page-ctl.bin > $(SRC_DIR)/run/tls-update-page-ctl.run
	$(CAPSIM_BIN) -c $(INC_DIRS) $(SRC_DIR)/ctl/tls-alloc-rnmdr.ctl $(OBJ_DIR)/tls-alloc-rnmdr.bin > $(SRC_DIR)/run/tls-alloc-rnmdr.run
	$(CAPSIM_BIN) -c $(INC_DIRS) $(SRC_DIR)/ctl/tls-alloc-rnmpr.ctl $(OBJ_DIR)/tls-alloc-rnmpr.bin > $(SRC_DIR)/run/tls-alloc-rnmpr.run
	$(CAPSIM_BIN) -c $(INC_DIRS) $(SRC_DIR)/ctl/tls-alloc-brq-pi.ctl $(OBJ_DIR)/tls-alloc-brq-pi.bin > $(SRC_DIR)/run/tls-alloc-brq-pi.run
	$(CAPSIM_BIN) -c $(INC_DIRS) $(SRC_DIR)/ctl/tls-queue-brq.ctl $(OBJ_DIR)/tls-queue-brq.bin > $(SRC_DIR)/run/tls-queue-brq.run

clean:
	@rm -fR $(OBJ_DIR)/*.bin
	@rm -fR $(GEN_DIR)/.tls_post_ce_asm

