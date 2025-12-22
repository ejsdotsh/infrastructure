<!--
SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>

SPDX-License-Identifier: Apache-2.0
-->

# README

## Reanimating my personal infrastructure...as code

This is the public repository for rewriting the 20+ years of ad-hoc scripts and automation used to manage my personal
infrastructure. One of the primary goals of this project is to replace (or at least significantly reduce) *click-ops*
(having to 'click' in a graphical/web UI) with a model-driven and GitOps-like workflow.

I decided to use [Pulumi](https://www.pulumi.com/) for provisioning, primarily so I wouldn't have to use HCL. As a
configuration language, I find that HCL is...strongly opinionated in ways that I would prefer to avoid, and I prefer
[Nickel Lang](https://nickel-lang.org/) or [CUE](https://cuelang.org) for defining configuration. That said, I am using
[Go](https://go.dev/) as the initial language for my Pulumi-driven infrastructure. This is to reduce the number of
languages and tool-chains in the build pipeline and minimize the number of new things to learn simultaneously.

## My personal infrastructure

For the purposes of this project, my personal infrastructure is my home network/lab, and various *aaS providers for DNS,
websites, email, and such.

The *reanimation* of my infrastructure will include:

- Programmatic management and migration of infrastructure secrets from 1Password to ProtonPass
- Defining models for all managed infrastructure components
- Replacing my existing webserver and websites
- Rebuilding my NAS using ProxmoxVE
- Rebuilding [Octopik3s](https://github.com/ejsdotsh/octopik3s)
- Implementing and integrating Netbox as the source-of-truth
- Monitoring, observability, and dashboards

Following the maxim, *"Make it work; Make it right; Make it fast"*, the first step is importing my current DNS, compute,
and storage infrastructure into Pulumi. Once the existing infrastructure is imported (i.e. it works), the next step will
be to refactor and rewrite the code with Pulumi [Component Resources](https://www.pulumi.com/docs/iac/concepts/components/),
to remove duplication and make it work more-correctly, as well as utilizing CUE to validate correctness.

### Infrastructure components

My infrastructure *"tech stack"* currently consists of:

- [Digital Ocean](https://www.digitalocean.com/)
- [Linode](https://www.linode.com/)
- [Proxmox VE](https://www.proxmox.com/en/products/proxmox-virtual-environment/overview)
- [Raspberry Pi](https://www.raspberrypi.com/)
- [OPNsense](https://opnsense.org/)
- [Debian](https://www.debian.org/)
- [Netbox](https://netboxlabs.com/)
- TP-Link [Omada Controller](https://github.com/mbentley/docker-omada-controller)
- TP-Link EAP660 HD(US) v1.0
- Juniper EX2200-C12T-2G
