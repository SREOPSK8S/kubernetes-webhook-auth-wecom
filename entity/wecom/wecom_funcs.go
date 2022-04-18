package wecom

func (ac *AccessTokenResponse) GetAccessToken() (accessToken string) {
	accessToken = ac.AccessToken
	return
}

func (base *BaseResponse) GetErrorCode() (errorCode int) {
	errorCode = base.ErrorCode
	return
}

func (base *BaseResponse) GetErrorMessage() (msg string) {
	msg = base.ErrorMessage
	return
}

func (IDS *CorpIDAndSecret) GetCorpIDAndSecret() *CorpIDAndSecret {
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
		Text: MessageContent{},
	}
}

func GetSendAppMessageTextCardRequest() *SendAppMessageTextCardRequest {
	return &SendAppMessageTextCardRequest{
		SendAppMessageBase: SendAppMessageBase{
			Touser:  "",
			Toparty: "",
			Totag:   "",
			Msgtype: TextCardMsgType,
			Agentid: 0,
			Safe:    1,
		},
		Textcard: Textcard{
			Title:       "",
			Description: "",
			Url:         "",
			Btntxt:      "更多",
		},
	}
}

func GetSendAppMessageMarkDownRequest() *SendAppMessageMarkDownRequest {
	return &SendAppMessageMarkDownRequest{
		SendAppMessageBase: SendAppMessageBase{
			Touser:  "",
			Toparty: "",
			Totag:   "",
			Msgtype: MarkDownMsgType,
			Agentid: 0,
			Safe:    0,
		},
		Markdown: MessageContent{
			Content: "",
		},
	}
}
func (appR *SendAppMessageRequestText) GetSendAppMessageRequestMsgType() string {
	return appR.Msgtype
}

func (appR *SendAppMessageRequestText) GetSendAppMessageRequestTextContent() string {
	return appR.Text.Content
}

func GetMessageTypeRequest(msgType string) interface{} {
	switch msgType {
	case TextMsgType:
		return GetSendAppMessageTextRequest()
	//case TextCardMsgType:
	//	// todo 完成TextCardMsgType 请求结构体
	//	return GetSendAppMessageTextCardRequest()
	case MarkDownMsgType:
		return GetSendAppMessageMarkDownRequest()
	}
	return GetSendAppMessageTextRequest()
}
