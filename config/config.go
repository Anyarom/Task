package config

import "github.com/spf13/viper"

type (
	Config struct {
		Quantity int
	}
)

// заполнение конфига
func InitConfig(file string) (*Config, error) {
	config := &Config{}

	viper.SetConfigFile(file)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config.Quantity = viper.GetInt("worker.quantity")

	return config, nil
}
