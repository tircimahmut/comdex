package market

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/comdex-official/comdex/x/market/keeper"
	"github.com/comdex-official/comdex/x/market/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(_ keeper.Keeper) sdk.Handler {
	return func(_ sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		default:
			return nil, errorsmod.Wrapf(types.ErrorUnknownMsgType, "%T", msg)
		}
	}
}
