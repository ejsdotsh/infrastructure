// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/danslimmon/donothing"
)

func main() {
	pcd := donothing.NewProcedure()
	pcd.Short(`reanimating my personal infrastructure`)
	pcd.Long(`

		reanimating my personal infrastructure...as code

  `)

	pcd.AddStep(func(step *donothing.Step) {
		step.Name("createGitRepo")
		step.Short("Create a new directory and Git repo")
		step.Long(`

			mkdir -p $PROJECT_DIR && cd $PROJECT_DIR && git init .

    `)
	})

	pcd.AddStep(func(step *donothing.Step) {
		step.Name("addGitIgnore")
		step.Short("Add .gitignore")
		step.Long(`

			Add .gitignore

    `)
	})

	pcd.AddStep(func(step *donothing.Step) {
		step.Name("createBuildSSHKey")
		step.Short("Create an SSH key to use for the initial connection")
		step.Long(`

			ssh-keyget -t ed_25519

    `)
	})

	pcd.AddStep(func(step *donothing.Step) {
		step.Name("createNewPulumiProject")
		step.Short("Create new Pulumi project")
		step.Long(`

			Create a new Pulumi project using the 'Go' template

			'pulumi new go --force'

    `)
	})

	pcd.AddStep(func(step *donothing.Step) {
		step.Name("convertRuntime")
		step.Short("Convert the Pulumi runtime")
		step.Long(`

			Edit 'Pulumi.yaml' and convert the runtime to 'yaml'

			'''yaml
			runtime:
			  name: yaml
			  options:
			    compiler: /home/ejs/.cargo/bin/nickel export -f yaml main.ncl
			'''

    `)
	})

	pcd.AddStep(func(step *donothing.Step) {
		step.Name("installProviders")
		step.Short("Install the Pulumi providers")
		step.Long(`

			Install provider(s) for Go (future)

			- '$ go get github.com/pulumi/pulumi-linode/sdk/v4'

    `)
	})

	pcd.AddStep(func(step *donothing.Step) {
		step.Name("addSecretsToESC")
		step.Short("Add secrets and configuration to Pulumi ESC")
		step.Long(`

			Add secrets and configuration to Pulumi IaC and ESC

			- 'pulumi esc init'
			- 'pulumi esc set unicorns.wtf/esc linode:token {{LINODE_API_TOKEN}} --secret'
			- 'pulumi esc set unicorns.wtf/esc digitalocean:token {{DIGITALOCEAN_TOKEN}} --secret'
			- 'pulumi esc set unicorns.wtf/esc environmentVariables "{"DIGITALOCEAN_TOKEN": "${digitalocean:token}"}"'
    `)
	})

	pcd.AddStep(func(step *donothing.Step) {
		step.Name("createNewDODroplet")
		step.Short("Create a new Digital Ocean Droplet")
		step.Long(`

			Create Digital Ocean Droplet, using Caddy as the HTTP/HTTPS server

    `)
	})

	pcd.AddStep(func(step *donothing.Step) {
		step.Name("createNewDODNSRecord")
		step.Short("Create/update DNS records for websites")
		step.Long(`

			Create/update DNS records for websites

    `)
	})

	pcd.AddStep(func(step *donothing.Step) {
		step.Name("createNewDOApp")
		step.Short("Future -- Create a new Digital Ocean App")
		step.Long(`

			Create Digital Ocean 'App' for 'website' using the 'Hugo' template

    `)
	})

	pcd.Execute()
}
