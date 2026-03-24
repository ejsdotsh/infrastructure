// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package loader

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// LoadLinodeMachines reads Linode machine definitions from a YAML file.
func LoadLinodeMachines(path string) ([]LinodeMachine, error) {
	var machines []LinodeMachine
	if err := loadYAML(path, &machines); err != nil {
		return nil, fmt.Errorf("loading linode machines from %s: %w", path, err)
	}
	return machines, nil
}

// LoadDODroplets reads DigitalOcean Droplet definitions from a YAML file.
func LoadDODroplets(path string) ([]DODroplet, error) {
	var droplets []DODroplet
	if err := loadYAML(path, &droplets); err != nil {
		return nil, fmt.Errorf("loading DO droplets from %s: %w", path, err)
	}
	return droplets, nil
}

// LoadLinodeDomains reads Linode DNS domain definitions from a YAML file.
func LoadLinodeDomains(path string) ([]LinodeDomain, error) {
	var domains []LinodeDomain
	if err := loadYAML(path, &domains); err != nil {
		return nil, fmt.Errorf("loading linode domains from %s: %w", path, err)
	}
	return domains, nil
}

// LoadDODomains reads DigitalOcean DNS domain definitions from a YAML file.
func LoadDODomains(path string) ([]DODomain, error) {
	var domains []DODomain
	if err := loadYAML(path, &domains); err != nil {
		return nil, fmt.Errorf("loading DO domains from %s: %w", path, err)
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
	// yaml.Unmarshal returns nil for empty/comment-only files,
	// which leaves the target at its zero value (empty slice). That's fine.
	if err := yaml.Unmarshal(data, target); err != nil {
		return fmt.Errorf("parsing YAML: %w", err)
	}
	return nil
}
