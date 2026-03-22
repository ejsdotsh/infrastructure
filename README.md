<!--
SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>

SPDX-License-Identifier: Apache-2.0
-->

# Reanimating My Personal Infrastructure...as Code

This is the repository for rewriting the 20+ years of ad-hoc scripts and automation used to manage my personal
infrastructure. The overall goals for this project are to:

- Replace (or at least significantly reduce) my *click-ops* (having to 'click' in a graphical/web UI) with a
  model-driven and GitOps-like workflow
- Implement comprehensive observability, monitoring, testing, and alerting
- Learn and improve my knowledge

Following the maxim, *"Make it work; Make it right; Make it fast"*, the first major milestone is a monolithic Pulumi
program, where a single `pulumi up` command ensures that my existing cloud (DNS, web, etc.) and NAS resources are
provisioned and configured. For simplicity, I will initially use Pulumi's ESC and cloud for configuration data and
resource state. The second major milestone will be programmatic management of infrastructure secrets with 1Password and
implementation and integration of Netbox as the source-of-truth for inventory and configuration data.

I decided to use [Pulumi](https://www.pulumi.com/) for provisioning so I wouldn't have to use HCL. As a
configuration language (and API?), I find that HCL is...strongly opinionated... in ways that I would prefer to avoid, and
I think [Nickel Lang](https://nickel-lang.org/) or [CUE](https://cuelang.org) are preferable for defining configuration.
Although both Typescript and Python are more-featureful with Pulumi, for this project, I am using [Go](https://go.dev/)
with some CUE for validation. This is both to reduce the number of programming languages/tool-chains in the build
pipeline, as well as to minimize the number of new things to learn simultaneously.

## My Personal Infrastructure

Essentially, if there is a reliable and programmatic way to manage it, then it is in scope. This includes both my home
network/lab, and the various *aaS providers for DNS, websites, email, and such. The *reanimation* of my infrastructure
will include:

- Replacing my existing webserver and websites
- Implementation and integration of 1Password for secrets management
- Defining models for all managed infrastructure components
- Implementing and integrating Netbox as the source-of-truth
- Rebuilding my NAS/application server
- Monitoring, observability, and dashboards
- Rebuilding [Octopik3s](https://github.com/ejsdotsh/octopik3s)

### Infrastructure Components

My infrastructure *"tech stack"* currently consists of:

- [Digital Ocean](https://www.digitalocean.com/)
- [Linode](https://www.linode.com/)
- [Alpine Linux](https://alpinelinux.org/)
- [Raspberry Pi](https://www.raspberrypi.com/)
- [OPNsense](https://opnsense.org/)
- [Debian](https://www.debian.org/)
- [Netbox](https://netboxlabs.com/)
- TP-Link [Omada Controller](https://github.com/mbentley/docker-omada-controller)
- TP-Link EAP660 HD(US) v1.0
- Juniper EX2200-C12T-2G
