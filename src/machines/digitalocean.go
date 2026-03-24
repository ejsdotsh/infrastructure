package machines

import (
	"fmt"

	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NewMachineArgs struct {
	DomainName pulumi.StringInput
	IpAddress  pulumi.StringInput
	Region     pulumi.StringInput
	Size       pulumi.StringInput
	Image      pulumi.StringInput
	SshKey     pulumi.StringInput
}

type NewMachine struct {
	pulumi.ResourceState

	Domain  *digitalocean.Domain
	Droplet *digitalocean.Droplet
}

func NewDOMachine(
	ctx *pulumi.Context,
	name string,
	args *NewMachineArgs,
	opts ...pulumi.ResourceOption,
) (*NewMachine, error) {

	comp := &NewMachine{}
	err := ctx.RegisterComponentResource("pkg:index:WebInfrastructure", name, comp, opts...)
	if err != nil {
		return nil, err
	}

	// Define the DigitalOcean domain
	domain, err := digitalocean.NewDomain(ctx, name+"-domain", &digitalocean.DomainArgs{
		Name:      args.DomainName,
		IpAddress: args.IpAddress,
	}, pulumi.Parent(comp))
	if err != nil {
		return nil, err
	}
	comp.Domain = domain

	// Define the DigitalOcean Droplet with both IPv4 and IPv6 enabled
	droplet, err := digitalocean.NewDroplet(ctx, name+"-droplet", &digitalocean.DropletArgs{
		Image:  args.Image,
		Region: args.Region,
		Size:   args.Size,
		Ipv6:   pulumi.Bool(true),
		Tags:   pulumi.StringArray{pulumi.String("web-server")},
		SshKeys: pulumi.StringArray{
			args.SshKey,
		},
	}, pulumi.Parent(comp))
	if err != nil {
		return nil, err
	}
	comp.Droplet = droplet

	// Create an A record for the droplet's IPv4 address
	_, err = digitalocean.NewDnsRecord(ctx, name+"-dns-ipv4", &digitalocean.DnsRecordArgs{
		Domain: domain.Name,
		Name:   pulumi.String("web"),
		Type:   pulumi.String("A"),
		Value:  droplet.Ipv4Address,
	}, pulumi.Parent(comp))
	if err != nil {
		return nil, err
	}

	// Create an AAAA record for the droplet's IPv6 address
	_, err = digitalocean.NewDnsRecord(ctx, name+"-dns-ipv6", &digitalocean.DnsRecordArgs{
		Domain: domain.Name,
		Name:   pulumi.String("web"),
		Type:   pulumi.String("AAAA"),
		Value:  droplet.Ipv6Address,
	}, pulumi.Parent(comp))
	if err != nil {
		return nil, err
	}

	// Export the Droplet's IP addresses
	ctx.Export("ipv4Address", droplet.Ipv4Address)
	ctx.Export("ipv6Address", droplet.Ipv6Address)

	return comp, nil
}

// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := NewDOMachine(ctx, "mywebinfra", &NewMachineArgs{
// 			DomainName: pulumi.String("example.com"),
// 			IpAddress:  pulumi.String("203.0.113.1"),
// 			Region:     pulumi.String("nyc3"),
// 			Size:       pulumi.String("s-1vcpu-1gb"),
// 			Image:      pulumi.String("alpine-3.15.0"),
// 			SshKey:     pulumi.String("your-ssh-key"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }

func manageDOMachines(ctx *pulumi.Context, m Machine) error {
	fmt.Println("i manage machines on DO")
	return nil
}
