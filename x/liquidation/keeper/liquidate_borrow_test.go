package keeper_test

import (
	assetTypes "github.com/comdex-official/comdex/x/asset/types"
	lendkeeper "github.com/comdex-official/comdex/x/lend/keeper"
	lendtypes "github.com/comdex-official/comdex/x/lend/types"
	liquidationTypes "github.com/comdex-official/comdex/x/liquidation/types"
	markettypes "github.com/comdex-official/comdex/x/market/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func Dec(s string) sdk.Dec {
	dec, err := sdk.NewDecFromStr(s)
	if err != nil {
		panic(err)
	}
	return dec
}

func (s *KeeperTestSuite) AddAppAssetLend() {
	assetKeeper, ctx := &s.assetKeeper, &s.ctx
	lendKeeper := &s.lendKeeper
	msg1 := assetTypes.AppData{
		Name:             "cswap",
		ShortName:        "cswap",
		MinGovDeposit:    sdk.NewIntFromUint64(10000000),
		GovTimeInSeconds: 900,
	}
	err := assetKeeper.AddAppRecords(*ctx, msg1)
	s.Require().NoError(err)

	msg2 := assetTypes.AppData{
		Name:             "harbor",
		ShortName:        "harbor",
		MinGovDeposit:    sdk.NewIntFromUint64(10000000),
		GovTimeInSeconds: 900,
	}
	err = assetKeeper.AddAppRecords(*ctx, msg2)
	s.Require().NoError(err)
	msg3 := assetTypes.AppData{
		Name:             "commodo",
		ShortName:        "comdo",
		MinGovDeposit:    sdk.NewIntFromUint64(10000000),
		GovTimeInSeconds: 900,
	}
	err = assetKeeper.AddAppRecords(*ctx, msg3)
	s.Require().NoError(err)

	msg4 := assetTypes.Asset{
		Name:          "ATOM",
		Denom:         "uatom",
		Decimals:      sdk.NewInt(1000000),
		IsOnChain:     true,
		IsCdpMintable: true,
	}

	err = assetKeeper.AddAssetRecords(*ctx, msg4)
	s.Require().NoError(err)
	market1 := markettypes.TimeWeightedAverage{
		AssetID:       1,
		ScriptID:      12,
		Twa:           1000000,
		CurrentIndex:  0,
		IsPriceActive: true,
		PriceValue:    []uint64{1000000},
	}
	s.app.MarketKeeper.SetTwa(s.ctx, market1)
	_, err = s.app.MarketKeeper.GetLatestPrice(s.ctx, 1)
	s.Suite.NoError(err)

	msg5 := assetTypes.Asset{
		Name:          "CMDX",
		Denom:         "ucmdx",
		Decimals:      sdk.NewInt(1000000),
		IsOnChain:     true,
		IsCdpMintable: true,
	}

	err = assetKeeper.AddAssetRecords(*ctx, msg5)
	s.Require().NoError(err)
	market2 := markettypes.TimeWeightedAverage{
		AssetID:       2,
		ScriptID:      12,
		Twa:           1000000,
		CurrentIndex:  0,
		IsPriceActive: true,
		PriceValue:    []uint64{1000000},
	}
	s.app.MarketKeeper.SetTwa(s.ctx, market2)
	_, err = s.app.MarketKeeper.GetLatestPrice(s.ctx, 2)
	s.Suite.NoError(err)

	msg6 := assetTypes.Asset{
		Name:          "CMST",
		Denom:         "ucmst",
		Decimals:      sdk.NewInt(1000000),
		IsOnChain:     true,
		IsCdpMintable: true,
	}
	err = assetKeeper.AddAssetRecords(*ctx, msg6)
	s.Require().NoError(err)

	market3 := markettypes.TimeWeightedAverage{
		AssetID:       3,
		ScriptID:      12,
		Twa:           1000000,
		CurrentIndex:  0,
		IsPriceActive: true,
		PriceValue:    []uint64{1000000},
	}
	s.app.MarketKeeper.SetTwa(s.ctx, market3)
	_, err = s.app.MarketKeeper.GetLatestPrice(s.ctx, 3)
	s.Suite.NoError(err)

	msg7 := assetTypes.Asset{
		Name:      "HARBOR",
		Denom:     "uharbor",
		Decimals:  sdk.NewInt(1000000),
		IsOnChain: true,
	}
	err = assetKeeper.AddAssetRecords(*ctx, msg7)
	s.Require().NoError(err)

	market4 := markettypes.TimeWeightedAverage{
		AssetID:       4,
		ScriptID:      12,
		Twa:           1000000,
		CurrentIndex:  0,
		IsPriceActive: true,
		PriceValue:    []uint64{1000000},
	}
	s.app.MarketKeeper.SetTwa(s.ctx, market4)
	_, err = s.app.MarketKeeper.GetLatestPrice(s.ctx, 4)
	s.Suite.NoError(err)

	msg11 := assetTypes.Asset{
		Name:      "CATOM",
		Denom:     "ucatom",
		Decimals:  sdk.NewInt(1000000),
		IsOnChain: true,
	}
	err = assetKeeper.AddAssetRecords(*ctx, msg11)
	s.Require().NoError(err)

	msg12 := assetTypes.Asset{
		Name:      "CCMDX",
		Denom:     "uccmdx",
		Decimals:  sdk.NewInt(1000000),
		IsOnChain: true,
	}
	err = assetKeeper.AddAssetRecords(*ctx, msg12)
	s.Require().NoError(err)

	msg13 := assetTypes.Asset{
		Name:      "CCMST",
		Denom:     "uccmst",
		Decimals:  sdk.NewInt(1000000),
		IsOnChain: true,
	}
	err = assetKeeper.AddAssetRecords(*ctx, msg13)
	s.Require().NoError(err)

	cmstRatesParams := lendtypes.AssetRatesParams{
		AssetID:              3,
		UOptimal:             Dec("0.8"),
		Base:                 Dec("0.002"),
		Slope1:               Dec("0.06"),
		Slope2:               Dec("0.6"),
		EnableStableBorrow:   false,
		StableBase:           Dec("0.0"),
		StableSlope1:         Dec("0.0"),
		StableSlope2:         Dec("0.0"),
		Ltv:                  Dec("0.8"),
		LiquidationThreshold: Dec("0.85"),
		LiquidationPenalty:   Dec("0.025"),
		LiquidationBonus:     Dec("0.025"),
		ReserveFactor:        Dec("0.1"),
		CAssetID:             7,
	}
	lendKeeper.SetAssetRatesParams(s.ctx, cmstRatesParams)
	atomRatesParams := lendtypes.AssetRatesParams{
		AssetID:              1,
		UOptimal:             Dec("0.75"),
		Base:                 Dec("0.002"),
		Slope1:               Dec("0.07"),
		Slope2:               Dec("1.25"),
		EnableStableBorrow:   false,
		StableBase:           Dec("0.0"),
		StableSlope1:         Dec("0.0"),
		StableSlope2:         Dec("0.0"),
		Ltv:                  Dec("0.7"),
		LiquidationThreshold: Dec("0.75"),
		LiquidationPenalty:   Dec("0.05"),
		LiquidationBonus:     Dec("0.05"),
		ReserveFactor:        Dec("0.2"),
		CAssetID:             5,
	}
	lendKeeper.SetAssetRatesParams(s.ctx, atomRatesParams)

	cmdxRatesParams := lendtypes.AssetRatesParams{
		AssetID:              2,
		UOptimal:             Dec("0.5"),
		Base:                 Dec("0.002"),
		Slope1:               Dec("0.08"),
		Slope2:               Dec("2.0"),
		EnableStableBorrow:   false,
		StableBase:           Dec("0.0"),
		StableSlope1:         Dec("0.0"),
		StableSlope2:         Dec("0.0"),
		Ltv:                  Dec("0.5"),
		LiquidationThreshold: Dec("0.55"),
		LiquidationPenalty:   Dec("0.05"),
		LiquidationBonus:     Dec("0.05"),
		ReserveFactor:        Dec("0.2"),
		CAssetID:             6,
	}
	lendKeeper.SetAssetRatesParams(s.ctx, cmdxRatesParams)

	var (
		assetDataCMDXPool []*lendtypes.AssetDataPoolMapping
	)
	assetDataPoolOneAssetOne := &lendtypes.AssetDataPoolMapping{
		AssetID:          1,
		AssetTransitType: 3,
		SupplyCap:        sdk.NewDec(5000000000000),
	}
	assetDataPoolOneAssetTwo := &lendtypes.AssetDataPoolMapping{
		AssetID:          2,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(1000000000000),
	}
	assetDataPoolOneAssetThree := &lendtypes.AssetDataPoolMapping{
		AssetID:          3,
		AssetTransitType: 2,
		SupplyCap:        sdk.NewDec(5000000000000),
	}

	assetDataCMDXPool = append(assetDataCMDXPool, assetDataPoolOneAssetOne, assetDataPoolOneAssetTwo, assetDataPoolOneAssetThree)
	cmdxPool := lendtypes.Pool{
		ModuleName: "cmdx",
		CPoolName:  "CMDX-ATOM-CMST",
		AssetData:  assetDataCMDXPool,
	}
	err = lendKeeper.AddPoolRecords(s.ctx, cmdxPool)
	if err != nil {
		panic(err)
	}

	cmdxcmstPair := lendtypes.Extended_Pair{ // 1
		AssetIn:         2,
		AssetOut:        3,
		IsInterPool:     false,
		AssetOutPoolID:  1,
		MinUsdValueLeft: 100000,
	}
	err = lendKeeper.AddLendPairsRecords(s.ctx, cmdxcmstPair)
	if err != nil {
		panic(err)
	}
	cmdxatomPair := lendtypes.Extended_Pair{ // 2
		AssetIn:         2,
		AssetOut:        1,
		IsInterPool:     false,
		AssetOutPoolID:  1,
		MinUsdValueLeft: 100000,
	}
	err = lendKeeper.AddLendPairsRecords(s.ctx, cmdxatomPair)
	if err != nil {
		panic(err)
	}
	atomcmdxPair := lendtypes.Extended_Pair{ // 3
		AssetIn:         1,
		AssetOut:        2,
		IsInterPool:     false,
		AssetOutPoolID:  1,
		MinUsdValueLeft: 100000,
	}
	err = lendKeeper.AddLendPairsRecords(s.ctx, atomcmdxPair)
	if err != nil {
		panic(err)
	}
	atomcmstPair := lendtypes.Extended_Pair{ // 4
		AssetIn:         1,
		AssetOut:        3,
		IsInterPool:     false,
		AssetOutPoolID:  1,
		MinUsdValueLeft: 100000,
	}
	err = lendKeeper.AddLendPairsRecords(s.ctx, atomcmstPair)
	if err != nil {
		panic(err)
	}
	cmstcmdxPair := lendtypes.Extended_Pair{ // 5
		AssetIn:         3,
		AssetOut:        2,
		IsInterPool:     false,
		AssetOutPoolID:  1,
		MinUsdValueLeft: 100000,
	}
	err = lendKeeper.AddLendPairsRecords(s.ctx, cmstcmdxPair)
	if err != nil {
		panic(err)
	}
	cmstatomPair := lendtypes.Extended_Pair{ // 6
		AssetIn:         3,
		AssetOut:        1,
		IsInterPool:     false,
		AssetOutPoolID:  1,
		MinUsdValueLeft: 100000,
	}
	err = lendKeeper.AddLendPairsRecords(s.ctx, cmstatomPair)
	if err != nil {
		panic(err)
	}

	// Adding Lend Pair Mapping
	map1 := lendtypes.AssetToPairMapping{
		PoolID:  1,
		AssetID: 1,
		PairID:  []uint64{3, 4},
	}
	lendKeeper.SetAssetToPair(s.ctx, map1)
	map2 := lendtypes.AssetToPairMapping{
		PoolID:  1,
		AssetID: 2,
		PairID:  []uint64{1, 2},
	}
	lendKeeper.SetAssetToPair(s.ctx, map2)
	map3 := lendtypes.AssetToPairMapping{
		PoolID:  1,
		AssetID: 3,
		PairID:  []uint64{5, 6},
	}
	lendKeeper.SetAssetToPair(s.ctx, map3)

	auctionParams := lendtypes.AuctionParams{
		AppId:                  3,
		AuctionDurationSeconds: 21600,
		Buffer:                 Dec("1.2"),
		Cusp:                   Dec("0.7"),
		Step:                   sdk.NewInt(360),
		PriceFunctionType:      1,
		DutchId:                3,
		BidDurationSeconds:     3600,
	}
	err = lendKeeper.AddAuctionParamsData(s.ctx, auctionParams)
	if err != nil {
		return
	}

	userAddr := "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"
	userAddress, _ := sdk.AccAddressFromBech32(userAddr)
	s.fundAddr(userAddress, sdk.NewCoin("uatom", sdk.NewIntFromUint64(100000000000000)))
	err = lendKeeper.FundModAcc(s.ctx, 1, 1, userAddr, sdk.NewCoin("uatom", sdk.NewInt(100000000000)))
	s.Require().NoError(err)
	s.fundAddr(userAddress, sdk.NewCoin("ucmdx", sdk.NewIntFromUint64(100000000000000)))

	err = lendKeeper.FundModAcc(s.ctx, 1, 2, userAddr, sdk.NewCoin("ucmdx", sdk.NewInt(1000000000000)))
	s.Require().NoError(err)
	s.fundAddr(userAddress, sdk.NewCoin("ucmst", sdk.NewIntFromUint64(100000000000000)))

	err = lendKeeper.FundModAcc(s.ctx, 1, 3, userAddr, sdk.NewCoin("ucmst", sdk.NewInt(100000000000)))
	s.Require().NoError(err)

	lendKeeper, ctx = &s.lendKeeper, &s.ctx
	server := lendkeeper.NewMsgServerImpl(*lendKeeper)

	msg20 := lendtypes.MsgBorrowAlternate{
		Lender:         "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
		AssetId:        1,
		PoolId:         1,
		AmountIn:       sdk.NewCoin("uatom", sdk.NewInt(100000000000)),
		PairId:         3,
		IsStableBorrow: false,
		AmountOut:      sdk.NewCoin("ucmdx", sdk.NewInt(70000000000)),
		AppId:          3,
	}

	s.fundAddr(userAddress, sdk.NewCoin("uatom", sdk.NewIntFromUint64(1000000000000)))
	_, err = server.BorrowAlternate(sdk.WrapSDKContext(*ctx), &msg20)
	s.Require().NoError(err)
	msg30 := lendtypes.MsgBorrowAlternate{
		Lender:         "cosmos1kwtdrjkwu6y87vlylaeatzmc5p4jhvn7qwqnkp",
		AssetId:        1,
		PoolId:         1,
		AmountIn:       sdk.NewCoin("uatom", sdk.NewInt(100000000000)),
		PairId:         3,
		IsStableBorrow: false,
		AmountOut:      sdk.NewCoin("ucmdx", sdk.NewInt(70000000000)),
		AppId:          3,
	}
	userAddr2, _ := sdk.AccAddressFromBech32("cosmos1kwtdrjkwu6y87vlylaeatzmc5p4jhvn7qwqnkp")
	s.fundAddr(userAddr2, sdk.NewCoin("uatom", sdk.NewIntFromUint64(1000000000000)))
	_, err = server.BorrowAlternate(sdk.WrapSDKContext(*ctx), &msg30)
	s.Require().NoError(err)
}

