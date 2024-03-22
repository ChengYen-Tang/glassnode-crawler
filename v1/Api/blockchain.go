package api

import (
	"path/filepath"

	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func SetBlockchainApi(tc *modules.TaskController, rootFolder *string) {
	folder := filepath.Join(*rootFolder, "Blockchain")

	symbols := [...]string{"BTC", "ETH"}
	for _, symbol := range symbols {
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/block_height?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/block_interval_mean?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/block_interval_median?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/block_size_mean?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/block_size_sum?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/block_count?a=%s&i=1h&f=CSV", &folder, &symbol, false)
	}

	symbol := "BTC"
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/utxo_profit_relative?a=%s&i=1h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/utxo_created_value_mean?a=%s&i=1h&f=CSV", &folder, &symbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/utxo_created_value_median?a=%s&i=1h&f=CSV", &folder, &symbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/utxo_created_value_sum?a=%s&i=1h&f=CSV", &folder, &symbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/utxo_spent_value_mean?a=%s&i=1h&f=CSV", &folder, &symbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/utxo_spent_value_median?a=%s&i=1h&f=CSV", &folder, &symbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/utxo_spent_value_sum?a=%s&i=1h&f=CSV", &folder, &symbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/utxo_count?a=%s&i=1h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/utxo_created_count?a=%s&i=1h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/utxo_loss_count?a=%s&i=1h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/utxo_profit_count?a=%s&i=1h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/utxo_spent_count?a=%s&i=1h&f=CSV", &folder, &symbol, false)
}
