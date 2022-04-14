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
		Endpoints:            []string{"localhost:2379"},
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

func (EtcdImpl) DeleteAccessToken() error {
	return nil
}

func (EtcdImpl) GetSoreAccessToken() (string, bool) {
	result := ""
	ctx := context.Background()
	client := NewStore()
	if client == nil {
		return "", false
	}
	response, err := client.Get(ctx, "")
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
	logs.Logger.Info("GetSoreAccessToken success", zap.Any("response", response))
	for _, item := range response.Kvs {
		result = string(item.Value)
	}
	return result, true
}

func (EtcdImpl) SetSoreAccessToken(token string) bool {
	ctx := context.Background()
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
	response, err := client.Put(ctx, "", token)
	if err != nil {
		logs.Logger.Error("SetSoreAccessToken failure", zap.Any("error_msg", err))
		return false
	}
	logs.Logger.Info("SetSoreAccessToken success", zap.Any("response", response))
	return true
}
