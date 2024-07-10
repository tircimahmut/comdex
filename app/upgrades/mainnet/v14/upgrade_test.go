package v14_test

import (
	"github.com/comdex-official/comdex/app"
	// v14 "github.com/comdex-official/comdex/app/upgrades/mainnet/v14"
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
	// s.Setup()

	// preUpgradeChecks(s)

	// upgradeHeight := int64(5)
	// s.ConfirmUpgradeSucceeded(v14.UpgradeName, upgradeHeight)

	// postUpgradeChecks(s)
}

func preUpgradeChecks(s *UpgradeTestSuite) {
}

func postUpgradeChecks(s *UpgradeTestSuite) {

}
