package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var AppConfig *config

type config struct {
	Database Database `mapstructure:"database"`
}

type Database struct {
	Postgres Postgres `mapstucture:"postgres"`
}
type Postgres struct {
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Pass         string `mapstructure:"pass"`
	DBNAME       string `mapstructure:"database_name"`
	SSLMODE      string `mapstructure:"ssl_mode"`
	MaxOpenConns int    `mapstructure:"max_open_cons"`
	MaxIdleConns int    `mapstructure:"max_idle_cons"`
	Timeout      string `mapstructure:"timeout"`
}

func LoadConfigFile(path string) {
	viper.SetConfigFile("configs")
	viper.SetConfigType("json")
	if path == "" {
		viper.AddConfigPath("../confis")
		viper.AddConfigPath("./configs")
	} else {
		viper.AddConfigPath(path)
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	AppConfig = &config{}
	if err := viper.Unmarshal(&AppConfig); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
