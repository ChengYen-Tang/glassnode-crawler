package api

import (
	"path/filepath"

	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func SetEth20Api(tc *modules.TaskController, rootFolder *string) {
	folder := filepath.Join(*rootFolder, "ETH2.0")
	symbol := "ETH"

	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/eth2/active_validators_count?a=%s&i=1h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/eth2/attestation_count?a=%s&i=1h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/eth2/validator_balance_mean?a=%s&i=1h&f=CSV", &folder, &symbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/eth2/epoch_height?a=%s&i=1h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/eth2/estimated_annual_issuance?a=%s&i=1h&f=CSV", &folder, &symbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/eth2/estimated_annual_issuance_roi_per_validator?a=%s&i=1h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/eth2/staking_deposits_count?a=%s&i=1h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/eth2/staking_validators_count?a=%s&i=1h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/eth2/staking_volume_sum?a=%s&i=1h&f=CSV", &folder, &symbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/eth2/staking_phase_0_goal_percent?a=%s&i=1h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/eth2/staking_total_deposits_count?a=%s&i=1h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/eth2/staking_total_validators_count?a=%s&i=1h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/eth2/staking_total_volume_sum?a=%s&i=1h&f=CSV", &folder, &symbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/eth2/deposited_by_provider_volume_sum?a=%s&i=24h&f=CSV", &folder, &symbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/eth2/missed_blocks_count?a=%s&i=1h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/eth2/orphaned_blocks_count?a=%s&i=1h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/eth2/participation_rate_mean?a=%s&i=1h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/eth2/slashings_count?a=%s&i=1h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/eth2/slot_height?a=%s&i=1h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/eth2/stake_effectiveness_mean?a=%s&i=1h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/eth2/deposits_count?a=%s&i=1h&f=CSV", &folder, &symbol, false)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/eth2/effective_balance_sum?a=%s&i=1h&f=CSV", &folder, &symbol, true)
	tc.GetAPIInfoTag("https://api.glassnode.com/v1/metrics/eth2/voluntary_exit_count?a=%s&i=1h&f=CSV", &folder, &symbol, false)
}
