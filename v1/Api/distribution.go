package api

import (
	"container/list"
	"path/filepath"
	"strings"

	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func SetDistributionApi(apiInfos *list.List, rootFolder *string) {
	folder := filepath.Join(*rootFolder, "Distribution")

	symbols := [...]string{"BTC", "ETH"}
	exchange1 := strings.Split("aggregated, bibox, bigone, binance, bitfinex, bitget, bithumb, bitmex, bitstamp, bittrex, bybit, checksig.com, cobinhood, coinbase, coincheck, coinex, crypto.com, deribit, ftx, gate.io, gemini, hitbtc, huobi, korbit, kraken, kucoin, luno, nexo, okex, poloniex, swissborg, zb.com", ", ")
	exchange2 := strings.Split("aggregated, bibox, bigone, binance, bitfinex, bitget, bithumb, bitmex, bitstamp, bittrex, bybit, cobinhood, coinbase, coincheck, coinex, crypto.com, deribit, ftx, gate.io, gemini, hitbtc, huobi, korbit, kraken, kucoin, luno, nexo, okex, poloniex, swissborg, zb.com", ", ")
	exchange3 := strings.Split("binance, bitfinex, bitget, bitmex, bybit, checksig.com, coinex, crypto.com, deribit, gemini, huobi, korbit, kucoin, okex, swissborg", ", ")
	btcMiners := [...]string{"aggregated", "other", "1THash&58COIN", "AntPool", "ArkPool", "BinancePool", "BitFury", "BitMinter", "Bixin", "BTC.com", "BTC.TOP", "DPool", "F2Pool", "FoundryUSAPool", "Genesis", "HuobiPool", "KuCoinPool", "Lubian.com", "LuxorTech", "MaraPool", "NovaBlock", "OKExPool", "Patoshi", "PegaPool", "Poolin", "SBICrypto", "SigmaPool", "SlushPool", "SpiderPool", "TerraPool", "UKRPool", "Ultimus", "ViaBTC"}

	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/distribution/balance_grayscale_trust?a=BTC&f=CSV", &folder))
	for _, symbol := range symbols {
		for _, exchange := range exchange1 {
			apiInfos.PushBack(modules.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/distribution/balance_exchanges_relative?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange))
			apiInfos.PushBack(modules.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/distribution/balance_exchanges?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange))
		}
		for _, exchange := range exchange2 {
			apiInfos.PushBack(modules.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/distribution/exchange_net_position_change?a=%s&e=%s&i=24h&f=CSV", &folder, &symbol, &exchange))
		}
		for _, exchange := range exchange3 {
			apiInfos.PushBack(modules.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/distribution/proof_of_reserves?a=%s&e=%s&i=24h&f=CSV", &folder, &symbol, &exchange))
			apiInfos.PushBack(modules.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/distribution/proof_of_reserves_all?a=%s&e=%s&i=24h&f=CSV", &folder, &symbol, &exchange))
		}
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/distribution/supply_contracts?a=%s&i=24h&f=CSV", &folder, &symbol))
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/distribution/balance_1pct_holders?a=%s&i=24h&f=CSV", &folder, &symbol))
	}
	a := "ETH"
	apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/distribution/gini?a=%s&i=24h&f=CSV", &folder, &a))
	apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/distribution/herfindahl?a=%s&i=24h&f=CSV", &folder, &a))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/distribution/balance_luna_foundation_guard?a=BTC&i=1h&f=CSV", &folder))
	for _, miner := range btcMiners {
		apiInfos.PushBack(modules.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/distribution/balance_miners_sum?a=BTC&miner=%s&i=24h&f=CSV", &folder, &miner))
	}
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/distribution/balance_miners_all?a=BTC&i=24h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/distribution/balance_miners_change?a=BTC&i=24h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/distribution/balance_mtgox_trustee?a=BTC&i=1h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/distribution/outflows_mtgox_trustee?a=BTC&i=24h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/distribution/balance_otc_desks?a=BTC&i=1h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/distribution/balance_us_government?a=BTC&i=1h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/distribution/balance_wbtc?a=ETH&i=1h&f=CSV", &folder))
}
