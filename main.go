// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ejsdotsh/infrastructure/src/pkg/dns"
	"github.com/ejsdotsh/infrastructure/src/pkg/machines"

	"github.com/ejsdotsh/infrastructure/src/internal/netbox"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Ensure that the required environment variables are set
		ctx.Log.Info(("=== PRE-CHECKS: load ENV vars ==="), nil)
		if err := checkRequiredEnvVars(); err != nil {
			ctx.Log.Error((fmt.Sprintf("=== PHASE 1: ERROR ===\n\n%v", err)), nil)
			return err
		}

		// Initialize the Netbox client (reads NETBOX_URL/TOKEN from env)
		ctx.Log.Info(("=== PHASE 1: Initialize Netbox client ==="), nil)
		ntbx := netbox.NewClient()
		cctx := context.Background()

		ctx.Log.Info("Getting DNS Domains and Records from Netbox", nil)
		zones, err := ntbx.ListZones(cctx)
		if err != nil {
			ctx.Log.Error((fmt.Sprintf("=== ERROR pulling from Netbox ===\n\n%v", err)), nil)
			return err
		}

		// Create DNS domains, MX, CNAME DKIM records
		ctx.Log.Info(("=== PHASE 2: manage DNS ==="), nil)
		for _, z := range zones {
			prov := providerFromTags([]string{}) // empty slice
		}
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

		// write a README to the project
		readmeBytes, err := os.ReadFile("../README.md")
		if err != nil {
			return fmt.Errorf("failed to read readme: %w", err)
		}
		ctx.Export("readme", pulumi.String(string(readmeBytes)))

		return nil
	})
}
