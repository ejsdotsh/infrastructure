// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package dns

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pulumi/pulumi-linode/sdk/v5/go/linode"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// LinodeDNSArgs defines the inputs for the LinodeDNS component.
type LinodeDNSArgs struct {
	// DomainName is the DNS domain name (e.g., "saha.la").
	DomainName string
	// SoaEmail is the SOA email address for the domain.
	SoaEmail string
	// MXRecords defines MX records to create for this domain.
	MXRecords []MXRecord
	// TXTRecords defines TXT records to create for this domain.
	TXTRecords []TXTRecord
	// CNAMERecords defines CNAME records to create for this domain.
	CNAMERecords []CNAMERecord
}

// LinodeDNS is a component resource that groups a Linode domain and its
// associated DNS records (MX, TXT, CNAME) under a single logical unit.
type LinodeDNS struct {
	pulumi.ResourceState

	// DomainID is the Linode domain ID.
	DomainID pulumi.IntOutput `pulumi:"domainId"`
	// DomainName is the domain name managed by this component.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
}

// NewLinodeDNS creates a new LinodeDNS component resource.
func NewLinodeDNS(ctx *pulumi.Context, name string, args *LinodeDNSArgs, opts ...pulumi.ResourceOption) (*LinodeDNS, error) {
	component := &LinodeDNS{}
	err := ctx.RegisterComponentResource("ejsdotsh:dns:LinodeDNS", name, component, opts...)
	if err != nil {
		fmt.Printf("there was an error: %v\n", err)
		return nil, err
	}

	// Derive the resource name prefix used by the original flat code.
	domainSlug := strings.ReplaceAll(args.DomainName, ".", "-")
	domainResourceName := fmt.Sprintf("domain-%s", domainSlug)
	recordResourceBase := fmt.Sprintf("domain-record-%s", domainSlug)

	// Create the Linode domain as a child of this component, with an alias
	// pointing to the old stack-root URN so Pulumi recognizes the existing resource.
	domain, err := linode.NewDomain(ctx, domainResourceName, &linode.DomainArgs{
		Domain:   pulumi.String(args.DomainName),
		SoaEmail: pulumi.String(args.SoaEmail),
		Type:     pulumi.String("master"),
	}, pulumi.Parent(component), pulumi.Aliases([]pulumi.Alias{{NoParent: pulumi.Bool(true)}}))
	if err != nil {
		fmt.Printf("there was an error: %v\n", err)
		return nil, err
	}

	// Convert the domain ID from IDOutput to IntOutput for use in DomainRecord args.
	domainID := domain.ID().ApplyT(func(id pulumi.ID) (int, error) {
		i, err := strconv.Atoi(string(id))
		if err != nil {
			return 0, err
		}
		return i, nil
	}).(pulumi.IntOutput)

	// Create MX records.
	for _, mx := range args.MXRecords {
		recordName := recordResourceBase + mx.ResourceSuffix
		_, err := linode.NewDomainRecord(ctx, recordName, &linode.DomainRecordArgs{
			DomainId:   domainID,
			Priority:   pulumi.Int(mx.Priority),
			RecordType: pulumi.String("MX"),
			Target:     pulumi.String(mx.Target),
		}, pulumi.Parent(component), pulumi.Aliases([]pulumi.Alias{{NoParent: pulumi.Bool(true)}}))
		if err != nil {
			fmt.Printf("there was an error: %v\n", err)
			return nil, err
		}
	}

	// Create TXT records.
	for _, txt := range args.TXTRecords {
		recordName := recordResourceBase + txt.ResourceSuffix
		recordArgs := &linode.DomainRecordArgs{
			DomainId:   domainID,
			RecordType: pulumi.String("TXT"),
			Target:     pulumi.String(txt.Target),
		}
		if txt.Name != "" {
			recordArgs.Name = pulumi.String(txt.Name)
		}
		_, err := linode.NewDomainRecord(ctx, recordName, recordArgs,
			pulumi.Parent(component), pulumi.Aliases([]pulumi.Alias{{NoParent: pulumi.Bool(true)}}))
		if err != nil {
			fmt.Printf("there was an error: %v\n", err)
			return nil, err
		}
	}

	// Create CNAME records.
	for _, cname := range args.CNAMERecords {
		recordName := recordResourceBase + cname.ResourceSuffix
		_, err := linode.NewDomainRecord(ctx, recordName, &linode.DomainRecordArgs{
			DomainId:   domainID,
			Name:       pulumi.String(cname.Name),
			RecordType: pulumi.String("CNAME"),
			Target:     pulumi.String(cname.Target),
		}, pulumi.Parent(component), pulumi.Aliases([]pulumi.Alias{{NoParent: pulumi.Bool(true)}}))
		if err != nil {
			fmt.Printf("there was an error: %v\n", err)
			return nil, err
		}
	}

	component.DomainID = domainID
	component.DomainName = domain.Domain

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"domainId":   domainID,
		"domainName": domain.Domain,
	}); err != nil {
		fmt.Printf("there was an error: %v\n", err)
		return nil, err
	}

	return component, nil
}

// manageLinodeDNS creates a LinodeDNS component for each Linode-hosted domain.
func manageLinodeDNS(ctx *pulumi.Context) error {
	for _, cfg := range linodeDomainConfigs {
		_, err := NewLinodeDNS(ctx, fmt.Sprintf("dns-%s", cfg.DomainName), &cfg)
		if err != nil {
			fmt.Printf("there was an error: %v\n", err)
			return err
		}
	}
	return nil
}
