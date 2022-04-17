package wecom

func (ac *AccessTokenResponse) GetAccessToken() (accessToken string) {
	accessToken = ac.AccessToken
	return
}

func (base *BaseResponse) GetErrorCode() (errorCode int){
	errorCode = base.ErrorCode
	return
}

func (base *BaseResponse) GetErrorMessage() (msg string)  {
	msg = base.ErrorMessage
	return
}

func (IDS *CorpIDAndSecret)GetCorpIDAndSecret() *CorpIDAndSecret {
	return &CorpIDAndSecret{
		CorpID:     IDS.CorpID,
		CorpSecret: IDS.CorpSecret,
	}
}

