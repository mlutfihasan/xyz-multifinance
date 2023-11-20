package helper

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Configuration struct {
	DbDsnSource1      string `mapstructure:"DB_DSN_SOURCE_1"`
	DbDsnReplication1 string `mapstructure:"DB_DSN_Replication_1"`
	Port              string `mapstructure:"PORT"`
}

func LoadConfig() (config Configuration, err error) {
	viper.SetConfigFile("./configuration/.env")
	err = viper.ReadInConfig()
	godotenv.Load("./configuration/.env")
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
