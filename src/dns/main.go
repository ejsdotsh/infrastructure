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

// ManageDomains provisions all DNS domains and records, dispatching by provider.
func ManageDomains(ctx *pulumi.Context, domains []loader.Domain) error {
	for _, domain := range domains {
		slug := strings.ReplaceAll(domain.DomainName, ".", "-")
		switch domain.Provider {
		case loader.ProviderLinode:
			_, err := NewLinodeDNS(ctx, fmt.Sprintf("dns-linode-%s", slug), domain)
			if err != nil {
				return fmt.Errorf("linode domain %s: %w", domain.DomainName, err)
			}
		case loader.ProviderDigitalOcean:
			_, err := NewDigitalOceanDNS(ctx, fmt.Sprintf("dns-do-%s", slug), domain)
			if err != nil {
				return fmt.Errorf("digitalocean domain %s: %w", domain.DomainName, err)
			}
		default:
			return fmt.Errorf("unknown DNS provider %q for domain %s", domain.Provider, domain.DomainName)
		}
	}
	return nil
}
