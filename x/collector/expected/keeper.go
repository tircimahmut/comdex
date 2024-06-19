package expected

import (
	"github.com/comdex-official/comdex/x/asset/types"
	auctiontypes "github.com/comdex-official/comdex/x/auction/types"
	esmtypes "github.com/comdex-official/comdex/x/esm/types"
	lockertypes "github.com/comdex-official/comdex/x/locker/types"
	rewardstypes "github.com/comdex-official/comdex/x/rewards/types"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
)

type BankKeeper interface {
	SendCoinsFromAccountToModule(ctx sdktypes.Context, address sdktypes.AccAddress, name string, coins sdktypes.Coins) error
	SendCoinsFromModuleToAccount(ctx sdktypes.Context, name string, address sdktypes.AccAddress, coins sdktypes.Coins) error

	SendCoinsFromModuleToModule(ctx sdktypes.Context, senderModule, recipientModule string, amt sdktypes.Coins) error
	GetBalance(ctx sdktypes.Context, addr sdktypes.AccAddress, denom string) sdktypes.Coin
}

type AssetKeeper interface {
	HasAssetForDenom(ctx sdktypes.Context, id string) bool
	HasAsset(ctx sdktypes.Context, id uint64) bool
	GetAssetForDenom(ctx sdktypes.Context, denom string) (types.Asset, bool)
	GetApp(ctx sdktypes.Context, id uint64) (types.AppData, bool)
	GetAsset(ctx sdktypes.Context, id uint64) (types.Asset, bool)
	GetMintGenesisTokenData(ctx sdktypes.Context, appID, assetID uint64) (mintData types.MintGenesisToken, found bool)
}

type AuctionKeeper interface {
	GetAuctionParams(ctx sdktypes.Context, appID uint64) (asset auctiontypes.AuctionParams, found bool)
}

type LockerKeeper interface {
	GetLockerLookupTable(ctx sdktypes.Context, appID, assetID uint64) (lockerLookupData lockertypes.LockerLookupTableData, found bool)
	GetLocker(ctx sdktypes.Context, lockerID uint64) (locker lockertypes.Locker, found bool)
	SetLocker(ctx sdktypes.Context, locker lockertypes.Locker)
	SetLockerLookupTable(ctx sdktypes.Context, lockerLookupData lockertypes.LockerLookupTableData)
	SetLockerTotalRewardsByAssetAppWise(ctx sdktypes.Context, lockerRewardsMapping lockertypes.LockerTotalRewardsByAssetAppWise) error
	GetLockerTotalRewardsByAssetAppWise(ctx sdktypes.Context, appID, assetID uint64) (lockerRewardsMapping lockertypes.LockerTotalRewardsByAssetAppWise, found bool)
}

type RewardsKeeper interface {
	GetReward(ctx sdktypes.Context, appID, assetID uint64) (rewards rewardstypes.InternalRewards, found bool)
	CalculationOfRewards(ctx sdktypes.Context, amount sdktypes.Int, lsr sdktypes.Dec, bTime int64) (sdktypes.Dec, error)
	SetLockerRewardTracker(ctx sdktypes.Context, rewards rewardstypes.LockerRewardsTracker)
	GetLockerRewardTracker(ctx sdktypes.Context, id, appID uint64) (rewards rewardstypes.LockerRewardsTracker, found bool)
}

type EsmKeeper interface {
	GetParams(ctx sdktypes.Context) esmtypes.Params
}
