from "registry.test.pensando.io:5000/pensando/nic:1.42"

inside "/etc" do
  run "rm localtime"
  run "ln -s /usr/share/zoneinfo/US/Pacific localtime"
end

copy "asset-build/asset-push.sh", "/asset-push.sh"
run "chmod +x /asset-push.sh"
copy "asset-build/apulu-asset-push.sh", "/apulu-asset-push.sh"
run "chmod +x /apulu-asset-push.sh"

copy "asset-build/entrypoint.sh", "/entrypoint.sh"
run "chmod +x /entrypoint.sh"

entrypoint "/entrypoint.sh"
