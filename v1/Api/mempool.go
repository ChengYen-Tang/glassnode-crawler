package api

import (
	"path/filepath"

	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func SetMempoolApi(tc *modules.TaskController, rootFolder *string) {
	folder := filepath.Join(*rootFolder, "Mempool")

	btcSymbol := "BTC"

	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/mempool/fees_average_relative?a=%s&i=1h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/mempool/fees_median_relative?a=%s&i=1h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/mempool/txs_value_sum?a=%s&i=1h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/mempool/txs_value_distribution?a=%s&i=1h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/mempool/fees_sum?a=%s&i=1h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/mempool/fees_distribution?a=%s&i=1h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/mempool/txs_count_sum?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/mempool/txs_count_distribution?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/mempool/txs_size_distribution?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
}
