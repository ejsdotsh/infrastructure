// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"

	unicornsdns "github.com/ejsdotsh/infrastructure/src/dns"
	"github.com/ejsdotsh/infrastructure/src/machines"

	// ntbx "github.com/ejsdotsh/infrastructure/src/netbox"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		// Check that required environment variables are set
		ctx.Log.Info(("=== PHASE 1: load ENV vars ==="), nil)
		if err := checkRequiredEnvVars(); err != nil {
			ctx.Log.Error((fmt.Sprintf("=== PHASE 1: ERROR ===\n\n%v", err)), nil)
			return err
		}

		// TODO - the netbox provider import/sdk creation doesn't work as desired.
		// Create infrastructure in Netbox Cloud
		// if err := ntbx.SetupNetbox(ctx); err != nil {
		// 	return err
		// }

		// Create DNS domains, MX, CNAME DKIM records
		ctx.Log.Info(("=== PHASE 2: manage DNS ==="), nil)
		if err := unicornsdns.ManageDomains(ctx); err != nil {
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
