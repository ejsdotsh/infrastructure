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

	return nil
}

// Domain is a struct representing a DNS domain.
type Domain struct {
	Name string // Domain name
}

// DomainRecord represents a DNS record.
type DomainRecord struct {
	Domain string // Domain name
	Type   string // Record type (A, AAAA, CNAME, etc.)
	Value  string // Record value
	Name   string // Hostname of the record
	Ttl    int    // Time to live for DNS records
}
