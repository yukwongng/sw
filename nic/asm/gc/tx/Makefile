TOPDIR = ../../..
GEN_DIR = $(TOPDIR)/gen
CUSTOM = $(GEN_DIR)/.gc_tx_asm
CUSTOM_CLEAN = my_clean

include $(TOPDIR)/make/Makefile.inc
include $(TOPDIR)/make/Make.inc

PROJ = gc_txdma
OBJ_DIR = $(TOPDIR)/obj/p4plus_bin
GEN_DIR = $(TOPDIR)/gen
ASM_HDR_DIR = $(GEN_DIR)/$(PROJ)/asm_out/
ALT_ASM_HDR_DIR = $(GEN_DIR)/$(PROJ)/alt_asm_out/
P4_DIR = $(TOPDIR)/p4/gc
COMMON_P4_DIR = $(TOPDIR)/p4/common-p4+
INC_DIRS = -I. -I$(ALT_ASM_HDR_DIR) -I$(ASM_HDR_DIR) -I../include -I../../common-p4+/include -I../../../include -I../../proxy/include
SRC_DIR = .

CAPSIM_BIN = ../../../asic/capri/model/capsim-gen/bin/capsim

SRC_FILES := \
    $(SRC_DIR)/gc_tx.s \
    $(SRC_DIR)/gc_tx_read_descr_addr.s \
    $(SRC_DIR)/gc_tx_dummy.s \
    $(SRC_DIR)/gc_tx_inc_descr_free_pair_pi.s

BIN_FILES := $(patsubst %.s, %.bin, $(SRC_FILES))

FILES = $(basename $(notdir $(BIN_FILES)))
COMP_BIN_FILES = $(FILES:%=$(OBJ_DIR)/%.bin)

$(GEN_DIR)/.gc_tx_asm: $(OBJ_DIR)/.dir $(COMP_BIN_FILES)
	touch $@

DEPS_DIRS = $(subst -I,, $(INC_DIRS))
DEPS = $(shell find $(DEPS_DIRS) -name '*.h')
$(OBJ_DIR)/%.bin: %.s $(DEPS) $(CAPAS_BIN)
	$(CAPAS_BIN) $(CAPAS_COV) $(INC_DIRS) $< -o $@

TEST_FILES := \

TMP_FILES  := $(patsubst %.ctl, %.x, $(TEST_FILES))

gc_initial_action_ARGS :=

$(TMP_FILES): %.x: %.ctl
	@echo Testing ...
	@echo $<
	$(CAPSIM_BIN) -e ${$(basename ${<})_ARGS} $(INC_DIRS) $< $(OBJ_DIR)/$(patsubst %.x,%.bin,$@)

test: $(TMP_FILES)

my_clean:
	@rm -fR $(OBJ_DIR)/*.bin
