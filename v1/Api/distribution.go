package api

import (
	"path/filepath"
	"strings"

	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func SetDistributionApi(tc *modules.TaskController, rootFolder *string) {
	folder := filepath.Join(*rootFolder, "Distribution")
	btc_symbol := "BTC"

	symbols := [...]string{"BTC", "ETH"}
	exchange1 := strings.Split("aggregated, bibox, bigone, binance, bitfinex, bitget, bithumb, bitmex, bitstamp, bittrex, bybit, checksig.com, cobinhood, coinbase, coincheck, coinex, crypto.com, deribit, ftx, gate.io, gemini, hitbtc, huobi, korbit, kraken, kucoin, luno, nexo, okex, poloniex, swissborg, zb.com", ", ")
	exchange2 := strings.Split("aggregated, bibox, bigone, binance, bitfinex, bitget, bithumb, bitmex, bitstamp, bittrex, bybit, cobinhood, coinbase, coincheck, coinex, crypto.com, deribit, ftx, gate.io, gemini, hitbtc, huobi, korbit, kraken, kucoin, luno, nexo, okex, poloniex, swissborg, zb.com", ", ")
	exchange3 := strings.Split("binance, bitfinex, bitget, bitmex, bybit, checksig.com, coinex, crypto.com, deribit, gemini, huobi, korbit, kucoin, okex, swissborg", ", ")
	btcMiners := [...]string{"aggregated", "other", "1THash&58COIN", "AntPool", "ArkPool", "BinancePool", "BitFury", "BitMinter", "Bixin", "BTC.com", "BTC.TOP", "DPool", "F2Pool", "FoundryUSAPool", "Genesis", "HuobiPool", "KuCoinPool", "Lubian.com", "LuxorTech", "MaraPool", "NovaBlock", "OKExPool", "Patoshi", "PegaPool", "Poolin", "SBICrypto", "SigmaPool", "SlushPool", "SpiderPool", "TerraPool", "UKRPool", "Ultimus", "ViaBTC"}

	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/distribution/balance_grayscale_trust?a=%s&f=CSV", &folder, &btc_symbol, true)
	for _, symbol := range symbols {
		for _, exchange := range exchange1 {
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/distribution/balance_exchanges_relative?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, false)
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/distribution/balance_exchanges_all?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, true)
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/distribution/balance_exchanges?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, true)
		}
		for _, exchange := range exchange2 {
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/distribution/exchange_net_position_change?a=%s&e=%s&i=24h&f=CSV", &folder, &symbol, &exchange, true)
		}
		for _, exchange := range exchange3 {
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/distribution/proof_of_reserves?a=%s&e=%s&i=24h&f=CSV", &folder, &symbol, &exchange, true)
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/distribution/proof_of_reserves_all?a=%s&e=%s&i=24h&f=CSV", &folder, &symbol, &exchange, false)
		}
	}
	a := "ETH"
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/distribution/gini?a=%s&i=24h&f=CSV", &folder, &a, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/distribution/herfindahl?a=%s&i=24h&f=CSV", &folder, &a, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/distribution/balance_luna_foundation_guard?a=%s&i=1h&f=CSV", &folder, &btc_symbol, true)
	for _, miner := range btcMiners {
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/distribution/balance_miners_sum?a=%s&miner=%s&i=24h&f=CSV", &folder, &btc_symbol, &miner, false)
	}
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/distribution/balance_miners_all?a=%s&i=24h&f=CSV", &folder, &btc_symbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/distribution/balance_miners_change?a=%s&i=24h&f=CSV", &folder, &btc_symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/distribution/balance_mtgox_trustee?a=%s&i=1h&f=CSV", &folder, &btc_symbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/distribution/outflows_mtgox_trustee?a=%s&i=24h&f=CSV", &folder, &btc_symbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/distribution/balance_otc_desks?a=%s&i=1h&f=CSV", &folder, &btc_symbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/distribution/supply_contracts?a=%s&i=24h&f=CSV", &folder, &a, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/distribution/balance_1pct_holders?a=%s&i=24h&f=CSV", &folder, &a, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/distribution/balance_us_government?a=%s&i=1h&f=CSV", &folder, &btc_symbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/distribution/balance_wbtc?a=%s&i=1h&f=CSV", &folder, &btc_symbol, true)
}
