package oauth2

import (
	"github.com/chikin14niwa/gomniauth/common"
	"github.com/chikin14niwa/gomniauth/test"
	"github.com/chikin14niwa/testify/assert"
	"github.com/chikin14niwa/testify/mock"
	"testing"
)

func TestOAuth2ProviderGetClient(t *testing.T) {

	testTripperFactory := new(test.TestTripperFactory)
	testTripper := new(test.TestTripper)
	testProvider := new(test.TestProvider)

	creds := new(common.Credentials)

	testTripperFactory.On("NewTripper", creds, mock.Anything).Return(testTripper, nil)

	client, clientErr := GetClient(testTripperFactory, creds, testProvider)

	if assert.NoError(t, clientErr) {
		if assert.NotNil(t, client) {
			assert.Equal(t, client.Transport, testTripper)
		}
	}

	mock.AssertExpectationsForObjects(t, testTripperFactory.Mock, testTripper.Mock, testProvider.Mock)

}
