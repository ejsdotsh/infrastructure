// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package machines

import (
	"github.com/pulumi/pulumi-linode/sdk/v5/go/linode"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
	// IPv4Addresses are the IPv4 addresses assigned to the instance.
	IPv4Addresses pulumi.StringArray
	// PrivateIP enables a private IP address on the instance.
	PrivateIP pulumi.BoolInput
	// Alerts configures the instance alert thresholds.
	Alerts *linode.InstanceAlertsArgs
	// BootConfigLabel is the label of the boot configuration profile.
	BootConfigLabel pulumi.StringInput
	// Booted indicates whether the instance should be booted.
	Booted pulumi.BoolInput
	// Configs defines the instance configuration profiles.
	Configs linode.InstanceConfigTypeArray
	// DiskEncryption sets the disk encryption mode.
	DiskEncryption pulumi.StringInput
	// Disks defines the instance disk layout.
	Disks linode.InstanceDiskTypeArray
}

// LinodeMachine is a component resource that groups a Linode compute instance
// and its configuration under a single logical unit.
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
	instance, err := linode.NewInstance(ctx, args.InstanceResourceName, &linode.InstanceArgs{
		Alerts:          args.Alerts,
		BootConfigLabel: args.BootConfigLabel,
		Booted:          args.Booted,
		Configs:         args.Configs,
		DiskEncryption:  args.DiskEncryption,
		Disks:           args.Disks,
		Ipv4s:           args.IPv4Addresses,
		Label:           args.Label,
		PrivateIp:       args.PrivateIP,
		Region:          args.Region,
		Type:            args.InstanceType,
	}, pulumi.Parent(component),
		pulumi.Protect(true),
		pulumi.Aliases([]pulumi.Alias{{NoParent: pulumi.Bool(true)}}),
	)
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
