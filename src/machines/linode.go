// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
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
		PrivateIP:            machine.PrivateIP,
		Alerts: &linode.InstanceAlertsArgs{
			Cpu:           pulumi.Int(90),
			Io:            pulumi.Int(10000),
			NetworkIn:     pulumi.Int(10),
			NetworkOut:    pulumi.Int(10),
			TransferQuota: pulumi.Int(80),
		},
		DiskEncryption: pulumi.String("disabled"),
		// Standalone InstanceDisk resources replace the deprecated embedded disks.
		Disks: []DiskDef{
			{
				ResourceSuffix: "-disk-boot",
				Label:          "Debian 10 Disk",
				Size:           25088,
				Filesystem:     "ext4",
				ImportID:       "19333075,39567975",
			},
			{
				ResourceSuffix: "-disk-swap",
				Label:          "512 MB Swap Image",
				Size:           512,
				Filesystem:     "swap",
				ImportID:       "19333075,39567976",
			},
		},
		// Standalone InstanceConfig resource replaces the deprecated embedded config.
		Config: ConfigDef{
			ResourceSuffix: "-config",
			Label:          "My Debian 10 Disk Profile",
			Kernel:         "linode/grub2",
			RootDevice:     "/dev/sda",
			Booted:         true,
			Helpers: &linode.InstanceConfigHelperArgs{
				DevtmpfsAutomount: pulumi.Bool(true),
				Network:           pulumi.Bool(false),
			},
			DeviceMap: map[string]string{
				"sda": "-disk-boot",
				"sdb": "-disk-swap",
			},
			ImportID: "19333075,20647231",
		},
	})
	if err != nil {
		return err
	}

	return nil
}
