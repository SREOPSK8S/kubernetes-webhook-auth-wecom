package config

import (
	_ "github.com/spf13/viper"
)
func InitAndLoad() {
	viper.SetConfigName("config")
}
