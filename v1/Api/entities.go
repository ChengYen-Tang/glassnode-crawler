package api

import (
	"path/filepath"

	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func SetEntitiesApi(tc *modules.TaskController, rootFolder *string) {
	folder := filepath.Join(*rootFolder, "Entities")
	symbol := "BTC"

	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/entities/active_count?a=%s&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/entities/net_growth_count?a=%s&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/entities/supply_distribution_relative?a=%s&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/entities/new_count?a=%s&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/entities/min_1k_count?a=%s&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/entities/profit_relative?a=%s&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/entities/receiving_count?a=%s&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/entities/sending_count?a=%s&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/entities/supply_balance_less_0001?a=%s&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/entities/supply_balance_more_100k?a=%s&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/entities/supply_balance_0001_001?a=%s&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/entities/supply_balance_001_01?a=%s&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/entities/supply_balance_01_1?a=%s&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/entities/supply_balance_1_10?a=%s&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/entities/supply_balance_10_100?a=%s&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/entities/supply_balance_100_1k?a=%s&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/entities/supply_balance_10k_100k?a=%s&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/entities/supply_balance_1k_10k?a=%s&f=CSV", &folder, &symbol, false)
}
