#!/usr/bin/env bash

# Run Sanity for golden config on naples sim environment.

TOPDIR=$GOPATH/src/github.com/pensando/sw/nic/e2etests/go/agent/
HEIMDALL=$TOPDIR/cmd/heimdall
HEIMDALL_RUN=$HEIMDALL/main.go
HEIMDALL_GOLDEN_CONFIG_DIR=$TOPDIR/configs/golden
HEIMDALL_GOLDEN_CONFIG=$HEIMDALL_GOLDEN_CONFIG_DIR/golden_cfg.yml
SIM_DEVICE_JSON=$HEIMDALL_GOLDEN_CONFIG_DIR/sim_device.json
UPLINK_MAP_JSON=$HEIMDALL_GOLDEN_CONFIG_DIR/uplink_map.json


echo "Starting Heimdall run..."

set +x
echo "Pushing golden config..."
go run $HEIMDALL_RUN run --config-file $HEIMDALL_GOLDEN_CONFIG --enable-sim

OUT=$?
if [ $OUT -ne 0 ];then
   echo "Pushing configuration to agent failed!"
   exit $OUT
fi

echo "Starting Traffic Tests..."



echo "Starting Network-Network Traffic tests.."
go run $HEIMDALL_RUN traffic --uplink-map $UPLINK_MAP_JSON --config-file $HEIMDALL_GOLDEN_CONFIG --sim-mode

OUT=$?
if [ $OUT -ne 0 ];then
   echo "Network-Network traffic test failed"
   exit $OUT
fi


echo "Starting Host-Host Traffic tests.."

echo "Pushing golden config..."
go run $HEIMDALL_RUN run --device-file $SIM_DEVICE_JSON --config-file $HEIMDALL_GOLDEN_CONFIG --enable-sim

OUT=$?
if [ $OUT -ne 0 ];then
   echo "Pushing configuration to agent failed!"
   exit $OUT
fi


go run $HEIMDALL_RUN traffic --device-file $SIM_DEVICE_JSON --config-file $HEIMDALL_GOLDEN_CONFIG --sim-mode

OUT=$?
if [ $OUT -ne 0 ];then
   echo "Host-Host traffic test failed"
   exit $OUT
fi

go run $HEIMDALL_RUN traffic --device-file $SIM_DEVICE_JSON  --uplink-map $UPLINK_MAP_JSON --config-file $HEIMDALL_GOLDEN_CONFIG --sim-mode

OUT=$?
if [ $OUT -ne 0 ];then
   echo "Host-Network traffic test failed"
   exit $OUT
fi


echo "Golden sanity passed..."
