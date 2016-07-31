package common

import (
	"testing"

	"github.com/chikin14niwa/codecs"
	"github.com/chikin14niwa/objx"
	"github.com/chikin14niwa/testify/assert"
)

func TestCredentials_PublicData(t *testing.T) {

	creds := &Credentials{objx.MSI("auth", "ABC123", CredentialsKeyID, 123)}

	publicData, _ := codecs.PublicData(creds, nil)

	if assert.NotNil(t, publicData) {
		assert.Equal(t, "ABC123", publicData.(objx.Map)["auth"])
		assert.Equal(t, "123", publicData.(objx.Map)[CredentialsKeyID], "CredentialsKeyID ("+CredentialsKeyID+") must be turned into a string.")
	}

}
