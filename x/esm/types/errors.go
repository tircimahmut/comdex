package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/esm module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrorUnknownProposalType = sdkerrors.Register(ModuleName, 401, "unknown proposal type")
	
)
