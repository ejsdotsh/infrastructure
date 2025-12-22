<!--
SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>

SPDX-License-Identifier: Apache-2.0
-->

# Reanimating Notes

Keeping track of what I'm doing

## Import existing infrastructure as Go

Keeping the following statement in mind:

> Make it work. Make it right. Make it fast.

The imports will focus only on the first two, and includes:

- cloud node for web and shell
- internal and external DNS
- firewall (OPNSense)
- 2x Juniper EX2200
- octopik3s (Kubernetes)
- TP-Link AP (that will be replaced when personal finances allow)
- a NAS that isn't being used as a NAS
- a gaming PC with no monitor

### DNS imports

There are two types of resources for the initial import:

- `Domain`
- `DomainRecord` (on Linode), or `DNSRecord` (on DigitalOcean)

Initial import resource naming uses the form:

- `domain-{{domain-name}}`
- `domain-record-{{domain-name}}-{{recordType}}-{{record-name}}`

```txt
pulumi import linode:index/domain:Domain domain-${domain-name} ${ID}
```

Each domain import results in code similar to:

```go
package main

import (
 "github.com/pulumi/pulumi-linode/sdk/v4/go/linode"
 "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
 pulumi.Run(func(ctx *pulumi.Context) error {
  _, err := linode.NewDomain(ctx, "domain-sahala-org", &linode.DomainArgs{
   Domain:   pulumi.String("sahala.org"),
   SoaEmail: pulumi.String("domains@sahala.org"),
   Status:   pulumi.String("active"),
   Type:     pulumi.String("master"),
  }, pulumi.Protect(true))
  if err != nil {
   return err
  }
  return nil
 })
}
```

Leading to the creation of a simple `Domain` type:

```go
// Domain is a struct representing a DNS domain.
type Domain struct {
  Name string // Domain name
}
```

Allowing the domains to be iterated over:

```go
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

  ...
```

The `domainID` is used for creating `DomainRecords` in Linode, and is the DNS domain associated with the record:

```go
  ...

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

  ...

  return nil
 }
```

A basic `DomainRecord`:

```go
// DomainRecord represents a DNS record.
type DomainRecord struct {
 DomainId int    // the ID of the Domain
 Domain   string // Domain name
 Type     string // Record type (A, AAAA, CNAME, etc.)
 Value    string // Record value
 Name     string // Hostname of the record
 Ttl      int    // Time to live for DNS records
}
```

### Machine imports

A simple `Machine`:

```go
type Machine struct {
 name pulumi.String
 machineType pulumi.String // one of `vm` or `lxc`
}
```

Existing machines to import:

- Linode
- ProxmoxVE
- OPNSense

#### Linode

Install the provider: `go get github.com/pulumi/pulumi-linode/sdk/v5/go/linode`

Import the instance:

```txt
pulumi import linode:index/instance:Instance ${resource-name} ${id}
```

#### ProxmoxVE

Install the provider: `go get github.com/muhlba91/pulumi-proxmoxve/sdk/v7/go/proxmoxve`

- it is much easier to use the default ENV variable names

Issues:

- Certificate error on import

```txt
$ pulumi import proxmoxve:CT/container:Container lxc-pulse pve/101

    ...

     Type                       Name                 Plan       Info
     pulumi:pulumi:Stack        infrastructure-main             1 error
 =   └─ proxmoxve:CT:Container  lxc-pulse            import     2 errors

Diagnostics:
  pulumi:pulumi:Stack (infrastructure-main):
    error: preview failed

  proxmoxve:CT:Container (lxc-pulse):
    error:   sdk-v2/provider2.go:572: sdk.helper_schema: error retrieving container: failed to authenticate HTTP GET request (path: nodes/pve/lxc/101/config) - Reason: failed to authenticate: failed to retrieve authentication response: Post "[secret]/api2/json/access/ticket": tls: failed to verify certificate: x509: certificate signed by unknown authority: provider=proxmox@v7.9.0
    error: Preview failed: refreshing urn:pulumi:main::infrastructure::proxmoxve:CT/container:Container::lxc-pulse: 1 error occurred:
     * error retrieving container: failed to authenticate HTTP GET request (path: nodes/pve/lxc/101/config) - Reason: failed to authenticate: failed to retrieve authentication response: Post "[secret]/api2/json/access/ticket": tls: failed to verify certificate: x509: certificate signed by unknown authority

Resources:
    61 unchanged
```

### Inventory data

Using Pulumi's ESC for secrets and data...loading the keys as objects, and unmarshalling JSON into structs

## References

- DigitalOcean [API docs](https://docs.digitalocean.com/reference/api/digitalocean)
