package keeper_test

import (
	"fmt"
	"time"

	errorsmod "cosmossdk.io/errors"
	utils "github.com/comdex-official/comdex/types"
	"github.com/comdex-official/comdex/x/lend/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/pkg/errors"
)

func (s *KeeperTestSuite) TestMsgLend() {
	assetOneID := s.CreateNewAsset("ASSETONE", "uasset1", 1000000)
	assetTwoID := s.CreateNewAsset("ASSETTWO", "uasset2", 2000000)
	assetThreeID := s.CreateNewAsset("ASSETTHREE", "uasset3", 2000000)
	assetFourID := s.CreateNewAsset("ASSETFOUR", "uasset4", 2000000)
	cAssetOneID := s.CreateNewAsset("CASSETONE", "ucasset1", 1000000)
	cAssetTwoID := s.CreateNewAsset("CASSETTWO", "ucasset2", 2000000)
	cAssetThreeID := s.CreateNewAsset("CASSETTHRE", "ucasset3", 2000000)
	// cAssetFourID := s.CreateNewAsset("CASSETFOUR", "ucasset4", 2000000)

	var (
		assetDataPoolOne []*types.AssetDataPoolMapping
		assetDataPoolTwo []*types.AssetDataPoolMapping
	)
	assetDataPoolOneAssetOne := &types.AssetDataPoolMapping{
		AssetID:          assetOneID,
		AssetTransitType: 3,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolOneAssetTwo := &types.AssetDataPoolMapping{
		AssetID:          assetTwoID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(1000000000000000000),
	}
	assetDataPoolOneAssetThree := &types.AssetDataPoolMapping{
		AssetID:          assetThreeID,
		AssetTransitType: 2,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolTwoAssetFour := &types.AssetDataPoolMapping{
		AssetID:          assetFourID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(3000000000000000000),
	}
	assetDataPoolOne = append(assetDataPoolOne, assetDataPoolOneAssetOne, assetDataPoolOneAssetTwo, assetDataPoolOneAssetThree)
	assetDataPoolTwo = append(assetDataPoolTwo, assetDataPoolTwoAssetFour, assetDataPoolOneAssetOne, assetDataPoolOneAssetThree)

	poolOneID := s.CreateNewPool("cmdx", "CMDX-ATOM-CMST", assetDataPoolOne)
	poolTwoID := s.CreateNewPool("osmo", "OSMO-ATOM-CMST", assetDataPoolTwo)

	s.AddAssetRatesStats(assetThreeID, newDec("0.8"), newDec("0.002"), newDec("0.06"), newDec("0.6"), true, newDec("0.04"), newDec("0.04"), newDec("0.06"), newDec("0.8"), newDec("0.85"), newDec("0.025"), newDec("0.025"), newDec("0.1"), cAssetThreeID)
	s.AddAssetRatesStats(assetOneID, newDec("0.75"), newDec("0.002"), newDec("0.07"), newDec("1.25"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.7"), newDec("0.75"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetOneID)
	// s.AddAssetRatesStats(assetFourID, newDec("0.65"), newDec("0.002"), newDec("0.08"), newDec("1.5"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.6"), newDec("0.65"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetFourID)
	s.AddAssetRatesStats(assetTwoID, newDec("0.5"), newDec("0.002"), newDec("0.08"), newDec("2.0"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.5"), newDec("0.55"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetTwoID)

	pairOneID := s.AddExtendedLendPair(assetTwoID, assetThreeID, false, poolOneID, 1000000)
	pairTwoID := s.AddExtendedLendPair(assetTwoID, assetOneID, false, poolOneID, 1000000)
	pairThreeID := s.AddExtendedLendPair(assetOneID, assetTwoID, false, poolOneID, 1000000)
	pairFourID := s.AddExtendedLendPair(assetOneID, assetThreeID, false, poolOneID, 1000000)
	pairFiveID := s.AddExtendedLendPair(assetThreeID, assetTwoID, false, poolOneID, 1000000)
	pairSixID := s.AddExtendedLendPair(assetThreeID, assetOneID, false, poolOneID, 1000000)
	pairSevenID := s.AddExtendedLendPair(assetFourID, assetThreeID, false, poolTwoID, 1000000)
	pairEightID := s.AddExtendedLendPair(assetFourID, assetOneID, false, poolTwoID, 1000000)
	pairNineID := s.AddExtendedLendPair(assetOneID, assetFourID, false, poolTwoID, 1000000)
	pairTenID := s.AddExtendedLendPair(assetOneID, assetThreeID, false, poolTwoID, 1000000)
	pairElevenID := s.AddExtendedLendPair(assetThreeID, assetFourID, false, poolTwoID, 1000000)
	pairTwelveID := s.AddExtendedLendPair(assetThreeID, assetOneID, false, poolTwoID, 1000000)
	pairThirteenID := s.AddExtendedLendPair(assetTwoID, assetFourID, true, poolTwoID, 1000000)
	pairFourteenID := s.AddExtendedLendPair(assetThreeID, assetFourID, true, poolTwoID, 1000000)
	pairFifteenID := s.AddExtendedLendPair(assetOneID, assetFourID, true, poolTwoID, 1000000)
	pairSixteenID := s.AddExtendedLendPair(assetFourID, assetTwoID, true, poolOneID, 1000000)
	pairSeventeenID := s.AddExtendedLendPair(assetThreeID, assetTwoID, true, poolOneID, 1000000)
	pairEighteenID := s.AddExtendedLendPair(assetOneID, assetTwoID, true, poolOneID, 1000000)

	s.AddAssetToPair(assetOneID, poolOneID, []uint64{pairThreeID, pairFourID, pairFifteenID})
	s.AddAssetToPair(assetTwoID, poolOneID, []uint64{pairOneID, pairTwoID, pairThirteenID})
	s.AddAssetToPair(assetThreeID, poolOneID, []uint64{pairFiveID, pairSixID, pairFourteenID})
	s.AddAssetToPair(assetFourID, poolTwoID, []uint64{pairSevenID, pairEightID, pairSixteenID})
	s.AddAssetToPair(assetOneID, poolTwoID, []uint64{pairNineID, pairTenID, pairEighteenID})
	s.AddAssetToPair(assetThreeID, poolTwoID, []uint64{pairElevenID, pairTwelveID, pairSeventeenID})

	appOneID := s.CreateNewApp("commodo", "cmmdo")
	appTwoID := s.CreateNewApp("cswap", "cswap")

	msg3 := types.NewMsgFundModuleAccounts(1, assetOneID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset1", newInt(10000000000)))
	msg4 := types.NewMsgFundModuleAccounts(1, assetTwoID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset2", newInt(10000000000)))
	msg5 := types.NewMsgFundModuleAccounts(1, assetThreeID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset3", newInt(120000000)))
	// msg6 := types.NewMsgFundModuleAccounts(2, assetThreeID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset3", newInt(10000000000)))
	msg7 := types.NewMsgFundModuleAccounts(2, assetOneID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset1", newInt(10000000000)))
	msg8 := types.NewMsgFundModuleAccounts(2, assetFourID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset4", newInt(10000000000)))

	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset1", newInt(100000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset2", newInt(100000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset3", newInt(100000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset4", newInt(100000000000))))

	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg3)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg4)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg5)
	//_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg6)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg7)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg8)

	msg := types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetOneID, sdk.NewCoin("uasset1", newInt(300)), poolOneID, appOneID)
	_, _ = s.msgServer.Lend(sdk.WrapSDKContext(s.ctx), msg)
	msg2 := types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, 3, false, sdk.NewCoin("ucasset1", newInt(100)), sdk.NewCoin("uasset2", newInt(10)))
	_, err := s.msgServer.Borrow(sdk.WrapSDKContext(s.ctx), msg2)
	fmt.Println("err", err)

	testCases := []struct {
		Name               string
		Msg                types.MsgLend
		ExpErr             error
		ExpResp            *types.MsgLendResponse
		QueryResponseIndex uint64
		QueryResponse      *types.LendAsset
		AvailableBalance   sdk.Coins
	}{
		{
			Name:               "asset does not exist",
			Msg:                *types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 10, sdk.NewCoin("uasset1", sdk.NewInt(100)), poolOneID, 1),
			ExpErr:             types.ErrorAssetDoesNotExist,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "Pool Not Found",
			Msg:                *types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetOneID, sdk.NewCoin("uasset1", sdk.NewInt(100)), 3, 1),
			ExpErr:             types.ErrPoolNotFound,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "App Mapping Id does not exists",
			Msg:                *types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetOneID, sdk.NewCoin("uasset1", sdk.NewInt(100)), poolOneID, 10),
			ExpErr:             types.ErrorAppMappingDoesNotExist,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "App Mapping Id mismatch, use the correct App Mapping ID in request",
			Msg:                *types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetOneID, sdk.NewCoin("uasset1", sdk.NewInt(100)), poolOneID, appTwoID),
			ExpErr:             types.ErrorAppMappingIDMismatch,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(100))),
		},
		{
			Name:               "invalid offer coin amount",
			Msg:                *types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetOneID, sdk.NewCoin("uasset2", sdk.NewInt(100)), poolOneID, appOneID),
			ExpErr:             errorsmod.Wrapf(types.ErrBadOfferCoinAmount, "uasset2"),
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(100))),
		},
		{
			Name:               "Asset Id not defined in the pool",
			Msg:                *types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetFourID, sdk.NewCoin("uasset4", sdk.NewInt(100)), poolOneID, appOneID),
			ExpErr:             errorsmod.Wrapf(types.ErrInvalidAssetIDForPool, "4"),
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(100))),
		},
		{
			Name:               "Asset Rates Stats not found",
			Msg:                *types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetFourID, sdk.NewCoin("uasset4", sdk.NewInt(100)), poolTwoID, appOneID),
			ExpErr:             errorsmod.Wrapf(types.ErrorAssetRatesParamsNotFound, "4"),
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(100))),
		},
		{
			Name:               "success valid case",
			Msg:                *types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetOneID, sdk.NewCoin("uasset1", sdk.NewInt(100)), poolTwoID, appOneID),
			ExpErr:             nil,
			ExpResp:            &types.MsgLendResponse{},
			QueryResponseIndex: 0,
			QueryResponse: &types.LendAsset{
				ID:                2,
				AssetID:           assetOneID,
				PoolID:            poolTwoID,
				Owner:             "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
				AmountIn:          sdk.NewCoin("uasset1", sdk.NewInt(100)),
				LendingTime:       time.Time{},
				AvailableToBorrow: sdk.NewInt(100),
				AppID:             appOneID,
				CPoolName:         "OSMO-ATOM-CMST",
			},
			AvailableBalance: sdk.NewCoins(sdk.NewCoin("ucasset1", newInt(100))),
		},
		{
			Name:               "success valid case",
			Msg:                *types.NewMsgLend("cosmos14edpcw6ptcqd2vct9rkjf7lgyvrlwdtd0rqrtx", assetTwoID, sdk.NewCoin("uasset2", sdk.NewInt(100)), poolOneID, appOneID),
			ExpErr:             nil,
			ExpResp:            &types.MsgLendResponse{},
			QueryResponseIndex: 0,
			QueryResponse: &types.LendAsset{
				ID:                3,
				AssetID:           assetOneID,
				PoolID:            poolOneID,
				Owner:             "cosmos14edpcw6ptcqd2vct9rkjf7lgyvrlwdtd0rqrtx",
				AmountIn:          sdk.NewCoin("uasset1", sdk.NewInt(100)),
				LendingTime:       time.Time{},
				AvailableToBorrow: sdk.NewInt(100),
				AppID:             appOneID,
				CPoolName:         "OSMO-ATOM-CMST",
			},
			AvailableBalance: sdk.NewCoins(sdk.NewCoin("ucasset1", newInt(100))),
		},
		//{
		//	Name:               "Duplicate lend Position",
		//	Msg:                *types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetOneID, sdk.NewCoin("uasset1", sdk.NewInt(100)), poolOneID, appOneID),
		//	ExpErr:             types.ErrorDuplicateLend,
		//	ExpResp:            nil,
		//	QueryResponseIndex: 0,
		//	QueryResponse:      nil,
		//	AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(100))),
		//},
	}
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.Name, func() {
			// add funds to acount for valid case
			if tc.ExpErr == nil {
				s.fundAddr(sdk.MustAccAddressFromBech32(tc.Msg.Lender), sdk.NewCoins(sdk.NewCoin("uasset1", tc.Msg.Amount.Amount), sdk.NewCoin("uasset2", tc.Msg.Amount.Amount)))
			}

			ctx := sdk.WrapSDKContext(s.ctx)
			resp, err := s.msgServer.Lend(ctx, &tc.Msg)
			if tc.ExpErr != nil {
				s.Require().Error(err)
				s.Require().EqualError(err, tc.ExpErr.Error())
				s.Require().Equal(tc.ExpResp, resp)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(resp)
				s.Require().Equal(tc.ExpResp, resp)

				// availableBalances := s.getBalances(sdk.MustAccAddressFromBech32(tc.Msg.Lender))
				// s.Require().True(tc.AvailableBalance.IsEqual(availableBalances))
			}
		})
	}
}

