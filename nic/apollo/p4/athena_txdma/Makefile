TOPDIR = ../../..
SWDIR = $(TOPDIR)/../
GEN_DIR = $(TOPDIR)/gen
CUSTOM = $(GEN_DIR)/.apollo_txdma
CUSTOM_CLEAN = apollo_txdma_clean

include $(TOPDIR)/make/Makefile.inc

NCC = $(TOPDIR)/tools/ncc/capri-ncc.py
PROJ = apollo_txdma
OBJ_DIR = $(TOPDIR)/obj/$(PROJ)
NCC_OPTIONS = --pipeline=apollo --asm-out --p4-plus --no-ohi --two-byte-profile --pd-gen --gen-dir $(GEN_DIR) --cfg-dir $(OBJ_DIR)/pgm_bin
NCC_OPTIONS += --fe-flags="-I$(SWDIR)"

p4_src = $(shell find $(SOURCEDIR) -name '*.p4')
hdr_src = $(shell find $(SOURCEDIR) -name '*.h')
inc_src = $(shell find ../include -name '*')
ncc_deps = $(shell find $(TOPDIR)/tools/ncc -path $(TOPDIR)/tools/ncc/utest -prune -o -name '*.c' -print -o -name '*.h' -print -o -name '*.py' -print)
$(GEN_DIR)/.apollo_txdma: $(p4_src) $(hdr_src) $(inc_src) $(ncc_deps)
	$(NCC) $(NCC_OPTIONS) --p4-plus-module txdma $(PROJ).p4
	touch $@

apollo_txdma_clean:
	rm -rf $(GEN_DIR)/$(PROJ)
	rm -rf $(OBJ_DIR)/pgm_bin/$(PROJ)*.bin
	rm -rf $(GEN_DIR)/.apollo_txdma

default:
	all
