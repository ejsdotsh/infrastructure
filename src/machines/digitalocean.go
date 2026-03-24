// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package machines

import (
	"fmt"
	"strconv"

	"github.com/ejsdotsh/infrastructure/src/loader"

	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// DODroplet is a component resource for a DigitalOcean Droplet.
// DNS is managed separately in the dns package.
type DODroplet struct {
	pulumi.ResourceState

	DropletID   pulumi.IntOutput         `pulumi:"dropletId"`
	IPv4Address pulumi.StringOutput      `pulumi:"ipv4Address"`
	IPv6Address pulumi.StringOutput      `pulumi:"ipv6Address"`
	Tags        pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewDODroplet creates a new DODroplet component from a loader.DODroplet.
func NewDODroplet(ctx *pulumi.Context, name string, droplet loader.DODroplet, opts ...pulumi.ResourceOption) (*DODroplet, error) {
	component := &DODroplet{}
	if err := ctx.RegisterComponentResource("ejsdotsh:machines:DODroplet", name, component, opts...); err != nil {
		return nil, err
	}

	// Build tags.
	var tags pulumi.StringArray
	for _, t := range droplet.Tags {
		tags = append(tags, pulumi.String(t))
	}

	// Create the DigitalOcean Droplet.
	d, err := digitalocean.NewDroplet(ctx, fmt.Sprintf("do-droplet-%s", droplet.Name), &digitalocean.DropletArgs{
		Image:  pulumi.String(droplet.Image),
		Region: pulumi.String(droplet.Region),
		Size:   pulumi.String(droplet.Size),
		Ipv6:   pulumi.Bool(droplet.IPv6),
		Tags:   tags,
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	component.DropletID = d.ID().ApplyT(func(id pulumi.ID) (int, error) {
		return strconv.Atoi(string(id))
	}).(pulumi.IntOutput)
	component.IPv4Address = d.Ipv4Address
	component.IPv6Address = d.Ipv6Address
	component.Tags = d.Tags

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"dropletId":   d.ID(),
		"ipv4Address": d.Ipv4Address,
		"ipv6Address": d.Ipv6Address,
		"tags":        d.Tags,
	}); err != nil {
		return nil, err
	}

	return component, nil
}
