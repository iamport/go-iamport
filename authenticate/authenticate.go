package authenticate

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/iamport/go-iamport/util"
	"github.com/iamport/interface/build/go/authenticate"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	URLGetToken = "/users/getToken"

	IMPKey    = "imp_key"
	IMPSecret = "imp_secret"

	ErrRestAPIURLMissing    = "iamport: REST API URL is missing"
	ErrRestAPIKeyMissing    = "iamport: REST API Key is missing"
	ErrRestAPISecretMissing = "iamport: REST API Secret is missing"
)

type Authenticate struct {
	APIUrl              string
	Client              *http.Client
	RestAPIKeyAndSecret url.Values
	Token               string
	Expired             time.Time
}

// NewAuthenticate 는 api url, http.Client, rest api key, rest api seceret을 파라미터로 받아
// rest api token을 발급받아 authenticate 모듈을 return 해준다.
func NewAuthenticate(apiURL string, cli *http.Client, restAPIKey string, restAPISecret string) (*Authenticate, error) {
	if apiURL == "" {
		return nil, errors.New(ErrRestAPIURLMissing)
	}

	if restAPIKey == "" {
		return nil, errors.New(ErrRestAPIKeyMissing)
	}

	if restAPISecret == "" {
		return nil, errors.New(ErrRestAPISecretMissing)
	}

	form := url.Values{}
	form.Set(IMPKey, restAPIKey)
	form.Set(IMPSecret, restAPISecret)
	auth := &Authenticate{
		APIUrl:              apiURL,
		Client:              cli,
		RestAPIKeyAndSecret: form,
	}

	err := auth.RequestToken()
	if err != nil {
		return nil, err
	}

	return auth, nil
}

// GetToken rest api를 호출할 수 있는 token을 return해준다.
// token이 없거나 만료된 경우 RequestToken을 하여 새로운 토큰을 발급받아 return해준다
func (a *Authenticate) GetToken() (string, error) {
	now := time.Now()

	if a.Token == "" || a.Expired.IsZero() || !a.Expired.Before(now) {
		err := a.RequestToken()
		if err != nil {
			return "", nil
		}
	}

	return a.Token, nil
}

// RequestToken APIKey와 APISecret을 사용하여 AccessToken을 받아 온다.
// POST /users/getToken
func (a *Authenticate) RequestToken() error {
	urls := []string{a.APIUrl, URLGetToken}
	urlGetToken := strings.Join(urls, "")

	res, err := util.Call(a.Client, "", urlGetToken, util.POST, &a.RestAPIKeyAndSecret)
	if err != nil {
		return err
	}

	tokenRes := authenticate.TokenResponse{}
	err = protojson.Unmarshal(res, &tokenRes)
	if err != nil {
		return err
	}

	if tokenRes.Code != util.CodeOK {
		return err
	}

	a.Token = tokenRes.Response.AccessToken
	a.Expired = time.Unix(int64(tokenRes.Response.ExpiredAt), 0)

	return nil
}
