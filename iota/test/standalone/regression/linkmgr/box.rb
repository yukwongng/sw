from "registry.test.pensando.io:5000/pensando/nic:1.37"

workdir "/sw"

copy "iota/entrypoint.sh", "/entrypoint.sh"
run "chmod +x /entrypoint.sh"

entrypoint "/entrypoint.sh"