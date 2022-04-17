package stores

import (
	"context"
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/entity/wecom"
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/interfaces/logs"
	store "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"time"
)

var _ wecom.StoreAccessToken = EtcdImpl{}

func NewStore() *store.Client {
	client, err := store.New(store.Config{
		Endpoints:            []string{"172.16.100.99:2379"},
		DialTimeout:          3 * time.Second,
		DialKeepAliveTime:    3 * time.Second,
		DialKeepAliveTimeout: 5 * time.Second,
		MaxCallSendMsgSize:   0,
		MaxCallRecvMsgSize:   0,
	})
	if err != nil {
		logs.Logger.Error("create store client is error", zap.Any("error_msg", err))
		return nil
	}
	return client
}

type EtcdImpl struct {
}
func (EtcdImpl) SetSoreAccessToken(ctx context.Context,token string,ttl int64) bool {
	var expireTime int64
	if ttl == 0 {
		expireTime = wecom.WorkChatAccessTokenExpire
	} else {
		expireTime = ttl
	}
	client := NewStore()
	if client == nil {
		return false
	}
	defer func() {
		err := client.Close()
		if err != nil {
			logs.Logger.Error("store client close  failure", zap.Any("error_msg", err))
			return
		}
	}()
	accessTokenName := wecom.WorkChatAccessTokenKeyName
	grantRes,errs := client.Grant(context.TODO(),expireTime)
	if errs != nil {
		logs.Logger.Error("SetSoreAccessToken Grant failure", zap.Any("error_msg", errs),zap.Any("response",grantRes))
		return false
	}
	response, err := client.Put(context.TODO(), accessTokenName, token,store.WithLease(grantRes.ID))
	if err != nil {
		logs.Logger.Error("SetSoreAccessToken failure", zap.Any("error_msg", err))
		return false
	}
	logs.Logger.Info("SetSoreAccessToken success", zap.Any("response", response))
	return true
}

func (EtcdImpl) GetSoreAccessToken(ctx context.Context) (string, bool) {
	result := ""
	client := NewStore()
	if client == nil {
		return "", false
	}
	accessTokenName := wecom.WorkChatAccessTokenKeyName
	response, err := client.Get(ctx, accessTokenName)
	if err != nil {
		logs.Logger.Error("GetSoreAccessToken failure", zap.Any("error_msg", err))
		return "", false
	}
	defer func() {
		err := client.Close()
		if err != nil {
			logs.Logger.Error("store client close  failure", zap.Any("error_msg", err))
			return
		}
	}()
	logs.Logger.Debug("GetSoreAccessToken success", zap.Any("response", response))
	for _, item := range response.Kvs {
		result = string(item.Value)
	}
	if result == "" {
		logs.Logger.Warn("GetSoreAccessToken success result,but result is empty", zap.Any("result", result))
		return result,false
	}
	return result, true
}

func (EtcdImpl) DeleteAccessToken(ctx context.Context) bool {
	client := NewStore()
	if client == nil {
		return false
	}
	accessTokenName := wecom.WorkChatAccessTokenKeyName
	response, err := client.Delete(ctx, accessTokenName)
	if err != nil {
		return false
	}
	logs.Logger.Debug("DeleteAccessToken response", zap.Any("response", response))
	return true
}

