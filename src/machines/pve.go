package machines

import (
	"os"

	"github.com/muhlba91/pulumi-proxmoxve/sdk/v7/go/proxmoxve"
	// "github.com/muhlba91/pulumi-proxmoxve/sdk/v7/go/proxmoxve/ct"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// "github.com/muhlba91/pulumi-proxmoxve/sdk/v7/go/proxmoxve/download"
	// "github.com/pulumi/pulumi-random/sdk/v4/go/random"
	// "github.com/pulumi/pulumi-std/sdk/go/std"
	// "github.com/pulumi/pulumi-tls/sdk/v5/go/tls"
)

// setupProxmoxProvider returns a new PVE provider
func setupProxmoxProvider(ctx *pulumi.Context) (*proxmoxve.Provider, error) {
	provider, err := proxmoxve.NewProvider(ctx, "pve-provider", &proxmoxve.ProviderArgs{
		// Username: pulumi.String(os.Getenv("PROXMOX_VE_USERNAME")),
		ApiToken: pulumi.String(os.Getenv("PROXMOX_VE_PASSWORD")),
		Endpoint: pulumi.String(os.Getenv("PROXMOX_VE_ENDPOINT")),
		Insecure: pulumi.Bool(true), // true for self-signed certs
	})
	if err != nil {
		return nil, err
	}
	return provider, nil
}

func createProxmoxMachines(ctx *pulumi.Context, provider *proxmoxve.Provider) error {

	// container, err := proxmoxve.GetContainer(ctx, "pulse", 101, *ct.ContainerState{}, []ct.ResourceOptions{})
	// if err != nil {
	// 	return err
	// }

	// fmt.Printf("the result of GetContainer:\n\n\t%v\n", container)

	// debian13LxcImg, err := download.NewFile(ctx, "debian_13_lxc_img", &download.FileArgs{
	// 		ContentType: pulumi.String("vztmpl"),
	// 		DatastoreId: pulumi.String("local-lvm"),
	// 		NodeName:    pulumi.String("pve"),
	// 		Url:         pulumi.String("https://mirrors.servercentral.com/ubuntu-cloud-images/releases/25.04/release/ubuntu-25.04-server-cloudimg-amd64-root.tar.xz"),
	// 	})
	// 	if err != nil {
	// 		return err
	// 	}
	// 	ubuntuContainerPassword, err := random.NewRandomPassword(ctx, "ubuntu_container_password", &random.RandomPasswordArgs{
	// 		Length:          pulumi.Int(16),
	// 		OverrideSpecial: pulumi.String("_%@"),
	// 		Special:         pulumi.Bool(true),
	// 	})
	// 	if err != nil {
	// 		return err
	// 	}
	// 	ubuntuContainerKey, err := tls.NewPrivateKey(ctx, "ubuntu_container_key", &tls.PrivateKeyArgs{
	// 		Algorithm: pulumi.String("RSA"),
	// 		RsaBits:   pulumi.Int(2048),
	// 	})
	// 	if err != nil {
	// 		return err
	// 	}
	// 	_, err = ct.NewContainer(ctx, "ubuntu_container", &ct.ContainerArgs{
	// 		Description:  pulumi.String("Managed by Pulumi"),
	// 		NodeName:     pulumi.String("pve"),
	// 		VmId:         pulumi.Int(101),
	// 		Unprivileged: pulumi.Bool(true),
	// 		Features: &ct.ContainerFeaturesArgs{
	// 			Nesting: pulumi.Bool(false),
	// 		},
	// 		Initialization: &ct.ContainerInitializationArgs{
	// 			Hostname: pulumi.String("pulse"),
	// 			IpConfigs: ct.ContainerInitializationIpConfigArray{
	// 				&ct.ContainerInitializationIpConfigArgs{
	// 					Ipv4: &ct.ContainerInitializationIpConfigIpv4Args{
	// 						Address: pulumi.String("dhcp"),
	// 					},
	// 				},
	// 			},
	// 			UserAccount: &ct.ContainerInitializationUserAccountArgs{
	// 				Keys: pulumi.StringArray{
	// 					std.TrimspaceOutput(ctx, std.TrimspaceOutputArgs{
	// 						Input: ubuntuContainerKey.PublicKeyOpenssh,
	// 					}, nil).ApplyT(func(invoke std.TrimspaceResult) (*string, error) {
	// 						return invoke.Result, nil
	// 					}).(pulumi.StringPtrOutput),
	// 				},
	// 				Password: ubuntuContainerPassword.Result,
	// 			},
	// 		},
	// 		NetworkInterfaces: ct.ContainerNetworkInterfaceArray{
	// 			&ct.ContainerNetworkInterfaceArgs{
	// 				Name: pulumi.String("eth0"),
	// 			},
	// 		},
	// 		Disk: &ct.ContainerDiskArgs{
	// 			DatastoreId: pulumi.String("local-lvm"),
	// 			Size:        pulumi.Int(4),
	// 		},
	// 		OperatingSystem: &ct.ContainerOperatingSystemArgs{
	// 			TemplateFileId: debian13LxcImg.ID(),
	// 			Type:           pulumi.String("debian"),
	// 		},
	// 		MountPoints: ct.ContainerMountPointArray{
	// 			&ct.ContainerMountPointArgs{
	// 				Volume: pulumi.String("/mnt/bindmounts/shared"),
	// 				Path:   pulumi.String("/mnt/shared"),
	// 			},
	// 			&ct.ContainerMountPointArgs{
	// 				Volume: pulumi.String("local-lvm"),
	// 				Size:   pulumi.String("10G"),
	// 				Path:   pulumi.String("/mnt/volume"),
	// 			},
	// 		},
	// 		Startup: &ct.ContainerStartupArgs{
	// 			Order:     pulumi.Int(3),
	// 			UpDelay:   pulumi.Int(60),
	// 			DownDelay: pulumi.Int(60),
	// 		},
	// 	})
	// 	if err != nil {
	// 		return err
	// 	}
	// 	// ctx.Export("ubuntuContainerPassword", ubuntuContainerPassword.Result)
	// 	// ctx.Export("ubuntuContainerPrivateKey", ubuntuContainerKey.PrivateKeyPem)
	// 	// ctx.Export("ubuntuContainerPublicKey", ubuntuContainerKey.PublicKeyOpenssh)
	// 	return nil
	// })

	return nil
}
