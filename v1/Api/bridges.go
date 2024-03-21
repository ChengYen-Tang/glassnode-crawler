package api

import (
	"container/list"
	"path/filepath"

	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func SetBridgesApi(apiInfos *list.List, rootFolder *string) {
	folder := filepath.Join(*rootFolder, "Bridges")

	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/bridges/deposits_by_chain?a=ETH&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/bridges/net_volume_by_chain?a=ETH&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/bridges/total_value_locked_by_chain?a=ETH&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/bridges/total_value_locked_by_chain_relative?a=ETH&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/bridges/withdrawals_by_chain?a=ETH&f=CSV", &folder))
}
