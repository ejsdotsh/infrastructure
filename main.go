// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"

	unet "github.com/ejsdotsh/infrastructure/network"
	"github.com/ejsdotsh/infrastructure/src/dns"
	"github.com/ejsdotsh/infrastructure/src/machines"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		ctx.Log.Info(("=== PRE-CHECKS: load ENV vars ==="), nil)
		// Ensure that the required environment variables are set
		// if err := CheckRequiredEnvVars(); err != nil {
		// 	ctx.Log.Error((fmt.Sprintf("=== PHASE 1: ERROR ===\n\n%v", err)), nil)
		// 	panic(err)
		// }

		// Initialize the Netbox client (reads NETBOX_URL/TOKEN from env)
		ctx.Log.Info(("=== PHASE 1: initialize inventory client ==="), nil)
		// ntbx := netbox.NewClient()
		// cctx := context.Background()

		ctx.Log.Info(("=== PHASE 1.a: getting inventory data ==="), nil)

		// ctx.Log.Info("Getting DNS Domains and Records from Netbox", nil)
		// zones, err := ntbx.ListZones(cctx)
		// if err != nil {
		// 	ctx.Log.Error((fmt.Sprintf("=== ERROR pulling from Netbox ===\n\n%v", err)), nil)
		// 	return err
		// }

		// Create DNS domains, MX, CNAME DKIM records
		ctx.Log.Info(("=== PHASE 2: manage DNS ==="), nil)
		// for _, z := range zones {
		// 	prov := providerFromTags([]string{}) // empty slice
		// }
		if err := dns.ManageDomains(ctx); err != nil {
			ctx.Log.Error((fmt.Sprintf("=== PHASE 2: ERROR ===\n\n%v", err)), nil)
			return err
		}

		// Create the Machines
		ctx.Log.Info(("=== PHASE 3: manage machines ==="), nil)
		if err := machines.ManageMachines(ctx); err != nil {
			ctx.Log.Error((fmt.Sprintf("=== PHASE 3: ERROR ===\n\n%v", err)), nil)
			return err
		}

		ctx.Log.Info(("=== PHASE 4: manage network ==="), nil)
		if err := unet.ManageNetwork(); err != nil {
			ctx.Log.Error((fmt.Sprintf("=== PHASE 4: ERROR ===\n\n%v", err)), nil)
			return (err)
		}

		// write a README to the project
		readmeBytes, err := os.ReadFile("README.md")
		if err != nil {
			return fmt.Errorf("failed to read readme: %w", err)
		}
		ctx.Export("readme", pulumi.String(string(readmeBytes)))

		return nil
	})
}
