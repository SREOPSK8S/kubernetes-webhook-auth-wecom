package wecom

type CorpIDAndSecret struct {
	CorpID     string
	CorpSecret string
}

type BaseResponse struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}
type AccessTokenResponse struct {
	BaseResponse
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
}
// ReadMemberResponse https://developer.work.weixin.qq.com/document/path/90196
type ReadMemberResponse struct {
	BaseResponse
	Userid string `json:"userid"`
	Name string `json:"name"`
	Department []int `json:"department"`
	Mobile string `json:"mobile"`
	Gender string `json:"gender"`
	Email string `json:"email"`
	Telephone string `json:"telephone"`
	Alias string `json:"alias"`
	Status int `json:"status"`
}
// GetDepartmentDetailsResponse https://developer.work.weixin.qq.com/document/path/95351
type GetDepartmentDetailsResponse struct {
	BaseResponse
	Department Department `json:"department"`
}
type Department struct {
	ID int `json:"id"`
	Name string `json:"name"`
	NameEn string `json:"name_en"`
	DepartmentLeader []string `json:"department_leader"`
	ParentID int `json:"parentid"`
	Order int `json:"order"`
}

func (ac *AccessTokenResponse) GetAccessToken() (accessToken string) {
	accessToken = ac.AccessToken
	return
}

