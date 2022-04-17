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

func GetSendAppMessageRequest(msgType string) *SendAppMessageRequest {
	return &SendAppMessageRequest{
		Touser:  "",
		Toparty: "",
		Totag:   "",
		Msgtype: msgType,
		Agentid: 0,
		Text:    MessageContent{},
		Safe:    1,
	}
}
func (appR *SendAppMessageRequest) GetSendAppMessageRequestMsgtype()  string {
	return appR.Msgtype
}

func (appR *SendAppMessageRequest) GetSendAppMessageRequestTextContent() string {
	return appR.Text.Content
}