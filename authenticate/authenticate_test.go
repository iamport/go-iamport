package authenticate

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAuthenticate(t *testing.T) {
	cli := &http.Client{}

	auth, err := NewAuthenticate("", cli, RestApiKey, RestApiSecret)
	assert.EqualError(t, err, ErrRestAPIURLMissing)
	assert.Nil(t, auth)

	auth, err = NewAuthenticate(BaseURL, cli, "", RestApiSecret)
	assert.EqualError(t, err, ErrRestAPIKeyMissing)
	assert.Nil(t, auth)

	auth, err = NewAuthenticate(BaseURL, cli, RestApiKey, "")
	assert.EqualError(t, err, ErrRestAPISecretMissing)
	assert.Nil(t, auth)

	auth, err = NewAuthenticate(BaseURL, cli, RestApiKey, RestApiSecret)
	assert.NoError(t, err)
}

func TestGetToken(t *testing.T) {
	auth := GetMockBaseAuthenticate()

	token, err := auth.GetToken()
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestBeforeExpiredGetTokenMustbeSameToken(t *testing.T) {
	auth := GetMockBaseAuthenticate()

	token, err := auth.GetToken()
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	token2, err := auth.GetToken()
	assert.NoError(t, err)
	assert.NotEmpty(t, token2)

	assert.Equal(t, token, token2)
}
