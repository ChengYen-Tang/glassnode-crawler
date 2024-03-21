package api

import (
	"container/list"
	"path/filepath"

	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func SetBlockchainApi(apiInfos *list.List, rootFolder *string) {
	folder := filepath.Join(*rootFolder, "Blockchain")

	symbols := [...]string{"BTC", "ETH"}
	for _, symbol := range symbols {
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/block_height?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/block_interval_mean?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/block_interval_median?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/block_size_mean?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/block_size_sum?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/block_count?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/utxo_profit_relative?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/utxo_created_value_mean?a=%s&i=1h&c=USD&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/utxo_created_value_median?a=%s&i=1h&c=USD&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/utxo_created_value_sum?a=%s&i=1h&c=USD&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/utxo_spent_value_mean?a=%s&i=1h&c=USD&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/utxo_spent_value_median?a=%s&i=1h&c=USD&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/utxo_spent_value_sum?a=%s&i=1h&c=USD&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/utxo_count?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/utxo_created_count?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/utxo_loss_count?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/utxo_profit_count?a=%s&i=1h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/blockchain/utxo_spent_count?a=%s&i=1h&f=CSV", &folder, &symbol))
	}
}
