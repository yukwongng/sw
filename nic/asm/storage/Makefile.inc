# This Makefile.inc is to be included within the subdirectories in storage

TOPDIR = ../../..
CUSTOM = all
CUSTOM_CLEAN = clean

include $(TOPDIR)/make/Makefile.inc
include $(TOPDIR)/make/Make.inc

PROJ = $(shell basename $(CURDIR))
OBJ_DIR = $(TOPDIR)/obj/p4plus_bin
GEN_DIR = $(TOPDIR)/gen
ASM_HDR_DIR = $(GEN_DIR)/$(PROJ)/asm_out/
P4_DIR = $(TOPDIR)/p4/storage
COMMON_P4_DIR = $(TOPDIR)/p4/common-p4+
INC_DIRS = -I. -I$(ASM_HDR_DIR) -I../include -I../../../common-p4+/include -I$(P4_DIR)/common
SRC_DIR = .
CAPSIM_BIN = $(TOPDIR)/asic/capri/model/capsim-gen/bin/capsim

all: deps asm

deps: $(ASM_HDR_DIR)/ingress.h

SRC_FILES := \
    $(wildcard *.s)

BIN_FILES := $(patsubst %.s, %.bin, $(SRC_FILES))

$(BIN_FILES): %.bin: %.s
	@#echo $@
	@mkdir -p $(OBJ_DIR)
	$(CAPAS_BIN) $(CAPAS_COV) $(INC_DIRS) $< -o $(OBJ_DIR)/$(shell basename $@)

%.h: $(P4_DIR)/$(PROJ).p4 $(COMMON_P4_DIR)/common_headers.p4
	$(MAKE) -C $(P4_DIR) ../../gen/$(PROJ)/asm_out/ingress.h

asm: $(BIN_FILES)

clean:
	@rm -fR $(OBJ_DIR)/*.bin
	@rm -fR $(ASM_HDR_DIR)/*.h
