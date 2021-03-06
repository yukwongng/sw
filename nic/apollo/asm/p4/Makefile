TOPDIR = ../../..
SWDIR = $(TOPDIR)/../
GEN_DIR = $(TOPDIR)/gen
CUSTOM = $(GEN_DIR)/.apollo_asm
CUSTOM_CLEAN = asm_clean

include $(TOPDIR)/make/Makefile.inc
include $(TOPDIR)/make/Make.inc

PROJ = apollo
OBJ_DIR = $(TOPDIR)/obj/$(PROJ)_p4/asm_bin/
ASM_HDR_DIR = $(TOPDIR)/gen/$(PROJ)/asm_out/
INC_DIRS = -I. -I$(ASM_HDR_DIR) -I$(SWDIR)
SRC_DIR = .

CAPSIM_BIN = ../../../asic/capri/model/capsim-gen/bin/capsim

SRC_FILES := $(wildcard $(SRC_DIR)/*.asm)

BIN_FILES := $(patsubst %.asm, %.bin, $(SRC_FILES))

DEPS_DIRS = $(subst -I,, $(INC_DIRS))
DEPS = $(shell find $(DEPS_DIRS) -name '*.h')

FILES = $(basename $(notdir $(BIN_FILES)))
COMP_BIN_FILES = $(FILES:%=$(OBJ_DIR)/%.bin)

$(GEN_DIR)/.apollo_asm: $(OBJ_DIR)/.dir $(COMP_BIN_FILES)
	touch $@

$(OBJ_DIR)/%.bin: %.asm $(DEPS) $(CAPAS_BIN)
	@echo "CAPAS   " $<
	@$(CAPAS_BIN) $(CAPAS_COV) $(INC_DIRS) $< -o $@

TEST_FILES := \
	$(SRC_DIR)/local_vnic_by_vlan_tx.ctl \
	$(SRC_DIR)/flow.ctl \
	$(SRC_DIR)/flow_info.ctl \
	$(SRC_DIR)/vnic_tx_stats.ctl \
	$(SRC_DIR)/tep_rx.ctl \
	$(SRC_DIR)/tep_tx.ctl \
	$(NULL)
TMP_FILES  := $(patsubst %.ctl, %.x, $(TEST_FILES))

flow_ARGS := -r flow_hash_info
flow_info_ARGS := -r flow_info

$(TMP_FILES): %.x: %.ctl
	@echo Testing ...
	@echo $<
	$(CAPSIM_BIN) -e ${$(basename ${<})_ARGS} $(INC_DIRS) $< $(OBJ_DIR)/$(patsubst %.x,%.bin,$@)

gtest: $(TMP_FILES)

asm_clean:
	@rm -fR $(OBJ_DIR)/*.bin
	@rm -fR $(GEN_DIR)/.apollo_asm

.asm_test:
	@echo DEPS_DIRS = $(DEPS_DIRS)
	@echo DEPS = $(DEPS)
	@echo FILES = $(FILES)
	@echo BIN_FILES = $(BIN_FILES)
	@echo COMP_BIN_FILES = $(COMP_BIN_FILES)
