#! /bin/bash
export SWDIR=`git rev-parse --show-toplevel`
export NICDIR=$SWDIR/nic/
export CAPRI_MOCK_MODE=1
export CAPRI_MOCK_MEMORY_MODE=1
export COVFILE=${NICDIR}/coverage/sim_bullseye_hal.cov

rm -f ${NICDIR}/core.*

if [[ "$1" ==  --coveragerun ]]; then
    CMD_OPTS="COVFILE\=${COVFILE}"
fi

set -e
# PI gtests
export PATH=${PATH}:${NICDIR}/build/x86_64/apulu/bin
echo "Starting " `which apulu_ftl_test`
sleep 1
$ARGS apulu_ftl_test $*
valgrind --track-origins=yes --leak-check=full --show-leak-kinds=definite apulu_ftl_test --gtest_filter=basic*:collision* --log-file=ftl_test_valgrind.log
#perf record --call-graph fp ftl_test --gtest_filter="scale.insert1M"
#gdb --args ftl_test --gtest_filter="basic.insert"