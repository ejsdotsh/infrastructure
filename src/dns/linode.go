// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package dns

import (
	"fmt"
	"strings"

	"github.com/pulumi/pulumi-linode/sdk/v5/go/linode"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var (
	// LinodeDomains is a list of domains to be created in Linode DNS.
	linodeDomains = []Domain{
		{Name: "sahala.org"},
		{Name: "saha.la"},
		{Name: "ejs.sh"},
		{Name: "ejs.wtf"},
		{Name: "octopik3s.io"},
	}
)

func manageLinodeDNS(ctx *pulumi.Context) error {
	for _, domain := range linodeDomains {
		// Replace dots with hyphens for resource naming
		resourceName := fmt.Sprintf("domain-%s", strings.ReplaceAll(domain.Name, ".", "-"))
		soaEmail := fmt.Sprintf("domains@%s", domain.Name)
		if domain.Name == "octopik3s.io" {
			soaEmail = fmt.Sprintf("hostmaster@%s", domain.Name)
		}
		_, err := linode.NewDomain(ctx, resourceName, &linode.DomainArgs{
			Domain:   pulumi.String(domain.Name),
			SoaEmail: pulumi.String(soaEmail),
			Status:   pulumi.String("active"),
			Type:     pulumi.String("master"),
		}, pulumi.Protect(true))
		if err != nil {
			return err
		}
	}

	return nil
}
