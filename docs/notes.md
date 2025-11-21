<!--
SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>

SPDX-License-Identifier: Apache-2.0
-->

# Reanimating Notes

Keeping track of what I'm doing

## Import existing infrastructure as Go

The command:

```txt
pulumi import linode:index/domain:Domain domain-sahala-org ${ID}
```

Results in the following output:

```txt
Previewing import (main)

View in Browser (Ctrl+O): <>

     Type                    Name                 Plan       
     pulumi:pulumi:Stack     infrastructure-main             
 =   └─ linode:index:Domain  domain-sahala-org    import     

Resources:
    = 1 to import
    27 unchanged

Do you want to perform this import? yes
Importing (main)

View in Browser (Ctrl+O): <>

     Type                    Name                 Status               
     pulumi:pulumi:Stack     infrastructure-main                       
 =   └─ linode:index:Domain  domain-sahala-org    imported (0.19s)     

Outputs:

<>

Resources:
    = 1 imported
    27 unchanged

Duration: 1s

Please copy the following code into your Pulumi application. Not doing so
will cause Pulumi to report that an update will happen on the next update command.

Please note that the imported resources are marked as protected. To destroy them
you will need to remove the `protect` option and run `pulumi update` *before*
the destroy will take effect.

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

## References

- DigitalOcean [API docs](https://docs.digitalocean.com/reference/api/digitalocean)
