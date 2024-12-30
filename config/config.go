package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

var AppConfig *config

type config struct {
	Database Database `mapstructure:"database"`
	General  General  `mapstructure:"general"`
}

type Database struct {
	Postgres Postgres `mapstucture:"postgres"`
}
type General struct {
	Listen           string        `mapstructure:"listen"`           // rest listen port
	LogLevel         int8          `mapstructure:"log_level"`        // logger level
	ShutdownTimeout  time.Duration `mapstructure:"shutdown_timeout"` // shutdown timeout
	DictionariesPath string        `mapstructure:"dictionaries_path"`
	ExportsPath      string        `mapstructure:"exports_path"`
	InternalMCC      string        `mapstructure:"internal_mcc"`
	InternalMNC      string        `mapstructure:"internal_mnc"`
	Postpaid         bool          `mapstructure:"postpaid"`
}
type Postgres struct {
	Host         string        `mapstructure:"host"`
	Port         string        `mapstructure:"port"`
	User         string        `mapstructure:"user"`
	Pass         string        `mapstructure:"pass"`
	DBNAME       string        `mapstructure:"database_name"`
	SSLMODE      string        `mapstructure:"ssl_mode"`
	MaxOpenConns int           `mapstructure:"max_open_cons"`
	MaxIdleConns int           `mapstructure:"max_idle_cons"`
	Timeout      time.Duration `mapstructure:"timeout"`
}

func LoadConfigFile(path string) {
	viper.SetConfigName("configs")
	viper.SetConfigType("json")
	if path == "" {
		viper.AddConfigPath("/home/amir/GitHub/phoonebook/config")
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
