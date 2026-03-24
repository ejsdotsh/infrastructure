// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

// Package machines defines the general purpose compute components.
package machines

import (
	"fmt"

	"github.com/ejsdotsh/infrastructure/src/loader"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// MachineOutput captures the outputs from a provisioned machine.
type MachineOutput struct {
	Name        string
	Provider    string
	IPv4Address pulumi.StringOutput
	IPv6Address pulumi.StringOutput
}

// ManageMachines provisions all compute resources, dispatching by provider.
// Returns a slice of MachineOutput for use in DNS A record creation and exports.
func ManageMachines(ctx *pulumi.Context, machines []loader.Machine) ([]MachineOutput, error) {
	var outputs []MachineOutput

	for _, machine := range machines {
		switch machine.Provider {
		case loader.ProviderLinode:
			lm, err := NewLinodeMachine(ctx, fmt.Sprintf("machine-linode-%s", machine.Name), machine)
			if err != nil {
				return nil, fmt.Errorf("linode machine %s: %w", machine.Name, err)
			}
			outputs = append(outputs, MachineOutput{
				Name:        machine.Name,
				Provider:    machine.Provider,
				IPv4Address: lm.IPv4.Index(pulumi.Int(0)),
				IPv6Address: pulumi.String("").ToStringOutput(), // Linode IPv6 via separate RDNS; placeholder.
			})
		case loader.ProviderDigitalOcean:
			d, err := NewDODroplet(ctx, fmt.Sprintf("machine-do-%s", machine.Name), machine)
			if err != nil {
				return nil, fmt.Errorf("do droplet %s: %w", machine.Name, err)
			}
			outputs = append(outputs, MachineOutput{
				Name:        machine.Name,
				Provider:    machine.Provider,
				IPv4Address: d.IPv4Address,
				IPv6Address: d.IPv6Address,
			})
		default:
			return nil, fmt.Errorf("unknown machine provider %q for %s", machine.Provider, machine.Name)
		}
	}

	return outputs, nil
}
