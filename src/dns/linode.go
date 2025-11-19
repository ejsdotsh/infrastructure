// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package dns

import (
	"github.com/pulumi/pulumi-linode/sdk/v5/go/linode"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func manageLinodeDNS(ctx *pulumi.Context) error {
	// Replace dots with hyphens for resource naming
	// resourceName := fmt.Sprintf("domain-%s", strings.ReplaceAll(domainName, ".", "-"))
	_, err := linode.NewDomain(ctx, "domain-sahala-org", &linode.DomainArgs{
		Domain:   pulumi.String("sahala.org"),
		SoaEmail: pulumi.String("domains@sahala.org"),
		Status:   pulumi.String("active"),
		Type:     pulumi.String("master"),
	}, pulumi.Protect(true))
	if err != nil {
		return err
	}

	_, err = linode.NewDomain(ctx, "domain-saha-la", &linode.DomainArgs{
		Domain:   pulumi.String("saha.la"),
		SoaEmail: pulumi.String("domains@saha.la"),
		Status:   pulumi.String("active"),
		Type:     pulumi.String("master"),
	}, pulumi.Protect(true))
	if err != nil {
		return err
	}

	_, err = linode.NewDomain(ctx, "domain-ejs-sh", &linode.DomainArgs{
		Domain:   pulumi.String("ejs.sh"),
		SoaEmail: pulumi.String("domains@ejs.sh"),
		Status:   pulumi.String("active"),
		Type:     pulumi.String("master"),
	}, pulumi.Protect(true))
	if err != nil {
		return err
	}

	_, err = linode.NewDomain(ctx, "domain-ejs-wtf", &linode.DomainArgs{
		Domain:   pulumi.String("ejs.wtf"),
		SoaEmail: pulumi.String("domains@ejs.wtf"),
		Status:   pulumi.String("active"),
		Type:     pulumi.String("master"),
	}, pulumi.Protect(true))
	if err != nil {
		return err
	}

	_, err = linode.NewDomain(ctx, "domain-octopik3s-io", &linode.DomainArgs{
		Domain:   pulumi.String("octopik3s.io"),
		SoaEmail: pulumi.String("hostmaster@octopik3s.io"),
		Status:   pulumi.String("active"),
		Type:     pulumi.String("master"),
	}, pulumi.Protect(true))
	if err != nil {
		return err
	}

	return nil
}
