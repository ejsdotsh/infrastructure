// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

// CUE schema for machine definitions.
// Validates data/machines/*.yaml files.

package machines

// #LinodeAlerts defines alert thresholds for a Linode instance.
#LinodeAlerts: {
	cpu:           int & >0
	io:            int & >0
	networkIn:     int & >0
	networkOut:    int & >0
	transferQuota: int & >0 & <=100
}

// #LinodeDisk defines a disk attached to a Linode instance.
#LinodeDisk: {
	suffix:     string & =~"^-"
	label:      string & !=""
	size:       int & >0
	filesystem: "ext4" | "swap" | "raw" | "initrd"
}

// #LinodeConfigHelpers defines config profile helper settings.
#LinodeConfigHelpers: {
	devtmpfsAutomount?: bool
	network?:           bool
	updateDBDisabled?:  bool
	distro?:            bool
	moduleDep?:         bool
}

// #LinodeConfig defines a config profile for a Linode instance.
#LinodeConfig: {
	suffix:     string & =~"^-"
	label:      string & !=""
	kernel:     string & !=""
	rootDevice: string & !=""
	booted:     bool
	helpers?:   #LinodeConfigHelpers
	deviceMap: {[string]: string & =~"^-"}
}

// #LinodeMachine defines a complete Linode instance specification.
#LinodeMachine: {
	name:            string & !=""
	region:          string & !=""
	type:            string & !=""
	privateIP?:      bool
	diskEncryption?: string
	alerts?:         #LinodeAlerts
	disks: [...#LinodeDisk] & [_, ...]
	config: #LinodeConfig
}

// #DODroplet defines a DigitalOcean Droplet specification.
#DODroplet: {
	name:    string & !=""
	region:  string & !=""
	size:    string & !=""
	image:   string & !=""
	ipv6?:   bool
	tags?: [...string]
}
