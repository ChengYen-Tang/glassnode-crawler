package api

import (
	"path/filepath"

	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func SetBridgesApi(tc *modules.TaskController, rootFolder *string) {
	folder := filepath.Join(*rootFolder, "Bridges")
	symbol := "ETH"

	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/bridges/deposits_by_chain?a=%s&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/bridges/net_volume_by_chain?a=%s&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/bridges/total_value_locked_by_chain?a=%s&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/bridges/total_value_locked_by_chain_relative?a=%s&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/bridges/withdrawals_by_chain?a=%s&f=CSV", &folder, &symbol, false)
}
