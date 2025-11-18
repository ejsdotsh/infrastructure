// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package netbox

// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

// 	"github.com/pulumi/pulumi-terraform-provider/sdks/go/netbox/v5/netbox"
// )

// var (
// 	netboxRegions = []string{"uw-west", "uw-east"}
// )

// func SetupNetbox(ctx *pulumi.Context) error {
// 	for _, region := range netboxRegions {
// 		_, err := netbox.NewRegion(ctx, region, &netbox.RegionArgs{
// 			Description: pulumi.String("Region for " + region + " resources"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
