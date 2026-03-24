package dns

// linodeDomainConfigs defines each Linode-hosted domain and its associated DNS records.
// Resource suffixes must match the original flat resource names to preserve state via aliases.
var linodeDomainConfigs = []LinodeDNSArgs{
	{
		DomainName: "sahala.org",
		SoaEmail:   "domains@sahala.org",
		// sahala.org has no MX, TXT, or CNAME records managed here.
	},
	{
		DomainName: "saha.la",
		SoaEmail:   "domains@saha.la",
		MXRecords: []MXRecord{
			{ResourceSuffix: "-mx1", Priority: 10, Target: "mail.protonmail.ch"},
			{ResourceSuffix: "-mx2", Priority: 20, Target: "mailsec.protonmail.ch"},
		},
		TXTRecords: []TXTRecord{
			{ResourceSuffix: "-txt-protonmail-verification", Target: "protonmail-verification=667be840d8b900eb75aabe21850897781b4083fc"},
			{ResourceSuffix: "-txt-spf", Target: "v=spf1 include:_spf.protonmail.ch mx ~all"},
			{ResourceSuffix: "-txt-dmarc", Name: "_dmarc", Target: "v=DMARC1; p=none"},
		},
		CNAMERecords: []CNAMERecord{
			{ResourceSuffix: "-cname-dkim1", Name: "protonmail._domainkey", Target: "protonmail.domainkey.dwyfiejqmt25ignanwa5si6rxkctauvcmvghmwwrdjj3q2ezokgeq.domains.proton.ch"},
			{ResourceSuffix: "-cname-dkim2", Name: "protonmail2._domainkey", Target: "protonmail2.domainkey.dwyfiejqmt25ignanwa5si6rxkctauvcmvghmwwrdjj3q2ezokgeq.domains.proton.ch"},
			{ResourceSuffix: "-cname-dkim3", Name: "protonmail3._domainkey", Target: "protonmail3.domainkey.dwyfiejqmt25ignanwa5si6rxkctauvcmvghmwwrdjj3q2ezokgeq.domains.proton.ch"},
		},
	},
	{
		DomainName: "ejs.sh",
		SoaEmail:   "domains@ejs.sh",
		MXRecords: []MXRecord{
			{ResourceSuffix: "-mx1", Priority: 10, Target: "mail.protonmail.ch"},
			{ResourceSuffix: "-mx2", Priority: 20, Target: "mailsec.protonmail.ch"},
		},
		TXTRecords: []TXTRecord{
			{ResourceSuffix: "-txt-protonmail-verification", Target: "protonmail-verification=3321abe05036b281248028d9da2b98d78d65df6e"},
			{ResourceSuffix: "-txt-spf", Target: "v=spf1 include:_spf.protonmail.ch mx ~all"},
			{ResourceSuffix: "-txt-dmarc", Name: "_dmarc", Target: "v=DMARC1; p=none"},
		},
		CNAMERecords: []CNAMERecord{
			{ResourceSuffix: "-cname-dkim1", Name: "protonmail._domainkey", Target: "protonmail.domainkey.dg3e6xadvn5cldvphwhfcspfbrainisw7uirf5e6l4d3jyfqkbgaa.domains.proton.ch"},
			{ResourceSuffix: "-cname-dkim2", Name: "protonmail2._domainkey", Target: "protonmail2.domainkey.dg3e6xadvn5cldvphwhfcspfbrainisw7uirf5e6l4d3jyfqkbgaa.domains.proton.ch"},
			{ResourceSuffix: "-cname-dkim3", Name: "protonmail3._domainkey", Target: "protonmail3.domainkey.dg3e6xadvn5cldvphwhfcspfbrainisw7uirf5e6l4d3jyfqkbgaa.domains.proton.ch"},
		},
	},
	{
		DomainName: "ejs.wtf",
		SoaEmail:   "domains@ejs.wtf",
		MXRecords: []MXRecord{
			{ResourceSuffix: "-mx1", Priority: 10, Target: "mail.protonmail.ch"},
			{ResourceSuffix: "-mx2", Priority: 20, Target: "mailsec.protonmail.ch"},
		},
		TXTRecords: []TXTRecord{
			{ResourceSuffix: "-txt-protonmail-verification", Target: "protonmail-verification=861bb2edab816ab48789a843460f74021b11d175"},
			{ResourceSuffix: "-txt-spf", Target: "v=spf1 include:_spf.protonmail.ch mx ~all"},
			{ResourceSuffix: "-txt-dmarc", Name: "_dmarc", Target: "v=DMARC1; p=none"},
		},
		CNAMERecords: []CNAMERecord{
			{ResourceSuffix: "-cname-dkim1", Name: "protonmail._domainkey", Target: "protonmail.domainkey.dijjcu5rapsqd4kwwvqg2n4pweerealedzosuhqrklurblrtgtxpq.domains.proton.ch"},
			{ResourceSuffix: "-cname-dkim2", Name: "protonmail2._domainkey", Target: "protonmail2.domainkey.dijjcu5rapsqd4kwwvqg2n4pweerealedzosuhqrklurblrtgtxpq.domains.proton.ch"},
			{ResourceSuffix: "-cname-dkim3", Name: "protonmail3._domainkey", Target: "protonmail3.domainkey.dijjcu5rapsqd4kwwvqg2n4pweerealedzosuhqrklurblrtgtxpq.domains.proton.ch"},
		},
	},
	{
		DomainName: "octopik3s.io",
		SoaEmail:   "hostmaster@octopik3s.io",
		MXRecords: []MXRecord{
			{ResourceSuffix: "-mx1", Priority: 10, Target: "mail.protonmail.ch"},
			{ResourceSuffix: "-mx2", Priority: 20, Target: "mailsec.protonmail.ch"},
		},
		TXTRecords: []TXTRecord{
			{ResourceSuffix: "-txt-protonmail-verification", Target: "protonmail-verification=b869e15412ae9ebda8ded7d2994b16a829d5ee12"},
		},
	},
}
