package pkg_test

import (
	"fmt"
	"github.com/cloud-club/Aviator-service/pkg"
	"testing"
)

func (suite *NcpSuite) TestList() {
	mockInterface := &pkg.MockServerInterface{}

	// given
	tests := []struct {
		name          string
		url           string
		payload       interface{}
		expectedError string
		actualError   string
		expectedData  interface{}
	}{
		{
			name:          "성공",
			url:           "http://localhost:8080",
			payload:       nil,
			actualError:   "",
			expectedError: "",
			expectedData:  nil,
		},
		{
			name:          "(실패) url 입력 안함",
			url:           "",
			payload:       nil,
			actualError:   "please input url",
			expectedError: "please input url",
			expectedData:  nil,
		},
	}

	for i := range tests {
		suite.T().Helper()

		suite.T().Run(tests[i].name, func(t *testing.T) {
			suite.T().Logf("%s : running scenario %d", tests[i].name, i)

			mockInterface.On("List", tests[i].url).
				Return(nil, fmt.Errorf("%v", tests[i].actualError)).
				Once()

			suite.ncp.Server.Interface = mockInterface

			_, err := suite.ncp.Server.Interface.List(tests[i].url)
			if err.Error() != "" {
				suite.Assert().EqualError(err, tests[i].expectedError)
			}

		})
	}

}
