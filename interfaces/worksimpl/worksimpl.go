package worksimpl

import (
	"context"
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/config"
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/entity/auth"
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/entity/interfs"
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/entity/wecom"
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/interfaces/logs"
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/interfaces/stores"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
	"strings"
	"sync"
	"time"
)

var _ interfs.AuthenticationUserInfo = &WorkChatImpl{}
var _ wecom.ServerAccessToken = &WorkChatImpl{}
var _ wecom.StoreAgentID = &AppAgentID{}

type AppAgentID struct {
	once    sync.Once
	Locker  sync.RWMutex
	AgentID int
}

var agentID int = config.GetAgentId()

func (a *AppAgentID) SetStoreAgentID(ctx context.Context) bool {
	a.AgentID = config.GetAgentId()
	client := stores.NewStore()
	defer a.Locker.Unlock()
	a.Locker.Lock()
	response, err := client.Put(ctx, wecom.WorkChatAppAgentIDKeyName, strconv.Itoa(agentID))
	if err != nil {
		logs.Logger.Error("SetStoreAgentID failure", zap.Any("error_msg", err))
		return false
	}
	logs.Logger.Info("SetStoreAgentID success", zap.Any("response", response))
	return true
}
func (a *AppAgentID) InitStoreAgentID() {
	a.once.Do(func() {
		a.SetStoreAgentID(context.TODO())
	})
}

func (a *AppAgentID) GetStoreAgentID(ctx context.Context) (int, bool) {
	result := ""
	client := stores.NewStore()
	if client == nil {
		return 0, false
	}
	agentIDName := wecom.WorkChatAppAgentIDKeyName
	response, err := client.Get(ctx, agentIDName)
	if err != nil {
		logs.Logger.Error("GetStoreAgentID failure", zap.Any("error_msg", err))
		return 0, false
	}
	defer func() {
		err := client.Close()
		if err != nil {
			logs.Logger.Error("store client close  failure", zap.Any("error_msg", err))
			return
		}
	}()
	logs.Logger.Debug("GetStoreAgentID success", zap.Any("response", response))
	for _, item := range response.Kvs {
		result = string(item.Value)
	}
	if result == "" {
		logs.Logger.Warn("GetStoreAgentID success result,but result is empty", zap.Any("result", result))
		return 0, false
	}
	ID, errs := strconv.Atoi(result)
	if errs != nil {
		return 0, false
	}
	return ID, true
}

type WorkChatImpl struct {
	AccessTokenMap  map[string]string
	SuccessResponse *wecom.ReadMemberResponse
}

func (w *WorkChatImpl) SendMsgToUser(ctx context.Context, msg string, msgType string, users ...string) bool {
	select {
	case <-ctx.Done():
		logs.Logger.Warn("Cancel SendMsgToUser to users",zap.Any("users",users))
		return false
	default:
		result := wecom.SendAppMessageTypeResponse{}
		toUsers := strings.Join(users, "|")
		client := resty.New()
		client.SetQueryParams(w.AccessTokenMap)
		typeRequest := wecom.GetMessageTypeRequest(msgType)
		// 使用类型断言
		// todo 获取GetAgentId 从内存中获取
		// 使用sync.once去获取 AgentID
		BParamsText := new(wecom.SendAppMessageRequestText)
		BParamsMarkDown := new(wecom.SendAppMessageMarkDownRequest)
		switch typeRequestI := typeRequest.(type) {
		case *wecom.SendAppMessageRequestText:
			BParamsText = typeRequestI
			BParamsText.SetSendAppMessageRequestTextParam(msg, toUsers, agentID)
		case *wecom.SendAppMessageMarkDownRequest:
			BParamsMarkDown = typeRequestI
			BParamsMarkDown.SetSendAppMessageRequestMarkDownParam(msg, toUsers, agentID)
		}
		// 定义接口接收通用数据
		var BParams interface{}
		switch msgType {
		case "text":
			BParams = BParamsText
		case "markdown":
			BParams = BParamsMarkDown
		}
		response, err := client.R().SetBody(BParams).SetResult(&result).Post(wecom.SendAppMessageURL)
		if err != nil || result.ErrorCode != 0 && result.ErrorMessage != "ok" {
			logs.Logger.Warn("SendMsgToUser failure", zap.Any("response", response))
			return false
		}
		logs.Logger.Info("SendMsgToUser success", zap.Any("response", response))
		return true
	}
}

func (w *WorkChatImpl) GetServerAccessToken() (accessToken string, status bool) {
	// 需要完成从cache里面获取
	ctx := context.TODO()
	var store wecom.StoreAccessToken = stores.EtcdImpl{}
	// todo 从缓存读取
	accessToken, status = store.GetSoreAccessToken(ctx)
	if !status || accessToken == "" {
		// 不在缓存中，请求后端服务并重新写入缓存
		result, ok := w.GetAccessTokenFromWorkChat()
		status = ok
		accessToken = result.AccessToken
		// todo 写入缓存中
		setStatus := store.SetSoreAccessToken(ctx, accessToken, wecom.WorkChatAccessTokenExpire)
		if setStatus != true {
			logs.Logger.Error("Store SetSoreAccessToken Token failure")
		}
		return
	}
	return accessToken, status
}

