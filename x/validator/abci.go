package validator

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// slash/jail (remove from the validator set) a node with detected Byzantine behaviour
	// (double signing, light client attack, etc.)
	for _, tmEvidence := range req.ByzantineValidators {
		switch tmEvidence.Type {
		case abci.EvidenceType_DUPLICATE_VOTE, abci.EvidenceType_LIGHT_CLIENT_ATTACK:
			evidence := evidencetypes.FromABCIEvidence(tmEvidence)
			k.HandleDoubleSign(ctx, evidence.(*evidencetypes.Equivocation))
		default:
			k.Logger(ctx).Error(fmt.Sprintf("ignored unknown evidence type: %s", tmEvidence.Type))
		}
	}

	// https://github.com/cosmos/modules/tree/module/poa/incubator/poa and
	// default Cosmos https://github.com/cosmos/cosmos-sdk/tree/master/x/slashing modules
	// have logic to slash/jail validators which doesn't sign too many blocks in a row (downtime slashing)
	// See https://github.com/cosmos/cosmos-sdk/blob/master/x/slashing/keeper/infractions.go.
	// this may be OK in a permisionless network, but not OK in a persmissioned network like DCL

	// Why it's not OK:
	// - Let's assume there is a number of malicious nodes in the pool.
	// - That malicious nodes star to DoS valid nodes one by one, so valid nodes can not sign blocks, slashed and removed from the validators set
	// => validators consists of malicious nodes only (or at least 2/3 of nodes are malicious)
	// It's more safe to just stop writing in case of such DoS attack than let malicious nodes take control of the pool
	// Moreover, the attack is easy to reproduce, as it's sufficient to DoS valid nodes one by one,
	// while > 1/3 nodes must be DoS to prevent the pool from writing.

	// That's why we don't implement/call/copy HandleValidatorSignature (downtime slashing) logic in DCL.
}

// Called every block, update validator set.
func EndBlocker(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	res, err := k.ApplyAndReturnValidatorSetUpdates(ctx)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Error when applying a new validator set: %s", err.Error()))
	}
	return res
}
