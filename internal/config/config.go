package config

import (
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)


type Config struct {
	App    AppConfig
	Postgres PostgresConfig
}

func LoadConfig() (*Config, error){
	viper.BindEnv("consul_url")
	viper.BindEnv("consul_path")

	consulUrl := viper.GetString("consul_url")
	consulPath := viper.GetString("consul_path")


	viper.SetConfigType("yaml")
	viper.AddRemoteProvider("consul", consulUrl, consulPath)
	err := viper.ReadRemoteConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = viper.UnmarshalKey("app", &config.App)
	if err != nil {
		return nil, err
	}

	err = viper.UnmarshalKey("postgres", &config.Postgres)
	if err != nil {
		return nil, err
	}

	return config, nil
}