func (w *WorkChatImpl) GetAccessTokenFromWorkChat() (result *wecom.AccessTokenResponse, status bool) {
	params := w.getAccessTokenFromWorkChatPre()
	client := resty.New()
	client.SetQueryParams(params)
	response, err := client.R().SetResult(&result).Get(wecom.GetWorkChatAccessTokenURL)
	if err != nil {
		return nil, false
	}
	if result.ErrorCode != 0 && result.ErrorMessage != "ok" || response.RawResponse.StatusCode != 200 {
		logs.Logger.Info("Get Token failure,", zap.Any("response", response))
		return result, false
	}
	w.AccessTokenMap = map[string]string{}
	w.AccessTokenMap["access_token"] = result.AccessToken
	logs.Logger.Info("Get Token success", zap.Any("response", response))
	return result, true
}

func (WorkChatImpl) TokenReviewFailure(review auth.TokenReview) auth.TokenReviewResponse {
	return auth.TokenReviewResponse{
		TypeMeta: metav1.TypeMeta{
			Kind:       review.Kind,
			APIVersion: review.APIVersion,
		},
		Status: auth.TokenReviewStatus{
			Authenticated: false,
			User:          auth.UserInfo{},
			Audiences:     []string{},
			Error:         "",
		},
	}
}

func (w *WorkChatImpl) TokenReviewSuccess(review auth.TokenReview) (successResponse auth.TokenReviewResponse) {
	usersInfo := auth.UserInfo{
		Username: w.SuccessResponse.Userid,
		UID:      w.SuccessResponse.Name,
		Groups:   w.GetDepartmentDetails(),
		Extra:    map[string]auth.ExtraValue{},
	}
	reviewStatus := auth.TokenReviewStatus{
		Authenticated: true,
		User:          usersInfo,
		Audiences:     []string{},
		Error:         "",
	}
	successResponse = auth.TokenReviewResponse{
		TypeMeta: metav1.TypeMeta{
			Kind:       review.Kind,
			APIVersion: review.APIVersion,
		},
		Status: reviewStatus,
	}
	// todo 消息推送给用户，通知用户结果
	ctxT, _ := context.WithTimeout(context.TODO(), time.Second*1)
	//ctxT, _ := context.WithTimeout(context.TODO(), time.Microsecond*1)
	go w.SendMsgToUser(ctxT, "auth success", "markdown", w.SuccessResponse.Userid)
	return
}

func (w *WorkChatImpl) TokenReviewVerify(review interface{}) bool {
	data, ok := review.(auth.TokenReview)
	if !ok {
		return false
	}
	// todo 实现为用户名生成Token，并验证Token
	// 当前根据用户名进行判断
	status, readResponse := w.GetReadMember(data.Spec.Token)
	w.SuccessResponse = readResponse
	return status
}

// 获取成员
func (w *WorkChatImpl) GetReadMember(token string) (status bool, readMemberResponse *wecom.ReadMemberResponse) {
	client := resty.New()
	// todo 需要转换Token到用户名
	w.AccessTokenMap["userid"] = token
	client.SetQueryParams(w.AccessTokenMap)
	response, err := client.R().SetResult(&readMemberResponse).Get(wecom.GetReadMemberURL)
	if err != nil {
		return false, nil
	}
	if readMemberResponse.ErrorCode != 0 && readMemberResponse.ErrorMessage != "ok" || response.RawResponse.StatusCode != 200 {
		logs.Logger.Info("GetReadMember failure,", zap.Any("response", readMemberResponse))
		//"response":{"errcode":42001,"errmsg":"access_token expired, more info at https://open.work.weixin.qq.com/devtool/query?e=42001"}
		if readMemberResponse.ErrorCode == 42001 {
			deleteExpireToken := stores.EtcdImpl{}
			deleteExpireToken.DeleteAccessToken(context.TODO())
			_, sts := w.GetServerAccessToken()
			if sts {
				logs.Logger.Info("token already expire,get new access token success")
				return
			}
			return
		}
		return false, nil
	}
	if readMemberResponse.Status != 1 {
		logs.Logger.Info("GetReadMember user not work,", zap.Any("response", readMemberResponse))
		return false, nil
	}
	logs.Logger.Info("GetReadMember success details", zap.Any("response", readMemberResponse))
	if strings.ToLower(w.AccessTokenMap["userid"]) == strings.ToLower(readMemberResponse.Userid) {
		return true, readMemberResponse
	}
	return false, nil
}

func (w *WorkChatImpl) GetDepartmentDetails() (nameList []string) {
	var result wecom.GetDepartmentDetailsResponse
	accessTokenMap := w.AccessTokenMap
	client := resty.New()
	for _, item := range w.SuccessResponse.Department {
		accessTokenMap["id"] = strconv.Itoa(item)
		client.SetQueryParams(accessTokenMap)
		response, err := client.R().SetResult(&result).Get(wecom.GetDepartmentDetailsURL)
		if err != nil {
			return
		}
		if result.ErrorCode != 0 && result.ErrorMessage != "ok" || response.StatusCode() != 200 {
			return
		}
		nameList = append(nameList, result.Department.Name)
	}
	logs.Logger.Debug("GetDepartmentDetails success details", zap.String("uid", w.SuccessResponse.Userid), zap.Strings("department", nameList))
	return
}

func (w *WorkChatImpl) getAccessTokenFromWorkChatPre() (params map[string]string) {
	corpPre := wecom.CorpIDAndSecret{
		CorpID:     config.GetCorpID(),
		CorpSecret: config.GetCorpSecret(),
	}
	params = make(map[string]string, 4)
	params["corpid"] = corpPre.CorpID
	params["corpsecret"] = corpPre.CorpSecret
	return params
}

func NewReadMemberResponse() *wecom.ReadMemberResponse {
	return &wecom.ReadMemberResponse{
		Department: []int{},
	}
}
