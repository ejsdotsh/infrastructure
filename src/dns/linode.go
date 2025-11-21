// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package dns

import (
	"fmt"
	"strconv"
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

		_domain, err := linode.NewDomain(ctx, resourceName, &linode.DomainArgs{
			Domain:   pulumi.String(domain.Name),
			SoaEmail: pulumi.String(soaEmail),
			Type:     pulumi.String("master"),
		})
		if err != nil {
			return err
		}

		// Get the domain ID for later use; the `ID()` method returns a `pulumi.IDOutput`,
		// rather than a `pulumi.IntOutput`
		domainID := _domain.ID().ApplyT(func(id pulumi.ID) (int, error) {
			i, err := strconv.Atoi(string(id))
			if err != nil {
				return 0, err
			}
			return i, nil
		}).(pulumi.IntOutput)

		// Add MX records for Proton Mail domains
		if domain.Name != "sahala.org" {
			resourceName = fmt.Sprintf("domain-record-%s", strings.ReplaceAll(domain.Name, ".", "-"))
			_, err = linode.NewDomainRecord(ctx, resourceName+"-mx1", &linode.DomainRecordArgs{
				DomainId:   domainID,
				Priority:   pulumi.Int(10),
				RecordType: pulumi.String("MX"),
				Target:     pulumi.String("mail.protonmail.ch"),
			})
			if err != nil {
				return err
			}
			_, err = linode.NewDomainRecord(ctx, resourceName+"-mx2", &linode.DomainRecordArgs{
				DomainId:   domainID,
				Priority:   pulumi.Int(20),
				RecordType: pulumi.String("MX"),
				Target:     pulumi.String("mailsec.protonmail.ch"),
			})
			if err != nil {
				return err
			}
		}

		if domain.Name == "saha.la" {
			_, err := linode.NewDomainRecord(ctx, "domain-record-saha-la-txt-protonmail-verification", &linode.DomainRecordArgs{
				DomainId:   domainID,
				RecordType: pulumi.String("TXT"),
				Target:     pulumi.String("protonmail-verification=667be840d8b900eb75aabe21850897781b4083fc"),
			})
			if err != nil {
				return err
			}
			_, err = linode.NewDomainRecord(ctx, "domain-record-saha-la-txt-spf", &linode.DomainRecordArgs{
				DomainId:   domainID,
				RecordType: pulumi.String("TXT"),
				Target:     pulumi.String("v=spf1 include:_spf.protonmail.ch mx ~all"),
			})
			if err != nil {
				return err
			}
			_, err = linode.NewDomainRecord(ctx, "domain-record-saha-la-txt-dmarc", &linode.DomainRecordArgs{
				DomainId:   domainID,
				Name:       pulumi.String("_dmarc"),
				RecordType: pulumi.String("TXT"),
				Target:     pulumi.String("v=DMARC1; p=none"),
			})
			if err != nil {
				return err
			}
			_, err = linode.NewDomainRecord(ctx, "domain-record-saha-la-cname-dkim1", &linode.DomainRecordArgs{
				DomainId:   domainID,
				Name:       pulumi.String("protonmail._domainkey"),
				RecordType: pulumi.String("CNAME"),
				Target:     pulumi.String("protonmail.domainkey.dwyfiejqmt25ignanwa5si6rxkctauvcmvghmwwrdjj3q2ezokgeq.domains.proton.ch"),
			})
			if err != nil {
				return err
			}
			_, err = linode.NewDomainRecord(ctx, "domain-record-saha-la-cname-dkim2", &linode.DomainRecordArgs{
				DomainId:   domainID,
				Name:       pulumi.String("protonmail2._domainkey"),
				RecordType: pulumi.String("CNAME"),
				Target:     pulumi.String("protonmail2.domainkey.dwyfiejqmt25ignanwa5si6rxkctauvcmvghmwwrdjj3q2ezokgeq.domains.proton.ch"),
			})
			if err != nil {
				return err
			}
			_, err = linode.NewDomainRecord(ctx, "domain-record-saha-la-cname-dkim3", &linode.DomainRecordArgs{
				DomainId:   domainID,
				Name:       pulumi.String("protonmail3._domainkey"),
				RecordType: pulumi.String("CNAME"),
				Target:     pulumi.String("protonmail3.domainkey.dwyfiejqmt25ignanwa5si6rxkctauvcmvghmwwrdjj3q2ezokgeq.domains.proton.ch"),
			})
			if err != nil {
				return err
			}
		}

		if domain.Name == "ejs.wtf" {
			_, err = linode.NewDomainRecord(ctx, "domain-record-ejs-wtf-txt-protonmail-verification", &linode.DomainRecordArgs{
				DomainId:   domainID,
				RecordType: pulumi.String("TXT"),
				Target:     pulumi.String("protonmail-verification=861bb2edab816ab48789a843460f74021b11d175"),
			})
			if err != nil {
				return err
			}
			_, err = linode.NewDomainRecord(ctx, "domain-record-ejs-wtf-txt-spf", &linode.DomainRecordArgs{
				DomainId:   domainID,
				RecordType: pulumi.String("TXT"),
				Target:     pulumi.String("v=spf1 include:_spf.protonmail.ch mx ~all"),
			})
			if err != nil {
				return err
			}
			_, err = linode.NewDomainRecord(ctx, "domain-record-ejs-wtf-txt-dmarc", &linode.DomainRecordArgs{
				DomainId:   domainID,
				Name:       pulumi.String("_dmarc"),
				RecordType: pulumi.String("TXT"),
				Target:     pulumi.String("v=DMARC1; p=none"),
			})
			if err != nil {
				return err
			}
			_, err = linode.NewDomainRecord(ctx, "domain-record-ejs-wtf-cname-dkim1", &linode.DomainRecordArgs{
				DomainId:   domainID,
				Name:       pulumi.String("protonmail._domainkey"),
				RecordType: pulumi.String("CNAME"),
				Target:     pulumi.String("protonmail.domainkey.dijjcu5rapsqd4kwwvqg2n4pweerealedzosuhqrklurblrtgtxpq.domains.proton.ch"),
			})
			if err != nil {
				return err
			}
			_, err = linode.NewDomainRecord(ctx, "domain-record-ejs-wtf-cname-dkim2", &linode.DomainRecordArgs{
				DomainId:   domainID,
				Name:       pulumi.String("protonmail2._domainkey"),
				RecordType: pulumi.String("CNAME"),
				Target:     pulumi.String("protonmail2.domainkey.dijjcu5rapsqd4kwwvqg2n4pweerealedzosuhqrklurblrtgtxpq.domains.proton.ch"),
			})
			if err != nil {
				return err
			}
			_, err = linode.NewDomainRecord(ctx, "domain-record-ejs-wtf-cname-dkim3", &linode.DomainRecordArgs{
				DomainId:   domainID,
				Name:       pulumi.String("protonmail3._domainkey"),
				RecordType: pulumi.String("CNAME"),
				Target:     pulumi.String("protonmail3.domainkey.dijjcu5rapsqd4kwwvqg2n4pweerealedzosuhqrklurblrtgtxpq.domains.proton.ch"),
			})
			if err != nil {
				return err
			}
		}

		if domain.Name == "ejs.sh" {
			_, err = linode.NewDomainRecord(ctx, "domain-record-ejs-sh-txt-protonmail-verification", &linode.DomainRecordArgs{
				DomainId:   domainID,
				RecordType: pulumi.String("TXT"),
				Target:     pulumi.String("protonmail-verification=3321abe05036b281248028d9da2b98d78d65df6e"),
			})
			if err != nil {
				return err
			}
			_, err = linode.NewDomainRecord(ctx, "domain-record-ejs-sh-txt-spf", &linode.DomainRecordArgs{
				DomainId:   domainID,
				RecordType: pulumi.String("TXT"),
				Target:     pulumi.String("v=spf1 include:_spf.protonmail.ch mx ~all"),
			})
			if err != nil {
				return err
			}
			_, err = linode.NewDomainRecord(ctx, "domain-record-ejs-sh-txt-dmarc", &linode.DomainRecordArgs{
				DomainId:   domainID,
				Name:       pulumi.String("_dmarc"),
				RecordType: pulumi.String("TXT"),
				Target:     pulumi.String("v=DMARC1; p=none"),
			})
			if err != nil {
				return err
			}
			_, err = linode.NewDomainRecord(ctx, "domain-record-ejs-sh-cname-dkim1", &linode.DomainRecordArgs{
				DomainId:   domainID,
				Name:       pulumi.String("protonmail._domainkey"),
				RecordType: pulumi.String("CNAME"),
				Target:     pulumi.String("protonmail.domainkey.dg3e6xadvn5cldvphwhfcspfbrainisw7uirf5e6l4d3jyfqkbgaa.domains.proton.ch"),
			})
			if err != nil {
				return err
			}
			_, err = linode.NewDomainRecord(ctx, "domain-record-ejs-sh-cname-dkim2", &linode.DomainRecordArgs{
				DomainId:   domainID,
				Name:       pulumi.String("protonmail2._domainkey"),
				RecordType: pulumi.String("CNAME"),
				Target:     pulumi.String("protonmail2.domainkey.dg3e6xadvn5cldvphwhfcspfbrainisw7uirf5e6l4d3jyfqkbgaa.domains.proton.ch"),
			})
			if err != nil {
				return err
			}
			_, err = linode.NewDomainRecord(ctx, "domain-record-ejs-sh-cname-dkim3", &linode.DomainRecordArgs{
				DomainId:   domainID,
				Name:       pulumi.String("protonmail3._domainkey"),
				RecordType: pulumi.String("CNAME"),
				Target:     pulumi.String("protonmail3.domainkey.dg3e6xadvn5cldvphwhfcspfbrainisw7uirf5e6l4d3jyfqkbgaa.domains.proton.ch"),
			})
			if err != nil {
				return err
			}
		}

		if domain.Name == "octopik3s.io" {
			_, err = linode.NewDomainRecord(ctx, "domain-record-octopik3s-io-txt-protonmail-verification", &linode.DomainRecordArgs{
				DomainId:   domainID,
				RecordType: pulumi.String("TXT"),
				Target:     pulumi.String("protonmail-verification=b869e15412ae9ebda8ded7d2994b16a829d5ee12"),
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
