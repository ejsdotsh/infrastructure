// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package dns

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ejsdotsh/infrastructure/src/loader"

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

// NewLinodeDNS creates a new LinodeDNS component from a loader.Domain.
func NewLinodeDNS(ctx *pulumi.Context, name string, domain loader.Domain, opts ...pulumi.ResourceOption) (*LinodeDNS, error) {
	component := &LinodeDNS{}
	if err := ctx.RegisterComponentResource("ejsdotsh:dns:LinodeDNS", name, component, opts...); err != nil {
		return nil, err
	}

	domainSlug := strings.ReplaceAll(domain.DomainName, ".", "-")

	d, err := linode.NewDomain(ctx, fmt.Sprintf("linode-domain-%s", domainSlug), &linode.DomainArgs{
		Domain:   pulumi.String(domain.DomainName),
		SoaEmail: pulumi.String(domain.SoaEmail),
		Type:     pulumi.String("master"),
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	domainID := d.ID().ApplyT(func(id pulumi.ID) (int, error) {
		return strconv.Atoi(string(id))
	}).(pulumi.IntOutput)

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
