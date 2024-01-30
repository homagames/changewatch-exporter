package monitor

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/go-github/v58/github"
	"github.com/homagames/changewatch-exporter/logger"
	"github.com/homagames/changewatch-exporter/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
)

// See ./monitor/README.md for more information.

// Define prometheus metrics
var (
	itemVec = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "change",
		Name:      "processed_item",
		Help:      "Indicate if a change has been processed",
	},
		[]string{"name", "kind", "version", "remote_version", "delta"},
	)
)

// Helper: https://zhwt.github.io/yaml-to-go/
type Monitor struct {
	Kind       string `yaml:"kind"`
	APIVersion string `yaml:"apiversion"`
	Metadata   struct {
		Name   string         `yaml:"name"`
		Labels map[string]any `yaml:"labels"`
	} `yaml:"metadata"`
	Spec map[string]any `yaml:"spec"`
}

// Load all monitors from a yaml object
func LoadMonitors(yamlObject map[string]interface{}) []*Monitor {
	monitors := []*Monitor{}

	for _, monitor := range yamlObject["monitors"].([]interface{}) {
		s, _ := yaml.Marshal(monitor)
		var m Monitor
		err := yaml.Unmarshal([]byte(s), &m)
		if err != nil {
			panic(err)
		}
		monitors = append(monitors, &m)
	}

	return monitors
}

func (m *Monitor) Process() error {
	// Determine monitor kind
	switch m.Kind {
	case "FormatedRemoteFile":
		return m.ProcessFormatedRemoteFile()
	case "GithubRelease":
		return m.ProcessGithubRelease()
	default:
		return errors.New("Unknown monitor kind")
	}
}

func (m *Monitor) ProcessFormatedRemoteFile() error {
	url := m.Spec["url"].(string)

	// Download url in memory
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var formatedObject map[string]interface{}

	switch m.Spec["format"].(string) {
	case "yaml":
		err := yaml.Unmarshal(body, &formatedObject)
		if err != nil {
			return err
		}
	case "json":
		return errors.New("JSON format not implemented yet")
	default:
		return errors.New(fmt.Sprintf("Unknown format %v", m.Spec["format"].(string)))
	}

	remote_version := formatedObject[m.Spec["key"].(string)].(string)

	return m.ComputeVersion(remote_version)
}

func (m *Monitor) ProcessGithubRelease() error {

	repo := m.Spec["repo"].(string)
	org := m.Spec["org"].(string)

	//logger := logger.GetLogger()

	client := github.NewClient(nil)

	// Get the latest release from the repository
	release, _, err := client.Repositories.GetLatestRelease(context.Background(), org, repo)

	if err != nil {
		return err
	}

	// extract version from release
	remote_version := *release.TagName

	fmt.Println(remote_version)

	return m.ComputeVersion(remote_version)
}

func (m *Monitor) ComputeVersion(remote_version string) error {

	logger := logger.GetLogger()

	var change float64 = 1

	delta := utils.ComputeDelta(remote_version, m.Spec["version"].(string))

	if delta == "none" {
		change = 0
	}

	item := itemVec.WithLabelValues(m.Metadata.Name, m.Kind, m.Spec["version"].(string), remote_version, delta)
	item.Set(change)

	logger.Info("Processed item", "name", m.Metadata.Name, "change", change)

	return nil
}
