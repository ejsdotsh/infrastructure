// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"

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

		allMachines, err := loader.LoadMachines("data/machines.yaml")
		if err != nil {
			return fmt.Errorf("loading machines: %w", err)
		}

		allDomains, err := loader.LoadDomains("data/dns.yaml")
		if err != nil {
			return fmt.Errorf("loading domains: %w", err)
		}

		ctx.Log.Info(fmt.Sprintf("Loaded %d machines, %d domains", len(allMachines), len(allDomains)), nil)

		// Manage DNS domains and records.
		ctx.Log.Info("=== Managing DNS ===", nil)
		if err := dns.ManageDomains(ctx, allDomains); err != nil {
			return fmt.Errorf("managing DNS: %w", err)
		}

		// Manage compute resources.
		ctx.Log.Info("=== Managing machines ===", nil)
		machineOutputs, err := machines.ManageMachines(ctx, allMachines)
		if err != nil {
			return fmt.Errorf("managing machines: %w", err)
		}

		// Export machine IP addresses and build a summary.
		readmeLines := pulumi.StringArray{
			pulumi.String("# Infrastructure\n\n"),
			pulumi.String("## Machines\n\n"),
			pulumi.String("| Name | Provider | IPv4 |\n"),
			pulumi.String("|------|----------|------|\n"),
		}

		for _, mo := range machineOutputs {
			// Export each machine's IPv4 address.
			ctx.Export(fmt.Sprintf("%s-ipv4", mo.Name), mo.IPv4Address)

			// Build a table row for the README.
			row := mo.IPv4Address.ApplyT(func(ip string) string {
				return fmt.Sprintf("| %s | %s | %s |\n", mo.Name, mo.Provider, ip)
			}).(pulumi.StringOutput)
			readmeLines = append(readmeLines, row)
		}

		readmeLines = append(readmeLines,
			pulumi.Sprintf("\n## Domains (%d)\n\n", len(allDomains)),
		)
		for _, d := range allDomains {
			readmeLines = append(readmeLines,
				pulumi.String(fmt.Sprintf("- %s (%s)\n", d.DomainName, d.Provider)),
			)
		}

		// Join all lines into a single readme output.
		readme := pulumi.All(readmeLines).ApplyT(func(args []interface{}) string {
			var result string
			for _, a := range args {
				if lines, ok := a.([]interface{}); ok {
					for _, line := range lines {
						result += line.(string)
					}
				}
			}
			return result
		}).(pulumi.StringOutput)

		ctx.Export("readme", readme)

		// Manage network devices.
		ctx.Log.Info("=== Managing network ===", nil)
		if err := unet.ManageNetwork(); err != nil {
			return fmt.Errorf("managing network: %w", err)
		}

		return nil
	})
}
