package gomniauth

import (
	"github.com/chikin14niwa/gomniauth/test"
	"github.com/stretchr/objx"
	"github.com/chikin14niwa/testify/assert"
	"testing"
)

func TestProviderPublicData(t *testing.T) {

	provider := new(test.TestProvider)

	provider.On("Name").Return("TestName")
	provider.On("DisplayName").Return("TestDisplayName")

	publicData, _ := ProviderPublicData(provider, objx.MSI("loginpathFormat", "~auth/%s/login"))
	publicDataMap := publicData.(map[string]interface{})

	assert.Equal(t, publicDataMap["name"], "TestName")
	assert.Equal(t, publicDataMap["display"], "TestDisplayName")
	assert.Equal(t, publicDataMap["loginpath"], "~auth/TestName/login")

}
