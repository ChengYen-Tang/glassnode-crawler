package api

import (
	"container/list"
	"path/filepath"

	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func SetAddressApi(apiInfos *list.List, rootFolder *string) {
	folder := filepath.Join(*rootFolder, "Address")

	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/addresses/accumulation_count?a=BTC&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/addresses/accumulation_balance?a=BTC&c=USD&f=CSV", &folder))
	symbols := [...]string{"BTC", "ETH"}
	for _, symbol := range symbols {
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/active_count?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/supply_distribution_relative?a=%s&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/loss_count?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/profit_count?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/min_1_usd_count?a=%s&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/min_10_usd_count?a=%s&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/min_100_usd_count?a=%s&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/min_100k_usd_count?a=%s&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/min_10k_usd_count?a=%s&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/min_1k_usd_count?a=%s&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/min_1m_usd_count?a=%s&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/min_point_zero_1_count?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/min_point_1_count?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/min_1_count?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/min_10_count?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/min_100_count?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/min_10k_count?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/min_1k_count?a=%s&i=1h&f=CSV", &folder, &symbol))
	}
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/addresses/min_32_count?a=ETH&i=1h&f=CSV", &folder))
	for _, symbol := range symbols {
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/non_zero_count?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/sending_to_exchanges_count?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/new_non_zero_count?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/profit_relative?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/receiving_count?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/sending_count?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/supply_balance_less_0001?a=%s&c=USD&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/supply_balance_more_100k?a=%s&c=USD&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/supply_balance_0001_001?a=%s&c=USD&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/supply_balance_001_01?a=%s&c=USD&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/supply_balance_01_1?a=%s&c=USD&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/supply_balance_1_10?a=%s&c=USD&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/supply_balance_10_100?a=%s&c=USD&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/supply_balance_100_1k?a=%s&c=USD&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/supply_balance_10k_100k?a=%s&c=USD&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/supply_balance_1k_10k?a=%s&c=USD&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/count?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/addresses/receiving_from_exchanges_count?a=%s&i=1h&f=CSV", &folder, &symbol))
	}
}
