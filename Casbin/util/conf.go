package utils

import "github.com/spf13/viper"

type Config struct {
	GRPCPort        string `mapstructure:"GRPCPort"`
	DbDriverName    string `mapstructure:"DbDriverName"`
	DbConnectString string `mapstructure:"DbConnectString"`
	ModelPath       string `mapstructure:"ModelPath"`
}

func ReadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
