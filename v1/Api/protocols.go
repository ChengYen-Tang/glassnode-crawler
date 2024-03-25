package api

import (
	"path/filepath"
	"strings"

	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func SetProtocolsApi(tc *modules.TaskController, rootFolder *string) {
	folder := filepath.Join(*rootFolder, "Protocols")

	aave := "AAVE_PROTOCOL"
	compound := "COMPOUND_PROTOCOL"
	uniswap := "UNISWAP_PROTOCOL"
	lido := "LIDO_PROTOCOL"
	market := "MAKER_PROTOCOL"
	uniSymbols := strings.Split("ETH, 1INCH, AAVE, ABT, AMP, AMPL, ANT, APE, BADGER, BAL, BAND, BAT, BNT, BOBA, BOND, BRD, BUSD, CAKE, CELR, CHSB, COMP, CREAM, CRO, CRV, CVC, CVP, CVX, CVXCRV, DAI, DDX, DENT, DHT, DODO, DPI, DRGN, DYDX, ELF, ENG, ENJ, EURS, FET, FLX, FRAX, FTM, FTT, FUN, FXS, GNO, GUSD, HEGIC, HOT, HT, HUSD, IMX, INDEX, KCS, LAMB, LBA, LDO, LEO, LINK, LOOM, LRC, MANA, MATIC, MCB, METIS, MIR, MKR, MLN, MTA, MTL, NDX, NEXO, NFTX, NMR, Nsure, OCEAN, OKB, OMG, PAY, PERP, PICKLE, PNK, PNT, POLY, POWR, PPT, PYUSD, QASH, QKC, QNT, RAI, RDN, REN, REP, rETH, RLC, ROOK, RPL, RSR, SAND, SFRXETH, SHIB, SNT, SNX, SSV, STAKE, stETH, STORJ, sUSD, SUSHI, TEL, TUSD, UBT, UMA, UNI, USDC, USDD, USDK, USDP, USDT, UTK, VERI, WaBi, WBTC, WETH, wNXM, YAM, YFI, ZRX", ", ")

	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v1_available_liquidity_sum_by_token?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v1_available_liquidity_perc_by_token?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v1_collateral_rate_by_token?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v1_count_by_action?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v1_stable_borrow_rate_by_token?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v1_variable_borrow_rate_by_token?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v1_volume_sum_by_action?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v2_total_value_locked_sum?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v2_total_value_locked_perc?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v2_available_liquidity_perc_by_token?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v2_available_liquidity_sum_by_token?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v2_borrows_sum_by_token?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v2_borrows_perc_by_token?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v2_collateral_rate_by_token?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v2_count_by_action?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v2_stable_borrow_rate_by_token?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v2_variable_borrow_rate_by_token?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v2_volume_sum_by_action?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v3_total_value_locked_sum?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v3_total_value_locked_perc?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v3_available_liquidity_perc_by_token?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v3_available_liquidity_sum_by_token?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v3_borrows_sum_by_token?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v3_borrows_perc_by_token?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v3_collateral_rate_by_token?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v3_count_by_action?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v3_stable_borrow_rate_by_token?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v3_variable_borrow_rate_by_token?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/aave_v3_volume_sum_by_action?a=%s&i=24h&f=CSV", &folder, &aave, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/compound_deposit_flow?a=%s&i=24h&f=CSV", &folder, &compound, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/compound_loan_flow?a=%s&i=24h&f=CSV", &folder, &compound, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/compound_total_value_locked?a=%s&i=24h&f=CSV", &folder, &compound, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/compound_count_by_action?a=%s&i=24h&f=CSV", &folder, &compound, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/compound_volume_by_action?a=%s&i=24h&f=CSV", &folder, &compound, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/compound_utilization_ratio?a=%s&i=24h&f=CSV", &folder, &compound, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/uniswap_fees_sum?a=%s&i=24h&f=CSV", &folder, &uniswap, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/lido_deposits_volume_sum?a=%s&i=1h&f=CSV", &folder, &lido, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/lido_volume_net?a=%s&i=1h&f=CSV", &folder, &lido, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/lido_total_value_locked?a=%s&i=1h&f=CSV", &folder, &lido, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/lido_withdrawn_volume_sum?a=%s&i=1h&f=CSV", &folder, &lido, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/maker_event_count_by_type?a=%s&i=24h&f=CSV", &folder, &market, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/maker_event_volume_usd_by_type?a=%s&i=24h&f=CSV", &folder, &market, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/maker_new_vaults_created_count?a=%s&i=24h&f=CSV", &folder, &market, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/maker_psm_balance_sum?a=%s&i=24h&f=CSV", &folder, &market, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/maker_total_value_locked_sum?a=%s&i=24h&f=CSV", &folder, &market, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/maker_total_value_locked_sum_by_token?a=%s&i=24h&f=CSV", &folder, &market, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/maker_total_vaults_created_cum?a=%s&i=24h&f=CSV", &folder, &market, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/uniswap_total_value_locked?a=%s&i=24h&f=CSV", &folder, &uniswap, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/uniswap_volume_traded_sum?a=%s&i=24h&f=CSV", &folder, &uniswap, false)
	for _, uniSymbol := range uniSymbols {
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/uniswap_liquidity_latest?a=%s&i=24h&f=CSV", &folder, &uniSymbol, true)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/uniswap_v3_liquidity_latest?a=%s&i=24h&f=CSV", &folder, &uniSymbol, true)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/uniswap_transaction_count?a=%s&i=24h&f=CSV", &folder, &uniSymbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/uniswap_v3_transaction_count?a=%s&i=24h&f=CSV", &folder, &uniSymbol, false)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/uniswap_volume_sum?a=%s&i=24h&f=CSV", &folder, &uniSymbol, true)
		tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/protocols/uniswap_v3_volume_sum?a=%s&i=24h&f=CSV", &folder, &uniSymbol, true)
	}
}
