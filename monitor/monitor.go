package monitor

import (
    "log/slog"
	"os"
    "gopkg.in/yaml.v3"
)
// See ./monitor/README.md for more information.

// Helper: https://zhwt.github.io/yaml-to-go/
type Monitor struct {
	Kind       string `yaml:"kind"`
	APIVersion string `yaml:"apiversion"`
	Metadata   struct {
		Name   string `yaml:"name"`
		Labels map[string]any `yaml:"labels"`
	} `yaml:"metadata"`
	Spec map[string]any `yaml:"spec"`
}

// Load all monitors from a yaml object
func LoadMonitors(yamlObject map[string]interface{}) []*Monitor {
	monitors := []*Monitor{}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))


	for _, monitor := range yamlObject["monitors"].([]interface{}) {
		s, _ := yaml.Marshal(monitor)
		var m Monitor
		err := yaml.Unmarshal([]byte(s), &m)
		if err != nil {
			panic(err)
		}
		logger.Info("Monitor loaded", "monitor", m)
		monitors = append(monitors, &m)
	}

	return monitors
}