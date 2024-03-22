package api

import (
	"path/filepath"

	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func SetDefiApi(tc *modules.TaskController, rootFolder *string) {
	folder := filepath.Join(*rootFolder, "Defi")

	tc.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/defi/total_value_locked?a=ETH&f=CSV", &folder, false)
}
