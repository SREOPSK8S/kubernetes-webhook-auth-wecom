package config

import (
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
	if tmpPort == "" && portVip != 0  {
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
	corpIDEnv := os.Getenv("CORP_ID")
	if corpIDEnv != "" {
		corpID = corpIDEnv
		return
	}
	corpID = viper.GetString("WeCom.CorpID")
	return
}

func GetCorpSecret() (corpSecret string ){
	corpSecretEnv :=  os.Getenv("CORP_SECRET")
	if corpSecretEnv != "" {
		corpSecret = corpSecretEnv
		return
	}
	corpSecret = viper.GetString("WeCom.CorpSecret")
	return
}
