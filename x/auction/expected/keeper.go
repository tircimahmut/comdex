package expected

import (
	assettypes "github.com/comdex-official/comdex/x/asset/types"
	auctiontypes "github.com/comdex-official/comdex/x/auction/types"
	"github.com/comdex-official/comdex/x/collector/types"
	esmtypes "github.com/comdex-official/comdex/x/esm/types"
	lendtypes "github.com/comdex-official/comdex/x/lend/types"
	liquidationtypes "github.com/comdex-official/comdex/x/liquidation/types"
	markettypes "github.com/comdex-official/comdex/x/market/types"
	vaulttypes "github.com/comdex-official/comdex/x/vault/types"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type AccountKeeper interface {
	GetModuleAccount(ctx sdktypes.Context, name string) authtypes.ModuleAccountI
	GetModuleAddress(name string) sdktypes.AccAddress
}

type BankKeeper interface {
	MintCoins(ctx sdktypes.Context, name string, coins sdktypes.Coins) error
	BurnCoins(ctx sdktypes.Context, name string, coins sdktypes.Coins) error
	SendCoinsFromModuleToModule(ctx sdktypes.Context, senderModule string, recipientModule string, amt sdktypes.Coins) error
	SendCoinsFromModuleToAccount(ctx sdktypes.Context, senderModule string, recipientAddr sdktypes.AccAddress, amt sdktypes.Coins) error
	SendCoinsFromAccountToModule(ctx sdktypes.Context, senderAddr sdktypes.AccAddress, recipientModule string, amt sdktypes.Coins) error
	GetBalance(ctx sdktypes.Context, addr sdktypes.AccAddress, denom string) sdktypes.Coin
}

type MarketKeeper interface {
	CalcAssetPrice(ctx sdktypes.Context, id uint64, amt sdktypes.Int) (price sdktypes.Dec, err error)
	GetTwa(ctx sdktypes.Context, id uint64) (twa markettypes.TimeWeightedAverage, found bool)
}

type LiquidationKeeper interface {
	SetFlagIsAuctionInProgress(ctx sdktypes.Context, appID, id uint64, flag bool) error
	SetFlagIsAuctionComplete(ctx sdktypes.Context, appID, id uint64, flag bool) error
	GetLockedVaults(ctx sdktypes.Context) (lockedVaults []liquidationtypes.LockedVault)
	GetLockedVault(ctx sdktypes.Context, appID, id uint64) (lockedVault liquidationtypes.LockedVault, found bool)
	SetLockedVault(ctx sdktypes.Context, lockedVault liquidationtypes.LockedVault)
	DeleteLockedVault(ctx sdktypes.Context, appID, id uint64)
	CreateLockedVaultHistory(ctx sdktypes.Context, lockedVault liquidationtypes.LockedVault) error
	UnLiquidateLockedBorrows(ctx sdktypes.Context, appID, id uint64, dutchAuction auctiontypes.DutchAuction) error
}

type AssetKeeper interface {
	GetAsset(ctx sdktypes.Context, id uint64) (assettypes.Asset, bool)
	GetPair(ctx sdktypes.Context, id uint64) (assettypes.Pair, bool)
	GetApps(ctx sdktypes.Context) (apps []assettypes.AppData, found bool)
	GetApp(ctx sdktypes.Context, id uint64) (app assettypes.AppData, found bool)
	GetPairsVault(ctx sdktypes.Context, id uint64) (pairs assettypes.ExtendedPairVault, found bool)
}

type VaultKeeper interface {
	GetAppExtendedPairVaultMappingData(ctx sdktypes.Context, appMappingID uint64, pairVaultsID uint64) (appExtendedPairVaultData vaulttypes.AppExtendedPairVaultMappingData, found bool)
	SetAppExtendedPairVaultMappingData(ctx sdktypes.Context, appExtendedPairVaultData vaulttypes.AppExtendedPairVaultMappingData)
	UpdateTokenMintedAmountLockerMapping(ctx sdktypes.Context, appMappingID uint64, extendedPairID uint64, amount sdktypes.Int, changeType bool)
	UpdateCollateralLockedAmountLockerMapping(ctx sdktypes.Context, appMappingID uint64, extendedPairID uint64, amount sdktypes.Int, changeType bool)
	DeleteUserVaultExtendedPairMapping(ctx sdktypes.Context, from string, appMapping uint64, extendedPairVault uint64)
	CreateNewVault(ctx sdktypes.Context, From string, AppID uint64, ExtendedPairVaultID uint64, AmountIn sdktypes.Int, AmountOut sdktypes.Int) error
	GetUserAppExtendedPairMappingData(ctx sdktypes.Context, from string, appMapping uint64, extendedPairVault uint64) (userVaultAssetData vaulttypes.OwnerAppExtendedPairVaultMappingData, found bool)
	GetUserAppMappingData(ctx sdktypes.Context, from string, appMapping uint64) (userVaultAssetData []vaulttypes.OwnerAppExtendedPairVaultMappingData, found bool)
	// CheckUserAppToExtendedPairMapping(ctx sdktypes.Context, userVaultAssetData vaulttypes.UserVaultAssetMapping, extendedPairVaultID uint64, appMappingID uint64) (vaultID uint64, found bool)
	SetVault(ctx sdktypes.Context, vault vaulttypes.Vault)
	GetVault(ctx sdktypes.Context, id uint64) (vault vaulttypes.Vault, found bool)
	GetAmountOfOtherToken(ctx sdktypes.Context, id1 uint64, rate1 sdktypes.Dec, amt1 sdktypes.Int, id2 uint64, rate2 sdktypes.Dec) (sdktypes.Dec, sdktypes.Int, error)
	GetLengthOfVault(ctx sdktypes.Context) uint64
	SetLengthOfVault(ctx sdktypes.Context, length uint64)
}

type CollectorKeeper interface {
	GetAppidToAssetCollectorMapping(ctx sdktypes.Context, appID, assetID uint64) (appAssetCollectorData types.AppToAssetIdCollectorMapping, found bool)
	UpdateCollector(ctx sdktypes.Context, appID, assetID uint64, CollectedStabilityFee, CollectedClosingFee, CollectedOpeningFee, LiquidationRewardsCollected sdktypes.Int) error
	// SetCollectorLookupTable(ctx sdktypes.Context, records ...types.CollectorLookupTable) error
	GetCollectorLookupTable(ctx sdktypes.Context, appID, assetID uint64) (collectorLookup types.CollectorLookupTableData, found bool)
	GetAuctionMappingForApp(ctx sdktypes.Context, appID, assetID uint64) (collectorAuctionLookupTable types.AppAssetIdToAuctionLookupTable, found bool)
	GetNetFeeCollectedData(ctx sdktypes.Context, appID, assetID uint64) (netFeeData types.AppAssetIdToFeeCollectedData, found bool)
	GetAmountFromCollector(ctx sdktypes.Context, appID, assetID uint64, amount sdktypes.Int) (sdktypes.Int, error)
	SetNetFeeCollectedData(ctx sdktypes.Context, appID, assetID uint64, fee sdktypes.Int) error
	SetAuctionMappingForApp(ctx sdktypes.Context, records types.AppAssetIdToAuctionLookupTable) error
	GetAllAuctionMappingForApp(ctx sdktypes.Context) (collectorAuctionLookupTable []types.AppAssetIdToAuctionLookupTable, found bool)
}

type TokenMintKeeper interface {
	MintNewTokensForApp(ctx sdktypes.Context, appMappingID uint64, assetID uint64, address string, amount sdktypes.Int) error
	BurnTokensForApp(ctx sdktypes.Context, appMappingID uint64, assetID uint64, amount sdktypes.Int) error
}

type EsmKeeper interface {
	GetKillSwitchData(ctx sdktypes.Context, appID uint64) (esmtypes.KillSwitchParams, bool)
	GetESMStatus(ctx sdktypes.Context, id uint64) (esmStatus esmtypes.ESMStatus, found bool)
	CalcDollarValueOfToken(ctx sdktypes.Context, rate uint64, amt sdktypes.Int, decimals sdktypes.Int) (price sdktypes.Dec)
	SetAssetToAmount(ctx sdktypes.Context, assetToAmount esmtypes.AssetToAmount)
	GetDataAfterCoolOff(ctx sdktypes.Context, id uint64) (esmDataAfterCoolOff esmtypes.DataAfterCoolOff, found bool)
	SetDataAfterCoolOff(ctx sdktypes.Context, esmDataAfterCoolOff esmtypes.DataAfterCoolOff)
	GetSnapshotOfPrices(ctx sdktypes.Context, appID, assetID uint64) (price uint64, found bool)
}

type LendKeeper interface {
	GetBorrow(ctx sdktypes.Context, id uint64) (borrow lendtypes.BorrowAsset, found bool)
	GetLendPair(ctx sdktypes.Context, id uint64) (pair lendtypes.Extended_Pair, found bool)
	GetAssetRatesParams(ctx sdktypes.Context, assetID uint64) (assetRatesStats lendtypes.AssetRatesParams, found bool)
	VerifyCollateralizationRatio(ctx sdktypes.Context, amountIn sdktypes.Int, assetIn assettypes.Asset, amountOut sdktypes.Int, assetOut assettypes.Asset, liquidationThreshold sdktypes.Dec) error
	CalculateCollateralizationRatio(ctx sdktypes.Context, amountIn sdktypes.Int, assetIn assettypes.Asset, amountOut sdktypes.Int, assetOut assettypes.Asset) (sdktypes.Dec, error)
	GetLend(ctx sdktypes.Context, id uint64) (lend lendtypes.LendAsset, found bool)
	GetPool(ctx sdktypes.Context, id uint64) (pool lendtypes.Pool, found bool)
	GetAddAuctionParamsData(ctx sdktypes.Context, appID uint64) (auctionParams lendtypes.AuctionParams, found bool)
	ModuleBalance(ctx sdktypes.Context, moduleName string, denom string) sdktypes.Int
	UpdateReserveBalances(ctx sdktypes.Context, assetID uint64, moduleName string, payment sdktypes.Coin, inc bool) error
	SetLend(ctx sdktypes.Context, lend lendtypes.LendAsset)
	SetAllReserveStatsByAssetID(ctx sdktypes.Context, allReserveStats lendtypes.AllReserveStats)
	GetAllReserveStatsByAssetID(ctx sdktypes.Context, id uint64) (allReserveStats lendtypes.AllReserveStats, found bool)
}
