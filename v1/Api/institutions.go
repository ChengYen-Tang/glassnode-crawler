package api

import (
	"path/filepath"

	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func SetInstitutionsApi(tc *modules.TaskController, rootFolder *string) {
	folder := filepath.Join(*rootFolder, "Institutions")

	btcSymbol := "BTC"
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/institutions/purpose_etf_aum_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/institutions/purpose_etf_flows_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/institutions/purpose_etf_holdings_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/institutions/btc_us_spot_etf_balances_all?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/institutions/btc_us_spot_etf_balances_latest?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/institutions/btc_us_spot_etf_flows_all?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
}
