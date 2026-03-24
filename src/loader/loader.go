// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package loader

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// LoadMachines reads machine definitions from a YAML file.
func LoadMachines(path string) ([]Machine, error) {
	var machines []Machine
	if err := loadYAML(path, &machines); err != nil {
		return nil, fmt.Errorf("loading machines from %s: %w", path, err)
	}
	return machines, nil
}

// LoadDomains reads DNS domain definitions from a YAML file.
func LoadDomains(path string) ([]Domain, error) {
	var domains []Domain
	if err := loadYAML(path, &domains); err != nil {
		return nil, fmt.Errorf("loading domains from %s: %w", path, err)
	}
	return domains, nil
}

// loadYAML reads a YAML file and unmarshals it into the provided target.
// Returns nil if the file is empty or contains only comments.
func loadYAML(path string, target interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(data, target); err != nil {
		return fmt.Errorf("parsing YAML: %w", err)
	}
	return nil
}
