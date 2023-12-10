package pkg_test

import (
	"github.com/cloud-club/Aviator-service/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// This is our suite
type NcpSuite struct {
	suite.Suite
	ncp *pkg.NcpService
}

// This gets run automatically by `go test` so we call `suite.Run` inside it
func TestNcpSuite(t *testing.T) {
	suite.Run(t, new(NcpSuite))
}

// This method gets run only first time before starting test
func (suite *NcpSuite) SetupSuite() {
	suite.ncp = pkg.NewNcpService("ncp service token")

	err := suite.ncp.CreateToken("admin", "CloudClubAdmin", []string{"admin"})
	assert.NoError(suite.T(), err, nil)
}
