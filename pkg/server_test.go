package pkg_test

import (
	"fmt"
	pkg2 "github.com/cloud-club/Aviator-service/pkg/mocks"
	"testing"
)

func (suite *NcpSuite) TestList() {
	mockInterface := &pkg2.MockServerInterface{}

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

			if tests[i].actualError == "" {
				mockInterface.On("List", tests[i].url).
					Return(nil, nil).
					Once()

			} else {
				mockInterface.On("List", tests[i].url).
					Return(nil, fmt.Errorf("%s", tests[i].actualError)).
					Once()
			}
			suite.ncp.Server.Interface = mockInterface

			_, err := suite.ncp.Server.Interface.List(tests[i].url)
			if err != nil {
				suite.Assert().EqualError(err, tests[i].expectedError)
			} else {
				suite.Assert().NoError(err)
			}
		})
	}

}
