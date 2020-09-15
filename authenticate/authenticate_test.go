package authenticate

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	BaseURL       = "https://api.iamport.kr"
	RestApiKey    = "imp_apikey"
	RestApiSecret = "ekKoeW8RyKuT0zgaZsUtXXTLQ4AhPFW3ZGseDA6bkA5lamv9OqDMnxyeB9wqOsuO9W3Mx9YSJ4dTqJ3f"
)

func TestNewAuthenticate(t *testing.T) {
	cli := &http.Client{}

	auth, err := NewAuthenticate(BaseURL, cli, "", RestApiSecret)
	assert.EqualError(t, err, ErrRestAPIKeyMissing)
	assert.Nil(t, auth)

	auth, err = NewAuthenticate(BaseURL, cli, RestApiKey, "")
	assert.EqualError(t, err, ErrRestAPISecretMissing)
	assert.Nil(t, auth)

	auth, err = NewAuthenticate(BaseURL, cli, RestApiKey, RestApiSecret)
	assert.NoError(t, err)
}

func TestGetToken(t *testing.T) {
	cli := &http.Client{}

	auth, err := NewAuthenticate(BaseURL, cli, RestApiKey, RestApiSecret)
	assert.NoError(t, err)

	token, err := auth.GetToken()
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestBeforeExpiredGetTokenMustbeSameToken(t *testing.T) {
	cli := &http.Client{}

	auth, err := NewAuthenticate(BaseURL, cli, RestApiKey, RestApiSecret)
	assert.NoError(t, err)

	token, err := auth.GetToken()
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	token2, err := auth.GetToken()
	assert.NoError(t, err)
	assert.NotEmpty(t, token2)

	assert.Equal(t, token, token2)
}
