package api

import (
	"path/filepath"

	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func SetDefiApi(tc *modules.TaskController, rootFolder *string) {
	folder := filepath.Join(*rootFolder, "Defi")
	symbol := "ETH"

	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/defi/total_value_locked?a=%s&f=CSV", &folder, &symbol, false)
}
