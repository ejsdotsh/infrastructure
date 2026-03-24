// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package dns

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ejsdotsh/infrastructure/src/loader"

	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi-linode/sdk/v5/go/linode"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// LinodeDNS is a component resource that groups a Linode domain and its
// associated DNS records under a single logical unit.
type LinodeDNS struct {
	pulumi.ResourceState

	DomainID   pulumi.IntOutput    `pulumi:"domainId"`
	DomainName pulumi.StringOutput `pulumi:"domainName"`
}

// NewLinodeDNS creates a new LinodeDNS component from a loader.LinodeDomain.
func NewLinodeDNS(ctx *pulumi.Context, name string, domain loader.LinodeDomain, opts ...pulumi.ResourceOption) (*LinodeDNS, error) {
	component := &LinodeDNS{}
	if err := ctx.RegisterComponentResource("ejsdotsh:dns:LinodeDNS", name, component, opts...); err != nil {
		return nil, err
	}

	// Derive consistent resource name from the domain.
	domainSlug := strings.ReplaceAll(domain.Domain, ".", "-")

	// Create the Linode domain.
	d, err := linode.NewDomain(ctx, fmt.Sprintf("linode-domain-%s", domainSlug), &linode.DomainArgs{
		Domain:   pulumi.String(domain.Domain),
		SoaEmail: pulumi.String(domain.SoaEmail),
		Type:     pulumi.String("master"),
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	// Convert domain ID for use in record args.
	domainID := d.ID().ApplyT(func(id pulumi.ID) (int, error) {
		return strconv.Atoi(string(id))
	}).(pulumi.IntOutput)

	// Create MX records.
	for i, mx := range domain.MX {
		_, err := linode.NewDomainRecord(ctx, fmt.Sprintf("linode-dns-%s-mx-%d", domainSlug, i), &linode.DomainRecordArgs{
			DomainId:   domainID,
			Priority:   pulumi.Int(mx.Priority),
			RecordType: pulumi.String("MX"),
			Target:     pulumi.String(mx.Target),
		}, pulumi.Parent(component))
		if err != nil {
			return nil, err
		}
	}

	// Create TXT records.
	for i, txt := range domain.TXT {
		args := &linode.DomainRecordArgs{
			DomainId:   domainID,
			RecordType: pulumi.String("TXT"),
			Target:     pulumi.String(txt.Target),
		}
		if txt.Name != "" {
			args.Name = pulumi.String(txt.Name)
		}
		_, err := linode.NewDomainRecord(ctx, fmt.Sprintf("linode-dns-%s-txt-%d", domainSlug, i), args, pulumi.Parent(component))
		if err != nil {
			return nil, err
		}
	}

	// Create CNAME records.
	for i, cname := range domain.CNAME {
		_, err := linode.NewDomainRecord(ctx, fmt.Sprintf("linode-dns-%s-cname-%d", domainSlug, i), &linode.DomainRecordArgs{
			DomainId:   domainID,
			Name:       pulumi.String(cname.Name),
			RecordType: pulumi.String("CNAME"),
			Target:     pulumi.String(cname.Target),
		}, pulumi.Parent(component))
		if err != nil {
			return nil, err
		}
	}

	// Create A records.
	for i, a := range domain.A {
		args := &linode.DomainRecordArgs{
			DomainId:   domainID,
			Name:       pulumi.String(a.Name),
			RecordType: pulumi.String("A"),
			Target:     pulumi.String(a.Target),
		}
		if a.TTL > 0 {
			args.TtlSec = pulumi.Int(a.TTL)
		}
		_, err := linode.NewDomainRecord(ctx, fmt.Sprintf("linode-dns-%s-a-%d", domainSlug, i), args, pulumi.Parent(component))
		if err != nil {
			return nil, err
		}
	}

	// Create AAAA records.
	for i, aaaa := range domain.AAAA {
		args := &linode.DomainRecordArgs{
			DomainId:   domainID,
			Name:       pulumi.String(aaaa.Name),
			RecordType: pulumi.String("AAAA"),
			Target:     pulumi.String(aaaa.Target),
		}
		if aaaa.TTL > 0 {
			args.TtlSec = pulumi.Int(aaaa.TTL)
		}
		_, err := linode.NewDomainRecord(ctx, fmt.Sprintf("linode-dns-%s-aaaa-%d", domainSlug, i), args, pulumi.Parent(component))
		if err != nil {
			return nil, err
		}
	}

	component.DomainID = domainID
	component.DomainName = d.Domain

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"domainId":   domainID,
		"domainName": d.Domain,
	}); err != nil {
		return nil, err
	}

	return component, nil
}

// DigitalOceanDNS is a component resource that groups a DigitalOcean domain
// and its associated DNS records under a single logical unit.
type DigitalOceanDNS struct {
	pulumi.ResourceState

	DomainName pulumi.StringOutput `pulumi:"domainName"`
}

