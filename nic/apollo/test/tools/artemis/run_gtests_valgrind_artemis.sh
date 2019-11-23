#! /bin/bash

export PIPELINE=artemis
export NICDIR=`pwd`

# initial setup
source ${NICDIR}/apollo/test/tools/setup_gtests.sh
setup

# unset mock mem mode for valgrind
unset CAPRI_MOCK_MEMORY_MODE

# run all gtests
run_valgrind_gtest CFG=" -i ${NICDIR}/apollo/test/scale/${PIPELINE}/scale_cfg_1vpc.json "

# end of script
clean_exit