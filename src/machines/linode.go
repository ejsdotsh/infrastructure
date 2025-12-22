package machines

import (
	"github.com/pulumi/pulumi-linode/sdk/v5/go/linode"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// manageLinodeMachines is based on the initial import of the resource.
func manageLinodeMachines(ctx *pulumi.Context, machine Machine) error {

	_, err := linode.NewInstance(ctx, "machine-linode01", &linode.InstanceArgs{
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
		Ipv4s: pulumi.StringArray{
			machine.IP4[0],
			machine.IP4[1],
		},
		Label:     machine.Name,
		PrivateIp: machine.PrivateIP,
		Region:    machine.Region,
		Type:      machine.Size,
	}, pulumi.Protect(true))
	if err != nil {
		return err
	}
	return nil
}
