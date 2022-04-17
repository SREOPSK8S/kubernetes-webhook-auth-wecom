package wecom

import "context"

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
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// ReadMemberResponse https://developer.work.weixin.qq.com/document/path/90196
type ReadMemberResponse struct {
	BaseResponse
	Userid     string `json:"userid"`
	Name       string `json:"name"`
	Department []int  `json:"department"`
	Mobile     string `json:"mobile"`
	Gender     string `json:"gender"`
	Email      string `json:"email"`
	Telephone  string `json:"telephone"`
	Alias      string `json:"alias"`
	Status     int    `json:"status"`
}

// GetDepartmentDetailsResponse https://developer.work.weixin.qq.com/document/path/95351
type GetDepartmentDetailsResponse struct {
	BaseResponse
	Department Department `json:"department"`
}
type Department struct {
	ID               int      `json:"id"`
	Name             string   `json:"name"`
	NameEn           string   `json:"name_en"`
	DepartmentLeader []string `json:"department_leader"`
	ParentID         int      `json:"parentid"`
	Order            int      `json:"order"`
}

type ServerAccessToken interface {
	GetServerAccessToken() (accessTokenAccess string, status bool)
}

const (
	WorkChatAccessTokenKeyName string = "workChatKubernetesAccessKey"                        // 设置访问企业微信access_token键
	WorkChatAccessTokenExpire int64 = 7134
	GetWorkChatAccessTokenURL  string = "https://qyapi.weixin.qq.com/cgi-bin/gettoken"       // 获取access_token
	GetReadMemberURL           string = "https://qyapi.weixin.qq.com/cgi-bin/user/get"       // 读取成员
	GetDepartmentDetailsURL    string = "https://qyapi.weixin.qq.com/cgi-bin/department/get" // 获取单个部门详情
)

type StoreAccessToken interface {
	SetSoreAccessToken(context.Context,string,int64) bool
	GetSoreAccessToken(ctx context.Context) (string, bool)
	DeleteAccessToken(ctx context.Context) bool
}
