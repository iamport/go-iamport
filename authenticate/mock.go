package authenticate

import "net/http"

const (
	BaseURL       = "https://api.iamport.kr"
	RestApiKey    = "2737886909191347"
	RestApiSecret = "DNZ25OYnAk9qaRobLy9SWEBGyJKP1PsEDrHIfF6QZfha6FmSevKa9mRI6Cx7s0L5rsOH8Ux8aPihvE9J"

	BingBongRestApiKey    = "2421367378124191"
	BingBongRestApiSecret = "TuL0mVHDmZMFDz3mncCQiA2idqkHxuLrZwrms9ZO8MuMmx8JKLz3lBlz2Fgza10aysw0BwPKijwEQFoA"
)

func GetMockBaseAuthenticate() *Authenticate {
	return getMockAuthenticate(RestApiKey, RestApiSecret)
}

func GetMockBingBongAuthenticate() *Authenticate {
	return getMockAuthenticate(BingBongRestApiKey, BingBongRestApiSecret)
}

func getMockAuthenticate(key string, secret string) *Authenticate {
	cli := &http.Client{}

	auth, _ := NewAuthenticate(BaseURL, cli, key, secret)

	return auth
}
