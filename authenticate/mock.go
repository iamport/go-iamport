package authenticate

import "net/http"

const (
	BaseURL       = "https://api.iamport.kr"
	RestApiKey    = "2737886909191347"
	RestApiSecret = "DNZ25OYnAk9qaRobLy9SWEBGyJKP1PsEDrHIfF6QZfha6FmSevKa9mRI6Cx7s0L5rsOH8Ux8aPihvE9J"
)

func GetMockBaseAuthenticate() *Authenticate {
	cli := &http.Client{}

	auth, _ := NewAuthenticate(BaseURL, cli, RestApiKey, RestApiSecret)

	return auth
}
