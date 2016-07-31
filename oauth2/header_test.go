package oauth2

import (
	"github.com/chikin14niwa/gomniauth/common"
	"github.com/chikin14niwa/objx"
	"github.com/chikin14niwa/testify/assert"
	"testing"
)

func TestAuthorizationHeader(t *testing.T) {

	creds := &common.Credentials{Map: objx.MSI()}
	accessTokenVal := "ACCESSTOKEN"
	creds.Set(OAuth2KeyAccessToken, accessTokenVal)
	k, v := AuthorizationHeader(creds)

	assert.Equal(t, k, "Authorization")
	assert.Equal(t, "Bearer "+accessTokenVal, v)

}
