// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package machines

import (
	"fmt"
	"strconv"

	"github.com/ejsdotsh/infrastructure/src/loader"

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

// NewLinodeMachine creates a new LinodeMachine component from a loader.Machine.
func NewLinodeMachine(ctx *pulumi.Context, name string, machine loader.Machine, opts ...pulumi.ResourceOption) (*LinodeMachine, error) {
	component := &LinodeMachine{}
	if err := ctx.RegisterComponentResource("ejsdotsh:machines:LinodeMachine", name, component, opts...); err != nil {
		return nil, err
	}

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

	instance, err := linode.NewInstance(ctx, fmt.Sprintf("linode-instance-%s", machine.Name), instanceArgs,
		pulumi.Parent(component),
	)
	if err != nil {
		return nil, err
	}

	linodeID := instance.ID().ApplyT(func(id pulumi.ID) (int, error) {
		return strconv.Atoi(string(id))
	}).(pulumi.IntOutput)

	if machine.MachineConfig != nil {
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

		var devices linode.InstanceConfigDeviceArray
		for deviceName, diskSuffix := range machine.MachineConfig.DeviceMap {
			disk, ok := diskResources[diskSuffix]
			if !ok {
				return nil, fmt.Errorf("config device %q references unknown disk suffix %q", deviceName, diskSuffix)
			}
			devices = append(devices, &linode.InstanceConfigDeviceArgs{
				DeviceName: pulumi.String(deviceName),
				DiskId:     disk.ID().ApplyT(func(id pulumi.ID) (int, error) { return strconv.Atoi(string(id)) }).(pulumi.IntOutput),
			})
		}

		helpers := linode.InstanceConfigHelperArray{}
		if machine.MachineConfig.Helpers.DevtmpfsAutomount != nil || machine.MachineConfig.Helpers.Network != nil {
			helperArgs := &linode.InstanceConfigHelperArgs{}
			if machine.MachineConfig.Helpers.DevtmpfsAutomount != nil {
				helperArgs.DevtmpfsAutomount = pulumi.Bool(*machine.MachineConfig.Helpers.DevtmpfsAutomount)
			}
			if machine.MachineConfig.Helpers.Network != nil {
				helperArgs.Network = pulumi.Bool(*machine.MachineConfig.Helpers.Network)
			}
			helpers = linode.InstanceConfigHelperArray{helperArgs}
		}

		_, err = linode.NewInstanceConfig(ctx, fmt.Sprintf("linode-config-%s%s", machine.Name, machine.MachineConfig.Suffix), &linode.InstanceConfigArgs{
			LinodeId:   linodeID,
			Label:      pulumi.String(machine.MachineConfig.Label),
			Kernel:     pulumi.String(machine.MachineConfig.Kernel),
			RootDevice: pulumi.String(machine.MachineConfig.RootDevice),
			Booted:     pulumi.Bool(machine.MachineConfig.Booted),
			Device:     devices,
			Helpers:    helpers,
		}, pulumi.Parent(component))
		if err != nil {
			return nil, err
		}
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
