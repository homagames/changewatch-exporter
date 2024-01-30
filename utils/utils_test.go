package utils

import (
	"testing"
)

func TestComputeDeltaNone1(t *testing.T) {
	// Define empty monitor config yamlObject map[string]interface{}
	v1 := "1.0.0"
	v2 := "1.0.0"

	delta := ComputeDelta(v1, v2)

	if delta != "none" {
		t.Errorf("Expected delta to be none, got %v", delta)
	}
}
func TestComputeDeltaNone2(t *testing.T) {
	// Define empty monitor config yamlObject map[string]interface{}
	v1 := "1.0.0"
	v2 := "v1.0.0"

	delta := ComputeDelta(v1, v2)

	if delta != "none" {
		t.Errorf("Expected delta to be none, got %v", delta)
	}
}

func TestComputeDeltaUnknown1(t *testing.T) {
	// Define empty monitor config yamlObject map[string]interface{}
	v1 := "aze"
	v2 := ""

	delta := ComputeDelta(v1, v2)

	if delta != "unknown" {
		t.Errorf("Expected delta to be unknown, got %v", delta)
	}
}

func TestComputeDeltaUnknown2(t *testing.T) {
	// Define empty monitor config yamlObject map[string]interface{}
	v1 := "v1.2.3"
	v2 := "v3.2.1"

	delta := ComputeDelta(v1, v2)

	if delta != "unknown" {
		t.Errorf("Expected delta to be unknown, got %v", delta)
	}
}
func TestComputeDeltaUnknown3(t *testing.T) {
	// Define empty monitor config yamlObject map[string]interface{}
	v1 := "1.2.3"
	v2 := "3.2.1"

	delta := ComputeDelta(v1, v2)

	if delta != "unknown" {
		t.Errorf("Expected delta to be unknown, got %v", delta)
	}
}

func TestComputeDeltaUnknown4(t *testing.T) {
	// Define empty monitor config yamlObject map[string]interface{}
	v1 := "4"
	v2 := "3"

	delta := ComputeDelta(v1, v2)

	if delta != "unknown" {
		t.Errorf("Expected delta to be unknown, got %v", delta)
	}
}
func TestComputeDeltaUnknown5(t *testing.T) {
	// Define empty monitor config yamlObject map[string]interface{}
	v1 := "4"
	v2 := "4"

	delta := ComputeDelta(v1, v2)

	if delta != "unknown" {
		t.Errorf("Expected delta to be unknown, got %v", delta)
	}
}

func TestComputeDeltaMajor1(t *testing.T) {
	// Define empty monitor config yamlObject map[string]interface{}
	v1 := "4.0.0"
	v2 := "3.0.0"

	delta := ComputeDelta(v1, v2)

	if delta != "major" {
		t.Errorf("Expected delta to be unknown, got %v", delta)
	}
}

func TestComputeDeltaMajor2(t *testing.T) {
	// Define empty monitor config yamlObject map[string]interface{}
	v1 := "v4.0.0"
	v2 := "v3.0.0"

	delta := ComputeDelta(v1, v2)

	if delta != "major" {
		t.Errorf("Expected delta to be unknown, got %v", delta)
	}
}

func TestComputeDeltaMajor3(t *testing.T) {
	// Define empty monitor config yamlObject map[string]interface{}
	v1 := "4.0.0"
	v2 := "v3.0.0"

	delta := ComputeDelta(v1, v2)

	if delta != "major" {
		t.Errorf("Expected delta to be unknown, got %v", delta)
	}
}

func TestComputeDeltaMinor1(t *testing.T) {
	// Define empty monitor config yamlObject map[string]interface{}
	v1 := "4.5.0"
	v2 := "4.0.0"

	delta := ComputeDelta(v1, v2)

	if delta != "minor" {
		t.Errorf("Expected delta to be unknown, got %v", delta)
	}
}
func TestComputeDeltaMinor2(t *testing.T) {
	// Define empty monitor config yamlObject map[string]interface{}
	v1 := "4.5.0"
	v2 := "v4.0.0"

	delta := ComputeDelta(v1, v2)

	if delta != "minor" {
		t.Errorf("Expected delta to be unknown, got %v", delta)
	}
}

func TestComputeDeltaPatch1(t *testing.T) {
	// Define empty monitor config yamlObject map[string]interface{}
	v1 := "4.0.7"
	v2 := "4.0.0"

	delta := ComputeDelta(v1, v2)

	if delta != "patch" {
		t.Errorf("Expected delta to be unknown, got %v", delta)
	}
}
func TestComputeDeltaPatch2(t *testing.T) {
	// Define empty monitor config yamlObject map[string]interface{}
	v1 := "v4.0.7"
	v2 := "v4.0.0"

	delta := ComputeDelta(v1, v2)

	if delta != "patch" {
		t.Errorf("Expected delta to be unknown, got %v", delta)
	}
}
func TestComputeDeltaPatch3(t *testing.T) {
	// Define empty monitor config yamlObject map[string]interface{}
	v1 := "4.0.7"
	v2 := "v4.0.0"

	delta := ComputeDelta(v1, v2)

	if delta != "patch" {
		t.Errorf("Expected delta to be unknown, got %v", delta)
	}
}
