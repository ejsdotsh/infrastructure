// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package dns

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// MXRecord defines an MX DNS record.
type MXRecord struct {
	// ResourceSuffix is appended to the base resource name (e.g., "-mx1").
	ResourceSuffix string
	Priority       int
	Target         string
}

// TXTRecord defines a TXT DNS record.
type TXTRecord struct {
	// ResourceSuffix is appended to the base resource name (e.g., "-txt-spf").
	ResourceSuffix string
	Name           string // Record name (empty for root)
	Target         string
}

// CNAMERecord defines a CNAME DNS record.
type CNAMERecord struct {
	// ResourceSuffix is appended to the base resource name (e.g., "-cname-dkim1").
	ResourceSuffix string
	Name           string
	Target         string
}

// ManageDomains sets up the DNS domains and records.
func ManageDomains(ctx *pulumi.Context) error {
	// Create domains in DigitalOcean
	if err := manageDigitalOceanDNS(ctx); err != nil {
		fmt.Printf("there was an error: %v\n", err)
		return err
	}

	// Create domains in Linode
	if err := manageLinodeDNS(ctx); err != nil {
		fmt.Printf("there was an error: %v\n", err)
		return err
	}

	return nil
}
