package api

import (
	"path/filepath"

	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func SetFeesApi(tc *modules.TaskController, rootFolder *string) {
	folder := filepath.Join(*rootFolder, "Fees")

	symbols := [...]string{"BTC", "ETH"}
	for _, symbol := range symbols {
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/exchanges_relative?a=%s&1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/exchanges_mean?a=%s&1h&f=CSV", &folder, &symbol, true)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/exchanges_sum?a=%s&1h&f=CSV", &folder, &symbol, true)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/volume_mean?a=%s&1h&f=CSV", &folder, &symbol, true)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/volume_median?a=%s&1h&f=CSV", &folder, &symbol, true)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/volume_sum?a=%s&1h&f=CSV", &folder, &symbol, true)
	}

	symbol := "ETH"
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/gas_price_mean?a=%s&1h&f=CSV", &folder, &symbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/gas_price_median?a=%s&1h&f=CSV", &folder, &symbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/gas_used_sum_bridges?a=%s&24h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/gas_used_sum_bridges_relative?a=%s&24h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/gas_used_sum_defi?a=%s&24h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/gas_used_sum_defi_relative?a=%s&24h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/gas_used_sum_erc20?a=%s&24h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/gas_used_sum_erc20_relative?a=%s&24h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/gas_used_sum_nfts?a=%s&24h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/gas_used_sum_nfts_relative?a=%s&24h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/gas_used_sum_stablecoins?a=%s&24h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/gas_used_sum_stablecoins_relative?a=%s&24h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/tx_types_breakdown_sum?a=%s&24h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/tx_types_breakdown_relative?a=%s&24h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/gas_used_sum_vanilla?a=%s&24h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/gas_used_sum_vanilla_relative?a=%s&24h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/gas_used_mean?a=%s&1h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/gas_used_median?a=%s&1h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/gas_used_sum?a=%s&1h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/gas_limit_tx_mean?a=%s&1h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/fees/gas_limit_tx_median?a=%s&1h&f=CSV", &folder, &symbol, false)
}
