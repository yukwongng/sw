#!/bin/bash -e
#cd `dirname $0`        
#NIC_DIR=$PWD      
TOOLS_DIR=`dirname $0`
ABS_TOOLS_DIR=`readlink -f $TOOLS_DIR`
NIC_DIR=`dirname $ABS_TOOLS_DIR`
# export LD_LIBRARY_PATH=$NIC_DIR/model_sim/:$NIC_DIR/model_sim/libs:/home/asic/tools/src/0.25/x86_64/lib64:$LD_LIBRARY_PATH        
SUDO=

while getopts t:p: option
do
 case "${option}"
 in
 t) server_type=${OPTARG};;
 p) portlist=${OPTARG};;
 esac
done 

if [ "$server_type" = "tls" ]; then
    $SUDO $GDB $NIC_DIR/../bazel-bin/nic/e2etests/proxy/nic_proxy-e2etest_tls-server $portlist
else
    $SUDO $GDB $NIC_DIR/../bazel-bin/nic/e2etests/proxy/nic_proxy-e2etest_tcp-server $portlist
fi

#$SUDO $GDB $NIC_DIR/../bazel-bin/nic/e2etests/proxy/nic_proxy-e2etest_tcp-server 80  2>&1 | tee $NIC_DIR/server.log
#$SUDO $GDB $NIC_DIR/../bazel-bin/nic/e2etests/proxy/nic_proxy-e2etest_tls-server 80


##! /bin/bash -e
#TOOLS_DIR=`dirname $0`
#ABS_TOOLS_DIR=`readlink -f $TOOLS_DIR`
#NIC_DIR=`dirname $ABS_TOOLS_DIR`
#export LD_LIBRARY_PATH=$NIC_DIR/model_sim/:$NIC_DIR/model_sim/libs:/home/asic/tools/src/0.25/x86_64/lib64:$LD_LIBRARY_PATH        
#$GDB $NIC_DIR/model_sim/build/cap_model +plog=info +model_debug=$NIC_DIR/gen/iris/dbg_out/model_debug.json 2>&1 | tee $NIC_DIR/model.log
