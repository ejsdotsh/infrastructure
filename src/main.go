package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// ntbx "github.com/ejsdotsh/infrastructure/src/netbox"
)

var (
	doDomains = []string{
		"ejs.sh",
		"ejs.wtf",
		"sahala.org",
		"saha.la",
		"unicorns.wtf",
	}
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create base infrastructure in Netbox Cloud
		// if err := ntbx.SetupNetbox(ctx); err != nil {
		// 	return err
		// }

		// Create DigitalOcean infrastructure
		// 		Provision and configure DNS
		// 		Provision and configure Droplets

		// Create unicornLANd infrastructure
		// 		Provision and configure OPNSense firewall
		// 		Provision and configure Proxmox VE
		// 		Provision and configure NAS
		// 		Provision and configure switches

		return nil
	})
}
