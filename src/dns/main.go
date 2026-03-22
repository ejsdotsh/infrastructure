// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package dns

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ManageDomains sets up the DNS domains and records.
func ManageDomains(ctx *pulumi.Context) error {
	// Create domains in DigitalOcean
	if err := manageDigitalOceanDNS(ctx); err != nil {
		return err
	}

	// Create domains in Linode
	if err := manageLinodeDNS(ctx); err != nil {
		return err
	}

	return nil
}
