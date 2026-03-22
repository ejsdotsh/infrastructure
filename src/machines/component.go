// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package machines

import (
	"fmt"
	"strconv"

	"github.com/pulumi/pulumi-linode/sdk/v5/go/linode"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// DiskDef defines a disk to be managed as a standalone InstanceDisk resource.
type DiskDef struct {
	// ResourceSuffix is appended to the component name to form the disk resource name.
	ResourceSuffix string
	// Label is the disk label.
	Label string
	// Size is the disk size in MB.
	Size int
	// Filesystem is the disk filesystem type (e.g., "ext4", "swap").
	Filesystem string
	// ImportID is the "linodeId,diskId" string used to import the existing disk.
	// Leave empty if the disk is being created fresh.
	ImportID string
}

// ConfigDef defines a config profile to be managed as a standalone InstanceConfig resource.
type ConfigDef struct {
	// ResourceSuffix is appended to the component name to form the config resource name.
	ResourceSuffix string
	// Label is the config profile label.
	Label string
	// Kernel is the kernel to boot (e.g., "linode/grub2").
	Kernel string
	// RootDevice is the root device path (e.g., "/dev/sda").
	RootDevice string
	// Booted indicates whether the instance should be booted into this config.
	Booted bool
	// Helpers configures the config profile helpers.
	Helpers *linode.InstanceConfigHelperArgs
	// DeviceMap maps device names (e.g., "sda", "sdb") to DiskDef ResourceSuffixes.
	// The disk IDs are resolved from the created InstanceDisk resources.
	DeviceMap map[string]string
	// ImportID is the "linodeId,configId" string used to import the existing config.
	// Leave empty if the config is being created fresh.
	ImportID string
}

// LinodeMachineArgs defines the inputs for the LinodeMachine component.
type LinodeMachineArgs struct {
	// InstanceResourceName is the logical Pulumi resource name for the Linode instance.
	// It must match the original name to preserve state via aliases.
	InstanceResourceName string
	// Label is the Linode instance label.
	Label pulumi.StringInput
	// Region is the Linode region (e.g., "us-central").
	Region pulumi.StringInput
	// InstanceType is the Linode plan type (e.g., "g6-nanode-1").
	InstanceType pulumi.StringInput
	// PrivateIP enables a private IP address on the instance.
	PrivateIP pulumi.BoolInput
	// Alerts configures the instance alert thresholds.
	Alerts *linode.InstanceAlertsArgs
	// DiskEncryption sets the disk encryption mode.
	DiskEncryption pulumi.StringInput
	// Disks defines the instance disks as standalone InstanceDisk resources.
	Disks []DiskDef
	// Config defines the instance config profile as a standalone InstanceConfig resource.
	Config ConfigDef
}

// LinodeMachine is a component resource that groups a Linode compute instance,
// its disks, and its configuration profile under a single logical unit.
type LinodeMachine struct {
	pulumi.ResourceState

	// InstanceID is the Linode instance ID.
	InstanceID pulumi.IDOutput `pulumi:"instanceId"`
	// Label is the instance label.
	Label pulumi.StringOutput `pulumi:"label"`
	// IPv4 contains the instance's IPv4 addresses.
	IPv4 pulumi.StringArrayOutput `pulumi:"ipv4"`
}

// NewLinodeMachine creates a new LinodeMachine component resource.
func NewLinodeMachine(ctx *pulumi.Context, name string, args *LinodeMachineArgs, opts ...pulumi.ResourceOption) (*LinodeMachine, error) {
	component := &LinodeMachine{}
	err := ctx.RegisterComponentResource("ejsdotsh:machines:LinodeMachine", name, component, opts...)
	if err != nil {
		return nil, err
	}

	// Create the Linode instance as a child of this component, with an alias
	// pointing to the old stack-root URN so Pulumi recognizes the existing resource.
	// Deprecated embedded configs, disks, bootConfigLabel, and ipv4s have been removed.
	instance, err := linode.NewInstance(ctx, args.InstanceResourceName, &linode.InstanceArgs{
		Alerts:         args.Alerts,
		DiskEncryption: args.DiskEncryption,
		Label:          args.Label,
		PrivateIp:      args.PrivateIP,
		Region:         args.Region,
		Type:           args.InstanceType,
	}, pulumi.Parent(component),
		pulumi.Protect(true),
		pulumi.Aliases([]pulumi.Alias{{NoParent: pulumi.Bool(true)}}),
	)
	if err != nil {
		return nil, err
	}

	// Convert the instance ID from IDOutput to IntOutput for use in disk/config args.
	linodeID := instance.ID().ApplyT(func(id pulumi.ID) (int, error) {
		return strconv.Atoi(string(id))
	}).(pulumi.IntOutput)

	// Create standalone InstanceDisk resources, keyed by ResourceSuffix for config device mapping.
	diskResources := make(map[string]*linode.InstanceDisk, len(args.Disks))
	for _, d := range args.Disks {
		diskName := fmt.Sprintf("%s%s", name, d.ResourceSuffix)
		diskOpts := []pulumi.ResourceOption{pulumi.Parent(component)}
		if d.ImportID != "" {
			diskOpts = append(diskOpts, pulumi.Import(pulumi.ID(d.ImportID)))
		}
		disk, err := linode.NewInstanceDisk(ctx, diskName, &linode.InstanceDiskArgs{
			Label:      pulumi.String(d.Label),
			LinodeId:   linodeID,
			Size:       pulumi.Int(d.Size),
			Filesystem: pulumi.String(d.Filesystem),
		}, diskOpts...)
		if err != nil {
			return nil, err
		}
		diskResources[d.ResourceSuffix] = disk
	}

	// Build the device list for the InstanceConfig from the DeviceMap.
	var devices linode.InstanceConfigDeviceArray
	for deviceName, diskSuffix := range args.Config.DeviceMap {
		disk, ok := diskResources[diskSuffix]
		if !ok {
			return nil, fmt.Errorf("config device %q references unknown disk suffix %q", deviceName, diskSuffix)
		}
		devices = append(devices, &linode.InstanceConfigDeviceArgs{
			DeviceName: pulumi.String(deviceName),
			DiskId:     disk.ID().ApplyT(func(id pulumi.ID) (int, error) { return strconv.Atoi(string(id)) }).(pulumi.IntOutput),
		})
	}

	// Create the standalone InstanceConfig resource.
	cfgName := fmt.Sprintf("%s%s", name, args.Config.ResourceSuffix)
	cfgOpts := []pulumi.ResourceOption{pulumi.Parent(component)}
	if args.Config.ImportID != "" {
		cfgOpts = append(cfgOpts, pulumi.Import(pulumi.ID(args.Config.ImportID)))
	}
	helpers := linode.InstanceConfigHelperArray{}
	if args.Config.Helpers != nil {
		helpers = linode.InstanceConfigHelperArray{args.Config.Helpers}
	}
	_, err = linode.NewInstanceConfig(ctx, cfgName, &linode.InstanceConfigArgs{
		LinodeId:   linodeID,
		Label:      pulumi.String(args.Config.Label),
		Kernel:     pulumi.String(args.Config.Kernel),
		RootDevice: pulumi.String(args.Config.RootDevice),
		Booted:     pulumi.Bool(args.Config.Booted),
		Device:     devices,
		Helpers:    helpers,
	}, cfgOpts...)
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
