TOPDIR = ../..
SWDIR = $(TOPDIR)/../
GEN_DIR = $(TOPDIR)/gen
CUSTOM = $(GEN_DIR)/.mpu_timestamp_asm
CUSTOM_CLEAN = asm_clean

include $(TOPDIR)/make/Makefile.inc
include $(TOPDIR)/make/Make.inc

PROJ = mpu_timestamp
OBJ_DIR = ../../obj/p4plus_bin
ASM_HDR_DIR = ../../gen/
INC_DIRS = -I. -I$(ASM_HDR_DIR) -I$(TOPDIR)/include -I$(SWDIR) -I$(TOPDIR)/nic/p4/ftl_dev/include -I../../build/iris/gen/datapath/
SRC_DIR = .


CAPSIM_BIN = ../../asic/capri/model/capsim-gen/bin/capsim

SRC_FILES := \
    $(SRC_DIR)/mpu_timestamp_start.s \
        $(NULL)

BIN_FILES := $(patsubst %.s, %.bin, $(SRC_FILES))

FILES = $(basename $(notdir $(BIN_FILES)))
COMP_BIN_FILES = $(FILES:%=$(OBJ_DIR)/%.bin)

$(GEN_DIR)/.mpu_timestamp_asm: $(OBJ_DIR)/.dir $(COMP_BIN_FILES)
        touch $@

DEPS_DIRS = $(subst -I,, $(INC_DIRS))
DEPS = $(shell find $(DEPS_DIRS) -name '*.h')
$(OBJ_DIR)/%.bin: %.s $(DEPS) $(CAPAS_BIN)
        $(CAPAS_BIN) $(CAPAS_COV) $(INC_DIRS) $< -o $@


#BIN_FILES := $(patsubst %.s, %.bin, $(SRC_FILES))
#
#$(BIN_FILES): %.bin: %.s
#       @echo "CAPAS $@"
#       @mkdir -p $(OBJ_DIR)
#       @$(CAPAS_BIN) $(CAPAS_COV) $(INC_DIRS) $< -o $(OBJ_DIR)/$(shell basename $@)
#
#asm: $(BIN_FILES)

TEST_FILES := $(patsubst %.s, %.ctl, $(SRC_FILES))
TMP_FILES  := $(patsubst %.ctl, %.x, $(TEST_FILES))

$(TMP_FILES): %.x: %.ctl
        @echo Testing ...
        @echo $<
        $(CAPSIM_BIN) -e ${$(basename ${<})_ARGS} $(INC_DIRS) $< $(OBJ_DIR)/$(patsubst %.x,%.bin,$@)

gtest:

asm_clean:
        @rm -f $(OBJ_DIR)/mpu_timestamp*.bin
        @rm -fR $(GEN_DIR)/.mpu_timestamp_asm
