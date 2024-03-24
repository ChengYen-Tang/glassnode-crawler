package api

import (
	"path/filepath"

	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func SetMiningApi(tc *modules.TaskController, rootFolder *string) {
	folder := filepath.Join(*rootFolder, "Mining")

	btcSymbol := "BTC"
	ethSymbol := "ETH"
	symbols := []string{btcSymbol, ethSymbol}
	btcMiners := [...]string{"aggregated", "other", "1THash&58COIN", "AntPool", "ArkPool", "BinancePool", "BitFury", "BitMinter", "Bixin", "BTC.com", "BTC.TOP", "DPool", "F2Pool", "FoundryUSAPool", "Genesis", "HuobiPool", "KuCoinPool", "Lubian.com", "LuxorTech", "MaraPool", "NovaBlock", "OKExPool", "Patoshi", "PegaPool", "Poolin", "SBICrypto", "SigmaPool", "SlushPool", "SpiderPool", "TerraPool", "UKRPool", "Ultimus", "ViaBTC"}
	ethMiners := [...]string{"aggregated"}

	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/mining/difficulty_latest?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/mining/hash_rate_mean?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/mining/marketcap_thermocap_ratio?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
	for _, miner := range btcMiners {
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/mining/miners_outflow_multiple?a=%s&miner=%s&i=1h&f=CSV", &folder, &btcSymbol, &miner, false)
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/mining/volume_mined_sum?a=%s&miner=%s&i=1h&f=CSV", &folder, &btcSymbol, &miner, true)
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/mining/revenue_sum?a=%s&miner=%s&i=1h&f=CSV", &folder, &btcSymbol, &miner, true)
	}
	for _, miner := range ethMiners {
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/mining/revenue_sum?a=%s&miner=%s&i=1h&f=CSV", &folder, &ethSymbol, &miner, true)
	}
	for _, symbol := range symbols {
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/mining/revenue_from_fees?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/mining/thermocap?a=%s&i=1h&f=CSV", &folder, &symbol, false)
	}
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/mining/miners_unspent_supply?a=%s&i=1h&f=CSV", &folder, &btcSymbol, true)
}
