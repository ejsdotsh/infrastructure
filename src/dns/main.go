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

// ManageDomains sets up the DNS domains and records.
func ManageDomains(ctx *pulumi.Context) error {
	// Create domains in DigitalOcean
	if err := createDODomains(ctx); err != nil {
		return err
	}

	// Add records for SPF, DKIM, and DMARC
	if err := updateRecords(ctx); err != nil {
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

// createDODomains creates the domains in DigitalOcean, as well as their MX and NS records.
func createDODomains(ctx *pulumi.Context) error {
	for _, domainName := range doDomains {
		// Replace dots with hyphens for resource naming
		resourceName := fmt.Sprintf("domain-%s", strings.ReplaceAll(domainName, ".", "-"))
		_default, err := digitalocean.NewDomain(ctx, resourceName, &digitalocean.DomainArgs{
			Name: pulumi.String(domainName),
		})
		if err != nil {
			return err
		}

		// Add MX records
		resourceName = fmt.Sprintf("domain-record-%s", strings.ReplaceAll(domainName, ".", "-"))
		_, err = digitalocean.NewDnsRecord(ctx, resourceName+"-mx1", &digitalocean.DnsRecordArgs{
			Domain:   _default.ID(),
			Type:     pulumi.String(digitalocean.RecordTypeMX),
			Name:     pulumi.String("@"),
			Value:    pulumi.String("mail.protonmail.ch."),
			Priority: pulumi.Int(10),
			Ttl:      pulumi.Int(14400),
		})
		if err != nil {
			return err
		}
		_, err = digitalocean.NewDnsRecord(ctx, resourceName+"-mx2", &digitalocean.DnsRecordArgs{
			Domain:   _default.ID(),
			Type:     pulumi.String(digitalocean.RecordTypeMX),
			Name:     pulumi.String("@"),
			Value:    pulumi.String("mailsec.protonmail.ch."),
			Priority: pulumi.Int(20),
			Ttl:      pulumi.Int(14400),
		})
		if err != nil {
			return err
		}

		// Add NS records
		for i := 1; i <= 3; i++ {
			_, err = digitalocean.NewDnsRecord(ctx, resourceName+fmt.Sprintf("-ns%d", i), &digitalocean.DnsRecordArgs{
				Domain: _default.ID(),
				Type:   pulumi.String(digitalocean.RecordTypeNS),
				Name:   pulumi.String("@"),
				Value:  pulumi.String(fmt.Sprintf("ns%d.digitalocean.com.", i)),
				Ttl:    pulumi.Int(1800),
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Add records for imported verify, SPF, DKIM, and DMARC
func updateRecords(ctx *pulumi.Context) error {
	_, err := digitalocean.NewDnsRecord(ctx, "domain-record-panemorfos-me-txt-protonmail-verification", &digitalocean.DnsRecordArgs{
		Domain: pulumi.String("panemorfos.me"),
		Name:   pulumi.String("@"),
		Ttl:    pulumi.Int(3600),
		Type:   pulumi.String(digitalocean.RecordTypeTXT),
		Value:  pulumi.String("protonmail-verification=fdfc3be39cbcc7aad30939cc525ca6c3ee38f61b"),
	})
	if err != nil {
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
		return err
	}
	return nil
}
