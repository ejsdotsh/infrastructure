// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

// Package loader reads machine and DNS definitions from YAML data files.
package loader

// LinodeAlerts defines alert thresholds for a Linode instance.
type LinodeAlerts struct {
	CPU           int `yaml:"cpu"`
	IO            int `yaml:"io"`
	NetworkIn     int `yaml:"networkIn"`
	NetworkOut    int `yaml:"networkOut"`
	TransferQuota int `yaml:"transferQuota"`
}

// LinodeDisk defines a disk attached to a Linode instance.
type LinodeDisk struct {
	Suffix     string `yaml:"suffix"`
	Label      string `yaml:"label"`
	Size       int    `yaml:"size"`
	Filesystem string `yaml:"filesystem"`
}

// LinodeConfigHelpers defines config profile helper settings.
type LinodeConfigHelpers struct {
	DevtmpfsAutomount *bool `yaml:"devtmpfsAutomount,omitempty"`
	Network           *bool `yaml:"network,omitempty"`
}

// LinodeConfig defines a config profile for a Linode instance.
type LinodeConfig struct {
	Suffix     string            `yaml:"suffix"`
	Label      string            `yaml:"label"`
	Kernel     string            `yaml:"kernel"`
	RootDevice string            `yaml:"rootDevice"`
	Booted     bool              `yaml:"booted"`
	Helpers    LinodeConfigHelpers `yaml:"helpers,omitempty"`
	DeviceMap  map[string]string `yaml:"deviceMap"`
}

// LinodeMachine defines a complete Linode instance specification.
type LinodeMachine struct {
	Name           string        `yaml:"name"`
	Region         string        `yaml:"region"`
	Type           string        `yaml:"type"`
	PrivateIP      bool          `yaml:"privateIP,omitempty"`
	DiskEncryption string        `yaml:"diskEncryption,omitempty"`
	Alerts         *LinodeAlerts `yaml:"alerts,omitempty"`
	Disks          []LinodeDisk  `yaml:"disks"`
	Config         LinodeConfig  `yaml:"config"`
}

// DODroplet defines a DigitalOcean Droplet specification.
type DODroplet struct {
	Name   string   `yaml:"name"`
	Region string   `yaml:"region"`
	Size   string   `yaml:"size"`
	Image  string   `yaml:"image"`
	IPv6   bool     `yaml:"ipv6,omitempty"`
	Tags   []string `yaml:"tags,omitempty"`
}

// MXRecord defines an MX DNS record.
type MXRecord struct {
	Priority int    `yaml:"priority"`
	Target   string `yaml:"target"`
}

// TXTRecord defines a TXT DNS record.
type TXTRecord struct {
	Name   string `yaml:"name,omitempty"`
	Target string `yaml:"target"`
}

// CNAMERecord defines a CNAME DNS record.
type CNAMERecord struct {
	Name   string `yaml:"name"`
	Target string `yaml:"target"`
}

// NSRecord defines an NS DNS record.
type NSRecord struct {
	Target string `yaml:"target"`
}

// ARecord defines an A DNS record.
type ARecord struct {
	Name   string `yaml:"name"`
	Target string `yaml:"target"`
	TTL    int    `yaml:"ttl,omitempty"`
}

// AAAARecord defines an AAAA DNS record.
type AAAARecord struct {
	Name   string `yaml:"name"`
	Target string `yaml:"target"`
	TTL    int    `yaml:"ttl,omitempty"`
}

// LinodeDomain defines a Linode-hosted DNS domain and its records.
type LinodeDomain struct {
	Domain   string        `yaml:"domain"`
	SoaEmail string        `yaml:"soaEmail"`
	MX       []MXRecord    `yaml:"mx,omitempty"`
	TXT      []TXTRecord   `yaml:"txt,omitempty"`
	CNAME    []CNAMERecord `yaml:"cname,omitempty"`
	A        []ARecord     `yaml:"a,omitempty"`
	AAAA     []AAAARecord  `yaml:"aaaa,omitempty"`
}

// DODomain defines a DigitalOcean-hosted DNS domain and its records.
type DODomain struct {
	Domain string        `yaml:"domain"`
	MX     []MXRecord    `yaml:"mx,omitempty"`
	TXT    []TXTRecord   `yaml:"txt,omitempty"`
	CNAME  []CNAMERecord `yaml:"cname,omitempty"`
	NS     []NSRecord    `yaml:"ns,omitempty"`
	A      []ARecord     `yaml:"a,omitempty"`
	AAAA   []AAAARecord  `yaml:"aaaa,omitempty"`
}
