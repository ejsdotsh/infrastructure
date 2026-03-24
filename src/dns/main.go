// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

// Package dns manages DNS domains and records across providers.
package dns

import (
	"fmt"
	"strings"

	"github.com/ejsdotsh/infrastructure/src/loader"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ManageDomains provisions all DNS domains and records from the loaded data.
func ManageDomains(ctx *pulumi.Context, linodeDomains []loader.LinodeDomain, doDomains []loader.DODomain) error {
	// Provision Linode-hosted domains.
	for _, domain := range linodeDomains {
		slug := strings.ReplaceAll(domain.Domain, ".", "-")
		_, err := NewLinodeDNS(ctx, fmt.Sprintf("dns-linode-%s", slug), domain)
		if err != nil {
			return fmt.Errorf("linode domain %s: %w", domain.Domain, err)
		}
	}

	// Provision DigitalOcean-hosted domains.
	for _, domain := range doDomains {
		slug := strings.ReplaceAll(domain.Domain, ".", "-")
		_, err := NewDigitalOceanDNS(ctx, fmt.Sprintf("dns-do-%s", slug), domain)
		if err != nil {
			return fmt.Errorf("digitalocean domain %s: %w", domain.Domain, err)
		}
	}

	return nil
}
