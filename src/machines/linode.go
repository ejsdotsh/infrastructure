i// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package machines

import (
	"fmt"

	"github.com/pulumi/pulumi-linode/sdk/v5/go/linode"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// manageLinodeMachines creates a LinodeMachine component for the given machine config.
func manageLinodeMachines(ctx *pulumi.Context, machine Machine) error {
	_, err := NewLinodeMachine(ctx, fmt.Sprintf("machine-%s", string(machine.Name)), &LinodeMachineArgs{
		// The original resource name must be preserved for alias matching.
		InstanceResourceName: "machine-linode01",
		Label:                machine.Name,
		Region:               machine.Region,
		InstanceType:         machine.Size,
		IPv4Addresses: pulumi.StringArray{
			machine.IP4[0],
			machine.IP4[1],
		},
		PrivateIP: machine.PrivateIP,
		Alerts: &linode.InstanceAlertsArgs{
			Cpu:           pulumi.Int(90),
			Io:            pulumi.Int(10000),
			NetworkIn:     pulumi.Int(10),
			NetworkOut:    pulumi.Int(10),
			TransferQuota: pulumi.Int(80),
		},
		BootConfigLabel: pulumi.String("My Debian 10 Disk Profile"),
		Booted:          pulumi.Bool(true),
		Configs: linode.InstanceConfigTypeArray{
			&linode.InstanceConfigTypeArgs{
				Devices: &linode.InstanceConfigDevicesArgs{
					Sda: &linode.InstanceConfigDevicesSdaArgs{
						DiskId: pulumi.Int(39567975),
					},
					Sdb: &linode.InstanceConfigDevicesSdbArgs{
						DiskId: pulumi.Int(39567976),
					},
				},
				Helpers: &linode.InstanceConfigHelpersArgs{
					DevtmpfsAutomount: pulumi.Bool(true),
					Network:           pulumi.Bool(false),
				},
				Kernel:     pulumi.String("linode/grub2"),
				Label:      pulumi.String("My Debian 10 Disk Profile"),
				RootDevice: pulumi.String("/dev/sda"),
			},
		},
		DiskEncryption: pulumi.String("disabled"),
		Disks: linode.InstanceDiskTypeArray{
			&linode.InstanceDiskTypeArgs{
				Filesystem: pulumi.String("ext4"),
				Label:      pulumi.String("Debian 10 Disk"),
				Size:       pulumi.Int(25088),
			},
			&linode.InstanceDiskTypeArgs{
				Filesystem: pulumi.String("swap"),
				Label:      pulumi.String("512 MB Swap Image"),
				Size:       pulumi.Int(512),
			},
		},
	})
	if err != nil {
		return err
	}

	return nil
}
