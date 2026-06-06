// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package keeper

import (
	"fmt"
	"math"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey storetypes.StoreKey
		memKey   storetypes.StoreKey
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// TODO issue 99: these getters were initially created
//
//	for tests needs: to link dependent keepers,
//	need to explore the alternatives
func (k Keeper) StoreKey() storetypes.StoreKey {
	return k.storeKey
}

func (k Keeper) MemKey() storetypes.StoreKey {
	return k.memKey
}

func (k Keeper) AccountApprovalsCount(ctx sdk.Context, percent float64) int {
	approvalsCount := int(math.Ceil(percent * float64(k.CountAccountsWithRole(ctx, types.Trustee))))
	if approvalsCount == 0 {
		return 1
	}

	return approvalsCount
}

func (k Keeper) AccountRejectApprovalsCount(ctx sdk.Context, percent float64) int {
	return k.CountAccountsWithRole(ctx, types.Trustee) - k.AccountApprovalsCount(ctx, percent) + 1
}

func (k Keeper) UnmarshalAccount(bytes []byte) (authtypes.AccountI, error) {
	var acc authtypes.AccountI

	return acc, k.cdc.UnmarshalInterface(bytes, &acc)
}

func (k Keeper) GetCodec() codec.BinaryCodec {
	return k.cdc
}
