// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package dns

import (
	"encoding/json"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

const (
	// DomainComponentType is the ID of the Pulumi component type.
	DomainComponentType = "ejsdotsh:dns:Domain"
)

// Domain is a simple definition for a high-level DNS resource.
type Domain struct {
	pulumi.ResourceState

	// Name       string `json:"name"` // Domain name
	// DNSRecords []DomainRecord
}

// DomainRecord represents a DNS record.
type DomainRecord struct {
	RecordResourceName pulumi.String    `json:"resourceName"` // manually generated the resource names
	DomainId           pulumi.IntOutput // the ID of the Domain
	Domain             string           // Domain name
	Type               pulumi.String    `json:"recordType"`  // Record type (A, AAAA, CNAME, etc.)
	Value              pulumi.String    `json:"recordValue"` // Record value
	Name               pulumi.String    `json:"recordName"`  // Name of the record
	Ttl                int              // Time to live for DNS records
}

// ManageDomains sets up the DNS domains and records.
func ManageDomains(ctx *pulumi.Context) error {
	// create a new Pulumi Config for Domains
	config := config.New(ctx, "")
	var domains []Domain
	jsonString := config.Get("domains")
	if err := json.Unmarshal([]byte(jsonString), &domains); err != nil {
		return err
	}

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

// NewDomain is a high-level abstraction that creates a Domain resource with
// the appropriate provider.
func NewDomain(ctx *pulumi.Context, name string, args *DomainSpec, opts ...pulumi.ResourceOption) (*Domain, error) {
	component := &Domain{}
	if err := ctx.RegisterComponentResourceV2(DomainComponentType, name, component, opts...); err != nil {
		return nil, err
	}

	return component, nil
}
