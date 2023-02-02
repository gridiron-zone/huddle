package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/gridiron-zone/huddle/x/posts/types"
)

// SetParams sets params on the store
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramsSubspace.SetParamSet(ctx, &params)
}

// GetParams returns the params from the store
func (k Keeper) GetParams(ctx sdk.Context) (p types.Params) {
	k.paramsSubspace.GetParamSet(ctx, &p)
	return p
}
