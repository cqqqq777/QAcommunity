package boot

import (
	"fmt"
	viper2 "github.com/spf13/viper"
	g "main/global"
)

const (
	configFile = "config/config.yaml"
)

func ViperSetup() {
	viper := viper2.New()
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed err:%v", err))
	}
	if err := viper.Unmarshal(&g.Config); err != nil {
		panic(fmt.Errorf("unmarshal config failed err:%v", err))
	}
}
