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

func GetSendAppMessageTextRequest() *SendAppMessageRequestText {
	return &SendAppMessageRequestText{
		SendAppMessageBase: SendAppMessageBase{
			Touser:  "",
			Toparty: "",
			Totag:   "",
			Msgtype: TextMsgType,
			Agentid: 0,
			Safe:    1,
		},
		Text:               MessageContent{},
	}
}
func (appR *SendAppMessageRequestText) GetSendAppMessageRequestMsgType()  string {
	return appR.Msgtype
}

func (appR *SendAppMessageRequestText) GetSendAppMessageRequestTextContent() string {
	return appR.Text.Content
}