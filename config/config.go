package config

import (
	"github.com/SREOPSK8S/kubernetes-webhook-auth-wecom/entity/wecom"
	"github.com/spf13/viper"
	"os"
	"strconv"
)

func InitAndLoad() {
	viper.SetConfigName("auth-wecom")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./conf")
	viper.AddConfigPath("../")
	if err := viper.ReadInConfig(); err != nil {
		panic("load config file error")
	}
}

func GetServerPort() (servicePort int) {
	tmpPort := os.Getenv("SERVICE_PORT")
	portVip := viper.GetInt("server.port")
	if tmpPort == "" && portVip != 0 {
		return portVip
	}
	portStr, err := strconv.Atoi(tmpPort)
	if err != nil {
		return portVip
	}
	servicePort = portStr
	if servicePort != 0 {
		return servicePort
	}
	return viper.GetInt("server.port")
}

func GetCorpID() (corpID string) {
	corpIDEnv := os.Getenv(wecom.CorpIDEnvKey)
	if corpIDEnv != "" {
		corpID = corpIDEnv
		return
	}
	corpID = viper.GetString("WeCom.CorpID")
	return
}

func GetCorpSecret() (corpSecret string) {
	corpSecretEnv := os.Getenv(wecom.CorpSecretEnvKey)
	if corpSecretEnv != "" {
		corpSecret = corpSecretEnv
		return
	}
	corpSecret = viper.GetString("WeCom.CorpSecret")
	return
}

func GetAgentId() (agentID int) {
	agentIDEnv := os.Getenv(wecom.AgentIDEnvKey)
	if agentIDEnv != "" {
		tmpID, err := strconv.Atoi(agentIDEnv)
		if err != nil {
			return
		}
		agentID = tmpID
		return
	}
	return viper.GetInt("WeCom.AgentId")
}

func GetEtcdEndpoints() (endpoints []string) {
	return viper.GetStringSlice("etcd.endpoints")
}
