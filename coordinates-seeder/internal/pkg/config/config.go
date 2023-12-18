package config

import (
	"fmt"
	"reflect"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`

	Broker string `mapstructure:"BROKER"`
	Topic  string `mapstructure:"TOPIC"`
}

// getEnvKeys takes the `mapstructure` tag value from all fields in the Config struct.
func getEnvKeys() []string {
	t := reflect.TypeOf(Config{})
	numField := t.NumField()
	envKeys := make([]string, 0, numField)

	for i := 0; i < numField; i++ {
		envKeys = append(envKeys, t.Field(i).Tag.Get("mapstructure"))
	}

	return envKeys
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("error reading server config from file, falling back to getting from environment variables instead")
		for _, k := range getEnvKeys() {
			_ = viper.BindEnv(k)
		}
	}

	cfg := Config{}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
