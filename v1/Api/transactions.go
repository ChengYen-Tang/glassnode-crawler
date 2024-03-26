package api

import (
	"path/filepath"
	"strings"

	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func SetTransactionsApi(tc *modules.TaskController, rootFolder *string) {
	folder := filepath.Join(*rootFolder, "Transactions")

	btcSymbol := "BTC"
	ethSymbol := "ETH"
	symbols := [...]string{btcSymbol, ethSymbol}
	exange1 := strings.Split("aggregated, bibox, bigone, binance, bitfinex, bitget, bithumb, bitmex, bitstamp, bittrex, bybit, cobinhood, coinbase, coincheck, coinex, crypto.com, deribit, ftx, gate.io, gemini, hitbtc, huobi, korbit, kraken, kucoin, luno, nexo, okex, poloniex, swissborg, zb.com", ", ")
	exange2 := strings.Split("aggregated, binance, bitfinex, bithumb, bitmex, bitstamp, bittrex, bybit, coinbase, coincheck, crypto.com, deribit, ftx, gate.io, gemini, hitbtc, huobi, kraken, kucoin, luno, okex, poloniex", ", ")
	exange3 := strings.Split("aggregated, binance, bitfinex, bitget, bithumb, bitmex, bitstamp, bittrex, bybit, coinbase, coincheck, coinex, crypto.com, deribit, ftx, gate.io, gemini, hitbtc, huobi, korbit, kraken, kucoin, luno, okex, poloniex", ", ")
	miners1 := [...]string{"aggregated", "other", "1THash&58COIN", "AntPool", "ArkPool", "BinancePool", "BitFury", "BitMinter", "Bixin", "BTC.com", "BTC.TOP", "DPool", "F2Pool", "FoundryUSAPool", "Genesis", "HuobiPool", "KuCoinPool", "Lubian.com", "LuxorTech", "MaraPool", "NovaBlock", "OKExPool", "Patoshi", "PegaPool", "Poolin", "SBICrypto", "SigmaPool", "SlushPool", "SpiderPool", "TerraPool", "UKRPool", "Ultimus", "ViaBTC"}

	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_count_bridges?a=%s&i=24h&f=CSV", &folder, &ethSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_count_bridges_relative?a=%s&i=24h&f=CSV", &folder, &ethSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_adjusted_mean?a=%s&i=1h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_adjusted_median?a=%s&i=1h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_adjusted_sum?a=%s&i=1h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_from_coinjoins_count?a=%s&i=24h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_from_coinjoins_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/contract_calls_external_count?a=%s&i=1h&f=CSV", &folder, &ethSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/contract_calls_internal_count?a=%s&i=1h&f=CSV", &folder, &ethSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_count_defi?a=%s&i=24h&f=CSV", &folder, &ethSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_count_defi_relative?a=%s&i=24h&f=CSV", &folder, &ethSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_entity_adjusted_from_lth_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_entity_adjusted_from_lth_loss_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_entity_adjusted_from_lth_profit_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_entity_adjusted_from_lth_sth_profit_loss_relative?a=%s&i=24h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_entity_adjusted_from_sth_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_entity_adjusted_from_sth_loss_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_entity_adjusted_from_sth_profit_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/entity_adjusted_count?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_entity_adjusted_mean?a=%s&i=1h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_entity_adjusted_median?a=%s&i=1h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_entity_adjusted_sum?a=%s&i=1h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_count_erc20?a=%s&i=24h&f=CSV", &folder, &ethSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_count_erc20_relative?a=%s&i=24h&f=CSV", &folder, &ethSymbol, false)
	for _, symbol := range symbols {
		for _, exange := range exange1 {
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_to_exchanges_count?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exange, false)
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_to_exchanges_mean?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exange, true)
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_to_exchanges_sum?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exange, true)
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_exchanges_net?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exange, true)
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_from_exchanges_mean?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exange, true)
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_from_exchanges_sum?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exange, true)
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_from_exchanges_count?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exange, false)
			tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_within_exchanges_sum?a=%s&e=%s&i=1h&f=CSV", &folder, &symbol, &exange, true)
		}
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/count?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/rate?a=%s&i=1h&f=CSV", &folder, &symbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_mean?a=%s&i=1h&f=CSV", &folder, &symbol, true)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_median?a=%s&i=1h&f=CSV", &folder, &symbol, true)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_sum?a=%s&i=1h&f=CSV", &folder, &symbol, true)
	}

	for _, exange := range exange2 {
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_exchanges_net_by_size?a=%s&e=%s&i=24h&f=CSV", &folder, &btcSymbol, &exange, true)
	}

	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/inscriptions_size_share?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/inscriptions_fee_share?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/inscriptions_fee?a=%s&i=1h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/inscriptions_size_sum?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/inscriptions_count_sum?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/inscriptions_count?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/inscriptions_count_share?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_lth_to_exchanges_loss_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_lth_to_exchanges_profit_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_lth_to_exchanges_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_lth_sth_to_exchanges_profit_loss_relative?a=%s&i=24h&f=CSV", &folder, &btcSymbol, false)

	for _, miner := range miners1 {
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_to_miners_count?a=%s&m=%s&i=1h&f=CSV", &folder, &btcSymbol, &miner, false)
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_to_miners_sum?a=%s&m=%s&i=1h&f=CSV", &folder, &btcSymbol, &miner, true)
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_miners_net?a=%s&m=%s&i=1h&f=CSV", &folder, &btcSymbol, &miner, true)
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_from_miners_sum?a=%s&m=%s&i=1h&f=CSV", &folder, &btcSymbol, &miner, true)
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_from_miners_count?a=%s&m=%s&i=1h&f=CSV", &folder, &btcSymbol, &miner, false)
	}

	for _, exange := range exange3 {
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_miners_to_exchanges?a=%s&e=%s&i=1h&f=CSV", &folder, &btcSymbol, &exange, true)
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_miners_to_exchanges_all?a=%s&e=%s&i=24h&f=CSV", &folder, &btcSymbol, &exange, true)
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_whales_to_exchanges_count?a=%s&e=%s&i=24h&f=CSV", &folder, &btcSymbol, &exange, false)
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_whales_to_exchanges_sum?a=%s&e=%s&i=24h&f=CSV", &folder, &btcSymbol, &exange, true)
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_exchanges_to_whales_count?a=%s&e=%s&i=24h&f=CSV", &folder, &btcSymbol, &exange, false)
		tc.GetAPIInfoTwoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_exchanges_to_whales_sum?a=%s&e=%s&i=24h&f=CSV", &folder, &btcSymbol, &exange, true)
	}
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_nft_marketplaces_count_all?a=%s&i=24h&f=CSV", &folder, &ethSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_nft_marketplaces_volume_sum_all?a=%s&i=24h&f=CSV", &folder, &ethSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_count_nfts?a=%s&i=24h&f=CSV", &folder, &ethSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_count_nfts_relative?a=%s&i=24h&f=CSV", &folder, &ethSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_to_otc_desks_count?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_to_otc_desks_sum?a=%s&i=1h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_from_otc_desks_sum?a=%s&i=1h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_from_otc_desks_count?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_profit_relative?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_by_size_entity_adjusted_relative?a=%s&i=24h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/segwit_adoption?a=%s&i=24h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_sth_to_exchanges_loss_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_sth_to_exchanges_profit_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_sth_to_exchanges_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/spent_output_types_share?a=%s&i=24h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_count_stablecoins?a=%s&i=24h&f=CSV", &folder, &ethSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_count_stablecoins_relative?a=%s&i=24h&f=CSV", &folder, &ethSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/taproot_adoption?a=%s&i=24h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_by_size_entity_adjusted_sum?a=%s&i=24h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/size_mean?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/size_sum?a=%s&i=1h&f=CSV", &folder, &btcSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/tx_types_breakdown_count?a=%s&i=24h&f=CSV", &folder, &ethSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/tx_types_breakdown_relative?a=%s&i=24h&f=CSV", &folder, &ethSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_count?a=%s&i=1h&f=CSV", &folder, &ethSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_rate?a=%s&i=1h&f=CSV", &folder, &ethSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_loss_sum?a=%s&i=1h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_volume_profit_sum?a=%s&i=1h&f=CSV", &folder, &btcSymbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_count_vanilla?a=%s&i=24h&f=CSV", &folder, &ethSymbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/transactions/transfers_count_vanilla_relative?a=%s&i=24h&f=CSV", &folder, &ethSymbol, false)
}
