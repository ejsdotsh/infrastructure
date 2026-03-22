// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

// Package machines defines the general purpose compute components.
package machines

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// Machine is a simple definition of a compute resource.
type Machine struct {
	Name      pulumi.String   `json:"machineName"`
	Type      pulumi.String   `json:"machineType"` // one of, `vm`, `lxc`, `pod`
	Provider  pulumi.String   `json:"machineProvider"`
	Region    pulumi.String   `json:"machineRegion"`
	Size      pulumi.String   `json:"machineSize"`
	IP4       []pulumi.String `json:"ip4"`
	IP6       []pulumi.String `json:"ip6"`
	PrivateIP pulumi.Bool     `json:"privateIP"`
}

// ManageMachines sets up all virtual machines, containers, and compute resources.
func ManageMachines(ctx *pulumi.Context) error {
	// create a new Pulumi Config
	config := config.New(ctx, "")
	machineData := config.Get("machines")

	var machines []Machine
	if err := json.Unmarshal([]byte(machineData), &machines); err != nil {
		return err
	}

	ctx.Log.Info(fmt.Sprintf("provisioning %d machines\n", len(machines)), nil)
	for _, machine := range machines {
		switch machine.Provider {
		case "linode":
			if err := manageLinodeMachines(ctx, machine); err != nil {
				return err
			}
		case "do":
			// fmt.Printf("manageDOMachines(%v)\n", machine)
			if err := manageDOMachines(ctx, machine); err != nil {
				return err
			}

		}
	}

	// Provision and configure NAS
	// Provision and configure switches

	// Provision and configure octopik3s Kubernetes cluster
	// if err := unicornsland.SetupUnicornsLANd(ctx); err != nil {
	// 	return err
	// }

	return nil

}
