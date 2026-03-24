// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package dns

import (
	"fmt"
	"strings"

	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var (
	doDomains = []string{
		"panemorfos.me",
		"pik3s.io",
		"unicorns.wtf",
	}
)

// manageDigitalOceanDNS creates the domains in DigitalOcean, as well as their MX and NS records.
func manageDigitalOceanDNS(ctx *pulumi.Context) error {
	for _, domainName := range doDomains {
		// Replace dots with hyphens for resource naming
		resourceName := fmt.Sprintf("domain-%s", strings.ReplaceAll(domainName, ".", "-"))
		_default, err := digitalocean.NewDomain(ctx, resourceName, &digitalocean.DomainArgs{
			Name: pulumi.String(domainName),
		})
		if err != nil {
			fmt.Printf("there was an error: %v\n", err)
			return err
		}
		// Add MX records
		resourceName = fmt.Sprintf("domain-record-%s", strings.ReplaceAll(domainName, ".", "-"))
		_, err = digitalocean.NewDnsRecord(ctx, resourceName+"-mx1", &digitalocean.DnsRecordArgs{
			Domain:   _default.Name,
			Type:     pulumi.String(digitalocean.RecordTypeMX),
			Name:     pulumi.String("@"),
			Value:    pulumi.String("mail.protonmail.ch."),
			Priority: pulumi.Int(10),
			Ttl:      pulumi.Int(14400),
		})
		if err != nil {
			fmt.Printf("there was an error: %v\n", err)
			return err
		}
		_, err = digitalocean.NewDnsRecord(ctx, resourceName+"-mx2", &digitalocean.DnsRecordArgs{
			Domain:   _default.Name,
			Type:     pulumi.String(digitalocean.RecordTypeMX),
			Name:     pulumi.String("@"),
			Value:    pulumi.String("mailsec.protonmail.ch."),
			Priority: pulumi.Int(20),
			Ttl:      pulumi.Int(14400),
		})
		if err != nil {
			fmt.Printf("there was an error: %v\n", err)
			return err
		}
		// Add NS records
		for i := 1; i <= 3; i++ {
			_, err = digitalocean.NewDnsRecord(ctx, resourceName+fmt.Sprintf("-ns%d", i), &digitalocean.DnsRecordArgs{
				Domain: _default.Name,
				Type:   pulumi.String(digitalocean.RecordTypeNS),
				Name:   pulumi.String("@"),
				Value:  pulumi.String(fmt.Sprintf("ns%d.digitalocean.com.", i)),
				Ttl:    pulumi.Int(1800),
			})
			if err != nil {
				fmt.Printf("there was an error: %v\n", err)
				return err
			}
		}
	}

	// Add records for imported verify, SPF, DKIM, and DMARC
	_, err := digitalocean.NewDnsRecord(ctx, "domain-record-panemorfos-me-txt-protonmail-verification", &digitalocean.DnsRecordArgs{
		Domain: pulumi.String("panemorfos.me"),
		Name:   pulumi.String("@"),
		Ttl:    pulumi.Int(3600),
		Type:   pulumi.String(digitalocean.RecordTypeTXT),
		Value:  pulumi.String("protonmail-verification=fdfc3be39cbcc7aad30939cc525ca6c3ee38f61b"),
	})
	if err != nil {
		fmt.Printf("there was an error: %v\n", err)
		return err
	}
	_, err = digitalocean.NewDnsRecord(ctx, "domain-record-pik3s-io-txt-protonmail-verification", &digitalocean.DnsRecordArgs{
		Domain: pulumi.String("pik3s.io"),
		Name:   pulumi.String("@"),
		Ttl:    pulumi.Int(3600),
		Type:   pulumi.String(digitalocean.RecordTypeTXT),
		Value:  pulumi.String("protonmail-verification=91adc44d657ca96a9c7327bfe7b9b2dc80b8261b"),
	})
	if err != nil {
		fmt.Printf("there was an error: %v\n", err)
		return err
	}
	_, err = digitalocean.NewDnsRecord(ctx, "domain-record-unicorns-wtf-txt-protonmail-verification", &digitalocean.DnsRecordArgs{
		Domain: pulumi.String("unicorns.wtf"),
		Name:   pulumi.String("@"),
		Ttl:    pulumi.Int(3600),
		Type:   pulumi.String(digitalocean.RecordTypeTXT),
		Value:  pulumi.String("protonmail-verification=e35f918ddaea3eae5ccb81ff300e7dd90713d5e7"),
	})
	if err != nil {
		fmt.Printf("there was an error: %v\n", err)
		return err
	}

	_, err = digitalocean.NewDnsRecord(ctx, "domain-record-unicorns-wtf-txt-spf", &digitalocean.DnsRecordArgs{
		Domain: pulumi.String("unicorns.wtf"),
		Name:   pulumi.String("@"),
		Ttl:    pulumi.Int(3600),
		Type:   pulumi.String(digitalocean.RecordTypeTXT),
		Value:  pulumi.String("v=spf1 include:_spf.protonmail.ch ~all"),
	})
	if err != nil {
		fmt.Printf("there was an error: %v\n", err)
		return err
	}
	_, err = digitalocean.NewDnsRecord(ctx, "domain-record-unicorns-wtf-txt-dmarc", &digitalocean.DnsRecordArgs{
		Domain: pulumi.String("unicorns.wtf"),
		Name:   pulumi.String("_dmarc"),
		Ttl:    pulumi.Int(3600),
		Type:   pulumi.String(digitalocean.RecordTypeTXT),
		Value:  pulumi.String("v=DMARC1; p=quarantine"),
	})
	if err != nil {
		fmt.Printf("there was an error: %v\n", err)
		return err
	}

	// DKIM CNAME records
	_, err = digitalocean.NewDnsRecord(ctx, "domain-record-unicorns-wtf-cname-dkim1", &digitalocean.DnsRecordArgs{
		Domain: pulumi.String("unicorns.wtf"),
		Name:   pulumi.String("protonmail._domainkey"),
		Ttl:    pulumi.Int(43200),
		Type:   pulumi.String(digitalocean.RecordTypeCNAME),
		Value:  pulumi.String("protonmail.domainkey.dryktxtupmrp5coofwzuib32r7l7msvlngqicqwbweu4szlvekd5q.domains.proton.ch."),
	})
	if err != nil {
		fmt.Printf("there was an error: %v\n", err)
		return err
	}
	_, err = digitalocean.NewDnsRecord(ctx, "domain-record-unicorns-wtf-cname-dkim2", &digitalocean.DnsRecordArgs{
		Domain: pulumi.String("unicorns.wtf"),
		Name:   pulumi.String("protonmail2._domainkey"),
		Ttl:    pulumi.Int(43200),
		Type:   pulumi.String(digitalocean.RecordTypeCNAME),
		Value:  pulumi.String("protonmail2.domainkey.dryktxtupmrp5coofwzuib32r7l7msvlngqicqwbweu4szlvekd5q.domains.proton.ch."),
	})
	if err != nil {
		fmt.Printf("there was an error: %v\n", err)
		return err
	}
	_, err = digitalocean.NewDnsRecord(ctx, "domain-record-unicorns-wtf-cname-dkim3", &digitalocean.DnsRecordArgs{
		Domain: pulumi.String("unicorns.wtf"),
		Name:   pulumi.String("protonmail3._domainkey"),
		Ttl:    pulumi.Int(43200),
		Type:   pulumi.String(digitalocean.RecordTypeCNAME),
		Value:  pulumi.String("protonmail3.domainkey.dryktxtupmrp5coofwzuib32r7l7msvlngqicqwbweu4szlvekd5q.domains.proton.ch."),
	})
	if err != nil {
		fmt.Printf("there was an error: %v\n", err)
		return err
	}

	return nil
}
