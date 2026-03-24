// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package machines

import (
	"fmt"
	"strconv"

	"github.com/ejsdotsh/infrastructure/src/loader"

	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi-linode/sdk/v5/go/linode"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// LinodeMachine is a component resource that groups a Linode compute instance,
// its disks, and its configuration profile under a single logical unit.
type LinodeMachine struct {
	pulumi.ResourceState

	InstanceID pulumi.IDOutput          `pulumi:"instanceId"`
	Label      pulumi.StringOutput      `pulumi:"label"`
	IPv4       pulumi.StringArrayOutput `pulumi:"ipv4"`
}

// NewLinodeMachine creates a new LinodeMachine component from a loader.LinodeMachine.
func NewLinodeMachine(ctx *pulumi.Context, name string, machine loader.LinodeMachine, opts ...pulumi.ResourceOption) (*LinodeMachine, error) {
	component := &LinodeMachine{}
	if err := ctx.RegisterComponentResource("ejsdotsh:machines:LinodeMachine", name, component, opts...); err != nil {
		return nil, err
	}

	// Build instance args.
	instanceArgs := &linode.InstanceArgs{
		Label:     pulumi.String(machine.Name),
		Region:    pulumi.String(machine.Region),
		Type:      pulumi.String(machine.Type),
		PrivateIp: pulumi.Bool(machine.PrivateIP),
	}
	if machine.DiskEncryption != "" {
		instanceArgs.DiskEncryption = pulumi.String(machine.DiskEncryption)
	}
	if machine.Alerts != nil {
		instanceArgs.Alerts = &linode.InstanceAlertsArgs{
			Cpu:           pulumi.Int(machine.Alerts.CPU),
			Io:            pulumi.Int(machine.Alerts.IO),
			NetworkIn:     pulumi.Int(machine.Alerts.NetworkIn),
			NetworkOut:    pulumi.Int(machine.Alerts.NetworkOut),
			TransferQuota: pulumi.Int(machine.Alerts.TransferQuota),
		}
	}

	// Create the Linode instance.
	instance, err := linode.NewInstance(ctx, fmt.Sprintf("linode-instance-%s", machine.Name), instanceArgs,
		pulumi.Parent(component),
	)
	if err != nil {
		return nil, err
	}

	// Convert instance ID for use in disk/config args.
	linodeID := instance.ID().ApplyT(func(id pulumi.ID) (int, error) {
		return strconv.Atoi(string(id))
	}).(pulumi.IntOutput)

	// Create standalone InstanceDisk resources.
	diskResources := make(map[string]*linode.InstanceDisk, len(machine.Disks))
	for _, d := range machine.Disks {
		diskName := fmt.Sprintf("linode-disk-%s%s", machine.Name, d.Suffix)
		disk, err := linode.NewInstanceDisk(ctx, diskName, &linode.InstanceDiskArgs{
			Label:      pulumi.String(d.Label),
			LinodeId:   linodeID,
			Size:       pulumi.Int(d.Size),
			Filesystem: pulumi.String(d.Filesystem),
		}, pulumi.Parent(component))
		if err != nil {
			return nil, err
		}
		diskResources[d.Suffix] = disk
	}

	// Build the device list for the InstanceConfig.
	var devices linode.InstanceConfigDeviceArray
	for deviceName, diskSuffix := range machine.Config.DeviceMap {
		disk, ok := diskResources[diskSuffix]
		if !ok {
			return nil, fmt.Errorf("config device %q references unknown disk suffix %q", deviceName, diskSuffix)
		}
		devices = append(devices, &linode.InstanceConfigDeviceArgs{
			DeviceName: pulumi.String(deviceName),
			DiskId:     disk.ID().ApplyT(func(id pulumi.ID) (int, error) { return strconv.Atoi(string(id)) }).(pulumi.IntOutput),
		})
	}

	// Build config helpers.
	helpers := linode.InstanceConfigHelperArray{}
	if machine.Config.Helpers.DevtmpfsAutomount != nil || machine.Config.Helpers.Network != nil {
		helperArgs := &linode.InstanceConfigHelperArgs{}
		if machine.Config.Helpers.DevtmpfsAutomount != nil {
			helperArgs.DevtmpfsAutomount = pulumi.Bool(*machine.Config.Helpers.DevtmpfsAutomount)
		}
		if machine.Config.Helpers.Network != nil {
			helperArgs.Network = pulumi.Bool(*machine.Config.Helpers.Network)
		}
		helpers = linode.InstanceConfigHelperArray{helperArgs}
	}

	// Create the standalone InstanceConfig resource.
	_, err = linode.NewInstanceConfig(ctx, fmt.Sprintf("linode-config-%s%s", machine.Name, machine.Config.Suffix), &linode.InstanceConfigArgs{
		LinodeId:   linodeID,
		Label:      pulumi.String(machine.Config.Label),
		Kernel:     pulumi.String(machine.Config.Kernel),
		RootDevice: pulumi.String(machine.Config.RootDevice),
		Booted:     pulumi.Bool(machine.Config.Booted),
		Device:     devices,
		Helpers:    helpers,
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	component.InstanceID = instance.ID()
	component.Label = instance.Label
	component.IPv4 = instance.Ipv4s

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"instanceId": instance.ID(),
		"label":      instance.Label,
		"ipv4":       instance.Ipv4s,
	}); err != nil {
		return nil, err
	}

	return component, nil
}

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
