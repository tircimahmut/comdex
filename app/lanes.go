package app

import (
	"cosmossdk.io/math"

	signerextraction "github.com/skip-mev/block-sdk/adapters/signer_extraction_adapter"
	"github.com/skip-mev/block-sdk/block/base"
	defaultlane "github.com/skip-mev/block-sdk/lanes/base"
	freelane "github.com/skip-mev/block-sdk/lanes/free"
	mevlane "github.com/skip-mev/block-sdk/lanes/mev"
)

// CreateLanes walks through the process of creating the lanes for the block sdk. In this function
// we create three separate lanes - MEV, Free, and Default - and then return them.
//
// NOTE: Application Developers should closely replicate this function in their own application.
func CreateLanes(app *App) (*mevlane.MEVLane, *base.BaseLane, *base.BaseLane) {
	// 1. Create the signer extractor. This is used to extract the expected signers from
	// a transaction. Each lane can have a different signer extractor if needed.
	signerAdapter := signerextraction.NewDefaultAdapter()
	encodingConfig := NewEncodingConfig()

	// 2. Create the configurations for each lane. These configurations determine how many
	// transactions the lane can store, the maximum block space the lane can consume, and
	// the signer extractor used to extract the expected signers from a transaction.
	//
	// IMPORTANT NOTE: If the block sdk module is utilized to store lanes, than the maximum
	// block space will be replaced with what is in state / in the genesis file.

	// Create a mev configuration that accepts 1000 transactions and consumes 20% of the
	// block space.
	mevConfig := base.LaneConfig{
		Logger:          app.Logger(),
		TxEncoder:       encodingConfig.TxConfig.TxEncoder(),
		TxDecoder:       encodingConfig.TxConfig.TxDecoder(),
		MaxBlockSpace:   math.LegacyMustNewDecFromStr("0.2"),
		SignerExtractor: signerAdapter,
		MaxTxs:          0,
	}

	// Create a free configuration that accepts 1000 transactions and consumes 20% of the
	// block space.
	freeConfig := base.LaneConfig{
		Logger:          app.Logger(),
		TxEncoder:       encodingConfig.TxConfig.TxEncoder(),
		TxDecoder:       encodingConfig.TxConfig.TxDecoder(),
		MaxBlockSpace:   math.LegacyMustNewDecFromStr("0.2"),
		SignerExtractor: signerAdapter,
		MaxTxs:          0,
	}

	// Create a default configuration that accepts 1000 transactions and consumes 60% of the
	// block space.
	defaultConfig := base.LaneConfig{
		Logger:          app.Logger(),
		TxEncoder:       encodingConfig.TxConfig.TxEncoder(),
		TxDecoder:       encodingConfig.TxConfig.TxDecoder(),
		MaxBlockSpace:   math.LegacyMustNewDecFromStr("0.6"),
		SignerExtractor: signerAdapter,
		MaxTxs:          0,
	}

	// 3. Create the match handlers for each lane. These match handlers determine whether or not
	// a transaction belongs in the lane.

	// Create the final match handler for the mev lane.
	factory := mevlane.NewDefaultAuctionFactory(encodingConfig.TxConfig.TxDecoder(), signerAdapter)
	mevMatchHandler := factory.MatchHandler()

	// Create the final match handler for the free lane.
	freeMatchHandler := freelane.DefaultMatchHandler()

	// Create the final match handler for the default lane.
	defaultMatchHandler := base.DefaultMatchHandler()

	// 4. Create the lanes.
	mevLane := mevlane.NewMEVLane(
		mevConfig,
		factory,
		mevMatchHandler,
	)

	freeLane := freelane.NewFreeLane(
		freeConfig,
		base.DefaultTxPriority(),
		freeMatchHandler,
	)

	defaultLane := defaultlane.NewDefaultLane(
		defaultConfig,
		defaultMatchHandler,
	)

	return mevLane, freeLane, defaultLane
}
