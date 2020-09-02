package configutils

import (
	"fmt"

	"github.com/spf13/viper"

	log "github.com/go-retail/common-utils/pkg/logutils"
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
			log.FailOnError(err, "Error reading Config file")
		}
	}
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())
}
