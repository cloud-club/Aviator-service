package pkg_test

import (
	"github.com/cloud-club/Aviator-service/pkg"
	"github.com/cloud-club/Aviator-service/types/auth"
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
	suite.ncp.Server = pkg.NewServerService("ncp server token")

	err := suite.ncp.CreateToken("admin", "CloudClubAdmin", []string{"admin"})
	assert.NoError(suite.T(), err, nil)
}

func (suite *NcpSuite) TestVerifyToken_success() {
	_, err := suite.ncp.VerifyToken(suite.ncp.GetToken(), &auth.AuthTokenClaims{})
	suite.Equal(err, nil, nil)
}

func (suite *NcpSuite) TestCreateAndVerifyToken_fail_token_is_invalid() {
	_, err := suite.ncp.VerifyToken(suite.ncp.GetToken()+"is failed", &auth.AuthTokenClaims{})
	suite.EqualError(err, "token signature is invalid")
}
