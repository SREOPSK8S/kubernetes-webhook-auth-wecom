package stores

import (
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/entity/wecom"
	"github.com/go-redis/redis/v8"
	"kubernetes-webhook-auth-wecom/vendor/github.com/go-redis/redis/v8"
)

var _ wecom.StoreAccessToken = RedisImpl{}

type RedisImpl struct {
}

func (RedisImpl) DeleteAccessToken() error {

	return nil
}

func (RedisImpl) GetSoreAccessToken() (string, bool) {

	return "", true
}

func (RedisImpl) SetSoreAccessToken(token string) bool {
	return true
}

func GetStoreClient() {
	client := redis.NewClient(&redis.Options{
		Addr:               "",
		Dialer:             nil,
		OnConnect:          nil,
		Username:           "",
		Password:           "",
		DB:                 0,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolFIFO:           false,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
		Limiter:            nil,
	})
	_ =client
}