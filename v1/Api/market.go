package api

import (
	"path/filepath"
	"strings"

	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func SetMarketApi(tc *modules.TaskController, rootFolder *string) {
	folder := filepath.Join(*rootFolder, "Market")

	btcSymbol := "BTC"
	ethSymbol := "ETH"
	symbols := []string{btcSymbol, ethSymbol}
	exanges := strings.Split("aggregated, binance, bitfinex, bitstamp, bybit, coinbase, crypto.com, gemini, huobi, kraken, kucoin, okex, poloniex", ", ")

	for _, symbol := range symbols {
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/market/apac_30d_price_change?a=%s&i=24h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/market/btc_dominance?a=%s&i=24h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/market/emea_30d_price_change?a=%s&i=24h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/market/marketcap_usd?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/market/mvrv?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/market/mvrv_z_score?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/market/price_usd_close?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/market/price_drawdown_relative?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/market/price_usd_ohlc?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/market/marketcap_realized_usd?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/market/price_realized_usd?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/market/realized_volatility_1_month?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/market/realized_volatility_1_week?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/market/realized_volatility_1_year?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/market/realized_volatility_2_weeks?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/market/realized_volatility_3_months?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/market/realized_volatility_6_months?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/market/realized_volatility_all?a=%s&i=1h&f=CSV", &folder, &symbol, false)

		for _, exchange := range exanges {
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/market/spot_buying_volume_sum?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, true)
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/market/spot_cvd_sum?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, true)
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/market/spot_selling_volume_sum?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, true)
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/market/spot_volume_daily_sum?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, true)
		}

		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/market/spot_volume_daily_latest?a=%s&i=10m&f=CSV", &folder, &symbol, true)
	}
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/market/deltacap_usd?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/market/mvrv_more_155?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
}
