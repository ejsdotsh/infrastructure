// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

// CUE schema for machine definitions.
// Validates data/machines.yaml.

package data

import "encoding/yaml"

// #LinodeAlerts defines alert thresholds for a Linode instance.
#LinodeAlerts: {
	cpu:           int & >0
	io:            int & >0
	networkIn:     int & >0
	networkOut:    int & >0
	transferQuota: int & >0 & <=100
}

// #Disk defines a disk attached to a Linode instance.
#Disk: {
	suffix:     string & =~"^-"
	label:      string & !=""
	size:       int & >0
	filesystem: "ext4" | "swap" | "raw" | "initrd"
}

// #ConfigHelpers defines config profile helper settings.
#ConfigHelpers: {
	devtmpfsAutomount?: bool
	network?:           bool
	updateDBDisabled?:  bool
	distro?:            bool
	moduleDep?:         bool
}

// #Config defines a config profile for a Linode instance.
#Config: {
	suffix:     string & =~"^-"
	label:      string & !=""
	kernel:     string & !=""
	rootDevice: string & !=""
	booted:     bool
	helpers?:   #ConfigHelpers
	deviceMap: {[string]: string & =~"^-"}
}

// #Machine defines a compute instance specification.
// The provider field determines which cloud to provision on.
#Machine: {
	name:     string & !=""
	provider: "linode" | "digitalocean"
	region:   string & !=""

	// Linode-specific fields (required when provider is "linode").
	type?:            string
	privateIP?:       bool
	diskEncryption?:  string
	alerts?:          #LinodeAlerts
	disks?: [...#Disk]
	config?: #Config

	// DigitalOcean-specific fields (required when provider is "digitalocean").
	size?:  string
	image?: string
	ipv6?:  bool
	tags?: [...string]
}

_machinesRaw: string @tag(machinesFile, type=string)
_machinesData: yaml.Unmarshal(_machinesRaw)
_machinesData: [...#Machine]
