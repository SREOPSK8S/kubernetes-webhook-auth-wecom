package config

import (
	"github.com/spf13/viper"
	"os"
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

func GetServerPort() int {
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