func (s *KeeperTestSuite) TestMsgWithdraw() {
	assetOneID := s.CreateNewAsset("ASSETONE", "uasset1", 1000000)
	assetTwoID := s.CreateNewAsset("ASSETTWO", "uasset2", 2000000)
	assetThreeID := s.CreateNewAsset("ASSETTHREE", "uasset3", 2000000)
	assetFourID := s.CreateNewAsset("ASSETFOUR", "uasset4", 2000000)
	cAssetOneID := s.CreateNewAsset("CASSETONE", "ucasset1", 1000000)
	cAssetTwoID := s.CreateNewAsset("CASSETTWO", "ucasset2", 2000000)
	cAssetThreeID := s.CreateNewAsset("CASSETTHRE", "ucasset3", 2000000)
	// cAssetFourID := s.CreateNewAsset("CASSETFOUR", "ucasset4", 2000000)

	var (
		assetDataPoolOne []*types.AssetDataPoolMapping
		assetDataPoolTwo []*types.AssetDataPoolMapping
	)
	assetDataPoolOneAssetOne := &types.AssetDataPoolMapping{
		AssetID:          assetOneID,
		AssetTransitType: 3,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolOneAssetTwo := &types.AssetDataPoolMapping{
		AssetID:          assetTwoID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(1000000000000000000),
	}
	assetDataPoolOneAssetThree := &types.AssetDataPoolMapping{
		AssetID:          assetThreeID,
		AssetTransitType: 2,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolTwoAssetFour := &types.AssetDataPoolMapping{
		AssetID:          assetFourID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(3000000000000000000),
	}

	assetDataPoolOne = append(assetDataPoolOne, assetDataPoolOneAssetOne, assetDataPoolOneAssetTwo, assetDataPoolOneAssetThree)
	assetDataPoolTwo = append(assetDataPoolOne, assetDataPoolTwoAssetFour, assetDataPoolOneAssetTwo, assetDataPoolOneAssetThree)

	poolOneID := s.CreateNewPool("cmdx", "CMDX-ATOM-CMST", assetDataPoolOne)
	poolTwoID := s.CreateNewPool("osmo", "OSMO-ATOM-CMST", assetDataPoolTwo)

	s.AddAssetRatesStats(assetThreeID, newDec("0.8"), newDec("0.002"), newDec("0.06"), newDec("0.6"), true, newDec("0.04"), newDec("0.04"), newDec("0.06"), newDec("0.8"), newDec("0.85"), newDec("0.025"), newDec("0.025"), newDec("0.1"), cAssetThreeID)
	s.AddAssetRatesStats(assetOneID, newDec("0.75"), newDec("0.002"), newDec("0.07"), newDec("1.25"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.7"), newDec("0.75"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetOneID)
	// s.AddAssetRatesStats(assetFourID, newDec("0.65"), newDec("0.002"), newDec("0.08"), newDec("1.5"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.6"), newDec("0.65"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetFourID)
	s.AddAssetRatesStats(assetTwoID, newDec("0.5"), newDec("0.002"), newDec("0.08"), newDec("2.0"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.5"), newDec("0.55"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetTwoID)

	pairOneID := s.AddExtendedLendPair(assetTwoID, assetThreeID, false, poolOneID, 1000000)
	pairTwoID := s.AddExtendedLendPair(assetTwoID, assetOneID, false, poolOneID, 1000000)
	pairThreeID := s.AddExtendedLendPair(assetOneID, assetTwoID, false, poolOneID, 1000000)
	pairFourID := s.AddExtendedLendPair(assetOneID, assetThreeID, false, poolOneID, 1000000)
	pairFiveID := s.AddExtendedLendPair(assetThreeID, assetTwoID, false, poolOneID, 1000000)
	pairSixID := s.AddExtendedLendPair(assetThreeID, assetOneID, false, poolOneID, 1000000)
	pairSevenID := s.AddExtendedLendPair(assetFourID, assetThreeID, false, poolTwoID, 1000000)
	pairEightID := s.AddExtendedLendPair(assetFourID, assetOneID, false, poolTwoID, 1000000)
	pairNineID := s.AddExtendedLendPair(assetOneID, assetFourID, false, poolTwoID, 1000000)
	pairTenID := s.AddExtendedLendPair(assetOneID, assetThreeID, false, poolTwoID, 1000000)
	pairElevenID := s.AddExtendedLendPair(assetThreeID, assetFourID, false, poolTwoID, 1000000)
	pairTwelveID := s.AddExtendedLendPair(assetThreeID, assetOneID, false, poolTwoID, 1000000)
	pairThirteenID := s.AddExtendedLendPair(assetTwoID, assetFourID, true, poolTwoID, 1000000)
	pairFourteenID := s.AddExtendedLendPair(assetThreeID, assetFourID, true, poolTwoID, 1000000)
	pairFifteenID := s.AddExtendedLendPair(assetOneID, assetFourID, true, poolTwoID, 1000000)
	pairSixteenID := s.AddExtendedLendPair(assetFourID, assetTwoID, true, poolOneID, 1000000)
	pairSeventeenID := s.AddExtendedLendPair(assetThreeID, assetTwoID, true, poolOneID, 1000000)
	pairEighteenID := s.AddExtendedLendPair(assetOneID, assetTwoID, true, poolOneID, 1000000)

	s.AddAssetToPair(assetOneID, poolOneID, []uint64{pairThreeID, pairFourID, pairFifteenID})
	s.AddAssetToPair(assetTwoID, poolOneID, []uint64{pairOneID, pairTwoID, pairThirteenID})
	s.AddAssetToPair(assetThreeID, poolOneID, []uint64{pairFiveID, pairSixID, pairFourteenID})
	s.AddAssetToPair(assetFourID, poolTwoID, []uint64{pairSevenID, pairEightID, pairSixteenID})
	s.AddAssetToPair(assetOneID, poolTwoID, []uint64{pairNineID, pairTenID, pairEighteenID})
	s.AddAssetToPair(assetThreeID, poolTwoID, []uint64{pairElevenID, pairTwelveID, pairSeventeenID})

	appOneID := s.CreateNewApp("commodo", "cmmdo")

	msg := types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetOneID, sdk.NewCoin("uasset1", newInt(100)), poolOneID, appOneID)
	msg2 := types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetTwoID, sdk.NewCoin("uasset2", newInt(100)), poolOneID, appOneID)

	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset1", newInt(100))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset2", newInt(100))))
	_, _ = s.msgServer.Lend(sdk.WrapSDKContext(s.ctx), msg)
	_, _ = s.msgServer.Lend(sdk.WrapSDKContext(s.ctx), msg2)

	testCases := []struct {
		Name               string
		Msg                types.MsgWithdraw
		ExpErr             error
		ExpResp            *types.MsgWithdrawResponse
		QueryResponseIndex uint64
		QueryResponse      *types.MsgWithdraw
		AvailableBalance   sdk.Coins
	}{
		{
			Name:               "Lend Position not found",
			Msg:                *types.NewMsgWithdraw("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 3, sdk.NewCoin("uasset1", sdk.NewInt(100))),
			ExpErr:             types.ErrLendNotFound,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
		},
		{
			Name:               "invalid offer coin amount",
			Msg:                *types.NewMsgWithdraw("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, sdk.NewCoin("uasset2", sdk.NewInt(10))),
			ExpErr:             errorsmod.Wrap(types.ErrBadOfferCoinAmount, "uasset2"),
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
		},
		{
			Name:               "Withdraw Amount Limit Exceeded",
			Msg:                *types.NewMsgWithdraw("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, sdk.NewCoin("uasset1", sdk.NewInt(101))),
			ExpErr:             types.ErrWithdrawAmountLimitExceeds,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
		},
		{
			Name:               "success valid case",
			Msg:                *types.NewMsgWithdraw("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, sdk.NewCoin("uasset1", sdk.NewInt(10))),
			ExpErr:             nil,
			ExpResp:            &types.MsgWithdrawResponse{},
			QueryResponseIndex: 0,
			QueryResponse: &types.MsgWithdraw{
				Lender: "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
				LendId: 1,
				Amount: sdk.NewCoin("uasset1", sdk.NewInt(10)),
			},
			AvailableBalance: sdk.NewCoins(sdk.NewCoin("uasset1", newInt(10)), sdk.NewCoin("ucasset1", newInt(90)), sdk.NewCoin("ucasset2", newInt(100))),
		},
	}
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.Name, func() {
			// add funds to acount for valid case
			//if tc.ExpErr == nil {
			//	s.fundAddr(sdk.MustAccAddressFromBech32(tc.Msg.Lender), sdk.NewCoins(sdk.NewCoin("uasset1", tc.Msg.Amount.Amount)))
			//}

			ctx := sdk.WrapSDKContext(s.ctx)
			resp, err := s.msgServer.Withdraw(ctx, &tc.Msg)
			if tc.ExpErr != nil {
				s.Require().Error(err)
				s.Require().EqualError(err, tc.ExpErr.Error())
				s.Require().Equal(tc.ExpResp, resp)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(resp)
				s.Require().Equal(tc.ExpResp, resp)

				availableBalances := s.getBalances(sdk.MustAccAddressFromBech32(tc.Msg.Lender))
				s.Require().True(tc.AvailableBalance.IsEqual(availableBalances))
			}
		})
	}
}

func (s *KeeperTestSuite) TestMsgDeposit() {
	assetOneID := s.CreateNewAsset("ASSETONE", "uasset1", 1000000)
	assetTwoID := s.CreateNewAsset("ASSETTWO", "uasset2", 2000000)
	assetThreeID := s.CreateNewAsset("ASSETTHREE", "uasset3", 2000000)
	assetFourID := s.CreateNewAsset("ASSETFOUR", "uasset4", 2000000)
	cAssetOneID := s.CreateNewAsset("CASSETONE", "ucasset1", 1000000)
	cAssetTwoID := s.CreateNewAsset("CASSETTWO", "ucasset2", 2000000)
	cAssetThreeID := s.CreateNewAsset("CASSETTHRE", "ucasset3", 2000000)
	// cAssetFourID := s.CreateNewAsset("CASSETFOUR", "ucasset4", 2000000)

	var (
		assetDataPoolOne []*types.AssetDataPoolMapping
		assetDataPoolTwo []*types.AssetDataPoolMapping
	)
	assetDataPoolOneAssetOne := &types.AssetDataPoolMapping{
		AssetID:          assetOneID,
		AssetTransitType: 3,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolOneAssetTwo := &types.AssetDataPoolMapping{
		AssetID:          assetTwoID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(1000000000000000000),
	}
	assetDataPoolOneAssetThree := &types.AssetDataPoolMapping{
		AssetID:          assetThreeID,
		AssetTransitType: 2,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolTwoAssetFour := &types.AssetDataPoolMapping{
		AssetID:          assetFourID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(3000000000000000000),
	}

	assetDataPoolOne = append(assetDataPoolOne, assetDataPoolOneAssetOne, assetDataPoolOneAssetTwo, assetDataPoolOneAssetThree)
	assetDataPoolTwo = append(assetDataPoolOne, assetDataPoolTwoAssetFour, assetDataPoolOneAssetTwo, assetDataPoolOneAssetThree)

	poolOneID := s.CreateNewPool("cmdx", "CMDX-ATOM-CMST", assetDataPoolOne)
	poolTwoID := s.CreateNewPool("osmo", "OSMO-ATOM-CMST", assetDataPoolTwo)

	s.AddAssetRatesStats(assetThreeID, newDec("0.8"), newDec("0.002"), newDec("0.06"), newDec("0.6"), true, newDec("0.04"), newDec("0.04"), newDec("0.06"), newDec("0.8"), newDec("0.85"), newDec("0.025"), newDec("0.025"), newDec("0.1"), cAssetThreeID)
	s.AddAssetRatesStats(assetOneID, newDec("0.75"), newDec("0.002"), newDec("0.07"), newDec("1.25"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.7"), newDec("0.75"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetOneID)
	// s.AddAssetRatesStats(assetFourID, newDec("0.65"), newDec("0.002"), newDec("0.08"), newDec("1.5"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.6"), newDec("0.65"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetFourID)
	s.AddAssetRatesStats(assetTwoID, newDec("0.5"), newDec("0.002"), newDec("0.08"), newDec("2.0"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.5"), newDec("0.55"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetTwoID)

	pairOneID := s.AddExtendedLendPair(assetTwoID, assetThreeID, false, poolOneID, 1000000)
	pairTwoID := s.AddExtendedLendPair(assetTwoID, assetOneID, false, poolOneID, 1000000)
	pairThreeID := s.AddExtendedLendPair(assetOneID, assetTwoID, false, poolOneID, 1000000)
	pairFourID := s.AddExtendedLendPair(assetOneID, assetThreeID, false, poolOneID, 1000000)
	pairFiveID := s.AddExtendedLendPair(assetThreeID, assetTwoID, false, poolOneID, 1000000)
	pairSixID := s.AddExtendedLendPair(assetThreeID, assetOneID, false, poolOneID, 1000000)
	pairSevenID := s.AddExtendedLendPair(assetFourID, assetThreeID, false, poolTwoID, 1000000)
	pairEightID := s.AddExtendedLendPair(assetFourID, assetOneID, false, poolTwoID, 1000000)
	pairNineID := s.AddExtendedLendPair(assetOneID, assetFourID, false, poolTwoID, 1000000)
	pairTenID := s.AddExtendedLendPair(assetOneID, assetThreeID, false, poolTwoID, 1000000)
	pairElevenID := s.AddExtendedLendPair(assetThreeID, assetFourID, false, poolTwoID, 1000000)
	pairTwelveID := s.AddExtendedLendPair(assetThreeID, assetOneID, false, poolTwoID, 1000000)
	pairThirteenID := s.AddExtendedLendPair(assetTwoID, assetFourID, true, poolTwoID, 1000000)
	pairFourteenID := s.AddExtendedLendPair(assetThreeID, assetFourID, true, poolTwoID, 1000000)
	pairFifteenID := s.AddExtendedLendPair(assetOneID, assetFourID, true, poolTwoID, 1000000)
	pairSixteenID := s.AddExtendedLendPair(assetFourID, assetTwoID, true, poolOneID, 1000000)
	pairSeventeenID := s.AddExtendedLendPair(assetThreeID, assetTwoID, true, poolOneID, 1000000)
	pairEighteenID := s.AddExtendedLendPair(assetOneID, assetTwoID, true, poolOneID, 1000000)

	s.AddAssetToPair(assetOneID, poolOneID, []uint64{pairThreeID, pairFourID, pairFifteenID})
	s.AddAssetToPair(assetTwoID, poolOneID, []uint64{pairOneID, pairTwoID, pairThirteenID})
	s.AddAssetToPair(assetThreeID, poolOneID, []uint64{pairFiveID, pairSixID, pairFourteenID})
	s.AddAssetToPair(assetFourID, poolTwoID, []uint64{pairSevenID, pairEightID, pairSixteenID})
	s.AddAssetToPair(assetOneID, poolTwoID, []uint64{pairNineID, pairTenID, pairEighteenID})
	s.AddAssetToPair(assetThreeID, poolTwoID, []uint64{pairElevenID, pairTwelveID, pairSeventeenID})

	appOneID := s.CreateNewApp("commodo", "cmmdo")

	msg := types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetOneID, sdk.NewCoin("uasset1", newInt(100)), poolOneID, appOneID)
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset1", newInt(100))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset2", newInt(100))))
	_, _ = s.msgServer.Lend(sdk.WrapSDKContext(s.ctx), msg)

	testCases := []struct {
		Name               string
		Msg                types.MsgDeposit
		ExpErr             error
		ExpResp            *types.MsgDepositResponse
		QueryResponseIndex uint64
		QueryResponse      *types.MsgDeposit
		AvailableBalance   sdk.Coins
	}{
		{
			Name:               "Lend Position not found",
			Msg:                *types.NewMsgDeposit("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 2, sdk.NewCoin("uasset1", sdk.NewInt(100))),
			ExpErr:             types.ErrLendNotFound,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "invalid offer coin amount",
			Msg:                *types.NewMsgDeposit("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, sdk.NewCoin("uasset2", sdk.NewInt(100))),
			ExpErr:             errorsmod.Wrap(types.ErrBadOfferCoinAmount, "uasset2"),
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "success valid case",
			Msg:                *types.NewMsgDeposit("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, sdk.NewCoin("uasset1", sdk.NewInt(10))),
			ExpErr:             nil,
			ExpResp:            &types.MsgDepositResponse{},
			QueryResponseIndex: 0,
			QueryResponse: &types.MsgDeposit{
				Lender: "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
				LendId: 1,
				Amount: sdk.NewCoin("uasset2", sdk.NewInt(10)),
			},
			// AvailableBalance: sdk.NewCoins(sdk.NewCoin("ucasset1", newInt(90))),
		},
	}
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.Name, func() {
			// add funds to acount for valid case
			if tc.ExpErr == nil {
				s.fundAddr(sdk.MustAccAddressFromBech32(tc.Msg.Lender), sdk.NewCoins(sdk.NewCoin("uasset1", tc.Msg.Amount.Amount)))
			}

			ctx := sdk.WrapSDKContext(s.ctx)
			resp, err := s.msgServer.Deposit(ctx, &tc.Msg)
			if tc.ExpErr != nil {
				s.Require().Error(err)
				s.Require().EqualError(err, tc.ExpErr.Error())
				s.Require().Equal(tc.ExpResp, resp)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(resp)
				s.Require().Equal(tc.ExpResp, resp)

				// availableBalances := s.getBalances(sdk.MustAccAddressFromBech32(tc.Msg.Lender))
				// s.Require().True(tc.AvailableBalance.IsEqual(availableBalances))
			}
		})
	}
}

func (s *KeeperTestSuite) TestMsgCloseLend() {
	assetOneID := s.CreateNewAsset("ASSETONE", "uasset1", 1000000)
	assetTwoID := s.CreateNewAsset("ASSETTWO", "uasset2", 2000000)
	assetThreeID := s.CreateNewAsset("ASSETTHREE", "uasset3", 2000000)
	assetFourID := s.CreateNewAsset("ASSETFOUR", "uasset4", 2000000)
	cAssetOneID := s.CreateNewAsset("CASSETONE", "ucasset1", 1000000)
	cAssetTwoID := s.CreateNewAsset("CASSETTWO", "ucasset2", 2000000)
	cAssetThreeID := s.CreateNewAsset("CASSETTHRE", "ucasset3", 2000000)
	// cAssetFourID := s.CreateNewAsset("CASSETFOUR", "ucasset4", 2000000)

	var (
		assetDataPoolOne []*types.AssetDataPoolMapping
		assetDataPoolTwo []*types.AssetDataPoolMapping
	)
	assetDataPoolOneAssetOne := &types.AssetDataPoolMapping{
		AssetID:          assetOneID,
		AssetTransitType: 3,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolOneAssetTwo := &types.AssetDataPoolMapping{
		AssetID:          assetTwoID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(1000000000000000000),
	}
	assetDataPoolOneAssetThree := &types.AssetDataPoolMapping{
		AssetID:          assetThreeID,
		AssetTransitType: 2,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolTwoAssetFour := &types.AssetDataPoolMapping{
		AssetID:          assetFourID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(3000000000000000000),
	}

	assetDataPoolOne = append(assetDataPoolOne, assetDataPoolOneAssetOne, assetDataPoolOneAssetTwo, assetDataPoolOneAssetThree)
	assetDataPoolTwo = append(assetDataPoolOne, assetDataPoolTwoAssetFour, assetDataPoolOneAssetTwo, assetDataPoolOneAssetThree)

	poolOneID := s.CreateNewPool("cmdx", "CMDX-ATOM-CMST", assetDataPoolOne)
	poolTwoID := s.CreateNewPool("osmo", "OSMO-ATOM-CMST", assetDataPoolTwo)

	s.AddAssetRatesStats(assetThreeID, newDec("0.8"), newDec("0.002"), newDec("0.06"), newDec("0.6"), true, newDec("0.04"), newDec("0.04"), newDec("0.06"), newDec("0.8"), newDec("0.85"), newDec("0.025"), newDec("0.025"), newDec("0.1"), cAssetThreeID)
	s.AddAssetRatesStats(assetOneID, newDec("0.75"), newDec("0.002"), newDec("0.07"), newDec("1.25"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.7"), newDec("0.75"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetOneID)
	// s.AddAssetRatesStats(assetFourID, newDec("0.65"), newDec("0.002"), newDec("0.08"), newDec("1.5"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.6"), newDec("0.65"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetFourID)
	s.AddAssetRatesStats(assetTwoID, newDec("0.5"), newDec("0.002"), newDec("0.08"), newDec("2.0"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.5"), newDec("0.55"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetTwoID)

	pairOneID := s.AddExtendedLendPair(assetTwoID, assetThreeID, false, poolOneID, 1000000)
	pairTwoID := s.AddExtendedLendPair(assetTwoID, assetOneID, false, poolOneID, 1000000)
	pairThreeID := s.AddExtendedLendPair(assetOneID, assetTwoID, false, poolOneID, 1000000)
	pairFourID := s.AddExtendedLendPair(assetOneID, assetThreeID, false, poolOneID, 1000000)
	pairFiveID := s.AddExtendedLendPair(assetThreeID, assetTwoID, false, poolOneID, 1000000)
	pairSixID := s.AddExtendedLendPair(assetThreeID, assetOneID, false, poolOneID, 1000000)
	pairSevenID := s.AddExtendedLendPair(assetFourID, assetThreeID, false, poolTwoID, 1000000)
	pairEightID := s.AddExtendedLendPair(assetFourID, assetOneID, false, poolTwoID, 1000000)
	pairNineID := s.AddExtendedLendPair(assetOneID, assetFourID, false, poolTwoID, 1000000)
	pairTenID := s.AddExtendedLendPair(assetOneID, assetThreeID, false, poolTwoID, 1000000)
	pairElevenID := s.AddExtendedLendPair(assetThreeID, assetFourID, false, poolTwoID, 1000000)
	pairTwelveID := s.AddExtendedLendPair(assetThreeID, assetOneID, false, poolTwoID, 1000000)
	pairThirteenID := s.AddExtendedLendPair(assetTwoID, assetFourID, true, poolTwoID, 1000000)
	pairFourteenID := s.AddExtendedLendPair(assetThreeID, assetFourID, true, poolTwoID, 1000000)
	pairFifteenID := s.AddExtendedLendPair(assetOneID, assetFourID, true, poolTwoID, 1000000)
	pairSixteenID := s.AddExtendedLendPair(assetFourID, assetTwoID, true, poolOneID, 1000000)
	pairSeventeenID := s.AddExtendedLendPair(assetThreeID, assetTwoID, true, poolOneID, 1000000)
	pairEighteenID := s.AddExtendedLendPair(assetOneID, assetTwoID, true, poolOneID, 1000000)

	s.AddAssetToPair(assetOneID, poolOneID, []uint64{pairThreeID, pairFourID, pairFifteenID})
	s.AddAssetToPair(assetTwoID, poolOneID, []uint64{pairOneID, pairTwoID, pairThirteenID})
	s.AddAssetToPair(assetThreeID, poolOneID, []uint64{pairFiveID, pairSixID, pairFourteenID})
	s.AddAssetToPair(assetFourID, poolTwoID, []uint64{pairSevenID, pairEightID, pairSixteenID})
	s.AddAssetToPair(assetOneID, poolTwoID, []uint64{pairNineID, pairTenID, pairEighteenID})
	s.AddAssetToPair(assetThreeID, poolTwoID, []uint64{pairElevenID, pairTwelveID, pairSeventeenID})

	appOneID := s.CreateNewApp("commodo", "cmmdo")

	msg := types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetOneID, sdk.NewCoin("uasset1", newInt(100)), poolOneID, appOneID)
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset1", newInt(100))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset2", newInt(100))))
	_, _ = s.msgServer.Lend(sdk.WrapSDKContext(s.ctx), msg)

	testCases := []struct {
		Name               string
		Msg                types.MsgCloseLend
		ExpErr             error
		ExpResp            *types.MsgCloseLendResponse
		QueryResponseIndex uint64
		QueryResponse      *types.MsgCloseLend
		AvailableBalance   sdk.Coins
	}{
		{
			Name:               "Lend Position not found",
			Msg:                *types.NewMsgCloseLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 2),
			ExpErr:             types.ErrLendNotFound,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "success valid case",
			Msg:                *types.NewMsgCloseLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1),
			ExpErr:             nil,
			ExpResp:            &types.MsgCloseLendResponse{},
			QueryResponseIndex: 0,
			QueryResponse: &types.MsgCloseLend{
				Lender: "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
				LendId: 1,
			},
			// AvailableBalance: sdk.NewCoins(sdk.NewCoin("ucasset1", newInt(90))),
		},
	}
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.Name, func() {
			// add funds to acount for valid case
			if tc.ExpErr == nil {
				s.fundAddr(sdk.MustAccAddressFromBech32(tc.Msg.Lender), sdk.NewCoins(sdk.NewCoin("uasset1", sdk.NewIntFromUint64(100))))
			}

			ctx := sdk.WrapSDKContext(s.ctx)
			resp, err := s.msgServer.CloseLend(ctx, &tc.Msg)
			if tc.ExpErr != nil {
				s.Require().Error(err)
				s.Require().EqualError(err, tc.ExpErr.Error())
				s.Require().Equal(tc.ExpResp, resp)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(resp)
				s.Require().Equal(tc.ExpResp, resp)

				// availableBalances := s.getBalances(sdk.MustAccAddressFromBech32(tc.Msg.Lender))
				// s.Require().True(tc.AvailableBalance.IsEqual(availableBalances))
			}
		})
	}
}

func (s *KeeperTestSuite) TestMsgBorrow() {
	assetOneID := s.CreateNewAsset("ASSETONE", "uasset1", 1000000)
	assetTwoID := s.CreateNewAsset("ASSETTWO", "uasset2", 2000000)
	assetThreeID := s.CreateNewAsset("ASSETTHREE", "uasset3", 2000000)
	assetFourID := s.CreateNewAsset("ASSETFOUR", "uasset4", 2000000)
	cAssetOneID := s.CreateNewAsset("CASSETONE", "ucasset1", 1000000)
	cAssetTwoID := s.CreateNewAsset("CASSETTWO", "ucasset2", 2000000)
	cAssetThreeID := s.CreateNewAsset("CASSETTHRE", "ucasset3", 2000000)
	cAssetFourID := s.CreateNewAsset("CASSETFOUR", "ucasset4", 2000000)

	var (
		assetDataPoolOne []*types.AssetDataPoolMapping
		assetDataPoolTwo []*types.AssetDataPoolMapping
	)
	assetDataPoolOneAssetOne := &types.AssetDataPoolMapping{
		AssetID:          assetOneID,
		AssetTransitType: 3,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolOneAssetTwo := &types.AssetDataPoolMapping{
		AssetID:          assetTwoID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(1000000000000000000),
	}
	assetDataPoolOneAssetThree := &types.AssetDataPoolMapping{
		AssetID:          assetThreeID,
		AssetTransitType: 2,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolTwoAssetFour := &types.AssetDataPoolMapping{
		AssetID:          assetFourID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(3000000000000000000),
	}

	assetDataPoolOne = append(assetDataPoolOne, assetDataPoolOneAssetOne, assetDataPoolOneAssetTwo, assetDataPoolOneAssetThree)
	assetDataPoolTwo = append(assetDataPoolOne, assetDataPoolTwoAssetFour, assetDataPoolOneAssetOne, assetDataPoolOneAssetThree)

	poolOneID := s.CreateNewPool("cmdx", "CMDX-ATOM-CMST", assetDataPoolOne)
	poolTwoID := s.CreateNewPool("osmo", "OSMO-ATOM-CMST", assetDataPoolTwo)

	s.AddAssetRatesStats(assetThreeID, newDec("0.8"), newDec("0.002"), newDec("0.06"), newDec("0.6"), true, newDec("0.04"), newDec("0.04"), newDec("0.06"), newDec("0.8"), newDec("0.85"), newDec("0.025"), newDec("0.025"), newDec("0.1"), cAssetThreeID)
	s.AddAssetRatesStats(assetOneID, newDec("0.75"), newDec("0.002"), newDec("0.07"), newDec("1.25"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.7"), newDec("0.75"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetOneID)
	s.AddAssetRatesStats(assetFourID, newDec("0.65"), newDec("0.002"), newDec("0.08"), newDec("1.5"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.6"), newDec("0.65"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetFourID)
	s.AddAssetRatesStats(assetTwoID, newDec("0.5"), newDec("0.002"), newDec("0.08"), newDec("2.0"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.5"), newDec("0.55"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetTwoID)

	pairOneID := s.AddExtendedLendPair(assetTwoID, assetThreeID, false, poolOneID, 1000000)
	pairTwoID := s.AddExtendedLendPair(assetTwoID, assetOneID, false, poolOneID, 1000000)
	pairThreeID := s.AddExtendedLendPair(assetOneID, assetTwoID, false, poolOneID, 1000000)
	pairFourID := s.AddExtendedLendPair(assetOneID, assetThreeID, false, poolOneID, 1000000)
	pairFiveID := s.AddExtendedLendPair(assetThreeID, assetTwoID, false, poolOneID, 1000000)
	pairSixID := s.AddExtendedLendPair(assetThreeID, assetOneID, false, poolOneID, 1000000)
	pairSevenID := s.AddExtendedLendPair(assetFourID, assetThreeID, false, poolTwoID, 1000000)
	pairEightID := s.AddExtendedLendPair(assetFourID, assetOneID, false, poolTwoID, 1000000)
	pairNineID := s.AddExtendedLendPair(assetOneID, assetFourID, false, poolTwoID, 1000000)
	pairTenID := s.AddExtendedLendPair(assetOneID, assetThreeID, false, poolTwoID, 1000000)
	pairElevenID := s.AddExtendedLendPair(assetThreeID, assetFourID, false, poolTwoID, 1000000)
	pairTwelveID := s.AddExtendedLendPair(assetThreeID, assetOneID, false, poolTwoID, 1000000)
	pairThirteenID := s.AddExtendedLendPair(assetTwoID, assetFourID, true, poolTwoID, 1000000)
	pairFourteenID := s.AddExtendedLendPair(assetThreeID, assetFourID, true, poolTwoID, 1000000)
	pairFifteenID := s.AddExtendedLendPair(assetOneID, assetFourID, true, poolTwoID, 1000000)
	pairSixteenID := s.AddExtendedLendPair(assetFourID, assetTwoID, true, poolOneID, 1000000)
	pairSeventeenID := s.AddExtendedLendPair(assetThreeID, assetTwoID, true, poolOneID, 1000000)
	pairEighteenID := s.AddExtendedLendPair(assetOneID, assetTwoID, true, poolOneID, 1000000)

	s.AddAssetToPair(assetOneID, poolOneID, []uint64{pairThreeID, pairFourID, pairFifteenID})
	s.AddAssetToPair(assetTwoID, poolOneID, []uint64{pairOneID, pairTwoID, pairThirteenID})
	s.AddAssetToPair(assetThreeID, poolOneID, []uint64{pairFiveID, pairSixID, pairFourteenID})
	s.AddAssetToPair(assetFourID, poolTwoID, []uint64{pairSevenID, pairEightID, pairSixteenID})
	s.AddAssetToPair(assetOneID, poolTwoID, []uint64{pairNineID, pairTenID, pairEighteenID})
	s.AddAssetToPair(assetThreeID, poolTwoID, []uint64{pairElevenID, pairTwelveID, pairSeventeenID})

	appOneID := s.CreateNewApp("commodo", "cmmdo")
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset1", newInt(1000000000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset2", newInt(1000000000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset3", newInt(1000000000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset4", newInt(1000000000000000))))

	msg := types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetOneID, sdk.NewCoin("uasset1", newInt(300)), poolOneID, appOneID)
	msgLend2 := types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetTwoID, sdk.NewCoin("uasset2", newInt(10000000000)), poolOneID, appOneID)

	msg3 := types.NewMsgFundModuleAccounts(1, assetOneID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset1", newInt(10000000000)))
	msg4 := types.NewMsgFundModuleAccounts(1, assetTwoID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset2", newInt(10000000000)))
	msg5 := types.NewMsgFundModuleAccounts(1, assetThreeID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset3", newInt(120000000)))
	// msg6 := types.NewMsgFundModuleAccounts(2, assetThreeID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset3", newInt(10000000000)))
	msg7 := types.NewMsgFundModuleAccounts(2, assetOneID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset1", newInt(10000000000)))
	msg8 := types.NewMsgFundModuleAccounts(2, assetFourID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset4", newInt(10000000000)))

	_, _ = s.msgServer.Lend(sdk.WrapSDKContext(s.ctx), msg)
	_, _ = s.msgServer.Lend(sdk.WrapSDKContext(s.ctx), msgLend2)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg3)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg4)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg5)
	//_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg6)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg7)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg8)

	msg2 := types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, 3, false, sdk.NewCoin("ucasset1", newInt(100)), sdk.NewCoin("uasset2", newInt(10)))
	_, _ = s.msgServer.Borrow(sdk.WrapSDKContext(s.ctx), msg2)

	testCases := []struct {
		Name               string
		Msg                types.MsgBorrow
		ExpErr             error
		ExpResp            *types.MsgBorrowResponse
		QueryResponseIndex uint64
		QueryResponse      *types.MsgBorrow
		AvailableBalance   sdk.Coins
	}{
		{
			Name:               "Pair Not Found",
			Msg:                *types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, 90, false, sdk.NewCoin("uasset2", newInt(100)), sdk.NewCoin("uasset2", newInt(100))),
			ExpErr:             types.ErrorPairNotFound,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "Pair Not Found",
			Msg:                *types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, 18, false, sdk.NewCoin("uasset2", newInt(100)), sdk.NewCoin("uasset2", newInt(100))),
			ExpErr:             types.ErrorPairNotFound,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "invalid offer coin Type",
			Msg:                *types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, 1, false, sdk.NewCoin("uasset2", newInt(100)), sdk.NewCoin("uasset2", newInt(100))),
			ExpErr:             types.ErrBadOfferCoinType,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "Available To Borrow Insufficient",
			Msg:                *types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, 4, false, sdk.NewCoin("ucasset1", newInt(201)), sdk.NewCoin("uasset3", newInt(10))),
			ExpErr:             types.ErrAvailableToBorrowInsufficient,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "invalid asset",
			Msg:                *types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, 4, false, sdk.NewCoin("ucasset1", newInt(100)), sdk.NewCoin("uasset1", newInt(10))),
			ExpErr:             types.ErrInvalidAsset,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "Stable Borrow Rate Not Enabled for This Asset",
			Msg:                *types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, 4, true, sdk.NewCoin("ucasset1", newInt(100)), sdk.NewCoin("uasset3", newInt(10))),
			ExpErr:             errorsmod.Wrap(types.ErrStableBorrowDisabled, "10uasset3"),
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "Error Invalid Collaterallization Ratio",
			Msg:                *types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, 4, false, sdk.NewCoin("ucasset1", newInt(100)), sdk.NewCoin("uasset3", newInt(100))),
			ExpErr:             types.ErrorInvalidCollateralizationRatio,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "Unauthorized User",
			Msg:                *types.NewMsgBorrow("cosmos14edpcw6ptcqd2vct9rkjf7lgyvrlwdtd0rqrtx", 1, 4, false, sdk.NewCoin("ucasset1", newInt(100)), sdk.NewCoin("uasset3", newInt(100))),
			ExpErr:             types.ErrLendAccessUnauthorized,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "success valid case",
			Msg:                *types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, 4, false, sdk.NewCoin("ucasset1", newInt(100)), sdk.NewCoin("uasset3", newInt(10))),
			ExpErr:             nil,
			ExpResp:            &types.MsgBorrowResponse{},
			QueryResponseIndex: 0,
			QueryResponse: &types.MsgBorrow{
				Borrower:       "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
				LendId:         1,
				PairId:         4,
				IsStableBorrow: false,
				AmountIn:       sdk.NewCoin("ucasset1", newInt(100)),
				AmountOut:      sdk.NewCoin("uasset3", newInt(10)),
			},
			AvailableBalance: sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "success valid case 2",
			Msg:                *types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, 15, false, sdk.NewCoin("ucasset1", newInt(100)), sdk.NewCoin("uasset4", newInt(1))),
			ExpErr:             nil,
			ExpResp:            &types.MsgBorrowResponse{},
			QueryResponseIndex: 0,
			QueryResponse: &types.MsgBorrow{
				Borrower:       "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
				LendId:         1,
				PairId:         15,
				IsStableBorrow: false,
				AmountIn:       sdk.NewCoin("ucasset1", newInt(100)),
				AmountOut:      sdk.NewCoin("uasset4", newInt(10)),
			},
			AvailableBalance: sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "success valid case 3",
			Msg:                *types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 2, 13, false, sdk.NewCoin("ucasset2", newInt(1000000000)), sdk.NewCoin("uasset4", newInt(1000000))),
			ExpErr:             nil,
			ExpResp:            &types.MsgBorrowResponse{},
			QueryResponseIndex: 0,
			QueryResponse: &types.MsgBorrow{
				Borrower:       "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
				LendId:         2,
				PairId:         13,
				IsStableBorrow: false,
				AmountIn:       sdk.NewCoin("ucasset2", newInt(1000000000)),
				AmountOut:      sdk.NewCoin("uasset4", newInt(100000000)),
			},
			AvailableBalance: sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "success valid case 4",
			Msg:                *types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 2, 13, false, sdk.NewCoin("ucasset2", newInt(1000000000)), sdk.NewCoin("uasset4", newInt(1000000))),
			ExpErr:             nil,
			ExpResp:            &types.MsgBorrowResponse{},
			QueryResponseIndex: 0,
			QueryResponse: &types.MsgBorrow{
				Borrower:       "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
				LendId:         2,
				PairId:         13,
				IsStableBorrow: false,
				AmountIn:       sdk.NewCoin("ucasset2", newInt(2000000000)),
				AmountOut:      sdk.NewCoin("uasset4", newInt(200000000)),
			},
			AvailableBalance: sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
	}
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.Name, func() {
			// add funds to acount for valid case
			if tc.ExpErr == nil {
				s.fundAddr(sdk.MustAccAddressFromBech32(tc.Msg.Borrower), sdk.NewCoins(sdk.NewCoin("uasset1", sdk.NewIntFromUint64(300))))
				s.fundAddr(sdk.MustAccAddressFromBech32(tc.Msg.Borrower), sdk.NewCoins(sdk.NewCoin("ucasset2", sdk.NewIntFromUint64(1000000000))))
			}

			ctx := sdk.WrapSDKContext(s.ctx)
			resp, err := s.msgServer.Borrow(ctx, &tc.Msg)
			if tc.ExpErr != nil {
				s.Require().Error(err)
				s.Require().EqualError(err, tc.ExpErr.Error())
				s.Require().Equal(tc.ExpResp, resp)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(resp)
				s.Require().Equal(tc.ExpResp, resp)

				// availableBalances := s.getBalances(sdk.MustAccAddressFromBech32(tc.Msg.Lender))
				// s.Require().True(tc.AvailableBalance.IsEqual(availableBalances))
			}
		})
	}
}

