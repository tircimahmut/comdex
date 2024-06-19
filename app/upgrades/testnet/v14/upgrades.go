package v14

import (
	"fmt"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	bandoraclemodulekeeper "github.com/comdex-official/comdex/x/bandoracle/keeper"
	lendkeeper "github.com/comdex-official/comdex/x/lend/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	auctionkeeperskip "github.com/skip-mev/block-sdk/x/auction/keeper"
	auctionmoduleskiptypes "github.com/skip-mev/block-sdk/x/auction/types"
	"strings"
)

func CreateUpgradeHandlerV14(
	mm *module.Manager,
	configurator module.Configurator,
	auctionkeeperskip auctionkeeperskip.Keeper,
	lendKeeper lendkeeper.Keeper,
	wasmKeeper wasmkeeper.Keeper,
	StakingKeeper stakingkeeper.Keeper,
	MintKeeper mintkeeper.Keeper,
	SlashingKeeper slashingkeeper.Keeper,
	bandoracleKeeper bandoraclemodulekeeper.Keeper,
	accountKeeper authkeeper.AccountKeeper,

) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {

		ctx.Logger().Info("Applying test net upgrade - v14.0.0")
		logger := ctx.Logger().With("upgrade", UpgradeName)
		vm, err := mm.RunMigrations(ctx, configurator, fromVM)
		if err != nil {
			return vm, err
		}

		ctx.Logger().Info("setting default params for MEV module (x/auction)")
		if err = setDefaultMEVParams(ctx, auctionkeeperskip); err != nil {
			return nil, err
		}

		// x/Mint
		// from 3 seconds to 6 = 1/2x blocks per year
		mintParams := MintKeeper.GetParams(ctx)
		mintParams.BlocksPerYear /= 2
		if err = MintKeeper.SetParams(ctx, mintParams); err != nil {
			return nil, err
		}
		logger.Info(fmt.Sprintf("updated minted blocks per year logic to %v", mintParams))

		// x/Slashing
		slashingParams := SlashingKeeper.GetParams(ctx)
		slashingParams.SignedBlocksWindow /= 2
		if err := SlashingKeeper.SetParams(ctx, slashingParams); err != nil {
			return nil, err
		}
		logger.Info(fmt.Sprintf("updated slashing params to %v", slashingParams))

		// update wasm to permission nobody
		wasmParams := wasmKeeper.GetParams(ctx)
		wasmParams.CodeUploadAccess = wasmtypes.AllowNobody
		wasmKeeper.SetParams(ctx, wasmParams)
		logger.Info(fmt.Sprintf("updated wasm params to %v", wasmParams))

		// update discard BH of oracle
		bandData := bandoracleKeeper.GetFetchPriceMsg(ctx)
		if bandData.Size() > 0 {
			bandData.AcceptedHeightDiff = 3000
			bandoracleKeeper.SetFetchPriceMsg(ctx, bandData)
			logger.Info(fmt.Sprintf("updated bandData to %v", bandData))
		}

		// update tx size cost per byte
		authParams := accountKeeper.GetParams(ctx)
		authParams.TxSizeCostPerByte = authParams.TxSizeCostPerByte * 2
		if err = accountKeeper.SetParams(ctx, authParams); err != nil {
			return nil, err
		}
		logger.Info(fmt.Sprintf("updated auth params to %v", accountKeeper.GetParams(ctx)))

		//TODO: uncomment this before mainnet upgrade
		//UpdateLendParams(ctx, lendKeeper)
		return vm, err
	}
}

func setDefaultMEVParams(ctx sdk.Context, auctionkeeperskip auctionkeeperskip.Keeper) error {
	nativeDenom := getChainBondDenom(ctx.ChainID())

	// Skip MEV (x/auction)
	return auctionkeeperskip.SetParams(ctx, auctionmoduleskiptypes.Params{
		MaxBundleSize:          auctionmoduleskiptypes.DefaultMaxBundleSize,
		EscrowAccountAddress:   authtypes.NewModuleAddress(auctionmoduleskiptypes.ModuleName), // TODO: revisit
		ReserveFee:             sdk.NewCoin(nativeDenom, sdk.NewInt(10)),
		MinBidIncrement:        sdk.NewCoin(nativeDenom, sdk.NewInt(5)),
		FrontRunningProtection: auctionmoduleskiptypes.DefaultFrontRunningProtection,
		ProposerFee:            auctionmoduleskiptypes.DefaultProposerFee,
	})
}

// getChainBondDenom returns expected bond denom based on chainID.
func getChainBondDenom(chainID string) string {
	if strings.HasPrefix(chainID, "comdex-") {
		return "ucmdx"
	}
	return "stake"
}

func UpdateLendParams(
	ctx sdk.Context,
	lendKeeper lendkeeper.Keeper,
) {
	assetRatesParamsStAtom, _ := lendKeeper.GetAssetRatesParams(ctx, 14)
	assetRatesParamsStAtom.CAssetID = 23
	lendKeeper.SetAssetRatesParams(ctx, assetRatesParamsStAtom)
}