func (s *KeeperTestSuite) GetBorrowCount() int {
	ctx := &s.ctx
	res, err := s.lendQuerier.QueryBorrows(sdk.WrapSDKContext(*ctx), &lendtypes.QueryBorrowsRequest{})
	s.Require().NoError(err)
	return len(res.Borrows)
}

func (s *KeeperTestSuite) GetBorrowCountByPoolIDAssetID(poolID, assetID uint64) int {
	lendKeeper, ctx := &s.lendKeeper, &s.ctx
	res, found := lendKeeper.GetAssetStatsByPoolIDAndAssetID(*ctx, poolID, assetID)
	s.Require().True(found)
	return len(res.BorrowIds)
}

func (s *KeeperTestSuite) ChangeOraclePriceLend(asset, price uint64) {
	market1 := markettypes.TimeWeightedAverage{
		AssetID:       asset,
		ScriptID:      12,
		Twa:           price,
		CurrentIndex:  0,
		IsPriceActive: true,
		PriceValue:    []uint64{price},
	}
	s.app.MarketKeeper.SetTwa(s.ctx, market1)
	_, err := s.app.MarketKeeper.GetLatestPrice(s.ctx, asset)
	s.Suite.NoError(err)
}

func (s *KeeperTestSuite) TestLiquidateBorrows() {
	liquidationKeeper, ctx := &s.liquidationKeeper, &s.ctx
	s.AddAppAssetLend()
	currentBorrowsCount := 2
	s.Require().Equal(s.GetBorrowCount(), currentBorrowsCount)
	s.Require().Equal(s.GetBorrowCountByPoolIDAssetID(1, 2), currentBorrowsCount)
	beforeBorrow, found := s.lendKeeper.GetBorrow(*ctx, 1)
	s.Require().True(found)

	// Liquidation shouldn't happen as price not changed
	err := liquidationKeeper.LiquidateBorrows(*ctx)
	s.Require().NoError(err)
	id := liquidationKeeper.GetLockedVaultID(*ctx)
	s.Require().Equal(id, uint64(0))

	// Liquidation should happen as price changed
	beforeAmtIn := beforeBorrow.AmountIn.Amount

	s.ChangeOraclePriceLend(2, 1200000)
	err = liquidationKeeper.LiquidateBorrows(*ctx)
	s.Require().NoError(err)
	id = liquidationKeeper.GetLockedVaultID(*ctx)
	s.Require().Equal(id, uint64(2))

	lockedVault := liquidationKeeper.GetLockedVaults(*ctx)
	s.Require().Equal(lockedVault[0].OriginalVaultId, beforeBorrow.ID)
	s.Require().Equal(lockedVault[0].ExtendedPairId, beforeBorrow.PairID)
	//s.Require().Equal(lockedVault[0].AmountIn, beforeBorrow.AmountIn.Amount)
	s.Require().Equal(lockedVault[0].AmountOut, beforeBorrow.AmountOut.Amount)
	s.Require().Equal(lockedVault[0].UpdatedAmountOut, beforeBorrow.AmountOut.Amount.Add(beforeBorrow.InterestAccumulated.TruncateInt()))
	s.Require().Equal(lockedVault[0].Initiator, liquidationTypes.ModuleName)
	s.Require().Equal(lockedVault[0].IsAuctionInProgress, true)
	s.Require().Equal(lockedVault[0].IsAuctionComplete, false)
	s.Require().Equal(lockedVault[0].SellOffHistory, []string(nil))
	price, err := s.app.MarketKeeper.CalcAssetPrice(*ctx, uint64(1), beforeAmtIn.Sub(lockedVault[0].AmountIn))
	s.Require().NoError(err)
	updatedPrice := price.Sub(price.Mul(Dec("0.09090909090")))

	s.Require().Equal(lockedVault[0].CollateralToBeAuctioned.TruncateInt(), updatedPrice.TruncateInt())
	s.Require().Equal(lockedVault[0].CrAtLiquidation, lockedVault[0].AmountOut.ToDec().Mul(s.GetAssetPrice(2)).Quo(beforeAmtIn.ToDec().Mul(s.GetAssetPrice(1))))
}