func (s *KeeperTestSuite) TestMsgRepay() {
	assetOneID := s.CreateNewAsset("ASSETONE", "uasset1", 1000000)
	assetTwoID := s.CreateNewAsset("ASSETTWO", "uasset2", 2000000)
	assetThreeID := s.CreateNewAsset("ASSETTHREE", "uasset3", 2000000)
	assetFourID := s.CreateNewAsset("ASSETFOUR", "uasset4", 2000000)
	cAssetOneID := s.CreateNewAsset("CASSETONE", "ucasset1", 1000000)
	cAssetTwoID := s.CreateNewAsset("CASSETTWO", "ucasset2", 2000000)
	cAssetThreeID := s.CreateNewAsset("CASSETTHRE", "ucasset3", 2000000)
	cAssetFourID := s.CreateNewAsset("CASSETFOUR", "ucasset4", 2000000)

	var (
		assetDataPoolOne []*types.AssetDataPoolMapping
		assetDataPoolTwo []*types.AssetDataPoolMapping
	)
	assetDataPoolOneAssetOne := &types.AssetDataPoolMapping{
		AssetID:          assetOneID,
		AssetTransitType: 3,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolOneAssetTwo := &types.AssetDataPoolMapping{
		AssetID:          assetTwoID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(1000000000000000000),
	}
	assetDataPoolOneAssetThree := &types.AssetDataPoolMapping{
		AssetID:          assetThreeID,
		AssetTransitType: 2,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolTwoAssetFour := &types.AssetDataPoolMapping{
		AssetID:          assetFourID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(3000000000000000000),
	}
	assetDataPoolOne = append(assetDataPoolOne, assetDataPoolOneAssetOne, assetDataPoolOneAssetTwo, assetDataPoolOneAssetThree)
	assetDataPoolTwo = append(assetDataPoolOne, assetDataPoolTwoAssetFour, assetDataPoolOneAssetOne, assetDataPoolOneAssetThree)

	poolOneID := s.CreateNewPool("cmdx", "CMDX-ATOM-CMST", assetDataPoolOne)
	poolTwoID := s.CreateNewPool("osmo", "OSMO-ATOM-CMST", assetDataPoolTwo)

	s.AddAssetRatesStats(assetThreeID, newDec("0.8"), newDec("0.002"), newDec("0.06"), newDec("0.6"), true, newDec("0.04"), newDec("0.04"), newDec("0.06"), newDec("0.8"), newDec("0.85"), newDec("0.025"), newDec("0.025"), newDec("0.1"), cAssetThreeID)
	s.AddAssetRatesStats(assetOneID, newDec("0.75"), newDec("0.002"), newDec("0.07"), newDec("1.25"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.7"), newDec("0.75"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetOneID)
	s.AddAssetRatesStats(assetFourID, newDec("0.65"), newDec("0.002"), newDec("0.08"), newDec("1.5"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.6"), newDec("0.65"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetFourID)
	s.AddAssetRatesStats(assetTwoID, newDec("0.5"), newDec("0.002"), newDec("0.08"), newDec("2.0"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.5"), newDec("0.55"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetTwoID)

	pairOneID := s.AddExtendedLendPair(assetTwoID, assetThreeID, false, poolOneID, 1000000)
	pairTwoID := s.AddExtendedLendPair(assetTwoID, assetOneID, false, poolOneID, 1000000)
	pairThreeID := s.AddExtendedLendPair(assetOneID, assetTwoID, false, poolOneID, 1000000)
	pairFourID := s.AddExtendedLendPair(assetOneID, assetThreeID, false, poolOneID, 1000000)
	pairFiveID := s.AddExtendedLendPair(assetThreeID, assetTwoID, false, poolOneID, 1000000)
	pairSixID := s.AddExtendedLendPair(assetThreeID, assetOneID, false, poolOneID, 1000000)
	pairSevenID := s.AddExtendedLendPair(assetFourID, assetThreeID, false, poolTwoID, 1000000)
	pairEightID := s.AddExtendedLendPair(assetFourID, assetOneID, false, poolTwoID, 1000000)
	pairNineID := s.AddExtendedLendPair(assetOneID, assetFourID, false, poolTwoID, 1000000)
	pairTenID := s.AddExtendedLendPair(assetOneID, assetThreeID, false, poolTwoID, 1000000)
	pairElevenID := s.AddExtendedLendPair(assetThreeID, assetFourID, false, poolTwoID, 1000000)
	pairTwelveID := s.AddExtendedLendPair(assetThreeID, assetOneID, false, poolTwoID, 1000000)
	pairThirteenID := s.AddExtendedLendPair(assetTwoID, assetFourID, true, poolTwoID, 1000000)
	pairFourteenID := s.AddExtendedLendPair(assetThreeID, assetFourID, true, poolTwoID, 1000000)
	pairFifteenID := s.AddExtendedLendPair(assetOneID, assetFourID, true, poolTwoID, 1000000)
	pairSixteenID := s.AddExtendedLendPair(assetFourID, assetTwoID, true, poolOneID, 1000000)
	pairSeventeenID := s.AddExtendedLendPair(assetThreeID, assetTwoID, true, poolOneID, 1000000)
	pairEighteenID := s.AddExtendedLendPair(assetOneID, assetTwoID, true, poolOneID, 1000000)

	s.AddAssetToPair(assetOneID, poolOneID, []uint64{pairThreeID, pairFourID, pairFifteenID})
	s.AddAssetToPair(assetTwoID, poolOneID, []uint64{pairOneID, pairTwoID, pairThirteenID})
	s.AddAssetToPair(assetThreeID, poolOneID, []uint64{pairFiveID, pairSixID, pairFourteenID})
	s.AddAssetToPair(assetFourID, poolTwoID, []uint64{pairSevenID, pairEightID, pairSixteenID})
	s.AddAssetToPair(assetOneID, poolTwoID, []uint64{pairNineID, pairTenID, pairEighteenID})
	s.AddAssetToPair(assetThreeID, poolTwoID, []uint64{pairElevenID, pairTwelveID, pairSeventeenID})

	appOneID := s.CreateNewApp("commodo", "cmmdo")

	msg := types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetOneID, sdk.NewCoin("uasset1", newInt(300)), poolOneID, appOneID)
	msgLend2 := types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetTwoID, sdk.NewCoin("uasset2", newInt(10000000000)), poolOneID, appOneID)

	msg3 := types.NewMsgFundModuleAccounts(1, assetOneID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset1", newInt(10000000000)))
	msg4 := types.NewMsgFundModuleAccounts(1, assetTwoID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset2", newInt(10000000000)))
	msg5 := types.NewMsgFundModuleAccounts(1, assetThreeID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset3", newInt(120000000)))
	// msg6 := types.NewMsgFundModuleAccounts("osmo", assetThreeID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset3", newInt(10000000000)))
	msg7 := types.NewMsgFundModuleAccounts(2, assetOneID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset1", newInt(10000000000)))
	msg8 := types.NewMsgFundModuleAccounts(2, assetFourID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset4", newInt(10000000000)))

	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset1", newInt(100000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset2", newInt(100000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset3", newInt(100000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset4", newInt(100000000000))))
	_, _ = s.msgServer.Lend(sdk.WrapSDKContext(s.ctx), msg)
	_, _ = s.msgServer.Lend(sdk.WrapSDKContext(s.ctx), msgLend2)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg3)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg4)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg5)
	//_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg6)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg7)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg8)

	msg2 := types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, 3, false, sdk.NewCoin("ucasset1", newInt(100)), sdk.NewCoin("uasset2", newInt(10)))
	_, _ = s.msgServer.Borrow(sdk.WrapSDKContext(s.ctx), msg2)

	testCases := []struct {
		Name               string
		Msg                types.MsgRepay
		ExpErr             error
		ExpResp            *types.MsgRepayResponse
		QueryResponseIndex uint64
		QueryResponse      *types.MsgRepay
		AvailableBalance   sdk.Coins
	}{
		{
			Name:               "Borrow Not Found",
			Msg:                *types.NewMsgRepay("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 2, sdk.NewCoin("uasset2", newInt(100))),
			ExpErr:             types.ErrBorrowNotFound,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "invalid offer coin amount",
			Msg:                *types.NewMsgRepay("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, sdk.NewCoin("uasset3", newInt(100))),
			ExpErr:             types.ErrBadOfferCoinAmount,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "invalid repayment",
			Msg:                *types.NewMsgRepay("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, sdk.NewCoin("uasset2", newInt(100))),
			ExpErr:             types.ErrInvalidRepayment,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "different user",
			Msg:                *types.NewMsgRepay("cosmos14edpcw6ptcqd2vct9rkjf7lgyvrlwdtd0rqrtx", 1, sdk.NewCoin("uasset2", newInt(100))),
			ExpErr:             types.ErrLendAccessUnauthorized,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "success valid case",
			Msg:                *types.NewMsgRepay("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, sdk.NewCoin("uasset2", newInt(5))),
			ExpErr:             nil,
			ExpResp:            &types.MsgRepayResponse{},
			QueryResponseIndex: 0,
			QueryResponse: &types.MsgRepay{
				Borrower: "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
				BorrowId: 1,
				Amount:   sdk.NewCoin("uasset2", newInt(5)),
			},
			AvailableBalance: sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "success valid case",
			Msg:                *types.NewMsgRepay("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, sdk.NewCoin("uasset2", newInt(5))),
			ExpErr:             nil,
			ExpResp:            &types.MsgRepayResponse{},
			QueryResponseIndex: 0,
			QueryResponse: &types.MsgRepay{
				Borrower: "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
				BorrowId: 1,
				Amount:   sdk.NewCoin("uasset2", newInt(5)),
			},
			AvailableBalance: sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
	}
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.Name, func() {
			// add funds to acount for valid case
			if tc.ExpErr == nil {
				s.fundAddr(sdk.MustAccAddressFromBech32(tc.Msg.Borrower), sdk.NewCoins(sdk.NewCoin("uasset1", sdk.NewIntFromUint64(300))))
				s.fundAddr(sdk.MustAccAddressFromBech32(tc.Msg.Borrower), sdk.NewCoins(sdk.NewCoin("ucasset2", sdk.NewIntFromUint64(1000000000))))
			}

			ctx := sdk.WrapSDKContext(s.ctx)
			resp, err := s.msgServer.Repay(ctx, &tc.Msg)
			if tc.ExpErr != nil {
				s.Require().Error(err)
				s.Require().EqualError(err, tc.ExpErr.Error())
				s.Require().Equal(tc.ExpResp, resp)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(resp)
				s.Require().Equal(tc.ExpResp, resp)

				// availableBalances := s.getBalances(sdk.MustAccAddressFromBech32(tc.Msg.Lender))
				// s.Require().True(tc.AvailableBalance.IsEqual(availableBalances))
			}
		})
	}
}

func (s *KeeperTestSuite) TestMsgDepositBorrow() {
	assetOneID := s.CreateNewAsset("ASSETONE", "uasset1", 1000000)
	assetTwoID := s.CreateNewAsset("ASSETTWO", "uasset2", 2000000)
	assetThreeID := s.CreateNewAsset("ASSETTHREE", "uasset3", 2000000)
	assetFourID := s.CreateNewAsset("ASSETFOUR", "uasset4", 2000000)
	cAssetOneID := s.CreateNewAsset("CASSETONE", "ucasset1", 1000000)
	cAssetTwoID := s.CreateNewAsset("CASSETTWO", "ucasset2", 2000000)
	cAssetThreeID := s.CreateNewAsset("CASSETTHRE", "ucasset3", 2000000)
	cAssetFourID := s.CreateNewAsset("CASSETFOUR", "ucasset4", 2000000)

	var (
		assetDataPoolOne []*types.AssetDataPoolMapping
		assetDataPoolTwo []*types.AssetDataPoolMapping
	)
	assetDataPoolOneAssetOne := &types.AssetDataPoolMapping{
		AssetID:          assetOneID,
		AssetTransitType: 3,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolOneAssetTwo := &types.AssetDataPoolMapping{
		AssetID:          assetTwoID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(1000000000000000000),
	}
	assetDataPoolOneAssetThree := &types.AssetDataPoolMapping{
		AssetID:          assetThreeID,
		AssetTransitType: 2,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolTwoAssetFour := &types.AssetDataPoolMapping{
		AssetID:          assetFourID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(3000000000000000000),
	}
	assetDataPoolOne = append(assetDataPoolOne, assetDataPoolOneAssetOne, assetDataPoolOneAssetTwo, assetDataPoolOneAssetThree)
	assetDataPoolTwo = append(assetDataPoolOne, assetDataPoolTwoAssetFour, assetDataPoolOneAssetOne, assetDataPoolOneAssetThree)

	poolOneID := s.CreateNewPool("cmdx", "CMDX-ATOM-CMST", assetDataPoolOne)
	poolTwoID := s.CreateNewPool("osmo", "OSMO-ATOM-CMST", assetDataPoolTwo)

	s.AddAssetRatesStats(assetThreeID, newDec("0.8"), newDec("0.002"), newDec("0.06"), newDec("0.6"), true, newDec("0.04"), newDec("0.04"), newDec("0.06"), newDec("0.8"), newDec("0.85"), newDec("0.025"), newDec("0.025"), newDec("0.1"), cAssetThreeID)
	s.AddAssetRatesStats(assetOneID, newDec("0.75"), newDec("0.002"), newDec("0.07"), newDec("1.25"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.7"), newDec("0.75"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetOneID)
	s.AddAssetRatesStats(assetFourID, newDec("0.65"), newDec("0.002"), newDec("0.08"), newDec("1.5"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.6"), newDec("0.65"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetFourID)
	s.AddAssetRatesStats(assetTwoID, newDec("0.5"), newDec("0.002"), newDec("0.08"), newDec("2.0"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.5"), newDec("0.55"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetTwoID)

	pairOneID := s.AddExtendedLendPair(assetTwoID, assetThreeID, false, poolOneID, 1000000)
	pairTwoID := s.AddExtendedLendPair(assetTwoID, assetOneID, false, poolOneID, 1000000)
	pairThreeID := s.AddExtendedLendPair(assetOneID, assetTwoID, false, poolOneID, 1000000)
	pairFourID := s.AddExtendedLendPair(assetOneID, assetThreeID, false, poolOneID, 1000000)
	pairFiveID := s.AddExtendedLendPair(assetThreeID, assetTwoID, false, poolOneID, 1000000)
	pairSixID := s.AddExtendedLendPair(assetThreeID, assetOneID, false, poolOneID, 1000000)
	pairSevenID := s.AddExtendedLendPair(assetFourID, assetThreeID, false, poolTwoID, 1000000)
	pairEightID := s.AddExtendedLendPair(assetFourID, assetOneID, false, poolTwoID, 1000000)
	pairNineID := s.AddExtendedLendPair(assetOneID, assetFourID, false, poolTwoID, 1000000)
	pairTenID := s.AddExtendedLendPair(assetOneID, assetThreeID, false, poolTwoID, 1000000)
	pairElevenID := s.AddExtendedLendPair(assetThreeID, assetFourID, false, poolTwoID, 1000000)
	pairTwelveID := s.AddExtendedLendPair(assetThreeID, assetOneID, false, poolTwoID, 1000000)
	pairThirteenID := s.AddExtendedLendPair(assetTwoID, assetFourID, true, poolTwoID, 1000000)
	pairFourteenID := s.AddExtendedLendPair(assetThreeID, assetFourID, true, poolTwoID, 1000000)
	pairFifteenID := s.AddExtendedLendPair(assetOneID, assetFourID, true, poolTwoID, 1000000)
	pairSixteenID := s.AddExtendedLendPair(assetFourID, assetTwoID, true, poolOneID, 1000000)
	pairSeventeenID := s.AddExtendedLendPair(assetThreeID, assetTwoID, true, poolOneID, 1000000)
	pairEighteenID := s.AddExtendedLendPair(assetOneID, assetTwoID, true, poolOneID, 1000000)

	s.AddAssetToPair(assetOneID, poolOneID, []uint64{pairThreeID, pairFourID, pairFifteenID})
	s.AddAssetToPair(assetTwoID, poolOneID, []uint64{pairOneID, pairTwoID, pairThirteenID})
	s.AddAssetToPair(assetThreeID, poolOneID, []uint64{pairFiveID, pairSixID, pairFourteenID})
	s.AddAssetToPair(assetFourID, poolTwoID, []uint64{pairSevenID, pairEightID, pairSixteenID})
	s.AddAssetToPair(assetOneID, poolTwoID, []uint64{pairNineID, pairTenID, pairEighteenID})
	s.AddAssetToPair(assetThreeID, poolTwoID, []uint64{pairElevenID, pairTwelveID, pairSeventeenID})

	appOneID := s.CreateNewApp("commodo", "cmmdo")

	msg := types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetOneID, sdk.NewCoin("uasset1", newInt(300)), poolOneID, appOneID)
	msgLend2 := types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetTwoID, sdk.NewCoin("uasset2", newInt(10000000000)), poolOneID, appOneID)

	msg3 := types.NewMsgFundModuleAccounts(1, assetOneID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset1", newInt(10000000000)))
	msg4 := types.NewMsgFundModuleAccounts(1, assetTwoID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset2", newInt(10000000000)))
	msg5 := types.NewMsgFundModuleAccounts(1, assetThreeID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset3", newInt(120000000)))
	// msg6 := types.NewMsgFundModuleAccounts("osmo", assetThreeID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset3", newInt(10000000000)))
	msg7 := types.NewMsgFundModuleAccounts(2, assetOneID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset1", newInt(10000000000)))
	msg8 := types.NewMsgFundModuleAccounts(2, assetFourID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset4", newInt(10000000000)))

	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset1", newInt(100000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset2", newInt(100000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset3", newInt(100000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset4", newInt(100000000000))))
	_, _ = s.msgServer.Lend(sdk.WrapSDKContext(s.ctx), msg)
	_, _ = s.msgServer.Lend(sdk.WrapSDKContext(s.ctx), msgLend2)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg3)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg4)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg5)
	//_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg6)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg7)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg8)

	msg2 := types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, 3, false, sdk.NewCoin("ucasset1", newInt(100)), sdk.NewCoin("uasset2", newInt(1)))
	msgBorrow2 := types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, 15, false, sdk.NewCoin("ucasset1", newInt(100)), sdk.NewCoin("uasset4", newInt(1)))

	_, _ = s.msgServer.Borrow(sdk.WrapSDKContext(s.ctx), msg2)
	_, _ = s.msgServer.Borrow(sdk.WrapSDKContext(s.ctx), msgBorrow2)

	testCases := []struct {
		Name               string
		Msg                types.MsgDepositBorrow
		ExpErr             error
		ExpResp            *types.MsgDepositBorrowResponse
		QueryResponseIndex uint64
		QueryResponse      *types.MsgDepositBorrow
		AvailableBalance   sdk.Coins
	}{
		{
			Name:               "Borrow Not Found",
			Msg:                *types.NewMsgDepositBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 12, sdk.NewCoin("ucasset1", newInt(100))),
			ExpErr:             types.ErrBorrowNotFound,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "invalid offer coin amount",
			Msg:                *types.NewMsgDepositBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, sdk.NewCoin("uasset3", newInt(100))),
			ExpErr:             errorsmod.Wrap(types.ErrBadOfferCoinAmount, "uasset3"),
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "Available To Borrow Insufficient",
			Msg:                *types.NewMsgDepositBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, sdk.NewCoin("ucasset1", newInt(201))),
			ExpErr:             types.ErrAvailableToBorrowInsufficient,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "success valid case",
			Msg:                *types.NewMsgDepositBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, sdk.NewCoin("ucasset1", newInt(5))),
			ExpErr:             nil,
			ExpResp:            &types.MsgDepositBorrowResponse{},
			QueryResponseIndex: 0,
			QueryResponse: &types.MsgDepositBorrow{
				Borrower: "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
				BorrowId: 1,
				Amount:   sdk.NewCoin("uasset2", newInt(5)),
			},
			AvailableBalance: sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "success valid case 2",
			Msg:                *types.NewMsgDepositBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 2, sdk.NewCoin("ucasset1", newInt(5))),
			ExpErr:             nil,
			ExpResp:            &types.MsgDepositBorrowResponse{},
			QueryResponseIndex: 0,
			QueryResponse: &types.MsgDepositBorrow{
				Borrower: "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
				BorrowId: 2,
				Amount:   sdk.NewCoin("uasset1", newInt(5)),
			},
			AvailableBalance: sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
	}
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.Name, func() {
			// add funds to acount for valid case
			if tc.ExpErr == nil {
				s.fundAddr(sdk.MustAccAddressFromBech32(tc.Msg.Borrower), sdk.NewCoins(sdk.NewCoin("uasset1", sdk.NewIntFromUint64(300))))
				s.fundAddr(sdk.MustAccAddressFromBech32(tc.Msg.Borrower), sdk.NewCoins(sdk.NewCoin("ucasset2", sdk.NewIntFromUint64(1000000000))))
			}

			ctx := sdk.WrapSDKContext(s.ctx)
			resp, err := s.msgServer.DepositBorrow(ctx, &tc.Msg)
			if tc.ExpErr != nil {
				s.Require().Error(err)
				s.Require().EqualError(err, tc.ExpErr.Error())
				s.Require().Equal(tc.ExpResp, resp)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(resp)
				s.Require().Equal(tc.ExpResp, resp)

				// availableBalances := s.getBalances(sdk.MustAccAddressFromBech32(tc.Msg.Lender))
				// s.Require().True(tc.AvailableBalance.IsEqual(availableBalances))
			}
		})
	}
}

func (s *KeeperTestSuite) TestMsgDraw() {
	assetOneID := s.CreateNewAsset("ASSETONE", "uasset1", 1000000)
	assetTwoID := s.CreateNewAsset("ASSETTWO", "uasset2", 2000000)
	assetThreeID := s.CreateNewAsset("ASSETTHREE", "uasset3", 2000000)
	assetFourID := s.CreateNewAsset("ASSETFOUR", "uasset4", 2000000)
	cAssetOneID := s.CreateNewAsset("CASSETONE", "ucasset1", 1000000)
	cAssetTwoID := s.CreateNewAsset("CASSETTWO", "ucasset2", 2000000)
	cAssetThreeID := s.CreateNewAsset("CASSETTHRE", "ucasset3", 2000000)
	cAssetFourID := s.CreateNewAsset("CASSETFOUR", "ucasset4", 2000000)

	var (
		assetDataPoolOne []*types.AssetDataPoolMapping
		assetDataPoolTwo []*types.AssetDataPoolMapping
	)
	assetDataPoolOneAssetOne := &types.AssetDataPoolMapping{
		AssetID:          assetOneID,
		AssetTransitType: 3,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolOneAssetTwo := &types.AssetDataPoolMapping{
		AssetID:          assetTwoID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(1000000000000000000),
	}
	assetDataPoolOneAssetThree := &types.AssetDataPoolMapping{
		AssetID:          assetThreeID,
		AssetTransitType: 2,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolTwoAssetFour := &types.AssetDataPoolMapping{
		AssetID:          assetFourID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(3000000000000000000),
	}

	assetDataPoolOne = append(assetDataPoolOne, assetDataPoolOneAssetOne, assetDataPoolOneAssetTwo, assetDataPoolOneAssetThree)
	assetDataPoolTwo = append(assetDataPoolOne, assetDataPoolTwoAssetFour, assetDataPoolOneAssetOne, assetDataPoolOneAssetThree)

	poolOneID := s.CreateNewPool("cmdx", "CMDX-ATOM-CMST", assetDataPoolOne)
	poolTwoID := s.CreateNewPool("osmo", "OSMO-ATOM-CMST", assetDataPoolTwo)

	s.AddAssetRatesStats(assetThreeID, newDec("0.8"), newDec("0.002"), newDec("0.06"), newDec("0.6"), true, newDec("0.04"), newDec("0.04"), newDec("0.06"), newDec("0.8"), newDec("0.85"), newDec("0.025"), newDec("0.025"), newDec("0.1"), cAssetThreeID)
	s.AddAssetRatesStats(assetOneID, newDec("0.75"), newDec("0.002"), newDec("0.07"), newDec("1.25"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.7"), newDec("0.75"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetOneID)
	s.AddAssetRatesStats(assetFourID, newDec("0.65"), newDec("0.002"), newDec("0.08"), newDec("1.5"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.6"), newDec("0.65"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetFourID)
	s.AddAssetRatesStats(assetTwoID, newDec("0.5"), newDec("0.002"), newDec("0.08"), newDec("2.0"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.5"), newDec("0.55"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetTwoID)

	pairOneID := s.AddExtendedLendPair(assetTwoID, assetThreeID, false, poolOneID, 1000000)
	pairTwoID := s.AddExtendedLendPair(assetTwoID, assetOneID, false, poolOneID, 1000000)
	pairThreeID := s.AddExtendedLendPair(assetOneID, assetTwoID, false, poolOneID, 1000000)
	pairFourID := s.AddExtendedLendPair(assetOneID, assetThreeID, false, poolOneID, 1000000)
	pairFiveID := s.AddExtendedLendPair(assetThreeID, assetTwoID, false, poolOneID, 1000000)
	pairSixID := s.AddExtendedLendPair(assetThreeID, assetOneID, false, poolOneID, 1000000)
	pairSevenID := s.AddExtendedLendPair(assetFourID, assetThreeID, false, poolTwoID, 1000000)
	pairEightID := s.AddExtendedLendPair(assetFourID, assetOneID, false, poolTwoID, 1000000)
	pairNineID := s.AddExtendedLendPair(assetOneID, assetFourID, false, poolTwoID, 1000000)
	pairTenID := s.AddExtendedLendPair(assetOneID, assetThreeID, false, poolTwoID, 1000000)
	pairElevenID := s.AddExtendedLendPair(assetThreeID, assetFourID, false, poolTwoID, 1000000)
	pairTwelveID := s.AddExtendedLendPair(assetThreeID, assetOneID, false, poolTwoID, 1000000)
	pairThirteenID := s.AddExtendedLendPair(assetTwoID, assetFourID, true, poolTwoID, 1000000)
	pairFourteenID := s.AddExtendedLendPair(assetThreeID, assetFourID, true, poolTwoID, 1000000)
	pairFifteenID := s.AddExtendedLendPair(assetOneID, assetFourID, true, poolTwoID, 1000000)
	pairSixteenID := s.AddExtendedLendPair(assetFourID, assetTwoID, true, poolOneID, 1000000)
	pairSeventeenID := s.AddExtendedLendPair(assetThreeID, assetTwoID, true, poolOneID, 1000000)
	pairEighteenID := s.AddExtendedLendPair(assetOneID, assetTwoID, true, poolOneID, 1000000)

	s.AddAssetToPair(assetOneID, poolOneID, []uint64{pairThreeID, pairFourID, pairFifteenID})
	s.AddAssetToPair(assetTwoID, poolOneID, []uint64{pairOneID, pairTwoID, pairThirteenID})
	s.AddAssetToPair(assetThreeID, poolOneID, []uint64{pairFiveID, pairSixID, pairFourteenID})
	s.AddAssetToPair(assetFourID, poolTwoID, []uint64{pairSevenID, pairEightID, pairSixteenID})
	s.AddAssetToPair(assetOneID, poolTwoID, []uint64{pairNineID, pairTenID, pairEighteenID})
	s.AddAssetToPair(assetThreeID, poolTwoID, []uint64{pairElevenID, pairTwelveID, pairSeventeenID})

	appOneID := s.CreateNewApp("commodo", "cmmdo")

	msg := types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetOneID, sdk.NewCoin("uasset1", newInt(300)), poolOneID, appOneID)
	msgLend2 := types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetTwoID, sdk.NewCoin("uasset2", newInt(10000000000)), poolOneID, appOneID)

	msg3 := types.NewMsgFundModuleAccounts(1, assetOneID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset1", newInt(10000000000)))
	msg4 := types.NewMsgFundModuleAccounts(1, assetTwoID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset2", newInt(10000000000)))
	msg5 := types.NewMsgFundModuleAccounts(1, assetThreeID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset3", newInt(120000000)))
	// msg6 := types.NewMsgFundModuleAccounts(2", assetThreeID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset3", newInt(10000000000)))
	msg7 := types.NewMsgFundModuleAccounts(2, assetOneID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset1", newInt(10000000000)))
	msg8 := types.NewMsgFundModuleAccounts(2, assetFourID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset4", newInt(10000000000)))

	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset1", newInt(100000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset2", newInt(100000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset3", newInt(100000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset4", newInt(100000000000))))
	_, _ = s.msgServer.Lend(sdk.WrapSDKContext(s.ctx), msg)
	_, _ = s.msgServer.Lend(sdk.WrapSDKContext(s.ctx), msgLend2)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg3)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg4)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg5)
	//_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg6)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg7)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg8)

	msg2 := types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, 3, false, sdk.NewCoin("ucasset1", newInt(100)), sdk.NewCoin("uasset2", newInt(10)))
	_, _ = s.msgServer.Borrow(sdk.WrapSDKContext(s.ctx), msg2)

	testCases := []struct {
		Name               string
		Msg                types.MsgDraw
		ExpErr             error
		ExpResp            *types.MsgDrawResponse
		QueryResponseIndex uint64
		QueryResponse      *types.MsgDraw
		AvailableBalance   sdk.Coins
	}{
		{
			Name:               "Borrow Not Found",
			Msg:                *types.NewMsgDraw("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 2, sdk.NewCoin("uasset2", newInt(100))),
			ExpErr:             types.ErrBorrowNotFound,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "invalid offer coin amount",
			Msg:                *types.NewMsgDraw("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, sdk.NewCoin("uasset3", newInt(100))),
			ExpErr:             types.ErrBadOfferCoinAmount,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "Error Invalid Collaterallization Ratio",
			Msg:                *types.NewMsgDraw("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, sdk.NewCoin("uasset2", newInt(100))),
			ExpErr:             types.ErrorInvalidCollateralizationRatio,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "success valid case",
			Msg:                *types.NewMsgDraw("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, sdk.NewCoin("uasset2", newInt(1))),
			ExpErr:             nil,
			ExpResp:            &types.MsgDrawResponse{},
			QueryResponseIndex: 0,
			QueryResponse: &types.MsgDraw{
				Borrower: "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
				BorrowId: 1,
				Amount:   sdk.NewCoin("uasset2", newInt(1)),
			},
			AvailableBalance: sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
	}
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.Name, func() {
			// add funds to acount for valid case
			if tc.ExpErr == nil {
				s.fundAddr(sdk.MustAccAddressFromBech32(tc.Msg.Borrower), sdk.NewCoins(sdk.NewCoin("uasset1", sdk.NewIntFromUint64(300))))
				s.fundAddr(sdk.MustAccAddressFromBech32(tc.Msg.Borrower), sdk.NewCoins(sdk.NewCoin("ucasset2", sdk.NewIntFromUint64(1000000000))))
			}

			ctx := sdk.WrapSDKContext(s.ctx)
			resp, err := s.msgServer.Draw(ctx, &tc.Msg)
			if tc.ExpErr != nil {
				s.Require().Error(err)
				s.Require().EqualError(err, tc.ExpErr.Error())
				s.Require().Equal(tc.ExpResp, resp)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(resp)
				s.Require().Equal(tc.ExpResp, resp)

				// availableBalances := s.getBalances(sdk.MustAccAddressFromBech32(tc.Msg.Lender))
				// s.Require().True(tc.AvailableBalance.IsEqual(availableBalances))
			}
		})
	}
}

func (s *KeeperTestSuite) TestMsgCloseBorrow() {
	assetOneID := s.CreateNewAsset("ASSETONE", "uasset1", 1000000)
	assetTwoID := s.CreateNewAsset("ASSETTWO", "uasset2", 2000000)
	assetThreeID := s.CreateNewAsset("ASSETTHREE", "uasset3", 2000000)
	assetFourID := s.CreateNewAsset("ASSETFOUR", "uasset4", 2000000)
	cAssetOneID := s.CreateNewAsset("CASSETONE", "ucasset1", 1000000)
	cAssetTwoID := s.CreateNewAsset("CASSETTWO", "ucasset2", 2000000)
	cAssetThreeID := s.CreateNewAsset("CASSETTHRE", "ucasset3", 2000000)
	cAssetFourID := s.CreateNewAsset("CASSETFOUR", "ucasset4", 2000000)

	var (
		assetDataPoolOne []*types.AssetDataPoolMapping
		assetDataPoolTwo []*types.AssetDataPoolMapping
	)
	assetDataPoolOneAssetOne := &types.AssetDataPoolMapping{
		AssetID:          assetOneID,
		AssetTransitType: 3,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolOneAssetTwo := &types.AssetDataPoolMapping{
		AssetID:          assetTwoID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(1000000000000000000),
	}
	assetDataPoolOneAssetThree := &types.AssetDataPoolMapping{
		AssetID:          assetThreeID,
		AssetTransitType: 2,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolTwoAssetFour := &types.AssetDataPoolMapping{
		AssetID:          assetFourID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(3000000000000000000),
	}

	assetDataPoolOne = append(assetDataPoolOne, assetDataPoolOneAssetOne, assetDataPoolOneAssetTwo, assetDataPoolOneAssetThree)
	assetDataPoolTwo = append(assetDataPoolOne, assetDataPoolTwoAssetFour, assetDataPoolOneAssetOne, assetDataPoolOneAssetThree)

	poolOneID := s.CreateNewPool("cmdx", "CMDX-ATOM-CMST", assetDataPoolOne)
	poolTwoID := s.CreateNewPool("osmo", "OSMO-ATOM-CMST", assetDataPoolTwo)

	s.AddAssetRatesStats(assetThreeID, newDec("0.8"), newDec("0.002"), newDec("0.06"), newDec("0.6"), true, newDec("0.04"), newDec("0.04"), newDec("0.06"), newDec("0.8"), newDec("0.85"), newDec("0.025"), newDec("0.025"), newDec("0.1"), cAssetThreeID)
	s.AddAssetRatesStats(assetOneID, newDec("0.75"), newDec("0.002"), newDec("0.07"), newDec("1.25"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.7"), newDec("0.75"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetOneID)
	s.AddAssetRatesStats(assetFourID, newDec("0.65"), newDec("0.002"), newDec("0.08"), newDec("1.5"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.6"), newDec("0.65"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetFourID)
	s.AddAssetRatesStats(assetTwoID, newDec("0.5"), newDec("0.002"), newDec("0.08"), newDec("2.0"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.5"), newDec("0.55"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetTwoID)

	pairOneID := s.AddExtendedLendPair(assetTwoID, assetThreeID, false, poolOneID, 1000000)
	pairTwoID := s.AddExtendedLendPair(assetTwoID, assetOneID, false, poolOneID, 1000000)
	pairThreeID := s.AddExtendedLendPair(assetOneID, assetTwoID, false, poolOneID, 1000000)
	pairFourID := s.AddExtendedLendPair(assetOneID, assetThreeID, false, poolOneID, 1000000)
	pairFiveID := s.AddExtendedLendPair(assetThreeID, assetTwoID, false, poolOneID, 1000000)
	pairSixID := s.AddExtendedLendPair(assetThreeID, assetOneID, false, poolOneID, 1000000)
	pairSevenID := s.AddExtendedLendPair(assetFourID, assetThreeID, false, poolTwoID, 1000000)
	pairEightID := s.AddExtendedLendPair(assetFourID, assetOneID, false, poolTwoID, 1000000)
	pairNineID := s.AddExtendedLendPair(assetOneID, assetFourID, false, poolTwoID, 1000000)
	pairTenID := s.AddExtendedLendPair(assetOneID, assetThreeID, false, poolTwoID, 1000000)
	pairElevenID := s.AddExtendedLendPair(assetThreeID, assetFourID, false, poolTwoID, 1000000)
	pairTwelveID := s.AddExtendedLendPair(assetThreeID, assetOneID, false, poolTwoID, 1000000)
	pairThirteenID := s.AddExtendedLendPair(assetTwoID, assetFourID, true, poolTwoID, 1000000)
	pairFourteenID := s.AddExtendedLendPair(assetThreeID, assetFourID, true, poolTwoID, 1000000)
	pairFifteenID := s.AddExtendedLendPair(assetOneID, assetFourID, true, poolTwoID, 1000000)
	pairSixteenID := s.AddExtendedLendPair(assetFourID, assetTwoID, true, poolOneID, 1000000)
	pairSeventeenID := s.AddExtendedLendPair(assetThreeID, assetTwoID, true, poolOneID, 1000000)
	pairEighteenID := s.AddExtendedLendPair(assetOneID, assetTwoID, true, poolOneID, 1000000)

	s.AddAssetToPair(assetOneID, poolOneID, []uint64{pairThreeID, pairFourID, pairFifteenID})
	s.AddAssetToPair(assetTwoID, poolOneID, []uint64{pairOneID, pairTwoID, pairThirteenID})
	s.AddAssetToPair(assetThreeID, poolOneID, []uint64{pairFiveID, pairSixID, pairFourteenID})
	s.AddAssetToPair(assetFourID, poolTwoID, []uint64{pairSevenID, pairEightID, pairSixteenID})
	s.AddAssetToPair(assetOneID, poolTwoID, []uint64{pairNineID, pairTenID, pairEighteenID})
	s.AddAssetToPair(assetThreeID, poolTwoID, []uint64{pairElevenID, pairTwelveID, pairSeventeenID})

	appOneID := s.CreateNewApp("commodo", "cmmdo")

	msg := types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetOneID, sdk.NewCoin("uasset1", newInt(300)), poolOneID, appOneID)
	msgLend2 := types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetTwoID, sdk.NewCoin("uasset2", newInt(10000000000)), poolOneID, appOneID)

	msg3 := types.NewMsgFundModuleAccounts(1, assetOneID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset1", newInt(10000000000)))
	msg4 := types.NewMsgFundModuleAccounts(1, assetTwoID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset2", newInt(10000000000)))
	msg5 := types.NewMsgFundModuleAccounts(1, assetThreeID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset3", newInt(120000000)))
	// msg6 := types.NewMsgFundModuleAccounts(2", assetThreeID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset3", newInt(10000000000)))
	msg7 := types.NewMsgFundModuleAccounts(2, assetOneID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset1", newInt(10000000000)))
	msg8 := types.NewMsgFundModuleAccounts(2, assetFourID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset4", newInt(10000000000)))

	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset1", newInt(100000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset2", newInt(100000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset3", newInt(100000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset4", newInt(100000000000))))
	_, _ = s.msgServer.Lend(sdk.WrapSDKContext(s.ctx), msg)
	_, _ = s.msgServer.Lend(sdk.WrapSDKContext(s.ctx), msgLend2)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg3)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg4)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg5)
	//_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg6)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg7)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg8)

	msg2 := types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, 3, false, sdk.NewCoin("ucasset1", newInt(100)), sdk.NewCoin("uasset2", newInt(1)))
	borrowMsg3 := types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, 15, false, sdk.NewCoin("ucasset1", newInt(100)), sdk.NewCoin("uasset4", newInt(1)))
	_, _ = s.msgServer.Borrow(sdk.WrapSDKContext(s.ctx), msg2)
	_, _ = s.msgServer.Borrow(sdk.WrapSDKContext(s.ctx), borrowMsg3)

	testCases := []struct {
		Name               string
		Msg                types.MsgCloseBorrow
		ExpErr             error
		ExpResp            *types.MsgCloseBorrowResponse
		QueryResponseIndex uint64
		QueryResponse      *types.MsgCloseBorrow
		AvailableBalance   sdk.Coins
	}{
		{
			Name:               "Borrow Not Found",
			Msg:                *types.NewMsgCloseBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 12),
			ExpErr:             types.ErrBorrowNotFound,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "Different user",
			Msg:                *types.NewMsgCloseBorrow("cosmos14edpcw6ptcqd2vct9rkjf7lgyvrlwdtd0rqrtx", 1),
			ExpErr:             types.ErrLendAccessUnauthorized,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "success valid case",
			Msg:                *types.NewMsgCloseBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1),
			ExpErr:             nil,
			ExpResp:            &types.MsgCloseBorrowResponse{},
			QueryResponseIndex: 0,
			QueryResponse: &types.MsgCloseBorrow{
				Borrower: "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
				BorrowId: 1,
			},
			AvailableBalance: sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "success valid case 2",
			Msg:                *types.NewMsgCloseBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 2),
			ExpErr:             nil,
			ExpResp:            &types.MsgCloseBorrowResponse{},
			QueryResponseIndex: 0,
			QueryResponse: &types.MsgCloseBorrow{
				Borrower: "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
				BorrowId: 1,
			},
			AvailableBalance: sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
	}
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.Name, func() {
			// add funds to acount for valid case
			if tc.ExpErr == nil {
				s.fundAddr(sdk.MustAccAddressFromBech32(tc.Msg.Borrower), sdk.NewCoins(sdk.NewCoin("uasset1", sdk.NewIntFromUint64(300))))
				s.fundAddr(sdk.MustAccAddressFromBech32(tc.Msg.Borrower), sdk.NewCoins(sdk.NewCoin("ucasset2", sdk.NewIntFromUint64(1000000000))))
			}

			ctx := sdk.WrapSDKContext(s.ctx)
			resp, err := s.msgServer.CloseBorrow(ctx, &tc.Msg)
			if tc.ExpErr != nil {
				s.Require().Error(err)
				s.Require().EqualError(err, tc.ExpErr.Error())
				s.Require().Equal(tc.ExpResp, resp)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(resp)
				s.Require().Equal(tc.ExpResp, resp)

				// availableBalances := s.getBalances(sdk.MustAccAddressFromBech32(tc.Msg.Lender))
				// s.Require().True(tc.AvailableBalance.IsEqual(availableBalances))
			}
		})
	}
}

func (s *KeeperTestSuite) TestMsgBorrowAlternate() {
	assetOneID := s.CreateNewAsset("ASSETONE", "uasset1", 1000000)
	assetTwoID := s.CreateNewAsset("ASSETTWO", "uasset2", 2000000)
	assetThreeID := s.CreateNewAsset("ASSETTHREE", "uasset3", 2000000)
	assetFourID := s.CreateNewAsset("ASSETFOUR", "uasset4", 2000000)
	cAssetOneID := s.CreateNewAsset("CASSETONE", "ucasset1", 1000000)
	// cAssetTwoID := s.CreateNewAsset("CASSETTWO", "ucasset2", 2000000)
	cAssetThreeID := s.CreateNewAsset("CASSETTHRE", "ucasset3", 2000000)
	// cAssetFourID := s.CreateNewAsset("CASSETFOUR", "ucasset4", 2000000)

	var (
		assetDataPoolOne []*types.AssetDataPoolMapping
		assetDataPoolTwo []*types.AssetDataPoolMapping
	)
	assetDataPoolOneAssetOne := &types.AssetDataPoolMapping{
		AssetID:          assetOneID,
		AssetTransitType: 3,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolOneAssetTwo := &types.AssetDataPoolMapping{
		AssetID:          assetTwoID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(1000000000000000000),
	}
	assetDataPoolOneAssetThree := &types.AssetDataPoolMapping{
		AssetID:          assetThreeID,
		AssetTransitType: 2,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolTwoAssetFour := &types.AssetDataPoolMapping{
		AssetID:          assetFourID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(3000000000000000000),
	}

	assetDataPoolOne = append(assetDataPoolOne, assetDataPoolOneAssetOne, assetDataPoolOneAssetTwo, assetDataPoolOneAssetThree)
	assetDataPoolTwo = append(assetDataPoolOne, assetDataPoolTwoAssetFour, assetDataPoolOneAssetTwo, assetDataPoolOneAssetThree)

	poolOneID := s.CreateNewPool("cmdx", "CMDX-ATOM-CMST", assetDataPoolOne)
	poolTwoID := s.CreateNewPool("osmo", "OSMO-ATOM-CMST", assetDataPoolTwo)

	s.AddAssetRatesStats(assetThreeID, newDec("0.8"), newDec("0.002"), newDec("0.06"), newDec("0.6"), true, newDec("0.04"), newDec("0.04"), newDec("0.06"), newDec("0.8"), newDec("0.85"), newDec("0.025"), newDec("0.025"), newDec("0.1"), cAssetThreeID)
	s.AddAssetRatesStats(assetOneID, newDec("0.75"), newDec("0.002"), newDec("0.07"), newDec("1.25"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.7"), newDec("0.75"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetOneID)
	// s.AddAssetRatesStats(assetFourID, newDec("0.65"), newDec("0.002"), newDec("0.08"), newDec("1.5"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.6"), newDec("0.65"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetFourID)
	// s.AddAssetRatesStats(assetTwoID, newDec("0.5"), newDec("0.002"), newDec("0.08"), newDec("2.0"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.5"), newDec("0.55"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetTwoID)

	pairOneID := s.AddExtendedLendPair(assetTwoID, assetThreeID, false, poolOneID, 1000000)
	pairTwoID := s.AddExtendedLendPair(assetTwoID, assetOneID, false, poolOneID, 1000000)
	pairThreeID := s.AddExtendedLendPair(assetOneID, assetTwoID, false, poolOneID, 1000000)
	pairFourID := s.AddExtendedLendPair(assetOneID, assetThreeID, false, poolOneID, 1000000)
	pairFiveID := s.AddExtendedLendPair(assetThreeID, assetTwoID, false, poolOneID, 1000000)
	pairSixID := s.AddExtendedLendPair(assetThreeID, assetOneID, false, poolOneID, 1000000)
	pairSevenID := s.AddExtendedLendPair(assetFourID, assetThreeID, false, poolTwoID, 1000000)
	pairEightID := s.AddExtendedLendPair(assetFourID, assetOneID, false, poolTwoID, 1000000)
	pairNineID := s.AddExtendedLendPair(assetOneID, assetFourID, false, poolTwoID, 1000000)
	pairTenID := s.AddExtendedLendPair(assetOneID, assetThreeID, false, poolTwoID, 1000000)
	pairElevenID := s.AddExtendedLendPair(assetThreeID, assetFourID, false, poolTwoID, 1000000)
	pairTwelveID := s.AddExtendedLendPair(assetThreeID, assetOneID, false, poolTwoID, 1000000)
	pairThirteenID := s.AddExtendedLendPair(assetTwoID, assetFourID, true, poolTwoID, 1000000)
	pairFourteenID := s.AddExtendedLendPair(assetThreeID, assetFourID, true, poolTwoID, 1000000)
	pairFifteenID := s.AddExtendedLendPair(assetOneID, assetFourID, true, poolTwoID, 1000000)
	pairSixteenID := s.AddExtendedLendPair(assetFourID, assetTwoID, true, poolOneID, 1000000)
	pairSeventeenID := s.AddExtendedLendPair(assetThreeID, assetTwoID, true, poolOneID, 1000000)
	pairEighteenID := s.AddExtendedLendPair(assetOneID, assetTwoID, true, poolOneID, 1000000)

	s.AddAssetToPair(assetOneID, poolOneID, []uint64{pairThreeID, pairFourID, pairFifteenID})
	s.AddAssetToPair(assetTwoID, poolOneID, []uint64{pairOneID, pairTwoID, pairThirteenID})
	s.AddAssetToPair(assetThreeID, poolOneID, []uint64{pairFiveID, pairSixID, pairFourteenID})
	s.AddAssetToPair(assetFourID, poolTwoID, []uint64{pairSevenID, pairEightID, pairSixteenID})
	s.AddAssetToPair(assetOneID, poolTwoID, []uint64{pairNineID, pairTenID, pairEighteenID})
	s.AddAssetToPair(assetThreeID, poolTwoID, []uint64{pairElevenID, pairTwelveID, pairSeventeenID})

	appOneID := s.CreateNewApp("commodo", "cmmdo")
	appTwoID := s.CreateNewApp("cswap", "cswap")
	msg := types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetThreeID, sdk.NewCoin("uasset3", newInt(300)), poolOneID, appOneID)

	msg3 := types.NewMsgFundModuleAccounts(1, assetOneID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset1", newInt(10000000000)))
	msg4 := types.NewMsgFundModuleAccounts(1, assetTwoID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset2", newInt(10000000000)))
	msg5 := types.NewMsgFundModuleAccounts(1, assetThreeID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset3", newInt(120000000)))
	// msg6 := types.NewMsgFundModuleAccounts(2, assetThreeID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset3", newInt(10000000000)))
	msg7 := types.NewMsgFundModuleAccounts(2, assetOneID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset1", newInt(10000000000)))
	msg8 := types.NewMsgFundModuleAccounts(2, assetFourID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset4", newInt(10000000000)))

	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset1", newInt(100000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset2", newInt(100000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset3", newInt(100000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset4", newInt(100000000000))))

	_, _ = s.msgServer.Lend(sdk.WrapSDKContext(s.ctx), msg)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg3)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg4)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg5)
	//_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg6)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg7)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg8)

	testCases := []struct {
		Name               string
		Msg                types.MsgBorrowAlternate
		ExpErr             error
		ExpResp            *types.MsgBorrowAlternateResponse
		QueryResponseIndex uint64
		QueryResponse      *types.MsgBorrowAlternate
		AvailableBalance   sdk.Coins
	}{
		{
			Name:               "asset does not exist",
			Msg:                *types.NewMsgBorrowAlternate("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 10, poolOneID, sdk.NewCoin("uasset1", sdk.NewInt(100)), pairOneID, false, sdk.NewCoin("uasset1", sdk.NewInt(100)), appOneID),
			ExpErr:             types.ErrorAssetDoesNotExist,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "Pool Not Found",
			Msg:                *types.NewMsgBorrowAlternate("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetOneID, 3, sdk.NewCoin("uasset1", sdk.NewInt(100)), pairOneID, false, sdk.NewCoin("uasset1", sdk.NewInt(100)), appOneID),
			ExpErr:             types.ErrPoolNotFound,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "App Mapping Id does not exists",
			Msg:                *types.NewMsgBorrowAlternate("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetOneID, poolOneID, sdk.NewCoin("uasset1", sdk.NewInt(100)), pairOneID, false, sdk.NewCoin("uasset1", sdk.NewInt(100)), 3),
			ExpErr:             types.ErrorAppMappingDoesNotExist,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "App Mapping Id mismatch, use the correct App Mapping ID in request",
			Msg:                *types.NewMsgBorrowAlternate("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetOneID, poolOneID, sdk.NewCoin("uasset1", sdk.NewInt(100)), pairThreeID, false, sdk.NewCoin("uasset2", sdk.NewInt(10)), appTwoID),
			ExpErr:             types.ErrorAppMappingIDMismatch,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(100))),
		},
		{
			Name:               "invalid offer coin amount",
			Msg:                *types.NewMsgBorrowAlternate("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetOneID, poolOneID, sdk.NewCoin("uasset2", sdk.NewInt(100)), pairThreeID, false, sdk.NewCoin("uasset2", sdk.NewInt(10)), appOneID),
			ExpErr:             errorsmod.Wrapf(types.ErrBadOfferCoinAmount, "uasset2"),
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(100))),
		},
		//{
		//	Name:               "Duplicate lend Position",
		//	Msg:                *types.NewMsgBorrowAlternate("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetThreeID, poolOneID, sdk.NewCoin("uasset3", sdk.NewInt(100)), pairFiveID, false, sdk.NewCoin("uasset2", sdk.NewInt(10)), appOneID),
		//	ExpErr:             types.ErrorDuplicateLend,
		//	ExpResp:            nil,
		//	QueryResponseIndex: 0,
		//	QueryResponse:      nil,
		//	AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(100))),
		//},
		{
			Name:               "Asset Id not defined in the pool",
			Msg:                *types.NewMsgBorrowAlternate("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetFourID, poolOneID, sdk.NewCoin("uasset4", sdk.NewInt(100)), pairSevenID, false, sdk.NewCoin("uasset3", sdk.NewInt(10)), appOneID),
			ExpErr:             errorsmod.Wrap(types.ErrInvalidAssetIDForPool, "4"),
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(100))),
		},
		{
			Name:               "Asset Rates Stats not found",
			Msg:                *types.NewMsgBorrowAlternate("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetFourID, poolTwoID, sdk.NewCoin("uasset4", sdk.NewInt(100)), pairSevenID, false, sdk.NewCoin("uasset3", sdk.NewInt(10)), appOneID),
			ExpErr:             errorsmod.Wrap(types.ErrorAssetRatesParamsNotFound, "4"),
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(100))),
		},
		{
			Name:               "success valid case",
			Msg:                *types.NewMsgBorrowAlternate("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetOneID, poolOneID, sdk.NewCoin("uasset1", sdk.NewInt(100)), pairThreeID, false, sdk.NewCoin("uasset2", sdk.NewInt(10)), appOneID),
			ExpErr:             nil,
			ExpResp:            &types.MsgBorrowAlternateResponse{},
			QueryResponseIndex: 0,
			QueryResponse: &types.MsgBorrowAlternate{
				Lender:         "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
				AssetId:        assetOneID,
				PoolId:         poolOneID,
				AmountIn:       sdk.NewCoin("uasset1", sdk.NewInt(100)),
				PairId:         pairThreeID,
				IsStableBorrow: false,
				AmountOut:      sdk.NewCoin("uasset2", sdk.NewInt(10)),
				AppId:          appOneID,
			},
			// AvailableBalance: sdk.NewCoins(sdk.NewCoin("uasset2", newInt(90000000010))),
		},
	}
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.Name, func() {
			// add funds to acount for valid case
			if tc.ExpErr == nil {
				s.fundAddr(sdk.MustAccAddressFromBech32(tc.Msg.Lender), sdk.NewCoins(sdk.NewCoin("uasset1", tc.Msg.AmountIn.Amount)))
			}

			ctx := sdk.WrapSDKContext(s.ctx)
			resp, err := s.msgServer.BorrowAlternate(ctx, &tc.Msg)
			if tc.ExpErr != nil {
				s.Require().Error(err)
				s.Require().EqualError(err, tc.ExpErr.Error())
				s.Require().Equal(tc.ExpResp, resp)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(resp)
				s.Require().Equal(tc.ExpResp, resp)
			}
		})
	}
}

func (s *KeeperTestSuite) TestMsgInitializeLend() {
	assetOneID := s.CreateNewAsset("ASSETONE", "uasset1", 1000000)
	assetTwoID := s.CreateNewAsset("ASSETTWO", "uasset2", 2000000)
	assetThreeID := s.CreateNewAsset("ASSETTHREE", "uasset3", 2000000)
	assetFourID := s.CreateNewAsset("ASSETFOUR", "uasset4", 2000000)
	cAssetOneID := s.CreateNewAsset("CASSETONE", "ucasset1", 1000000)
	cAssetTwoID := s.CreateNewAsset("CASSETTWO", "ucasset2", 2000000)
	cAssetThreeID := s.CreateNewAsset("CASSETTHRE", "ucasset3", 2000000)
	cAssetFourID := s.CreateNewAsset("CASSETFOUR", "ucasset4", 2000000)

	var (
		assetDataPoolOne []*types.AssetDataPoolMapping
		assetDataPoolTwo []*types.AssetDataPoolMapping
	)
	assetDataPoolOneAssetOne := &types.AssetDataPoolMapping{
		AssetID:          assetOneID,
		AssetTransitType: 3,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolOneAssetTwo := &types.AssetDataPoolMapping{
		AssetID:          assetTwoID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(1000000000000000000),
	}
	assetDataPoolOneAssetThree := &types.AssetDataPoolMapping{
		AssetID:          assetThreeID,
		AssetTransitType: 2,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolTwoAssetFour := &types.AssetDataPoolMapping{
		AssetID:          assetFourID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(3000000000000000000),
	}

	assetDataPoolOne = append(assetDataPoolOne, assetDataPoolOneAssetOne, assetDataPoolOneAssetTwo, assetDataPoolOneAssetThree)
	assetDataPoolTwo = append(assetDataPoolOne, assetDataPoolTwoAssetFour, assetDataPoolOneAssetTwo, assetDataPoolOneAssetThree)

	poolOneID := s.CreateNewPool("cmdx", "CMDX-ATOM-CMST", assetDataPoolOne)
	poolTwoID := s.CreateNewPool("osmo", "OSMO-ATOM-CMST", assetDataPoolTwo)

	s.AddAssetRatesStats(assetThreeID, newDec("0.8"), newDec("0.002"), newDec("0.06"), newDec("0.6"), true, newDec("0.04"), newDec("0.04"), newDec("0.06"), newDec("0.8"), newDec("0.85"), newDec("0.025"), newDec("0.025"), newDec("0.1"), cAssetThreeID)
	s.AddAssetRatesStats(assetOneID, newDec("0.75"), newDec("0.002"), newDec("0.07"), newDec("1.25"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.7"), newDec("0.75"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetOneID)
	s.AddAssetRatesStats(assetFourID, newDec("0.65"), newDec("0.002"), newDec("0.08"), newDec("1.5"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.6"), newDec("0.65"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetFourID)
	s.AddAssetRatesStats(assetTwoID, newDec("0.5"), newDec("0.002"), newDec("0.08"), newDec("2.0"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.5"), newDec("0.55"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetTwoID)

	pairOneID := s.AddExtendedLendPair(assetTwoID, assetThreeID, false, poolOneID, 1000000)
	pairTwoID := s.AddExtendedLendPair(assetTwoID, assetOneID, false, poolOneID, 1000000)
	pairThreeID := s.AddExtendedLendPair(assetOneID, assetTwoID, false, poolOneID, 1000000)
	pairFourID := s.AddExtendedLendPair(assetOneID, assetThreeID, false, poolOneID, 1000000)
	pairFiveID := s.AddExtendedLendPair(assetThreeID, assetTwoID, false, poolOneID, 1000000)
	pairSixID := s.AddExtendedLendPair(assetThreeID, assetOneID, false, poolOneID, 1000000)
	pairSevenID := s.AddExtendedLendPair(assetFourID, assetThreeID, false, poolTwoID, 1000000)
	pairEightID := s.AddExtendedLendPair(assetFourID, assetOneID, false, poolTwoID, 1000000)
	pairNineID := s.AddExtendedLendPair(assetOneID, assetFourID, false, poolTwoID, 1000000)
	pairTenID := s.AddExtendedLendPair(assetOneID, assetThreeID, false, poolTwoID, 1000000)
	pairElevenID := s.AddExtendedLendPair(assetThreeID, assetFourID, false, poolTwoID, 1000000)
	pairTwelveID := s.AddExtendedLendPair(assetThreeID, assetOneID, false, poolTwoID, 1000000)
	pairThirteenID := s.AddExtendedLendPair(assetTwoID, assetFourID, true, poolTwoID, 1000000)
	pairFourteenID := s.AddExtendedLendPair(assetThreeID, assetFourID, true, poolTwoID, 1000000)
	pairFifteenID := s.AddExtendedLendPair(assetOneID, assetFourID, true, poolTwoID, 1000000)
	pairSixteenID := s.AddExtendedLendPair(assetFourID, assetTwoID, true, poolOneID, 1000000)
	pairSeventeenID := s.AddExtendedLendPair(assetThreeID, assetTwoID, true, poolOneID, 1000000)
	pairEighteenID := s.AddExtendedLendPair(assetOneID, assetTwoID, true, poolOneID, 1000000)

	s.AddAssetToPair(assetOneID, poolOneID, []uint64{pairThreeID, pairFourID, pairFifteenID})
	s.AddAssetToPair(assetTwoID, poolOneID, []uint64{pairOneID, pairTwoID, pairThirteenID})
	s.AddAssetToPair(assetThreeID, poolOneID, []uint64{pairFiveID, pairSixID, pairFourteenID})
	s.AddAssetToPair(assetFourID, poolTwoID, []uint64{pairSevenID, pairEightID, pairSixteenID})
	s.AddAssetToPair(assetOneID, poolTwoID, []uint64{pairNineID, pairTenID, pairEighteenID})
	s.AddAssetToPair(assetThreeID, poolTwoID, []uint64{pairElevenID, pairTwelveID, pairSeventeenID})

	s.CreateNewApp("commodo", "cmmdo")
	msg3 := types.NewMsgFundModuleAccounts(1, assetOneID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset1", newInt(10000000000)))
	msg4 := types.NewMsgFundModuleAccounts(1, assetTwoID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset2", newInt(10000000000)))
	msg5 := types.NewMsgFundModuleAccounts(1, assetThreeID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset3", newInt(120000000)))
	msg6 := types.NewMsgFundModuleAccounts(2, assetThreeID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset3", newInt(10000000000)))
	msg7 := types.NewMsgFundModuleAccounts(2, assetOneID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset1", newInt(10000000000)))
	msg8 := types.NewMsgFundModuleAccounts(2, assetFourID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset4", newInt(10000000000)))

	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset1", newInt(100000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset2", newInt(100000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset3", newInt(100000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset4", newInt(100000000000))))

	_, err := s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg3)
	s.Require().NoError(err)

	_, err = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg4)
	s.Require().NoError(err)

	_, err = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg5)
	s.Require().NoError(err)

	_, err = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg6)
	s.Require().NoError(err)

	_, err = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg7)
	s.Require().NoError(err)

	_, err = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg8)
	s.Require().NoError(err)

}

func (s *KeeperTestSuite) TestMsgCalculateInterestAndRewards() {
	s.ctx = s.ctx.WithBlockTime(utils.ParseTime("2022-03-01T12:00:00Z"))
	s.TestMsgInitializeLend()
	msg := types.NewMsgBorrowAlternate("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, 1, sdk.NewCoin("uasset1", sdk.NewInt(100000000)), 3, false, sdk.NewCoin("uasset2", sdk.NewInt(10000000)), 1)
	_, err := s.msgServer.BorrowAlternate(sdk.WrapSDKContext(s.ctx), msg)
	s.Require().NoError(err)
	s.ctx = s.ctx.WithBlockTime(utils.ParseTime("2023-03-01T12:00:00Z"))

	testCases := []struct {
		Name               string
		Msg                types.MsgCalculateInterestAndRewards
		ExpErr             error
		ExpResp            *types.MsgCalculateInterestAndRewardsResponse
		QueryResponseIndex uint64
		QueryResponse      *types.MsgCalculateInterestAndRewards
		AvailableBalance   sdk.Coins
	}{
		{
			Name:               "Valid Case",
			Msg:                *types.NewMsgCalculateInterestAndRewards("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"),
			ExpErr:             nil,
			ExpResp:            &types.MsgCalculateInterestAndRewardsResponse{},
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(100))),
		},
	}
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.Name, func() {
			// add funds to acount for valid case
			ctx := sdk.WrapSDKContext(s.ctx)
			resp, err := s.msgServer.CalculateInterestAndRewards(ctx, &tc.Msg)
			if tc.ExpErr != nil {
				s.Require().Error(err)
				s.Require().EqualError(err, tc.ExpErr.Error())
				s.Require().Equal(tc.ExpResp, resp)

			} else {
				s.Require().NoError(err)
				s.Require().NotNil(resp)
				s.Require().Equal(tc.ExpResp, resp)
				borrowPos, _ := s.keeper.GetBorrow(s.ctx, 1)
				if borrowPos.InterestAccumulated == sdk.ZeroDec() {
					s.Error(errors.Wrap(err, "Interest Not Calculated"))
				}
			}
		})
	}
}

func (s *KeeperTestSuite) TestMsgBorrowForIsolatedMode() {
	assetOneID := s.CreateNewAsset("ASSETONE", "uasset1", 1000000)
	assetTwoID := s.CreateNewAsset("ASSETTWO", "uasset2", 2000000)
	assetThreeID := s.CreateNewAsset("ASSETTHREE", "uasset3", 2000000)
	assetFourID := s.CreateNewAsset("ASSETFOUR", "uasset4", 2000000)
	cAssetOneID := s.CreateNewAsset("CASSETONE", "ucasset1", 1000000)
	cAssetTwoID := s.CreateNewAsset("CASSETTWO", "ucasset2", 2000000)
	cAssetThreeID := s.CreateNewAsset("CASSETTHRE", "ucasset3", 2000000)
	cAssetFourID := s.CreateNewAsset("CASSETFOUR", "ucasset4", 2000000)

	var (
		assetDataPoolOne []*types.AssetDataPoolMapping
		assetDataPoolTwo []*types.AssetDataPoolMapping
	)
	assetDataPoolOneAssetOne := &types.AssetDataPoolMapping{
		AssetID:          assetOneID,
		AssetTransitType: 3,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolOneAssetTwo := &types.AssetDataPoolMapping{
		AssetID:          assetTwoID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(1000000000000000000),
	}
	assetDataPoolOneAssetThree := &types.AssetDataPoolMapping{
		AssetID:          assetThreeID,
		AssetTransitType: 2,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolTwoAssetFour := &types.AssetDataPoolMapping{
		AssetID:          assetFourID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(3000000000000000000),
	}

	assetDataPoolOne = append(assetDataPoolOne, assetDataPoolOneAssetOne, assetDataPoolOneAssetTwo, assetDataPoolOneAssetThree)
	assetDataPoolTwo = append(assetDataPoolTwo, assetDataPoolTwoAssetFour, assetDataPoolOneAssetOne, assetDataPoolOneAssetThree)

	s.AddAssetRatesStats(assetThreeID, newDec("0.8"), newDec("0.002"), newDec("0.06"), newDec("0.6"), true, newDec("0.04"), newDec("0.04"), newDec("0.06"), newDec("0.8"), newDec("0.85"), newDec("0.025"), newDec("0.025"), newDec("0.1"), cAssetThreeID)
	s.AddAssetRatesStats(assetOneID, newDec("0.75"), newDec("0.002"), newDec("0.07"), newDec("1.25"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.7"), newDec("0.75"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetOneID)
	s.AddAssetRatesPoolPairs(assetTwoID, newDec("0.5"), newDec("0.002"), newDec("0.08"), newDec("2.0"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.5"), newDec("0.55"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetTwoID, "cmdx", "CMDX-ATOM-CMST", assetDataPoolOne, 1000000, true)
	s.AddAssetRatesPoolPairs(assetFourID, newDec("0.65"), newDec("0.002"), newDec("0.08"), newDec("1.5"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.6"), newDec("0.65"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetFourID, "osmo", "OSMO-ATOM-CMST", assetDataPoolTwo, 1000000, false)

	appOneID := s.CreateNewApp("commodo", "cmmdo")
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset1", newInt(1000000000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset2", newInt(1000000000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset3", newInt(1000000000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset4", newInt(1000000000000000))))

	msg := types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetOneID, sdk.NewCoin("uasset1", newInt(300)), 1, appOneID)
	msgLend2 := types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetTwoID, sdk.NewCoin("uasset2", newInt(10000000000)), 1, appOneID)

	msg3 := types.NewMsgFundModuleAccounts(1, assetOneID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset1", newInt(10000000000)))
	msg4 := types.NewMsgFundModuleAccounts(1, assetTwoID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset2", newInt(10000000000)))
	msg5 := types.NewMsgFundModuleAccounts(1, assetThreeID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset3", newInt(120000000)))
	msg7 := types.NewMsgFundModuleAccounts(2, assetOneID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset1", newInt(10000000000)))
	msg8 := types.NewMsgFundModuleAccounts(2, assetFourID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset4", newInt(10000000000)))

	_, _ = s.msgServer.Lend(sdk.WrapSDKContext(s.ctx), msg)
	_, _ = s.msgServer.Lend(sdk.WrapSDKContext(s.ctx), msgLend2)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg3)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg4)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg5)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg7)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg8)

	msg2 := types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, 1, false, sdk.NewCoin("ucasset1", newInt(100)), sdk.NewCoin("uasset2", newInt(10)))
	_, _ = s.msgServer.Borrow(sdk.WrapSDKContext(s.ctx), msg2)

	testCases := []struct {
		Name               string
		Msg                types.MsgBorrow
		ExpErr             error
		ExpResp            *types.MsgBorrowResponse
		QueryResponseIndex uint64
		QueryResponse      *types.MsgBorrow
		AvailableBalance   sdk.Coins
	}{
		{
			Name:               "success valid case",
			Msg:                *types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, 3, false, sdk.NewCoin("ucasset1", newInt(100)), sdk.NewCoin("uasset3", newInt(10))),
			ExpErr:             nil,
			ExpResp:            &types.MsgBorrowResponse{},
			QueryResponseIndex: 0,
			QueryResponse: &types.MsgBorrow{
				Borrower:       "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
				LendId:         1,
				PairId:         4,
				IsStableBorrow: false,
				AmountIn:       sdk.NewCoin("ucasset1", newInt(100)),
				AmountOut:      sdk.NewCoin("uasset3", newInt(10)),
			},
			AvailableBalance: sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "success valid case 2",
			Msg:                *types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 2, 2, false, sdk.NewCoin("ucasset2", newInt(100)), sdk.NewCoin("uasset1", newInt(1))),
			ExpErr:             nil,
			ExpResp:            &types.MsgBorrowResponse{},
			QueryResponseIndex: 0,
			QueryResponse: &types.MsgBorrow{
				Borrower:       "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
				LendId:         1,
				PairId:         15,
				IsStableBorrow: false,
				AmountIn:       sdk.NewCoin("ucasset1", newInt(100)),
				AmountOut:      sdk.NewCoin("uasset4", newInt(10)),
			},
			AvailableBalance: sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "Isolated Mode",
			Msg:                *types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 2, 5, false, sdk.NewCoin("ucasset2", newInt(100)), sdk.NewCoin("uasset3", newInt(1))),
			ExpErr:             types.ErrorIsolatedModeActivated,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
	}
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.Name, func() {
			// add funds to acount for valid case
			if tc.ExpErr == nil {
				s.fundAddr(sdk.MustAccAddressFromBech32(tc.Msg.Borrower), sdk.NewCoins(sdk.NewCoin("uasset1", sdk.NewIntFromUint64(300))))
				s.fundAddr(sdk.MustAccAddressFromBech32(tc.Msg.Borrower), sdk.NewCoins(sdk.NewCoin("ucasset2", sdk.NewIntFromUint64(1000000000))))
			}

			ctx := sdk.WrapSDKContext(s.ctx)
			resp, err := s.msgServer.Borrow(ctx, &tc.Msg)
			if tc.ExpErr != nil {
				s.Require().Error(err)
				s.Require().EqualError(err, tc.ExpErr.Error())
				s.Require().Equal(tc.ExpResp, resp)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(resp)
				s.Require().Equal(tc.ExpResp, resp)
			}
		})
	}
}

func (s *KeeperTestSuite) TestMsgBorrowForEMode() {
	assetOneID := s.CreateNewAsset("ASSETONE", "uasset1", 1000000)
	assetTwoID := s.CreateNewAsset("ASSETTWO", "uasset2", 2000000)
	assetThreeID := s.CreateNewAsset("ASSETTHREE", "uasset3", 2000000)
	assetFourID := s.CreateNewAsset("ASSETFOUR", "uasset4", 2000000)
	cAssetOneID := s.CreateNewAsset("CASSETONE", "ucasset1", 1000000)
	cAssetTwoID := s.CreateNewAsset("CASSETTWO", "ucasset2", 2000000)
	cAssetThreeID := s.CreateNewAsset("CASSETTHRE", "ucasset3", 2000000)
	cAssetFourID := s.CreateNewAsset("CASSETFOUR", "ucasset4", 2000000)

	var (
		assetDataPoolOne []*types.AssetDataPoolMapping
		assetDataPoolTwo []*types.AssetDataPoolMapping
	)
	assetDataPoolOneAssetOne := &types.AssetDataPoolMapping{
		AssetID:          assetOneID,
		AssetTransitType: 3,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolOneAssetTwo := &types.AssetDataPoolMapping{
		AssetID:          assetTwoID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(1000000000000000000),
	}
	assetDataPoolOneAssetThree := &types.AssetDataPoolMapping{
		AssetID:          assetThreeID,
		AssetTransitType: 2,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolTwoAssetFour := &types.AssetDataPoolMapping{
		AssetID:          assetFourID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(3000000000000000000),
	}

	assetDataPoolOne = append(assetDataPoolOne, assetDataPoolOneAssetOne, assetDataPoolOneAssetTwo, assetDataPoolOneAssetThree)
	assetDataPoolTwo = append(assetDataPoolTwo, assetDataPoolTwoAssetFour, assetDataPoolOneAssetOne, assetDataPoolOneAssetThree)

	s.AddAssetRatesStats(assetThreeID, newDec("0.8"), newDec("0.002"), newDec("0.06"), newDec("0.6"), true, newDec("0.04"), newDec("0.04"), newDec("0.06"), newDec("0.8"), newDec("0.85"), newDec("0.025"), newDec("0.025"), newDec("0.1"), cAssetThreeID)
	s.AddAssetRatesStats(assetOneID, newDec("0.75"), newDec("0.002"), newDec("0.07"), newDec("1.25"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.7"), newDec("0.75"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetOneID)
	s.AddAssetRatesPoolPairs(assetTwoID, newDec("0.5"), newDec("0.002"), newDec("0.08"), newDec("2.0"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.5"), newDec("0.55"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetTwoID, "cmdx", "CMDX-ATOM-CMST", assetDataPoolOne, 1000000, false)
	s.AddAssetRatesPoolPairs(assetFourID, newDec("0.65"), newDec("0.002"), newDec("0.08"), newDec("1.5"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.6"), newDec("0.65"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetFourID, "osmo", "OSMO-ATOM-CMST", assetDataPoolTwo, 1000000, false)

	var EModePairs []types.EModePairs
	EModePairs = append(EModePairs, types.EModePairs{
		PairID:                3,
		ELtv:                  newDec("0.9"),
		ELiquidationThreshold: newDec("0.85"),
		ELiquidationPenalty:   newDec("0.05"),
	})

	s.AddEModePairs(EModePairs)

	appOneID := s.CreateNewApp("commodo", "cmmdo")
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset1", newInt(1000000000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset2", newInt(1000000000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset3", newInt(1000000000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset4", newInt(1000000000000000))))

	msg := types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetOneID, sdk.NewCoin("uasset1", newInt(300)), 1, appOneID)
	msgLend2 := types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetTwoID, sdk.NewCoin("uasset2", newInt(10000000000)), 1, appOneID)

	msg3 := types.NewMsgFundModuleAccounts(1, assetOneID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset1", newInt(10000000000)))
	msg4 := types.NewMsgFundModuleAccounts(1, assetTwoID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset2", newInt(10000000000)))
	msg5 := types.NewMsgFundModuleAccounts(1, assetThreeID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset3", newInt(120000000)))
	msg7 := types.NewMsgFundModuleAccounts(2, assetOneID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset1", newInt(10000000000)))
	msg8 := types.NewMsgFundModuleAccounts(2, assetFourID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset4", newInt(10000000000)))

	_, _ = s.msgServer.Lend(sdk.WrapSDKContext(s.ctx), msg)
	_, _ = s.msgServer.Lend(sdk.WrapSDKContext(s.ctx), msgLend2)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg3)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg4)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg5)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg7)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg8)

	msg2 := types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, 1, false, sdk.NewCoin("ucasset1", newInt(100)), sdk.NewCoin("uasset2", newInt(10)))
	_, _ = s.msgServer.Borrow(sdk.WrapSDKContext(s.ctx), msg2)

	testCases := []struct {
		Name               string
		Msg                types.MsgBorrow
		ExpErr             error
		ExpResp            *types.MsgBorrowResponse
		QueryResponseIndex uint64
		QueryResponse      *types.MsgBorrow
		AvailableBalance   sdk.Coins
	}{
		{
			Name:               "success valid case",
			Msg:                *types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, 3, false, sdk.NewCoin("ucasset1", newInt(100)), sdk.NewCoin("uasset3", newInt(90))),
			ExpErr:             nil,
			ExpResp:            &types.MsgBorrowResponse{},
			QueryResponseIndex: 0,
			QueryResponse: &types.MsgBorrow{
				Borrower:       "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
				LendId:         1,
				PairId:         4,
				IsStableBorrow: false,
				AmountIn:       sdk.NewCoin("ucasset1", newInt(100)),
				AmountOut:      sdk.NewCoin("uasset3", newInt(10)),
			},
			AvailableBalance: sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
		{
			Name:               "error invalid case 2",
			Msg:                *types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, 1, false, sdk.NewCoin("ucasset1", newInt(100)), sdk.NewCoin("uasset2", newInt(140))),
			ExpErr:             types.ErrorInvalidCollateralizationRatio,
			ExpResp:            nil,
			QueryResponseIndex: 0,
			QueryResponse:      nil,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
	}
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.Name, func() {
			// add funds to acount for valid case
			if tc.ExpErr == nil {
				s.fundAddr(sdk.MustAccAddressFromBech32(tc.Msg.Borrower), sdk.NewCoins(sdk.NewCoin("uasset1", sdk.NewIntFromUint64(300))))
				s.fundAddr(sdk.MustAccAddressFromBech32(tc.Msg.Borrower), sdk.NewCoins(sdk.NewCoin("ucasset2", sdk.NewIntFromUint64(1000000000))))
			}

			ctx := sdk.WrapSDKContext(s.ctx)
			resp, err := s.msgServer.Borrow(ctx, &tc.Msg)
			if tc.ExpErr != nil {
				s.Require().Error(err)
				s.Require().EqualError(err, tc.ExpErr.Error())
				s.Require().Equal(tc.ExpResp, resp)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(resp)
				s.Require().Equal(tc.ExpResp, resp)
			}
		})
	}
}

func (s *KeeperTestSuite) TestMsgRepayWithdraw() {
	assetOneID := s.CreateNewAsset("ASSETONE", "uasset1", 1000000)
	assetTwoID := s.CreateNewAsset("ASSETTWO", "uasset2", 2000000)
	assetThreeID := s.CreateNewAsset("ASSETTHREE", "uasset3", 2000000)
	assetFourID := s.CreateNewAsset("ASSETFOUR", "uasset4", 2000000)
	cAssetOneID := s.CreateNewAsset("CASSETONE", "ucasset1", 1000000)
	cAssetTwoID := s.CreateNewAsset("CASSETTWO", "ucasset2", 2000000)
	cAssetThreeID := s.CreateNewAsset("CASSETTHRE", "ucasset3", 2000000)
	cAssetFourID := s.CreateNewAsset("CASSETFOUR", "ucasset4", 2000000)

	var (
		assetDataPoolOne []*types.AssetDataPoolMapping
		assetDataPoolTwo []*types.AssetDataPoolMapping
	)
	assetDataPoolOneAssetOne := &types.AssetDataPoolMapping{
		AssetID:          assetOneID,
		AssetTransitType: 3,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolOneAssetTwo := &types.AssetDataPoolMapping{
		AssetID:          assetTwoID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(1000000000000000000),
	}
	assetDataPoolOneAssetThree := &types.AssetDataPoolMapping{
		AssetID:          assetThreeID,
		AssetTransitType: 2,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolTwoAssetFour := &types.AssetDataPoolMapping{
		AssetID:          assetFourID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(3000000000000000000),
	}

	assetDataPoolOne = append(assetDataPoolOne, assetDataPoolOneAssetOne, assetDataPoolOneAssetTwo, assetDataPoolOneAssetThree)
	assetDataPoolTwo = append(assetDataPoolTwo, assetDataPoolTwoAssetFour, assetDataPoolOneAssetOne, assetDataPoolOneAssetThree)

	s.AddAssetRatesStats(assetThreeID, newDec("0.8"), newDec("0.002"), newDec("0.06"), newDec("0.6"), true, newDec("0.04"), newDec("0.04"), newDec("0.06"), newDec("0.8"), newDec("0.85"), newDec("0.025"), newDec("0.025"), newDec("0.1"), cAssetThreeID)
	s.AddAssetRatesStats(assetOneID, newDec("0.75"), newDec("0.002"), newDec("0.07"), newDec("1.25"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.7"), newDec("0.75"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetOneID)
	s.AddAssetRatesPoolPairs(assetTwoID, newDec("0.5"), newDec("0.002"), newDec("0.08"), newDec("2.0"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.5"), newDec("0.55"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetTwoID, "cmdx", "CMDX-ATOM-CMST", assetDataPoolOne, 1000000, false)
	s.AddAssetRatesPoolPairs(assetFourID, newDec("0.65"), newDec("0.002"), newDec("0.08"), newDec("1.5"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.6"), newDec("0.65"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetFourID, "osmo", "OSMO-ATOM-CMST", assetDataPoolTwo, 1000000, false)

	var EModePairs []types.EModePairs
	EModePairs = append(EModePairs, types.EModePairs{
		PairID:                3,
		ELtv:                  newDec("0.9"),
		ELiquidationThreshold: newDec("0.85"),
		ELiquidationPenalty:   newDec("0.05"),
	})

	s.AddEModePairs(EModePairs)

	appOneID := s.CreateNewApp("commodo", "cmmdo")
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset1", newInt(1000000000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset2", newInt(1000000000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset3", newInt(1000000000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset4", newInt(1000000000000000))))

	msg := types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetOneID, sdk.NewCoin("uasset1", newInt(300)), 1, appOneID)
	msgLend2 := types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetTwoID, sdk.NewCoin("uasset2", newInt(10000000000)), 1, appOneID)

	msg3 := types.NewMsgFundModuleAccounts(1, assetOneID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset1", newInt(10000000000)))
	msg4 := types.NewMsgFundModuleAccounts(1, assetTwoID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset2", newInt(10000000000)))
	msg5 := types.NewMsgFundModuleAccounts(1, assetThreeID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset3", newInt(120000000)))
	msg7 := types.NewMsgFundModuleAccounts(2, assetOneID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset1", newInt(10000000000)))
	msg8 := types.NewMsgFundModuleAccounts(2, assetFourID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset4", newInt(10000000000)))

	_, _ = s.msgServer.Lend(sdk.WrapSDKContext(s.ctx), msg)
	_, _ = s.msgServer.Lend(sdk.WrapSDKContext(s.ctx), msgLend2)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg3)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg4)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg5)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg7)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg8)

	msg2 := types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, 1, false, sdk.NewCoin("ucasset1", newInt(100)), sdk.NewCoin("uasset2", newInt(10)))
	_, _ = s.msgServer.Borrow(sdk.WrapSDKContext(s.ctx), msg2)

	testCases := []struct {
		Name               string
		Msg                types.MsgRepayWithdraw
		ExpErr             error
		ExpResp            *types.MsgRepayWithdrawResponse
		QueryResponseIndex uint64
		AvailableBalance   sdk.Coins
	}{
		{
			Name:               "success valid case",
			Msg:                *types.NewMsgRepayWithdraw("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1),
			ExpErr:             nil,
			ExpResp:            &types.MsgRepayWithdrawResponse{},
			QueryResponseIndex: 0,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
	}
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.Name, func() {
			// add funds to acount for valid case
			if tc.ExpErr == nil {
				s.fundAddr(sdk.MustAccAddressFromBech32(tc.Msg.Borrower), sdk.NewCoins(sdk.NewCoin("uasset1", sdk.NewIntFromUint64(300))))
				s.fundAddr(sdk.MustAccAddressFromBech32(tc.Msg.Borrower), sdk.NewCoins(sdk.NewCoin("ucasset2", sdk.NewIntFromUint64(1000000000))))
			}

			ctx := sdk.WrapSDKContext(s.ctx)
			resp, err := s.msgServer.RepayWithdraw(ctx, &tc.Msg)
			if tc.ExpErr != nil {
				s.Require().Error(err)
				s.Require().EqualError(err, tc.ExpErr.Error())
				s.Require().Equal(tc.ExpResp, resp)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(resp)
				s.Require().Equal(tc.ExpResp, resp)
				lend, _ := s.keeper.GetLend(s.ctx, 1)
				s.Require().Equal(lend.AmountIn.Amount, sdk.NewInt(200))
				_, found := s.keeper.GetBorrow(s.ctx, 1)
				s.Require().Equal(found, false)
			}
		})
	}
}

func (s *KeeperTestSuite) TestMsgPoolDepreciate() {
	assetOneID := s.CreateNewAsset("ASSETONE", "uasset1", 1000000)
	assetTwoID := s.CreateNewAsset("ASSETTWO", "uasset2", 2000000)
	assetThreeID := s.CreateNewAsset("ASSETTHREE", "uasset3", 2000000)
	assetFourID := s.CreateNewAsset("ASSETFOUR", "uasset4", 2000000)
	cAssetOneID := s.CreateNewAsset("CASSETONE", "ucasset1", 1000000)
	cAssetTwoID := s.CreateNewAsset("CASSETTWO", "ucasset2", 2000000)
	cAssetThreeID := s.CreateNewAsset("CASSETTHRE", "ucasset3", 2000000)
	cAssetFourID := s.CreateNewAsset("CASSETFOUR", "ucasset4", 2000000)

	var (
		assetDataPoolOne []*types.AssetDataPoolMapping
		assetDataPoolTwo []*types.AssetDataPoolMapping
	)
	assetDataPoolOneAssetOne := &types.AssetDataPoolMapping{
		AssetID:          assetOneID,
		AssetTransitType: 3,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolOneAssetTwo := &types.AssetDataPoolMapping{
		AssetID:          assetTwoID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(1000000000000000000),
	}
	assetDataPoolOneAssetThree := &types.AssetDataPoolMapping{
		AssetID:          assetThreeID,
		AssetTransitType: 2,
		SupplyCap:        sdk.NewDec(5000000000000000000),
	}
	assetDataPoolTwoAssetFour := &types.AssetDataPoolMapping{
		AssetID:          assetFourID,
		AssetTransitType: 1,
		SupplyCap:        sdk.NewDec(3000000000000000000),
	}

	assetDataPoolOne = append(assetDataPoolOne, assetDataPoolOneAssetOne, assetDataPoolOneAssetTwo, assetDataPoolOneAssetThree)
	assetDataPoolTwo = append(assetDataPoolTwo, assetDataPoolTwoAssetFour, assetDataPoolOneAssetOne, assetDataPoolOneAssetThree)

	s.AddAssetRatesStats(assetThreeID, newDec("0.8"), newDec("0.002"), newDec("0.06"), newDec("0.6"), true, newDec("0.04"), newDec("0.04"), newDec("0.06"), newDec("0.8"), newDec("0.85"), newDec("0.025"), newDec("0.025"), newDec("0.1"), cAssetThreeID)
	s.AddAssetRatesStats(assetOneID, newDec("0.75"), newDec("0.002"), newDec("0.07"), newDec("1.25"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.7"), newDec("0.75"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetOneID)
	s.AddAssetRatesPoolPairs(assetTwoID, newDec("0.5"), newDec("0.002"), newDec("0.08"), newDec("2.0"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.5"), newDec("0.55"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetTwoID, "cmdx", "CMDX-ATOM-CMST", assetDataPoolOne, 1000000, false)
	appOneID := s.CreateNewApp("commodo", "cmmdo")
	msg := types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetOneID, sdk.NewCoin("uasset1", newInt(300)), 1, appOneID)
	msgLend2 := types.NewMsgLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", assetTwoID, sdk.NewCoin("uasset2", newInt(10000000000)), 1, appOneID)
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset1", newInt(1000000000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset2", newInt(1000000000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset3", newInt(1000000000000000))))
	s.fundAddr(sdk.MustAccAddressFromBech32("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"), sdk.NewCoins(sdk.NewCoin("uasset4", newInt(1000000000000000))))

	msg3 := types.NewMsgFundModuleAccounts(1, assetOneID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset1", newInt(10000000000)))
	msg4 := types.NewMsgFundModuleAccounts(1, assetTwoID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset2", newInt(10000000000)))
	msg5 := types.NewMsgFundModuleAccounts(1, assetThreeID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset3", newInt(120000000)))
	msg7 := types.NewMsgFundModuleAccounts(2, assetOneID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset1", newInt(10000000000)))
	msg8 := types.NewMsgFundModuleAccounts(2, assetFourID, "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", sdk.NewCoin("uasset4", newInt(10000000000)))

	_, _ = s.msgServer.Lend(sdk.WrapSDKContext(s.ctx), msg)
	_, _ = s.msgServer.Lend(sdk.WrapSDKContext(s.ctx), msgLend2)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg3)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg4)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg5)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg7)
	_, _ = s.msgServer.FundModuleAccounts(sdk.WrapSDKContext(s.ctx), msg8)
	//pairs := s.keeper.GetLendPairs(s.ctx)
	//fmt.Println("pairs", pairs)
	msg2 := types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, 1, false, sdk.NewCoin("ucasset1", newInt(100)), sdk.NewCoin("uasset2", newInt(10)))
	msg22 := types.NewMsgBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 2, 2, false, sdk.NewCoin("ucasset2", newInt(100)), sdk.NewCoin("uasset1", newInt(10)))
	_, _ = s.msgServer.Borrow(sdk.WrapSDKContext(s.ctx), msg2)
	_, _ = s.msgServer.Borrow(sdk.WrapSDKContext(s.ctx), msg22)

	var individualPool []types.IndividualPoolDepreciate
	individualPool = append(individualPool, types.IndividualPoolDepreciate{
		PoolID:            1,
		IsPoolDepreciated: false,
	})
	s.PoolDepreciate(individualPool)
	s.AddAssetRatesPoolPairs(assetFourID, newDec("0.65"), newDec("0.002"), newDec("0.08"), newDec("1.5"), false, newDec("0.0"), newDec("0.0"), newDec("0.0"), newDec("0.6"), newDec("0.65"), newDec("0.05"), newDec("0.05"), newDec("0.2"), cAssetFourID, "osmo", "OSMO-ATOM-CMST", assetDataPoolTwo, 1000000, false)
	msgDeposit := types.NewMsgDeposit("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1, sdk.NewCoin("uasset1", sdk.NewInt(100)))
	msgDraw := types.NewMsgDraw("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 2, sdk.NewCoin("uasset1", newInt(1)))
	testCases := []struct {
		Name               string
		Msg                types.MsgRepayWithdraw
		ExpErr             error
		ExpResp            *types.MsgRepayWithdrawResponse
		QueryResponseIndex uint64
		AvailableBalance   sdk.Coins
	}{
		{
			Name:               "success valid case",
			Msg:                *types.NewMsgRepayWithdraw("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1),
			ExpErr:             nil,
			ExpResp:            &types.MsgRepayWithdrawResponse{},
			QueryResponseIndex: 0,
			AvailableBalance:   sdk.NewCoins(sdk.NewCoin("uasset1", newInt(0))),
		},
	}
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.Name, func() {
			// add funds to acount for valid case
			if tc.ExpErr == nil {
				s.fundAddr(sdk.MustAccAddressFromBech32(tc.Msg.Borrower), sdk.NewCoins(sdk.NewCoin("uasset1", sdk.NewIntFromUint64(300))))
				s.fundAddr(sdk.MustAccAddressFromBech32(tc.Msg.Borrower), sdk.NewCoins(sdk.NewCoin("ucasset2", sdk.NewIntFromUint64(1000000000))))
			}

			ctx := sdk.WrapSDKContext(s.ctx)
			resp, err := s.msgServer.RepayWithdraw(ctx, &tc.Msg)
			if tc.ExpErr != nil {
				s.Require().Error(err)
				s.Require().EqualError(err, tc.ExpErr.Error())
				s.Require().Equal(tc.ExpResp, resp)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(resp)
				s.Require().Equal(tc.ExpResp, resp)

				_, err1 := s.msgServer.Lend(sdk.WrapSDKContext(s.ctx), msg)
				s.Require().NotNil(err1)
				_, err2 := s.msgServer.Borrow(sdk.WrapSDKContext(s.ctx), msg2)
				s.Require().NotNil(err2)
				_, err3 := s.msgServer.Deposit(sdk.WrapSDKContext(s.ctx), msgDeposit)
				s.Require().NotNil(err3)
				_, err4 := s.msgServer.Draw(sdk.WrapSDKContext(s.ctx), msgDraw)
				s.Require().NotNil(err4)
				msgCloseBorrow := types.NewMsgCloseBorrow("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 2)
				msgCloseLend1 := types.NewMsgCloseLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 1)
				msgCloseLend2 := types.NewMsgCloseLend("cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t", 2)
				_, _ = s.msgServer.CloseBorrow(sdk.WrapSDKContext(s.ctx), msgCloseBorrow)
				_, _ = s.msgServer.CloseLend(sdk.WrapSDKContext(s.ctx), msgCloseLend1)
				_, _ = s.msgServer.CloseLend(sdk.WrapSDKContext(s.ctx), msgCloseLend2)

			}
		})
	}

	pools := s.keeper.GetPools(s.ctx)
	s.Require().Equal(len(pools), 2)
	modBalFirstAsset := s.app.BankKeeper.GetBalance(s.ctx, authtypes.NewModuleAddress(pools[0].ModuleName), "uasset1")
	modBalSecondAsset := s.app.BankKeeper.GetBalance(s.ctx, authtypes.NewModuleAddress(pools[0].ModuleName), "uasset2")
	modBalThirdAsset := s.app.BankKeeper.GetBalance(s.ctx, authtypes.NewModuleAddress(pools[0].ModuleName), "uasset3")
	modBal1 := s.app.BankKeeper.GetBalance(s.ctx, authtypes.NewModuleAddress("lendV2"), "uasset1")
	modBal2 := s.app.BankKeeper.GetBalance(s.ctx, authtypes.NewModuleAddress("lendV2"), "uasset2")
	modBal3 := s.app.BankKeeper.GetBalance(s.ctx, authtypes.NewModuleAddress("lendV2"), "uasset3")

	fmt.Println("modBalFirstAsset", modBalFirstAsset)
	fmt.Println("modBalSecondAsset", modBalSecondAsset)
	fmt.Println("modBalThirdAsset", modBalThirdAsset)
	fmt.Println("modBal1", modBal1)
	fmt.Println("modBal2", modBal2)
	fmt.Println("modBal3", modBal3)
	s.nextBlock()
	modBalFirstAsset = s.app.BankKeeper.GetBalance(s.ctx, authtypes.NewModuleAddress(pools[0].ModuleName), "uasset1")
	modBalSecondAsset = s.app.BankKeeper.GetBalance(s.ctx, authtypes.NewModuleAddress(pools[0].ModuleName), "uasset2")
	modBalThirdAsset = s.app.BankKeeper.GetBalance(s.ctx, authtypes.NewModuleAddress(pools[0].ModuleName), "uasset3")
	modBal1 = s.app.BankKeeper.GetBalance(s.ctx, authtypes.NewModuleAddress("lendV2"), "uasset1")
	modBal2 = s.app.BankKeeper.GetBalance(s.ctx, authtypes.NewModuleAddress("lendV2"), "uasset2")
	modBal3 = s.app.BankKeeper.GetBalance(s.ctx, authtypes.NewModuleAddress("lendV2"), "uasset3")
	fmt.Println("modBalFirstAsset", modBalFirstAsset)
	fmt.Println("modBalSecondAsset", modBalSecondAsset)
	fmt.Println("modBalThirdAsset", modBalThirdAsset)
	fmt.Println("modBal1", modBal1)
	fmt.Println("modBal2", modBal2)
	fmt.Println("modBal3", modBal3)
	pools = s.keeper.GetPools(s.ctx)
	s.Require().Equal(len(pools), 1)

}
