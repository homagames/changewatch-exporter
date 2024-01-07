package command

import (
	"github.com/spf13/viper"
)

// Load config
func LoadConfig(configFile string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(configFile)
	// Set logvel as info for viper by defaut
	v.SetDefault("loglevel", "info")
	// Set monitors as empty list by default
	v.SetDefault("monitors", []string{})

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return v, nil
}