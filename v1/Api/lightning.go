package api

import (
	"path/filepath"

	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func SetLightningApi(tc *modules.TaskController, rootFolder *string) {
	folder := filepath.Join(*rootFolder, "Lightning")

	btcSymbol := "BTC"
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/lightning/base_fee_median?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/lightning/network_capacity_sum?a=%s&i=1h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/lightning/channel_size_mean?a=%s&i=1h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/lightning/channel_size_median?a=%s&i=1h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/lightning/fee_rate_median?a=%s&i=24h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/lightning/gini_capacity_distribution?a=%s&i=24h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/lightning/gini_channel_distribution?a=%s&i=24h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/lightning/node_connectivity?a=%s&i=24h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/lightning/channels_count?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/lightning/nodes_count?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
}
