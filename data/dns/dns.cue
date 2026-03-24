// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

// CUE schema for DNS domain definitions.
// Validates data/dns/*.yaml files.

package dns

// #MXRecord defines an MX DNS record.
#MXRecord: {
	priority: int & >0
	target:   string & !=""
}

// #TXTRecord defines a TXT DNS record.
#TXTRecord: {
	name?:  string
	target: string & !=""
}

// #CNAMERecord defines a CNAME DNS record.
#CNAMERecord: {
	name:   string & !=""
	target: string & !=""
}

// #NSRecord defines an NS DNS record.
#NSRecord: {
	target: string & !=""
}

// #ARecord defines an A DNS record.
#ARecord: {
	name:   string & !=""
	target: string & !=""
	ttl?:   int & >0
}

// #AAAARecord defines an AAAA DNS record.
#AAAARecord: {
	name:   string & !=""
	target: string & !=""
	ttl?:   int & >0
}

// #LinodeDomain defines a Linode-hosted DNS domain and its records.
#LinodeDomain: {
	domain:   string & =~"^[a-zA-Z0-9]([a-zA-Z0-9-]*\\.)+[a-zA-Z]{2,}$"
	soaEmail: string & =~"@"
	mx?: [...#MXRecord]
	txt?: [...#TXTRecord]
	cname?: [...#CNAMERecord]
	a?: [...#ARecord]
	aaaa?: [...#AAAARecord]
}

// #DODomain defines a DigitalOcean-hosted DNS domain and its records.
#DODomain: {
	domain: string & =~"^[a-zA-Z0-9]([a-zA-Z0-9-]*\\.)+[a-zA-Z]{2,}$"
	mx?: [...#MXRecord]
	txt?: [...#TXTRecord]
	cname?: [...#CNAMERecord]
	ns?: [...#NSRecord]
	a?: [...#ARecord]
	aaaa?: [...#AAAARecord]
}
