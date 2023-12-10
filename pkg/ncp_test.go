package pkg_test

import (
	"github.com/cloud-club/Aviator-service/types/auth"
)

func (suite *NcpSuite) TestVerifyToken_success() {
	_, err := suite.ncp.VerifyToken(suite.ncp.GetToken(), &auth.AuthTokenClaims{})
	suite.Equal(err, nil, nil)
}

func (suite *NcpSuite) TestCreateAndVerifyToken_fail_token_is_invalid() {
	_, err := suite.ncp.VerifyToken(suite.ncp.GetToken()+"is failed", &auth.AuthTokenClaims{})
	suite.EqualError(err, "token signature is invalid")
}
