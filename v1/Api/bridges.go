package api

import (
	"path/filepath"

	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func SetBridgesApi(tc *modules.TaskController, rootFolder *string) {
	folder := filepath.Join(*rootFolder, "Bridges")

	tc.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/bridges/deposits_by_chain?a=ETH&f=CSV", &folder, false)
	tc.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/bridges/net_volume_by_chain?a=ETH&f=CSV", &folder, false)
	tc.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/bridges/total_value_locked_by_chain?a=ETH&f=CSV", &folder, false)
	tc.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/bridges/total_value_locked_by_chain_relative?a=ETH&f=CSV", &folder, false)
	tc.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/bridges/withdrawals_by_chain?a=ETH&f=CSV", &folder, false)
}
