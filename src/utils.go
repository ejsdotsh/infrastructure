// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"
)

// checkRequiredEnvVars checks that all required environment variables are set.
func checkRequiredEnvVars() error {
	required := []string{
		"DIGITALOCEAN_TOKEN",
	}
	missingVars := []string{}
	for _, v := range required {
		if os.Getenv(v) == "" {
			missingVars = append(missingVars, v)
		}
	}
	if len(missingVars) > 0 {
		return fmt.Errorf("missing required environment variables: %v", missingVars)
	}
	return nil
}
