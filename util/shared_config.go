package util

import (
	"github.com/spf13/viper"
)

type SharedConfig struct {
	Environment     string `mapstructure:"ENVIRONMENT"`
	DBDriver        string `mapstructure:"DB_DRIVER"`
	DBSource        string `mapstructure:"DB_SOURCE"`
	RedisAddress    string `mapstructure:"REDIS_ADDRESS"`
	BookRepoAddress string `mapstructure:"BOOK_REPO_SVC_ADDRESS"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config SharedConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
