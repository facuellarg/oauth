package config

import (
	"sync"

	"github.com/spf13/viper"
)

type Server struct {
	Host string
	Port string
}

var (
	config *Config
	mut    sync.Mutex
)

type OAuthCredentials struct {
	ClientID     string
	ClientSecret string
}
type Config struct {
	AppServer                Server
	GithubOAuthCredentials   OAuthCredentials
	GoogleOAuthCredentials   OAuthCredentials
	LinkedInOAuthCredentials OAuthCredentials
}

func ReadConfig() (*Config, error) {
	// viper.AddConfigPath("./infrastructure/config")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	config = &Config{}
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}
	return config, nil
}

func GetConfig() *Config {
	var err error
	mut.Lock()
	defer mut.Unlock()
	if config == nil {
		config, err = ReadConfig()
		if err != nil {
			panic(err)
		}
	}
	return config
}
