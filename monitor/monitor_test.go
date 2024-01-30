package monitor

import (
	"testing"
)

// Test for monitor.go

func TestLoadMonitorsEmpty(t *testing.T) {
	// Define empty monitor config yamlObject map[string]interface{}
	yamlObject := map[string]interface{}{}

	yamlObject["monitors"] = []interface{}{}

	monitors := LoadMonitors(yamlObject)

	if len(monitors) != 0 {
		t.Errorf("Expected empty list of monitors, got %v", monitors)
	}

}

func TestLoadMonitors1(t *testing.T) {
	// Define empty monitor config yamlObject map[string]interface{}
	yamlObject := map[string]interface{}{}

	url := "github.com/test.yaml"
	format := "yaml"
	key := "version"
	version := "1.0.0"

	yamlObject["monitors"] = []interface{}{
		map[string]interface{}{
			"kind":       "FormatedRemoteFile",
			"apiversion": "v1",
			"metadata": map[string]interface{}{
				"name": "test",
				"labels": map[string]interface{}{
					"test": "test",
				},
			},
			"spec": map[string]interface{}{
				"url":     url,
				"format":  format,
				"key":     key,
				"version": version,
			},
		},
	}

	monitors := LoadMonitors(yamlObject)

	if len(monitors) != 1 {
		t.Errorf("Expected 1 item list of monitors, got %v", monitors)
	}

	if monitors[0].Kind != "FormatedRemoteFile" {
		t.Errorf("Expected monitor kind to be FormatedRemoteFile, got %v", monitors[0].Kind)
	}

	if monitors[0].APIVersion != "v1" {
		t.Errorf("Expected monitor apiversion to be v1, got %v", monitors[0].APIVersion)
	}

	if monitors[0].Metadata.Name != "test" {
		t.Errorf("Expected monitor metadata name to be test, got %v", monitors[0].Metadata.Name)
	}

	if monitors[0].Metadata.Labels["test"] != "test" {
		t.Errorf("Expected monitor metadata labels test to be test, got %v", monitors[0].Metadata.Labels["test"])
	}

	if monitors[0].Spec["url"] != url {
		t.Errorf("Expected monitor spec url to be %v", url)
	}

	if monitors[0].Spec["format"] != format {
		t.Errorf("Expected monitor spec format to be %v", format)
	}

	if monitors[0].Spec["key"] != key {
		t.Errorf("Expected monitor spec key to be %v", key)
	}

	if monitors[0].Spec["version"] != version {
		t.Errorf("Expected monitor spec version to be %v", version)
	}

}
