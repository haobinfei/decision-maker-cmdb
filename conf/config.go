package conf

import (
	"log"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func InitConfig() {
	Config = viper.New()
	Config.AddConfigPath("./conf/")
	Config.SetConfigName("config")
	Config.SetConfigType("toml")

	if err := Config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalln("配置文件不存在")
		} else {
			log.Fatalln("日志文件格式错误")
		}
	}

}
