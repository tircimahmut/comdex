package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = (*MsgMintNewTokensRequest)(nil)
	_ sdk.Msg = (*MsgBurnHarborTokensRequest)(nil)
)

func NewMsgMintNewTokensRequest(from string, appID uint64, assetID uint64) *MsgMintNewTokensRequest {
	return &MsgMintNewTokensRequest{
		From:    from,
		AppId:   appID,
		AssetId: assetID,
	}
}

func (m *MsgMintNewTokensRequest) Route() string {
	return RouterKey
}

func (m *MsgMintNewTokensRequest) Type() string {
	return TypeMsgMintNewTokensRequest
}

func (m *MsgMintNewTokensRequest) ValidateBasic() error {
	if m.From == "" {
		return errorsmod.Wrap(ErrorInvalidFrom, "from cannot be empty")
	}
	if m.AppId == 0 {
		return errorsmod.Wrap(ErrorInvalidAppID, "app id can not be zero")
	}
	if m.AssetId == 0 {
		return errorsmod.Wrap(ErrorInvalidAssetID, "asset id can not be zero")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errorsmod.Wrapf(ErrorInvalidFrom, "%s", err)
	}

	return nil
}

func (m *MsgMintNewTokensRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m *MsgMintNewTokensRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

func NewMsgBurnHarborTokensRequest(from string, appID uint64, burnCoin sdk.Coin) *MsgBurnHarborTokensRequest {
	return &MsgBurnHarborTokensRequest{
		From:      from,
		AppId:     appID,
		BurnCoins: burnCoin,
	}
}

func (m *MsgBurnHarborTokensRequest) Route() string {
	return RouterKey
}

func (m *MsgBurnHarborTokensRequest) Type() string {
	return TypeMsgBurnHarborTokensRequest
}

func (m *MsgBurnHarborTokensRequest) ValidateBasic() error {
	if m.From == "" {
		return errorsmod.Wrap(ErrorInvalidFrom, "from cannot be empty")
	}
	if m.AppId == 0 {
		return errorsmod.Wrap(ErrorInvalidAppID, "app id can not be zero or negative")
	}
	if !m.BurnCoins.IsValid() {
		return errorsmod.Wrapf(errors.ErrInvalidCoins, "bid amount %s", m.BurnCoins)
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errorsmod.Wrapf(ErrorInvalidFrom, "%s", err)
	}

	return nil
}

func (m *MsgBurnHarborTokensRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m *MsgBurnHarborTokensRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}
