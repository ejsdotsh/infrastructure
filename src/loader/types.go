// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

// Package loader reads machine and DNS definitions from YAML data files.
package loader

// Provider constants.
const (
	ProviderLinode       = "linode"
	ProviderDigitalOcean = "digitalocean"
)

// LinodeAlerts defines alert thresholds for a Linode instance.
type LinodeAlerts struct {
	CPU           int `yaml:"cpu"`
	IO            int `yaml:"io"`
	NetworkIn     int `yaml:"networkIn"`
	NetworkOut    int `yaml:"networkOut"`
	TransferQuota int `yaml:"transferQuota"`
}

// Disk defines a disk attached to a Linode instance.
type Disk struct {
	Suffix     string `yaml:"suffix"`
	Label      string `yaml:"label"`
	Size       int    `yaml:"size"`
	Filesystem string `yaml:"filesystem"`
}

// ConfigHelpers defines config profile helper settings.
type ConfigHelpers struct {
	DevtmpfsAutomount *bool `yaml:"devtmpfsAutomount,omitempty"`
	Network           *bool `yaml:"network,omitempty"`
}

// Config defines a config profile for a Linode instance.
type Config struct {
	Suffix     string            `yaml:"suffix"`
	Label      string            `yaml:"label"`
	Kernel     string            `yaml:"kernel"`
	RootDevice string            `yaml:"rootDevice"`
	Booted     bool              `yaml:"booted"`
	Helpers    ConfigHelpers     `yaml:"helpers,omitempty"`
	DeviceMap  map[string]string `yaml:"deviceMap"`
}

// Machine defines a compute instance specification.
// The Provider field determines which cloud to provision on.
type Machine struct {
	Name     string `yaml:"name"`
	Provider string `yaml:"provider"` // "linode" or "digitalocean"
	Region   string `yaml:"region"`

	// Linode-specific fields.
	Type           string        `yaml:"type,omitempty"`
	PrivateIP      bool          `yaml:"privateIP,omitempty"`
	DiskEncryption string        `yaml:"diskEncryption,omitempty"`
	Alerts         *LinodeAlerts `yaml:"alerts,omitempty"`
	Disks          []Disk        `yaml:"disks,omitempty"`
	MachineConfig  *Config       `yaml:"config,omitempty"`

	// DigitalOcean-specific fields.
	Size  string   `yaml:"size,omitempty"`
	Image string   `yaml:"image,omitempty"`
	IPv6  bool     `yaml:"ipv6,omitempty"`
	Tags  []string `yaml:"tags,omitempty"`
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

// Domain defines a DNS domain and its records.
// The Provider field determines which cloud DNS provider to use.
type Domain struct {
	DomainName string `yaml:"domain"`
	Provider   string `yaml:"provider"` // "linode" or "digitalocean"
	SoaEmail   string `yaml:"soaEmail,omitempty"`

	MX    []MXRecord    `yaml:"mx,omitempty"`
	TXT   []TXTRecord   `yaml:"txt,omitempty"`
	CNAME []CNAMERecord `yaml:"cname,omitempty"`
	NS    []NSRecord    `yaml:"ns,omitempty"`
	A     []ARecord     `yaml:"a,omitempty"`
	AAAA  []AAAARecord  `yaml:"aaaa,omitempty"`
}
