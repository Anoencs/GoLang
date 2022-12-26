package util

import "github.com/spf13/viper"

type Config struct {
	DBSource   string `mapstructure:"DB_SOURCE"`
	DBNAME     string `mapstructure:"DB_NAME"`
	LOG_FILE   string `mapstructure:"LOG_FILE"`
	OKR_CHECK  string `mapstructure:"OKR_CHECK"`
	SHEET_SKIP string `mapstructure:"SHEET_SKIP"`
	File_SKIP  string `mapstructure:"FILE_SKIP"`
	BOD_ROLE   string `mapstructure:"BOD_ROLE"`
}

func LoadConfig(path string) (config Config, err error) {
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
