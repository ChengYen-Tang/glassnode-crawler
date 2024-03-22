package api

import (
	"container/list"
	"path/filepath"

	"github.com/ChengYen-Tang/glassnode-crawler/modules"
)

func SetEth20Api(apiInfos *list.List, rootFolder *string) {
	folder := filepath.Join(*rootFolder, "ETH2.0")

	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/eth2/active_validators_count?a=ETH&i=1h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/eth2/attestation_count?a=ETH&i=1h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/eth2/validator_balance_mean?a=ETH&i=1h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/eth2/epoch_height?a=ETH&i=1h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/eth2/estimated_annual_issuance?a=ETH&i=1h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/eth2/estimated_annual_issuance_roi_per_validator?a=ETH&i=1h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/eth2/staking_deposits_count?a=ETH&i=1h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/eth2/staking_validators_count?a=ETH&i=1h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/eth2/staking_volume_sum?a=ETH&i=1h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/eth2/staking_phase_0_goal_percent?a=ETH&i=1h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/eth2/staking_total_deposits_count?a=ETH&i=1h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/eth2/staking_total_validators_count?a=ETH&i=1h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/eth2/staking_total_volume_sum?a=ETH&i=1h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/eth2/deposited_by_provider_volume_sum?a=ETH&i=24h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/eth2/missed_blocks_count?a=ETH&i=1h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/eth2/orphaned_blocks_count?a=ETH&i=1h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/eth2/participation_rate_mean?a=ETH&i=1h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/eth2/slashings_count?a=ETH&i=1h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/eth2/slot_height?a=ETH&i=1h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/eth2/stake_effectiveness_mean?a=ETH&i=1h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/eth2/deposits_count?a=ETH&i=1h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/eth2/effective_balance_sum?a=ETH&i=1h&f=CSV", &folder))
	apiInfos.PushBack(modules.GetAPIInfoDefault("https://api.glassnode.com/v1/metrics/eth2/voluntary_exit_count?a=ETH&i=1h&f=CSV", &folder))

}
