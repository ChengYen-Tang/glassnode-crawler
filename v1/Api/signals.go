package api

import (
	"path/filepath"

	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func SetSignalsApi(tc *modules.TaskController, rootFolder *string) {
	folder := filepath.Join(*rootFolder, "Signals")

	btcSymbol := "BTC"
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/signals/altcoin_index?a=%s&i=24h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/signals/btc_risk_index?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/signals/btc_sharpe_signal?a=%s&i=24h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/signals/mcap_returns_vs_btc?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/signals/l1_native_tokens_returns_vs_btc?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
}
