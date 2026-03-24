// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"fmt"
	"os"
)

// CheckRequiredEnvVars checks that all required environment variables are set.
func CheckRequiredEnvVars() error {
	required := []string{
		"DIGITALOCEAN_TOKEN",
		"NETBOX_TOKEN",
		"NETBOX_URL",
		// "PROXMOX_VE_ENDPOINT",
		// "PROXMOX_VE_USERNAME",
		// "PROXMOX_VE_PASSWORD",
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
