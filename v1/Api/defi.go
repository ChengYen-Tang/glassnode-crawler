package api

import (
	"container/list"
	"path/filepath"

	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func SetDefiApi(apiInfos *list.List, rootFolder *string) {
	folder := filepath.Join(*rootFolder, "Defi")

	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/defi/total_value_locked?a=ETH&f=CSV", &folder))
}
