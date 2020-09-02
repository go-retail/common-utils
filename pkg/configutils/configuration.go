package configutils

import (
	"fmt"

	"github.com/go-retail/common-utils/pkg/logutils"
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
			logutils.FailOnError(err, "Error reading Config file")
		}
	}
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())
}
