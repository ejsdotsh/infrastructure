package main

import (
	unicornsdns "github.com/ejsdotsh/infrastructure/src/dns"
	// ntbx "github.com/ejsdotsh/infrastructure/src/netbox"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Check that required environment variables are set
		if err := checkRequiredEnvVars(); err != nil {
			return err
		}

		// TODO - the netbox provider import/sdk creation doesn't work as desired.
		// Create infrastructure in Netbox Cloud
		// if err := ntbx.SetupNetbox(ctx); err != nil {
		// 	return err
		// }

		// Create DNS domains, MX, CNAME DKIM records
		if err := unicornsdns.ManageDomains(ctx); err != nil {
			return err
		}

		// TODO Create unicornsLANd infrastructure
		// 		Provision and configure OPNSense firewall
		// 		Provision and configure Proxmox VE
		// 		Provision and configure NAS
		// 		Provision and configure switches

		// Provision and configure octopik3s Kubernetes cluster
		// if err := unicornsland.SetupUnicornsLANd(ctx); err != nil {
		// 	return err
		// }

		return nil
	})
}