// NewDigitalOceanDNS creates a new DigitalOceanDNS component from a loader.DODomain.
func NewDigitalOceanDNS(ctx *pulumi.Context, name string, domain loader.DODomain, opts ...pulumi.ResourceOption) (*DigitalOceanDNS, error) {
	component := &DigitalOceanDNS{}
	if err := ctx.RegisterComponentResource("ejsdotsh:dns:DigitalOceanDNS", name, component, opts...); err != nil {
		return nil, err
	}

	// Derive consistent resource name from the domain.
	domainSlug := strings.ReplaceAll(domain.Domain, ".", "-")

	// Create the DigitalOcean domain.
	d, err := digitalocean.NewDomain(ctx, fmt.Sprintf("do-domain-%s", domainSlug), &digitalocean.DomainArgs{
		Name: pulumi.String(domain.Domain),
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	// Create MX records.
	for i, mx := range domain.MX {
		_, err := digitalocean.NewDnsRecord(ctx, fmt.Sprintf("do-dns-%s-mx-%d", domainSlug, i), &digitalocean.DnsRecordArgs{
			Domain:   d.Name,
			Type:     pulumi.String(digitalocean.RecordTypeMX),
			Name:     pulumi.String("@"),
			Value:    pulumi.String(mx.Target),
			Priority: pulumi.Int(mx.Priority),
			Ttl:      pulumi.Int(14400),
		}, pulumi.Parent(component))
		if err != nil {
			return nil, err
		}
	}

	// Create NS records.
	for i, ns := range domain.NS {
		_, err := digitalocean.NewDnsRecord(ctx, fmt.Sprintf("do-dns-%s-ns-%d", domainSlug, i), &digitalocean.DnsRecordArgs{
			Domain: d.Name,
			Type:   pulumi.String(digitalocean.RecordTypeNS),
			Name:   pulumi.String("@"),
			Value:  pulumi.String(ns.Target),
			Ttl:    pulumi.Int(1800),
		}, pulumi.Parent(component))
		if err != nil {
			return nil, err
		}
	}

	// Create TXT records.
	for i, txt := range domain.TXT {
		args := &digitalocean.DnsRecordArgs{
			Domain: d.Name,
			Type:   pulumi.String(digitalocean.RecordTypeTXT),
			Name:   pulumi.String("@"),
			Value:  pulumi.String(txt.Target),
			Ttl:    pulumi.Int(3600),
		}
		if txt.Name != "" {
			args.Name = pulumi.String(txt.Name)
		}
		_, err := digitalocean.NewDnsRecord(ctx, fmt.Sprintf("do-dns-%s-txt-%d", domainSlug, i), args, pulumi.Parent(component))
		if err != nil {
			return nil, err
		}
	}

	// Create CNAME records.
	for i, cname := range domain.CNAME {
		_, err := digitalocean.NewDnsRecord(ctx, fmt.Sprintf("do-dns-%s-cname-%d", domainSlug, i), &digitalocean.DnsRecordArgs{
			Domain: d.Name,
			Type:   pulumi.String(digitalocean.RecordTypeCNAME),
			Name:   pulumi.String(cname.Name),
			Value:  pulumi.String(cname.Target),
			Ttl:    pulumi.Int(43200),
		}, pulumi.Parent(component))
		if err != nil {
			return nil, err
		}
	}

	// Create A records.
	for i, a := range domain.A {
		args := &digitalocean.DnsRecordArgs{
			Domain: d.Name,
			Type:   pulumi.String(digitalocean.RecordTypeA),
			Name:   pulumi.String(a.Name),
			Value:  pulumi.String(a.Target),
		}
		if a.TTL > 0 {
			args.Ttl = pulumi.Int(a.TTL)
		}
		_, err := digitalocean.NewDnsRecord(ctx, fmt.Sprintf("do-dns-%s-a-%d", domainSlug, i), args, pulumi.Parent(component))
		if err != nil {
			return nil, err
		}
	}

	// Create AAAA records.
	for i, aaaa := range domain.AAAA {
		args := &digitalocean.DnsRecordArgs{
			Domain: d.Name,
			Type:   pulumi.String(digitalocean.RecordTypeAAAA),
			Name:   pulumi.String(aaaa.Name),
			Value:  pulumi.String(aaaa.Target),
		}
		if aaaa.TTL > 0 {
			args.Ttl = pulumi.Int(aaaa.TTL)
		}
		_, err := digitalocean.NewDnsRecord(ctx, fmt.Sprintf("do-dns-%s-aaaa-%d", domainSlug, i), args, pulumi.Parent(component))
		if err != nil {
			return nil, err
		}
	}

	component.DomainName = d.Name

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"domainName": d.Name,
	}); err != nil {
		return nil, err
	}

	return component, nil
}
