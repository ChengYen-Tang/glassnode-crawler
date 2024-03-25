package api

import (
	"path/filepath"

	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func SetSupplyApi(tc *modules.TaskController, rootFolder *string) {
	folder := filepath.Join(*rootFolder, "Supply")

	btcSymbol := "BTC"
	ethSymbol := "ETH"
	symbols := [...]string{btcSymbol, ethSymbol}
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/current_adjusted?a=%s&i=1h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/apac_1y_supply_change?a=%s&i=24h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/burn_rate?a=%s&i=1h&f=CSV", &folder, &ethSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/burned?a=%s&i=1h&f=CSV", &folder, &ethSymbol, true)
	for _, symbol := range symbols {
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/current?a=%s&i=1h&f=CSV", &folder, &symbol, true)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/hodl_waves?a=%s&i=24h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/inflation_rate?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/issued?a=%s&i=1h&f=CSV", &folder, &symbol, true)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/profit_relative?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/loss_sum?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/profit_sum?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/active_24h?a=%s&i=1h&f=CSV", &folder, &symbol, true)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/active_more_10y?a=%s&i=1h&f=CSV", &folder, &symbol, true)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/active_1d_1w?a=%s&i=1h&f=CSV", &folder, &symbol, true)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/active_1m_3m?a=%s&i=1h&f=CSV", &folder, &symbol, true)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/active_1w_1m?a=%s&i=1h&f=CSV", &folder, &symbol, true)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/active_1y_2y?a=%s&i=1h&f=CSV", &folder, &symbol, true)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/active_2y_3y?a=%s&i=1h&f=CSV", &folder, &symbol, true)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/active_3m_6m?a=%s&i=1h&f=CSV", &folder, &symbol, true)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/active_3y_5y?a=%s&i=1h&f=CSV", &folder, &symbol, true)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/active_5y_7y?a=%s&i=1h&f=CSV", &folder, &symbol, true)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/active_6m_12m?a=%s&i=1h&f=CSV", &folder, &symbol, true)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/active_7y_10y?a=%s&i=1h&f=CSV", &folder, &symbol, true)
	}
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/emea_1y_supply_change?a=%s&i=24h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/highly_liquid_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/illiquid_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/illiquid_change?a=%s&i=24h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/liquid_illiquid_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/liquid_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/liquid_change?a=%s&i=24h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/lth_net_change?a=%s&i=24h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/lth_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/lth_loss_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/lth_profit_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/minted?a=%s&i=1h&f=CSV", &folder, &ethSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/probably_lost?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/provably_lost?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/rcap_hodl_waves?a=%s&i=24h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/lth_sth_profit_loss_relative?a=%s&i=24h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/revived_more_1y_sum?a=%s&i=1h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/revived_more_2y_sum?a=%s&i=1h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/revived_more_3y_sum?a=%s&i=1h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/revived_more_5y_sum?a=%s&i=1h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/sth_lth_realized_value_ratio?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/sth_profit_loss_ratio?a=%s&i=24h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/sth_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/sth_loss_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/sth_profit_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/supply_by_txout_type?a=%s&i=24h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/active_more_1y_percent?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/active_more_2y_percent?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/active_more_3y_percent?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/active_more_5y_percent?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/tips?a=%s&i=1h&f=CSV", &folder, &ethSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/supply/amer_1y_supply_change?a=%s&i=24h&f=CSV", &folder, &btcSymbol, false)
}
