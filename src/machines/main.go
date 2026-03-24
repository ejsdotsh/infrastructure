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

// ManageMachines provisions all compute resources from the loaded data.
func ManageMachines(ctx *pulumi.Context, linodeMachines []loader.LinodeMachine, doDroplets []loader.DODroplet) error {
	// Provision Linode instances.
	for _, machine := range linodeMachines {
		_, err := NewLinodeMachine(ctx, fmt.Sprintf("machine-linode-%s", machine.Name), machine)
		if err != nil {
			return fmt.Errorf("linode machine %s: %w", machine.Name, err)
		}
	}

	// Provision DigitalOcean Droplets.
	for _, droplet := range doDroplets {
		_, err := NewDODroplet(ctx, fmt.Sprintf("machine-do-%s", droplet.Name), droplet)
		if err != nil {
			return fmt.Errorf("do droplet %s: %w", droplet.Name, err)
		}
	}

	return nil
}
