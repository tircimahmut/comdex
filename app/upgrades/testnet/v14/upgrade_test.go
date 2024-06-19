package v14_test

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/comdex-official/comdex/app"
	v14 "github.com/comdex-official/comdex/app/upgrades/testnet/v14"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UpgradeTestSuite struct {
	app.KeeperTestHelper
}

func (s *UpgradeTestSuite) SetupTest() {
	s.Setup()
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}

// Ensures the test does not error out.
func (s *UpgradeTestSuite) TestUpgrade() {
	s.Setup()

	preUpgradeChecks(s)

	upgradeHeight := int64(5)
	s.ConfirmUpgradeSucceeded(v14.UpgradeName, upgradeHeight)

	postUpgradeChecks(s)
}

func preUpgradeChecks(s *UpgradeTestSuite) {

	mp := s.App.MintKeeper.GetParams(s.Ctx)
	s.Require().Equal(mp.BlocksPerYear, uint64(6311520))

	sp := s.App.SlashingKeeper.GetParams(s.Ctx)
	s.Require().Equal(sp.SignedBlocksWindow, int64(100))

}

func postUpgradeChecks(s *UpgradeTestSuite) {

	// Ensure the mint params have halved
	mp := s.App.MintKeeper.GetParams(s.Ctx)
	s.Require().Equal(mp.BlocksPerYear, uint64(6311520/2))

	// Ensure the slashing params have halved
	sp := s.App.SlashingKeeper.GetParams(s.Ctx)
	s.Require().Equal(sp.SignedBlocksWindow, int64(100/2))

	// Ensure the wasm Permission nobody
	wp := s.App.WasmKeeper.GetParams(s.Ctx)
	s.Require().Equal(wp.CodeUploadAccess, wasmtypes.AllowNobody)

}
