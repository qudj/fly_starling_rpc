package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Name string `mapstructure:"name"`
	Host struct {
		BaiduApi string `mapstructure:"baidu_api"`
	} `mapstructure:"host"`
	DB struct {
		StarlingRead DBConfig `mapstructure:"starling_read"`
		StarlingWrite     DBConfig `mapstructure:"starling_write"`
	} `mapstructure:"db"`
	Redis struct {
	} `mapstructure:"redis"`
}

var Global *Config

func InitConfFile() {
	configName := os.Getenv("conf_file")
	if configName == "" {
		configName = "config"
	}

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	viper.SetConfigName(configName)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Global)
	if err != nil {
		panic(err)
	}
	Global.Name = configName
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

func InitConfig()  {
	InitConfFile()
	InitDb()
}