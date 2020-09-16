package authenticate

import "net/http"

const (
	BaseURL       = "https://api.iamport.kr"
	RestApiKey    = "imp_apikey"
	RestApiSecret = "ekKoeW8RyKuT0zgaZsUtXXTLQ4AhPFW3ZGseDA6bkA5lamv9OqDMnxyeB9wqOsuO9W3Mx9YSJ4dTqJ3f"
)

func GetMockBaseAuthenticate() *Authenticate {
	cli := &http.Client{}

	auth, _ := NewAuthenticate(BaseURL, cli, RestApiKey, RestApiSecret)

	return auth
}
