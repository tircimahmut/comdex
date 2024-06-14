package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/comdex-official/comdex/x/tokenmint/types"
)

// GetParams get all parameters as types.Params.
func (k Keeper) GetParams(_ sdk.Context) types.Params {
	return types.NewParams()
}

// SetParams set the params.
func (k Keeper) SetParams(_ sdk.Context, params types.Params) {
}
