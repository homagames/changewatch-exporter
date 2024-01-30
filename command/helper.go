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
	// Set prometheus port as 8080 by default
	v.SetDefault("prometheus.port", 8080)
	// Set prometheus path as /metrics by default
	v.SetDefault("prometheus.path", "/metrics")
	// Set mode as "daemon" by default
	v.SetDefault("mode", "daemon")
	// Set interval as 60 seconds by default
	v.SetDefault("interval", 60)

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return v, nil
}
