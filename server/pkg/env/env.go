package env

import (
	"github.com/spf13/viper"
	"github.com/ssibrahimbas/ssi-we/pkg/helper"
)

func LoadEnv(path string, c interface{}) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	helper.CheckErr(err)
	err = viper.Unmarshal(&c)
	helper.CheckErr(err)
}
