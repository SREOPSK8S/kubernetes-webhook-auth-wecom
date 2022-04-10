package wecom

type CorpIDAndSecret struct {
	CorpID     string
	CorpSecret string
}

type AccessTokenResponse struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
}

func (ac *AccessTokenResponse) GetAccessToken() (accessToken string) {
	accessToken = ac.AccessToken
	return
}

