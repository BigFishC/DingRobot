package config

import "github.com/spf13/viper"

//ParseConfig 解析配置文件
func ParseConfig(configpath string) (string, string) {
	//设置配置文件的名字和路径
	viper.SetConfigName("config")
	viper.AddConfigPath(configpath)

	//读取配置文件
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
	return viper.GetString("server.token"), viper.GetString("server.appsecret")
}
