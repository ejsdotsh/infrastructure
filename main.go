// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"

	unet "github.com/ejsdotsh/infrastructure/network"
	"github.com/ejsdotsh/infrastructure/src/dns"
	"github.com/ejsdotsh/infrastructure/src/loader"
	"github.com/ejsdotsh/infrastructure/src/machines"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Load data files.
		ctx.Log.Info("=== Loading data files ===", nil)

		linodeMachines, err := loader.LoadLinodeMachines("data/machines/linode.yaml")
		if err != nil {
			return fmt.Errorf("loading linode machines: %w", err)
		}

		doDroplets, err := loader.LoadDODroplets("data/machines/digitalocean.yaml")
		if err != nil {
			return fmt.Errorf("loading DO droplets: %w", err)
		}

		linodeDomains, err := loader.LoadLinodeDomains("data/dns/linode.yaml")
		if err != nil {
			return fmt.Errorf("loading linode domains: %w", err)
		}

		doDomains, err := loader.LoadDODomains("data/dns/digitalocean.yaml")
		if err != nil {
			return fmt.Errorf("loading DO domains: %w", err)
		}

		ctx.Log.Info(fmt.Sprintf("Loaded %d Linode machines, %d DO droplets, %d Linode domains, %d DO domains",
			len(linodeMachines), len(doDroplets), len(linodeDomains), len(doDomains)), nil)

		// Manage DNS domains and records.
		ctx.Log.Info("=== Managing DNS ===", nil)
		if err := dns.ManageDomains(ctx, linodeDomains, doDomains); err != nil {
			return fmt.Errorf("managing DNS: %w", err)
		}

		// Manage compute resources.
		ctx.Log.Info("=== Managing machines ===", nil)
		if err := machines.ManageMachines(ctx, linodeMachines, doDroplets); err != nil {
			return fmt.Errorf("managing machines: %w", err)
		}

		// Manage network devices.
		ctx.Log.Info("=== Managing network ===", nil)
		if err := unet.ManageNetwork(); err != nil {
			return fmt.Errorf("managing network: %w", err)
		}

		// Export the README.
		readmeBytes, err := os.ReadFile("README.md")
		if err != nil {
			return fmt.Errorf("failed to read readme: %w", err)
		}
		ctx.Export("readme", pulumi.String(string(readmeBytes)))

		return nil
	})
}
