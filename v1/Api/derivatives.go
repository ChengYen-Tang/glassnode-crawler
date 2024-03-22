package api

import (
	"path/filepath"

	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func SetDerivativesApi(tc *modules.TaskController, rootFolder *string) {
	folder := filepath.Join(*rootFolder, "Derivatives")

	symbols := [...]string{"BTC", "ETH"}
	exchange1 := [...]string{"aggregated", "binance", "bitfinex", "bitmex", "bybit", "deribit", "huobi", "kraken", "okex"}
	exchange2 := [...]string{"aggregated", "binance", "bybit", "bitmex", "deribit", "ftx", "okex"}
	exchange3 := [...]string{"aggregated", "binance", "bitfinex", "bitmex", "bybit", "deribit", "ftx", "huobi", "kraken", "okex"}
	exchange4 := [...]string{"aggregated", "binance", "bitmex", "okex"}
	exchange5 := [...]string{"aggregated", "binance", "bybit", "deribit", "huobi", "okex"}
	exchange6 := [...]string{"binance", "bybit", "deribit", "huobi", "okex"}
	exchange7 := [...]string{"aggregated", "binance", "bitfinex", "bitflyer", "bitmex", "bybit", "deribit", "ftx", "huobi", "kraken", "okex"}
	exchange8 := [...]string{"aggregated", "deribit", "ftx", "okex"}
	exchange9 := [...]string{"aggregated", "deribit"}
	deribit_exchange := "deribit"

	for _, symbol := range symbols {
		for _, exchange := range exchange1 {
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/futures_open_interest_cash_margin_sum?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, false)
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/futures_open_interest_crypto_margin_sum?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, false)
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/futures_open_interest_crypto_margin_relative?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, false)
		}
		for _, exchange := range exchange2 {
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/futures_annualized_basis_3m?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, false)
		}
		for _, exchange := range exchange3 {
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/futures_estimated_leverage_ratio?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, false)
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/futures_open_interest_sum?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, true)
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/futures_open_interest_perpetual_sum?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, true)
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/futures_funding_rate_perpetual?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, false)
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/futures_funding_rate_perpetual_v2?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, false)
		}
		for _, exchange := range exchange4 {
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/futures_liquidated_volume_long_mean?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, true)
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/futures_liquidated_volume_long_sum?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, true)
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/futures_liquidated_volume_long_relative?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, true)
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/futures_liquidated_volume_short_mean?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, true)
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/futures_liquidated_volume_short_sum?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, true)
		}
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/derivatives/futures_open_interest_sum_all?a=%s&i=1h&f=CSV", &folder, &symbol, true)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/derivatives/futures_open_interest_perpetual_sum_all?a=%s&i=1h&f=CSV", &folder, &symbol, true)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/derivatives/futures_funding_rate_perpetual_all?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/derivatives/futures_funding_rate_perpetual_all_v2?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		for _, exchange := range exchange5 {
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/futures_term_structure?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, false)
		}
		for _, exchange := range exchange6 {
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/futures_term_structure_by_exchange?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, false)
		}
		for _, exchange := range exchange7 {
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/futures_volume_daily_sum?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, true)
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/futures_volume_daily_perpetual_sum?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, true)
		}
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/derivatives/futures_volume_daily_sum_all?a=%s&i=1h&f=CSV", &folder, &symbol, true)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/derivatives/futures_volume_daily_perpetual_sum_all?a=%s&i=1h&f=CSV", &folder, &symbol, true)
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/options_25delta_skew_1_month?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &deribit_exchange, false)
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/options_25delta_skew_1_week?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &deribit_exchange, false)
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/options_25delta_skew_3_months?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &deribit_exchange, false)
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/options_25delta_skew_6_months?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &deribit_exchange, false)
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/options_25delta_skew_all?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &deribit_exchange, false)
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/options_atm_implied_volatility_1_month?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &deribit_exchange, false)
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/options_atm_implied_volatility_1_week?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &deribit_exchange, false)
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/options_atm_implied_volatility_3_months?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &deribit_exchange, false)
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/options_atm_implied_volatility_6_months?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &deribit_exchange, false)
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/options_atm_implied_volatility_all?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &deribit_exchange, false)
		for _, exchange := range exchange8 {
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/options_open_interest_sum?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, true)
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/options_volume_daily_sum?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, true)
		}
		for _, exchange := range exchange9 {
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/options_open_interest_put_call_ratio?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, false)
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/derivatives/options_volume_put_call_ratio?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exchange, false)
		}
	}
}
