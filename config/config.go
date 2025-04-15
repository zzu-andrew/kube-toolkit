package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zzu-andrew/toolkit/pkg/errors"
	"log"
)

type Config struct {
	Type     string `mapstructure:"type"`
	Address  string `mapstructure:"address"`
	Timeout  int    `mapstructure:"timeout"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

var globalConfig Config

func GetConfig() *Config {
	return &globalConfig
}

func ParseConfig() error {
	configFile := viper.GetString("configFile")
	if configFile == "" {
		return fmt.Errorf("no config file specified")
	}
	viper.AddConfigPath(".") // add the current directory as a search path
	viper.SetConfigFile(configFile)
	// load the xml config file info
	err := viper.ReadInConfig()
	if err != nil {
		return errors.ErrorNew("Error reading config file: %v", err)
	}

	// 将配置文件绑定到结构体
	if err := viper.Unmarshal(&globalConfig); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
		return errors.ErrInvalidConfig
	}
	return nil
}
