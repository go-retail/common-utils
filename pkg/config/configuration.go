package config

import (
	"fmt"

	"github.com/go-retail/pos-server/pkg/utils"
	"github.com/spf13/viper"
)

//GetConfig ..
func GetConfig() {
	//TODO may be , I should parameterize these path
	viper.SetConfigName("config")
	viper.AddConfigPath("conf")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		} else {
			utils.FailOnError(err, "Error reading Config file")
		}
	}
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())
}
